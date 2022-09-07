// TODO: better import syntax?
import {BaseAPIRequestFactory, RequiredError} from './baseapi';
import {Configuration} from '../configuration';
import {RequestContext, HttpMethod, ResponseContext, HttpFile} from '../http/http';
import {ObjectSerializer} from '../models/ObjectSerializer';
import {ApiException} from './exception';
import {canConsumeForm, isCodeInRange} from '../util';
import {SecurityAuthentication} from '../auth/auth';


import { AutopilotConfiguration } from '../models/AutopilotConfiguration';
import { OperatorHealthReply } from '../models/OperatorHealthReply';
import { RaftConfiguration } from '../models/RaftConfiguration';
import { SchedulerConfiguration } from '../models/SchedulerConfiguration';
import { SchedulerConfigurationResponse } from '../models/SchedulerConfigurationResponse';
import { SchedulerSetConfigurationResponse } from '../models/SchedulerSetConfigurationResponse';

/**
 * no description
 */
export class OperatorApiRequestFactory extends BaseAPIRequestFactory {

    /**
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public async deleteOperatorRaftPeer(region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;





        // Path Params
        const localVarPath = '/operator/raft/peer';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.DELETE);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (idempotencyToken !== undefined) {
            requestContext.setQueryParam("idempotency_token", ObjectSerializer.serialize(idempotencyToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
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
    public async getOperatorAutopilotConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;










        // Path Params
        const localVarPath = '/operator/autopilot/configuration';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }

        // Query Params
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }

        // Query Params
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }

        // Query Params
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }

        // Query Params
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
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
    public async getOperatorAutopilotHealth(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;










        // Path Params
        const localVarPath = '/operator/autopilot/health';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }

        // Query Params
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }

        // Query Params
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }

        // Query Params
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }

        // Query Params
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
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
    public async getOperatorRaftConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;










        // Path Params
        const localVarPath = '/operator/raft/configuration';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }

        // Query Params
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }

        // Query Params
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }

        // Query Params
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }

        // Query Params
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
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
    public async getOperatorSchedulerConfiguration(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;










        // Path Params
        const localVarPath = '/operator/scheduler/configuration';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }

        // Query Params
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }

        // Query Params
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }

        // Query Params
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }

        // Query Params
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

    /**
     * @param schedulerConfiguration 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public async postOperatorSchedulerConfiguration(schedulerConfiguration: SchedulerConfiguration, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'schedulerConfiguration' is not null or undefined
        if (schedulerConfiguration === null || schedulerConfiguration === undefined) {
            throw new RequiredError("OperatorApi", "postOperatorSchedulerConfiguration", "schedulerConfiguration");
        }






        // Path Params
        const localVarPath = '/operator/scheduler/configuration';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.POST);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (idempotencyToken !== undefined) {
            requestContext.setQueryParam("idempotency_token", ObjectSerializer.serialize(idempotencyToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        // Body Params
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(schedulerConfiguration, "SchedulerConfiguration", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

    /**
     * @param autopilotConfiguration 
     * @param region Filters results based on the specified region.
     * @param namespace Filters results based on the specified namespace.
     * @param xNomadToken A Nomad ACL token.
     * @param idempotencyToken Can be used to ensure operations are only run once.
     */
    public async putOperatorAutopilotConfiguration(autopilotConfiguration: AutopilotConfiguration, region?: string, namespace?: string, xNomadToken?: string, idempotencyToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'autopilotConfiguration' is not null or undefined
        if (autopilotConfiguration === null || autopilotConfiguration === undefined) {
            throw new RequiredError("OperatorApi", "putOperatorAutopilotConfiguration", "autopilotConfiguration");
        }






        // Path Params
        const localVarPath = '/operator/autopilot/configuration';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.PUT);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }

        // Query Params
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }

        // Query Params
        if (idempotencyToken !== undefined) {
            requestContext.setQueryParam("idempotency_token", ObjectSerializer.serialize(idempotencyToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));


        // Body Params
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(autopilotConfiguration, "AutopilotConfiguration", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        let authMethod: SecurityAuthentication | undefined;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod?.applySecurityAuthentication) {
            await authMethod?.applySecurityAuthentication(requestContext);
        }
        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

}

export class OperatorApiResponseProcessor {

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to deleteOperatorRaftPeer
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async deleteOperatorRaftPeer(response: ResponseContext): Promise<void > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            return;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: void = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "void", ""
            ) as void;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getOperatorAutopilotConfiguration
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getOperatorAutopilotConfiguration(response: ResponseContext): Promise<AutopilotConfiguration > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: AutopilotConfiguration = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "AutopilotConfiguration", ""
            ) as AutopilotConfiguration;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: AutopilotConfiguration = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "AutopilotConfiguration", ""
            ) as AutopilotConfiguration;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getOperatorAutopilotHealth
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getOperatorAutopilotHealth(response: ResponseContext): Promise<OperatorHealthReply > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: OperatorHealthReply = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "OperatorHealthReply", ""
            ) as OperatorHealthReply;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: OperatorHealthReply = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "OperatorHealthReply", ""
            ) as OperatorHealthReply;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getOperatorRaftConfiguration
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getOperatorRaftConfiguration(response: ResponseContext): Promise<RaftConfiguration > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: RaftConfiguration = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "RaftConfiguration", ""
            ) as RaftConfiguration;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: RaftConfiguration = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "RaftConfiguration", ""
            ) as RaftConfiguration;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getOperatorSchedulerConfiguration
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getOperatorSchedulerConfiguration(response: ResponseContext): Promise<SchedulerConfigurationResponse > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: SchedulerConfigurationResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SchedulerConfigurationResponse", ""
            ) as SchedulerConfigurationResponse;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: SchedulerConfigurationResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SchedulerConfigurationResponse", ""
            ) as SchedulerConfigurationResponse;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to postOperatorSchedulerConfiguration
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async postOperatorSchedulerConfiguration(response: ResponseContext): Promise<SchedulerSetConfigurationResponse > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: SchedulerSetConfigurationResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SchedulerSetConfigurationResponse", ""
            ) as SchedulerSetConfigurationResponse;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: SchedulerSetConfigurationResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SchedulerSetConfigurationResponse", ""
            ) as SchedulerSetConfigurationResponse;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to putOperatorAutopilotConfiguration
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async putOperatorAutopilotConfiguration(response: ResponseContext): Promise<boolean > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: boolean = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "boolean", ""
            ) as boolean;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Bad request", undefined, response.headers);
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Forbidden", undefined, response.headers);
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Method not allowed", undefined, response.headers);
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<undefined>(response.httpStatusCode, "Internal server error", undefined, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: boolean = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "boolean", ""
            ) as boolean;
            return body;
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

}
