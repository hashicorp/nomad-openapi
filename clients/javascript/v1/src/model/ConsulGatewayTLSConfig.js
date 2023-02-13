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
 * The ConsulGatewayTLSConfig model module.
 * @module model/ConsulGatewayTLSConfig
 * @version 1.1.4
 */
class ConsulGatewayTLSConfig {
    /**
     * Constructs a new <code>ConsulGatewayTLSConfig</code>.
     * @alias module:model/ConsulGatewayTLSConfig
     */
    constructor() { 
        
        ConsulGatewayTLSConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulGatewayTLSConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulGatewayTLSConfig} obj Optional instance to populate.
     * @return {module:model/ConsulGatewayTLSConfig} The populated <code>ConsulGatewayTLSConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulGatewayTLSConfig();

            if (data.hasOwnProperty('CipherSuites')) {
                obj['CipherSuites'] = ApiClient.convertToType(data['CipherSuites'], ['String']);
            }
            if (data.hasOwnProperty('Enabled')) {
                obj['Enabled'] = ApiClient.convertToType(data['Enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('TLSMaxVersion')) {
                obj['TLSMaxVersion'] = ApiClient.convertToType(data['TLSMaxVersion'], 'String');
            }
            if (data.hasOwnProperty('TLSMinVersion')) {
                obj['TLSMinVersion'] = ApiClient.convertToType(data['TLSMinVersion'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<String>} CipherSuites
 */
ConsulGatewayTLSConfig.prototype['CipherSuites'] = undefined;

/**
 * @member {Boolean} Enabled
 */
ConsulGatewayTLSConfig.prototype['Enabled'] = undefined;

/**
 * @member {String} TLSMaxVersion
 */
ConsulGatewayTLSConfig.prototype['TLSMaxVersion'] = undefined;

/**
 * @member {String} TLSMinVersion
 */
ConsulGatewayTLSConfig.prototype['TLSMinVersion'] = undefined;






export default ConsulGatewayTLSConfig;

