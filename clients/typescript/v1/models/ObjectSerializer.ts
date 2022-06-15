export * from './ACLPolicy';
export * from './ACLPolicyListStub';
export * from './ACLToken';
export * from './ACLTokenListStub';
export * from './Affinity';
export * from './AllocDeploymentStatus';
export * from './AllocStopResponse';
export * from './AllocatedCpuResources';
export * from './AllocatedDeviceResource';
export * from './AllocatedMemoryResources';
export * from './AllocatedResources';
export * from './AllocatedSharedResources';
export * from './AllocatedTaskResources';
export * from './Allocation';
export * from './AllocationListStub';
export * from './AllocationMetric';
export * from './Attribute';
export * from './AutopilotConfiguration';
export * from './CSIControllerInfo';
export * from './CSIInfo';
export * from './CSIMountOptions';
export * from './CSINodeInfo';
export * from './CSIPlugin';
export * from './CSIPluginListStub';
export * from './CSISnapshot';
export * from './CSISnapshotCreateRequest';
export * from './CSISnapshotCreateResponse';
export * from './CSISnapshotListResponse';
export * from './CSITopology';
export * from './CSITopologyRequest';
export * from './CSIVolume';
export * from './CSIVolumeCapability';
export * from './CSIVolumeCreateRequest';
export * from './CSIVolumeExternalStub';
export * from './CSIVolumeListExternalResponse';
export * from './CSIVolumeListStub';
export * from './CSIVolumeRegisterRequest';
export * from './CheckRestart';
export * from './Constraint';
export * from './Consul';
export * from './ConsulConnect';
export * from './ConsulExposeConfig';
export * from './ConsulExposePath';
export * from './ConsulGateway';
export * from './ConsulGatewayBindAddress';
export * from './ConsulGatewayProxy';
export * from './ConsulGatewayTLSConfig';
export * from './ConsulIngressConfigEntry';
export * from './ConsulIngressListener';
export * from './ConsulIngressService';
export * from './ConsulLinkedService';
export * from './ConsulMeshGateway';
export * from './ConsulProxy';
export * from './ConsulSidecarService';
export * from './ConsulTerminatingConfigEntry';
export * from './ConsulUpstream';
export * from './DNSConfig';
export * from './Deployment';
export * from './DeploymentAllocHealthRequest';
export * from './DeploymentPauseRequest';
export * from './DeploymentPromoteRequest';
export * from './DeploymentState';
export * from './DeploymentUnblockRequest';
export * from './DeploymentUpdateResponse';
export * from './DesiredTransition';
export * from './DesiredUpdates';
export * from './DispatchPayloadConfig';
export * from './DrainMetadata';
export * from './DrainSpec';
export * from './DrainStrategy';
export * from './DriverInfo';
export * from './EphemeralDisk';
export * from './EvalOptions';
export * from './Evaluation';
export * from './EvaluationStub';
export * from './FieldDiff';
export * from './FuzzyMatch';
export * from './FuzzySearchRequest';
export * from './FuzzySearchResponse';
export * from './GaugeValue';
export * from './HostNetworkInfo';
export * from './HostVolumeInfo';
export * from './Job';
export * from './JobChildrenSummary';
export * from './JobDeregisterResponse';
export * from './JobDiff';
export * from './JobDispatchRequest';
export * from './JobDispatchResponse';
export * from './JobEvaluateRequest';
export * from './JobListStub';
export * from './JobPlanRequest';
export * from './JobPlanResponse';
export * from './JobRegisterRequest';
export * from './JobRegisterResponse';
export * from './JobRevertRequest';
export * from './JobScaleStatusResponse';
export * from './JobStabilityRequest';
export * from './JobStabilityResponse';
export * from './JobSummary';
export * from './JobValidateRequest';
export * from './JobValidateResponse';
export * from './JobVersionsResponse';
export * from './JobsParseRequest';
export * from './LogConfig';
export * from './MetricsSummary';
export * from './MigrateStrategy';
export * from './Multiregion';
export * from './MultiregionRegion';
export * from './MultiregionStrategy';
export * from './Namespace';
export * from './NamespaceCapabilities';
export * from './NetworkResource';
export * from './Node';
export * from './NodeCpuResources';
export * from './NodeDevice';
export * from './NodeDeviceLocality';
export * from './NodeDeviceResource';
export * from './NodeDiskResources';
export * from './NodeDrainUpdateResponse';
export * from './NodeEligibilityUpdateResponse';
export * from './NodeEvent';
export * from './NodeListStub';
export * from './NodeMemoryResources';
export * from './NodePurgeResponse';
export * from './NodeReservedCpuResources';
export * from './NodeReservedDiskResources';
export * from './NodeReservedMemoryResources';
export * from './NodeReservedNetworkResources';
export * from './NodeReservedResources';
export * from './NodeResources';
export * from './NodeScoreMeta';
export * from './NodeUpdateDrainRequest';
export * from './NodeUpdateEligibilityRequest';
export * from './ObjectDiff';
export * from './OneTimeToken';
export * from './OneTimeTokenExchangeRequest';
export * from './OperatorHealthReply';
export * from './ParameterizedJobConfig';
export * from './PeriodicConfig';
export * from './PeriodicForceResponse';
export * from './PlanAnnotations';
export * from './PointValue';
export * from './Port';
export * from './PortMapping';
export * from './PreemptionConfig';
export * from './QuotaLimit';
export * from './QuotaSpec';
export * from './RaftConfiguration';
export * from './RaftServer';
export * from './RequestedDevice';
export * from './RescheduleEvent';
export * from './ReschedulePolicy';
export * from './RescheduleTracker';
export * from './Resources';
export * from './RestartPolicy';
export * from './SampledValue';
export * from './ScalingEvent';
export * from './ScalingPolicy';
export * from './ScalingPolicyListStub';
export * from './ScalingRequest';
export * from './SchedulerConfiguration';
export * from './SchedulerConfigurationResponse';
export * from './SchedulerSetConfigurationResponse';
export * from './SearchRequest';
export * from './SearchResponse';
export * from './ServerHealth';
export * from './Service';
export * from './ServiceCheck';
export * from './ServiceRegistration';
export * from './SidecarTask';
export * from './Spread';
export * from './SpreadTarget';
export * from './Task';
export * from './TaskArtifact';
export * from './TaskCSIPluginConfig';
export * from './TaskDiff';
export * from './TaskEvent';
export * from './TaskGroup';
export * from './TaskGroupDiff';
export * from './TaskGroupScaleStatus';
export * from './TaskGroupSummary';
export * from './TaskHandle';
export * from './TaskLifecycle';
export * from './TaskState';
export * from './Template';
export * from './UpdateStrategy';
export * from './Vault';
export * from './VolumeMount';
export * from './VolumeRequest';
export * from './WaitConfig';

