from typing import NamedTuple


def run_diagnose_me(target_apis: str, bucket: str) -> NamedTuple('Outputs', [('project_id', str), ('bucket', str)]):
    """ Performs enviroment verification specific to this pipeline.

    Raises:
        RuntimeError: If configuration is not setup properly and
        HALT_ON_ERROR flag is set.
    """

    # Installing pip3 and kfp, since the base image 'google/cloud-sdk:276.0.0' does not come with pip3 pre-installed.
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
        print('GCP credentials are configured with access to project: %s ...\n' %
              (project_id))
        print('Following account(s) are active under this pipeline:\n')
        subprocess.run(['gcloud', 'auth', 'list'])
        print('\n')
    else:
        print('Project configuration is not accessible with error  %s\n' %
              (project_config.stderr) + 'Follow the instructions at\n' +
              'https://github.com/kubeflow/pipelines/blob/master/manifests/gcp_marketplace/guide.md#gcp-service-account-credentials \n'
              + 'to verify you have configured the required gcp secret.', file=sys.stderr)
        config_error_observed = True

    # Get project buckets
    get_project_bucket_results = gcp.get_gcp_configuration(
        gcp.Commands.GET_STORAGE_BUCKETS, human_readable=False)

    if get_project_bucket_results.has_error:
        print('could not retrieve project buckets with error: %s' %
              (get_project_bucket_results.stderr), file=sys.stderr)
        config_error_observed = True

    # Get the root of the user provided bucket i.e. gs://root.
    bucket_root = '/'.join(bucket.split('/')[0:3])

    print('Checking to see if the provided GCS bucket\n  %s\nis accessible ...\n' % (bucket))

    if bucket_root in get_project_bucket_results.json_output:
        print('Provided bucket \n   %s\nis accessible within the project\n   %s.' % (
            bucket, project_id))
        print('\n\n')

    else:
        print('''
    Could not find the bucket %s in project %s
    Please verify that you have provided the correct GCS bucket name.

    Only the following buckets are visible in this project:
    \n%s
                '''
              % (bucket, project_id, get_project_bucket_results.parsed_output), file=sys.stderr)
        config_error_observed = True

    # Verify APIs that are required are enabled
    api_config_results = gcp.get_gcp_configuration(
        gcp.Commands.GET_APIS)

    api_status = {}

    if api_config_results.has_error:
        raise RuntimeError('could not retrieve API status with error: %s' % (
            api_config_results.stderr))

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
                'API \"%s\" is not enabled. To enable this api go to https://pantheon.corp.google.com/apis/library/%s?project=%s' % (api, api, project_id))

    if not api_check_results:
        config_error_observed = True
        print('Required APIs are not enabled:\n' +
              '\n'.join(error_list), file=sys.stderr)

    return (project_id, bucket)


if __name__ == '__main__':
    import kfp.components as comp

    comp.func_to_container_op(
        run_diagnose_me,
        base_image='google/cloud-sdk:276.0.0',
        output_component_file='component.yaml',
    )
