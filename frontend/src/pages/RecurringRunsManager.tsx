/*
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as React from 'react';
import BusyButton from '../atoms/BusyButton';
import CustomTable, { Column, Row } from '../components/CustomTable';
import Toolbar, { ToolbarActionConfig } from '../components/Toolbar';
import { ApiJob } from '../apis/job';
import { Apis, JobSortKeys, BaseListRequest, ListJobsRequest } from '../lib/Apis';
import { DialogProps, RoutePage, RouteParams } from '../components/Router';
import { Link } from 'react-router-dom';
import { RouteComponentProps } from 'react-router';
import { SnackbarProps } from '@material-ui/core/Snackbar';
import { commonCss } from '../Css';
import { logger, formatDateString } from '../lib/Utils';

interface RecurringRunListProps extends RouteComponentProps {
  experimentId: string;
  updateDialog: (dialogProps: DialogProps) => void;
  updateSnackbar: (snackbarProps: SnackbarProps) => void;
}

interface RecurringRunListState {
  busyIds: Set<string>;
  orderAscending: boolean;
  pageSize: number;
  pageToken: string;
  runs: ApiJob[];
  selectedIds: string[];
  sortBy: string;
  toolbarActions: ToolbarActionConfig[];
  viewIndex: number;
}

class RecurringRunsManager extends React.Component<RecurringRunListProps, RecurringRunListState> {

  constructor(props: any) {
    super(props);

    this.state = {
      busyIds: new Set(),
      orderAscending: false,
      pageSize: 10,
      pageToken: '',
      runs: [],
      selectedIds: [],
      sortBy: JobSortKeys.CREATED_AT,
      toolbarActions: [],
      viewIndex: 1,
    };
  }

  public render() {
    const { runs, orderAscending, pageSize, selectedIds, sortBy, toolbarActions } = this.state;

    const columns: Column[] = [
      {
        customRenderer: this._nameCustomRenderer.bind(this),
        flex: 2,
        label: 'Run name',
        sortKey: JobSortKeys.NAME,
      },
      { label: 'Created at', flex: 2, sortKey: JobSortKeys.CREATED_AT },
      { customRenderer: this._enabledCustomRenderer.bind(this), label: '', flex: 1 },
    ];

    const rows: Row[] = runs.map(r => {
      return {
        error: r.error,
        id: r.id!,
        otherFields: [
          r.name,
          formatDateString(r.created_at),
          r.enabled,
        ],
      };
    });

    return (<React.Fragment>
      <Toolbar actions={toolbarActions} breadcrumbs={[{ displayName: 'Recurring runs', href: '' }]} />
      <CustomTable columns={columns} rows={rows} orderAscending={orderAscending}
        pageSize={pageSize} selectedIds={selectedIds} disableSelection={true}
        updateSelection={ids => this.setState({ selectedIds: ids })} sortBy={sortBy}
        reload={this._loadRuns.bind(this)} emptyMessage={'No recurring runs found in this experiment.'} />
    </React.Fragment>);
  }

  public async load() {
    await this._loadRuns();
  }

  private async _loadRuns(loadRequest?: BaseListRequest): Promise<string> {
    // Override the current state with incoming request
    const request: ListJobsRequest = Object.assign({
      experimentId: this.props.experimentId,
      orderAscending: this.state.orderAscending,
      pageSize: this.state.pageSize,
      pageToken: this.state.pageToken,
      sortBy: this.state.sortBy,
    }, loadRequest);

    let runs: ApiJob[] = [];
    let nextPageToken = '';
    try {
      const response = await Apis.jobServiceApi.listJobs(
        request.pageToken,
        request.pageSize,
        request.sortBy ? request.sortBy + (request.orderAscending ? ' asc' : ' desc') : '',
      );
      runs = response.jobs || [];
      nextPageToken = response.next_page_token || '';
    } catch (err) {
      // TODO: better error experience here
      logger.error('Could not get list of recurring runs');
    }

    this.setState({
      orderAscending: request.orderAscending!,
      pageSize: request.pageSize!,
      pageToken: request.pageToken!,
      runs,
      sortBy: request.sortBy!,
    });

    return nextPageToken;
  }

  private _nameCustomRenderer(value: string, id: string) {
    return <Link className={commonCss.link}
      to={RoutePage.RECURRING_RUN.replace(':' + RouteParams.runId, id)}>{value}</Link>;
  }

  private async _setEnabledState(id: string, enabled: boolean) {
    try {
      await (enabled ? Apis.jobServiceApi.enableJob(id) : Apis.jobServiceApi.disableJob(id));
    } catch (err) {
      this.props.updateDialog({
        buttons: [{ text: 'Close' }],
        content: 'Error changing enabled state of recurring run',
        title: 'Error',
      });
    }
  }

  private _enabledCustomRenderer(value: boolean | undefined, id: string) {
    const isBusy = this.state.busyIds.has(id);
    return <BusyButton outlined={value} title={value === true ? 'Enabled' : 'Disabled'}
      busy={isBusy} onClick={() => {
        let busyIds = this.state.busyIds;
        busyIds.add(id);
        this.setState({ busyIds }, async () => {
          await this._setEnabledState(id, !value);
          busyIds = this.state.busyIds;
          busyIds.delete(id);
          this.setState({ busyIds });
          await this.load();
        });
      }} />;
  }
}

export default RecurringRunsManager;