import { ACLPolicy } from './ACLPolicy';
import { ACLPolicyListStub } from './ACLPolicyListStub';
import { ACLToken } from './ACLToken';
import { ACLTokenListStub } from './ACLTokenListStub';
import { Affinity } from './Affinity';
import { AllocDeploymentStatus } from './AllocDeploymentStatus';
import { AllocStopResponse } from './AllocStopResponse';
import { AllocatedCpuResources } from './AllocatedCpuResources';
import { AllocatedDeviceResource } from './AllocatedDeviceResource';
import { AllocatedMemoryResources } from './AllocatedMemoryResources';
import { AllocatedResources } from './AllocatedResources';
import { AllocatedSharedResources } from './AllocatedSharedResources';
import { AllocatedTaskResources } from './AllocatedTaskResources';
import { Allocation } from './Allocation';
import { AllocationListStub } from './AllocationListStub';
import { AllocationMetric } from './AllocationMetric';
import { Attribute } from './Attribute';
import { AutopilotConfiguration } from './AutopilotConfiguration';
import { CSIControllerInfo } from './CSIControllerInfo';
import { CSIInfo } from './CSIInfo';
import { CSIMountOptions } from './CSIMountOptions';
import { CSINodeInfo } from './CSINodeInfo';
import { CSIPlugin } from './CSIPlugin';
import { CSIPluginListStub } from './CSIPluginListStub';
import { CSISnapshot } from './CSISnapshot';
import { CSISnapshotCreateRequest } from './CSISnapshotCreateRequest';
import { CSISnapshotCreateResponse } from './CSISnapshotCreateResponse';
import { CSISnapshotListResponse } from './CSISnapshotListResponse';
import { CSITopology } from './CSITopology';
import { CSITopologyRequest } from './CSITopologyRequest';
import { CSIVolume } from './CSIVolume';
import { CSIVolumeCapability } from './CSIVolumeCapability';
import { CSIVolumeCreateRequest } from './CSIVolumeCreateRequest';
import { CSIVolumeExternalStub } from './CSIVolumeExternalStub';
import { CSIVolumeListExternalResponse } from './CSIVolumeListExternalResponse';
import { CSIVolumeListStub } from './CSIVolumeListStub';
import { CSIVolumeRegisterRequest } from './CSIVolumeRegisterRequest';
import { CheckRestart } from './CheckRestart';
import { Constraint } from './Constraint';
import { Consul } from './Consul';
import { ConsulConnect } from './ConsulConnect';
import { ConsulExposeConfig } from './ConsulExposeConfig';
import { ConsulExposePath } from './ConsulExposePath';
import { ConsulGateway } from './ConsulGateway';
import { ConsulGatewayBindAddress } from './ConsulGatewayBindAddress';
import { ConsulGatewayProxy } from './ConsulGatewayProxy';
import { ConsulGatewayTLSConfig } from './ConsulGatewayTLSConfig';
import { ConsulIngressConfigEntry } from './ConsulIngressConfigEntry';
import { ConsulIngressListener } from './ConsulIngressListener';
import { ConsulIngressService } from './ConsulIngressService';
import { ConsulLinkedService } from './ConsulLinkedService';
import { ConsulMeshGateway } from './ConsulMeshGateway';
import { ConsulProxy } from './ConsulProxy';
import { ConsulSidecarService } from './ConsulSidecarService';
import { ConsulTerminatingConfigEntry } from './ConsulTerminatingConfigEntry';
import { ConsulUpstream } from './ConsulUpstream';
import { DNSConfig } from './DNSConfig';
import { Deployment } from './Deployment';
import { DeploymentAllocHealthRequest } from './DeploymentAllocHealthRequest';
import { DeploymentPauseRequest } from './DeploymentPauseRequest';
import { DeploymentPromoteRequest } from './DeploymentPromoteRequest';
import { DeploymentState } from './DeploymentState';
import { DeploymentUnblockRequest } from './DeploymentUnblockRequest';
import { DeploymentUpdateResponse } from './DeploymentUpdateResponse';
import { DesiredTransition } from './DesiredTransition';
import { DesiredUpdates } from './DesiredUpdates';
import { DispatchPayloadConfig } from './DispatchPayloadConfig';
import { DrainMetadata } from './DrainMetadata';
import { DrainSpec } from './DrainSpec';
import { DrainStrategy } from './DrainStrategy';
import { DriverInfo } from './DriverInfo';
import { EphemeralDisk } from './EphemeralDisk';
import { EvalOptions } from './EvalOptions';
import { Evaluation } from './Evaluation';
import { EvaluationStub } from './EvaluationStub';
import { FieldDiff } from './FieldDiff';
import { FuzzyMatch } from './FuzzyMatch';
import { FuzzySearchRequest } from './FuzzySearchRequest';
import { FuzzySearchResponse } from './FuzzySearchResponse';
import { GaugeValue } from './GaugeValue';
import { HostNetworkInfo } from './HostNetworkInfo';
import { HostVolumeInfo } from './HostVolumeInfo';
import { Job } from './Job';
import { JobChildrenSummary } from './JobChildrenSummary';
import { JobDeregisterResponse } from './JobDeregisterResponse';
import { JobDiff } from './JobDiff';
import { JobDispatchRequest } from './JobDispatchRequest';
import { JobDispatchResponse } from './JobDispatchResponse';
import { JobEvaluateRequest } from './JobEvaluateRequest';
import { JobListStub } from './JobListStub';
import { JobPlanRequest } from './JobPlanRequest';
import { JobPlanResponse } from './JobPlanResponse';
import { JobRegisterRequest } from './JobRegisterRequest';
import { JobRegisterResponse } from './JobRegisterResponse';
import { JobRevertRequest } from './JobRevertRequest';
import { JobScaleStatusResponse } from './JobScaleStatusResponse';
import { JobStabilityRequest } from './JobStabilityRequest';
import { JobStabilityResponse } from './JobStabilityResponse';
import { JobSummary } from './JobSummary';
import { JobValidateRequest } from './JobValidateRequest';
import { JobValidateResponse } from './JobValidateResponse';
import { JobVersionsResponse } from './JobVersionsResponse';
import { JobsParseRequest } from './JobsParseRequest';
import { LogConfig } from './LogConfig';
import { MetricsSummary } from './MetricsSummary';
import { MigrateStrategy } from './MigrateStrategy';
import { Multiregion } from './Multiregion';
import { MultiregionRegion } from './MultiregionRegion';
import { MultiregionStrategy } from './MultiregionStrategy';
import { Namespace } from './Namespace';
import { NamespaceCapabilities } from './NamespaceCapabilities';
import { NetworkResource } from './NetworkResource';
import { Node } from './Node';
import { NodeCpuResources } from './NodeCpuResources';
import { NodeDevice } from './NodeDevice';
import { NodeDeviceLocality } from './NodeDeviceLocality';
import { NodeDeviceResource } from './NodeDeviceResource';
import { NodeDiskResources } from './NodeDiskResources';
import { NodeDrainUpdateResponse } from './NodeDrainUpdateResponse';
import { NodeEligibilityUpdateResponse } from './NodeEligibilityUpdateResponse';
import { NodeEvent } from './NodeEvent';
import { NodeListStub } from './NodeListStub';
import { NodeMemoryResources } from './NodeMemoryResources';
import { NodePurgeResponse } from './NodePurgeResponse';
import { NodeReservedCpuResources } from './NodeReservedCpuResources';
import { NodeReservedDiskResources } from './NodeReservedDiskResources';
import { NodeReservedMemoryResources } from './NodeReservedMemoryResources';
import { NodeReservedNetworkResources } from './NodeReservedNetworkResources';
import { NodeReservedResources } from './NodeReservedResources';
import { NodeResources } from './NodeResources';
import { NodeScoreMeta } from './NodeScoreMeta';
import { NodeUpdateDrainRequest } from './NodeUpdateDrainRequest';
import { NodeUpdateEligibilityRequest } from './NodeUpdateEligibilityRequest';
import { ObjectDiff } from './ObjectDiff';
import { OneTimeToken } from './OneTimeToken';
import { OneTimeTokenExchangeRequest } from './OneTimeTokenExchangeRequest';
import { OperatorHealthReply } from './OperatorHealthReply';
import { ParameterizedJobConfig } from './ParameterizedJobConfig';
import { PeriodicConfig } from './PeriodicConfig';
import { PeriodicForceResponse } from './PeriodicForceResponse';
import { PlanAnnotations } from './PlanAnnotations';
import { PointValue } from './PointValue';
import { Port } from './Port';
import { PortMapping } from './PortMapping';
import { PreemptionConfig } from './PreemptionConfig';
import { QuotaLimit } from './QuotaLimit';
import { QuotaSpec } from './QuotaSpec';
import { RaftConfiguration } from './RaftConfiguration';
import { RaftServer } from './RaftServer';
import { RequestedDevice } from './RequestedDevice';
import { RescheduleEvent } from './RescheduleEvent';
import { ReschedulePolicy } from './ReschedulePolicy';
import { RescheduleTracker } from './RescheduleTracker';
import { Resources } from './Resources';
import { RestartPolicy } from './RestartPolicy';
import { SampledValue } from './SampledValue';
import { ScalingEvent } from './ScalingEvent';
import { ScalingPolicy } from './ScalingPolicy';
import { ScalingPolicyListStub } from './ScalingPolicyListStub';
import { ScalingRequest } from './ScalingRequest';
import { SchedulerConfiguration } from './SchedulerConfiguration';
import { SchedulerConfigurationResponse } from './SchedulerConfigurationResponse';
import { SchedulerSetConfigurationResponse } from './SchedulerSetConfigurationResponse';
import { SearchRequest } from './SearchRequest';
import { SearchResponse } from './SearchResponse';
import { ServerHealth } from './ServerHealth';
import { Service } from './Service';
import { ServiceCheck } from './ServiceCheck';
import { ServiceRegistration } from './ServiceRegistration';
import { SidecarTask } from './SidecarTask';
import { Spread } from './Spread';
import { SpreadTarget } from './SpreadTarget';
import { Task } from './Task';
import { TaskArtifact } from './TaskArtifact';
import { TaskCSIPluginConfig } from './TaskCSIPluginConfig';
import { TaskDiff } from './TaskDiff';
import { TaskEvent } from './TaskEvent';
import { TaskGroup } from './TaskGroup';
import { TaskGroupDiff } from './TaskGroupDiff';
import { TaskGroupScaleStatus } from './TaskGroupScaleStatus';
import { TaskGroupSummary } from './TaskGroupSummary';
import { TaskHandle } from './TaskHandle';
import { TaskLifecycle } from './TaskLifecycle';
import { TaskState } from './TaskState';
import { Template } from './Template';
import { UpdateStrategy } from './UpdateStrategy';
import { Vault } from './Vault';
import { VolumeMount } from './VolumeMount';
import { VolumeRequest } from './VolumeRequest';
import { WaitConfig } from './WaitConfig';

