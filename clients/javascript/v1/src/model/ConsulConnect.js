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
import ConsulGateway from './ConsulGateway';
import ConsulSidecarService from './ConsulSidecarService';
import SidecarTask from './SidecarTask';

/**
 * The ConsulConnect model module.
 * @module model/ConsulConnect
 * @version 1.1.4
 */
class ConsulConnect {
    /**
     * Constructs a new <code>ConsulConnect</code>.
     * @alias module:model/ConsulConnect
     */
    constructor() { 
        
        ConsulConnect.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulConnect</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulConnect} obj Optional instance to populate.
     * @return {module:model/ConsulConnect} The populated <code>ConsulConnect</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulConnect();

            if (data.hasOwnProperty('Gateway')) {
                obj['Gateway'] = ConsulGateway.constructFromObject(data['Gateway']);
            }
            if (data.hasOwnProperty('Native')) {
                obj['Native'] = ApiClient.convertToType(data['Native'], 'Boolean');
            }
            if (data.hasOwnProperty('SidecarService')) {
                obj['SidecarService'] = ConsulSidecarService.constructFromObject(data['SidecarService']);
            }
            if (data.hasOwnProperty('SidecarTask')) {
                obj['SidecarTask'] = SidecarTask.constructFromObject(data['SidecarTask']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ConsulGateway} Gateway
 */
ConsulConnect.prototype['Gateway'] = undefined;

/**
 * @member {Boolean} Native
 */
ConsulConnect.prototype['Native'] = undefined;

/**
 * @member {module:model/ConsulSidecarService} SidecarService
 */
ConsulConnect.prototype['SidecarService'] = undefined;

/**
 * @member {module:model/SidecarTask} SidecarTask
 */
ConsulConnect.prototype['SidecarTask'] = undefined;






export default ConsulConnect;

