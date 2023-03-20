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
 * The JobStabilityRequest model module.
 * @module model/JobStabilityRequest
 * @version 1.1.4
 */
class JobStabilityRequest {
    /**
     * Constructs a new <code>JobStabilityRequest</code>.
     * @alias module:model/JobStabilityRequest
     */
    constructor() { 
        
        JobStabilityRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobStabilityRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobStabilityRequest} obj Optional instance to populate.
     * @return {module:model/JobStabilityRequest} The populated <code>JobStabilityRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobStabilityRequest();

            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
            }
            if (data.hasOwnProperty('JobVersion')) {
                obj['JobVersion'] = ApiClient.convertToType(data['JobVersion'], 'Number');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Region')) {
                obj['Region'] = ApiClient.convertToType(data['Region'], 'String');
            }
            if (data.hasOwnProperty('SecretID')) {
                obj['SecretID'] = ApiClient.convertToType(data['SecretID'], 'String');
            }
            if (data.hasOwnProperty('Stable')) {
                obj['Stable'] = ApiClient.convertToType(data['Stable'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} JobID
 */
JobStabilityRequest.prototype['JobID'] = undefined;

/**
 * @member {Number} JobVersion
 */
JobStabilityRequest.prototype['JobVersion'] = undefined;

/**
 * @member {String} Namespace
 */
JobStabilityRequest.prototype['Namespace'] = undefined;

/**
 * @member {String} Region
 */
JobStabilityRequest.prototype['Region'] = undefined;

/**
 * @member {String} SecretID
 */
JobStabilityRequest.prototype['SecretID'] = undefined;

/**
 * @member {Boolean} Stable
 */
JobStabilityRequest.prototype['Stable'] = undefined;






export default JobStabilityRequest;

