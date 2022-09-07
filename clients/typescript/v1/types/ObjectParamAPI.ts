import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { ACLPolicy } from '../models/ACLPolicy';
import { ACLPolicyListStub } from '../models/ACLPolicyListStub';
import { ACLToken } from '../models/ACLToken';
import { ACLTokenListStub } from '../models/ACLTokenListStub';
import { Affinity } from '../models/Affinity';
import { AllocDeploymentStatus } from '../models/AllocDeploymentStatus';
import { AllocStopResponse } from '../models/AllocStopResponse';
import { AllocatedCpuResources } from '../models/AllocatedCpuResources';
import { AllocatedDeviceResource } from '../models/AllocatedDeviceResource';
import { AllocatedMemoryResources } from '../models/AllocatedMemoryResources';
import { AllocatedResources } from '../models/AllocatedResources';
import { AllocatedSharedResources } from '../models/AllocatedSharedResources';
import { AllocatedTaskResources } from '../models/AllocatedTaskResources';
import { Allocation } from '../models/Allocation';
import { AllocationListStub } from '../models/AllocationListStub';
import { AllocationMetric } from '../models/AllocationMetric';
import { Attribute } from '../models/Attribute';
import { AutopilotConfiguration } from '../models/AutopilotConfiguration';
import { CSIControllerInfo } from '../models/CSIControllerInfo';
import { CSIInfo } from '../models/CSIInfo';
import { CSIMountOptions } from '../models/CSIMountOptions';
import { CSINodeInfo } from '../models/CSINodeInfo';
import { CSIPlugin } from '../models/CSIPlugin';
import { CSIPluginListStub } from '../models/CSIPluginListStub';
import { CSISnapshot } from '../models/CSISnapshot';
import { CSISnapshotCreateRequest } from '../models/CSISnapshotCreateRequest';
import { CSISnapshotCreateResponse } from '../models/CSISnapshotCreateResponse';
import { CSISnapshotListResponse } from '../models/CSISnapshotListResponse';
import { CSITopology } from '../models/CSITopology';
import { CSITopologyRequest } from '../models/CSITopologyRequest';
import { CSIVolume } from '../models/CSIVolume';
import { CSIVolumeCapability } from '../models/CSIVolumeCapability';
import { CSIVolumeCreateRequest } from '../models/CSIVolumeCreateRequest';
import { CSIVolumeExternalStub } from '../models/CSIVolumeExternalStub';
import { CSIVolumeListExternalResponse } from '../models/CSIVolumeListExternalResponse';
import { CSIVolumeListStub } from '../models/CSIVolumeListStub';
import { CSIVolumeRegisterRequest } from '../models/CSIVolumeRegisterRequest';
import { CheckRestart } from '../models/CheckRestart';
import { Constraint } from '../models/Constraint';
import { Consul } from '../models/Consul';
import { ConsulConnect } from '../models/ConsulConnect';
import { ConsulExposeConfig } from '../models/ConsulExposeConfig';
import { ConsulExposePath } from '../models/ConsulExposePath';
import { ConsulGateway } from '../models/ConsulGateway';
import { ConsulGatewayBindAddress } from '../models/ConsulGatewayBindAddress';
import { ConsulGatewayProxy } from '../models/ConsulGatewayProxy';
import { ConsulGatewayTLSConfig } from '../models/ConsulGatewayTLSConfig';
import { ConsulIngressConfigEntry } from '../models/ConsulIngressConfigEntry';
import { ConsulIngressListener } from '../models/ConsulIngressListener';
import { ConsulIngressService } from '../models/ConsulIngressService';
import { ConsulLinkedService } from '../models/ConsulLinkedService';
import { ConsulMeshGateway } from '../models/ConsulMeshGateway';
import { ConsulProxy } from '../models/ConsulProxy';
import { ConsulSidecarService } from '../models/ConsulSidecarService';
import { ConsulTerminatingConfigEntry } from '../models/ConsulTerminatingConfigEntry';
import { ConsulUpstream } from '../models/ConsulUpstream';
import { DNSConfig } from '../models/DNSConfig';
import { Deployment } from '../models/Deployment';
import { DeploymentAllocHealthRequest } from '../models/DeploymentAllocHealthRequest';
import { DeploymentPauseRequest } from '../models/DeploymentPauseRequest';
import { DeploymentPromoteRequest } from '../models/DeploymentPromoteRequest';
import { DeploymentState } from '../models/DeploymentState';
import { DeploymentUnblockRequest } from '../models/DeploymentUnblockRequest';
import { DeploymentUpdateResponse } from '../models/DeploymentUpdateResponse';
import { DesiredTransition } from '../models/DesiredTransition';
import { DesiredUpdates } from '../models/DesiredUpdates';
import { DispatchPayloadConfig } from '../models/DispatchPayloadConfig';
import { DrainMetadata } from '../models/DrainMetadata';
import { DrainSpec } from '../models/DrainSpec';
import { DrainStrategy } from '../models/DrainStrategy';
import { DriverInfo } from '../models/DriverInfo';
import { EphemeralDisk } from '../models/EphemeralDisk';
import { EvalOptions } from '../models/EvalOptions';
import { Evaluation } from '../models/Evaluation';
import { EvaluationStub } from '../models/EvaluationStub';
import { FieldDiff } from '../models/FieldDiff';
import { FuzzyMatch } from '../models/FuzzyMatch';
import { FuzzySearchRequest } from '../models/FuzzySearchRequest';
import { FuzzySearchResponse } from '../models/FuzzySearchResponse';
import { GaugeValue } from '../models/GaugeValue';
import { HostNetworkInfo } from '../models/HostNetworkInfo';
import { HostVolumeInfo } from '../models/HostVolumeInfo';
import { Job } from '../models/Job';
import { JobChildrenSummary } from '../models/JobChildrenSummary';
import { JobDeregisterResponse } from '../models/JobDeregisterResponse';
import { JobDiff } from '../models/JobDiff';
import { JobDispatchRequest } from '../models/JobDispatchRequest';
import { JobDispatchResponse } from '../models/JobDispatchResponse';
import { JobEvaluateRequest } from '../models/JobEvaluateRequest';
import { JobListStub } from '../models/JobListStub';
import { JobPlanRequest } from '../models/JobPlanRequest';
import { JobPlanResponse } from '../models/JobPlanResponse';
import { JobRegisterRequest } from '../models/JobRegisterRequest';
import { JobRegisterResponse } from '../models/JobRegisterResponse';
import { JobRevertRequest } from '../models/JobRevertRequest';
import { JobScaleStatusResponse } from '../models/JobScaleStatusResponse';
import { JobStabilityRequest } from '../models/JobStabilityRequest';
import { JobStabilityResponse } from '../models/JobStabilityResponse';
import { JobSummary } from '../models/JobSummary';
import { JobValidateRequest } from '../models/JobValidateRequest';
import { JobValidateResponse } from '../models/JobValidateResponse';
import { JobVersionsResponse } from '../models/JobVersionsResponse';
import { JobsParseRequest } from '../models/JobsParseRequest';
import { LogConfig } from '../models/LogConfig';
import { MetricsSummary } from '../models/MetricsSummary';
import { MigrateStrategy } from '../models/MigrateStrategy';
import { Multiregion } from '../models/Multiregion';
import { MultiregionRegion } from '../models/MultiregionRegion';
import { MultiregionStrategy } from '../models/MultiregionStrategy';
import { Namespace } from '../models/Namespace';
import { NamespaceCapabilities } from '../models/NamespaceCapabilities';
import { NetworkResource } from '../models/NetworkResource';
import { Node } from '../models/Node';
import { NodeCpuResources } from '../models/NodeCpuResources';
import { NodeDevice } from '../models/NodeDevice';
import { NodeDeviceLocality } from '../models/NodeDeviceLocality';
import { NodeDeviceResource } from '../models/NodeDeviceResource';
import { NodeDiskResources } from '../models/NodeDiskResources';
import { NodeDrainUpdateResponse } from '../models/NodeDrainUpdateResponse';
import { NodeEligibilityUpdateResponse } from '../models/NodeEligibilityUpdateResponse';
import { NodeEvent } from '../models/NodeEvent';
import { NodeListStub } from '../models/NodeListStub';
import { NodeMemoryResources } from '../models/NodeMemoryResources';
import { NodePurgeResponse } from '../models/NodePurgeResponse';
import { NodeReservedCpuResources } from '../models/NodeReservedCpuResources';
import { NodeReservedDiskResources } from '../models/NodeReservedDiskResources';
import { NodeReservedMemoryResources } from '../models/NodeReservedMemoryResources';
import { NodeReservedNetworkResources } from '../models/NodeReservedNetworkResources';
import { NodeReservedResources } from '../models/NodeReservedResources';
import { NodeResources } from '../models/NodeResources';
import { NodeScoreMeta } from '../models/NodeScoreMeta';
import { NodeUpdateDrainRequest } from '../models/NodeUpdateDrainRequest';
import { NodeUpdateEligibilityRequest } from '../models/NodeUpdateEligibilityRequest';
import { ObjectDiff } from '../models/ObjectDiff';
import { OneTimeToken } from '../models/OneTimeToken';
import { OneTimeTokenExchangeRequest } from '../models/OneTimeTokenExchangeRequest';
import { OperatorHealthReply } from '../models/OperatorHealthReply';
import { ParameterizedJobConfig } from '../models/ParameterizedJobConfig';
import { PeriodicConfig } from '../models/PeriodicConfig';
import { PeriodicForceResponse } from '../models/PeriodicForceResponse';
import { PlanAnnotations } from '../models/PlanAnnotations';
import { PointValue } from '../models/PointValue';
import { Port } from '../models/Port';
import { PortMapping } from '../models/PortMapping';
import { PreemptionConfig } from '../models/PreemptionConfig';
import { QuotaLimit } from '../models/QuotaLimit';
import { QuotaSpec } from '../models/QuotaSpec';
import { RaftConfiguration } from '../models/RaftConfiguration';
import { RaftServer } from '../models/RaftServer';
import { RequestedDevice } from '../models/RequestedDevice';
import { RescheduleEvent } from '../models/RescheduleEvent';
import { ReschedulePolicy } from '../models/ReschedulePolicy';
import { RescheduleTracker } from '../models/RescheduleTracker';
import { Resources } from '../models/Resources';
import { RestartPolicy } from '../models/RestartPolicy';
import { SampledValue } from '../models/SampledValue';
import { ScalingEvent } from '../models/ScalingEvent';
import { ScalingPolicy } from '../models/ScalingPolicy';
import { ScalingPolicyListStub } from '../models/ScalingPolicyListStub';
import { ScalingRequest } from '../models/ScalingRequest';
import { SchedulerConfiguration } from '../models/SchedulerConfiguration';
import { SchedulerConfigurationResponse } from '../models/SchedulerConfigurationResponse';
import { SchedulerSetConfigurationResponse } from '../models/SchedulerSetConfigurationResponse';
import { SearchRequest } from '../models/SearchRequest';
import { SearchResponse } from '../models/SearchResponse';
import { ServerHealth } from '../models/ServerHealth';
import { Service } from '../models/Service';
import { ServiceCheck } from '../models/ServiceCheck';
import { ServiceRegistration } from '../models/ServiceRegistration';
import { SidecarTask } from '../models/SidecarTask';
import { Spread } from '../models/Spread';
import { SpreadTarget } from '../models/SpreadTarget';
import { Task } from '../models/Task';
import { TaskArtifact } from '../models/TaskArtifact';
import { TaskCSIPluginConfig } from '../models/TaskCSIPluginConfig';
import { TaskDiff } from '../models/TaskDiff';
import { TaskEvent } from '../models/TaskEvent';
import { TaskGroup } from '../models/TaskGroup';
import { TaskGroupDiff } from '../models/TaskGroupDiff';
import { TaskGroupScaleStatus } from '../models/TaskGroupScaleStatus';
import { TaskGroupSummary } from '../models/TaskGroupSummary';
import { TaskHandle } from '../models/TaskHandle';
import { TaskLifecycle } from '../models/TaskLifecycle';
import { TaskState } from '../models/TaskState';
import { Template } from '../models/Template';
import { UpdateStrategy } from '../models/UpdateStrategy';
import { Vault } from '../models/Vault';
import { VolumeMount } from '../models/VolumeMount';
import { VolumeRequest } from '../models/VolumeRequest';
import { WaitConfig } from '../models/WaitConfig';

