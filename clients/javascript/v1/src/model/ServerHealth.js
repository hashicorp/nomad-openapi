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
 * The ServerHealth model module.
 * @module model/ServerHealth
 * @version 1.1.4
 */
class ServerHealth {
    /**
     * Constructs a new <code>ServerHealth</code>.
     * @alias module:model/ServerHealth
     */
    constructor() { 
        
        ServerHealth.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServerHealth</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServerHealth} obj Optional instance to populate.
     * @return {module:model/ServerHealth} The populated <code>ServerHealth</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServerHealth();

            if (data.hasOwnProperty('Address')) {
                obj['Address'] = ApiClient.convertToType(data['Address'], 'String');
            }
            if (data.hasOwnProperty('Healthy')) {
                obj['Healthy'] = ApiClient.convertToType(data['Healthy'], 'Boolean');
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('LastContact')) {
                obj['LastContact'] = ApiClient.convertToType(data['LastContact'], 'Number');
            }
            if (data.hasOwnProperty('LastIndex')) {
                obj['LastIndex'] = ApiClient.convertToType(data['LastIndex'], 'Number');
            }
            if (data.hasOwnProperty('LastTerm')) {
                obj['LastTerm'] = ApiClient.convertToType(data['LastTerm'], 'Number');
            }
            if (data.hasOwnProperty('Leader')) {
                obj['Leader'] = ApiClient.convertToType(data['Leader'], 'Boolean');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('SerfStatus')) {
                obj['SerfStatus'] = ApiClient.convertToType(data['SerfStatus'], 'String');
            }
            if (data.hasOwnProperty('StableSince')) {
                obj['StableSince'] = ApiClient.convertToType(data['StableSince'], 'Date');
            }
            if (data.hasOwnProperty('Version')) {
                obj['Version'] = ApiClient.convertToType(data['Version'], 'String');
            }
            if (data.hasOwnProperty('Voter')) {
                obj['Voter'] = ApiClient.convertToType(data['Voter'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} Address
 */
ServerHealth.prototype['Address'] = undefined;

/**
 * @member {Boolean} Healthy
 */
ServerHealth.prototype['Healthy'] = undefined;

/**
 * @member {String} ID
 */
ServerHealth.prototype['ID'] = undefined;

/**
 * @member {Number} LastContact
 */
ServerHealth.prototype['LastContact'] = undefined;

/**
 * @member {Number} LastIndex
 */
ServerHealth.prototype['LastIndex'] = undefined;

/**
 * @member {Number} LastTerm
 */
ServerHealth.prototype['LastTerm'] = undefined;

/**
 * @member {Boolean} Leader
 */
ServerHealth.prototype['Leader'] = undefined;

/**
 * @member {String} Name
 */
ServerHealth.prototype['Name'] = undefined;

/**
 * @member {String} SerfStatus
 */
ServerHealth.prototype['SerfStatus'] = undefined;

/**
 * @member {Date} StableSince
 */
ServerHealth.prototype['StableSince'] = undefined;

/**
 * @member {String} Version
 */
ServerHealth.prototype['Version'] = undefined;

/**
 * @member {Boolean} Voter
 */
ServerHealth.prototype['Voter'] = undefined;






export default ServerHealth;

