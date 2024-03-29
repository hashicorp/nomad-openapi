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
import Deployment from '../model/Deployment';
import DeploymentAllocHealthRequest from '../model/DeploymentAllocHealthRequest';
import DeploymentPauseRequest from '../model/DeploymentPauseRequest';
import DeploymentPromoteRequest from '../model/DeploymentPromoteRequest';
import DeploymentUnblockRequest from '../model/DeploymentUnblockRequest';
import DeploymentUpdateResponse from '../model/DeploymentUpdateResponse';

/**
* Deployments service.
* @module api/DeploymentsApi
* @version 1.1.4
*/
export default class DeploymentsApi {

    /**
    * Constructs a new DeploymentsApi. 
    * @alias module:api/DeploymentsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the getDeployment operation.
     * @callback module:api/DeploymentsApi~getDeploymentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Deployment} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
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
     * @param {module:api/DeploymentsApi~getDeploymentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Deployment}
     */
    getDeployment(deploymentID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling getDeployment");
      }

      let pathParams = {
        'deploymentID': deploymentID
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
      let returnType = Deployment;
      return this.apiClient.callApi(
        '/deployment/{deploymentID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getDeploymentAllocations operation.
     * @callback module:api/DeploymentsApi~getDeploymentAllocationsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/AllocationListStub>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
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
     * @param {module:api/DeploymentsApi~getDeploymentAllocationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/AllocationListStub>}
     */
    getDeploymentAllocations(deploymentID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling getDeploymentAllocations");
      }

      let pathParams = {
        'deploymentID': deploymentID
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
        '/deployment/allocations/{deploymentID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getDeployments operation.
     * @callback module:api/DeploymentsApi~getDeploymentsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Deployment>} data The data returned by the service call.
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
     * @param {module:api/DeploymentsApi~getDeploymentsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Deployment>}
     */
    getDeployments(opts, callback) {
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
      let returnType = [Deployment];
      return this.apiClient.callApi(
        '/deployments', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postDeploymentAllocationHealth operation.
     * @callback module:api/DeploymentsApi~postDeploymentAllocationHealthCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeploymentUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
     * @param {module:model/DeploymentAllocHealthRequest} deploymentAllocHealthRequest 
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {String} opts.idempotencyToken Can be used to ensure operations are only run once.
     * @param {module:api/DeploymentsApi~postDeploymentAllocationHealthCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeploymentUpdateResponse}
     */
    postDeploymentAllocationHealth(deploymentID, deploymentAllocHealthRequest, opts, callback) {
      opts = opts || {};
      let postBody = deploymentAllocHealthRequest;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling postDeploymentAllocationHealth");
      }
      // verify the required parameter 'deploymentAllocHealthRequest' is set
      if (deploymentAllocHealthRequest === undefined || deploymentAllocHealthRequest === null) {
        throw new Error("Missing the required parameter 'deploymentAllocHealthRequest' when calling postDeploymentAllocationHealth");
      }

      let pathParams = {
        'deploymentID': deploymentID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'idempotency_token': opts['idempotencyToken']
      };
      let headerParams = {
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = DeploymentUpdateResponse;
      return this.apiClient.callApi(
        '/deployment/allocation-health/{deploymentID}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postDeploymentFail operation.
     * @callback module:api/DeploymentsApi~postDeploymentFailCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeploymentUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {String} opts.idempotencyToken Can be used to ensure operations are only run once.
     * @param {module:api/DeploymentsApi~postDeploymentFailCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeploymentUpdateResponse}
     */
    postDeploymentFail(deploymentID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling postDeploymentFail");
      }

      let pathParams = {
        'deploymentID': deploymentID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'idempotency_token': opts['idempotencyToken']
      };
      let headerParams = {
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = DeploymentUpdateResponse;
      return this.apiClient.callApi(
        '/deployment/fail/{deploymentID}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postDeploymentPause operation.
     * @callback module:api/DeploymentsApi~postDeploymentPauseCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeploymentUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
     * @param {module:model/DeploymentPauseRequest} deploymentPauseRequest 
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {String} opts.idempotencyToken Can be used to ensure operations are only run once.
     * @param {module:api/DeploymentsApi~postDeploymentPauseCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeploymentUpdateResponse}
     */
    postDeploymentPause(deploymentID, deploymentPauseRequest, opts, callback) {
      opts = opts || {};
      let postBody = deploymentPauseRequest;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling postDeploymentPause");
      }
      // verify the required parameter 'deploymentPauseRequest' is set
      if (deploymentPauseRequest === undefined || deploymentPauseRequest === null) {
        throw new Error("Missing the required parameter 'deploymentPauseRequest' when calling postDeploymentPause");
      }

      let pathParams = {
        'deploymentID': deploymentID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'idempotency_token': opts['idempotencyToken']
      };
      let headerParams = {
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = DeploymentUpdateResponse;
      return this.apiClient.callApi(
        '/deployment/pause/{deploymentID}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postDeploymentPromote operation.
     * @callback module:api/DeploymentsApi~postDeploymentPromoteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeploymentUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
     * @param {module:model/DeploymentPromoteRequest} deploymentPromoteRequest 
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {String} opts.idempotencyToken Can be used to ensure operations are only run once.
     * @param {module:api/DeploymentsApi~postDeploymentPromoteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeploymentUpdateResponse}
     */
    postDeploymentPromote(deploymentID, deploymentPromoteRequest, opts, callback) {
      opts = opts || {};
      let postBody = deploymentPromoteRequest;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling postDeploymentPromote");
      }
      // verify the required parameter 'deploymentPromoteRequest' is set
      if (deploymentPromoteRequest === undefined || deploymentPromoteRequest === null) {
        throw new Error("Missing the required parameter 'deploymentPromoteRequest' when calling postDeploymentPromote");
      }

      let pathParams = {
        'deploymentID': deploymentID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'idempotency_token': opts['idempotencyToken']
      };
      let headerParams = {
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = DeploymentUpdateResponse;
      return this.apiClient.callApi(
        '/deployment/promote/{deploymentID}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postDeploymentUnblock operation.
     * @callback module:api/DeploymentsApi~postDeploymentUnblockCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeploymentUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} deploymentID Deployment ID.
     * @param {module:model/DeploymentUnblockRequest} deploymentUnblockRequest 
     * @param {Object} opts Optional parameters
     * @param {String} opts.region Filters results based on the specified region.
     * @param {String} opts.namespace Filters results based on the specified namespace.
     * @param {String} opts.xNomadToken A Nomad ACL token.
     * @param {String} opts.idempotencyToken Can be used to ensure operations are only run once.
     * @param {module:api/DeploymentsApi~postDeploymentUnblockCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeploymentUpdateResponse}
     */
    postDeploymentUnblock(deploymentID, deploymentUnblockRequest, opts, callback) {
      opts = opts || {};
      let postBody = deploymentUnblockRequest;
      // verify the required parameter 'deploymentID' is set
      if (deploymentID === undefined || deploymentID === null) {
        throw new Error("Missing the required parameter 'deploymentID' when calling postDeploymentUnblock");
      }
      // verify the required parameter 'deploymentUnblockRequest' is set
      if (deploymentUnblockRequest === undefined || deploymentUnblockRequest === null) {
        throw new Error("Missing the required parameter 'deploymentUnblockRequest' when calling postDeploymentUnblock");
      }

      let pathParams = {
        'deploymentID': deploymentID
      };
      let queryParams = {
        'region': opts['region'],
        'namespace': opts['namespace'],
        'idempotency_token': opts['idempotencyToken']
      };
      let headerParams = {
        'X-Nomad-Token': opts['xNomadToken']
      };
      let formParams = {
      };

      let authNames = ['X-Nomad-Token'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = DeploymentUpdateResponse;
      return this.apiClient.callApi(
        '/deployment/unblock/{deploymentID}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
