import json
import unittest

from unittest.mock import patch, call, Mock, MagicMock, mock_open
from botocore.exceptions import ClientError
from datetime import datetime

from model.src import create_model
from common import _utils
from . import test_utils


required_args = [
  '--region', 'us-west-2',
  '--model_name', 'model_test',
  '--role', 'arn:aws:iam::123456789012:user/Development/product_1234/*',
  '--image', 'test-image',
  '--model_artifact_url', 's3://fake-bucket/model_artifact'
]

class ModelTestCase(unittest.TestCase):
  @classmethod
  def setUpClass(cls):
    parser = create_model.create_parser()
    cls.parser = parser

  def test_create_parser(self):
    self.assertIsNotNone(self.parser)

  def test_main(self):
    # Mock out all of utils except parser
    create_model._utils = MagicMock()
    create_model._utils.add_default_client_arguments = _utils.add_default_client_arguments

    # Set some static returns
    create_model._utils.create_model.return_value = 'model_test'

    with patch('builtins.open', mock_open()) as file_open:
      create_model.main(required_args)

    # Check if correct requests were created and triggered
    create_model._utils.create_model.assert_called()

    # Check the file outputs
    file_open.assert_has_calls([
      call('/tmp/model_name.txt', 'w')
    ])

    file_open().write.assert_has_calls([
      call('model_test')
    ])

  def test_create_model(self):
    mock_client = MagicMock()
    mock_args = self.parser.parse_args(required_args)
    response = _utils.create_model(mock_client, vars(mock_args))

    mock_client.create_model.assert_called_once_with(
      EnableNetworkIsolation=True,
      ExecutionRoleArn='arn:aws:iam::123456789012:user/Development/product_1234/*',
      ModelName='model_test',
      PrimaryContainer={'Image': 'test-image', 'ModelDataUrl': 's3://fake-bucket/model_artifact', 'Environment': {}},
      Tags=[]
    )


  def test_sagemaker_exception_in_create_model(self):
    mock_client = MagicMock()
    mock_exception = ClientError({"Error": {"Message": "SageMaker broke"}}, "create_model")
    mock_client.create_model.side_effect = mock_exception
    mock_args = self.parser.parse_args(required_args)

    with self.assertRaises(Exception):
      _utils.create_model(mock_client, vars(mock_args))