/* tslint:disable:no-unused-variable */
let primitives = [
                    "string",
                    "boolean",
                    "double",
                    "integer",
                    "long",
                    "float",
                    "number",
                    "any"
                 ];

const supportedMediaTypes: { [mediaType: string]: number } = {
  "application/json": Infinity,
  "application/octet-stream": 0,
  "application/x-www-form-urlencoded": 0
}


let enumsMap: Set<string> = new Set<string>([
]);

let typeMap: {[index: string]: any} = {
    "ACLPolicy": ACLPolicy,
    "ACLPolicyListStub": ACLPolicyListStub,
    "ACLToken": ACLToken,
    "ACLTokenListStub": ACLTokenListStub,
    "Affinity": Affinity,
    "AllocDeploymentStatus": AllocDeploymentStatus,
    "AllocStopResponse": AllocStopResponse,
    "AllocatedCpuResources": AllocatedCpuResources,
    "AllocatedDeviceResource": AllocatedDeviceResource,
    "AllocatedMemoryResources": AllocatedMemoryResources,
    "AllocatedResources": AllocatedResources,
    "AllocatedSharedResources": AllocatedSharedResources,
    "AllocatedTaskResources": AllocatedTaskResources,
    "Allocation": Allocation,
    "AllocationListStub": AllocationListStub,
    "AllocationMetric": AllocationMetric,
    "Attribute": Attribute,
    "AutopilotConfiguration": AutopilotConfiguration,
    "CSIControllerInfo": CSIControllerInfo,
    "CSIInfo": CSIInfo,
    "CSIMountOptions": CSIMountOptions,
    "CSINodeInfo": CSINodeInfo,
    "CSIPlugin": CSIPlugin,
    "CSIPluginListStub": CSIPluginListStub,
    "CSISnapshot": CSISnapshot,
    "CSISnapshotCreateRequest": CSISnapshotCreateRequest,
    "CSISnapshotCreateResponse": CSISnapshotCreateResponse,
    "CSISnapshotListResponse": CSISnapshotListResponse,
    "CSITopology": CSITopology,
    "CSITopologyRequest": CSITopologyRequest,
    "CSIVolume": CSIVolume,
    "CSIVolumeCapability": CSIVolumeCapability,
    "CSIVolumeCreateRequest": CSIVolumeCreateRequest,
    "CSIVolumeExternalStub": CSIVolumeExternalStub,
    "CSIVolumeListExternalResponse": CSIVolumeListExternalResponse,
    "CSIVolumeListStub": CSIVolumeListStub,
    "CSIVolumeRegisterRequest": CSIVolumeRegisterRequest,
    "CheckRestart": CheckRestart,
    "Constraint": Constraint,
    "Consul": Consul,
    "ConsulConnect": ConsulConnect,
    "ConsulExposeConfig": ConsulExposeConfig,
    "ConsulExposePath": ConsulExposePath,
    "ConsulGateway": ConsulGateway,
    "ConsulGatewayBindAddress": ConsulGatewayBindAddress,
    "ConsulGatewayProxy": ConsulGatewayProxy,
    "ConsulGatewayTLSConfig": ConsulGatewayTLSConfig,
    "ConsulIngressConfigEntry": ConsulIngressConfigEntry,
    "ConsulIngressListener": ConsulIngressListener,
    "ConsulIngressService": ConsulIngressService,
    "ConsulLinkedService": ConsulLinkedService,
    "ConsulMeshGateway": ConsulMeshGateway,
    "ConsulProxy": ConsulProxy,
    "ConsulSidecarService": ConsulSidecarService,
    "ConsulTerminatingConfigEntry": ConsulTerminatingConfigEntry,
    "ConsulUpstream": ConsulUpstream,
    "DNSConfig": DNSConfig,
    "Deployment": Deployment,
    "DeploymentAllocHealthRequest": DeploymentAllocHealthRequest,
    "DeploymentPauseRequest": DeploymentPauseRequest,
    "DeploymentPromoteRequest": DeploymentPromoteRequest,
    "DeploymentState": DeploymentState,
    "DeploymentUnblockRequest": DeploymentUnblockRequest,
    "DeploymentUpdateResponse": DeploymentUpdateResponse,
    "DesiredTransition": DesiredTransition,
    "DesiredUpdates": DesiredUpdates,
    "DispatchPayloadConfig": DispatchPayloadConfig,
    "DrainMetadata": DrainMetadata,
    "DrainSpec": DrainSpec,
    "DrainStrategy": DrainStrategy,
    "DriverInfo": DriverInfo,
    "EphemeralDisk": EphemeralDisk,
    "EvalOptions": EvalOptions,
    "Evaluation": Evaluation,
    "EvaluationStub": EvaluationStub,
    "FieldDiff": FieldDiff,
    "FuzzyMatch": FuzzyMatch,
    "FuzzySearchRequest": FuzzySearchRequest,
    "FuzzySearchResponse": FuzzySearchResponse,
    "GaugeValue": GaugeValue,
    "HostNetworkInfo": HostNetworkInfo,
    "HostVolumeInfo": HostVolumeInfo,
    "Job": Job,
    "JobChildrenSummary": JobChildrenSummary,
    "JobDeregisterResponse": JobDeregisterResponse,
    "JobDiff": JobDiff,
    "JobDispatchRequest": JobDispatchRequest,
    "JobDispatchResponse": JobDispatchResponse,
    "JobEvaluateRequest": JobEvaluateRequest,
    "JobListStub": JobListStub,
    "JobPlanRequest": JobPlanRequest,
    "JobPlanResponse": JobPlanResponse,
    "JobRegisterRequest": JobRegisterRequest,
    "JobRegisterResponse": JobRegisterResponse,
    "JobRevertRequest": JobRevertRequest,
    "JobScaleStatusResponse": JobScaleStatusResponse,
    "JobStabilityRequest": JobStabilityRequest,
    "JobStabilityResponse": JobStabilityResponse,
    "JobSummary": JobSummary,
    "JobValidateRequest": JobValidateRequest,
    "JobValidateResponse": JobValidateResponse,
    "JobVersionsResponse": JobVersionsResponse,
    "JobsParseRequest": JobsParseRequest,
    "LogConfig": LogConfig,
    "MetricsSummary": MetricsSummary,
    "MigrateStrategy": MigrateStrategy,
    "Multiregion": Multiregion,
    "MultiregionRegion": MultiregionRegion,
    "MultiregionStrategy": MultiregionStrategy,
    "Namespace": Namespace,
    "NamespaceCapabilities": NamespaceCapabilities,
    "NetworkResource": NetworkResource,
    "Node": Node,
    "NodeCpuResources": NodeCpuResources,
    "NodeDevice": NodeDevice,
    "NodeDeviceLocality": NodeDeviceLocality,
    "NodeDeviceResource": NodeDeviceResource,
    "NodeDiskResources": NodeDiskResources,
    "NodeDrainUpdateResponse": NodeDrainUpdateResponse,
    "NodeEligibilityUpdateResponse": NodeEligibilityUpdateResponse,
    "NodeEvent": NodeEvent,
    "NodeListStub": NodeListStub,
    "NodeMemoryResources": NodeMemoryResources,
    "NodePurgeResponse": NodePurgeResponse,
    "NodeReservedCpuResources": NodeReservedCpuResources,
    "NodeReservedDiskResources": NodeReservedDiskResources,
    "NodeReservedMemoryResources": NodeReservedMemoryResources,
    "NodeReservedNetworkResources": NodeReservedNetworkResources,
    "NodeReservedResources": NodeReservedResources,
    "NodeResources": NodeResources,
    "NodeScoreMeta": NodeScoreMeta,
    "NodeUpdateDrainRequest": NodeUpdateDrainRequest,
    "NodeUpdateEligibilityRequest": NodeUpdateEligibilityRequest,
    "ObjectDiff": ObjectDiff,
    "OneTimeToken": OneTimeToken,
    "OneTimeTokenExchangeRequest": OneTimeTokenExchangeRequest,
    "OperatorHealthReply": OperatorHealthReply,
    "ParameterizedJobConfig": ParameterizedJobConfig,
    "PeriodicConfig": PeriodicConfig,
    "PeriodicForceResponse": PeriodicForceResponse,
    "PlanAnnotations": PlanAnnotations,
    "PointValue": PointValue,
    "Port": Port,
    "PortMapping": PortMapping,
    "PreemptionConfig": PreemptionConfig,
    "QuotaLimit": QuotaLimit,
    "QuotaSpec": QuotaSpec,
    "RaftConfiguration": RaftConfiguration,
    "RaftServer": RaftServer,
    "RequestedDevice": RequestedDevice,
    "RescheduleEvent": RescheduleEvent,
    "ReschedulePolicy": ReschedulePolicy,
    "RescheduleTracker": RescheduleTracker,
    "Resources": Resources,
    "RestartPolicy": RestartPolicy,
    "SampledValue": SampledValue,
    "ScalingEvent": ScalingEvent,
    "ScalingPolicy": ScalingPolicy,
    "ScalingPolicyListStub": ScalingPolicyListStub,
    "ScalingRequest": ScalingRequest,
    "SchedulerConfiguration": SchedulerConfiguration,
    "SchedulerConfigurationResponse": SchedulerConfigurationResponse,
    "SchedulerSetConfigurationResponse": SchedulerSetConfigurationResponse,
    "SearchRequest": SearchRequest,
    "SearchResponse": SearchResponse,
    "ServerHealth": ServerHealth,
    "Service": Service,
    "ServiceCheck": ServiceCheck,
    "ServiceRegistration": ServiceRegistration,
    "SidecarTask": SidecarTask,
    "Spread": Spread,
    "SpreadTarget": SpreadTarget,
    "Task": Task,
    "TaskArtifact": TaskArtifact,
    "TaskCSIPluginConfig": TaskCSIPluginConfig,
    "TaskDiff": TaskDiff,
    "TaskEvent": TaskEvent,
    "TaskGroup": TaskGroup,
    "TaskGroupDiff": TaskGroupDiff,
    "TaskGroupScaleStatus": TaskGroupScaleStatus,
    "TaskGroupSummary": TaskGroupSummary,
    "TaskHandle": TaskHandle,
    "TaskLifecycle": TaskLifecycle,
    "TaskState": TaskState,
    "Template": Template,
    "UpdateStrategy": UpdateStrategy,
    "Vault": Vault,
    "VolumeMount": VolumeMount,
    "VolumeRequest": VolumeRequest,
    "WaitConfig": WaitConfig,
}

