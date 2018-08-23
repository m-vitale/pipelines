# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# A program to generate confusion matrix data out of prediction results.
# Usage:
# python confusion_matrix.py  \
#   --predictions=gs://bradley-playground/sfpd/predictions/part-* \
#   --output=gs://bradley-playground/sfpd/cm/ \
#   --target=resolution \
#   --analysis=gs://bradley-playground/sfpd/analysis \


import argparse
import json
import os
import pandas as pd
from sklearn.metrics import confusion_matrix
from tensorflow.python.lib.io import file_io


def main(argv=None):
  parser = argparse.ArgumentParser(description='ML Trainer')
  parser.add_argument('--predictions', type=str, help='GCS path of prediction file pattern.')
  parser.add_argument('--output', type=str, help='GCS path of the output directory.')
  args = parser.parse_args()

  schema_file = os.path.join(os.path.dirname(args.predictions), 'schema.json')
  schema = json.loads(file_io.read_file_to_string(schema_file))
  names = [x['name'] for x in schema]
  dfs = []
  files = file_io.get_matching_files(args.predictions)
  for file in files:
    with file_io.FileIO(file, 'r') as f:
      dfs.append(pd.read_csv(f, names=names))
    
  df = pd.concat(dfs)
  vocab = list(df['target'].unique())
  cm = confusion_matrix(df['target'], df['predicted'], labels=vocab)
  data = []
  for target_index, target_row in enumerate(cm):
    for predicted_index, count in enumerate(target_row):
      data.append((vocab[target_index], vocab[predicted_index], count))

  df_cm = pd.DataFrame(data, columns=['target', 'predicted', 'count'])
  cm_file = os.path.join(args.output, 'confusion_matrix.csv')
  with file_io.FileIO(cm_file, 'w') as f:
    df_cm.to_csv(f, columns=['target', 'predicted', 'count'], header=False, index=False)

  metadata = {
    'outputs' : [{
      'type': 'confusion_matrix',
      'storage': 'gcs',
      'format': 'csv',
      'schema': [
        {'name': 'target', 'type': 'CATEGORY'},
        {'name': 'predicted', 'type': 'CATEGORY'},
        {'name': 'count', 'type': 'NUMBER'},
      ],
      'source': cm_file,
      'labels': vocab,
    }]
  }
  with file_io.FileIO('/mlpipeline-ui-metadata.json', 'w') as f:
    json.dump(metadata, f)


if __name__== "__main__":
  main()
