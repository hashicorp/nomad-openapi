"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import nomad_client
from nomad_client.api.scaling_api import ScalingApi  # noqa: E501


class TestScalingApi(unittest.TestCase):
    """ScalingApi unit test stubs"""

    def setUp(self):
        self.api = ScalingApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_scaling_policies(self):
        """Test case for get_scaling_policies

        """
        pass

    def test_get_scaling_policy(self):
        """Test case for get_scaling_policy

        """
        pass


if __name__ == '__main__':
    unittest.main()