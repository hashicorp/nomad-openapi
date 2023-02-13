/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

/*
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.nomadproject.client.api;

import io.nomadproject.client.ApiCallback;
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.ApiResponse;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.Pair;
import io.nomadproject.client.ProgressRequestBody;
import io.nomadproject.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import io.nomadproject.client.models.FuzzySearchRequest;
import io.nomadproject.client.models.FuzzySearchResponse;
import io.nomadproject.client.models.SearchRequest;
import io.nomadproject.client.models.SearchResponse;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class SearchApi {
    private ApiClient localVarApiClient;

    public SearchApi() {
        this(Configuration.getDefaultApiClient());
    }

    public SearchApi(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for getFuzzySearch
     * @param fuzzySearchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getFuzzySearchCall(FuzzySearchRequest fuzzySearchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = fuzzySearchRequest;

        // create path and map variables
        String localVarPath = "/search/fuzzy";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        if (region != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("region", region));
        }

        if (namespace != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("namespace", namespace));
        }

        if (wait != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("wait", wait));
        }

        if (stale != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("stale", stale));
        }

        if (prefix != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("prefix", prefix));
        }

        if (perPage != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("per_page", perPage));
        }

        if (nextToken != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("next_token", nextToken));
        }

        if (index != null) {
            localVarHeaderParams.put("index", localVarApiClient.parameterToString(index));
        }

        if (xNomadToken != null) {
            localVarHeaderParams.put("X-Nomad-Token", localVarApiClient.parameterToString(xNomadToken));
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "X-Nomad-Token" };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getFuzzySearchValidateBeforeCall(FuzzySearchRequest fuzzySearchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'fuzzySearchRequest' is set
        if (fuzzySearchRequest == null) {
            throw new ApiException("Missing the required parameter 'fuzzySearchRequest' when calling getFuzzySearch(Async)");
        }
        

        okhttp3.Call localVarCall = getFuzzySearchCall(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param fuzzySearchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return FuzzySearchResponse
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public FuzzySearchResponse getFuzzySearch(FuzzySearchRequest fuzzySearchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        ApiResponse<FuzzySearchResponse> localVarResp = getFuzzySearchWithHttpInfo(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param fuzzySearchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return ApiResponse&lt;FuzzySearchResponse&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<FuzzySearchResponse> getFuzzySearchWithHttpInfo(FuzzySearchRequest fuzzySearchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        okhttp3.Call localVarCall = getFuzzySearchValidateBeforeCall(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, null);
        Type localVarReturnType = new TypeToken<FuzzySearchResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
     * @param fuzzySearchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getFuzzySearchAsync(FuzzySearchRequest fuzzySearchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback<FuzzySearchResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = getFuzzySearchValidateBeforeCall(fuzzySearchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        Type localVarReturnType = new TypeToken<FuzzySearchResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getSearch
     * @param searchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getSearchCall(SearchRequest searchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = searchRequest;

        // create path and map variables
        String localVarPath = "/search";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        if (region != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("region", region));
        }

        if (namespace != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("namespace", namespace));
        }

        if (wait != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("wait", wait));
        }

        if (stale != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("stale", stale));
        }

        if (prefix != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("prefix", prefix));
        }

        if (perPage != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("per_page", perPage));
        }

        if (nextToken != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("next_token", nextToken));
        }

        if (index != null) {
            localVarHeaderParams.put("index", localVarApiClient.parameterToString(index));
        }

        if (xNomadToken != null) {
            localVarHeaderParams.put("X-Nomad-Token", localVarApiClient.parameterToString(xNomadToken));
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "X-Nomad-Token" };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getSearchValidateBeforeCall(SearchRequest searchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'searchRequest' is set
        if (searchRequest == null) {
            throw new ApiException("Missing the required parameter 'searchRequest' when calling getSearch(Async)");
        }
        

        okhttp3.Call localVarCall = getSearchCall(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param searchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return SearchResponse
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public SearchResponse getSearch(SearchRequest searchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        ApiResponse<SearchResponse> localVarResp = getSearchWithHttpInfo(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param searchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return ApiResponse&lt;SearchResponse&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<SearchResponse> getSearchWithHttpInfo(SearchRequest searchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        okhttp3.Call localVarCall = getSearchValidateBeforeCall(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, null);
        Type localVarReturnType = new TypeToken<SearchResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
     * @param searchRequest  (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td>  </td><td>  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  </td></tr>
        <tr><td> 400 </td><td> Bad request </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> Forbidden </td><td>  -  </td></tr>
        <tr><td> 405 </td><td> Method not allowed </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Internal server error </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getSearchAsync(SearchRequest searchRequest, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback<SearchResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = getSearchValidateBeforeCall(searchRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        Type localVarReturnType = new TypeToken<SearchResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}
