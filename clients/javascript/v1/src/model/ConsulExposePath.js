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
 * The ConsulExposePath model module.
 * @module model/ConsulExposePath
 * @version 1.1.4
 */
class ConsulExposePath {
    /**
     * Constructs a new <code>ConsulExposePath</code>.
     * @alias module:model/ConsulExposePath
     */
    constructor() { 
        
        ConsulExposePath.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulExposePath</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulExposePath} obj Optional instance to populate.
     * @return {module:model/ConsulExposePath} The populated <code>ConsulExposePath</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulExposePath();

            if (data.hasOwnProperty('ListenerPort')) {
                obj['ListenerPort'] = ApiClient.convertToType(data['ListenerPort'], 'String');
            }
            if (data.hasOwnProperty('LocalPathPort')) {
                obj['LocalPathPort'] = ApiClient.convertToType(data['LocalPathPort'], 'Number');
            }
            if (data.hasOwnProperty('Path')) {
                obj['Path'] = ApiClient.convertToType(data['Path'], 'String');
            }
            if (data.hasOwnProperty('Protocol')) {
                obj['Protocol'] = ApiClient.convertToType(data['Protocol'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} ListenerPort
 */
ConsulExposePath.prototype['ListenerPort'] = undefined;

/**
 * @member {Number} LocalPathPort
 */
ConsulExposePath.prototype['LocalPathPort'] = undefined;

/**
 * @member {String} Path
 */
ConsulExposePath.prototype['Path'] = undefined;

/**
 * @member {String} Protocol
 */
ConsulExposePath.prototype['Protocol'] = undefined;






export default ConsulExposePath;

