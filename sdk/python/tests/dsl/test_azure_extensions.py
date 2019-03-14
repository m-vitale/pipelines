# Copyright 2019 The Kubeflow Authors
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

from kfp.dsl import Pipeline, ContainerOp
from kfp.az import use_az_secret
import unittest
import inspect

class AzExtensionTests(unittest.TestCase):
    def test_default_secret_name(self):
        spec = inspect.getfullargspec(use_az_secret)
        assert len(spec.defaults) == 1
        assert spec.defaults[0] == 'azcreds'

    def test_use_az_secret(self):
        with Pipeline('somename') as p:
            op1 = ContainerOp(name='op1', image='image')
            op1 = op1.apply(use_az_secret('foo'))
            assert len(op1.env_variables) == 4
            
            index = 0
            for expected in ['AZ_SUBSCRIPTION_ID', 'AZ_TENANT_ID', 'AZ_CLIENT_ID', 'AZ_CLIENT_SECRET']:
                assert op1.env_variables[index].name == expected
                assert op1.env_variables[index].value_from.secret_key_ref.name == 'foo'
                assert op1.env_variables[index].value_from.secret_key_ref.key == expected
                index += 1

if __name__ == '__main__':
    unittest.main()
