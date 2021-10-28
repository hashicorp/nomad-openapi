// TODO: better import syntax?
import { BaseAPIRequestFactory, RequiredError } from './baseapi';
import {Configuration} from '../configuration';
import { RequestContext, HttpMethod, ResponseContext, HttpFile} from '../http/http';
import {ObjectSerializer} from '../models/ObjectSerializer';
import {ApiException} from './exception';
import {isCodeInRange} from '../util';

import { CSIPlugin } from '../models/CSIPlugin';
import { CSIPluginListStub } from '../models/CSIPluginListStub';

/**
 * no description
 */
export class PluginsApiRequestFactory extends BaseAPIRequestFactory {

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
    public async getPluginCSI(pluginID: string, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'pluginID' is not null or undefined
        if (pluginID === null || pluginID === undefined) {
            throw new RequiredError('Required parameter pluginID was null or undefined when calling getPluginCSI.');
        }











        // Path Params
        const localVarPath = '/plugin/csi/{pluginID}'
            .replace('{' + 'pluginID' + '}', encodeURIComponent(String(pluginID)));

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));

        // Form Params


        // Body Params

        let authMethod = null;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod) {
            await authMethod.applySecurityAuthentication(requestContext);
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
    public async getPlugins(region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;










        // Path Params
        const localVarPath = '/plugins';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.GET);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")

        // Query Params
        if (region !== undefined) {
            requestContext.setQueryParam("region", ObjectSerializer.serialize(region, "string", ""));
        }
        if (namespace !== undefined) {
            requestContext.setQueryParam("namespace", ObjectSerializer.serialize(namespace, "string", ""));
        }
        if (wait !== undefined) {
            requestContext.setQueryParam("wait", ObjectSerializer.serialize(wait, "string", ""));
        }
        if (stale !== undefined) {
            requestContext.setQueryParam("stale", ObjectSerializer.serialize(stale, "string", ""));
        }
        if (prefix !== undefined) {
            requestContext.setQueryParam("prefix", ObjectSerializer.serialize(prefix, "string", ""));
        }
        if (perPage !== undefined) {
            requestContext.setQueryParam("per_page", ObjectSerializer.serialize(perPage, "number", ""));
        }
        if (nextToken !== undefined) {
            requestContext.setQueryParam("next_token", ObjectSerializer.serialize(nextToken, "string", ""));
        }

        // Header Params
        requestContext.setHeaderParam("index", ObjectSerializer.serialize(index, "number", ""));
        requestContext.setHeaderParam("X-Nomad-Token", ObjectSerializer.serialize(xNomadToken, "string", ""));

        // Form Params


        // Body Params

        let authMethod = null;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod) {
            await authMethod.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

}

export class PluginsApiResponseProcessor {

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getPluginCSI
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getPluginCSI(response: ResponseContext): Promise<Array<CSIPlugin> > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: Array<CSIPlugin> = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "Array<CSIPlugin>", ""
            ) as Array<CSIPlugin>;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Bad request");
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Forbidden");
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Method not allowed");
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Internal server error");
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: Array<CSIPlugin> = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "Array<CSIPlugin>", ""
            ) as Array<CSIPlugin>;
            return body;
        }

        let body = response.body || "";
        throw new ApiException<string>(response.httpStatusCode, "Unknown API Status Code!\nBody: \"" + body + "\"");
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getPlugins
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getPlugins(response: ResponseContext): Promise<Array<CSIPluginListStub> > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: Array<CSIPluginListStub> = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "Array<CSIPluginListStub>", ""
            ) as Array<CSIPluginListStub>;
            return body;
        }
        if (isCodeInRange("400", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Bad request");
        }
        if (isCodeInRange("403", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Forbidden");
        }
        if (isCodeInRange("405", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Method not allowed");
        }
        if (isCodeInRange("500", response.httpStatusCode)) {
            throw new ApiException<string>(response.httpStatusCode, "Internal server error");
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: Array<CSIPluginListStub> = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "Array<CSIPluginListStub>", ""
            ) as Array<CSIPluginListStub>;
            return body;
        }

        let body = response.body || "";
        throw new ApiException<string>(response.httpStatusCode, "Unknown API Status Code!\nBody: \"" + body + "\"");
    }

}
