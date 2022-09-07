import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'
import { Observable, of, from } from '../rxjsStub';
import {mergeMap, map} from  '../rxjsStub';
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

import { ACLApiRequestFactory, ACLApiResponseProcessor} from "../apis/ACLApi";
export class ObservableACLApi {
    private requestFactory: ACLApiRequestFactory;
    private responseProcessor: ACLApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: ACLApiRequestFactory,
        responseProcessor?: ACLApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new ACLApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new ACLApiResponseProcessor();
    }

    /**
     * @param policyName The ACL policy name.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public deleteACLPolicy(policyName: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteACLPolicy(policyName, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteACLPolicy(rsp)));
            }));
    }

    /**
     * @param tokenAccessor The token accessor ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public deleteACLToken(tokenAccessor: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteACLToken(tokenAccessor, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteACLToken(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getACLPolicies(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<ACLPolicyListStub>> {
        const requestContextPromise = this.requestFactory.getACLPolicies(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getACLPolicies(rsp)));
            }));
    }

    /**
     * @param policyName The ACL policy name.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getACLPolicy(policyName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<ACLPolicy> {
        const requestContextPromise = this.requestFactory.getACLPolicy(policyName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getACLPolicy(rsp)));
            }));
    }

    /**
     * @param tokenAccessor The token accessor ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getACLToken(tokenAccessor: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<ACLToken> {
        const requestContextPromise = this.requestFactory.getACLToken(tokenAccessor, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getACLToken(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getACLTokenSelf(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<ACLToken> {
        const requestContextPromise = this.requestFactory.getACLTokenSelf(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getACLTokenSelf(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getACLTokens(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<ACLTokenListStub>> {
        const requestContextPromise = this.requestFactory.getACLTokens(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getACLTokens(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postACLBootstrap(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<ACLToken> {
        const requestContextPromise = this.requestFactory.postACLBootstrap(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postACLBootstrap(rsp)));
            }));
    }

    /**
     * @param policyName The ACL policy name.
     * @param aCLPolicy 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postACLPolicy(policyName: string, aCLPolicy: ACLPolicy, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.postACLPolicy(policyName, aCLPolicy, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postACLPolicy(rsp)));
            }));
    }

    /**
     * @param tokenAccessor The token accessor ID.
     * @param aCLToken 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postACLToken(tokenAccessor: string, aCLToken: ACLToken, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<ACLToken> {
        const requestContextPromise = this.requestFactory.postACLToken(tokenAccessor, aCLToken, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postACLToken(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postACLTokenOnetime(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<OneTimeToken> {
        const requestContextPromise = this.requestFactory.postACLTokenOnetime(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postACLTokenOnetime(rsp)));
            }));
    }

    /**
     * @param oneTimeTokenExchangeRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postACLTokenOnetimeExchange(oneTimeTokenExchangeRequest: OneTimeTokenExchangeRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<ACLToken> {
        const requestContextPromise = this.requestFactory.postACLTokenOnetimeExchange(oneTimeTokenExchangeRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postACLTokenOnetimeExchange(rsp)));
            }));
    }

}

import { AllocationsApiRequestFactory, AllocationsApiResponseProcessor} from "../apis/AllocationsApi";
export class ObservableAllocationsApi {
    private requestFactory: AllocationsApiRequestFactory;
    private responseProcessor: AllocationsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: AllocationsApiRequestFactory,
        responseProcessor?: AllocationsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new AllocationsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new AllocationsApiResponseProcessor();
    }

    /**
     * @param allocID Allocation ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getAllocation(allocID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Allocation> {
        const requestContextPromise = this.requestFactory.getAllocation(allocID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getAllocation(rsp)));
            }));
    }

    /**
     * @param allocID Allocation ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getAllocationServices(allocID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<ServiceRegistration>> {
        const requestContextPromise = this.requestFactory.getAllocationServices(allocID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getAllocationServices(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param resources Flag indicating whether to include resources in response.
     * @param taskStates Flag indicating whether to include task states in response.
     */
    public getAllocations(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, resources?: boolean, taskStates?: boolean, _options?: Configuration): Observable<Array<AllocationListStub>> {
        const requestContextPromise = this.requestFactory.getAllocations(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, resources, taskStates, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getAllocations(rsp)));
            }));
    }

    /**
     * @param allocID Allocation ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param noShutdownDelay Flag indicating whether to delay shutdown when requesting an allocation stop.
     */
    public postAllocationStop(allocID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, noShutdownDelay?: boolean, _options?: Configuration): Observable<AllocStopResponse> {
        const requestContextPromise = this.requestFactory.postAllocationStop(allocID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, noShutdownDelay, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postAllocationStop(rsp)));
            }));
    }

}

