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
import EvalOptions from './EvalOptions';

/**
 * The JobEvaluateRequest model module.
 * @module model/JobEvaluateRequest
 * @version 1.1.4
 */
class JobEvaluateRequest {
    /**
     * Constructs a new <code>JobEvaluateRequest</code>.
     * @alias module:model/JobEvaluateRequest
     */
    constructor() { 
        
        JobEvaluateRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobEvaluateRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobEvaluateRequest} obj Optional instance to populate.
     * @return {module:model/JobEvaluateRequest} The populated <code>JobEvaluateRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobEvaluateRequest();

            if (data.hasOwnProperty('EvalOptions')) {
                obj['EvalOptions'] = EvalOptions.constructFromObject(data['EvalOptions']);
            }
            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
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
        }
        return obj;
    }


}

/**
 * @member {module:model/EvalOptions} EvalOptions
 */
JobEvaluateRequest.prototype['EvalOptions'] = undefined;

/**
 * @member {String} JobID
 */
JobEvaluateRequest.prototype['JobID'] = undefined;

/**
 * @member {String} Namespace
 */
JobEvaluateRequest.prototype['Namespace'] = undefined;

/**
 * @member {String} Region
 */
JobEvaluateRequest.prototype['Region'] = undefined;

/**
 * @member {String} SecretID
 */
JobEvaluateRequest.prototype['SecretID'] = undefined;






export default JobEvaluateRequest;

