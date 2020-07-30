from dataclasses import dataclass

from typing import (
    List,
    NamedTuple,
    NewType,
    Optional,
    Union,
)
from .spec_validators import SpecValidators

# NamedTuple does not support generic types yet, otherwise we could replace
# `object` with T.
@dataclass(frozen=True)
class SageMakerComponentInputValidator:
    """Defines the structure of a component input to be used for validation."""

    input_type: object
    description: str
    required: bool = False
    choices: Optional[List[object]] = None
    default: Optional[object] = None

    def to_argparse_mapping(self):
        """Maps each property to an argparse argument.

        See: https://docs.python.org/3/library/argparse.html#adding-arguments
        """
        return {
            "type": self.input_type,
            "help": self.description,
            "required": self.required,
            "choices": self.choices,
            "default": self.default
        }


@dataclass(frozen=True)
class SageMakerComponentOutputValidator:
    """Defines the structure of a component output."""

    description: str


# The value of an input or output can be arbitrary.
# This could be replaced with a generic, as well.
SageMakerIOValue = NewType("SageMakerIOValue", object)

# Allow the component input to represent either the validator or the value itself
# This saves on having to rewrite the struct for both types.
# Could replace `object` with T if TypedDict supported generics (see above).
SageMakerComponentInput = NewType(
    "SageMakerComponentInput", Union[SageMakerComponentInputValidator, SageMakerIOValue]
)
SageMakerComponentOutput = NewType(
    "SageMakerComponentOutput",
    Union[SageMakerComponentOutputValidator, SageMakerIOValue],
)


@dataclass(frozen=True)
class SageMakerComponentBaseInputs:
    """The base class for all component inputs."""

    pass


@dataclass
class SageMakerComponentBaseOutputs:
    """A base class for all component outputs."""

    pass


@dataclass()
class SageMakerComponentCommonInputs(SageMakerComponentBaseInputs):
    """A set of common inputs for all components."""

    region: SageMakerComponentInput
    endpoint_url: SageMakerComponentInput
    tags: SageMakerComponentInput


COMMON_INPUTS = SageMakerComponentCommonInputs(
    region=SageMakerComponentInputValidator(
        input_type=str,
        required=True,
        description="The region where the training job launches.",
    ),
    endpoint_url=SageMakerComponentInputValidator(
        input_type=SpecValidators.nullable_string_argument,
        required=False,
        description="The URL to use when communicating with the SageMaker service.",
    ),
    tags=SageMakerComponentInputValidator(
        input_type=SpecValidators.yaml_or_json_dict,
        required=False,
        description="An array of key-value pairs, to categorize AWS resources.",
        default={},
    ),
)


@dataclass(frozen=True)
class SpotInstanceInputs(SageMakerComponentBaseInputs):
    """Inputs to enable spot instance support."""

    spot_instance: SageMakerComponentInput
    max_wait_time: SageMakerComponentInput
    checkpoint_config: SageMakerComponentInput


SPOT_INSTANCE_INPUTS = SpotInstanceInputs(
    spot_instance=SageMakerComponentInputValidator(
        input_type=SpecValidators.str_to_bool,
        description="Use managed spot training.",
        default=False,
    ),
    max_wait_time=SageMakerComponentInputValidator(
        input_type=int,
        description="The maximum time in seconds you are willing to wait for a managed spot training job to complete.",
        default=86400,
    ),
    checkpoint_config=SageMakerComponentInputValidator(
        input_type=SpecValidators.yaml_or_json_dict,
        description="Dictionary of information about the output location for managed spot training checkpoint data.",
        default={},
    ),
)