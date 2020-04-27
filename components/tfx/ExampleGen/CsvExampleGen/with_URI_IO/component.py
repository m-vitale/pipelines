# flake8: noqa TODO

from typing import NamedTuple

def CsvExampleGen_GCS( #
    # Inputs
    #input_base_path: InputPath('ExternalPath'),
    input_base_path: 'ExternalPath', # A Channel of 'ExternalPath' type, which includes one artifact whose uri is an external directory with csv files inside (required).

    # Outputs
    #example_artifacts_path: OutputPath('ExamplesPath'),
    example_artifacts_path: 'ExamplesPath',

    # Execution properties
    #input_config_splits: {'List' : {'item_type': 'ExampleGen.Input.Split'}},
    input_config: 'ExampleGen.Input' = None, # = '{"splits": []}', # JSON-serialized example_gen_pb2.Input instance, providing input configuration. If unset, the files under input_base will be treated as a single split.
    #output_config_splits: {'List' : {'item_type': 'ExampleGen.SplitConfig'}},
    output_config: 'ExampleGen.Output' = None, # = '{"splitConfig": {"splits": []}}', # JSON-serialized example_gen_pb2.Output instance, providing output configuration. If unset, default splits will be 'train' and 'eval' with size 2:1.
    #custom_config: 'ExampleGen.CustomConfig' = None,
) -> NamedTuple('Outputs', [
    ('example_artifacts', 'ExamplesPath'),
]):
    """Executes the CsvExampleGen component.

    Args:
      input_base: A Channel of 'ExternalPath' type, which includes one artifact
        whose uri is an external directory with csv files inside (required).
      input_config: An example_gen_pb2.Input instance, providing input
        configuration. If unset, the files under input_base will be treated as a
        single split.
      output_config: An example_gen_pb2.Output instance, providing output
        configuration. If unset, default splits will be 'train' and 'eval' with
        size 2:1.
      ??? input: Forwards compatibility alias for the 'input_base' argument.
    Returns:
      example_artifacts: Artifact of type 'ExamplesPath' for output train and
        eval examples.
    """
    from tfx.components.example_gen.csv_example_gen.component import CsvExampleGen as component_class

    #Generated code
    import json
    import os
    import tensorflow
    from google.protobuf import json_format, message
    from tfx.types import Artifact, channel_utils, artifact_utils

    arguments = locals().copy()

    component_class_args = {}

    for name, execution_parameter in component_class.SPEC_CLASS.PARAMETERS.items():
        argument_value_obj = argument_value = arguments.get(name, None)
        if argument_value is None:
            continue
        parameter_type = execution_parameter.type
        if isinstance(parameter_type, type) and issubclass(parameter_type, message.Message): # Maybe FIX: execution_parameter.type can also be a tuple
            argument_value_obj = parameter_type()
            json_format.Parse(argument_value, argument_value_obj)
        component_class_args[name] = argument_value_obj

    for name, channel_parameter in component_class.SPEC_CLASS.INPUTS.items():
        artifact_path = arguments[name + '_path']
        if artifact_path:
            artifact = channel_parameter.type()
            artifact.uri = artifact_path + '/' # ?
            if channel_parameter.type.PROPERTIES and 'split_names' in channel_parameter.type.PROPERTIES:
                # Recovering splits
                subdirs = tensorflow.io.gfile.listdir(artifact_path)
                artifact.split_names = artifact_utils.encode_split_names(sorted(subdirs))
            component_class_args[name] = channel_utils.as_channel([artifact])

    component_class_instance = component_class(**component_class_args)

    input_dict = {name: channel.get() for name, channel in component_class_instance.inputs.get_all().items()}
    output_dict = {name: channel.get() for name, channel in component_class_instance.outputs.get_all().items()}
    exec_properties = component_class_instance.exec_properties

    # Generating paths for output artifacts
    for name, artifacts in output_dict.items():
        base_artifact_path = arguments[name + '_path']
        # Are there still cases where output channel has multiple artifacts?
        for idx, artifact in enumerate(artifacts):
            subdir = str(idx + 1) if idx > 0 else ''
            artifact.uri = os.path.join(base_artifact_path, subdir)  # Ends with '/'

    print('component instance: ' + str(component_class_instance))

    #executor = component_class.EXECUTOR_SPEC.executor_class() # Same
    executor = component_class_instance.executor_spec.executor_class()
    executor.Do(
        input_dict=input_dict,
        output_dict=output_dict,
        exec_properties=exec_properties,
    )
    # End generated code
    return (example_artifacts_path,)

if __name__ == '__main__':
    import kfp
    kfp.components.func_to_container_op(
        CsvExampleGen_GCS,
        base_image='tensorflow/tfx:0.21.4',
        output_component_file='component.yaml'
    )
