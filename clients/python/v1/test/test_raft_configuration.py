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
from nomad_client.model.raft_server import RaftServer
globals()['RaftServer'] = RaftServer
from nomad_client.model.raft_configuration import RaftConfiguration


class TestRaftConfiguration(unittest.TestCase):
    """RaftConfiguration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testRaftConfiguration(self):
        """Test RaftConfiguration"""
        # FIXME: construct object with mandatory attributes with example values
        # model = RaftConfiguration()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
