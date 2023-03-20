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
import ConsulGatewayBindAddress from './ConsulGatewayBindAddress';

/**
 * The ConsulGatewayProxy model module.
 * @module model/ConsulGatewayProxy
 * @version 1.1.4
 */
class ConsulGatewayProxy {
    /**
     * Constructs a new <code>ConsulGatewayProxy</code>.
     * @alias module:model/ConsulGatewayProxy
     */
    constructor() { 
        
        ConsulGatewayProxy.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulGatewayProxy</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulGatewayProxy} obj Optional instance to populate.
     * @return {module:model/ConsulGatewayProxy} The populated <code>ConsulGatewayProxy</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulGatewayProxy();

            if (data.hasOwnProperty('Config')) {
                obj['Config'] = ApiClient.convertToType(data['Config'], {'String': Object});
            }
            if (data.hasOwnProperty('ConnectTimeout')) {
                obj['ConnectTimeout'] = ApiClient.convertToType(data['ConnectTimeout'], 'Number');
            }
            if (data.hasOwnProperty('EnvoyDNSDiscoveryType')) {
                obj['EnvoyDNSDiscoveryType'] = ApiClient.convertToType(data['EnvoyDNSDiscoveryType'], 'String');
            }
            if (data.hasOwnProperty('EnvoyGatewayBindAddresses')) {
                obj['EnvoyGatewayBindAddresses'] = ApiClient.convertToType(data['EnvoyGatewayBindAddresses'], {'String': ConsulGatewayBindAddress});
            }
            if (data.hasOwnProperty('EnvoyGatewayBindTaggedAddresses')) {
                obj['EnvoyGatewayBindTaggedAddresses'] = ApiClient.convertToType(data['EnvoyGatewayBindTaggedAddresses'], 'Boolean');
            }
            if (data.hasOwnProperty('EnvoyGatewayNoDefaultBind')) {
                obj['EnvoyGatewayNoDefaultBind'] = ApiClient.convertToType(data['EnvoyGatewayNoDefaultBind'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Object.<String, Object>} Config
 */
ConsulGatewayProxy.prototype['Config'] = undefined;

/**
 * @member {Number} ConnectTimeout
 */
ConsulGatewayProxy.prototype['ConnectTimeout'] = undefined;

/**
 * @member {String} EnvoyDNSDiscoveryType
 */
ConsulGatewayProxy.prototype['EnvoyDNSDiscoveryType'] = undefined;

/**
 * @member {Object.<String, module:model/ConsulGatewayBindAddress>} EnvoyGatewayBindAddresses
 */
ConsulGatewayProxy.prototype['EnvoyGatewayBindAddresses'] = undefined;

/**
 * @member {Boolean} EnvoyGatewayBindTaggedAddresses
 */
ConsulGatewayProxy.prototype['EnvoyGatewayBindTaggedAddresses'] = undefined;

/**
 * @member {Boolean} EnvoyGatewayNoDefaultBind
 */
ConsulGatewayProxy.prototype['EnvoyGatewayNoDefaultBind'] = undefined;






export default ConsulGatewayProxy;