export class ObjectSerializer {
    public static findCorrectType(data: any, expectedType: string) {
        if (data == undefined) {
            return expectedType;
        } else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        } else if (expectedType === "Date") {
            return expectedType;
        } else {
            if (enumsMap.has(expectedType)) {
                return expectedType;
            }

            if (!typeMap[expectedType]) {
                return expectedType; // w/e we don't know the type
            }

            // Check the discriminator
            let discriminatorProperty = typeMap[expectedType].discriminator;
            if (discriminatorProperty == null) {
                return expectedType; // the type does not have a discriminator. use it.
            } else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if(typeMap[discriminatorType]){
                        return discriminatorType; // use the type given in the discriminator
                    } else {
                        return expectedType; // discriminator did not map to a type
                    }
                } else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }

    public static serialize(data: any, type: string, format: string) {
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index in data) {
                let date = data[index];
                transformedData.push(ObjectSerializer.serialize(date, subType, format));
            }
            return transformedData;
        } else if (type === "Date") {
            if (format == "date") {
                let month = data.getMonth()+1
                month = month < 10 ? "0" + month.toString() : month.toString()
                let day = data.getDate();
                day = day < 10 ? "0" + day.toString() : day.toString();

                return data.getFullYear() + "-" + month + "-" + day;
            } else {
                return data.toISOString();
            }
        } else {
            if (enumsMap.has(type)) {
                return data;
            }
            if (!typeMap[type]) { // in case we dont know the type
                return data;
            }

            // Get the actual type of this object
            type = this.findCorrectType(data, type);

            // get the map for the correct type.
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            let instance: {[index: string]: any} = {};
            for (let index in attributeTypes) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type, attributeType.format);
            }
            return instance;
        }
    }

    public static deserialize(data: any, type: string, format: string) {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index in data) {
                let date = data[index];
                transformedData.push(ObjectSerializer.deserialize(date, subType, format));
            }
            return transformedData;
        } else if (type === "Date") {
            return new Date(data);
        } else {
            if (enumsMap.has(type)) {// is Enum
                return data;
            }

            if (!typeMap[type]) { // dont know the type
                return data;
            }
            let instance = new typeMap[type]();
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            for (let index in attributeTypes) {
                let attributeType = attributeTypes[index];
                let value = ObjectSerializer.deserialize(data[attributeType.baseName], attributeType.type, attributeType.format);
                if (value !== undefined) {
                    instance[attributeType.name] = value;
                }
            }
            return instance;
        }
    }


    /**
     * Normalize media type
     *
     * We currently do not handle any media types attributes, i.e. anything
     * after a semicolon. All content is assumed to be UTF-8 compatible.
     */
    public static normalizeMediaType(mediaType: string | undefined): string | undefined {
        if (mediaType === undefined) {
            return undefined;
        }
        return mediaType.split(";")[0].trim().toLowerCase();
    }

    /**
     * From a list of possible media types, choose the one we can handle best.
     *
     * The order of the given media types does not have any impact on the choice
     * made.
     */
    public static getPreferredMediaType(mediaTypes: Array<string>): string {
        /** According to OAS 3 we should default to json */
        if (!mediaTypes) {
            return "application/json";
        }

        const normalMediaTypes = mediaTypes.map(this.normalizeMediaType);
        let selectedMediaType: string | undefined = undefined;
        let selectedRank: number = -Infinity;
        for (const mediaType of normalMediaTypes) {
            if (supportedMediaTypes[mediaType!] > selectedRank) {
                selectedMediaType = mediaType;
                selectedRank = supportedMediaTypes[mediaType!];
            }
        }

        if (selectedMediaType === undefined) {
            throw new Error("None of the given media types are supported: " + mediaTypes.join(", "));
        }

        return selectedMediaType!;
    }

    /**
     * Convert data to a string according the given media type
     */
    public static stringify(data: any, mediaType: string): string {
        if (mediaType === "application/json") {
            return JSON.stringify(data);
        }

        throw new Error("The mediaType " + mediaType + " is not supported by ObjectSerializer.stringify.");
    }

    /**
     * Parse data from a string according to the given media type
     */
    public static parse(rawData: string, mediaType: string | undefined) {
        if (mediaType === undefined) {
            throw new Error("Cannot parse content. No Content-Type defined.");
        }

        if (mediaType === "application/json") {
            return JSON.parse(rawData);
        }

        throw new Error("The mediaType " + mediaType + " is not supported by ObjectSerializer.parse.");
    }
}
