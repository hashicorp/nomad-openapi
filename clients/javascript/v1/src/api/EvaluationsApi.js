/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

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


import ApiClient from "../ApiClient";
import AllocationListStub from '../model/AllocationListStub';
import Evaluation from '../model/Evaluation';

/**
* Evaluations service.
* @module api/EvaluationsApi
* @version 1.1.4
*/
export default class EvaluationsApi {

    /**
    * Constructs a new EvaluationsApi. 
    * @alias module:api/EvaluationsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the getEvaluation operation.
     * @callback module:api/EvaluationsApi~getEvaluationCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Evaluation} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} evalID Evaluation ID.
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {Number} opts.index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param {String} opts.wait Provided with IndexParam to wait for change.
     * @param {String} opts.stale If present, results will include stale reads.
     * @param {String} opts.prefix Constrains results to jobs that start with the defined prefix
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {Number} opts.perPage Maximum number of results to return.
     * @param {String} opts.nextToken Indicates where to start paging for queries that support pagination.
     * @param {module:api/EvaluationsApi~getEvaluationCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Evaluation}
     */
    getEvaluation(evalID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'evalID' is set
      if (evalID === undefined || evalID === null) {
        throw new Error("Missing the required parameter 'evalID' when calling getEvaluation");
      }

      let pathParams = {
        'evalID': evalID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'wait': opts['wait'],
        'stale': opts['stale'],
        'prefix': opts['prefix'],
        'per_page': opts['perPage'],
        'next_token': opts['nextToken']
      };
      let headerParams = {
        'index': opts['index'],
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Evaluation;
      return this.apiClient.callApi(
        '/evaluation/{evalID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getEvaluationAllocations operation.
     * @callback module:api/EvaluationsApi~getEvaluationAllocationsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/AllocationListStub>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} evalID Evaluation ID.
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {Number} opts.index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param {String} opts.wait Provided with IndexParam to wait for change.
     * @param {String} opts.stale If present, results will include stale reads.
     * @param {String} opts.prefix Constrains results to jobs that start with the defined prefix
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {Number} opts.perPage Maximum number of results to return.
     * @param {String} opts.nextToken Indicates where to start paging for queries that support pagination.
     * @param {module:api/EvaluationsApi~getEvaluationAllocationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/AllocationListStub>}
     */
    getEvaluationAllocations(evalID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'evalID' is set
      if (evalID === undefined || evalID === null) {
        throw new Error("Missing the required parameter 'evalID' when calling getEvaluationAllocations");
      }

      let pathParams = {
        'evalID': evalID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'wait': opts['wait'],
        'stale': opts['stale'],
        'prefix': opts['prefix'],
        'per_page': opts['perPage'],
        'next_token': opts['nextToken']
      };
      let headerParams = {
        'index': opts['index'],
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [AllocationListStub];
      return this.apiClient.callApi(
        '/evaluation/{evalID}/allocations', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getEvaluations operation.
     * @callback module:api/EvaluationsApi~getEvaluationsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Evaluation>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {Number} opts.index If set, wait until query exceeds given index. Must be provided with WaitParam.
     * @param {String} opts.wait Provided with IndexParam to wait for change.
     * @param {String} opts.stale If present, results will include stale reads.
     * @param {String} opts.prefix Constrains results to jobs that start with the defined prefix
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {Number} opts.perPage Maximum number of results to return.
     * @param {String} opts.nextToken Indicates where to start paging for queries that support pagination.
     * @param {module:api/EvaluationsApi~getEvaluationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Evaluation>}
     */
    getEvaluations(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'wait': opts['wait'],
        'stale': opts['stale'],
        'prefix': opts['prefix'],
        'per_page': opts['perPage'],
        'next_token': opts['nextToken']
      };
      let headerParams = {
        'index': opts['index'],
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Evaluation];
      return this.apiClient.callApi(
        '/evaluations', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
