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

import ApiClient from '../ApiClient';

/**
 * The NodePurgeResponse model module.
 * @module model/NodePurgeResponse
 * @version 1.1.4
 */
class NodePurgeResponse {
    /**
     * Constructs a new <code>NodePurgeResponse</code>.
     * @alias module:model/NodePurgeResponse
     */
    constructor() { 
        
        NodePurgeResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NodePurgeResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NodePurgeResponse} obj Optional instance to populate.
     * @return {module:model/NodePurgeResponse} The populated <code>NodePurgeResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NodePurgeResponse();

            if (data.hasOwnProperty('EvalCreateIndex')) {
                obj['EvalCreateIndex'] = ApiClient.convertToType(data['EvalCreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('EvalIDs')) {
                obj['EvalIDs'] = ApiClient.convertToType(data['EvalIDs'], ['String']);
            }
            if (data.hasOwnProperty('NodeModifyIndex')) {
                obj['NodeModifyIndex'] = ApiClient.convertToType(data['NodeModifyIndex'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} EvalCreateIndex
 */
NodePurgeResponse.prototype['EvalCreateIndex'] = undefined;

/**
 * @member {Array.<String>} EvalIDs
 */
NodePurgeResponse.prototype['EvalIDs'] = undefined;

/**
 * @member {Number} NodeModifyIndex
 */
NodePurgeResponse.prototype['NodeModifyIndex'] = undefined;






export default NodePurgeResponse;

