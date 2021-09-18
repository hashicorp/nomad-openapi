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
import Affinity from './Affinity';
import Constraint from './Constraint';
import Consul from './Consul';
import EphemeralDisk from './EphemeralDisk';
import MigrateStrategy from './MigrateStrategy';
import NetworkResource from './NetworkResource';
import ReschedulePolicy from './ReschedulePolicy';
import RestartPolicy from './RestartPolicy';
import ScalingPolicy from './ScalingPolicy';
import Service from './Service';
import Spread from './Spread';
import Task from './Task';
import UpdateStrategy from './UpdateStrategy';
import VolumeRequest from './VolumeRequest';

/**
 * The TaskGroup model module.
 * @module model/TaskGroup
 * @version 1.1.4
 */
class TaskGroup {
    /**
     * Constructs a new <code>TaskGroup</code>.
     * @alias module:model/TaskGroup
     */
    constructor() { 
        
        TaskGroup.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TaskGroup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TaskGroup} obj Optional instance to populate.
     * @return {module:model/TaskGroup} The populated <code>TaskGroup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TaskGroup();

            if (data.hasOwnProperty('Affinities')) {
                obj['Affinities'] = ApiClient.convertToType(data['Affinities'], [Affinity]);
            }
            if (data.hasOwnProperty('Constraints')) {
                obj['Constraints'] = ApiClient.convertToType(data['Constraints'], [Constraint]);
            }
            if (data.hasOwnProperty('Consul')) {
                obj['Consul'] = Consul.constructFromObject(data['Consul']);
            }
            if (data.hasOwnProperty('Count')) {
                obj['Count'] = ApiClient.convertToType(data['Count'], 'Number');
            }
            if (data.hasOwnProperty('EphemeralDisk')) {
                obj['EphemeralDisk'] = EphemeralDisk.constructFromObject(data['EphemeralDisk']);
            }
            if (data.hasOwnProperty('Meta')) {
                obj['Meta'] = ApiClient.convertToType(data['Meta'], {'String': 'String'});
            }
            if (data.hasOwnProperty('Migrate')) {
                obj['Migrate'] = MigrateStrategy.constructFromObject(data['Migrate']);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Networks')) {
                obj['Networks'] = ApiClient.convertToType(data['Networks'], [NetworkResource]);
            }
            if (data.hasOwnProperty('ReschedulePolicy')) {
                obj['ReschedulePolicy'] = ReschedulePolicy.constructFromObject(data['ReschedulePolicy']);
            }
            if (data.hasOwnProperty('RestartPolicy')) {
                obj['RestartPolicy'] = RestartPolicy.constructFromObject(data['RestartPolicy']);
            }
            if (data.hasOwnProperty('Scaling')) {
                obj['Scaling'] = ScalingPolicy.constructFromObject(data['Scaling']);
            }
            if (data.hasOwnProperty('Services')) {
                obj['Services'] = ApiClient.convertToType(data['Services'], [Service]);
            }
            if (data.hasOwnProperty('ShutdownDelay')) {
                obj['ShutdownDelay'] = ApiClient.convertToType(data['ShutdownDelay'], 'Number');
            }
            if (data.hasOwnProperty('Spreads')) {
                obj['Spreads'] = ApiClient.convertToType(data['Spreads'], [Spread]);
            }
            if (data.hasOwnProperty('StopAfterClientDisconnect')) {
                obj['StopAfterClientDisconnect'] = ApiClient.convertToType(data['StopAfterClientDisconnect'], 'Number');
            }
            if (data.hasOwnProperty('Tasks')) {
                obj['Tasks'] = ApiClient.convertToType(data['Tasks'], [Task]);
            }
            if (data.hasOwnProperty('Update')) {
                obj['Update'] = UpdateStrategy.constructFromObject(data['Update']);
            }
            if (data.hasOwnProperty('Volumes')) {
                obj['Volumes'] = ApiClient.convertToType(data['Volumes'], {'String': VolumeRequest});
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/Affinity>} Affinities
 */
TaskGroup.prototype['Affinities'] = undefined;

/**
 * @member {Array.<module:model/Constraint>} Constraints
 */
TaskGroup.prototype['Constraints'] = undefined;

/**
 * @member {module:model/Consul} Consul
 */
TaskGroup.prototype['Consul'] = undefined;

/**
 * @member {Number} Count
 */
TaskGroup.prototype['Count'] = undefined;

/**
 * @member {module:model/EphemeralDisk} EphemeralDisk
 */
TaskGroup.prototype['EphemeralDisk'] = undefined;

/**
 * @member {Object.<String, String>} Meta
 */
TaskGroup.prototype['Meta'] = undefined;

/**
 * @member {module:model/MigrateStrategy} Migrate
 */
TaskGroup.prototype['Migrate'] = undefined;

/**
 * @member {String} Name
 */
TaskGroup.prototype['Name'] = undefined;

/**
 * @member {Array.<module:model/NetworkResource>} Networks
 */
TaskGroup.prototype['Networks'] = undefined;

/**
 * @member {module:model/ReschedulePolicy} ReschedulePolicy
 */
TaskGroup.prototype['ReschedulePolicy'] = undefined;

/**
 * @member {module:model/RestartPolicy} RestartPolicy
 */
TaskGroup.prototype['RestartPolicy'] = undefined;

/**
 * @member {module:model/ScalingPolicy} Scaling
 */
TaskGroup.prototype['Scaling'] = undefined;

/**
 * @member {Array.<module:model/Service>} Services
 */
TaskGroup.prototype['Services'] = undefined;

/**
 * @member {Number} ShutdownDelay
 */
TaskGroup.prototype['ShutdownDelay'] = undefined;

/**
 * @member {Array.<module:model/Spread>} Spreads
 */
TaskGroup.prototype['Spreads'] = undefined;

/**
 * @member {Number} StopAfterClientDisconnect
 */
TaskGroup.prototype['StopAfterClientDisconnect'] = undefined;

/**
 * @member {Array.<module:model/Task>} Tasks
 */
TaskGroup.prototype['Tasks'] = undefined;

/**
 * @member {module:model/UpdateStrategy} Update
 */
TaskGroup.prototype['Update'] = undefined;

/**
 * @member {Object.<String, module:model/VolumeRequest>} Volumes
 */
TaskGroup.prototype['Volumes'] = undefined;






export default TaskGroup;

