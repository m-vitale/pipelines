/*
 * Copyright 2019 Google LLC
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
import { Page } from './Page';
import { ToolbarProps } from '../components/Toolbar';
import Markdown from 'markdown-to-jsx';
import { ExternalLink } from '../atoms/ExternalLink';
import { cssRaw, classes } from 'typestyle';
import { commonCss, padding } from '../Css';

const options = {
  overrides: { a: { component: ExternalLink } },
};

const PAGE_CONTENT_MD = `
Build a end-to-end ML pipeline with TFX  [Start Here](https://console.cloud.google.com/mlengine/notebooks/deploy-notebook?q=download_url%3Dhttps%253A%252F%252Fraw.githubusercontent.com%252Fkubeflow%252Fpipelines%252F0.1.40%252Fsamples%252Fcore%252Fparameterized_tfx_oss%252Ftaxi_pipeline_notebook.ipynb) (Alpha)

## Demos and Tutorials

This section contains demo and tutorial pipelines.

**Demos** - Try an end-to-end demonstration pipeline.

  * [TFX pipeline demo](https://www.google.com) \\- A trainer that does end-to-end distributed training for XGBoost models. [source code](https://github.com/kubeflow/pipelines/tree/master/samples/core/parameterized_tfx_oss)
  * [XGBoost Pipeline](https://www.google.com) \\- Example pipeline that does classification with model analysis based on a public taxi cab BigQuery dataset. [source code](https://github.com/kubeflow/pipelines/tree/master/samples/core/xgboost_training_cm)


**Tutorials** - Learn pipeline concepts by following a tutorial.

  * [Name of Tutorial 1] - \\<tutorial 1 description\\>. [source code]()
  * [Name of Tutorial 2] - \\<tutorial 2 description\\>. [source code]()

  You can find additional tutorials and samples [here]()

### Additional resources and documentation
  * [TFX Landing page]()
  * [Hosted Pipeline documentation]()
  * [Troubleshooting guide]()
`;

cssRaw(`
.kfp-start-page li {
  font-size: 14px;
  margin-block-start: 0.83em;
  margin-block-end: 0.83em;
  margin-left: 2em;
}
.kfp-start-page p {
  font-size: 14px;
  margin-block-start: 0.83em;
  margin-block-end: 0.83em;
}
.kfp-start-page h2 {
  font-size: 18px;
  margin-block-start: 1em;
  margin-block-end: 1em;
}
.kfp-start-page h3 {
  font-size: 16px;
  margin-block-start: 1em;
  margin-block-end: 1em;
}
`);

export class GettingStarted extends Page<{}, {}> {
  public getInitialToolbarState(): ToolbarProps {
    const buttons = new Buttons(this.props, this.refresh.bind(this));
    return {
      actions: buttons.getToolbarActionMap(),
      breadcrumbs: [],
      pageTitle: 'Getting Started: Build your own pipeline',
    };
  }

  public async refresh() {
    // do nothing
  }

  public render(): JSX.Element {
    return (
      <div className={classes(commonCss.page, padding(20, 'lr'), 'kfp-start-page')}>
        <Markdown options={options}>{PAGE_CONTENT_MD}</Markdown>
      </div>
    );
  }
}
