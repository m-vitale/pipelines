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
import AddIcon from '@material-ui/icons/Add';
import CustomTable, { Column, Row } from '../components/CustomTable';
import DeleteIcon from '@material-ui/icons/Delete';
import RefreshIcon from '@material-ui/icons/Refresh';
import UploadIcon from '@material-ui/icons/CloudUpload';
import UploadPipelineDialog from '../components/UploadPipelineDialog';
import { ApiPipeline, ApiListPipelinesResponse } from '../apis/pipeline';
import { Apis, PipelineSortKeys, BaseListRequest, ListPipelinesRequest } from '../lib/Apis';
import { Link } from 'react-router-dom';
import { Page } from './Page';
import { RoutePage, RouteParams } from '../components/Router';
import { URLParser, QUERY_PARAMS } from '../lib/URLParser';
import { classes } from 'typestyle';
import { commonCss, padding } from '../Css';
import { logger } from '../lib/Utils';

interface PipelineListState {
  orderAscending: boolean;
  pageSize: number;
  pageToken: string;
  pipelines: ApiPipeline[];
  selectedIds: string[];
  sortBy: string;
  uploadDialogOpen: boolean;
}

class PipelineList extends Page<{}, PipelineListState> {

  constructor(props: any) {
    super(props);

    this.state = {
      orderAscending: true,
      pageSize: 10,
      pageToken: '',
      pipelines: [],
      selectedIds: [],
      sortBy: PipelineSortKeys.NAME,
      uploadDialogOpen: false,
    };
  }

  public getInitialToolbarState() {
    return {
      actions: [
        {
          action: () => this.setState({ uploadDialogOpen: true }),
          disabled: false,
          icon: UploadIcon,
          id: 'uploadBtn',
          title: 'Upload pipeline',
          tooltip: 'Upload pipeline',
        },
        {
          action: this._createJob.bind(this),
          disabled: true,
          disabledTitle: 'Select a pipeline to create a job',
          icon: AddIcon,
          id: 'createJobBtn',
          title: 'Create job',
          tooltip: 'Create job',
        },
        {
          action: () => this.load(),
          disabled: false,
          icon: RefreshIcon,
          id: 'refreshBtn',
          title: 'Refresh',
          tooltip: 'Refresh',
        },
        {
          action: () => this.props.updateDialog({
            buttons: [
              { onClick: () => this._deleteDialogClosed(true), text: 'Delete' },
              { onClick: () => this._deleteDialogClosed(false), text: 'Cancel' },
            ],
            onClose: () => this._deleteDialogClosed(false),
            title: `Delete ${this.state.selectedIds.length} Pipeline${this.state.selectedIds.length === 1 ? '' : 's'}?`,
          }),
          disabled: true,
          disabledTitle: 'Select at least one pipeline to delete',
          icon: DeleteIcon,
          id: 'deleteBtn',
          title: 'Delete',
          tooltip: 'Delete',
        },
      ],
      breadcrumbs: [{ displayName: 'Pipelines', href: RoutePage.PIPELINES }],
    };
  }

  public render(): JSX.Element {
    const columns: Column[] = [
      {
        customRenderer: this._nameCustomRenderer.bind(this),
        flex: 1,
        label: 'Pipeline name',
        sortKey: PipelineSortKeys.NAME,
      },
      { label: 'Description', flex: 3 },
      { label: 'Uploaded on', sortKey: PipelineSortKeys.CREATED_AT, flex: 1 },
    ];

    const rows: Row[] = this.state.pipelines.map((p) => {
      return {
        id: p.id!,
        otherFields: [p.name!, p.description!, p.created_at!.toLocaleString()],
      };
    });

    return (
      <div className={classes(commonCss.page, padding(20, 'lr'))}>
        <CustomTable columns={columns} rows={rows} orderAscending={this.state.orderAscending}
          pageSize={this.state.pageSize} sortBy={this.state.sortBy}
          updateSelection={this._selectionChanged.bind(this)} selectedIds={this.state.selectedIds}
          reload={this._reload.bind(this)}
          emptyMessage='No pipelines found. Click "Upload pipeline" to start.' />

        <UploadPipelineDialog open={this.state.uploadDialogOpen}
          onClose={this._uploadDialogClosed.bind(this)} />
      </div>
    );
  }

  public async load(): Promise<void> {
    await this._reload();
  }

