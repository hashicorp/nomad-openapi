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
 * The Vault model module.
 * @module model/Vault
 * @version 1.1.4
 */
class Vault {
    /**
     * Constructs a new <code>Vault</code>.
     * @alias module:model/Vault
     */
    constructor() { 
        
        Vault.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Vault</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Vault} obj Optional instance to populate.
     * @return {module:model/Vault} The populated <code>Vault</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Vault();

            if (data.hasOwnProperty('ChangeMode')) {
                obj['ChangeMode'] = ApiClient.convertToType(data['ChangeMode'], 'String');
            }
            if (data.hasOwnProperty('ChangeSignal')) {
                obj['ChangeSignal'] = ApiClient.convertToType(data['ChangeSignal'], 'String');
            }
            if (data.hasOwnProperty('Env')) {
                obj['Env'] = ApiClient.convertToType(data['Env'], 'Boolean');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Policies')) {
                obj['Policies'] = ApiClient.convertToType(data['Policies'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} ChangeMode
 */
Vault.prototype['ChangeMode'] = undefined;

/**
 * @member {String} ChangeSignal
 */
Vault.prototype['ChangeSignal'] = undefined;

/**
 * @member {Boolean} Env
 */
Vault.prototype['Env'] = undefined;

/**
 * @member {String} Namespace
 */
Vault.prototype['Namespace'] = undefined;

/**
 * @member {Array.<String>} Policies
 */
Vault.prototype['Policies'] = undefined;






export default Vault;

