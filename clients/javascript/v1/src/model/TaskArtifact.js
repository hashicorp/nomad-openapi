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
 * The TaskArtifact model module.
 * @module model/TaskArtifact
 * @version 1.1.4
 */
class TaskArtifact {
    /**
     * Constructs a new <code>TaskArtifact</code>.
     * @alias module:model/TaskArtifact
     */
    constructor() { 
        
        TaskArtifact.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskArtifact</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskArtifact} obj Optional instance to populate.
     * @return {module:model/TaskArtifact} The populated <code>TaskArtifact</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskArtifact();

            if (data.hasOwnProperty('GetterHeaders')) {
                obj['GetterHeaders'] = ApiClient.convertToType(data['GetterHeaders'], {'String': 'String'});
            }
            if (data.hasOwnProperty('GetterMode')) {
                obj['GetterMode'] = ApiClient.convertToType(data['GetterMode'], 'String');
            }
            if (data.hasOwnProperty('GetterOptions')) {
                obj['GetterOptions'] = ApiClient.convertToType(data['GetterOptions'], {'String': 'String'});
            }
            if (data.hasOwnProperty('GetterSource')) {
                obj['GetterSource'] = ApiClient.convertToType(data['GetterSource'], 'String');
            }
            if (data.hasOwnProperty('RelativeDest')) {
                obj['RelativeDest'] = ApiClient.convertToType(data['RelativeDest'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Object.<String, String>} GetterHeaders
 */
TaskArtifact.prototype['GetterHeaders'] = undefined;

/**
 * @member {String} GetterMode
 */
TaskArtifact.prototype['GetterMode'] = undefined;

/**
 * @member {Object.<String, String>} GetterOptions
 */
TaskArtifact.prototype['GetterOptions'] = undefined;

/**
 * @member {String} GetterSource
 */
TaskArtifact.prototype['GetterSource'] = undefined;

/**
 * @member {String} RelativeDest
 */
TaskArtifact.prototype['RelativeDest'] = undefined;






export default TaskArtifact;

