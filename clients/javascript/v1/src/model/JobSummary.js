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
import JobChildrenSummary from './JobChildrenSummary';
import TaskGroupSummary from './TaskGroupSummary';

/**
 * The JobSummary model module.
 * @module model/JobSummary
 * @version 1.1.4
 */
class JobSummary {
    /**
     * Constructs a new <code>JobSummary</code>.
     * @alias module:model/JobSummary
     */
    constructor() { 
        
        JobSummary.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobSummary</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobSummary} obj Optional instance to populate.
     * @return {module:model/JobSummary} The populated <code>JobSummary</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobSummary();

            if (data.hasOwnProperty('Children')) {
                obj['Children'] = JobChildrenSummary.constructFromObject(data['Children']);
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('JobID')) {
                obj['JobID'] = ApiClient.convertToType(data['JobID'], 'String');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Summary')) {
                obj['Summary'] = ApiClient.convertToType(data['Summary'], {'String': TaskGroupSummary});
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/JobChildrenSummary} Children
 */
JobSummary.prototype['Children'] = undefined;

/**
 * @member {Number} CreateIndex
 */
JobSummary.prototype['CreateIndex'] = undefined;

/**
 * @member {String} JobID
 */
JobSummary.prototype['JobID'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
JobSummary.prototype['ModifyIndex'] = undefined;

/**
 * @member {String} Namespace
 */
JobSummary.prototype['Namespace'] = undefined;

/**
 * @member {Object.<String, module:model/TaskGroupSummary>} Summary
 */
JobSummary.prototype['Summary'] = undefined;






export default JobSummary;

