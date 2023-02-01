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
import CheckRestart from './CheckRestart';
import ConsulConnect from './ConsulConnect';
import ServiceCheck from './ServiceCheck';

/**
 * The Service model module.
 * @module model/Service
 * @version 1.1.4
 */
class Service {
    /**
     * Constructs a new <code>Service</code>.
     * @alias module:model/Service
     */
    constructor() { 
        
        Service.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Service</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Service} obj Optional instance to populate.
     * @return {module:model/Service} The populated <code>Service</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Service();

            if (data.hasOwnProperty('Address')) {
                obj['Address'] = ApiClient.convertToType(data['Address'], 'String');
            }
            if (data.hasOwnProperty('AddressMode')) {
                obj['AddressMode'] = ApiClient.convertToType(data['AddressMode'], 'String');
            }
            if (data.hasOwnProperty('CanaryMeta')) {
                obj['CanaryMeta'] = ApiClient.convertToType(data['CanaryMeta'], {'String': 'String'});
            }
            if (data.hasOwnProperty('CanaryTags')) {
                obj['CanaryTags'] = ApiClient.convertToType(data['CanaryTags'], ['String']);
            }
            if (data.hasOwnProperty('CheckRestart')) {
                obj['CheckRestart'] = CheckRestart.constructFromObject(data['CheckRestart']);
            }
            if (data.hasOwnProperty('Checks')) {
                obj['Checks'] = ApiClient.convertToType(data['Checks'], [ServiceCheck]);
            }
            if (data.hasOwnProperty('Connect')) {
                obj['Connect'] = ConsulConnect.constructFromObject(data['Connect']);
            }
            if (data.hasOwnProperty('EnableTagOverride')) {
                obj['EnableTagOverride'] = ApiClient.convertToType(data['EnableTagOverride'], 'Boolean');
            }
            if (data.hasOwnProperty('Meta')) {
                obj['Meta'] = ApiClient.convertToType(data['Meta'], {'String': 'String'});
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('OnUpdate')) {
                obj['OnUpdate'] = ApiClient.convertToType(data['OnUpdate'], 'String');
            }
            if (data.hasOwnProperty('PortLabel')) {
                obj['PortLabel'] = ApiClient.convertToType(data['PortLabel'], 'String');
            }
            if (data.hasOwnProperty('Provider')) {
                obj['Provider'] = ApiClient.convertToType(data['Provider'], 'String');
            }
            if (data.hasOwnProperty('TaggedAddresses')) {
                obj['TaggedAddresses'] = ApiClient.convertToType(data['TaggedAddresses'], {'String': 'String'});
            }
            if (data.hasOwnProperty('Tags')) {
                obj['Tags'] = ApiClient.convertToType(data['Tags'], ['String']);
            }
            if (data.hasOwnProperty('TaskName')) {
                obj['TaskName'] = ApiClient.convertToType(data['TaskName'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} Address
 */
Service.prototype['Address'] = undefined;

/**
 * @member {String} AddressMode
 */
Service.prototype['AddressMode'] = undefined;

/**
 * @member {Object.<String, String>} CanaryMeta
 */
Service.prototype['CanaryMeta'] = undefined;

/**
 * @member {Array.<String>} CanaryTags
 */
Service.prototype['CanaryTags'] = undefined;

/**
 * @member {module:model/CheckRestart} CheckRestart
 */
Service.prototype['CheckRestart'] = undefined;

/**
 * @member {Array.<module:model/ServiceCheck>} Checks
 */
Service.prototype['Checks'] = undefined;

/**
 * @member {module:model/ConsulConnect} Connect
 */
Service.prototype['Connect'] = undefined;

/**
 * @member {Boolean} EnableTagOverride
 */
Service.prototype['EnableTagOverride'] = undefined;

/**
 * @member {Object.<String, String>} Meta
 */
Service.prototype['Meta'] = undefined;

/**
 * @member {String} Name
 */
Service.prototype['Name'] = undefined;

/**
 * @member {String} OnUpdate
 */
Service.prototype['OnUpdate'] = undefined;

/**
 * @member {String} PortLabel
 */
Service.prototype['PortLabel'] = undefined;

/**
 * @member {String} Provider
 */
Service.prototype['Provider'] = undefined;

/**
 * @member {Object.<String, String>} TaggedAddresses
 */
Service.prototype['TaggedAddresses'] = undefined;

/**
 * @member {Array.<String>} Tags
 */
Service.prototype['Tags'] = undefined;

/**
 * @member {String} TaskName
 */
Service.prototype['TaskName'] = undefined;






export default Service;

