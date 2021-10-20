"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.affinity import Affinity
from openapi_client.model.constraint import Constraint
from openapi_client.model.migrate_strategy import MigrateStrategy
from openapi_client.model.multiregion import Multiregion
from openapi_client.model.parameterized_job_config import ParameterizedJobConfig
from openapi_client.model.periodic_config import PeriodicConfig
from openapi_client.model.reschedule_policy import ReschedulePolicy
from openapi_client.model.spread import Spread
from openapi_client.model.task_group import TaskGroup
from openapi_client.model.update_strategy import UpdateStrategy
globals()['Affinity'] = Affinity
globals()['Constraint'] = Constraint
globals()['MigrateStrategy'] = MigrateStrategy
globals()['Multiregion'] = Multiregion
globals()['ParameterizedJobConfig'] = ParameterizedJobConfig
globals()['PeriodicConfig'] = PeriodicConfig
globals()['ReschedulePolicy'] = ReschedulePolicy
globals()['Spread'] = Spread
globals()['TaskGroup'] = TaskGroup
globals()['UpdateStrategy'] = UpdateStrategy
from openapi_client.model.job import Job


class TestJob(unittest.TestCase):
    """Job unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testJob(self):
        """Test Job"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Job()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
