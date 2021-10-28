// TODO: better import syntax?
import { BaseAPIRequestFactory, RequiredError } from './baseapi';
import {Configuration} from '../configuration';
import { RequestContext, HttpMethod, ResponseContext, HttpFile} from '../http/http';
import {ObjectSerializer} from '../models/ObjectSerializer';
import {ApiException} from './exception';
import {isCodeInRange} from '../util';

import { FuzzySearchRequest } from '../models/FuzzySearchRequest';
import { FuzzySearchResponse } from '../models/FuzzySearchResponse';
import { SearchRequest } from '../models/SearchRequest';
import { SearchResponse } from '../models/SearchResponse';

/**
 * no description
 */
export class SearchApiRequestFactory extends BaseAPIRequestFactory {

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
    public async getFuzzySearch(fuzzySearchRequest: FuzzySearchRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'fuzzySearchRequest' is not null or undefined
        if (fuzzySearchRequest === null || fuzzySearchRequest === undefined) {
            throw new RequiredError('Required parameter fuzzySearchRequest was null or undefined when calling getFuzzySearch.');
        }











        // Path Params
        const localVarPath = '/search/fuzzy';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.POST);
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
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(fuzzySearchRequest, "FuzzySearchRequest", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        let authMethod = null;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod) {
            await authMethod.applySecurityAuthentication(requestContext);
        }

        return requestContext;
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
    public async getSearch(searchRequest: SearchRequest, region?: string, namespace?: string, index?: number, wait?: string, stale?: string, prefix?: string, xNomadToken?: string, perPage?: number, nextToken?: string, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'searchRequest' is not null or undefined
        if (searchRequest === null || searchRequest === undefined) {
            throw new RequiredError('Required parameter searchRequest was null or undefined when calling getSearch.');
        }











        // Path Params
        const localVarPath = '/search';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.POST);
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
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(searchRequest, "SearchRequest", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        let authMethod = null;
        // Apply auth methods
        authMethod = _config.authMethods["X-Nomad-Token"]
        if (authMethod) {
            await authMethod.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

}

export class SearchApiResponseProcessor {

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getFuzzySearch
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getFuzzySearch(response: ResponseContext): Promise<FuzzySearchResponse > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: FuzzySearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "FuzzySearchResponse", ""
            ) as FuzzySearchResponse;
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
            const body: FuzzySearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "FuzzySearchResponse", ""
            ) as FuzzySearchResponse;
            return body;
        }

        let body = response.body || "";
        throw new ApiException<string>(response.httpStatusCode, "Unknown API Status Code!\nBody: \"" + body + "\"");
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to getSearch
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async getSearch(response: ResponseContext): Promise<SearchResponse > {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
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
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
            return body;
        }

        let body = response.body || "";
        throw new ApiException<string>(response.httpStatusCode, "Unknown API Status Code!\nBody: \"" + body + "\"");
    }

}
