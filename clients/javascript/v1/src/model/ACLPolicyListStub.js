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
 * The ACLPolicyListStub model module.
 * @module model/ACLPolicyListStub
 * @version 1.1.4
 */
class ACLPolicyListStub {
    /**
     * Constructs a new <code>ACLPolicyListStub</code>.
     * @alias module:model/ACLPolicyListStub
     */
    constructor() { 
        
        ACLPolicyListStub.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ACLPolicyListStub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ACLPolicyListStub} obj Optional instance to populate.
     * @return {module:model/ACLPolicyListStub} The populated <code>ACLPolicyListStub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ACLPolicyListStub();

            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = ApiClient.convertToType(data['Description'], 'String');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} CreateIndex
 */
ACLPolicyListStub.prototype['CreateIndex'] = undefined;

/**
 * @member {String} Description
 */
ACLPolicyListStub.prototype['Description'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
ACLPolicyListStub.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Name
 */
ACLPolicyListStub.prototype['Name'] = undefined;






export default ACLPolicyListStub;

