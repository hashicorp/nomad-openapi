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
import Node from '../model/Node';
import NodeDrainUpdateResponse from '../model/NodeDrainUpdateResponse';
import NodeEligibilityUpdateResponse from '../model/NodeEligibilityUpdateResponse';
import NodeListStub from '../model/NodeListStub';
import NodePurgeResponse from '../model/NodePurgeResponse';
import NodeUpdateDrainRequest from '../model/NodeUpdateDrainRequest';
import NodeUpdateEligibilityRequest from '../model/NodeUpdateEligibilityRequest';

/**
* Nodes service.
* @module api/NodesApi
* @version 1.1.4
*/
export default class NodesApi {

    /**
    * Constructs a new NodesApi. 
    * @alias module:api/NodesApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the getNode operation.
     * @callback module:api/NodesApi~getNodeCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Node} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} nodeId The ID of the node.
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
     * @param {module:api/NodesApi~getNodeCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Node}
     */
    getNode(nodeId, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'nodeId' is set
      if (nodeId === undefined || nodeId === null) {
        throw new Error("Missing the required parameter 'nodeId' when calling getNode");
      }

      let pathParams = {
        'nodeId': nodeId
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
      let returnType = Node;
      return this.apiClient.callApi(
        '/node/{nodeId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getNodeAllocations operation.
     * @callback module:api/NodesApi~getNodeAllocationsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/AllocationListStub>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} nodeId The ID of the node.
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
     * @param {module:api/NodesApi~getNodeAllocationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/AllocationListStub>}
     */
    getNodeAllocations(nodeId, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'nodeId' is set
      if (nodeId === undefined || nodeId === null) {
        throw new Error("Missing the required parameter 'nodeId' when calling getNodeAllocations");
      }

      let pathParams = {
        'nodeId': nodeId
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
        '/node/{nodeId}/allocations', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getNodes operation.
     * @callback module:api/NodesApi~getNodesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/NodeListStub>} data The data returned by the service call.
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
     * @param {Boolean} opts.resources Whether or not to include the NodeResources and ReservedResources fields in the response.
     * @param {module:api/NodesApi~getNodesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/NodeListStub>}
     */
    getNodes(opts, callback) {
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
        'next_token': opts['nextToken'],
        'resources': opts['resources']
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
      let returnType = [NodeListStub];
      return this.apiClient.callApi(
        '/nodes', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateNodeDrain operation.
     * @callback module:api/NodesApi~updateNodeDrainCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NodeDrainUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} nodeId The ID of the node.
     * @param {module:model/NodeUpdateDrainRequest} nodeUpdateDrainRequest 
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
     * @param {module:api/NodesApi~updateNodeDrainCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NodeDrainUpdateResponse}
     */
    updateNodeDrain(nodeId, nodeUpdateDrainRequest, opts, callback) {
      opts = opts || {};
      let postBody = nodeUpdateDrainRequest;
      // verify the required parameter 'nodeId' is set
      if (nodeId === undefined || nodeId === null) {
        throw new Error("Missing the required parameter 'nodeId' when calling updateNodeDrain");
      }
      // verify the required parameter 'nodeUpdateDrainRequest' is set
      if (nodeUpdateDrainRequest === undefined || nodeUpdateDrainRequest === null) {
        throw new Error("Missing the required parameter 'nodeUpdateDrainRequest' when calling updateNodeDrain");
      }

      let pathParams = {
        'nodeId': nodeId
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
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = NodeDrainUpdateResponse;
      return this.apiClient.callApi(
        '/node/{nodeId}/drain', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateNodeEligibility operation.
     * @callback module:api/NodesApi~updateNodeEligibilityCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NodeEligibilityUpdateResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} nodeId The ID of the node.
     * @param {module:model/NodeUpdateEligibilityRequest} nodeUpdateEligibilityRequest 
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
     * @param {module:api/NodesApi~updateNodeEligibilityCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NodeEligibilityUpdateResponse}
     */
    updateNodeEligibility(nodeId, nodeUpdateEligibilityRequest, opts, callback) {
      opts = opts || {};
      let postBody = nodeUpdateEligibilityRequest;
      // verify the required parameter 'nodeId' is set
      if (nodeId === undefined || nodeId === null) {
        throw new Error("Missing the required parameter 'nodeId' when calling updateNodeEligibility");
      }
      // verify the required parameter 'nodeUpdateEligibilityRequest' is set
      if (nodeUpdateEligibilityRequest === undefined || nodeUpdateEligibilityRequest === null) {
        throw new Error("Missing the required parameter 'nodeUpdateEligibilityRequest' when calling updateNodeEligibility");
      }

      let pathParams = {
        'nodeId': nodeId
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
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = NodeEligibilityUpdateResponse;
      return this.apiClient.callApi(
        '/node/{nodeId}/eligibility', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateNodePurge operation.
     * @callback module:api/NodesApi~updateNodePurgeCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NodePurgeResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} nodeId The ID of the node.
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
     * @param {module:api/NodesApi~updateNodePurgeCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NodePurgeResponse}
     */
    updateNodePurge(nodeId, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'nodeId' is set
      if (nodeId === undefined || nodeId === null) {
        throw new Error("Missing the required parameter 'nodeId' when calling updateNodePurge");
      }

      let pathParams = {
        'nodeId': nodeId
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
      let returnType = NodePurgeResponse;
      return this.apiClient.callApi(
        '/node/{nodeId}/purge', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
