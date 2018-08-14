import 'paper-spinner/paper-spinner.html';
import 'polymer/polymer.html';

import * as Apis from '../../lib/apis';
import * as Utils from '../../lib/utils';

import { csvParseRows } from 'd3';
import { customElement, property } from 'polymer-decorators/src/decorators';
import { PlotMetadata, PlotType } from '../../model/output_metadata';
import { TableViewer } from '../table-viewer/table-viewer';
import { drawMatrix } from './confusion-matrix';
import { drawROC } from './roc-plot';

import '../table-viewer/table-viewer';

import './data-plot.html';

@customElement('data-plot')
export class DataPlot extends Polymer.Element {

  @property({ type: Object })
  public plotMetadata: PlotMetadata | null = null;

  @property({ type: String })
  public plotTitle = '';

  @property({ type: Boolean })
  protected _showTensorboardControls = false;

  @property({ type: String })
  protected _podAddress = '';

  @property({ type: Boolean })
  protected _tensorboardBusy = false;

  @property({ type: Boolean })
  protected _renderHtmlApp = false;

  @property({ type: Boolean })
  protected _renderTable = false;

  @property({ type: String })
  protected _staticHtmlSource = '';

  public get iframe(): HTMLIFrameElement {
    const root = this.shadowRoot as ShadowRoot;
    return root.querySelector('#iframe') as HTMLIFrameElement;
  }

  public ready(): void {
    super.ready();
    if (this.plotMetadata) {
      switch (this.plotMetadata.type) {
        case PlotType.CONFUSION_MATRIX:
          this._plotConfusionMatrix(this.plotMetadata);
          break;
        case PlotType.ROC:
          this._plotRocCurve(this.plotMetadata);
          break;
        case PlotType.TABLE:
          this._plotTable(this.plotMetadata);
          break;
        case PlotType.TENSORBOARD:
          this._addTensorboardControls();
          break;
        case PlotType.WEB_APP:
          this._renderStaticWebApp(this.plotMetadata);
          break;
        default:
          Utils.log.error('Unknown plotType:', this.plotMetadata.type);
      }
    }
  }

  protected async _startTensorboard(): Promise<void> {
    if (!this.plotMetadata) {
      return;
    }
    this._tensorboardBusy = true;
    try {
      await Apis.startTensorboardApp(encodeURIComponent(this.plotMetadata.source));
    } finally {
      this._tensorboardBusy = false;
    }
    this._addTensorboardControls();
  }

  private async _plotConfusionMatrix(metadata: PlotMetadata): Promise<void> {
    if (!metadata.source) {
      throw new Error('Malformed metadata, property "source" is required.');
    }
    if (!metadata.labels) {
      throw new Error('Malformed metadata, property "labels" is required.');
    }
    if (!metadata.schema) {
      throw new Error('Malformed metadata, property "schema" missing.');
    }
    const data = csvParseRows(await Apis.readFile(metadata.source));
    const labels = metadata.labels;
    const labelIndex: { [label: string]: number } = {};
    let index = 0;
    labels.forEach((l) => {
      labelIndex[l] = index++;
    });

    const matrix = Array.from(Array(labels.length), () => new Array(labels.length));
    data.forEach(([target, predicted, count]) => {
      const i = labelIndex[target];
      const j = labelIndex[predicted];
      matrix[i][j] = Number.parseInt(count, 10);
    });

    const axisLabels = metadata.schema.map((r) => r.name);

    this.plotTitle = 'Confusion Matrix from file: ' + metadata.source;

    // Render the confusion matrix
    drawMatrix({
      container: this.$.container as HTMLElement,
      data: matrix,
      endColor: getComputedStyle(this).getPropertyValue('--theme-primary-color'),
      labels,
      startColor: getComputedStyle(this).getPropertyValue('--background-contrast-color'),
      xAxisLabel: axisLabels[0],
      yAxisLabel: axisLabels[1],
    });
  }

  private async _plotRocCurve(metadata: PlotMetadata): Promise<void> {
    this.plotTitle = 'ROC curve from file: ' + metadata.source;

    // Render the ROC plot
    drawROC({
      container: this.$.container as HTMLElement,
      data: csvParseRows(await Apis.readFile(metadata.source)),
      height: 450,
      lineColor: getComputedStyle(this).getPropertyValue('--theme-primary-color'),
      margin: 50,
      width: 650
    });
  }

  private async _plotTable(metadata: PlotMetadata): Promise<void> {
    if (!metadata.source) {
      throw new Error('Malformed metadata, property "source" is required.');
    }
    if (!metadata.header) {
      throw new Error('Malformed metadata, property "header" is required.');
    }
    if (!metadata.format) {
      throw new Error('Malformed metadata, property "format" is required.');
    }
    this._renderTable = true;

    this.plotTitle = 'Table viewer from file: ' + metadata.source;
    switch (metadata.format) {
      case 'csv':
        const data = csvParseRows(await Apis.readFile(metadata.source));
        const tableViewer = this.shadowRoot!.querySelector('table-viewer') as TableViewer;
        tableViewer.header = metadata.header;
        tableViewer.rows = data;
        break;
      default:
        throw new Error('Unsupported table format: ' + metadata.format);
    }
  }

  private async _renderStaticWebApp(metadata: PlotMetadata): Promise<void> {
    this.plotTitle = 'HTML from file: ' + metadata.source;
    this._renderHtmlApp = true;
    const htmlContent = await Apis.readFile(metadata.source);
    // TODO: iframe.srcdoc doesn't work on Edge yet. It's been added, but not
    // yet rolled out as of the time of writing this (6/14/18):
    // https://developer.microsoft.com/en-us/microsoft-edge/platform/issues/12375527/
    // I'm using this since it seems like the safest way to insert HTML into an
    // iframe, while allowing Javascript, but without needing to require
    // "allow-same-origin" sandbox rule.
    (this.iframe as any).srcdoc = htmlContent;
  }

  private async _addTensorboardControls(): Promise<void> {
    if (!this.plotMetadata) {
      return;
    }
    this._podAddress = await Apis.getTensorboardApp(this.plotMetadata.source);
    this._showTensorboardControls = true;
    this.plotTitle = 'Tensorboard for logdir: ' + this.plotMetadata.source;
  }
}
