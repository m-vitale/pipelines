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
import Buttons from '../lib/Buttons';
import CustomTable, { Column, Row, ExpandState } from '../components/CustomTable';
import RunList from './RunList';
import produce from 'immer';
import { ApiFilter, PredicateOp } from '../apis/filter';
import { ApiListExperimentsResponse, ApiExperiment } from '../apis/experiment';
import { ApiResourceType, ApiRun, RunStorageState } from '../apis/run';
import { Apis, ExperimentSortKeys, ListRequest, RunSortKeys } from '../lib/Apis';
import { Link } from 'react-router-dom';
import { Page } from './Page';
import { RoutePage, RouteParams, QUERY_PARAMS } from '../components/Router';
import { ToolbarProps } from '../components/Toolbar';
import { URLParser } from '../lib/URLParser';
import { classes } from 'typestyle';
import { commonCss, padding } from '../Css';
import { logger, s, errorToMessage } from '../lib/Utils';
import { statusToIcon, NodePhase } from './Status';

interface DisplayExperiment extends ApiExperiment {
  last5Runs?: ApiRun[];
  error?: string;
  expandState?: ExpandState;
}

interface ExperimentListState {
  displayExperiments: DisplayExperiment[];
  selectedRunIds: string[];
  selectedTab: number;
}

class ExperimentList extends Page<{}, ExperimentListState> {
  private _tableRef = React.createRef<CustomTable>();

  constructor(props: any) {
    super(props);

    this.state = {
      displayExperiments: [],
      selectedRunIds: [],
      selectedTab: 0,
    };
  }

  public getInitialToolbarState(): ToolbarProps {
    return {
      actions: [
        Buttons.newExperiment(this._newExperimentClicked.bind(this)),
        Buttons.compareRuns(this._compareRuns.bind(this)),
        Buttons.cloneRun(this._cloneRun.bind(this)),
        Buttons.archive(() => this.props.updateDialog({
          buttons: [
            { onClick: async () => await this._archiveDialogClosed(true), text: 'Archive' },
            { onClick: async () => await this._archiveDialogClosed(false), text: 'Cancel' },
          ],
          onClose: async () => await this._archiveDialogClosed(false),
          title: `Archive ${this.state.selectedRunIds.length} run${s(this.state.selectedRunIds.length)}?`,
        })),
        Buttons.refresh(this.refresh.bind(this)),
      ],
      breadcrumbs: [],
      pageTitle: 'Experiments',
    };
  }

  public render(): JSX.Element {
    const columns: Column[] = [{
      customRenderer: this._nameCustomRenderer.bind(this),
      flex: 1,
      label: 'Experiment name',
      sortKey: ExperimentSortKeys.NAME,
    }, {
      flex: 2,
      label: 'Description',
    }, {
      customRenderer: this._last5RunsCustomRenderer.bind(this),
      flex: 1,
      label: 'Last 5 runs',
    }];

    const rows: Row[] = this.state.displayExperiments.map((exp) => {
      return {
        error: exp.error,
        expandState: exp.expandState,
        id: exp.id!,
        otherFields: [
          exp.name!,
          exp.description!,
          exp.expandState === ExpandState.EXPANDED ? [] : exp.last5Runs,
        ]
      };
    });

    return (
      <div className={classes(commonCss.page, padding(20, 'lr'))}>
        <CustomTable columns={columns} rows={rows} ref={this._tableRef}
          disableSelection={true} initialSortColumn={ExperimentSortKeys.CREATED_AT}
          reload={this._reload.bind(this)} toggleExpansion={this._toggleRowExpand.bind(this)}
          getExpandComponent={this._getExpandedExperimentComponent.bind(this)}
          filterLabel='Filter experiments'
          emptyMessage='No experiments found. Click "Create experiment" to start.' />
      </div>
    );
  }

  public async refresh(): Promise<void> {
    if (this._tableRef.current) {
      this.clearBanner();
      await this._tableRef.current.reload();
    }
  }

  private async _reload(request: ListRequest): Promise<string> {
    // Fetch the list of experiments
    let response: ApiListExperimentsResponse;
    let displayExperiments: DisplayExperiment[];
    try {
      response = await Apis.experimentServiceApi.listExperiment(
        request.pageToken, request.pageSize, request.sortBy, request.filter);
      displayExperiments = response.experiments || [];
      displayExperiments.forEach((exp) => exp.expandState = ExpandState.COLLAPSED);
    } catch (err) {
      await this.showPageError('Error: failed to retrieve list of experiments.', err);
      // No point in continuing if we couldn't retrieve any experiments.
      return '';
    }

    // Fetch and set last 5 runs' statuses for each experiment
    await Promise.all(displayExperiments.map(async experiment => {
      // TODO: should we aggregate errors here? What if they fail for different reasons?
      try {
        const listRunsResponse = await Apis.runServiceApi.listRuns(
          undefined /* pageToken */,
          5 /* pageSize */,
          RunSortKeys.CREATED_AT + ' desc',
          ApiResourceType.EXPERIMENT.toString(),
          experiment.id,
          encodeURIComponent(JSON.stringify({
            predicates: [{
              key: 'storageState',
              op: PredicateOp.EQUALS,
              string_value: RunStorageState.ARCHIVED.toString(),
            }]
          } as ApiFilter)),
        );
        experiment.last5Runs = listRunsResponse.runs || [];
      } catch (err) {
        experiment.error = 'Failed to load the last 5 runs of this experiment';
        logger.error(
          `Error: failed to retrieve run statuses for experiment: ${experiment.name}.`,
          err);
      }
    }));

    this.setState({ displayExperiments });
    return response.next_page_token || '';
  }

