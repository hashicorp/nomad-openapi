/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from './ApiClient';
import ACLPolicy from './model/ACLPolicy';
import ACLPolicyListStub from './model/ACLPolicyListStub';
import ACLToken from './model/ACLToken';
import ACLTokenListStub from './model/ACLTokenListStub';
import Affinity from './model/Affinity';
import AllocDeploymentStatus from './model/AllocDeploymentStatus';
import AllocatedCpuResources from './model/AllocatedCpuResources';
import AllocatedDeviceResource from './model/AllocatedDeviceResource';
import AllocatedMemoryResources from './model/AllocatedMemoryResources';
import AllocatedResources from './model/AllocatedResources';
import AllocatedSharedResources from './model/AllocatedSharedResources';
import AllocatedTaskResources from './model/AllocatedTaskResources';
import Allocation from './model/Allocation';
import AllocationListStub from './model/AllocationListStub';
import AllocationMetric from './model/AllocationMetric';
import Attribute from './model/Attribute';
import CSIControllerInfo from './model/CSIControllerInfo';
import CSIInfo from './model/CSIInfo';
import CSIMountOptions from './model/CSIMountOptions';
import CSINodeInfo from './model/CSINodeInfo';
import CSIPlugin from './model/CSIPlugin';
import CSIPluginListStub from './model/CSIPluginListStub';
import CSISnapshot from './model/CSISnapshot';
import CSISnapshotCreateRequest from './model/CSISnapshotCreateRequest';
import CSISnapshotCreateResponse from './model/CSISnapshotCreateResponse';
import CSISnapshotListResponse from './model/CSISnapshotListResponse';
import CSITopology from './model/CSITopology';
import CSIVolume from './model/CSIVolume';
import CSIVolumeCapability from './model/CSIVolumeCapability';
import CSIVolumeCreateRequest from './model/CSIVolumeCreateRequest';
import CSIVolumeExternalStub from './model/CSIVolumeExternalStub';
import CSIVolumeListExternalResponse from './model/CSIVolumeListExternalResponse';
import CSIVolumeListStub from './model/CSIVolumeListStub';
import CSIVolumeRegisterRequest from './model/CSIVolumeRegisterRequest';
import CheckRestart from './model/CheckRestart';
import Constraint from './model/Constraint';
import Consul from './model/Consul';
import ConsulConnect from './model/ConsulConnect';
import ConsulExposeConfig from './model/ConsulExposeConfig';
import ConsulExposePath from './model/ConsulExposePath';
import ConsulGateway from './model/ConsulGateway';
import ConsulGatewayBindAddress from './model/ConsulGatewayBindAddress';
import ConsulGatewayProxy from './model/ConsulGatewayProxy';
import ConsulGatewayTLSConfig from './model/ConsulGatewayTLSConfig';
import ConsulIngressConfigEntry from './model/ConsulIngressConfigEntry';
import ConsulIngressListener from './model/ConsulIngressListener';
import ConsulIngressService from './model/ConsulIngressService';
import ConsulLinkedService from './model/ConsulLinkedService';
import ConsulMeshGateway from './model/ConsulMeshGateway';
import ConsulProxy from './model/ConsulProxy';
import ConsulSidecarService from './model/ConsulSidecarService';
import ConsulTerminatingConfigEntry from './model/ConsulTerminatingConfigEntry';
import ConsulUpstream from './model/ConsulUpstream';
import DNSConfig from './model/DNSConfig';
import Deployment from './model/Deployment';
import DeploymentAllocHealthRequest from './model/DeploymentAllocHealthRequest';
import DeploymentPauseRequest from './model/DeploymentPauseRequest';
import DeploymentPromoteRequest from './model/DeploymentPromoteRequest';
import DeploymentState from './model/DeploymentState';
import DeploymentUnblockRequest from './model/DeploymentUnblockRequest';
import DeploymentUpdateResponse from './model/DeploymentUpdateResponse';
import DesiredTransition from './model/DesiredTransition';
import DesiredUpdates from './model/DesiredUpdates';
import DispatchPayloadConfig from './model/DispatchPayloadConfig';
import DrainMetadata from './model/DrainMetadata';
import DrainSpec from './model/DrainSpec';
import DrainStrategy from './model/DrainStrategy';
import DriverInfo from './model/DriverInfo';
import EphemeralDisk from './model/EphemeralDisk';
import EvalOptions from './model/EvalOptions';
import Evaluation from './model/Evaluation';
import FieldDiff from './model/FieldDiff';
import FuzzyMatch from './model/FuzzyMatch';
import FuzzySearchRequest from './model/FuzzySearchRequest';
import FuzzySearchResponse from './model/FuzzySearchResponse';
import GaugeValue from './model/GaugeValue';
import HostVolumeInfo from './model/HostVolumeInfo';
import Job from './model/Job';
import JobChildrenSummary from './model/JobChildrenSummary';
import JobDeregisterResponse from './model/JobDeregisterResponse';
import JobDiff from './model/JobDiff';
import JobDispatchRequest from './model/JobDispatchRequest';
import JobDispatchResponse from './model/JobDispatchResponse';
import JobEvaluateRequest from './model/JobEvaluateRequest';
import JobListStub from './model/JobListStub';
import JobPlanRequest from './model/JobPlanRequest';
import JobPlanResponse from './model/JobPlanResponse';
import JobRegisterRequest from './model/JobRegisterRequest';
import JobRegisterResponse from './model/JobRegisterResponse';
import JobRevertRequest from './model/JobRevertRequest';
import JobScaleStatusResponse from './model/JobScaleStatusResponse';
import JobStabilityRequest from './model/JobStabilityRequest';
import JobStabilityResponse from './model/JobStabilityResponse';
import JobSummary from './model/JobSummary';
import JobValidateRequest from './model/JobValidateRequest';
import JobValidateResponse from './model/JobValidateResponse';
import JobVersionsResponse from './model/JobVersionsResponse';
import JobsParseRequest from './model/JobsParseRequest';
import LogConfig from './model/LogConfig';
import MetricsSummary from './model/MetricsSummary';
import MigrateStrategy from './model/MigrateStrategy';
import Multiregion from './model/Multiregion';
import MultiregionRegion from './model/MultiregionRegion';
import MultiregionStrategy from './model/MultiregionStrategy';
import Namespace from './model/Namespace';
import NetworkResource from './model/NetworkResource';
import Node from './model/Node';
import NodeCpuResources from './model/NodeCpuResources';
import NodeDevice from './model/NodeDevice';
import NodeDeviceLocality from './model/NodeDeviceLocality';
import NodeDeviceResource from './model/NodeDeviceResource';
import NodeDiskResources from './model/NodeDiskResources';
import NodeDrainUpdateResponse from './model/NodeDrainUpdateResponse';
import NodeEligibilityUpdateResponse from './model/NodeEligibilityUpdateResponse';
import NodeEvent from './model/NodeEvent';
import NodeListStub from './model/NodeListStub';
import NodeMemoryResources from './model/NodeMemoryResources';
import NodePurgeResponse from './model/NodePurgeResponse';
import NodeReservedCpuResources from './model/NodeReservedCpuResources';
import NodeReservedDiskResources from './model/NodeReservedDiskResources';
import NodeReservedMemoryResources from './model/NodeReservedMemoryResources';
import NodeReservedNetworkResources from './model/NodeReservedNetworkResources';
import NodeReservedResources from './model/NodeReservedResources';
import NodeResources from './model/NodeResources';
import NodeScoreMeta from './model/NodeScoreMeta';
import NodeUpdateDrainRequest from './model/NodeUpdateDrainRequest';
import NodeUpdateEligibilityRequest from './model/NodeUpdateEligibilityRequest';
import ObjectDiff from './model/ObjectDiff';
import OneTimeToken from './model/OneTimeToken';
import OneTimeTokenExchangeRequest from './model/OneTimeTokenExchangeRequest';
import ParameterizedJobConfig from './model/ParameterizedJobConfig';
import PeriodicConfig from './model/PeriodicConfig';
import PeriodicForceResponse from './model/PeriodicForceResponse';
import PlanAnnotations from './model/PlanAnnotations';
import PointValue from './model/PointValue';
import Port from './model/Port';
import PortMapping from './model/PortMapping';
import QuotaLimit from './model/QuotaLimit';
import QuotaSpec from './model/QuotaSpec';
import RequestedDevice from './model/RequestedDevice';
import RescheduleEvent from './model/RescheduleEvent';
import ReschedulePolicy from './model/ReschedulePolicy';
import RescheduleTracker from './model/RescheduleTracker';
import Resources from './model/Resources';
import RestartPolicy from './model/RestartPolicy';
import SampledValue from './model/SampledValue';
import ScalingEvent from './model/ScalingEvent';
import ScalingPolicy from './model/ScalingPolicy';
import ScalingPolicyListStub from './model/ScalingPolicyListStub';
import ScalingRequest from './model/ScalingRequest';
import SearchRequest from './model/SearchRequest';
import SearchResponse from './model/SearchResponse';
import Service from './model/Service';
import ServiceCheck from './model/ServiceCheck';
import SidecarTask from './model/SidecarTask';
import Spread from './model/Spread';
import SpreadTarget from './model/SpreadTarget';
import Task from './model/Task';
import TaskArtifact from './model/TaskArtifact';
import TaskCSIPluginConfig from './model/TaskCSIPluginConfig';
import TaskDiff from './model/TaskDiff';
import TaskEvent from './model/TaskEvent';
import TaskGroup from './model/TaskGroup';
import TaskGroupDiff from './model/TaskGroupDiff';
import TaskGroupScaleStatus from './model/TaskGroupScaleStatus';
import TaskGroupSummary from './model/TaskGroupSummary';
import TaskHandle from './model/TaskHandle';
import TaskLifecycle from './model/TaskLifecycle';
import TaskState from './model/TaskState';
import Template from './model/Template';
import UpdateStrategy from './model/UpdateStrategy';
import Vault from './model/Vault';
import VolumeMount from './model/VolumeMount';
import VolumeRequest from './model/VolumeRequest';
import ACLApi from './api/ACLApi';
import AllocationsApi from './api/AllocationsApi';
import DeploymentsApi from './api/DeploymentsApi';
import EnterpriseApi from './api/EnterpriseApi';
import EvaluationsApi from './api/EvaluationsApi';
import JobsApi from './api/JobsApi';
import MetricsApi from './api/MetricsApi';
import NamespacesApi from './api/NamespacesApi';
import NodesApi from './api/NodesApi';
import PluginsApi from './api/PluginsApi';
import RegionsApi from './api/RegionsApi';
import ScalingApi from './api/ScalingApi';
import SearchApi from './api/SearchApi';
import StatusApi from './api/StatusApi';
<<<<<<< HEAD
>>>>>>> 8d74053 (updated generator/status.go)
=======
import SystemApi from './api/SystemApi';
>>>>>>> e3fb144 (added v1/status.go and v1/status_test.go, tests passing)
import VolumesApi from './api/VolumesApi';


