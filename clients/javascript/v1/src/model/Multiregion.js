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
import MultiregionRegion from './MultiregionRegion';
import MultiregionStrategy from './MultiregionStrategy';

/**
 * The Multiregion model module.
 * @module model/Multiregion
 * @version 1.1.4
 */
class Multiregion {
    /**
     * Constructs a new <code>Multiregion</code>.
     * @alias module:model/Multiregion
     */
    constructor() { 
        
        Multiregion.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Multiregion</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Multiregion} obj Optional instance to populate.
     * @return {module:model/Multiregion} The populated <code>Multiregion</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Multiregion();

            if (data.hasOwnProperty('Regions')) {
                obj['Regions'] = ApiClient.convertToType(data['Regions'], [MultiregionRegion]);
            }
            if (data.hasOwnProperty('Strategy')) {
                obj['Strategy'] = MultiregionStrategy.constructFromObject(data['Strategy']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/MultiregionRegion>} Regions
 */
Multiregion.prototype['Regions'] = undefined;

/**
 * @member {module:model/MultiregionStrategy} Strategy
 */
Multiregion.prototype['Strategy'] = undefined;






export default Multiregion;

