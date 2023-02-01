# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.3
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import nomad_client
from nomad_client.model.field_diff import FieldDiff
from nomad_client.model.object_diff import ObjectDiff
from nomad_client.model.task_group_diff import TaskGroupDiff
globals()['FieldDiff'] = FieldDiff
globals()['ObjectDiff'] = ObjectDiff
globals()['TaskGroupDiff'] = TaskGroupDiff
from nomad_client.model.job_diff import JobDiff


class TestJobDiff(unittest.TestCase):
    """JobDiff unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testJobDiff(self):
        """Test JobDiff"""
        # FIXME: construct object with mandatory attributes with example values
        # model = JobDiff()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
