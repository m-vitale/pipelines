/*
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React, { Component } from 'react';
import { Page } from './Page';
import { ToolbarProps } from '../components/Toolbar';
import { RoutePage, RouteParams } from '../components/Router';
import { classes, stylesheet } from 'typestyle';
import { commonCss, padding } from '../Css';
import { CircularProgress } from '@material-ui/core';
import { titleCase, getResourceProperty, serviceErrorToString } from '../lib/Utils';
import { ResourceInfo } from '../components/ResourceInfo';
import { Execution } from '../generated/src/apis/metadata/metadata_store_pb';
import { Apis, ExecutionProperties, ArtifactProperties } from '../lib/Apis';
import { GetExecutionsByIDRequest, GetEventsByExecutionIDsRequest, GetEventsByExecutionIDsResponse, GetArtifactsByIDRequest } from '../generated/src/apis/metadata/metadata_store_service_pb';
import { EventTypes } from 'src/lib/MetadataUtils';
import { Event } from '../generated/src/apis/metadata/metadata_store_pb';

type ArtifactIdList = number[];

interface ExecutionDetailsState {
  execution?: Execution;
  events?: Record<EventTypes, ArtifactIdList>;
}

export default class ExecutionDetails extends Page<{}, ExecutionDetailsState> {

  constructor(props: {}) {
    super(props);
    this.state = {};
    this.load = this.load.bind(this);
  }

  private get fullTypeName(): string {
    return this.props.match.params[RouteParams.EXECUTION_TYPE] || '';
  }

  private get properTypeName(): string {
    const parts = this.fullTypeName.split('/');
    if (!parts.length) {
      return '';
    }

    return titleCase(parts[parts.length - 1]);
  }

  private get id(): string {
    return this.props.match.params[RouteParams.ID];
  }

  public async componentDidMount(): Promise<void> {
    return this.load();
  }

  public render(): JSX.Element {
    if (!this.state.execution || !this.state.events) {
      return <CircularProgress />;
    }

    return (
      <div className={classes(commonCss.page, padding(20, 'lr'))}>
        {<ResourceInfo typeName={this.properTypeName}
          resource={this.state.execution} />}
        <SectionIO title={'Declared Inputs'} artifactIds={this.state.events[Event.Type.DECLARED_INPUT]} />
        <SectionIO title={'Inputs'} artifactIds={this.state.events[Event.Type.INPUT]} />
        <SectionIO title={'Declared Outputs'} artifactIds={this.state.events[Event.Type.DECLARED_OUTPUT]} />
        <SectionIO title={'Outputs'} artifactIds={this.state.events[Event.Type.OUTPUT]} />
      </div >
    );
  }

  public getInitialToolbarState(): ToolbarProps {
    return {
      actions: {},
      breadcrumbs: [{ displayName: 'Executions', href: RoutePage.EXECUTIONS }],
      pageTitle: `${this.properTypeName} ${this.id} details`
    };
  }

  public async refresh(): Promise<void> {
    return this.load();
  }

  private async load(): Promise<void> {
    const numberId = parseInt(this.id, 10);
    if (isNaN(numberId) || numberId < 0) {
      const error = new Error(`Invalid execution id: ${this.id}`);
      this.showPageError(error.message, error);
      return Promise.reject(error);
    }

    const getExecutionsRequest = new GetExecutionsByIDRequest();
    getExecutionsRequest.setExecutionIdsList([numberId]);
    const getEventsRequest = new GetEventsByExecutionIDsRequest();
    getEventsRequest.setExecutionIdsList([numberId]);

    const [executionResponse, eventResponse] = await Promise.all([
      Apis.getMetadataServicePromiseClient().getExecutionsByID(getExecutionsRequest),
      Apis.getMetadataServicePromiseClient().getEventsByExecutionIDs(getEventsRequest)
    ]);

    if (eventResponse.error) {
      this.showPageError(serviceErrorToString(eventResponse.error));
      // events data is optional, no need to skip the following
    }

    if (executionResponse.error) {
      this.showPageError(serviceErrorToString(executionResponse.error));
      return;
    }
    if (!executionResponse.response || !executionResponse.response.getExecutionsList().length) {
      this.showPageError(`No ${this.fullTypeName} identified by id: ${this.id}`);
      return;
    }
    if (executionResponse.response.getExecutionsList().length > 1) {
      this.showPageError(`Found multiple executions with ID: ${this.id}`);
      return;
    }

    const execution = executionResponse.response.getExecutionsList()[0];

    const executionName = getResourceProperty(execution, ExecutionProperties.COMPONENT_ID);
    this.props.updateToolbar({
      pageTitle: executionName ? executionName.toString() : ''
    });

    const events = parseEventsByType(eventResponse.response);
    this.setState({
      events,
      execution,
    });
  }
}

function parseEventsByType(response: GetEventsByExecutionIDsResponse | null): Record<EventTypes, ArtifactIdList> {
  const events: Record<EventTypes, ArtifactIdList> = {
    [Event.Type.UNKNOWN]: [],
    [Event.Type.DECLARED_INPUT]: [],
    [Event.Type.INPUT]: [],
    [Event.Type.DECLARED_OUTPUT]: [],
    [Event.Type.OUTPUT]: [],
  };

  if (!response) {
    return events;
  }

  response.getEventsList().forEach(event => {
    const type = event.getType();
    const id = event.getArtifactId();
    if (type != null && id != null) {
      events[type].push(id);
    }
  });

  return events;
}

interface SectionIOProps {
  title: string;
  artifactIds: number[];
}
class SectionIO extends Component<SectionIOProps, { artifactNameMap: {} }> {
  constructor(props: any) {
    super(props);

    this.state = {
      artifactNameMap: {},
    };
  }

  public async componentDidMount(): Promise<void> {
    const request = new GetArtifactsByIDRequest();
    request.setArtifactIdsList(this.props.artifactIds);
    const { error, response } = await Apis.getMetadataServicePromiseClient().getArtifactsByID(request);
    if (error || !response) {
      return;
    }

    const artifactNameMap = {};
    response.getArtifactsList().forEach(artifact => {
      const name = getResourceProperty(artifact, ArtifactProperties.NAME);
      const id = artifact.getId();
      if (!name || !id) {
        return;
      }
      artifactNameMap[id] = name;
    });
    this.setState({
      artifactNameMap,
    });
  }

  public render(): JSX.Element | null {
    const { title, artifactIds } = this.props;
    if (artifactIds.length === 0) {
      return null;
    }

    return <section>
      <h2 className={commonCss.header2}>{title}</h2>
      <table>
        <thead>
          <tr>
            <th className={css.tableCell}>Artifact ID</th>
            <th className={css.tableCell}>Name</th>
          </tr>
        </thead>
        <tbody>
          {artifactIds.map(id => <ArtifactRow key={id} id={id} name={this.state.artifactNameMap[id] || ''} />)}
        </tbody>
      </table>
    </section>;
  }
}

// tslint:disable-next-line:variable-name
const ArtifactRow: React.FC<{ id: number, name: string }> = ({ id, name }) => {
  return <tr>
    <td className={css.tableCell}>{id}</td>
    <td className={css.tableCell}>{name}</td>
  </tr>;
};

const css = stylesheet({
  tableCell: {
    padding: 4,
    textAlign: 'left',
  },
});