  private _showErrorDialog(title: string, content: string): void {
    this.props.updateDialog({
      buttons: [{ text: 'Dismiss' }],
      content,
      title,
    });
  }

  private async _reload(loadRequest?: BaseListRequest): Promise<string> {
    // Override the current state with incoming request
    const request: ListPipelinesRequest = Object.assign({
      orderAscending: this.state.orderAscending,
      pageSize: this.state.pageSize,
      pageToken: this.state.pageToken,
      sortBy: this.state.sortBy,
    }, loadRequest);


    let response: ApiListPipelinesResponse;
    try {
      response = await Apis.pipelineServiceApi.listPipelines(
        request.pageToken,
        request.pageSize,
        request.sortBy ? request.sortBy + (request.orderAscending ? ' asc' : ' desc') : ''
      );
    } catch (err) {
      this.props.updateBanner({
        additionalInfo: err.message,
        message: 'Error: failed to retrieve list of pipelines.'
          + (err.message ? ' Click Details for more information.' : ''),
        mode: 'error',
        refresh: this.load.bind(this),
      });
      // No point in continuing if we couldn't retrieve the pipelines.
      return '';
    }

    this.setState({
      orderAscending: request.orderAscending!,
      pageSize: request.pageSize!,
      pageToken: request.pageToken!,
      pipelines: response.pipelines || [],
      sortBy: request.sortBy!,
    });

    return response.next_page_token || '';
  }

  private _nameCustomRenderer(value: string, id: string): React.ReactElement<Link> {
    return (
      <Link onClick={(e) => e.stopPropagation()}
        className={commonCss.link}
        to={RoutePage.PIPELINE_DETAILS.replace(':' + RouteParams.pipelineId, id)}>{value}
      </Link>
    );
  }

  private _selectionChanged(selectedIds: string[]): void {
    const toolbarActions = [...this.props.toolbarProps.actions];
    toolbarActions[1].disabled = selectedIds.length !== 1;
    toolbarActions[3].disabled = !selectedIds.length;
    this.props.updateToolbar({ breadcrumbs: this.props.toolbarProps.breadcrumbs, actions: toolbarActions });
    this.setState({ selectedIds });
  }

  private async _deleteDialogClosed(deleteConfirmed: boolean): Promise<void> {
    if (deleteConfirmed) {
      const unsuccessfulDeleteIds: string[] = [];
      const errorMessages: string[] = [];
      // TODO: Show spinner during wait.
      await Promise.all(this.state.selectedIds.map(async (id) => {
        try {
          await Apis.pipelineServiceApi.deletePipeline(id);
        } catch (err) {
          unsuccessfulDeleteIds.push(id);
          const pipeline = this.state.pipelines.find((p) => p.id === id);
          errorMessages.push(
            `Deleting pipeline${pipeline ? ': ' + pipeline.name : ''} failed with error: "${err}"`);
        }
      }));

      const successfulDeletes = this.state.selectedIds.length - unsuccessfulDeleteIds.length;
      if (successfulDeletes > 0) {
        this.props.updateSnackbar({
          message: `Successfully deleted ${successfulDeletes} pipeline${successfulDeletes === 1 ? '' : 's'}!`,
          open: true,
        });
        this.load();
      }

      if (unsuccessfulDeleteIds.length > 0) {
        this._showErrorDialog(
          `Failed to delete ${unsuccessfulDeleteIds.length} pipeline${unsuccessfulDeleteIds.length === 1 ? '' : 's'}`,
          errorMessages.join('\n\n'));
      }

      this._selectionChanged(unsuccessfulDeleteIds);
    }
  }

  private _createJob() {
    if (this.state.selectedIds.length === 1) {
      const pipelineId = this.state.selectedIds[0];
      const search = new URLParser(this.props).build({ [QUERY_PARAMS.pipelineId]: pipelineId });
      this.props.history.push(RoutePage.NEW_JOB + search);
    }
  }

  private async _uploadDialogClosed(name: string, file: File | null): Promise<boolean> {
    if (!!file) {
      try {
        await Apis.uploadPipeline(name, file);
        this.setState({ uploadDialogOpen: false });
        this.load();
        return true;
      } catch (err) {
        this._showErrorDialog('Failed to upload pipeline', err.message);
        logger.error('Error uploading pipeline:', err);
        return false;
      }
    } else {
      this.setState({ uploadDialogOpen: false });
      return false;
    }
  }
}

// tslint:disable-next-line:no-default-export
export default PipelineList;
