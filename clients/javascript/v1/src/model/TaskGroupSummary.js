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
 * The TaskGroupSummary model module.
 * @module model/TaskGroupSummary
 * @version 1.1.4
 */
class TaskGroupSummary {
    /**
     * Constructs a new <code>TaskGroupSummary</code>.
     * @alias module:model/TaskGroupSummary
     */
    constructor() { 
        
        TaskGroupSummary.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskGroupSummary</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskGroupSummary} obj Optional instance to populate.
     * @return {module:model/TaskGroupSummary} The populated <code>TaskGroupSummary</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskGroupSummary();

            if (data.hasOwnProperty('Complete')) {
                obj['Complete'] = ApiClient.convertToType(data['Complete'], 'Number');
            }
            if (data.hasOwnProperty('Failed')) {
                obj['Failed'] = ApiClient.convertToType(data['Failed'], 'Number');
            }
            if (data.hasOwnProperty('Lost')) {
                obj['Lost'] = ApiClient.convertToType(data['Lost'], 'Number');
            }
            if (data.hasOwnProperty('Queued')) {
                obj['Queued'] = ApiClient.convertToType(data['Queued'], 'Number');
            }
            if (data.hasOwnProperty('Running')) {
                obj['Running'] = ApiClient.convertToType(data['Running'], 'Number');
            }
            if (data.hasOwnProperty('Starting')) {
                obj['Starting'] = ApiClient.convertToType(data['Starting'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} Complete
 */
TaskGroupSummary.prototype['Complete'] = undefined;

/**
 * @member {Number} Failed
 */
TaskGroupSummary.prototype['Failed'] = undefined;

/**
 * @member {Number} Lost
 */
TaskGroupSummary.prototype['Lost'] = undefined;

/**
 * @member {Number} Queued
 */
TaskGroupSummary.prototype['Queued'] = undefined;

/**
 * @member {Number} Running
 */
TaskGroupSummary.prototype['Running'] = undefined;

/**
 * @member {Number} Starting
 */
TaskGroupSummary.prototype['Starting'] = undefined;






export default TaskGroupSummary;

