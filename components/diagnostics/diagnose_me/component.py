inputs:
- {type: String, name: bucket}
- {type: String, name: execution_mode}
- {type: String, name: target_apis}
implementation:
  container:
    args:
    - --bucket
    - {inputValue: bucket}
    - --execution-mode
    - {inputValue: execution_mode}
    - --target-apis
    - {inputValue: target_apis}
    - '----output-paths'
    - {outputPath: bucket}
    - {outputPath: project_id}
    command:
    - python3
    - -u
    - -c
    - |
      from typing import NamedTuple

      def run_diagnose_me(
          bucket: str, execution_mode: str, target_apis: str
      ) -> NamedTuple('Outputs', [('bucket', str), ('project_id', 'GCPProjectID')]):
          """ Performs environment verification specific to this pipeline.

            args:
                bucket:
                    string name of the bucket to be checked. Must be of the format
                    gs://bucket_root/any/path/here/is/ignored where any path beyond root
                    is ignored.
                execution_mode:
                    If set to HALT_ON_ERROR will case any error to raise an exception.
                    This is intended to stop the data processing of a pipeline. Can set
                    to False to only report Errors/Warnings.
                target_apis:
                    String consisting of a comma separated list of apis to be verified.
            Raises:
                RuntimeError: If configuration is not setup properly and
                HALT_ON_ERROR flag is set.
            """

          # Installing pip3 and kfp, since the base image 'google/cloud-sdk:276.0.0'
          # does not come with pip3 pre-installed.
          import subprocess
          subprocess.run(
              ['curl', 'https://bootstrap.pypa.io/get-pip.py', '-o', 'get-pip.py'],
              capture_output=True)
          subprocess.run(['apt-get', 'install', 'python3-distutils', '--yes'],
                         capture_output=True)
          subprocess.run(['python3', 'get-pip.py'], capture_output=True)
          subprocess.run(['python3', '-m', 'pip', 'install', 'kfp>=0.1.31', '--quiet'],
                         capture_output=True)

          import sys
          from typing import List, Text
          import os
          from kfp.cli.diagnose_me import gcp
          from kfp.cli.diagnose_me import utility
          import json as json_library

          config_error_observed = False

          # Get the project ID
          # from project configuration
          project_config = gcp.get_gcp_configuration(
              gcp.Commands.GET_GCLOUD_DEFAULT, human_readable=False)
          project_id = ''
          if not project_config.has_error:
              project_id = project_config.parsed_output['core']['project']
              print('GCP credentials are configured with access to project: %s ...\n' % (project_id))
              print('Following account(s) are active under this pipeline:\n')
              subprocess.run(['gcloud', 'auth', 'list'])
              print('\n')
          else:
              print(
                  'Project configuration is not accessible with error  %s\n' %
                  (project_config.stderr),
                  file=sys.stderr)
              config_error_observed = True

          # Get project buckets
          get_project_bucket_results = gcp.get_gcp_configuration(
              gcp.Commands.GET_STORAGE_BUCKETS, human_readable=False)

          if get_project_bucket_results.has_error:
              print(
                  'could not retrieve project buckets with error: %s' %
                  (get_project_bucket_results.stderr),
                  file=sys.stderr)
              config_error_observed = True

          # Get the root of the user provided bucket i.e. gs://root.
          bucket_root = '/'.join(bucket.split('/')[0:3])

          print(
              'Checking to see if the provided GCS bucket\n  %s\nis accessible ...\n' %
              (bucket))

          if bucket_root in get_project_bucket_results.json_output:
              print('Provided bucket \n   %s\nis accessible within the project\n   %s\n' %
                    (bucket, project_id))

          else:
              print('Could not find the bucket %s in project %s' % (bucket, project_id) +
                    'Please verify that you have provided the correct GCS bucket name.\n' +
                    'Only the following buckets are visible in this project:\n%s' %
                    (get_project_bucket_results.parsed_output),
                    file=sys.stderr)
              config_error_observed = True

          # Verify APIs that are required are enabled
          api_config_results = gcp.get_gcp_configuration(gcp.Commands.GET_APIS)

          api_status = {}

          if api_config_results.has_error:
              print('could not retrieve API status with error: %s' %
                                 (api_config_results.stderr),sys.stderr)
              config_error_observed = True

          print('Checking APIs status ...')
          for item in api_config_results.parsed_output:
              api_status[item['config']['name']] = item['state']
              # printing the results in stdout for logging purposes
              print('%s %s' % (item['config']['name'], item['state']))

          # Check if target apis are enabled
          api_check_results = True
          error_list = []
          for api in target_apis.replace(' ', '').split(','):
              if 'ENABLED' != api_status.get(api, 'DISABLED'):
                  api_check_results = False
                  error_list.append(
                      'API \"%s\" is not enabled. To enable this api go to ' % (api) +
                      'https://pantheon.corp.google.com/apis/library/%s?project=%s' %
                      (api, project_id))

          if not api_check_results:
              config_error_observed = True
              print(
                  'Required APIs are not enabled:\n' + '\n'.join(error_list),
                  file=sys.stderr)

          if 'HALT_ON_ERROR' in execution_mode and config_error_observed:
              raise RuntimeError('There was an error in your environment configuration. See above for details.\n'+
                  'Note that resolving such issues generally require a deep knowledge of Kubernetes.\n'+
                  '\n'+
                  'We highly recommend that you recreate the cluster and check "Allow access ..." \n'+
                  'checkbox during cluster creation to have the cluster configured automatically.\n'+
                  'For more information on this and other troubleshooting instructions refer to\n'+
                  'our troubleshooting guide.\n'+
                  '\n'+
                  'If you have intentionally modified the cluster configuration, you may\n'+
                  'bypass this error by removing the execution_mode HALT_ON_ERROR flag.\n')

          return (project_id, bucket)

      def _serialize_str(str_value: str) -> str:
          if not isinstance(str_value, str):
              raise TypeError('Value "{}" has type "{}" instead of str.'.format(str(str_value), str(type(str_value))))
          return str_value

      import argparse
      _parser = argparse.ArgumentParser(prog='Run diagnose me', description='Performs environment verification specific to this pipeline.\n\n      args:\n          bucket:\n              string name of the bucket to be checked. Must be of the format\n              gs://bucket_root/any/path/here/is/ignored where any path beyond root\n              is ignored.\n          execution_mode:\n              If set to HALT_ON_ERROR will case any error to raise an exception.\n              This is intended to stop the data processing of a pipeline. Can set\n              to False to only report Errors/Warnings.\n          target_apis:\n              String consisting of a comma separated list of apis to be verified.\n      Raises:\n          RuntimeError: If configuration is not setup properly and\n          HALT_ON_ERROR flag is set.\n')
      _parser.add_argument("--bucket", dest="bucket", type=str, required=True, default=argparse.SUPPRESS)
      _parser.add_argument("--execution-mode", dest="execution_mode", type=str, required=True, default=argparse.SUPPRESS)
      _parser.add_argument("--target-apis", dest="target_apis", type=str, required=True, default=argparse.SUPPRESS)
      _parser.add_argument("----output-paths", dest="_output_paths", type=str, nargs=2)
      _parsed_args = vars(_parser.parse_args())
      _output_files = _parsed_args.pop("_output_paths", [])

      _outputs = run_diagnose_me(**_parsed_args)

      if not hasattr(_outputs, '__getitem__') or isinstance(_outputs, str):
          _outputs = [_outputs]

      _output_serializers = [
          _serialize_str,
          str,

      ]

      import os
      for idx, output_file in enumerate(_output_files):
          try:
              os.makedirs(os.path.dirname(output_file))
          except OSError:
              pass
          with open(output_file, 'w') as f:
              f.write(_output_serializers[idx](_outputs[idx]))
    image: google/cloud-sdk:276.0.0
description: |
  Performs environment verification specific to this pipeline.

        args:
            bucket:
                string name of the bucket to be checked. Must be of the format
                gs://bucket_root/any/path/here/is/ignored where any path beyond root
                is ignored.
            execution_mode:
                If set to HALT_ON_ERROR will case any error to raise an exception.
                This is intended to stop the data processing of a pipeline. Can set
                to False to only report Errors/Warnings.
            target_apis:
                String consisting of a comma separated list of apis to be verified.
        Raises:
            RuntimeError: If configuration is not setup properly and
            HALT_ON_ERROR flag is set.
name: Run diagnose me
outputs:
- {type: String, name: bucket}
- {type: GCPProjectID, name: project_id}
