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

from kubernetes.client import V1Toleration

def use_gcp_secret(secret_name='user-gcp-sa', secret_file_path_in_volume=None, volume_name=None, secret_volume_mount_path='/secret/gcp-credentials'):
    """An operator that configures the container to use GCP service account.

        The user-gcp-sa secret is created as part of the kubeflow deployment that
        stores the access token for kubeflow user service account.

        With this service account, the container has a range of GCP APIs to
        access to. This service account is automatically created as part of the
        kubeflow deployment.

        For the list of the GCP APIs this service account can access to, check
        https://github.com/kubeflow/kubeflow/blob/7b0db0d92d65c0746ac52b000cbc290dac7c62b1/deployment/gke/deployment_manager_configs/iam_bindings_template.yaml#L18

        If you want to call the GCP APIs in a different project, grant the kf-user
        service account access permission.
    """

    # permitted values for secret_name = ['admin-gcp-sa', 'user-gcp-sa']
    if secret_file_path_in_volume == None:
        secret_file_path_in_volume = '/' + secret_name + '.json'

    if volume_name == None:
        volume_name = 'gcp-credentials-' + secret_name

    else:
        import warnings
        warnings.warn('The volume_name parameter is deprecated and will be removed in next release. The volume names are now generated automatically.', DeprecationWarning)
    
    def _use_gcp_secret(task):
        from kubernetes import client as k8s_client
        return (
            task
                .add_volume(
                    k8s_client.V1Volume(
                        name=volume_name,
                        secret=k8s_client.V1SecretVolumeSource(
                            secret_name=secret_name,
                        )
                    )
                )
                .add_volume_mount(
                    k8s_client.V1VolumeMount(
                        name=volume_name,
                        mount_path=secret_volume_mount_path,
                    )
                )
                .add_env_variable(
                    k8s_client.V1EnvVar(
                        name='GOOGLE_APPLICATION_CREDENTIALS',
                        value=secret_volume_mount_path + secret_file_path_in_volume,
                    )
                )
                .add_env_variable(
                    k8s_client.V1EnvVar(
                        name='CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE',
                        value=secret_volume_mount_path + secret_file_path_in_volume,
                    )
                ) # Set GCloud Credentials by using the env var override.
                  # TODO: Is there a better way for GCloud to pick up the credential?
        )
    
    return _use_gcp_secret

def use_tpu(tpu_cores: int, tpu_resource: str, tf_version: str):
    """An operator that configures GCP TPU spec in a container op.

    Args:
      tpu_cores: Required. The number of cores of TPU resource. 
        For example, the value can be '8', '32', '128', etc.
        Check more details at: https://cloud.google.com/tpu/docs/kubernetes-engine-setup#pod-spec.
      tpu_resource: Required. The resource name of the TPU resource. 
        For example, the value can be 'v2', 'preemptible-v1', 'v3' or 'preemptible-v3'.
        Check more details at: https://cloud.google.com/tpu/docs/kubernetes-engine-setup#pod-spec.
      tf_version: Required. The TensorFlow version that the TPU nodes use.
        For example, the value can be '1.12', '1.11', '1.9' or '1.8'.
        Check more details at: https://cloud.google.com/tpu/docs/supported-versions.
    """

    def _set_tpu_spec(task):
        task.add_pod_annotation('tf-version.cloud-tpus.google.com', tf_version)
        task.add_resource_limit('cloud-tpus.google.com/{}'.format(tpu_resource), str(tpu_cores))
        return task

    return _set_tpu_spec

def use_preemptible_nodepool(toleration: V1Toleration = V1Toleration(effect='NoSchedule',
                                                             key='preemptible',
                                                             operator='Equal',
                                                             value='true')):
  """An operator that configures the GKE preemptible in a container op.
  """

  def _set_preemptible(task):
    task.add_toleration(toleration)
    task.add_node_selector_constraint("cloud.google.com/gke-preemptible", "true")
    return task

  return _set_preemptible