import { ObservableACLApi } from "./ObservableAPI";
import { ACLApiRequestFactory, ACLApiResponseProcessor} from "../apis/ACLApi";

export interface ACLApiDeleteACLPolicyRequest {
    /**
     * The ACL policy name.
     * @type string
     * @memberof ACLApideleteACLPolicy
     */
    policyName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApideleteACLPolicy
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApideleteACLPolicy
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApideleteACLPolicy
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApideleteACLPolicy
     */
    idempotencyToken?: string
}

export interface ACLApiDeleteACLTokenRequest {
    /**
     * The token accessor ID.
     * @type string
     * @memberof ACLApideleteACLToken
     */
    tokenAccessor: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApideleteACLToken
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApideleteACLToken
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApideleteACLToken
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApideleteACLToken
     */
    idempotencyToken?: string
}

export interface ACLApiGetACLPoliciesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ACLApigetACLPolicies
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ACLApigetACLPolicies
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ACLApigetACLPolicies
     */
    nextToken?: string
}

export interface ACLApiGetACLPolicyRequest {
    /**
     * The ACL policy name.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    policyName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ACLApigetACLPolicy
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ACLApigetACLPolicy
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ACLApigetACLPolicy
     */
    nextToken?: string
}

export interface ACLApiGetACLTokenRequest {
    /**
     * The token accessor ID.
     * @type string
     * @memberof ACLApigetACLToken
     */
    tokenAccessor: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApigetACLToken
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApigetACLToken
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ACLApigetACLToken
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ACLApigetACLToken
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ACLApigetACLToken
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ACLApigetACLToken
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApigetACLToken
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ACLApigetACLToken
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ACLApigetACLToken
     */
    nextToken?: string
}

export interface ACLApiGetACLTokenSelfRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ACLApigetACLTokenSelf
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ACLApigetACLTokenSelf
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ACLApigetACLTokenSelf
     */
    nextToken?: string
}

export interface ACLApiGetACLTokensRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ACLApigetACLTokens
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ACLApigetACLTokens
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ACLApigetACLTokens
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ACLApigetACLTokens
     */
    nextToken?: string
}

export interface ACLApiPostACLBootstrapRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApipostACLBootstrap
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApipostACLBootstrap
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApipostACLBootstrap
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApipostACLBootstrap
     */
    idempotencyToken?: string
}

export interface ACLApiPostACLPolicyRequest {
    /**
     * The ACL policy name.
     * @type string
     * @memberof ACLApipostACLPolicy
     */
    policyName: string
    /**
     * 
     * @type ACLPolicy
     * @memberof ACLApipostACLPolicy
     */
    aCLPolicy: ACLPolicy
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApipostACLPolicy
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApipostACLPolicy
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApipostACLPolicy
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApipostACLPolicy
     */
    idempotencyToken?: string
}

export interface ACLApiPostACLTokenRequest {
    /**
     * The token accessor ID.
     * @type string
     * @memberof ACLApipostACLToken
     */
    tokenAccessor: string
    /**
     * 
     * @type ACLToken
     * @memberof ACLApipostACLToken
     */
    aCLToken: ACLToken
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApipostACLToken
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApipostACLToken
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApipostACLToken
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApipostACLToken
     */
    idempotencyToken?: string
}

export interface ACLApiPostACLTokenOnetimeRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApipostACLTokenOnetime
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApipostACLTokenOnetime
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApipostACLTokenOnetime
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApipostACLTokenOnetime
     */
    idempotencyToken?: string
}

export interface ACLApiPostACLTokenOnetimeExchangeRequest {
    /**
     * 
     * @type OneTimeTokenExchangeRequest
     * @memberof ACLApipostACLTokenOnetimeExchange
     */
    oneTimeTokenExchangeRequest: OneTimeTokenExchangeRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ACLApipostACLTokenOnetimeExchange
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ACLApipostACLTokenOnetimeExchange
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ACLApipostACLTokenOnetimeExchange
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof ACLApipostACLTokenOnetimeExchange
     */
    idempotencyToken?: string
}

export class ObjectACLApi {
    private api: ObservableACLApi