  private _cloneRun(): void {
    if (this.state.selectedRunIds.length === 1) {
      const searchString = new URLParser(this.props).build({
        [QUERY_PARAMS.cloneFromRun]: this.state.selectedRunIds[0] || ''
      });
      this.props.history.push(RoutePage.NEW_RUN + searchString);
    }
  }

  private _nameCustomRenderer(value: string, id: string): JSX.Element {
    return <Link className={commonCss.link} onClick={(e) => e.stopPropagation()}
      to={RoutePage.EXPERIMENT_DETAILS.replace(':' + RouteParams.experimentId, id)}>{value}</Link>;
  }

  private _last5RunsCustomRenderer(runs: ApiRun[]): JSX.Element {
    return <div className={commonCss.flex}>
      {(runs || []).map((run, i) => (
        <span key={i} style={{ margin: '0 1px' }}>
          {statusToIcon(run.status as NodePhase || NodePhase.UNKNOWN, run.created_at)}
        </span>
      ))}
    </div>;
  }

  private _runSelectionChanged(selectedRunIds: string[]): void {
    const actions = produce(this.props.toolbarProps.actions, draft => {
      // Enable/Disable Run compare button
      draft[1].disabled = selectedRunIds.length <= 1 || selectedRunIds.length > 10;
      // Enable/Disable Clone button
      draft[2].disabled = selectedRunIds.length !== 1;
    });
    this.props.updateToolbar({ actions });
    this.setState({ selectedRunIds });
  }

  private _compareRuns(): void {
    const indices = this.state.selectedRunIds;
    if (indices.length > 1 && indices.length <= 10) {
      const runIds = this.state.selectedRunIds.join(',');
      const searchString = new URLParser(this.props).build({
        [QUERY_PARAMS.runlist]: runIds,
      });
      this.props.history.push(RoutePage.COMPARE + searchString);
    }
  }

  private _newExperimentClicked(): void {
    this.props.history.push(RoutePage.NEW_EXPERIMENT);
  }

  private _toggleRowExpand(rowIndex: number): void {
    const displayExperiments = produce(this.state.displayExperiments, draft => {
      draft[rowIndex].expandState =
        draft[rowIndex].expandState === ExpandState.COLLAPSED ?
          ExpandState.EXPANDED :
          ExpandState.COLLAPSED;
    });

    this.setState({ displayExperiments });
  }

  private _getExpandedExperimentComponent(experimentIndex: number): JSX.Element {
    const experiment = this.state.displayExperiments[experimentIndex];
    const runIds = (experiment.last5Runs || []).map((r) => r.id!);
    return <RunList runIdListMask={runIds} onError={() => null} {...this.props}
      disablePaging={true} selectedIds={this.state.selectedRunIds} noFilterBox={true}
      storageState={RunStorageState.AVAILABLE}
      onSelectionChange={this._runSelectionChanged.bind(this)} disableSorting={true} />;
  }

  private async _archiveDialogClosed(confirmed: boolean): Promise<void> {
    if (confirmed) {
      const unsuccessfulIds: string[] = [];
      const errorMessages: string[] = [];
      await Promise.all(this.state.selectedRunIds.map(async (id) => {
        try {
          await Apis.runServiceApi.archiveRun(id);
        } catch (err) {
          unsuccessfulIds.push(id);
          const errorMessage = await errorToMessage(err);
          errorMessages.push(`Deleting run failed with error: "${errorMessage}"`);
        }
      }));

      const successfulObjects = this.state.selectedRunIds.length - unsuccessfulIds.length;
      if (successfulObjects > 0) {
        this.props.updateSnackbar({
          message: `Successfully archived ${successfulObjects} run${s(successfulObjects)}!`,
          open: true,
        });
        this.refresh();
      }

      if (unsuccessfulIds.length > 0) {
        this.showErrorDialog(
          `Failed to archive ${unsuccessfulIds.length} run${s(unsuccessfulIds)}`,
          errorMessages.join('\n\n'));
      }

      this._runSelectionChanged(unsuccessfulIds);
    }
  }

}

export default ExperimentList;
