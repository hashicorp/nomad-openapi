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
 * The DrainMetadata model module.
 * @module model/DrainMetadata
 * @version 1.1.4
 */
class DrainMetadata {
    /**
     * Constructs a new <code>DrainMetadata</code>.
     * @alias module:model/DrainMetadata
     */
    constructor() { 
        
        DrainMetadata.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DrainMetadata</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DrainMetadata} obj Optional instance to populate.
     * @return {module:model/DrainMetadata} The populated <code>DrainMetadata</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DrainMetadata();

            if (data.hasOwnProperty('AccessorID')) {
                obj['AccessorID'] = ApiClient.convertToType(data['AccessorID'], 'String');
            }
            if (data.hasOwnProperty('Meta')) {
                obj['Meta'] = ApiClient.convertToType(data['Meta'], {'String': 'String'});
            }
            if (data.hasOwnProperty('StartedAt')) {
                obj['StartedAt'] = ApiClient.convertToType(data['StartedAt'], 'Date');
            }
            if (data.hasOwnProperty('Status')) {
                obj['Status'] = ApiClient.convertToType(data['Status'], 'String');
            }
            if (data.hasOwnProperty('UpdatedAt')) {
                obj['UpdatedAt'] = ApiClient.convertToType(data['UpdatedAt'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {String} AccessorID
 */
DrainMetadata.prototype['AccessorID'] = undefined;

/**
 * @member {Object.<String, String>} Meta
 */
DrainMetadata.prototype['Meta'] = undefined;

/**
 * @member {Date} StartedAt
 */
DrainMetadata.prototype['StartedAt'] = undefined;

/**
 * @member {String} Status
 */
DrainMetadata.prototype['Status'] = undefined;

/**
 * @member {Date} UpdatedAt
 */
DrainMetadata.prototype['UpdatedAt'] = undefined;






export default DrainMetadata;

