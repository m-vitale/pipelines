// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import { Stream } from 'stream';
import * as tar from 'tar-stream';
import * as gunzip from 'gunzip-maybe';
import { Client as MinioClient, ClientOptions as MinioClientOptions } from 'minio';
import { awsInstanceProfileCredentials } from './aws-helper';

/** MinioRequestConfig describes the info required to retrieve an artifact. */
export interface MinioRequestConfig {
  bucket: string;
  key: string;
  client: MinioClient;
}

/** MinioClientOptionsWithOptionalSecrets wraps around MinioClientOptions where only endPoint is required (accesskey and secretkey are optional). */
export interface MinioClientOptionsWithOptionalSecrets extends Partial<MinioClientOptions> {
  endPoint: string;
}

/**
 * Create minio client with aws instance profile credentials if needed.
 * @param config minio client options where `accessKey` and `secretKey` are optional.
 */
export async function createMinioClient(config: MinioClientOptionsWithOptionalSecrets) {
  if (!config.accessKey || !config.secretKey) {
    try {
      if (await awsInstanceProfileCredentials.ok()) {
        const credentials = await awsInstanceProfileCredentials.getCredentials();
        if (credentials) {
          const {
            AccessKeyId: accessKey,
            SecretAccessKey: secretKey,
            Token: sessionToken,
          } = credentials;
          return new MinioClient({ ...config, accessKey, secretKey, sessionToken });
        }
        console.error('unable to get credentials from AWS metadata store.');
      }
    } catch (err) {
      console.error('Unable to get aws instance profile credentials: ', err);
    }
  }
  return new MinioClient(config as MinioClientOptions);
}

export function getTarObjectAsString({ bucket, key, client }: MinioRequestConfig) {
  return new Promise<string>(async (resolve, reject) => {
    try {
      const stream = await getObjectStream({ bucket, key, client });
      const extract = tar.extract();
      let emitOnce = false;
      extract.once('entry', function(_header, stream, next) {
        let content = '';
        stream.on('data', buffer => (content += buffer.toString()));
        stream.on('end', () => {
          emitOnce = true;
          resolve(content);
          next();
        });
        stream.resume(); // just auto drain the stream
      });
      extract.on('end', function() {
        if (!emitOnce) {
          resolve('');
        }
      });
      extract.on('error', reject);
      stream.pipe(gunzip()).pipe(extract);
    } catch (err) {
      reject(err);
    }
  });
}

export function getObjectStream({ bucket, key, client }: MinioRequestConfig) {
  return client.getObject(bucket, key);
}
