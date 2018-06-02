import * as express from 'express';
import * as fs from 'fs';
import * as _path from 'path';
import proxyMiddleware from '../server/proxy-middleware';

import { JobMetadata } from '../src/model/job';
import { JobSortKeys } from '../src/model/list_jobs_request';
import { ListJobsResponse } from '../src/model/list_jobs_response';
import { PipelineSortKeys } from '../src/model/list_pipelines_request';
import { ListPipelinesResponse } from '../src/model/list_pipelines_response';
import { Pipeline } from '../src/model/pipeline';

const prefix = __dirname + '/pipeline-data';

const fixedData = require('./fixed-data').data;

const rocMetadataJsonPath = './eval-output/metadata.json';
const rocDataPath = './eval-output/roc.csv';

const confusionMatrixMetadataJsonPath = './model-output/metadata.json';
const confusionMatrixPath = './model-output/confusion_matrix.csv';

const apisPrefix = '/apis/v1alpha1';

let tensorboardPod = '';

let apiServerReady = false;

// Simulate API server not ready for 5 seconds
setTimeout(() => {
  apiServerReady = true;
}, 5000);

export default (app) => {

  app.set('json spaces', 2);
  app.use(express.json());

  app.get(apisPrefix + '/healthz', (req, res) => {
    if (apiServerReady) {
      res.send({ apiServerReady });
    } else {
      res.status(404).send();
    }
  });

  app.get(apisPrefix + '/pipelines', (req, res) => {
    if (!apiServerReady) {
      res.status(404).send();
      return;
    }

    res.header('Content-Type', 'application/json');
    // Note: the way that we use the nextPageToken here may not reflect the way the backend works.
    const response: ListPipelinesResponse = {
      nextPageToken: '',
      pipelines: [],
    };

    let pipelines: Pipeline[] = fixedData.pipelines;
    if (req.query.filterBy) {
      // NOTE: We do not mock fuzzy matching. E.g. 'ee' doesn't match 'Pipeline'
      // This may need to be updated when the backend implements filtering.
      pipelines = fixedData.pipelines.filter((p) => p.name.toLocaleLowerCase().match(
          decodeURIComponent(req.query.filterBy).toLocaleLowerCase()));

    }

    if (req.query.sortBy) {
      let sortByPropName = '';

      switch (req.query.sortBy) {
        case PipelineSortKeys.CREATED_AT:
          sortByPropName = 'createdAt';
          break;
        case PipelineSortKeys.ID:
          sortByPropName = 'id';
          break;
        case PipelineSortKeys.NAME:
          sortByPropName = 'name';
          break;
        case PipelineSortKeys.PACKAGE_ID:
          sortByPropName = 'packageId';
          break;
        default:
          res.status(405).send(`Unsupported sort string: ${req.query.sortBy}`);
          return;
      }

      pipelines.sort((a, b) => {
        let result = 1;
        if (a[sortByPropName] < b[sortByPropName]) {
          result = -1;
        }
        if (a[sortByPropName] === b[sortByPropName]) {
          result = 0;
        }
        return result * ((req.query.ascending === 'true') ? 1 : -1);
      });
    }

    const start = (req.query.pageToken ? +req.query.pageToken : 0);
    const end = start + (+req.query.pageSize) + 1;
    response.pipelines = pipelines.slice(start, end);

    if (end < pipelines.length) {
      response.nextPageToken = end + '';
    }

    res.json(response);
  });

  app.post(apisPrefix + '/pipelines', (req, res) => {
    const pipeline = req.body;
    pipeline.id = fixedData.pipelines.length;
    pipeline.createdAt = Math.floor(Date.now() / 1000);
    pipeline.jobs = [];
    pipeline.enabled = !!pipeline.schedule;
    fixedData.pipelines.push(pipeline);
    setTimeout(() => {
      res.send(fixedData.pipelines[0]);
    }, 1000);
  });

  app.all(apisPrefix + '/pipelines/:pid', (req, res) => {
    res.header('Content-Type', 'application/json');
    const pid = Number.parseInt(req.params.pid);
    switch (req.method) {
      case 'DELETE':
        const i = fixedData.pipelines.findIndex((p) => p.id === pid);
        if (fixedData.pipelines[i].name.startsWith('Cannot be deleted')) {
          res.status(502).send(`Deletion failed for Pipeline: '${fixedData.pipelines[i].name}'`);
        } else {
          // Delete the Pipeline from fixedData.
          fixedData.pipelines.splice(i, 1);
          res.send('ok');
        }
        break;
      case 'GET':
        res.json(fixedData.pipelines.find((p) => p.id === pid));
        break;
      default:
        res.status(405).send('Unsupported request type: ' + res.method);
    }
  });

  app.get(apisPrefix + '/pipelines/:pid/jobs', (req, res) => {
    res.header('Content-Type', 'application/json');
    const pid = Number.parseInt(req.params.pid);
    // Note: the way that we use the nextPageToken here may not reflect the way the backend works.
    const response: ListJobsResponse = {
      jobs: [],
      nextPageToken: '',
    };

    let jobs: JobMetadata[] =
        fixedData.pipelines.find((p) => p.id === pid).jobs.map((j) => j.metadata);

    if (req.query.filterBy) {
      // NOTE: We do not mock fuzzy matching. E.g. 'ee' doesn't match 'Pipeline'
      // This may need to be updated when the backend implements filtering.
      jobs = jobs.filter((j) => j.name.toLocaleLowerCase().match(
          decodeURIComponent(req.query.filterBy).toLocaleLowerCase()));
    }

    // The backend sorts by createdAt by default.
    req.query.sortBy = req.query.sortBy || JobSortKeys.CREATED_AT;

    if (req.query.sortBy) {
      let sortByPropName = '';

      switch (req.query.sortBy) {
        case JobSortKeys.NAME:
          sortByPropName = 'name';
          break;
        case JobSortKeys.CREATED_AT:
          // Intentionally falling through
        case '':
          sortByPropName = 'createdAt';
          break;
        default:
          res.status(405).send(`Unsupported sort string: ${req.query.sortBy}`);
          return;
      }

      jobs.sort((a, b) => {
        let result = 1;
        if (a[sortByPropName] < b[sortByPropName]) {
          result = -1;
        }
        if (a[sortByPropName] === b[sortByPropName]) {
          result = 0;
        }
        return result * ((req.query.ascending === 'true') ? 1 : -1);
      });
    }

    const start = (req.query.pageToken ? +req.query.pageToken : 0);
    const end = start + (+req.query.pageSize) + 1;
    response.jobs = jobs.slice(start, end);

    if (end < jobs.length) {
      response.nextPageToken = end + '';
    }

    res.json(response);
  });

  app.post(apisPrefix + '/pipelines/:pid/enable', (req, res) => {
    setTimeout(() => {
      const pid = Number.parseInt(req.params.pid);
      const pipeline = fixedData.pipelines.find((p) => p.id === pid);
      pipeline.enabled = true;
      res.send('ok');
    }, 1000);
  });

  app.post(apisPrefix + '/pipelines/:pid/disable', (req, res) => {
    setTimeout(() => {
      const pid = Number.parseInt(req.params.pid);
      const pipeline = fixedData.pipelines.find((p) => p.id === pid);
      pipeline.enabled = false;
      res.send('ok');
    }, 1000);
  });

  app.get(apisPrefix + '/pipelines/:pid/jobs/:jname', (req, res) => {
    const pid = Number.parseInt(req.params.pid);
    const jname = req.params.jname;
    const pipeline = fixedData.pipelines.find((p) => p.id === pid);
    const job = pipeline.jobs.find((j) => j.metadata.name === jname);
    if (!job) {
      res.status(404).send('Cannot find a job with name: ' + jname);
      return;
    }
    res.json(job.jobDetail);
  });

  app.get(apisPrefix + '/packages', (req, res) => {
    res.header('Content-Type', 'application/json');
    res.json(fixedData.packages);
  });

  app.get(apisPrefix + '/packages/:pid/templates', (req, res) => {
    res.header('Content-Type', 'text/x-yaml');
    res.send(fs.readFileSync('./mock-backend/mock-template.yaml'));
  });

  app.post(apisPrefix + '/packages/upload', (req, res) => {
    res.header('Content-Type', 'application/json');
    res.json(fixedData.packages[0]);
  });

  app.get(apisPrefix + '/artifacts/list/:path', (req, res) => {

    const path = decodeURIComponent(req.params.path);

    res.header('Content-Type', 'application/json');
    res.json([
      path + '/file1',
      path + '/file2',
      path + (path.match('analysis$|model$') ? '/metadata.json' : '/file3'),
    ]);
  });

  app.get(apisPrefix + '/artifacts/get/:path', (req, res) => {
    res.header('Content-Type', 'application/json');
    const path = decodeURIComponent(req.params.path);
    if (path.endsWith('roc.csv')) {
      res.sendFile(_path.resolve(__dirname, rocDataPath));
    } else if (path.endsWith('confusion_matrix.csv')) {
      res.sendFile(_path.resolve(__dirname, confusionMatrixPath));
    } else if (path.endsWith('analysis/metadata.json')) {
      res.sendFile(_path.resolve(__dirname, confusionMatrixMetadataJsonPath));
    } else if (path.endsWith('model/metadata.json')) {
      res.sendFile(_path.resolve(__dirname, rocMetadataJsonPath));
    } else {
      res.send('dummy file');
    }
  });

  app.get(apisPrefix + '/apps/tensorboard', (req, res) => {
    res.send(tensorboardPod);
  });

  app.post(apisPrefix + '/apps/tensorboard', (req, res) => {
    tensorboardPod = 'http://tensorboardserver:port';
    setTimeout(() => {
      res.send('ok');
    }, 1000);
  });

  app.get(apisPrefix + '/k8s/pod/logs', (req, res) => {
    setTimeout(() => {
      res.send(String.raw`
      _____________
    < hello world >
      -------------
           \
            \
             \
                           ##        .
                     ## ## ##       ==
                  ## ## ## ##      ===
              /""""""""""""""""___/ ===
         ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
              \______ o          __/
               \    \        __/
                 \____\______/

       `);
    }, 300);
  });

  app.get('/_componentTests*', (req, res) => {
    res.sendFile(_path.resolve('test', 'components', 'index.html'));
  });

  app.all(apisPrefix + '*', (req, res) => {
    res.status(404).send('Bad request endpoint.');
  });

  proxyMiddleware(app, apisPrefix);

};
