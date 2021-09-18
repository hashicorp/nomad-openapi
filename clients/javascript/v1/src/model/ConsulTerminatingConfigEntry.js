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
import ConsulLinkedService from './ConsulLinkedService';

/**
 * The ConsulTerminatingConfigEntry model module.
 * @module model/ConsulTerminatingConfigEntry
 * @version 1.1.4
 */
class ConsulTerminatingConfigEntry {
    /**
     * Constructs a new <code>ConsulTerminatingConfigEntry</code>.
     * @alias module:model/ConsulTerminatingConfigEntry
     */
    constructor() { 
        
        ConsulTerminatingConfigEntry.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulTerminatingConfigEntry</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulTerminatingConfigEntry} obj Optional instance to populate.
     * @return {module:model/ConsulTerminatingConfigEntry} The populated <code>ConsulTerminatingConfigEntry</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulTerminatingConfigEntry();

            if (data.hasOwnProperty('Services')) {
                obj['Services'] = ApiClient.convertToType(data['Services'], [ConsulLinkedService]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/ConsulLinkedService>} Services
 */
ConsulTerminatingConfigEntry.prototype['Services'] = undefined;






export default ConsulTerminatingConfigEntry;

