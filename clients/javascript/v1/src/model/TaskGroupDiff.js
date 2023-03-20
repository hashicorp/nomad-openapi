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
import FieldDiff from './FieldDiff';
import ObjectDiff from './ObjectDiff';
import TaskDiff from './TaskDiff';

/**
 * The TaskGroupDiff model module.
 * @module model/TaskGroupDiff
 * @version 1.1.4
 */
class TaskGroupDiff {
    /**
     * Constructs a new <code>TaskGroupDiff</code>.
     * @alias module:model/TaskGroupDiff
     */
    constructor() { 
        
        TaskGroupDiff.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskGroupDiff</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskGroupDiff} obj Optional instance to populate.
     * @return {module:model/TaskGroupDiff} The populated <code>TaskGroupDiff</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskGroupDiff();

            if (data.hasOwnProperty('Fields')) {
                obj['Fields'] = ApiClient.convertToType(data['Fields'], [FieldDiff]);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Objects')) {
                obj['Objects'] = ApiClient.convertToType(data['Objects'], [ObjectDiff]);
            }
            if (data.hasOwnProperty('Tasks')) {
                obj['Tasks'] = ApiClient.convertToType(data['Tasks'], [TaskDiff]);
            }
            if (data.hasOwnProperty('Type')) {
                obj['Type'] = ApiClient.convertToType(data['Type'], 'String');
            }
            if (data.hasOwnProperty('Updates')) {
                obj['Updates'] = ApiClient.convertToType(data['Updates'], {'String': 'Number'});
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/FieldDiff>} Fields
 */
TaskGroupDiff.prototype['Fields'] = undefined;

/**
 * @member {String} Name
 */
TaskGroupDiff.prototype['Name'] = undefined;

/**
 * @member {Array.<module:model/ObjectDiff>} Objects
 */
TaskGroupDiff.prototype['Objects'] = undefined;

/**
 * @member {Array.<module:model/TaskDiff>} Tasks
 */
TaskGroupDiff.prototype['Tasks'] = undefined;

/**
 * @member {String} Type
 */
TaskGroupDiff.prototype['Type'] = undefined;

/**
 * @member {Object.<String, Number>} Updates
 */
TaskGroupDiff.prototype['Updates'] = undefined;






export default TaskGroupDiff;

