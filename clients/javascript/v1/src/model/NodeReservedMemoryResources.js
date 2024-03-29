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
 * The NodeReservedMemoryResources model module.
 * @module model/NodeReservedMemoryResources
 * @version 1.1.4
 */
class NodeReservedMemoryResources {
    /**
     * Constructs a new <code>NodeReservedMemoryResources</code>.
     * @alias module:model/NodeReservedMemoryResources
     */
    constructor() { 
        
        NodeReservedMemoryResources.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NodeReservedMemoryResources</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NodeReservedMemoryResources} obj Optional instance to populate.
     * @return {module:model/NodeReservedMemoryResources} The populated <code>NodeReservedMemoryResources</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NodeReservedMemoryResources();

            if (data.hasOwnProperty('MemoryMB')) {
                obj['MemoryMB'] = ApiClient.convertToType(data['MemoryMB'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} MemoryMB
 */
NodeReservedMemoryResources.prototype['MemoryMB'] = undefined;






export default NodeReservedMemoryResources;

