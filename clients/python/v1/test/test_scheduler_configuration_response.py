# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import nomad_client
from nomad_client.model.scheduler_configuration import SchedulerConfiguration
globals()['SchedulerConfiguration'] = SchedulerConfiguration
from nomad_client.model.scheduler_configuration_response import SchedulerConfigurationResponse


class TestSchedulerConfigurationResponse(unittest.TestCase):
    """SchedulerConfigurationResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSchedulerConfigurationResponse(self):
        """Test SchedulerConfigurationResponse"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SchedulerConfigurationResponse()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
