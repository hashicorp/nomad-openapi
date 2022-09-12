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
 * The JobACL model module.
 * @module model/JobACL
 * @version 1.1.4
 */
class JobACL {
    /**
     * Constructs a new <code>JobACL</code>.
     * @alias module:model/JobACL
     */
    constructor() { 
        
        JobACL.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobACL</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobACL} obj Optional instance to populate.
     * @return {module:model/JobACL} The populated <code>JobACL</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobACL();

            if (data.hasOwnProperty('Group')) {
                obj['Group'] = ApiClient.convertToType(data['Group'], 'String');
            }
            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Task')) {
                obj['Task'] = ApiClient.convertToType(data['Task'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} Group
 */
JobACL.prototype['Group'] = undefined;

/**
 * @member {String} JobID
 */
JobACL.prototype['JobID'] = undefined;

/**
 * @member {String} Namespace
 */
JobACL.prototype['Namespace'] = undefined;

/**
 * @member {String} Task
 */
JobACL.prototype['Task'] = undefined;






export default JobACL;

