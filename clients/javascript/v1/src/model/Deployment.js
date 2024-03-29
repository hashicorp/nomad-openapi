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
import DeploymentState from './DeploymentState';

/**
 * The Deployment model module.
 * @module model/Deployment
 * @version 1.1.4
 */
class Deployment {
    /**
     * Constructs a new <code>Deployment</code>.
     * @alias module:model/Deployment
     */
    constructor() { 
        
        Deployment.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Deployment</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Deployment} obj Optional instance to populate.
     * @return {module:model/Deployment} The populated <code>Deployment</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Deployment();

            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('IsMultiregion')) {
                obj['IsMultiregion'] = ApiClient.convertToType(data['IsMultiregion'], 'Boolean');
            }
            if (data.hasOwnProperty('JobCreateIndex')) {
                obj['JobCreateIndex'] = ApiClient.convertToType(data['JobCreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
            }
            if (data.hasOwnProperty('JobModifyIndex')) {
                obj['JobModifyIndex'] = ApiClient.convertToType(data['JobModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('JobSpecModifyIndex')) {
                obj['JobSpecModifyIndex'] = ApiClient.convertToType(data['JobSpecModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('JobVersion')) {
                obj['JobVersion'] = ApiClient.convertToType(data['JobVersion'], 'Number');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Status')) {
                obj['Status'] = ApiClient.convertToType(data['Status'], 'String');
            }
            if (data.hasOwnProperty('StatusDescription')) {
                obj['StatusDescription'] = ApiClient.convertToType(data['StatusDescription'], 'String');
            }
            if (data.hasOwnProperty('TaskGroups')) {
                obj['TaskGroups'] = ApiClient.convertToType(data['TaskGroups'], {'String': DeploymentState});
            }
        }
        return obj;
    }


}

/**
 * @member {Number} CreateIndex
 */
Deployment.prototype['CreateIndex'] = undefined;

/**
 * @member {String} ID
 */
Deployment.prototype['ID'] = undefined;

/**
 * @member {Boolean} IsMultiregion
 */
Deployment.prototype['IsMultiregion'] = undefined;

/**
 * @member {Number} JobCreateIndex
 */
Deployment.prototype['JobCreateIndex'] = undefined;

/**
 * @member {String} JobID
 */
Deployment.prototype['JobID'] = undefined;

/**
 * @member {Number} JobModifyIndex
 */
Deployment.prototype['JobModifyIndex'] = undefined;

/**
 * @member {Number} JobSpecModifyIndex
 */
Deployment.prototype['JobSpecModifyIndex'] = undefined;

/**
 * @member {Number} JobVersion
 */
Deployment.prototype['JobVersion'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
Deployment.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Namespace
 */
Deployment.prototype['Namespace'] = undefined;

/**
 * @member {String} Status
 */
Deployment.prototype['Status'] = undefined;

/**
 * @member {String} StatusDescription
 */
Deployment.prototype['StatusDescription'] = undefined;

/**
 * @member {Object.<String, module:model/DeploymentState>} TaskGroups
 */
Deployment.prototype['TaskGroups'] = undefined;






export default Deployment;