/**
* JS API client generated by OpenAPI Generator.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var nomad-client = require('index'); // See note below*.
* var xxxSvc = new nomad-client.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new nomad-client.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new nomad-client.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new nomad-client.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 1.1.4
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The ACLPolicy model constructor.
     * @property {module:model/ACLPolicy}
     */
    ACLPolicy,

    /**
     * The ACLPolicyListStub model constructor.
     * @property {module:model/ACLPolicyListStub}
     */
    ACLPolicyListStub,

    /**
     * The ACLToken model constructor.
     * @property {module:model/ACLToken}
     */
    ACLToken,

    /**
     * The ACLTokenListStub model constructor.
     * @property {module:model/ACLTokenListStub}
     */
    ACLTokenListStub,

    /**
     * The Affinity model constructor.
     * @property {module:model/Affinity}
     */
    Affinity,

    /**
     * The AllocDeploymentStatus model constructor.
     * @property {module:model/AllocDeploymentStatus}
     */
    AllocDeploymentStatus,

    /**
     * The AllocatedCpuResources model constructor.
     * @property {module:model/AllocatedCpuResources}
     */
    AllocatedCpuResources,

    /**
     * The AllocatedDeviceResource model constructor.
     * @property {module:model/AllocatedDeviceResource}
     */
    AllocatedDeviceResource,

    /**
     * The AllocatedMemoryResources model constructor.
     * @property {module:model/AllocatedMemoryResources}
     */
    AllocatedMemoryResources,

    /**
     * The AllocatedResources model constructor.
     * @property {module:model/AllocatedResources}
     */
    AllocatedResources,

    /**
     * The AllocatedSharedResources model constructor.
     * @property {module:model/AllocatedSharedResources}
     */
    AllocatedSharedResources,

    /**
     * The AllocatedTaskResources model constructor.
     * @property {module:model/AllocatedTaskResources}
     */
    AllocatedTaskResources,

    /**
     * The Allocation model constructor.
     * @property {module:model/Allocation}
     */
    Allocation,

    /**
     * The AllocationListStub model constructor.
     * @property {module:model/AllocationListStub}
     */
    AllocationListStub,

    /**
     * The AllocationMetric model constructor.
     * @property {module:model/AllocationMetric}
     */
    AllocationMetric,

    /**
     * The Attribute model constructor.
     * @property {module:model/Attribute}
     */
    Attribute,

    /**
     * The CSIControllerInfo model constructor.
     * @property {module:model/CSIControllerInfo}
     */
    CSIControllerInfo,

    /**
     * The CSIInfo model constructor.
     * @property {module:model/CSIInfo}
     */
    CSIInfo,

    /**
     * The CSIMountOptions model constructor.
     * @property {module:model/CSIMountOptions}
     */
    CSIMountOptions,

    /**
     * The CSINodeInfo model constructor.
     * @property {module:model/CSINodeInfo}
     */
    CSINodeInfo,

    /**
     * The CSIPlugin model constructor.
     * @property {module:model/CSIPlugin}
     */
    CSIPlugin,

    /**
     * The CSIPluginListStub model constructor.
     * @property {module:model/CSIPluginListStub}
     */
    CSIPluginListStub,

    /**
     * The CSISnapshot model constructor.
     * @property {module:model/CSISnapshot}
     */
    CSISnapshot,

    /**
     * The CSISnapshotCreateRequest model constructor.
     * @property {module:model/CSISnapshotCreateRequest}
     */
    CSISnapshotCreateRequest,

    /**
     * The CSISnapshotCreateResponse model constructor.
     * @property {module:model/CSISnapshotCreateResponse}
     */
    CSISnapshotCreateResponse,

    /**
     * The CSISnapshotListResponse model constructor.
     * @property {module:model/CSISnapshotListResponse}
     */
    CSISnapshotListResponse,

    /**
     * The CSITopology model constructor.
     * @property {module:model/CSITopology}
     */
    CSITopology,

    /**
     * The CSIVolume model constructor.
     * @property {module:model/CSIVolume}
     */
    CSIVolume,

    /**
     * The CSIVolumeCapability model constructor.
     * @property {module:model/CSIVolumeCapability}
     */
    CSIVolumeCapability,

    /**
     * The CSIVolumeCreateRequest model constructor.
     * @property {module:model/CSIVolumeCreateRequest}
     */
    CSIVolumeCreateRequest,

    /**
     * The CSIVolumeExternalStub model constructor.
     * @property {module:model/CSIVolumeExternalStub}
     */
    CSIVolumeExternalStub,

    /**
     * The CSIVolumeListExternalResponse model constructor.
     * @property {module:model/CSIVolumeListExternalResponse}
     */
    CSIVolumeListExternalResponse,

    /**
     * The CSIVolumeListStub model constructor.
     * @property {module:model/CSIVolumeListStub}
     */
    CSIVolumeListStub,

    /**
     * The CSIVolumeRegisterRequest model constructor.
     * @property {module:model/CSIVolumeRegisterRequest}
     */
    CSIVolumeRegisterRequest,

    /**
     * The CheckRestart model constructor.
     * @property {module:model/CheckRestart}
     */
    CheckRestart,

    /**
     * The Constraint model constructor.
     * @property {module:model/Constraint}
     */
    Constraint,

    /**
     * The Consul model constructor.
     * @property {module:model/Consul}
     */
    Consul,

    /**
     * The ConsulConnect model constructor.
     * @property {module:model/ConsulConnect}
     */
    ConsulConnect,

    /**
     * The ConsulExposeConfig model constructor.
     * @property {module:model/ConsulExposeConfig}
     */
    ConsulExposeConfig,

    /**
     * The ConsulExposePath model constructor.
     * @property {module:model/ConsulExposePath}
     */
    ConsulExposePath,

    /**
     * The ConsulGateway model constructor.
     * @property {module:model/ConsulGateway}
     */
    ConsulGateway,

    /**
     * The ConsulGatewayBindAddress model constructor.
     * @property {module:model/ConsulGatewayBindAddress}
     */
    ConsulGatewayBindAddress,

    /**
     * The ConsulGatewayProxy model constructor.
     * @property {module:model/ConsulGatewayProxy}
     */
    ConsulGatewayProxy,

    /**
     * The ConsulGatewayTLSConfig model constructor.
     * @property {module:model/ConsulGatewayTLSConfig}
     */
    ConsulGatewayTLSConfig,

    /**
     * The ConsulIngressConfigEntry model constructor.
     * @property {module:model/ConsulIngressConfigEntry}
     */
    ConsulIngressConfigEntry,

    /**
     * The ConsulIngressListener model constructor.
     * @property {module:model/ConsulIngressListener}
     */
    ConsulIngressListener,

    /**
     * The ConsulIngressService model constructor.
     * @property {module:model/ConsulIngressService}
     */
    ConsulIngressService,

    /**
     * The ConsulLinkedService model constructor.
     * @property {module:model/ConsulLinkedService}
     */
    ConsulLinkedService,

    /**
     * The ConsulMeshGateway model constructor.
     * @property {module:model/ConsulMeshGateway}
     */
    ConsulMeshGateway,

    /**
     * The ConsulProxy model constructor.
     * @property {module:model/ConsulProxy}
     */
    ConsulProxy,

    /**
     * The ConsulSidecarService model constructor.
     * @property {module:model/ConsulSidecarService}
     */
    ConsulSidecarService,

    /**
     * The ConsulTerminatingConfigEntry model constructor.
     * @property {module:model/ConsulTerminatingConfigEntry}
     */
    ConsulTerminatingConfigEntry,

    /**
     * The ConsulUpstream model constructor.
     * @property {module:model/ConsulUpstream}
     */
    ConsulUpstream,

    /**
     * The DNSConfig model constructor.
     * @property {module:model/DNSConfig}
     */
    DNSConfig,

    /**
     * The Deployment model constructor.
     * @property {module:model/Deployment}
     */
    Deployment,

    /**
     * The DeploymentAllocHealthRequest model constructor.
     * @property {module:model/DeploymentAllocHealthRequest}
     */
    DeploymentAllocHealthRequest,

    /**
     * The DeploymentPauseRequest model constructor.
     * @property {module:model/DeploymentPauseRequest}
     */
    DeploymentPauseRequest,

    /**
     * The DeploymentPromoteRequest model constructor.
     * @property {module:model/DeploymentPromoteRequest}
     */
    DeploymentPromoteRequest,

    /**
     * The DeploymentState model constructor.
     * @property {module:model/DeploymentState}
     */
    DeploymentState,

    /**
     * The DeploymentUnblockRequest model constructor.
     * @property {module:model/DeploymentUnblockRequest}
     */
    DeploymentUnblockRequest,

    /**
     * The DeploymentUpdateResponse model constructor.
     * @property {module:model/DeploymentUpdateResponse}
     */
    DeploymentUpdateResponse,

    /**
     * The DesiredTransition model constructor.
     * @property {module:model/DesiredTransition}
     */
    DesiredTransition,

    /**
     * The DesiredUpdates model constructor.
     * @property {module:model/DesiredUpdates}
     */
    DesiredUpdates,

    /**
     * The DispatchPayloadConfig model constructor.
     * @property {module:model/DispatchPayloadConfig}
     */
    DispatchPayloadConfig,

    /**
     * The DrainMetadata model constructor.
     * @property {module:model/DrainMetadata}
     */
    DrainMetadata,

    /**
     * The DrainSpec model constructor.
     * @property {module:model/DrainSpec}
     */
    DrainSpec,

    /**
     * The DrainStrategy model constructor.
     * @property {module:model/DrainStrategy}
     */
    DrainStrategy,

    /**
     * The DriverInfo model constructor.
     * @property {module:model/DriverInfo}
     */
    DriverInfo,

    /**
     * The EphemeralDisk model constructor.
     * @property {module:model/EphemeralDisk}
     */
    EphemeralDisk,

    /**
     * The EvalOptions model constructor.
     * @property {module:model/EvalOptions}
     */
    EvalOptions,

    /**
     * The Evaluation model constructor.
     * @property {module:model/Evaluation}
     */
    Evaluation,

    /**
     * The FieldDiff model constructor.
     * @property {module:model/FieldDiff}
     */
    FieldDiff,

    /**
     * The FuzzyMatch model constructor.
     * @property {module:model/FuzzyMatch}
     */
    FuzzyMatch,

    /**
     * The FuzzySearchRequest model constructor.
     * @property {module:model/FuzzySearchRequest}
     */
    FuzzySearchRequest,

    /**
     * The FuzzySearchResponse model constructor.
     * @property {module:model/FuzzySearchResponse}
     */
    FuzzySearchResponse,

    /**
     * The GaugeValue model constructor.
     * @property {module:model/GaugeValue}
     */
    GaugeValue,

    /**
     * The HostVolumeInfo model constructor.
     * @property {module:model/HostVolumeInfo}
     */
    HostVolumeInfo,

    /**
     * The Job model constructor.
     * @property {module:model/Job}
     */
    Job,

    /**
     * The JobChildrenSummary model constructor.
     * @property {module:model/JobChildrenSummary}
     */
    JobChildrenSummary,

    /**
     * The JobDeregisterResponse model constructor.
     * @property {module:model/JobDeregisterResponse}
     */
    JobDeregisterResponse,

    /**
     * The JobDiff model constructor.
     * @property {module:model/JobDiff}
     */
    JobDiff,

    /**
     * The JobDispatchRequest model constructor.
     * @property {module:model/JobDispatchRequest}
     */
    JobDispatchRequest,

    /**
     * The JobDispatchResponse model constructor.
     * @property {module:model/JobDispatchResponse}
     */
    JobDispatchResponse,

    /**
     * The JobEvaluateRequest model constructor.
     * @property {module:model/JobEvaluateRequest}
     */
    JobEvaluateRequest,

    /**
     * The JobListStub model constructor.
     * @property {module:model/JobListStub}
     */
    JobListStub,

    /**
     * The JobPlanRequest model constructor.
     * @property {module:model/JobPlanRequest}
     */
    JobPlanRequest,

    /**
     * The JobPlanResponse model constructor.
     * @property {module:model/JobPlanResponse}
     */
    JobPlanResponse,

    /**
     * The JobRegisterRequest model constructor.
     * @property {module:model/JobRegisterRequest}
     */
    JobRegisterRequest,

    /**
     * The JobRegisterResponse model constructor.
     * @property {module:model/JobRegisterResponse}
     */
    JobRegisterResponse,

    /**
     * The JobRevertRequest model constructor.
     * @property {module:model/JobRevertRequest}
     */
    JobRevertRequest,

    /**
     * The JobScaleStatusResponse model constructor.
     * @property {module:model/JobScaleStatusResponse}
     */
    JobScaleStatusResponse,

    /**
     * The JobStabilityRequest model constructor.
     * @property {module:model/JobStabilityRequest}
     */
    JobStabilityRequest,

    /**
     * The JobStabilityResponse model constructor.
     * @property {module:model/JobStabilityResponse}
     */
    JobStabilityResponse,

    /**
     * The JobSummary model constructor.
     * @property {module:model/JobSummary}
     */
    JobSummary,

    /**
     * The JobValidateRequest model constructor.
     * @property {module:model/JobValidateRequest}
     */
    JobValidateRequest,

    /**
     * The JobValidateResponse model constructor.
     * @property {module:model/JobValidateResponse}
     */
    JobValidateResponse,

    /**
     * The JobVersionsResponse model constructor.
     * @property {module:model/JobVersionsResponse}
     */
    JobVersionsResponse,

    /**
     * The JobsParseRequest model constructor.
     * @property {module:model/JobsParseRequest}
     */
    JobsParseRequest,

    /**
     * The LogConfig model constructor.
     * @property {module:model/LogConfig}
     */
    LogConfig,

    /**
     * The MetricsSummary model constructor.
     * @property {module:model/MetricsSummary}
     */
    MetricsSummary,

    /**
     * The MigrateStrategy model constructor.
     * @property {module:model/MigrateStrategy}
     */
    MigrateStrategy,

    /**
     * The Multiregion model constructor.
     * @property {module:model/Multiregion}
     */
    Multiregion,

    /**
     * The MultiregionRegion model constructor.
     * @property {module:model/MultiregionRegion}
     */
    MultiregionRegion,

    /**
     * The MultiregionStrategy model constructor.
     * @property {module:model/MultiregionStrategy}
     */
    MultiregionStrategy,

    /**
     * The Namespace model constructor.
     * @property {module:model/Namespace}
     */
    Namespace,

    /**
     * The NetworkResource model constructor.
     * @property {module:model/NetworkResource}
     */
    NetworkResource,

    /**
     * The Node model constructor.
     * @property {module:model/Node}
     */
    Node,

    /**
     * The NodeCpuResources model constructor.
     * @property {module:model/NodeCpuResources}
     */
    NodeCpuResources,

    /**
     * The NodeDevice model constructor.
     * @property {module:model/NodeDevice}
     */
    NodeDevice,

    /**
     * The NodeDeviceLocality model constructor.
     * @property {module:model/NodeDeviceLocality}
     */
    NodeDeviceLocality,

    /**
     * The NodeDeviceResource model constructor.
     * @property {module:model/NodeDeviceResource}
     */
    NodeDeviceResource,

    /**
     * The NodeDiskResources model constructor.
     * @property {module:model/NodeDiskResources}
     */
    NodeDiskResources,

    /**
     * The NodeDrainUpdateResponse model constructor.
     * @property {module:model/NodeDrainUpdateResponse}
     */
    NodeDrainUpdateResponse,

    /**
     * The NodeEligibilityUpdateResponse model constructor.
     * @property {module:model/NodeEligibilityUpdateResponse}
     */
    NodeEligibilityUpdateResponse,

    /**
     * The NodeEvent model constructor.
     * @property {module:model/NodeEvent}
     */
    NodeEvent,

    /**
     * The NodeListStub model constructor.
     * @property {module:model/NodeListStub}
     */
    NodeListStub,

    /**
     * The NodeMemoryResources model constructor.
     * @property {module:model/NodeMemoryResources}
     */
    NodeMemoryResources,

    /**
     * The NodePurgeResponse model constructor.
     * @property {module:model/NodePurgeResponse}
     */
    NodePurgeResponse,

    /**
     * The NodeReservedCpuResources model constructor.
     * @property {module:model/NodeReservedCpuResources}
     */
    NodeReservedCpuResources,

    /**
     * The NodeReservedDiskResources model constructor.
     * @property {module:model/NodeReservedDiskResources}
     */
    NodeReservedDiskResources,

    /**
     * The NodeReservedMemoryResources model constructor.
     * @property {module:model/NodeReservedMemoryResources}
     */
    NodeReservedMemoryResources,

    /**
     * The NodeReservedNetworkResources model constructor.
     * @property {module:model/NodeReservedNetworkResources}
     */
    NodeReservedNetworkResources,

    /**
     * The NodeReservedResources model constructor.
     * @property {module:model/NodeReservedResources}
     */
    NodeReservedResources,

    /**
     * The NodeResources model constructor.
     * @property {module:model/NodeResources}
     */
    NodeResources,

    /**
     * The NodeScoreMeta model constructor.
     * @property {module:model/NodeScoreMeta}
     */
    NodeScoreMeta,

    /**
     * The NodeUpdateDrainRequest model constructor.
     * @property {module:model/NodeUpdateDrainRequest}
     */
    NodeUpdateDrainRequest,

    /**
     * The NodeUpdateEligibilityRequest model constructor.
     * @property {module:model/NodeUpdateEligibilityRequest}
     */
    NodeUpdateEligibilityRequest,

    /**
     * The ObjectDiff model constructor.
     * @property {module:model/ObjectDiff}
     */
    ObjectDiff,

    /**
     * The OneTimeToken model constructor.
     * @property {module:model/OneTimeToken}
     */
    OneTimeToken,

    /**
     * The OneTimeTokenExchangeRequest model constructor.
     * @property {module:model/OneTimeTokenExchangeRequest}
     */
    OneTimeTokenExchangeRequest,

    /**
     * The ParameterizedJobConfig model constructor.
     * @property {module:model/ParameterizedJobConfig}
     */
    ParameterizedJobConfig,

    /**
     * The PeriodicConfig model constructor.
     * @property {module:model/PeriodicConfig}
     */
    PeriodicConfig,

    /**
     * The PeriodicForceResponse model constructor.
     * @property {module:model/PeriodicForceResponse}
     */
    PeriodicForceResponse,

    /**
     * The PlanAnnotations model constructor.
     * @property {module:model/PlanAnnotations}
     */
    PlanAnnotations,

    /**
     * The PointValue model constructor.
     * @property {module:model/PointValue}
     */
    PointValue,

    /**
     * The Port model constructor.
     * @property {module:model/Port}
     */
    Port,

    /**
     * The PortMapping model constructor.
     * @property {module:model/PortMapping}
     */
    PortMapping,

    /**
     * The QuotaLimit model constructor.
     * @property {module:model/QuotaLimit}
     */
    QuotaLimit,

    /**
     * The QuotaSpec model constructor.
     * @property {module:model/QuotaSpec}
     */
    QuotaSpec,

    /**
     * The RequestedDevice model constructor.
     * @property {module:model/RequestedDevice}
     */
    RequestedDevice,

    /**
     * The RescheduleEvent model constructor.
     * @property {module:model/RescheduleEvent}
     */
    RescheduleEvent,

    /**
     * The ReschedulePolicy model constructor.
     * @property {module:model/ReschedulePolicy}
     */
    ReschedulePolicy,

    /**
     * The RescheduleTracker model constructor.
     * @property {module:model/RescheduleTracker}
     */
    RescheduleTracker,

    /**
     * The Resources model constructor.
     * @property {module:model/Resources}
     */
    Resources,

    /**
     * The RestartPolicy model constructor.
     * @property {module:model/RestartPolicy}
     */
    RestartPolicy,

    /**
     * The SampledValue model constructor.
     * @property {module:model/SampledValue}
     */
    SampledValue,

    /**
     * The ScalingEvent model constructor.
     * @property {module:model/ScalingEvent}
     */
    ScalingEvent,

    /**
     * The ScalingPolicy model constructor.
     * @property {module:model/ScalingPolicy}
     */
    ScalingPolicy,

    /**
     * The ScalingPolicyListStub model constructor.
     * @property {module:model/ScalingPolicyListStub}
     */
    ScalingPolicyListStub,

    /**
     * The ScalingRequest model constructor.
     * @property {module:model/ScalingRequest}
     */
    ScalingRequest,

    /**
     * The SearchRequest model constructor.
     * @property {module:model/SearchRequest}
     */
    SearchRequest,

    /**
     * The SearchResponse model constructor.
     * @property {module:model/SearchResponse}
     */
    SearchResponse,

    /**
     * The Service model constructor.
     * @property {module:model/Service}
     */
    Service,

    /**
     * The ServiceCheck model constructor.
     * @property {module:model/ServiceCheck}
     */
    ServiceCheck,

    /**
     * The SidecarTask model constructor.
     * @property {module:model/SidecarTask}
     */
    SidecarTask,

    /**
     * The Spread model constructor.
     * @property {module:model/Spread}
     */
    Spread,

    /**
     * The SpreadTarget model constructor.
     * @property {module:model/SpreadTarget}
     */
    SpreadTarget,

    /**
     * The Task model constructor.
     * @property {module:model/Task}
     */
    Task,

    /**
     * The TaskArtifact model constructor.
     * @property {module:model/TaskArtifact}
     */
    TaskArtifact,

    /**
     * The TaskCSIPluginConfig model constructor.
     * @property {module:model/TaskCSIPluginConfig}
     */
    TaskCSIPluginConfig,

    /**
     * The TaskDiff model constructor.
     * @property {module:model/TaskDiff}
     */
    TaskDiff,

    /**
     * The TaskEvent model constructor.
     * @property {module:model/TaskEvent}
     */
    TaskEvent,

    /**
     * The TaskGroup model constructor.
     * @property {module:model/TaskGroup}
     */
    TaskGroup,

    /**
     * The TaskGroupDiff model constructor.
     * @property {module:model/TaskGroupDiff}
     */
    TaskGroupDiff,

    /**
     * The TaskGroupScaleStatus model constructor.
     * @property {module:model/TaskGroupScaleStatus}
     */
    TaskGroupScaleStatus,

    /**
     * The TaskGroupSummary model constructor.
     * @property {module:model/TaskGroupSummary}
     */
    TaskGroupSummary,

    /**
     * The TaskHandle model constructor.
     * @property {module:model/TaskHandle}
     */
    TaskHandle,

    /**
     * The TaskLifecycle model constructor.
     * @property {module:model/TaskLifecycle}
     */
    TaskLifecycle,

    /**
     * The TaskState model constructor.
     * @property {module:model/TaskState}
     */
    TaskState,

    /**
     * The Template model constructor.
     * @property {module:model/Template}
     */
    Template,

    /**
     * The UpdateStrategy model constructor.
     * @property {module:model/UpdateStrategy}
     */
    UpdateStrategy,

    /**
     * The Vault model constructor.
     * @property {module:model/Vault}
     */
    Vault,

    /**
     * The VolumeMount model constructor.
     * @property {module:model/VolumeMount}
     */
    VolumeMount,

    /**
     * The VolumeRequest model constructor.
     * @property {module:model/VolumeRequest}
     */
    VolumeRequest,

    /**
    * The ACLApi service constructor.
    * @property {module:api/ACLApi}
    */
    ACLApi,

    /**
    * The AllocationsApi service constructor.
    * @property {module:api/AllocationsApi}
    */
    AllocationsApi,

    /**
    * The DeploymentsApi service constructor.
    * @property {module:api/DeploymentsApi}
    */
    DeploymentsApi,

    /**
    * The EnterpriseApi service constructor.
    * @property {module:api/EnterpriseApi}
    */
    EnterpriseApi,

    /**
    * The EvaluationsApi service constructor.
    * @property {module:api/EvaluationsApi}
    */
    EvaluationsApi,

    /**
    * The JobsApi service constructor.
    * @property {module:api/JobsApi}
    */
    JobsApi,

    /**
    * The MetricsApi service constructor.
    * @property {module:api/MetricsApi}
    */
    MetricsApi,

    /**
    * The NamespacesApi service constructor.
    * @property {module:api/NamespacesApi}
    */
    NamespacesApi,

    /**
    * The NodesApi service constructor.
    * @property {module:api/NodesApi}
    */
    NodesApi,

    /**
    * The PluginsApi service constructor.
    * @property {module:api/PluginsApi}
    */
    PluginsApi,

    /**
    * The RegionsApi service constructor.
    * @property {module:api/RegionsApi}
    */
    RegionsApi,

    /**
    * The ScalingApi service constructor.
    * @property {module:api/ScalingApi}
    */
    ScalingApi,

    /**
    * The SearchApi service constructor.
    * @property {module:api/SearchApi}
    */
    SearchApi,

    /**
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
    * The SystemApi service constructor.
    * @property {module:api/SystemApi}
    */
    SystemApi,
=======
=======
>>>>>>> 8d74053 (updated generator/status.go)
=======
>>>>>>> e3fb144 (added v1/status.go and v1/status_test.go, tests passing)
    * The StatusApi service constructor.
    * @property {module:api/StatusApi}
    */
    StatusApi,
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 8d74053 (updated generator/status.go)
=======
>>>>>>> 8d74053 (updated generator/status.go)
=======

    /**
    * The SystemApi service constructor.
    * @property {module:api/SystemApi}
    */
    SystemApi,
>>>>>>> e3fb144 (added v1/status.go and v1/status_test.go, tests passing)

    /**
    * The VolumesApi service constructor.
    * @property {module:api/VolumesApi}
    */
    VolumesApi
};
