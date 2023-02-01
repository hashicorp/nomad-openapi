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
 * The Attribute model module.
 * @module model/Attribute
 * @version 1.1.4
 */
class Attribute {
    /**
     * Constructs a new <code>Attribute</code>.
     * @alias module:model/Attribute
     */
    constructor() { 
        
        Attribute.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Attribute</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Attribute} obj Optional instance to populate.
     * @return {module:model/Attribute} The populated <code>Attribute</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Attribute();

            if (data.hasOwnProperty('Bool')) {
                obj['Bool'] = ApiClient.convertToType(data['Bool'], 'Boolean');
            }
            if (data.hasOwnProperty('Float')) {
                obj['Float'] = ApiClient.convertToType(data['Float'], 'Number');
            }
            if (data.hasOwnProperty('Int')) {
                obj['Int'] = ApiClient.convertToType(data['Int'], 'Number');
            }
            if (data.hasOwnProperty('String')) {
                obj['String'] = ApiClient.convertToType(data['String'], 'String');
            }
            if (data.hasOwnProperty('Unit')) {
                obj['Unit'] = ApiClient.convertToType(data['Unit'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} Bool
 */
Attribute.prototype['Bool'] = undefined;

/**
 * @member {Number} Float
 */
Attribute.prototype['Float'] = undefined;

/**
 * @member {Number} Int
 */
Attribute.prototype['Int'] = undefined;

/**
 * @member {String} String
 */
Attribute.prototype['String'] = undefined;

/**
 * @member {String} Unit
 */
Attribute.prototype['Unit'] = undefined;






export default Attribute;

