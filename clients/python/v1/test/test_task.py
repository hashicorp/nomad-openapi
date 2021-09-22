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
from nomad_client.model.affinity import Affinity
from nomad_client.model.constraint import Constraint
from nomad_client.model.dispatch_payload_config import DispatchPayloadConfig
from nomad_client.model.log_config import LogConfig
from nomad_client.model.resources import Resources
from nomad_client.model.restart_policy import RestartPolicy
from nomad_client.model.scaling_policy import ScalingPolicy
from nomad_client.model.service import Service
from nomad_client.model.task_artifact import TaskArtifact
from nomad_client.model.task_csi_plugin_config import TaskCSIPluginConfig
from nomad_client.model.task_lifecycle import TaskLifecycle
from nomad_client.model.template import Template
from nomad_client.model.vault import Vault
from nomad_client.model.volume_mount import VolumeMount
globals()['Affinity'] = Affinity
globals()['Constraint'] = Constraint
globals()['DispatchPayloadConfig'] = DispatchPayloadConfig
globals()['LogConfig'] = LogConfig
globals()['Resources'] = Resources
globals()['RestartPolicy'] = RestartPolicy
globals()['ScalingPolicy'] = ScalingPolicy
globals()['Service'] = Service
globals()['TaskArtifact'] = TaskArtifact
globals()['TaskCSIPluginConfig'] = TaskCSIPluginConfig
globals()['TaskLifecycle'] = TaskLifecycle
globals()['Template'] = Template
globals()['Vault'] = Vault
globals()['VolumeMount'] = VolumeMount
from nomad_client.model.task import Task


class TestTask(unittest.TestCase):
    """Task unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testTask(self):
        """Test Task"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Task()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()