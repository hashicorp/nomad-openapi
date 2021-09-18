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
import AllocatedCpuResources from './AllocatedCpuResources';
import AllocatedDeviceResource from './AllocatedDeviceResource';
import AllocatedMemoryResources from './AllocatedMemoryResources';
import NetworkResource from './NetworkResource';

/**
 * The AllocatedTaskResources model module.
 * @module model/AllocatedTaskResources
 * @version 1.1.4
 */
class AllocatedTaskResources {
    /**
     * Constructs a new <code>AllocatedTaskResources</code>.
     * @alias module:model/AllocatedTaskResources
     */
    constructor() { 
        
        AllocatedTaskResources.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AllocatedTaskResources</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AllocatedTaskResources} obj Optional instance to populate.
     * @return {module:model/AllocatedTaskResources} The populated <code>AllocatedTaskResources</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AllocatedTaskResources();

            if (data.hasOwnProperty('Cpu')) {
                obj['Cpu'] = AllocatedCpuResources.constructFromObject(data['Cpu']);
            }
            if (data.hasOwnProperty('Devices')) {
                obj['Devices'] = ApiClient.convertToType(data['Devices'], [AllocatedDeviceResource]);
            }
            if (data.hasOwnProperty('Memory')) {
                obj['Memory'] = AllocatedMemoryResources.constructFromObject(data['Memory']);
            }
            if (data.hasOwnProperty('Networks')) {
                obj['Networks'] = ApiClient.convertToType(data['Networks'], [NetworkResource]);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/AllocatedCpuResources} Cpu
 */
AllocatedTaskResources.prototype['Cpu'] = undefined;

/**
 * @member {Array.<module:model/AllocatedDeviceResource>} Devices
 */
AllocatedTaskResources.prototype['Devices'] = undefined;

/**
 * @member {module:model/AllocatedMemoryResources} Memory
 */
AllocatedTaskResources.prototype['Memory'] = undefined;

/**
 * @member {Array.<module:model/NetworkResource>} Networks
 */
AllocatedTaskResources.prototype['Networks'] = undefined;






export default AllocatedTaskResources;

