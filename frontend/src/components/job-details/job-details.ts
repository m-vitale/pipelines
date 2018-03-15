import 'iron-icons/iron-icons.html';
import 'paper-progress/paper-progress.html';
import 'paper-tabs/paper-tab.html';
import 'paper-tabs/paper-tabs.html';
import 'polymer/polymer.html';

import * as Apis from '../../lib/apis';
import * as Utils from '../../lib/utils';

import { csvParseRows } from 'd3';

// @ts-ignore
import prettyJson from 'json-pretty-html';
import { customElement, property } from '../../decorators';
import { PageElement } from '../../lib/page_element';
import { parseTemplateOuputPaths } from '../../lib/template_parser';
import { Job, JobStatus } from '../../model/job';
import { PlotMetadata } from '../../model/output_metadata';
import { DataPlotter } from '../data-plotter/data-plotter';

import './job-details.html';

const progressCssColors = {
  completed: '--success-color',
  errored: '--error-color',
  notStarted: '',
  running: '--progress-color',
};

@customElement('job-details')
export class JobDetails extends Polymer.Element implements PageElement {

  @property({ type: Object })
  public job: Job | null = null;

  @property({ type: Number })
  public selectedTab = 0;

  private _pipelineId = -1;
  private _jobId = '';

  public async refresh(_: string, queryParams: { jobId?: string, pipelineId: number }) {
    if (queryParams.jobId !== undefined && queryParams.pipelineId > -1) {
      this._pipelineId = queryParams.pipelineId;
      this._jobId = queryParams.jobId;
      this.job = await Apis.getJob(this._pipelineId, this._jobId);

      const pipeline = await Apis.getPipeline(this._pipelineId);
      const templateYaml = await Apis.getPackageTemplate(pipeline.packageId);

      const baseOutputPath = pipeline
        .parameters
        .filter((p) => p.name === 'output')[0]
        .value.toString();

      const outputPaths = parseTemplateOuputPaths(templateYaml, baseOutputPath, this._jobId);

      // TODO: this is a dummy function to get ouput data for this job
      outputPaths.forEach(async (path) => {
        const fileList = await Apis.listFiles(path);
        const metadataFile = fileList.filter((f) => f.endsWith('metadata.json'))[0];

        if (metadataFile) {
          const metadataJson = await Apis.readFile(metadataFile);
          const metadata = JSON.parse(metadataJson) as PlotMetadata;
          switch (metadata.type) {
            case 'roc':
              this.plotRoc(metadata);
              break;
            case 'confusion_matrix':
              this.plotConfusionMatrix(metadata);
              break;
            default:
          }
        }
      });

      this._colorProgressBar();
    }
  }

  protected async plotRoc(metadata: PlotMetadata) {
    const data = csvParseRows(await Apis.readFile(metadata.source));
    const d = new DataPlotter(this.$.plot as HTMLElement);
    const lineColor = getComputedStyle(this).getPropertyValue('--accent-color');
    await d.plotRocCurve(data, lineColor);
    (this.$.plotTitle as any).innerText = 'ROC curve from file: ' + metadata.source;
  }

  protected async plotConfusionMatrix(metadata: PlotMetadata) {
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
      matrix[i][j] = Number.parseInt(count);
    });

    const d = new DataPlotter(this.$.plot as HTMLElement);
    const startColor = getComputedStyle(this).getPropertyValue('--bg-color');
    const endColor = getComputedStyle(this).getPropertyValue('--accent-color');
    await d.plotConfusionMatrix(matrix, labels, startColor, endColor);
    (this.$.plotTitle as any).innerText = 'Confusion Matrix from file: ' + metadata.source;
  }

  protected _dateToString(date: number) {
    return Utils.dateToString(date);
  }

  protected _getStatusIcon(status: JobStatus) {
    return Utils.jobStatusToIcon(status);
  }

  protected _getRuntime(start: number, end: number, status: JobStatus) {
    if (status === 'Not started') {
      return '-';
    }
    if (end === -1) {
      end = Date.now();
    }
    return Utils.dateDiffToString(end - start);
  }

  private _colorProgressBar() {
    if (!this.job) {
      return;
    }
    let color = '';
    switch (this.job.status) {
      case 'Running':
        color = progressCssColors.running;
        break;
      case 'Succeeded':
        color = progressCssColors.completed;
        break;
      case 'Errored':
        color = progressCssColors.errored;
      default:
        color = progressCssColors.notStarted;
        break;
    }
    (this.$.progress as any).updateStyles({
      '--paper-progress-active-color': `var(${color})`,
    });
  }
}
