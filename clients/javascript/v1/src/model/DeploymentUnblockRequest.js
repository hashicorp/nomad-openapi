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
 * The DeploymentUnblockRequest model module.
 * @module model/DeploymentUnblockRequest
 * @version 1.1.4
 */
class DeploymentUnblockRequest {
    /**
     * Constructs a new <code>DeploymentUnblockRequest</code>.
     * @alias module:model/DeploymentUnblockRequest
     */
    constructor() { 
        
        DeploymentUnblockRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DeploymentUnblockRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DeploymentUnblockRequest} obj Optional instance to populate.
     * @return {module:model/DeploymentUnblockRequest} The populated <code>DeploymentUnblockRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DeploymentUnblockRequest();

            if (data.hasOwnProperty('DeploymentID')) {
                obj['DeploymentID'] = ApiClient.convertToType(data['DeploymentID'], 'String');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Region')) {
                obj['Region'] = ApiClient.convertToType(data['Region'], 'String');
            }
            if (data.hasOwnProperty('SecretID')) {
                obj['SecretID'] = ApiClient.convertToType(data['SecretID'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} DeploymentID
 */
DeploymentUnblockRequest.prototype['DeploymentID'] = undefined;

/**
 * @member {String} Namespace
 */
DeploymentUnblockRequest.prototype['Namespace'] = undefined;

/**
 * @member {String} Region
 */
DeploymentUnblockRequest.prototype['Region'] = undefined;

/**
 * @member {String} SecretID
 */
DeploymentUnblockRequest.prototype['SecretID'] = undefined;






export default DeploymentUnblockRequest;

