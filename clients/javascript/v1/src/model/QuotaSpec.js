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
import QuotaLimit from './QuotaLimit';

/**
 * The QuotaSpec model module.
 * @module model/QuotaSpec
 * @version 1.1.4
 */
class QuotaSpec {
    /**
     * Constructs a new <code>QuotaSpec</code>.
     * @alias module:model/QuotaSpec
     */
    constructor() { 
        
        QuotaSpec.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>QuotaSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/QuotaSpec} obj Optional instance to populate.
     * @return {module:model/QuotaSpec} The populated <code>QuotaSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new QuotaSpec();

            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = ApiClient.convertToType(data['Description'], 'String');
            }
            if (data.hasOwnProperty('Limits')) {
                obj['Limits'] = ApiClient.convertToType(data['Limits'], [QuotaLimit]);
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
QuotaSpec.prototype['CreateIndex'] = undefined;

/**
 * @member {String} Description
 */
QuotaSpec.prototype['Description'] = undefined;

/**
 * @member {Array.<module:model/QuotaLimit>} Limits
 */
QuotaSpec.prototype['Limits'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
QuotaSpec.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Name
 */
QuotaSpec.prototype['Name'] = undefined;






export default QuotaSpec;

