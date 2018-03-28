import { customElement, property } from 'polymer-decorators/src/decorators';
import 'polymer/polymer.html';

import * as Apis from '../../lib/apis';
import * as Utils from '../../lib/utils';

import { NodePhase } from '../../model/argo_template';
import { ItemClickEvent, RouteEvent } from '../../model/events';

import { ColumnTypeName, ItemListColumn, ItemListElement, ItemListRow } from '../item-list/item-list';
import './job-list.html';

@customElement('job-list')
export class JobList extends Polymer.Element {

  @property({ type: Array })
  public jobs: Apis.JobMetadata[] = [];

  protected jobListRows: ItemListRow[] = [];

  protected jobListColumns: ItemListColumn[] = [
    { name: 'Job Name', type: ColumnTypeName.STRING },
    { name: 'Scheduled Time', type: ColumnTypeName.DATE },
  ];

  private _pipelineId = -1;

  ready() {
    super.ready();
    const itemList = this.$.jobsItemList as ItemListElement;
    itemList.addEventListener('itemDoubleClick', this._navigate.bind(this));
  }

  // TODO: should these jobs be cached?
  public async loadJobs(pipelineId: number) {
    this._pipelineId = pipelineId;
    this.jobs = await Apis.getJobs(this._pipelineId);

    this.jobListRows = this.jobs.map((job) => {
      const row = new ItemListRow({
        columns: [
          job.name,
          job.scheduledAt,
        ],
        selected: false,
      });
      return row;
    });
  }

  protected _navigate(ev: ItemClickEvent) {
    const jobId = this.jobs[ev.detail.index].name;
    this.dispatchEvent(
      new RouteEvent(`/pipelineJob?pipelineId=${this._pipelineId}&jobId=${jobId}`));
  }

  protected _paramsToArray(paramsObject: {}) {
    return Utils.objectToArray(paramsObject);
  }

  protected _getStatusIcon(status: NodePhase) {
    return Utils.nodePhaseToIcon(status);
  }

  protected _getRuntime(start: string, end: string, status: NodePhase) {
    if (!status) {
      return '-';
    }
    const startDate = new Date(start);
    const endDate = end ? new Date(end) : new Date();
    return Utils.dateDiffToString(endDate.valueOf() - startDate.valueOf());
  }
}
