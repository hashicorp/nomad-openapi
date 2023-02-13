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


import io.nomadproject.client.models.AllocationListStub;
import io.nomadproject.client.models.Evaluation;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class EvaluationsApi {
    private ApiClient localVarApiClient;

    public EvaluationsApi() {
        this(Configuration.getDefaultApiClient());
    }

    public EvaluationsApi(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for getEvaluation
     * @param evalID Evaluation ID. (required)
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
    public okhttp3.Call getEvaluationCall(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/evaluation/{evalID}"
            .replaceAll("\\{" + "evalID" + "\\}", localVarApiClient.escapeString(evalID.toString()));

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
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "X-Nomad-Token" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getEvaluationValidateBeforeCall(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'evalID' is set
        if (evalID == null) {
            throw new ApiException("Missing the required parameter 'evalID' when calling getEvaluation(Async)");
        }
        

        okhttp3.Call localVarCall = getEvaluationCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param evalID Evaluation ID. (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return Evaluation
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
    public Evaluation getEvaluation(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        ApiResponse<Evaluation> localVarResp = getEvaluationWithHttpInfo(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param evalID Evaluation ID. (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return ApiResponse&lt;Evaluation&gt;
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
    public ApiResponse<Evaluation> getEvaluationWithHttpInfo(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        okhttp3.Call localVarCall = getEvaluationValidateBeforeCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, null);
        Type localVarReturnType = new TypeToken<Evaluation>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
     * @param evalID Evaluation ID. (required)
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
    public okhttp3.Call getEvaluationAsync(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback<Evaluation> _callback) throws ApiException {

        okhttp3.Call localVarCall = getEvaluationValidateBeforeCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        Type localVarReturnType = new TypeToken<Evaluation>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getEvaluationAllocations
     * @param evalID Evaluation ID. (required)
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
    public okhttp3.Call getEvaluationAllocationsCall(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/evaluation/{evalID}/allocations"
            .replaceAll("\\{" + "evalID" + "\\}", localVarApiClient.escapeString(evalID.toString()));

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
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "X-Nomad-Token" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getEvaluationAllocationsValidateBeforeCall(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'evalID' is set
        if (evalID == null) {
            throw new ApiException("Missing the required parameter 'evalID' when calling getEvaluationAllocations(Async)");
        }
        

        okhttp3.Call localVarCall = getEvaluationAllocationsCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param evalID Evaluation ID. (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return List&lt;AllocationListStub&gt;
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
    public List<AllocationListStub> getEvaluationAllocations(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        ApiResponse<List<AllocationListStub>> localVarResp = getEvaluationAllocationsWithHttpInfo(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param evalID Evaluation ID. (required)
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return ApiResponse&lt;List&lt;AllocationListStub&gt;&gt;
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
    public ApiResponse<List<AllocationListStub>> getEvaluationAllocationsWithHttpInfo(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        okhttp3.Call localVarCall = getEvaluationAllocationsValidateBeforeCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, null);
        Type localVarReturnType = new TypeToken<List<AllocationListStub>>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
     * @param evalID Evaluation ID. (required)
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
    public okhttp3.Call getEvaluationAllocationsAsync(String evalID, String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback<List<AllocationListStub>> _callback) throws ApiException {

        okhttp3.Call localVarCall = getEvaluationAllocationsValidateBeforeCall(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        Type localVarReturnType = new TypeToken<List<AllocationListStub>>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getEvaluations
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
    public okhttp3.Call getEvaluationsCall(String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/evaluations";

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
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "X-Nomad-Token" };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getEvaluationsValidateBeforeCall(String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = getEvaluationsCall(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return List&lt;Evaluation&gt;
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
    public List<Evaluation> getEvaluations(String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        ApiResponse<List<Evaluation>> localVarResp = getEvaluationsWithHttpInfo(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param region Filters results based on the specified region. (optional)
     * @param namespace Filters results based on the specified namespace. (optional)
     * @param index If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
     * @param wait Provided with IndexParam to wait for change. (optional)
     * @param stale If present, results will include stale reads. (optional)
     * @param prefix Constrains results to jobs that start with the defined prefix (optional)
     * @param xNomadToken A Nomad ACL token. (optional)
     * @param perPage Maximum number of results to return. (optional)
     * @param nextToken Indicates where to start paging for queries that support pagination. (optional)
     * @return ApiResponse&lt;List&lt;Evaluation&gt;&gt;
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
    public ApiResponse<List<Evaluation>> getEvaluationsWithHttpInfo(String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken) throws ApiException {
        okhttp3.Call localVarCall = getEvaluationsValidateBeforeCall(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, null);
        Type localVarReturnType = new TypeToken<List<Evaluation>>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
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
    public okhttp3.Call getEvaluationsAsync(String region, String namespace, Integer index, String wait, String stale, String prefix, String xNomadToken, Integer perPage, String nextToken, final ApiCallback<List<Evaluation>> _callback) throws ApiException {

        okhttp3.Call localVarCall = getEvaluationsValidateBeforeCall(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, _callback);
        Type localVarReturnType = new TypeToken<List<Evaluation>>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}
