# flake8: noqa

# import all models into this package
# if you have many models here with many references from one model to another this may
# raise a RecursionError
# to avoid this, import only the models that you directly need like:
# from from nomad_client.model.pet import Pet
# or import this package, but before doing it, use:
# import sys
# sys.setrecursionlimit(n)

from nomad_client.model.acl_policy import ACLPolicy
from nomad_client.model.acl_policy_list_stub import ACLPolicyListStub
from nomad_client.model.acl_token import ACLToken
from nomad_client.model.acl_token_list_stub import ACLTokenListStub
from nomad_client.model.affinity import Affinity
from nomad_client.model.alloc_deployment_status import AllocDeploymentStatus
from nomad_client.model.allocated_cpu_resources import AllocatedCpuResources
from nomad_client.model.allocated_device_resource import AllocatedDeviceResource
from nomad_client.model.allocated_memory_resources import AllocatedMemoryResources
from nomad_client.model.allocated_resources import AllocatedResources
from nomad_client.model.allocated_shared_resources import AllocatedSharedResources
from nomad_client.model.allocated_task_resources import AllocatedTaskResources
from nomad_client.model.allocation import Allocation
from nomad_client.model.allocation_list_stub import AllocationListStub
from nomad_client.model.allocation_metric import AllocationMetric
<<<<<<< HEAD
<<<<<<< HEAD
from nomad_client.model.attribute import Attribute
from nomad_client.model.csi_controller_info import CSIControllerInfo
from nomad_client.model.csi_info import CSIInfo
=======
from nomad_client.model.autopilot_configuration import AutopilotConfiguration
>>>>>>> 2fdd42e (updates to generator/operator.go)
=======
from nomad_client.model.autopilot_configuration import AutopilotConfiguration
>>>>>>> 4825049f87bd05f944cf6444d16c4a85fed43b46
from nomad_client.model.csi_mount_options import CSIMountOptions
from nomad_client.model.csi_node_info import CSINodeInfo
from nomad_client.model.csi_plugin import CSIPlugin
from nomad_client.model.csi_plugin_list_stub import CSIPluginListStub
from nomad_client.model.csi_secrets import CSISecrets
from nomad_client.model.csi_snapshot import CSISnapshot
from nomad_client.model.csi_snapshot_create_request import CSISnapshotCreateRequest
from nomad_client.model.csi_snapshot_create_response import CSISnapshotCreateResponse
from nomad_client.model.csi_snapshot_list_response import CSISnapshotListResponse
from nomad_client.model.csi_topology import CSITopology
from nomad_client.model.csi_volume import CSIVolume
from nomad_client.model.csi_volume_capability import CSIVolumeCapability
from nomad_client.model.csi_volume_create_request import CSIVolumeCreateRequest
from nomad_client.model.csi_volume_external_stub import CSIVolumeExternalStub
from nomad_client.model.csi_volume_list_external_response import CSIVolumeListExternalResponse
from nomad_client.model.csi_volume_list_stub import CSIVolumeListStub
from nomad_client.model.csi_volume_register_request import CSIVolumeRegisterRequest
from nomad_client.model.check_restart import CheckRestart
from nomad_client.model.constraint import Constraint
from nomad_client.model.consul import Consul
from nomad_client.model.consul_connect import ConsulConnect
from nomad_client.model.consul_expose_config import ConsulExposeConfig
from nomad_client.model.consul_expose_path import ConsulExposePath
from nomad_client.model.consul_gateway import ConsulGateway
from nomad_client.model.consul_gateway_bind_address import ConsulGatewayBindAddress
from nomad_client.model.consul_gateway_proxy import ConsulGatewayProxy
from nomad_client.model.consul_gateway_tls_config import ConsulGatewayTLSConfig
from nomad_client.model.consul_ingress_config_entry import ConsulIngressConfigEntry
from nomad_client.model.consul_ingress_listener import ConsulIngressListener
from nomad_client.model.consul_ingress_service import ConsulIngressService
from nomad_client.model.consul_linked_service import ConsulLinkedService
from nomad_client.model.consul_mesh_gateway import ConsulMeshGateway
from nomad_client.model.consul_proxy import ConsulProxy
from nomad_client.model.consul_sidecar_service import ConsulSidecarService
from nomad_client.model.consul_terminating_config_entry import ConsulTerminatingConfigEntry
from nomad_client.model.consul_upstream import ConsulUpstream
from nomad_client.model.dns_config import DNSConfig
from nomad_client.model.deployment import Deployment
from nomad_client.model.deployment_alloc_health_request import DeploymentAllocHealthRequest
from nomad_client.model.deployment_pause_request import DeploymentPauseRequest
from nomad_client.model.deployment_promote_request import DeploymentPromoteRequest
from nomad_client.model.deployment_state import DeploymentState
from nomad_client.model.deployment_unblock_request import DeploymentUnblockRequest
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
from nomad_client.model.desired_transition import DesiredTransition
from nomad_client.model.desired_updates import DesiredUpdates
from nomad_client.model.dispatch_payload_config import DispatchPayloadConfig
from nomad_client.model.drain_metadata import DrainMetadata
from nomad_client.model.drain_spec import DrainSpec
from nomad_client.model.drain_strategy import DrainStrategy
from nomad_client.model.driver_info import DriverInfo
from nomad_client.model.ephemeral_disk import EphemeralDisk
from nomad_client.model.eval_options import EvalOptions
from nomad_client.model.evaluation import Evaluation
from nomad_client.model.field_diff import FieldDiff
from nomad_client.model.fuzzy_match import FuzzyMatch
from nomad_client.model.fuzzy_search_request import FuzzySearchRequest
from nomad_client.model.fuzzy_search_response import FuzzySearchResponse
from nomad_client.model.gauge_value import GaugeValue
from nomad_client.model.host_volume_info import HostVolumeInfo
from nomad_client.model.int8 import Int8
from nomad_client.model.job import Job
from nomad_client.model.job_children_summary import JobChildrenSummary
from nomad_client.model.job_deregister_response import JobDeregisterResponse
from nomad_client.model.job_diff import JobDiff
from nomad_client.model.job_dispatch_request import JobDispatchRequest
from nomad_client.model.job_dispatch_response import JobDispatchResponse
from nomad_client.model.job_evaluate_request import JobEvaluateRequest
from nomad_client.model.job_list_stub import JobListStub
from nomad_client.model.job_plan_request import JobPlanRequest
from nomad_client.model.job_plan_response import JobPlanResponse
from nomad_client.model.job_register_request import JobRegisterRequest
from nomad_client.model.job_register_response import JobRegisterResponse
from nomad_client.model.job_revert_request import JobRevertRequest
from nomad_client.model.job_scale_status_response import JobScaleStatusResponse
from nomad_client.model.job_stability_request import JobStabilityRequest
from nomad_client.model.job_stability_response import JobStabilityResponse
from nomad_client.model.job_summary import JobSummary
from nomad_client.model.job_validate_request import JobValidateRequest
from nomad_client.model.job_validate_response import JobValidateResponse
from nomad_client.model.job_versions_response import JobVersionsResponse
from nomad_client.model.jobs_parse_request import JobsParseRequest
from nomad_client.model.log_config import LogConfig
from nomad_client.model.metrics_summary import MetricsSummary
from nomad_client.model.migrate_strategy import MigrateStrategy
from nomad_client.model.multiregion import Multiregion
from nomad_client.model.multiregion_region import MultiregionRegion
from nomad_client.model.multiregion_strategy import MultiregionStrategy
from nomad_client.model.namespace import Namespace
from nomad_client.model.network_resource import NetworkResource
from nomad_client.model.node import Node
from nomad_client.model.node_cpu_resources import NodeCpuResources
from nomad_client.model.node_device import NodeDevice
from nomad_client.model.node_device_locality import NodeDeviceLocality
from nomad_client.model.node_device_resource import NodeDeviceResource
from nomad_client.model.node_disk_resources import NodeDiskResources
from nomad_client.model.node_drain_update_response import NodeDrainUpdateResponse
from nomad_client.model.node_eligibility_update_response import NodeEligibilityUpdateResponse
from nomad_client.model.node_event import NodeEvent
from nomad_client.model.node_list_stub import NodeListStub
from nomad_client.model.node_memory_resources import NodeMemoryResources
from nomad_client.model.node_purge_response import NodePurgeResponse
from nomad_client.model.node_reserved_cpu_resources import NodeReservedCpuResources
from nomad_client.model.node_reserved_disk_resources import NodeReservedDiskResources
from nomad_client.model.node_reserved_memory_resources import NodeReservedMemoryResources
from nomad_client.model.node_reserved_network_resources import NodeReservedNetworkResources
from nomad_client.model.node_reserved_resources import NodeReservedResources
from nomad_client.model.node_resources import NodeResources
from nomad_client.model.node_score_meta import NodeScoreMeta
from nomad_client.model.node_update_drain_request import NodeUpdateDrainRequest
from nomad_client.model.node_update_eligibility_request import NodeUpdateEligibilityRequest
from nomad_client.model.object_diff import ObjectDiff
<<<<<<< HEAD
<<<<<<< HEAD
from nomad_client.model.one_time_token import OneTimeToken
from nomad_client.model.one_time_token_exchange_request import OneTimeTokenExchangeRequest
=======
from nomad_client.model.operator_health_reply import OperatorHealthReply
>>>>>>> 2fdd42e (updates to generator/operator.go)
=======
from nomad_client.model.operator_health_reply import OperatorHealthReply
>>>>>>> 4825049f87bd05f944cf6444d16c4a85fed43b46
from nomad_client.model.parameterized_job_config import ParameterizedJobConfig
from nomad_client.model.periodic_config import PeriodicConfig
from nomad_client.model.periodic_force_response import PeriodicForceResponse
from nomad_client.model.plan_annotations import PlanAnnotations
from nomad_client.model.point_value import PointValue
from nomad_client.model.port import Port
from nomad_client.model.port_mapping import PortMapping
from nomad_client.model.preemption_config import PreemptionConfig
from nomad_client.model.quota_limit import QuotaLimit
from nomad_client.model.quota_spec import QuotaSpec
from nomad_client.model.raft_server import RaftServer
from nomad_client.model.requested_device import RequestedDevice
from nomad_client.model.reschedule_event import RescheduleEvent
from nomad_client.model.reschedule_policy import ReschedulePolicy
from nomad_client.model.reschedule_tracker import RescheduleTracker
from nomad_client.model.resources import Resources
from nomad_client.model.restart_policy import RestartPolicy
from nomad_client.model.sampled_value import SampledValue
from nomad_client.model.scaling_event import ScalingEvent
from nomad_client.model.scaling_policy import ScalingPolicy
from nomad_client.model.scaling_policy_list_stub import ScalingPolicyListStub
from nomad_client.model.scaling_request import ScalingRequest
from nomad_client.model.scheduler_configuration import SchedulerConfiguration
from nomad_client.model.scheduler_configuration_response import SchedulerConfigurationResponse
from nomad_client.model.scheduler_set_configuration_response import SchedulerSetConfigurationResponse
from nomad_client.model.search_request import SearchRequest
from nomad_client.model.search_response import SearchResponse
from nomad_client.model.server_health import ServerHealth
from nomad_client.model.service import Service
from nomad_client.model.service_check import ServiceCheck
from nomad_client.model.sidecar_task import SidecarTask
from nomad_client.model.spread import Spread
from nomad_client.model.spread_target import SpreadTarget
from nomad_client.model.task import Task
from nomad_client.model.task_artifact import TaskArtifact
from nomad_client.model.task_csi_plugin_config import TaskCSIPluginConfig
from nomad_client.model.task_diff import TaskDiff
from nomad_client.model.task_event import TaskEvent
from nomad_client.model.task_group import TaskGroup
from nomad_client.model.task_group_diff import TaskGroupDiff
from nomad_client.model.task_group_scale_status import TaskGroupScaleStatus
from nomad_client.model.task_group_summary import TaskGroupSummary
from nomad_client.model.task_handle import TaskHandle
from nomad_client.model.task_lifecycle import TaskLifecycle
from nomad_client.model.task_state import TaskState
from nomad_client.model.template import Template
<<<<<<< HEAD
<<<<<<< HEAD
from nomad_client.model.uint16 import Uint16
=======
from nomad_client.model.uint import Uint
>>>>>>> 2fdd42e (updates to generator/operator.go)
=======
from nomad_client.model.uint import Uint
>>>>>>> 4825049f87bd05f944cf6444d16c4a85fed43b46
from nomad_client.model.uint64 import Uint64
from nomad_client.model.uint8 import Uint8
from nomad_client.model.update_strategy import UpdateStrategy
from nomad_client.model.vault import Vault
from nomad_client.model.volume_mount import VolumeMount
from nomad_client.model.volume_request import VolumeRequest