import { DeploymentsApiRequestFactory, DeploymentsApiResponseProcessor} from "../apis/DeploymentsApi";
export class ObservableDeploymentsApi {
    private requestFactory: DeploymentsApiRequestFactory;
    private responseProcessor: DeploymentsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: DeploymentsApiRequestFactory,
        responseProcessor?: DeploymentsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new DeploymentsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new DeploymentsApiResponseProcessor();
    }

    /**
     * @param deploymentID Deployment ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getDeployment(deploymentID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Deployment> {
        const requestContextPromise = this.requestFactory.getDeployment(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getDeployment(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getDeploymentAllocations(deploymentID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<AllocationListStub>> {
        const requestContextPromise = this.requestFactory.getDeploymentAllocations(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getDeploymentAllocations(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getDeployments(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<Deployment>> {
        const requestContextPromise = this.requestFactory.getDeployments(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getDeployments(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param deploymentAllocHealthRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postDeploymentAllocationHealth(deploymentID: string, deploymentAllocHealthRequest: DeploymentAllocHealthRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<DeploymentUpdateResponse> {
        const requestContextPromise = this.requestFactory.postDeploymentAllocationHealth(deploymentID, deploymentAllocHealthRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postDeploymentAllocationHealth(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postDeploymentFail(deploymentID: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<DeploymentUpdateResponse> {
        const requestContextPromise = this.requestFactory.postDeploymentFail(deploymentID, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postDeploymentFail(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param deploymentPauseRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postDeploymentPause(deploymentID: string, deploymentPauseRequest: DeploymentPauseRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<DeploymentUpdateResponse> {
        const requestContextPromise = this.requestFactory.postDeploymentPause(deploymentID, deploymentPauseRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postDeploymentPause(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param deploymentPromoteRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postDeploymentPromote(deploymentID: string, deploymentPromoteRequest: DeploymentPromoteRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<DeploymentUpdateResponse> {
        const requestContextPromise = this.requestFactory.postDeploymentPromote(deploymentID, deploymentPromoteRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postDeploymentPromote(rsp)));
            }));
    }

    /**
     * @param deploymentID Deployment ID.
     * @param deploymentUnblockRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postDeploymentUnblock(deploymentID: string, deploymentUnblockRequest: DeploymentUnblockRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<DeploymentUpdateResponse> {
        const requestContextPromise = this.requestFactory.postDeploymentUnblock(deploymentID, deploymentUnblockRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postDeploymentUnblock(rsp)));
            }));
    }

}

import { EnterpriseApiRequestFactory, EnterpriseApiResponseProcessor} from "../apis/EnterpriseApi";
export class ObservableEnterpriseApi {
    private requestFactory: EnterpriseApiRequestFactory;
    private responseProcessor: EnterpriseApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: EnterpriseApiRequestFactory,
        responseProcessor?: EnterpriseApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new EnterpriseApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new EnterpriseApiResponseProcessor();
    }

    /**
     * @param quotaSpec 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public createQuotaSpec(quotaSpec: QuotaSpec, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.createQuotaSpec(quotaSpec, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.createQuotaSpec(rsp)));
            }));
    }

    /**
     * @param specName The quota spec identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public deleteQuotaSpec(specName: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteQuotaSpec(specName, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteQuotaSpec(rsp)));
            }));
    }

    /**
     * @param specName The quota spec identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getQuotaSpec(specName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<QuotaSpec> {
        const requestContextPromise = this.requestFactory.getQuotaSpec(specName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getQuotaSpec(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getQuotas(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<any>> {
        const requestContextPromise = this.requestFactory.getQuotas(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getQuotas(rsp)));
            }));
    }

    /**
     * @param specName The quota spec identifier.
     * @param quotaSpec 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postQuotaSpec(specName: string, quotaSpec: QuotaSpec, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.postQuotaSpec(specName, quotaSpec, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postQuotaSpec(rsp)));
            }));
    }

}

import { EvaluationsApiRequestFactory, EvaluationsApiResponseProcessor} from "../apis/EvaluationsApi";
export class ObservableEvaluationsApi {
    private requestFactory: EvaluationsApiRequestFactory;
    private responseProcessor: EvaluationsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: EvaluationsApiRequestFactory,
        responseProcessor?: EvaluationsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new EvaluationsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new EvaluationsApiResponseProcessor();
    }

    /**
     * @param evalID Evaluation ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getEvaluation(evalID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Evaluation> {
        const requestContextPromise = this.requestFactory.getEvaluation(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getEvaluation(rsp)));
            }));
    }

    /**
     * @param evalID Evaluation ID.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getEvaluationAllocations(evalID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<AllocationListStub>> {
        const requestContextPromise = this.requestFactory.getEvaluationAllocations(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getEvaluationAllocations(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getEvaluations(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<Evaluation>> {
        const requestContextPromise = this.requestFactory.getEvaluations(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getEvaluations(rsp)));
            }));
    }

}

import { JobsApiRequestFactory, JobsApiResponseProcessor} from "../apis/JobsApi";
export class ObservableJobsApi {
    private requestFactory: JobsApiRequestFactory;
    private responseProcessor: JobsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: JobsApiRequestFactory,
        responseProcessor?: JobsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new JobsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new JobsApiResponseProcessor();
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     * @param purge Boolean flag indicating whether to purge allocations of the job after deleting.
     * @param global Boolean flag indicating whether the operation should apply to all instances of the job globally.
     */
    public deleteJob(jobName: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, purge?: boolean, global?: boolean, _options?: Configuration): Observable<JobDeregisterResponse> {
        const requestContextPromise = this.requestFactory.deleteJob(jobName, region, namespace, xNomadToken, idempotencyToken, purge, global, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteJob(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJob(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Job> {
        const requestContextPromise = this.requestFactory.getJob(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJob(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param all Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered.
     */
    public getJobAllocations(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, all?: boolean, _options?: Configuration): Observable<Array<AllocationListStub>> {
        const requestContextPromise = this.requestFactory.getJobAllocations(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, all, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobAllocations(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJobDeployment(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Deployment> {
        const requestContextPromise = this.requestFactory.getJobDeployment(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobDeployment(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param all Flag indicating whether to constrain by job creation index or not.
     */
    public getJobDeployments(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, all?: number, _options?: Configuration): Observable<Array<Deployment>> {
        const requestContextPromise = this.requestFactory.getJobDeployments(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, all, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobDeployments(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJobEvaluations(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<Evaluation>> {
        const requestContextPromise = this.requestFactory.getJobEvaluations(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobEvaluations(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJobScaleStatus(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<JobScaleStatusResponse> {
        const requestContextPromise = this.requestFactory.getJobScaleStatus(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobScaleStatus(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJobSummary(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<JobSummary> {
        const requestContextPromise = this.requestFactory.getJobSummary(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobSummary(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param diffs Boolean flag indicating whether to compute job diffs.
     */
    public getJobVersions(jobName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, diffs?: boolean, _options?: Configuration): Observable<JobVersionsResponse> {
        const requestContextPromise = this.requestFactory.getJobVersions(jobName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, diffs, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobVersions(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getJobs(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<JobListStub>> {
        const requestContextPromise = this.requestFactory.getJobs(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getJobs(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobRegisterRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJob(jobName: string, jobRegisterRequest: JobRegisterRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobRegisterResponse> {
        const requestContextPromise = this.requestFactory.postJob(jobName, jobRegisterRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJob(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobDispatchRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobDispatch(jobName: string, jobDispatchRequest: JobDispatchRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobDispatchResponse> {
        const requestContextPromise = this.requestFactory.postJobDispatch(jobName, jobDispatchRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobDispatch(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobEvaluateRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobEvaluate(jobName: string, jobEvaluateRequest: JobEvaluateRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobRegisterResponse> {
        const requestContextPromise = this.requestFactory.postJobEvaluate(jobName, jobEvaluateRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobEvaluate(rsp)));
            }));
    }

    /**
     * @param jobsParseRequest 
     */
    public postJobParse(jobsParseRequest: JobsParseRequest, _options?: Configuration): Observable<Job> {
        const requestContextPromise = this.requestFactory.postJobParse(jobsParseRequest, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobParse(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobPeriodicForce(jobName: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<PeriodicForceResponse> {
        const requestContextPromise = this.requestFactory.postJobPeriodicForce(jobName, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobPeriodicForce(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobPlanRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobPlan(jobName: string, jobPlanRequest: JobPlanRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobPlanResponse> {
        const requestContextPromise = this.requestFactory.postJobPlan(jobName, jobPlanRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobPlan(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobRevertRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobRevert(jobName: string, jobRevertRequest: JobRevertRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobRegisterResponse> {
        const requestContextPromise = this.requestFactory.postJobRevert(jobName, jobRevertRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobRevert(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param scalingRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobScalingRequest(jobName: string, scalingRequest: ScalingRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobRegisterResponse> {
        const requestContextPromise = this.requestFactory.postJobScalingRequest(jobName, scalingRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobScalingRequest(rsp)));
            }));
    }

    /**
     * @param jobName The job identifier.
     * @param jobStabilityRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobStability(jobName: string, jobStabilityRequest: JobStabilityRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobStabilityResponse> {
        const requestContextPromise = this.requestFactory.postJobStability(jobName, jobStabilityRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobStability(rsp)));
            }));
    }

    /**
     * @param jobValidateRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postJobValidateRequest(jobValidateRequest: JobValidateRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobValidateResponse> {
        const requestContextPromise = this.requestFactory.postJobValidateRequest(jobValidateRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postJobValidateRequest(rsp)));
            }));
    }

    /**
     * @param jobRegisterRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public registerJob(jobRegisterRequest: JobRegisterRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<JobRegisterResponse> {
        const requestContextPromise = this.requestFactory.registerJob(jobRegisterRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.registerJob(rsp)));
            }));
    }

}

import { MetricsApiRequestFactory, MetricsApiResponseProcessor} from "../apis/MetricsApi";
export class ObservableMetricsApi {
    private requestFactory: MetricsApiRequestFactory;
    private responseProcessor: MetricsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: MetricsApiRequestFactory,
        responseProcessor?: MetricsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new MetricsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new MetricsApiResponseProcessor();
    }

    /**
     * @param format The format the user requested for the metrics summary (e.g. prometheus)
     */
    public getMetricsSummary(format?: string, _options?: Configuration): Observable<MetricsSummary> {
        const requestContextPromise = this.requestFactory.getMetricsSummary(format, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getMetricsSummary(rsp)));
            }));
    }

}

import { NamespacesApiRequestFactory, NamespacesApiResponseProcessor} from "../apis/NamespacesApi";
export class ObservableNamespacesApi {
    private requestFactory: NamespacesApiRequestFactory;
    private responseProcessor: NamespacesApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: NamespacesApiRequestFactory,
        responseProcessor?: NamespacesApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new NamespacesApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new NamespacesApiResponseProcessor();
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public createNamespace(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.createNamespace(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.createNamespace(rsp)));
            }));
    }

    /**
     * @param namespaceName The namespace identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public deleteNamespace(namespaceName: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteNamespace(namespaceName, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteNamespace(rsp)));
            }));
    }

    /**
     * @param namespaceName The namespace identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getNamespace(namespaceName: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Namespace> {
        const requestContextPromise = this.requestFactory.getNamespace(namespaceName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getNamespace(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getNamespaces(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<Namespace>> {
        const requestContextPromise = this.requestFactory.getNamespaces(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getNamespaces(rsp)));
            }));
    }

    /**
     * @param namespaceName The namespace identifier.
     * @param namespace2 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postNamespace(namespaceName: string, namespace2: Namespace, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.postNamespace(namespaceName, namespace2, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postNamespace(rsp)));
            }));
    }

}

import { NodesApiRequestFactory, NodesApiResponseProcessor} from "../apis/NodesApi";
export class ObservableNodesApi {
    private requestFactory: NodesApiRequestFactory;
    private responseProcessor: NodesApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: NodesApiRequestFactory,
        responseProcessor?: NodesApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new NodesApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new NodesApiResponseProcessor();
    }

    /**
     * @param nodeId The ID of the node.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getNode(nodeId: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Node> {
        const requestContextPromise = this.requestFactory.getNode(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getNode(rsp)));
            }));
    }

    /**
     * @param nodeId The ID of the node.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getNodeAllocations(nodeId: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<AllocationListStub>> {
        const requestContextPromise = this.requestFactory.getNodeAllocations(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getNodeAllocations(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param resources Whether or not to include the NodeResources and ReservedResources fields in the response.
     */
    public getNodes(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, resources?: boolean, _options?: Configuration): Observable<Array<NodeListStub>> {
        const requestContextPromise = this.requestFactory.getNodes(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, resources, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getNodes(rsp)));
            }));
    }

    /**
     * @param nodeId The ID of the node.
     * @param nodeUpdateDrainRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public updateNodeDrain(nodeId: string, nodeUpdateDrainRequest: NodeUpdateDrainRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<NodeDrainUpdateResponse> {
        const requestContextPromise = this.requestFactory.updateNodeDrain(nodeId, nodeUpdateDrainRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.updateNodeDrain(rsp)));
            }));
    }

    /**
     * @param nodeId The ID of the node.
     * @param nodeUpdateEligibilityRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public updateNodeEligibility(nodeId: string, nodeUpdateEligibilityRequest: NodeUpdateEligibilityRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<NodeEligibilityUpdateResponse> {
        const requestContextPromise = this.requestFactory.updateNodeEligibility(nodeId, nodeUpdateEligibilityRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.updateNodeEligibility(rsp)));
            }));
    }

    /**
     * @param nodeId The ID of the node.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public updateNodePurge(nodeId: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<NodePurgeResponse> {
        const requestContextPromise = this.requestFactory.updateNodePurge(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.updateNodePurge(rsp)));
            }));
    }

}

import { OperatorApiRequestFactory, OperatorApiResponseProcessor} from "../apis/OperatorApi";
export class ObservableOperatorApi {
    private requestFactory: OperatorApiRequestFactory;
    private responseProcessor: OperatorApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: OperatorApiRequestFactory,
        responseProcessor?: OperatorApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new OperatorApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new OperatorApiResponseProcessor();
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public deleteOperatorRaftPeer(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteOperatorRaftPeer(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteOperatorRaftPeer(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getOperatorAutopilotConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<AutopilotConfiguration> {
        const requestContextPromise = this.requestFactory.getOperatorAutopilotConfiguration(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getOperatorAutopilotConfiguration(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getOperatorAutopilotHealth(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<OperatorHealthReply> {
        const requestContextPromise = this.requestFactory.getOperatorAutopilotHealth(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getOperatorAutopilotHealth(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getOperatorRaftConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<RaftConfiguration> {
        const requestContextPromise = this.requestFactory.getOperatorRaftConfiguration(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getOperatorRaftConfiguration(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getOperatorSchedulerConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<SchedulerConfigurationResponse> {
        const requestContextPromise = this.requestFactory.getOperatorSchedulerConfiguration(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getOperatorSchedulerConfiguration(rsp)));
            }));
    }

    /**
     * @param schedulerConfiguration 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postOperatorSchedulerConfiguration(schedulerConfiguration: SchedulerConfiguration, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<SchedulerSetConfigurationResponse> {
        const requestContextPromise = this.requestFactory.postOperatorSchedulerConfiguration(schedulerConfiguration, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postOperatorSchedulerConfiguration(rsp)));
            }));
    }

    /**
     * @param autopilotConfiguration 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public putOperatorAutopilotConfiguration(autopilotConfiguration: AutopilotConfiguration, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<boolean> {
        const requestContextPromise = this.requestFactory.putOperatorAutopilotConfiguration(autopilotConfiguration, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.putOperatorAutopilotConfiguration(rsp)));
            }));
    }

}

import { PluginsApiRequestFactory, PluginsApiResponseProcessor} from "../apis/PluginsApi";
export class ObservablePluginsApi {
    private requestFactory: PluginsApiRequestFactory;
    private responseProcessor: PluginsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: PluginsApiRequestFactory,
        responseProcessor?: PluginsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new PluginsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new PluginsApiResponseProcessor();
    }

    /**
     * @param pluginID The CSI plugin identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getPluginCSI(pluginID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<CSIPlugin>> {
        const requestContextPromise = this.requestFactory.getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getPluginCSI(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getPlugins(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<CSIPluginListStub>> {
        const requestContextPromise = this.requestFactory.getPlugins(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getPlugins(rsp)));
            }));
    }

}

import { RegionsApiRequestFactory, RegionsApiResponseProcessor} from "../apis/RegionsApi";
export class ObservableRegionsApi {
    private requestFactory: RegionsApiRequestFactory;
    private responseProcessor: RegionsApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: RegionsApiRequestFactory,
        responseProcessor?: RegionsApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new RegionsApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new RegionsApiResponseProcessor();
    }

    /**
     */
    public getRegions(_options?: Configuration): Observable<Array<string>> {
        const requestContextPromise = this.requestFactory.getRegions(_options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getRegions(rsp)));
            }));
    }

}

import { ScalingApiRequestFactory, ScalingApiResponseProcessor} from "../apis/ScalingApi";
export class ObservableScalingApi {
    private requestFactory: ScalingApiRequestFactory;
    private responseProcessor: ScalingApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: ScalingApiRequestFactory,
        responseProcessor?: ScalingApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new ScalingApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new ScalingApiResponseProcessor();
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getScalingPolicies(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<ScalingPolicyListStub>> {
        const requestContextPromise = this.requestFactory.getScalingPolicies(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getScalingPolicies(rsp)));
            }));
    }

    /**
     * @param policyID The scaling policy identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getScalingPolicy(policyID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<ScalingPolicy> {
        const requestContextPromise = this.requestFactory.getScalingPolicy(policyID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getScalingPolicy(rsp)));
            }));
    }

}

import { SearchApiRequestFactory, SearchApiResponseProcessor} from "../apis/SearchApi";
export class ObservableSearchApi {
    private requestFactory: SearchApiRequestFactory;
    private responseProcessor: SearchApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: SearchApiRequestFactory,
        responseProcessor?: SearchApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new SearchApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new SearchApiResponseProcessor();
    }

    /**
     * @param fuzzySearchRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getFuzzySearch(fuzzySearchRequest: FuzzySearchRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<FuzzySearchResponse> {
        const requestContextPromise = this.requestFactory.getFuzzySearch(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getFuzzySearch(rsp)));
            }));
    }

    /**
     * @param searchRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getSearch(searchRequest: SearchRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<SearchResponse> {
        const requestContextPromise = this.requestFactory.getSearch(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getSearch(rsp)));
            }));
    }

}

import { StatusApiRequestFactory, StatusApiResponseProcessor} from "../apis/StatusApi";
export class ObservableStatusApi {
    private requestFactory: StatusApiRequestFactory;
    private responseProcessor: StatusApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: StatusApiRequestFactory,
        responseProcessor?: StatusApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new StatusApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new StatusApiResponseProcessor();
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getStatusLeader(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<string> {
        const requestContextPromise = this.requestFactory.getStatusLeader(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getStatusLeader(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getStatusPeers(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<Array<string>> {
        const requestContextPromise = this.requestFactory.getStatusPeers(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getStatusPeers(rsp)));
            }));
    }

}

import { SystemApiRequestFactory, SystemApiResponseProcessor} from "../apis/SystemApi";
export class ObservableSystemApi {
    private requestFactory: SystemApiRequestFactory;
    private responseProcessor: SystemApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: SystemApiRequestFactory,
        responseProcessor?: SystemApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new SystemApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new SystemApiResponseProcessor();
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public putSystemGC(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.putSystemGC(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.putSystemGC(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public putSystemReconcileSummaries(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.putSystemReconcileSummaries(region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.putSystemReconcileSummaries(rsp)));
            }));
    }

}

import { VolumesApiRequestFactory, VolumesApiResponseProcessor} from "../apis/VolumesApi";
export class ObservableVolumesApi {
    private requestFactory: VolumesApiRequestFactory;
    private responseProcessor: VolumesApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: VolumesApiRequestFactory,
        responseProcessor?: VolumesApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new VolumesApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new VolumesApiResponseProcessor();
    }

    /**
     * @param volumeId Volume unique identifier.
     * @param action The action to perform on the Volume (create, detach, delete).
     * @param cSIVolumeCreateRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public createVolume(volumeId: string, action: string, cSIVolumeCreateRequest: CSIVolumeCreateRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.createVolume(volumeId, action, cSIVolumeCreateRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.createVolume(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     * @param pluginId Filters volume lists by plugin ID.
     * @param snapshotId The ID of the snapshot to target.
     */
    public deleteSnapshot(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, pluginId?: string, snapshotId?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteSnapshot(region, namespace, xNomadToken, idempotencyToken, pluginId, snapshotId, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteSnapshot(rsp)));
            }));
    }

    /**
     * @param volumeId Volume unique identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     * @param force Used to force the de-registration of a volume.
     */
    public deleteVolumeRegistration(volumeId: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, force?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.deleteVolumeRegistration(volumeId, region, namespace, xNomadToken, idempotencyToken, force, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.deleteVolumeRegistration(rsp)));
            }));
    }

    /**
     * @param volumeId Volume unique identifier.
     * @param action The action to perform on the Volume (create, detach, delete).
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     * @param node Specifies node to target volume operation for.
     */
    public detachOrDeleteVolume(volumeId: string, action: string, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, node?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.detachOrDeleteVolume(volumeId, action, region, namespace, xNomadToken, idempotencyToken, node, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.detachOrDeleteVolume(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param pluginId Filters volume lists by plugin ID.
     */
    public getExternalVolumes(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, pluginId?: string, _options?: Configuration): Observable<CSIVolumeListExternalResponse> {
        const requestContextPromise = this.requestFactory.getExternalVolumes(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, pluginId, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getExternalVolumes(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param pluginId Filters volume lists by plugin ID.
     */
    public getSnapshots(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, pluginId?: string, _options?: Configuration): Observable<CSISnapshotListResponse> {
        const requestContextPromise = this.requestFactory.getSnapshots(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, pluginId, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getSnapshots(rsp)));
            }));
    }

    /**
     * @param volumeId Volume unique identifier.
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     */
    public getVolume(volumeId: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Observable<CSIVolume> {
        const requestContextPromise = this.requestFactory.getVolume(volumeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getVolume(rsp)));
            }));
    }

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param wait Provided with IndexParam to wait for change.
     * @param stale If present, results will include stale reads.
     * @param prefix Constrains results to jobs that start with the defined prefix
     * @param xNomadToken A Nomad ACL token.
     * @param perPage Maximum number of results to return.
     * @param nextToken Indicates where to start paging for queries that support pagination.
     * @param nodeId Filters volume lists by node ID.
     * @param pluginId Filters volume lists by plugin ID.
     * @param type Filters volume lists to a specific type.
     */
    public getVolumes(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, nodeId?: string, pluginId?: string, type?: string, _options?: Configuration): Observable<Array<CSIVolumeListStub>> {
        const requestContextPromise = this.requestFactory.getVolumes(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, nodeId, pluginId, type, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.getVolumes(rsp)));
            }));
    }

    /**
     * @param cSISnapshotCreateRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postSnapshot(cSISnapshotCreateRequest: CSISnapshotCreateRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<CSISnapshotCreateResponse> {
        const requestContextPromise = this.requestFactory.postSnapshot(cSISnapshotCreateRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postSnapshot(rsp)));
            }));
    }

    /**
     * @param cSIVolumeRegisterRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postVolume(cSIVolumeRegisterRequest: CSIVolumeRegisterRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.postVolume(cSIVolumeRegisterRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postVolume(rsp)));
            }));
    }

    /**
     * @param volumeId Volume unique identifier.
     * @param cSIVolumeRegisterRequest 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public postVolumeRegistration(volumeId: string, cSIVolumeRegisterRequest: CSIVolumeRegisterRequest, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Observable<void> {
        const requestContextPromise = this.requestFactory.postVolumeRegistration(volumeId, cSIVolumeRegisterRequest, region, namespace, xNomadToken, idempotencyToken, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.postVolumeRegistration(rsp)));
            }));
    }

}