    public constructor(configuration: Configuration, requestFactory?: ACLApiRequestFactory, responseProcessor?: ACLApiResponseProcessor) {
        this.api = new ObservableACLApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public deleteACLPolicy(param: ACLApiDeleteACLPolicyRequest, options?: Configuration): Promise<void> {
        return this.api.deleteACLPolicy(param.policyName, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public deleteACLToken(param: ACLApiDeleteACLTokenRequest, options?: Configuration): Promise<void> {
        return this.api.deleteACLToken(param.tokenAccessor, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getACLPolicies(param: ACLApiGetACLPoliciesRequest = {}, options?: Configuration): Promise<Array<ACLPolicyListStub>> {
        return this.api.getACLPolicies(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getACLPolicy(param: ACLApiGetACLPolicyRequest, options?: Configuration): Promise<ACLPolicy> {
        return this.api.getACLPolicy(param.policyName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getACLToken(param: ACLApiGetACLTokenRequest, options?: Configuration): Promise<ACLToken> {
        return this.api.getACLToken(param.tokenAccessor, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getACLTokenSelf(param: ACLApiGetACLTokenSelfRequest = {}, options?: Configuration): Promise<ACLToken> {
        return this.api.getACLTokenSelf(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getACLTokens(param: ACLApiGetACLTokensRequest = {}, options?: Configuration): Promise<Array<ACLTokenListStub>> {
        return this.api.getACLTokens(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postACLBootstrap(param: ACLApiPostACLBootstrapRequest = {}, options?: Configuration): Promise<ACLToken> {
        return this.api.postACLBootstrap(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postACLPolicy(param: ACLApiPostACLPolicyRequest, options?: Configuration): Promise<void> {
        return this.api.postACLPolicy(param.policyName, param.aCLPolicy, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postACLToken(param: ACLApiPostACLTokenRequest, options?: Configuration): Promise<ACLToken> {
        return this.api.postACLToken(param.tokenAccessor, param.aCLToken, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postACLTokenOnetime(param: ACLApiPostACLTokenOnetimeRequest = {}, options?: Configuration): Promise<OneTimeToken> {
        return this.api.postACLTokenOnetime(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postACLTokenOnetimeExchange(param: ACLApiPostACLTokenOnetimeExchangeRequest, options?: Configuration): Promise<ACLToken> {
        return this.api.postACLTokenOnetimeExchange(param.oneTimeTokenExchangeRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableAllocationsApi } from "./ObservableAPI";
import { AllocationsApiRequestFactory, AllocationsApiResponseProcessor} from "../apis/AllocationsApi";

export interface AllocationsApiGetAllocationRequest {
    /**
     * Allocation ID.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    allocID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof AllocationsApigetAllocation
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof AllocationsApigetAllocation
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof AllocationsApigetAllocation
     */
    nextToken?: string
}

export interface AllocationsApiGetAllocationServicesRequest {
    /**
     * Allocation ID.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    allocID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof AllocationsApigetAllocationServices
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof AllocationsApigetAllocationServices
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof AllocationsApigetAllocationServices
     */
    nextToken?: string
}

export interface AllocationsApiGetAllocationsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof AllocationsApigetAllocations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof AllocationsApigetAllocations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof AllocationsApigetAllocations
     */
    nextToken?: string
    /**
     * Flag indicating whether to include resources in response.
     * @type boolean
     * @memberof AllocationsApigetAllocations
     */
    resources?: boolean
    /**
     * Flag indicating whether to include task states in response.
     * @type boolean
     * @memberof AllocationsApigetAllocations
     */
    taskStates?: boolean
}

export interface AllocationsApiPostAllocationStopRequest {
    /**
     * Allocation ID.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    allocID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof AllocationsApipostAllocationStop
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof AllocationsApipostAllocationStop
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof AllocationsApipostAllocationStop
     */
    nextToken?: string
    /**
     * Flag indicating whether to delay shutdown when requesting an allocation stop.
     * @type boolean
     * @memberof AllocationsApipostAllocationStop
     */
    noShutdownDelay?: boolean
}

export class ObjectAllocationsApi {
    private api: ObservableAllocationsApi

    public constructor(configuration: Configuration, requestFactory?: AllocationsApiRequestFactory, responseProcessor?: AllocationsApiResponseProcessor) {
        this.api = new ObservableAllocationsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getAllocation(param: AllocationsApiGetAllocationRequest, options?: Configuration): Promise<Allocation> {
        return this.api.getAllocation(param.allocID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getAllocationServices(param: AllocationsApiGetAllocationServicesRequest, options?: Configuration): Promise<Array<ServiceRegistration>> {
        return this.api.getAllocationServices(param.allocID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getAllocations(param: AllocationsApiGetAllocationsRequest = {}, options?: Configuration): Promise<Array<AllocationListStub>> {
        return this.api.getAllocations(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.resources, param.taskStates,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postAllocationStop(param: AllocationsApiPostAllocationStopRequest, options?: Configuration): Promise<AllocStopResponse> {
        return this.api.postAllocationStop(param.allocID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.noShutdownDelay,  options).toPromise();
    }

}

import { ObservableDeploymentsApi } from "./ObservableAPI";
import { DeploymentsApiRequestFactory, DeploymentsApiResponseProcessor} from "../apis/DeploymentsApi";

export interface DeploymentsApiGetDeploymentRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    deploymentID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof DeploymentsApigetDeployment
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof DeploymentsApigetDeployment
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof DeploymentsApigetDeployment
     */
    nextToken?: string
}

export interface DeploymentsApiGetDeploymentAllocationsRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    deploymentID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof DeploymentsApigetDeploymentAllocations
     */
    nextToken?: string
}

export interface DeploymentsApiGetDeploymentsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof DeploymentsApigetDeployments
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof DeploymentsApigetDeployments
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof DeploymentsApigetDeployments
     */
    nextToken?: string
}

export interface DeploymentsApiPostDeploymentAllocationHealthRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    deploymentID: string
    /**
     * 
     * @type DeploymentAllocHealthRequest
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    deploymentAllocHealthRequest: DeploymentAllocHealthRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof DeploymentsApipostDeploymentAllocationHealth
     */
    idempotencyToken?: string
}

export interface DeploymentsApiPostDeploymentFailRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApipostDeploymentFail
     */
    deploymentID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApipostDeploymentFail
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApipostDeploymentFail
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApipostDeploymentFail
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof DeploymentsApipostDeploymentFail
     */
    idempotencyToken?: string
}

export interface DeploymentsApiPostDeploymentPauseRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApipostDeploymentPause
     */
    deploymentID: string
    /**
     * 
     * @type DeploymentPauseRequest
     * @memberof DeploymentsApipostDeploymentPause
     */
    deploymentPauseRequest: DeploymentPauseRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApipostDeploymentPause
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApipostDeploymentPause
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApipostDeploymentPause
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof DeploymentsApipostDeploymentPause
     */
    idempotencyToken?: string
}

export interface DeploymentsApiPostDeploymentPromoteRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApipostDeploymentPromote
     */
    deploymentID: string
    /**
     * 
     * @type DeploymentPromoteRequest
     * @memberof DeploymentsApipostDeploymentPromote
     */
    deploymentPromoteRequest: DeploymentPromoteRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApipostDeploymentPromote
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApipostDeploymentPromote
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApipostDeploymentPromote
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof DeploymentsApipostDeploymentPromote
     */
    idempotencyToken?: string
}

export interface DeploymentsApiPostDeploymentUnblockRequest {
    /**
     * Deployment ID.
     * @type string
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    deploymentID: string
    /**
     * 
     * @type DeploymentUnblockRequest
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    deploymentUnblockRequest: DeploymentUnblockRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof DeploymentsApipostDeploymentUnblock
     */
    idempotencyToken?: string
}

export class ObjectDeploymentsApi {
    private api: ObservableDeploymentsApi

    public constructor(configuration: Configuration, requestFactory?: DeploymentsApiRequestFactory, responseProcessor?: DeploymentsApiResponseProcessor) {
        this.api = new ObservableDeploymentsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getDeployment(param: DeploymentsApiGetDeploymentRequest, options?: Configuration): Promise<Deployment> {
        return this.api.getDeployment(param.deploymentID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getDeploymentAllocations(param: DeploymentsApiGetDeploymentAllocationsRequest, options?: Configuration): Promise<Array<AllocationListStub>> {
        return this.api.getDeploymentAllocations(param.deploymentID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getDeployments(param: DeploymentsApiGetDeploymentsRequest = {}, options?: Configuration): Promise<Array<Deployment>> {
        return this.api.getDeployments(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postDeploymentAllocationHealth(param: DeploymentsApiPostDeploymentAllocationHealthRequest, options?: Configuration): Promise<DeploymentUpdateResponse> {
        return this.api.postDeploymentAllocationHealth(param.deploymentID, param.deploymentAllocHealthRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postDeploymentFail(param: DeploymentsApiPostDeploymentFailRequest, options?: Configuration): Promise<DeploymentUpdateResponse> {
        return this.api.postDeploymentFail(param.deploymentID, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postDeploymentPause(param: DeploymentsApiPostDeploymentPauseRequest, options?: Configuration): Promise<DeploymentUpdateResponse> {
        return this.api.postDeploymentPause(param.deploymentID, param.deploymentPauseRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postDeploymentPromote(param: DeploymentsApiPostDeploymentPromoteRequest, options?: Configuration): Promise<DeploymentUpdateResponse> {
        return this.api.postDeploymentPromote(param.deploymentID, param.deploymentPromoteRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postDeploymentUnblock(param: DeploymentsApiPostDeploymentUnblockRequest, options?: Configuration): Promise<DeploymentUpdateResponse> {
        return this.api.postDeploymentUnblock(param.deploymentID, param.deploymentUnblockRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableEnterpriseApi } from "./ObservableAPI";
import { EnterpriseApiRequestFactory, EnterpriseApiResponseProcessor} from "../apis/EnterpriseApi";

export interface EnterpriseApiCreateQuotaSpecRequest {
    /**
     * 
     * @type QuotaSpec
     * @memberof EnterpriseApicreateQuotaSpec
     */
    quotaSpec: QuotaSpec
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EnterpriseApicreateQuotaSpec
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EnterpriseApicreateQuotaSpec
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EnterpriseApicreateQuotaSpec
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof EnterpriseApicreateQuotaSpec
     */
    idempotencyToken?: string
}

export interface EnterpriseApiDeleteQuotaSpecRequest {
    /**
     * The quota spec identifier.
     * @type string
     * @memberof EnterpriseApideleteQuotaSpec
     */
    specName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EnterpriseApideleteQuotaSpec
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EnterpriseApideleteQuotaSpec
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EnterpriseApideleteQuotaSpec
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof EnterpriseApideleteQuotaSpec
     */
    idempotencyToken?: string
}

export interface EnterpriseApiGetQuotaSpecRequest {
    /**
     * The quota spec identifier.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    specName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof EnterpriseApigetQuotaSpec
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof EnterpriseApigetQuotaSpec
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof EnterpriseApigetQuotaSpec
     */
    nextToken?: string
}

export interface EnterpriseApiGetQuotasRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof EnterpriseApigetQuotas
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof EnterpriseApigetQuotas
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof EnterpriseApigetQuotas
     */
    nextToken?: string
}

export interface EnterpriseApiPostQuotaSpecRequest {
    /**
     * The quota spec identifier.
     * @type string
     * @memberof EnterpriseApipostQuotaSpec
     */
    specName: string
    /**
     * 
     * @type QuotaSpec
     * @memberof EnterpriseApipostQuotaSpec
     */
    quotaSpec: QuotaSpec
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EnterpriseApipostQuotaSpec
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EnterpriseApipostQuotaSpec
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EnterpriseApipostQuotaSpec
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof EnterpriseApipostQuotaSpec
     */
    idempotencyToken?: string
}

export class ObjectEnterpriseApi {
    private api: ObservableEnterpriseApi

    public constructor(configuration: Configuration, requestFactory?: EnterpriseApiRequestFactory, responseProcessor?: EnterpriseApiResponseProcessor) {
        this.api = new ObservableEnterpriseApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public createQuotaSpec(param: EnterpriseApiCreateQuotaSpecRequest, options?: Configuration): Promise<void> {
        return this.api.createQuotaSpec(param.quotaSpec, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public deleteQuotaSpec(param: EnterpriseApiDeleteQuotaSpecRequest, options?: Configuration): Promise<void> {
        return this.api.deleteQuotaSpec(param.specName, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getQuotaSpec(param: EnterpriseApiGetQuotaSpecRequest, options?: Configuration): Promise<QuotaSpec> {
        return this.api.getQuotaSpec(param.specName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getQuotas(param: EnterpriseApiGetQuotasRequest = {}, options?: Configuration): Promise<Array<any>> {
        return this.api.getQuotas(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postQuotaSpec(param: EnterpriseApiPostQuotaSpecRequest, options?: Configuration): Promise<void> {
        return this.api.postQuotaSpec(param.specName, param.quotaSpec, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableEvaluationsApi } from "./ObservableAPI";
import { EvaluationsApiRequestFactory, EvaluationsApiResponseProcessor} from "../apis/EvaluationsApi";

export interface EvaluationsApiGetEvaluationRequest {
    /**
     * Evaluation ID.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    evalID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof EvaluationsApigetEvaluation
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof EvaluationsApigetEvaluation
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof EvaluationsApigetEvaluation
     */
    nextToken?: string
}

export interface EvaluationsApiGetEvaluationAllocationsRequest {
    /**
     * Evaluation ID.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    evalID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof EvaluationsApigetEvaluationAllocations
     */
    nextToken?: string
}

export interface EvaluationsApiGetEvaluationsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof EvaluationsApigetEvaluations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof EvaluationsApigetEvaluations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof EvaluationsApigetEvaluations
     */
    nextToken?: string
}

export class ObjectEvaluationsApi {
    private api: ObservableEvaluationsApi

    public constructor(configuration: Configuration, requestFactory?: EvaluationsApiRequestFactory, responseProcessor?: EvaluationsApiResponseProcessor) {
        this.api = new ObservableEvaluationsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getEvaluation(param: EvaluationsApiGetEvaluationRequest, options?: Configuration): Promise<Evaluation> {
        return this.api.getEvaluation(param.evalID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getEvaluationAllocations(param: EvaluationsApiGetEvaluationAllocationsRequest, options?: Configuration): Promise<Array<AllocationListStub>> {
        return this.api.getEvaluationAllocations(param.evalID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getEvaluations(param: EvaluationsApiGetEvaluationsRequest = {}, options?: Configuration): Promise<Array<Evaluation>> {
        return this.api.getEvaluations(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableJobsApi } from "./ObservableAPI";
import { JobsApiRequestFactory, JobsApiResponseProcessor} from "../apis/JobsApi";

export interface JobsApiDeleteJobRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApideleteJob
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApideleteJob
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApideleteJob
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApideleteJob
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApideleteJob
     */
    idempotencyToken?: string
    /**
     * Boolean flag indicating whether to purge allocations of the job after deleting.
     * @type boolean
     * @memberof JobsApideleteJob
     */
    purge?: boolean
    /**
     * Boolean flag indicating whether the operation should apply to all instances of the job globally.
     * @type boolean
     * @memberof JobsApideleteJob
     */
    global?: boolean
}

export interface JobsApiGetJobRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJob
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJob
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJob
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJob
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJob
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJob
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJob
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJob
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJob
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJob
     */
    nextToken?: string
}

export interface JobsApiGetJobAllocationsRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobAllocations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobAllocations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobAllocations
     */
    nextToken?: string
    /**
     * Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered.
     * @type boolean
     * @memberof JobsApigetJobAllocations
     */
    all?: boolean
}

export interface JobsApiGetJobDeploymentRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobDeployment
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobDeployment
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobDeployment
     */
    nextToken?: string
}

export interface JobsApiGetJobDeploymentsRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobDeployments
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobDeployments
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobDeployments
     */
    nextToken?: string
    /**
     * Flag indicating whether to constrain by job creation index or not.
     * @type number
     * @memberof JobsApigetJobDeployments
     */
    all?: number
}

export interface JobsApiGetJobEvaluationsRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobEvaluations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobEvaluations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobEvaluations
     */
    nextToken?: string
}

export interface JobsApiGetJobScaleStatusRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobScaleStatus
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobScaleStatus
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobScaleStatus
     */
    nextToken?: string
}

export interface JobsApiGetJobSummaryRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobSummary
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobSummary
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobSummary
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobSummary
     */
    nextToken?: string
}

export interface JobsApiGetJobVersionsRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobVersions
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobVersions
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobVersions
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobVersions
     */
    nextToken?: string
    /**
     * Boolean flag indicating whether to compute job diffs.
     * @type boolean
     * @memberof JobsApigetJobVersions
     */
    diffs?: boolean
}

export interface JobsApiGetJobsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApigetJobs
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApigetJobs
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof JobsApigetJobs
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof JobsApigetJobs
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof JobsApigetJobs
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof JobsApigetJobs
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApigetJobs
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof JobsApigetJobs
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof JobsApigetJobs
     */
    nextToken?: string
}

export interface JobsApiPostJobRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJob
     */
    jobName: string
    /**
     * 
     * @type JobRegisterRequest
     * @memberof JobsApipostJob
     */
    jobRegisterRequest: JobRegisterRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJob
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJob
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJob
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJob
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobDispatchRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobDispatch
     */
    jobName: string
    /**
     * 
     * @type JobDispatchRequest
     * @memberof JobsApipostJobDispatch
     */
    jobDispatchRequest: JobDispatchRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobDispatch
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobDispatch
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobDispatch
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobDispatch
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobEvaluateRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobEvaluate
     */
    jobName: string
    /**
     * 
     * @type JobEvaluateRequest
     * @memberof JobsApipostJobEvaluate
     */
    jobEvaluateRequest: JobEvaluateRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobEvaluate
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobEvaluate
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobEvaluate
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobEvaluate
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobParseRequest {
    /**
     * 
     * @type JobsParseRequest
     * @memberof JobsApipostJobParse
     */
    jobsParseRequest: JobsParseRequest
}

export interface JobsApiPostJobPeriodicForceRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobPeriodicForce
     */
    jobName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobPeriodicForce
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobPeriodicForce
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobPeriodicForce
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobPeriodicForce
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobPlanRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobPlan
     */
    jobName: string
    /**
     * 
     * @type JobPlanRequest
     * @memberof JobsApipostJobPlan
     */
    jobPlanRequest: JobPlanRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobPlan
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobPlan
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobPlan
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobPlan
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobRevertRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobRevert
     */
    jobName: string
    /**
     * 
     * @type JobRevertRequest
     * @memberof JobsApipostJobRevert
     */
    jobRevertRequest: JobRevertRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobRevert
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobRevert
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobRevert
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobRevert
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobScalingRequestRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobScalingRequest
     */
    jobName: string
    /**
     * 
     * @type ScalingRequest
     * @memberof JobsApipostJobScalingRequest
     */
    scalingRequest: ScalingRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobScalingRequest
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobScalingRequest
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobScalingRequest
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobScalingRequest
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobStabilityRequest {
    /**
     * The job identifier.
     * @type string
     * @memberof JobsApipostJobStability
     */
    jobName: string
    /**
     * 
     * @type JobStabilityRequest
     * @memberof JobsApipostJobStability
     */
    jobStabilityRequest: JobStabilityRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobStability
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobStability
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobStability
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobStability
     */
    idempotencyToken?: string
}

export interface JobsApiPostJobValidateRequestRequest {
    /**
     * 
     * @type JobValidateRequest
     * @memberof JobsApipostJobValidateRequest
     */
    jobValidateRequest: JobValidateRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApipostJobValidateRequest
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApipostJobValidateRequest
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApipostJobValidateRequest
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApipostJobValidateRequest
     */
    idempotencyToken?: string
}

export interface JobsApiRegisterJobRequest {
    /**
     * 
     * @type JobRegisterRequest
     * @memberof JobsApiregisterJob
     */
    jobRegisterRequest: JobRegisterRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof JobsApiregisterJob
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof JobsApiregisterJob
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof JobsApiregisterJob
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof JobsApiregisterJob
     */
    idempotencyToken?: string
}

export class ObjectJobsApi {
    private api: ObservableJobsApi

    public constructor(configuration: Configuration, requestFactory?: JobsApiRequestFactory, responseProcessor?: JobsApiResponseProcessor) {
        this.api = new ObservableJobsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public deleteJob(param: JobsApiDeleteJobRequest, options?: Configuration): Promise<JobDeregisterResponse> {
        return this.api.deleteJob(param.jobName, param.region, param.namespace, param.xNomadToken, param.idempotencyToken, param.purge, param.global,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJob(param: JobsApiGetJobRequest, options?: Configuration): Promise<Job> {
        return this.api.getJob(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobAllocations(param: JobsApiGetJobAllocationsRequest, options?: Configuration): Promise<Array<AllocationListStub>> {
        return this.api.getJobAllocations(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.all,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobDeployment(param: JobsApiGetJobDeploymentRequest, options?: Configuration): Promise<Deployment> {
        return this.api.getJobDeployment(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobDeployments(param: JobsApiGetJobDeploymentsRequest, options?: Configuration): Promise<Array<Deployment>> {
        return this.api.getJobDeployments(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.all,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobEvaluations(param: JobsApiGetJobEvaluationsRequest, options?: Configuration): Promise<Array<Evaluation>> {
        return this.api.getJobEvaluations(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobScaleStatus(param: JobsApiGetJobScaleStatusRequest, options?: Configuration): Promise<JobScaleStatusResponse> {
        return this.api.getJobScaleStatus(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobSummary(param: JobsApiGetJobSummaryRequest, options?: Configuration): Promise<JobSummary> {
        return this.api.getJobSummary(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobVersions(param: JobsApiGetJobVersionsRequest, options?: Configuration): Promise<JobVersionsResponse> {
        return this.api.getJobVersions(param.jobName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.diffs,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getJobs(param: JobsApiGetJobsRequest = {}, options?: Configuration): Promise<Array<JobListStub>> {
        return this.api.getJobs(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJob(param: JobsApiPostJobRequest, options?: Configuration): Promise<JobRegisterResponse> {
        return this.api.postJob(param.jobName, param.jobRegisterRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobDispatch(param: JobsApiPostJobDispatchRequest, options?: Configuration): Promise<JobDispatchResponse> {
        return this.api.postJobDispatch(param.jobName, param.jobDispatchRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobEvaluate(param: JobsApiPostJobEvaluateRequest, options?: Configuration): Promise<JobRegisterResponse> {
        return this.api.postJobEvaluate(param.jobName, param.jobEvaluateRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobParse(param: JobsApiPostJobParseRequest, options?: Configuration): Promise<Job> {
        return this.api.postJobParse(param.jobsParseRequest,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobPeriodicForce(param: JobsApiPostJobPeriodicForceRequest, options?: Configuration): Promise<PeriodicForceResponse> {
        return this.api.postJobPeriodicForce(param.jobName, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobPlan(param: JobsApiPostJobPlanRequest, options?: Configuration): Promise<JobPlanResponse> {
        return this.api.postJobPlan(param.jobName, param.jobPlanRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobRevert(param: JobsApiPostJobRevertRequest, options?: Configuration): Promise<JobRegisterResponse> {
        return this.api.postJobRevert(param.jobName, param.jobRevertRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobScalingRequest(param: JobsApiPostJobScalingRequestRequest, options?: Configuration): Promise<JobRegisterResponse> {
        return this.api.postJobScalingRequest(param.jobName, param.scalingRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobStability(param: JobsApiPostJobStabilityRequest, options?: Configuration): Promise<JobStabilityResponse> {
        return this.api.postJobStability(param.jobName, param.jobStabilityRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postJobValidateRequest(param: JobsApiPostJobValidateRequestRequest, options?: Configuration): Promise<JobValidateResponse> {
        return this.api.postJobValidateRequest(param.jobValidateRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public registerJob(param: JobsApiRegisterJobRequest, options?: Configuration): Promise<JobRegisterResponse> {
        return this.api.registerJob(param.jobRegisterRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableMetricsApi } from "./ObservableAPI";
import { MetricsApiRequestFactory, MetricsApiResponseProcessor} from "../apis/MetricsApi";

export interface MetricsApiGetMetricsSummaryRequest {
    /**
     * The format the user requested for the metrics summary (e.g. prometheus)
     * @type string
     * @memberof MetricsApigetMetricsSummary
     */
    format?: string
}

export class ObjectMetricsApi {
    private api: ObservableMetricsApi

    public constructor(configuration: Configuration, requestFactory?: MetricsApiRequestFactory, responseProcessor?: MetricsApiResponseProcessor) {
        this.api = new ObservableMetricsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getMetricsSummary(param: MetricsApiGetMetricsSummaryRequest = {}, options?: Configuration): Promise<MetricsSummary> {
        return this.api.getMetricsSummary(param.format,  options).toPromise();
    }

}

import { ObservableNamespacesApi } from "./ObservableAPI";
import { NamespacesApiRequestFactory, NamespacesApiResponseProcessor} from "../apis/NamespacesApi";

export interface NamespacesApiCreateNamespaceRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NamespacesApicreateNamespace
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NamespacesApicreateNamespace
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NamespacesApicreateNamespace
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof NamespacesApicreateNamespace
     */
    idempotencyToken?: string
}

export interface NamespacesApiDeleteNamespaceRequest {
    /**
     * The namespace identifier.
     * @type string
     * @memberof NamespacesApideleteNamespace
     */
    namespaceName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NamespacesApideleteNamespace
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NamespacesApideleteNamespace
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NamespacesApideleteNamespace
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof NamespacesApideleteNamespace
     */
    idempotencyToken?: string
}

export interface NamespacesApiGetNamespaceRequest {
    /**
     * The namespace identifier.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    namespaceName: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NamespacesApigetNamespace
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NamespacesApigetNamespace
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NamespacesApigetNamespace
     */
    nextToken?: string
}

export interface NamespacesApiGetNamespacesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NamespacesApigetNamespaces
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NamespacesApigetNamespaces
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NamespacesApigetNamespaces
     */
    nextToken?: string
}

export interface NamespacesApiPostNamespaceRequest {
    /**
     * The namespace identifier.
     * @type string
     * @memberof NamespacesApipostNamespace
     */
    namespaceName: string
    /**
     * 
     * @type Namespace
     * @memberof NamespacesApipostNamespace
     */
    namespace2: Namespace
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NamespacesApipostNamespace
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NamespacesApipostNamespace
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NamespacesApipostNamespace
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof NamespacesApipostNamespace
     */
    idempotencyToken?: string
}

export class ObjectNamespacesApi {
    private api: ObservableNamespacesApi

    public constructor(configuration: Configuration, requestFactory?: NamespacesApiRequestFactory, responseProcessor?: NamespacesApiResponseProcessor) {
        this.api = new ObservableNamespacesApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public createNamespace(param: NamespacesApiCreateNamespaceRequest = {}, options?: Configuration): Promise<void> {
        return this.api.createNamespace(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public deleteNamespace(param: NamespacesApiDeleteNamespaceRequest, options?: Configuration): Promise<void> {
        return this.api.deleteNamespace(param.namespaceName, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getNamespace(param: NamespacesApiGetNamespaceRequest, options?: Configuration): Promise<Namespace> {
        return this.api.getNamespace(param.namespaceName, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getNamespaces(param: NamespacesApiGetNamespacesRequest = {}, options?: Configuration): Promise<Array<Namespace>> {
        return this.api.getNamespaces(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postNamespace(param: NamespacesApiPostNamespaceRequest, options?: Configuration): Promise<void> {
        return this.api.postNamespace(param.namespaceName, param.namespace2, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableNodesApi } from "./ObservableAPI";
import { NodesApiRequestFactory, NodesApiResponseProcessor} from "../apis/NodesApi";

export interface NodesApiGetNodeRequest {
    /**
     * The ID of the node.
     * @type string
     * @memberof NodesApigetNode
     */
    nodeId: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApigetNode
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApigetNode
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApigetNode
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApigetNode
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApigetNode
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApigetNode
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApigetNode
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApigetNode
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApigetNode
     */
    nextToken?: string
}

export interface NodesApiGetNodeAllocationsRequest {
    /**
     * The ID of the node.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    nodeId: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApigetNodeAllocations
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApigetNodeAllocations
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApigetNodeAllocations
     */
    nextToken?: string
}

export interface NodesApiGetNodesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApigetNodes
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApigetNodes
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApigetNodes
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApigetNodes
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApigetNodes
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApigetNodes
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApigetNodes
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApigetNodes
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApigetNodes
     */
    nextToken?: string
    /**
     * Whether or not to include the NodeResources and ReservedResources fields in the response.
     * @type boolean
     * @memberof NodesApigetNodes
     */
    resources?: boolean
}

export interface NodesApiUpdateNodeDrainRequest {
    /**
     * The ID of the node.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    nodeId: string
    /**
     * 
     * @type NodeUpdateDrainRequest
     * @memberof NodesApiupdateNodeDrain
     */
    nodeUpdateDrainRequest: NodeUpdateDrainRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApiupdateNodeDrain
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApiupdateNodeDrain
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApiupdateNodeDrain
     */
    nextToken?: string
}

export interface NodesApiUpdateNodeEligibilityRequest {
    /**
     * The ID of the node.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    nodeId: string
    /**
     * 
     * @type NodeUpdateEligibilityRequest
     * @memberof NodesApiupdateNodeEligibility
     */
    nodeUpdateEligibilityRequest: NodeUpdateEligibilityRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApiupdateNodeEligibility
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApiupdateNodeEligibility
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApiupdateNodeEligibility
     */
    nextToken?: string
}

export interface NodesApiUpdateNodePurgeRequest {
    /**
     * The ID of the node.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    nodeId: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof NodesApiupdateNodePurge
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof NodesApiupdateNodePurge
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof NodesApiupdateNodePurge
     */
    nextToken?: string
}

export class ObjectNodesApi {
    private api: ObservableNodesApi

    public constructor(configuration: Configuration, requestFactory?: NodesApiRequestFactory, responseProcessor?: NodesApiResponseProcessor) {
        this.api = new ObservableNodesApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getNode(param: NodesApiGetNodeRequest, options?: Configuration): Promise<Node> {
        return this.api.getNode(param.nodeId, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getNodeAllocations(param: NodesApiGetNodeAllocationsRequest, options?: Configuration): Promise<Array<AllocationListStub>> {
        return this.api.getNodeAllocations(param.nodeId, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getNodes(param: NodesApiGetNodesRequest = {}, options?: Configuration): Promise<Array<NodeListStub>> {
        return this.api.getNodes(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.resources,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public updateNodeDrain(param: NodesApiUpdateNodeDrainRequest, options?: Configuration): Promise<NodeDrainUpdateResponse> {
        return this.api.updateNodeDrain(param.nodeId, param.nodeUpdateDrainRequest, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public updateNodeEligibility(param: NodesApiUpdateNodeEligibilityRequest, options?: Configuration): Promise<NodeEligibilityUpdateResponse> {
        return this.api.updateNodeEligibility(param.nodeId, param.nodeUpdateEligibilityRequest, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public updateNodePurge(param: NodesApiUpdateNodePurgeRequest, options?: Configuration): Promise<NodePurgeResponse> {
        return this.api.updateNodePurge(param.nodeId, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableOperatorApi } from "./ObservableAPI";
import { OperatorApiRequestFactory, OperatorApiResponseProcessor} from "../apis/OperatorApi";

export interface OperatorApiDeleteOperatorRaftPeerRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApideleteOperatorRaftPeer
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApideleteOperatorRaftPeer
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApideleteOperatorRaftPeer
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof OperatorApideleteOperatorRaftPeer
     */
    idempotencyToken?: string
}

export interface OperatorApiGetOperatorAutopilotConfigurationRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotConfiguration
     */
    nextToken?: string
}

export interface OperatorApiGetOperatorAutopilotHealthRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof OperatorApigetOperatorAutopilotHealth
     */
    nextToken?: string
}

export interface OperatorApiGetOperatorRaftConfigurationRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof OperatorApigetOperatorRaftConfiguration
     */
    nextToken?: string
}

export interface OperatorApiGetOperatorSchedulerConfigurationRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof OperatorApigetOperatorSchedulerConfiguration
     */
    nextToken?: string
}

export interface OperatorApiPostOperatorSchedulerConfigurationRequest {
    /**
     * 
     * @type SchedulerConfiguration
     * @memberof OperatorApipostOperatorSchedulerConfiguration
     */
    schedulerConfiguration: SchedulerConfiguration
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApipostOperatorSchedulerConfiguration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApipostOperatorSchedulerConfiguration
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApipostOperatorSchedulerConfiguration
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof OperatorApipostOperatorSchedulerConfiguration
     */
    idempotencyToken?: string
}

export interface OperatorApiPutOperatorAutopilotConfigurationRequest {
    /**
     * 
     * @type AutopilotConfiguration
     * @memberof OperatorApiputOperatorAutopilotConfiguration
     */
    autopilotConfiguration: AutopilotConfiguration
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof OperatorApiputOperatorAutopilotConfiguration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof OperatorApiputOperatorAutopilotConfiguration
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof OperatorApiputOperatorAutopilotConfiguration
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof OperatorApiputOperatorAutopilotConfiguration
     */
    idempotencyToken?: string
}

export class ObjectOperatorApi {
    private api: ObservableOperatorApi

    public constructor(configuration: Configuration, requestFactory?: OperatorApiRequestFactory, responseProcessor?: OperatorApiResponseProcessor) {
        this.api = new ObservableOperatorApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public deleteOperatorRaftPeer(param: OperatorApiDeleteOperatorRaftPeerRequest = {}, options?: Configuration): Promise<void> {
        return this.api.deleteOperatorRaftPeer(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getOperatorAutopilotConfiguration(param: OperatorApiGetOperatorAutopilotConfigurationRequest = {}, options?: Configuration): Promise<AutopilotConfiguration> {
        return this.api.getOperatorAutopilotConfiguration(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getOperatorAutopilotHealth(param: OperatorApiGetOperatorAutopilotHealthRequest = {}, options?: Configuration): Promise<OperatorHealthReply> {
        return this.api.getOperatorAutopilotHealth(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getOperatorRaftConfiguration(param: OperatorApiGetOperatorRaftConfigurationRequest = {}, options?: Configuration): Promise<RaftConfiguration> {
        return this.api.getOperatorRaftConfiguration(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getOperatorSchedulerConfiguration(param: OperatorApiGetOperatorSchedulerConfigurationRequest = {}, options?: Configuration): Promise<SchedulerConfigurationResponse> {
        return this.api.getOperatorSchedulerConfiguration(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postOperatorSchedulerConfiguration(param: OperatorApiPostOperatorSchedulerConfigurationRequest, options?: Configuration): Promise<SchedulerSetConfigurationResponse> {
        return this.api.postOperatorSchedulerConfiguration(param.schedulerConfiguration, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public putOperatorAutopilotConfiguration(param: OperatorApiPutOperatorAutopilotConfigurationRequest, options?: Configuration): Promise<boolean> {
        return this.api.putOperatorAutopilotConfiguration(param.autopilotConfiguration, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservablePluginsApi } from "./ObservableAPI";
import { PluginsApiRequestFactory, PluginsApiResponseProcessor} from "../apis/PluginsApi";

export interface PluginsApiGetPluginCSIRequest {
    /**
     * The CSI plugin identifier.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    pluginID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof PluginsApigetPluginCSI
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof PluginsApigetPluginCSI
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof PluginsApigetPluginCSI
     */
    nextToken?: string
}

export interface PluginsApiGetPluginsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof PluginsApigetPlugins
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof PluginsApigetPlugins
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof PluginsApigetPlugins
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof PluginsApigetPlugins
     */
    nextToken?: string
}

export class ObjectPluginsApi {
    private api: ObservablePluginsApi

    public constructor(configuration: Configuration, requestFactory?: PluginsApiRequestFactory, responseProcessor?: PluginsApiResponseProcessor) {
        this.api = new ObservablePluginsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getPluginCSI(param: PluginsApiGetPluginCSIRequest, options?: Configuration): Promise<Array<CSIPlugin>> {
        return this.api.getPluginCSI(param.pluginID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getPlugins(param: PluginsApiGetPluginsRequest = {}, options?: Configuration): Promise<Array<CSIPluginListStub>> {
        return this.api.getPlugins(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableRegionsApi } from "./ObservableAPI";
import { RegionsApiRequestFactory, RegionsApiResponseProcessor} from "../apis/RegionsApi";

export interface RegionsApiGetRegionsRequest {
}

export class ObjectRegionsApi {
    private api: ObservableRegionsApi

    public constructor(configuration: Configuration, requestFactory?: RegionsApiRequestFactory, responseProcessor?: RegionsApiResponseProcessor) {
        this.api = new ObservableRegionsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getRegions(param: RegionsApiGetRegionsRequest = {}, options?: Configuration): Promise<Array<string>> {
        return this.api.getRegions( options).toPromise();
    }

}

import { ObservableScalingApi } from "./ObservableAPI";
import { ScalingApiRequestFactory, ScalingApiResponseProcessor} from "../apis/ScalingApi";

export interface ScalingApiGetScalingPoliciesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ScalingApigetScalingPolicies
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ScalingApigetScalingPolicies
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ScalingApigetScalingPolicies
     */
    nextToken?: string
}

export interface ScalingApiGetScalingPolicyRequest {
    /**
     * The scaling policy identifier.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    policyID: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof ScalingApigetScalingPolicy
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof ScalingApigetScalingPolicy
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof ScalingApigetScalingPolicy
     */
    nextToken?: string
}

export class ObjectScalingApi {
    private api: ObservableScalingApi

    public constructor(configuration: Configuration, requestFactory?: ScalingApiRequestFactory, responseProcessor?: ScalingApiResponseProcessor) {
        this.api = new ObservableScalingApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getScalingPolicies(param: ScalingApiGetScalingPoliciesRequest = {}, options?: Configuration): Promise<Array<ScalingPolicyListStub>> {
        return this.api.getScalingPolicies(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getScalingPolicy(param: ScalingApiGetScalingPolicyRequest, options?: Configuration): Promise<ScalingPolicy> {
        return this.api.getScalingPolicy(param.policyID, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableSearchApi } from "./ObservableAPI";
import { SearchApiRequestFactory, SearchApiResponseProcessor} from "../apis/SearchApi";

export interface SearchApiGetFuzzySearchRequest {
    /**
     * 
     * @type FuzzySearchRequest
     * @memberof SearchApigetFuzzySearch
     */
    fuzzySearchRequest: FuzzySearchRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof SearchApigetFuzzySearch
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof SearchApigetFuzzySearch
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof SearchApigetFuzzySearch
     */
    nextToken?: string
}

export interface SearchApiGetSearchRequest {
    /**
     * 
     * @type SearchRequest
     * @memberof SearchApigetSearch
     */
    searchRequest: SearchRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof SearchApigetSearch
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof SearchApigetSearch
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof SearchApigetSearch
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof SearchApigetSearch
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof SearchApigetSearch
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof SearchApigetSearch
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof SearchApigetSearch
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof SearchApigetSearch
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof SearchApigetSearch
     */
    nextToken?: string
}

export class ObjectSearchApi {
    private api: ObservableSearchApi

    public constructor(configuration: Configuration, requestFactory?: SearchApiRequestFactory, responseProcessor?: SearchApiResponseProcessor) {
        this.api = new ObservableSearchApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getFuzzySearch(param: SearchApiGetFuzzySearchRequest, options?: Configuration): Promise<FuzzySearchResponse> {
        return this.api.getFuzzySearch(param.fuzzySearchRequest, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getSearch(param: SearchApiGetSearchRequest, options?: Configuration): Promise<SearchResponse> {
        return this.api.getSearch(param.searchRequest, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableStatusApi } from "./ObservableAPI";
import { StatusApiRequestFactory, StatusApiResponseProcessor} from "../apis/StatusApi";

export interface StatusApiGetStatusLeaderRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof StatusApigetStatusLeader
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof StatusApigetStatusLeader
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof StatusApigetStatusLeader
     */
    nextToken?: string
}

export interface StatusApiGetStatusPeersRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof StatusApigetStatusPeers
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof StatusApigetStatusPeers
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof StatusApigetStatusPeers
     */
    nextToken?: string
}

export class ObjectStatusApi {
    private api: ObservableStatusApi

    public constructor(configuration: Configuration, requestFactory?: StatusApiRequestFactory, responseProcessor?: StatusApiResponseProcessor) {
        this.api = new ObservableStatusApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public getStatusLeader(param: StatusApiGetStatusLeaderRequest = {}, options?: Configuration): Promise<string> {
        return this.api.getStatusLeader(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getStatusPeers(param: StatusApiGetStatusPeersRequest = {}, options?: Configuration): Promise<Array<string>> {
        return this.api.getStatusPeers(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

}

import { ObservableSystemApi } from "./ObservableAPI";
import { SystemApiRequestFactory, SystemApiResponseProcessor} from "../apis/SystemApi";

export interface SystemApiPutSystemGCRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof SystemApiputSystemGC
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof SystemApiputSystemGC
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof SystemApiputSystemGC
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof SystemApiputSystemGC
     */
    idempotencyToken?: string
}

export interface SystemApiPutSystemReconcileSummariesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof SystemApiputSystemReconcileSummaries
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof SystemApiputSystemReconcileSummaries
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof SystemApiputSystemReconcileSummaries
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof SystemApiputSystemReconcileSummaries
     */
    idempotencyToken?: string
}

export class ObjectSystemApi {
    private api: ObservableSystemApi

    public constructor(configuration: Configuration, requestFactory?: SystemApiRequestFactory, responseProcessor?: SystemApiResponseProcessor) {
        this.api = new ObservableSystemApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public putSystemGC(param: SystemApiPutSystemGCRequest = {}, options?: Configuration): Promise<void> {
        return this.api.putSystemGC(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public putSystemReconcileSummaries(param: SystemApiPutSystemReconcileSummariesRequest = {}, options?: Configuration): Promise<void> {
        return this.api.putSystemReconcileSummaries(param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}

import { ObservableVolumesApi } from "./ObservableAPI";
import { VolumesApiRequestFactory, VolumesApiResponseProcessor} from "../apis/VolumesApi";

export interface VolumesApiCreateVolumeRequest {
    /**
     * Volume unique identifier.
     * @type string
     * @memberof VolumesApicreateVolume
     */
    volumeId: string
    /**
     * The action to perform on the Volume (create, detach, delete).
     * @type string
     * @memberof VolumesApicreateVolume
     */
    action: string
    /**
     * 
     * @type CSIVolumeCreateRequest
     * @memberof VolumesApicreateVolume
     */
    cSIVolumeCreateRequest: CSIVolumeCreateRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApicreateVolume
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApicreateVolume
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApicreateVolume
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApicreateVolume
     */
    idempotencyToken?: string
}

export interface VolumesApiDeleteSnapshotRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    idempotencyToken?: string
    /**
     * Filters volume lists by plugin ID.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    pluginId?: string
    /**
     * The ID of the snapshot to target.
     * @type string
     * @memberof VolumesApideleteSnapshot
     */
    snapshotId?: string
}

export interface VolumesApiDeleteVolumeRegistrationRequest {
    /**
     * Volume unique identifier.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    volumeId: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    idempotencyToken?: string
    /**
     * Used to force the de-registration of a volume.
     * @type string
     * @memberof VolumesApideleteVolumeRegistration
     */
    force?: string
}

export interface VolumesApiDetachOrDeleteVolumeRequest {
    /**
     * Volume unique identifier.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    volumeId: string
    /**
     * The action to perform on the Volume (create, detach, delete).
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    action: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    idempotencyToken?: string
    /**
     * Specifies node to target volume operation for.
     * @type string
     * @memberof VolumesApidetachOrDeleteVolume
     */
    node?: string
}

export interface VolumesApiGetExternalVolumesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof VolumesApigetExternalVolumes
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof VolumesApigetExternalVolumes
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    nextToken?: string
    /**
     * Filters volume lists by plugin ID.
     * @type string
     * @memberof VolumesApigetExternalVolumes
     */
    pluginId?: string
}

export interface VolumesApiGetSnapshotsRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof VolumesApigetSnapshots
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof VolumesApigetSnapshots
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    nextToken?: string
    /**
     * Filters volume lists by plugin ID.
     * @type string
     * @memberof VolumesApigetSnapshots
     */
    pluginId?: string
}

export interface VolumesApiGetVolumeRequest {
    /**
     * Volume unique identifier.
     * @type string
     * @memberof VolumesApigetVolume
     */
    volumeId: string
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApigetVolume
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApigetVolume
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof VolumesApigetVolume
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof VolumesApigetVolume
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof VolumesApigetVolume
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof VolumesApigetVolume
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApigetVolume
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof VolumesApigetVolume
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof VolumesApigetVolume
     */
    nextToken?: string
}

export interface VolumesApiGetVolumesRequest {
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    namespace?: string
    /**
     * If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @type number
     * @memberof VolumesApigetVolumes
     */
    index?: number
    /**
     * Provided with IndexParam to wait for change.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    wait?: string
    /**
     * If present, results will include stale reads.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    stale?: string
    /**
     * Constrains results to jobs that start with the defined prefix
     * @type string
     * @memberof VolumesApigetVolumes
     */
    prefix?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    xNomadToken?: string
    /**
     * Maximum number of results to return.
     * @type number
     * @memberof VolumesApigetVolumes
     */
    perPage?: number
    /**
     * Indicates where to start paging for queries that support pagination.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    nextToken?: string
    /**
     * Filters volume lists by node ID.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    nodeId?: string
    /**
     * Filters volume lists by plugin ID.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    pluginId?: string
    /**
     * Filters volume lists to a specific type.
     * @type string
     * @memberof VolumesApigetVolumes
     */
    type?: string
}

export interface VolumesApiPostSnapshotRequest {
    /**
     * 
     * @type CSISnapshotCreateRequest
     * @memberof VolumesApipostSnapshot
     */
    cSISnapshotCreateRequest: CSISnapshotCreateRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApipostSnapshot
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApipostSnapshot
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApipostSnapshot
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApipostSnapshot
     */
    idempotencyToken?: string
}

export interface VolumesApiPostVolumeRequest {
    /**
     * 
     * @type CSIVolumeRegisterRequest
     * @memberof VolumesApipostVolume
     */
    cSIVolumeRegisterRequest: CSIVolumeRegisterRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApipostVolume
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApipostVolume
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApipostVolume
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApipostVolume
     */
    idempotencyToken?: string
}

export interface VolumesApiPostVolumeRegistrationRequest {
    /**
     * Volume unique identifier.
     * @type string
     * @memberof VolumesApipostVolumeRegistration
     */
    volumeId: string
    /**
     * 
     * @type CSIVolumeRegisterRequest
     * @memberof VolumesApipostVolumeRegistration
     */
    cSIVolumeRegisterRequest: CSIVolumeRegisterRequest
    /**
     * Filters results based on the specified region.
     * @type string
     * @memberof VolumesApipostVolumeRegistration
     */
    region?: string
    /**
     * Filters results based on the specified namespace.
     * @type string
     * @memberof VolumesApipostVolumeRegistration
     */
    namespace?: string
    /**
     * A Nomad ACL token.
     * @type string
     * @memberof VolumesApipostVolumeRegistration
     */
    xNomadToken?: string
    /**
     * Can be used to ensure operations are only run once.
     * @type string
     * @memberof VolumesApipostVolumeRegistration
     */
    idempotencyToken?: string
}

export class ObjectVolumesApi {
    private api: ObservableVolumesApi

    public constructor(configuration: Configuration, requestFactory?: VolumesApiRequestFactory, responseProcessor?: VolumesApiResponseProcessor) {
        this.api = new ObservableVolumesApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public createVolume(param: VolumesApiCreateVolumeRequest, options?: Configuration): Promise<void> {
        return this.api.createVolume(param.volumeId, param.action, param.cSIVolumeCreateRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public deleteSnapshot(param: VolumesApiDeleteSnapshotRequest = {}, options?: Configuration): Promise<void> {
        return this.api.deleteSnapshot(param.region, param.namespace, param.xNomadToken, param.idempotencyToken, param.pluginId, param.snapshotId,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public deleteVolumeRegistration(param: VolumesApiDeleteVolumeRegistrationRequest, options?: Configuration): Promise<void> {
        return this.api.deleteVolumeRegistration(param.volumeId, param.region, param.namespace, param.xNomadToken, param.idempotencyToken, param.force,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public detachOrDeleteVolume(param: VolumesApiDetachOrDeleteVolumeRequest, options?: Configuration): Promise<void> {
        return this.api.detachOrDeleteVolume(param.volumeId, param.action, param.region, param.namespace, param.xNomadToken, param.idempotencyToken, param.node,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getExternalVolumes(param: VolumesApiGetExternalVolumesRequest = {}, options?: Configuration): Promise<CSIVolumeListExternalResponse> {
        return this.api.getExternalVolumes(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.pluginId,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getSnapshots(param: VolumesApiGetSnapshotsRequest = {}, options?: Configuration): Promise<CSISnapshotListResponse> {
        return this.api.getSnapshots(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.pluginId,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getVolume(param: VolumesApiGetVolumeRequest, options?: Configuration): Promise<CSIVolume> {
        return this.api.getVolume(param.volumeId, param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public getVolumes(param: VolumesApiGetVolumesRequest = {}, options?: Configuration): Promise<Array<CSIVolumeListStub>> {
        return this.api.getVolumes(param.region, param.namespace, param.index, param.wait, param.stale, param.prefix, param.xNomadToken, param.perPage, param.nextToken, param.nodeId, param.pluginId, param.type,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postSnapshot(param: VolumesApiPostSnapshotRequest, options?: Configuration): Promise<CSISnapshotCreateResponse> {
        return this.api.postSnapshot(param.cSISnapshotCreateRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postVolume(param: VolumesApiPostVolumeRequest, options?: Configuration): Promise<void> {
        return this.api.postVolume(param.cSIVolumeRegisterRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public postVolumeRegistration(param: VolumesApiPostVolumeRegistrationRequest, options?: Configuration): Promise<void> {
        return this.api.postVolumeRegistration(param.volumeId, param.cSIVolumeRegisterRequest, param.region, param.namespace, param.xNomadToken, param.idempotencyToken,  options).toPromise();
    }

}
