/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

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
import Allocation from './Allocation';
import AllocationListStub from './AllocationListStub';
import CSIMountOptions from './CSIMountOptions';
import CSITopology from './CSITopology';
import CSITopologyRequest from './CSITopologyRequest';
import CSIVolumeCapability from './CSIVolumeCapability';

/**
 * The CSIVolume model module.
 * @module model/CSIVolume
 * @version 1.1.4
 */
class CSIVolume {
    /**
     * Constructs a new <code>CSIVolume</code>.
     * @alias module:model/CSIVolume
     */
    constructor() { 
        
        CSIVolume.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSIVolume</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSIVolume} obj Optional instance to populate.
     * @return {module:model/CSIVolume} The populated <code>CSIVolume</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSIVolume();

            if (data.hasOwnProperty('AccessMode')) {
                obj['AccessMode'] = ApiClient.convertToType(data['AccessMode'], 'String');
            }
            if (data.hasOwnProperty('Allocations')) {
                obj['Allocations'] = ApiClient.convertToType(data['Allocations'], [AllocationListStub]);
            }
            if (data.hasOwnProperty('AttachmentMode')) {
                obj['AttachmentMode'] = ApiClient.convertToType(data['AttachmentMode'], 'String');
            }
            if (data.hasOwnProperty('Capacity')) {
                obj['Capacity'] = ApiClient.convertToType(data['Capacity'], 'Number');
            }
            if (data.hasOwnProperty('CloneID')) {
                obj['CloneID'] = ApiClient.convertToType(data['CloneID'], 'String');
            }
            if (data.hasOwnProperty('Context')) {
                obj['Context'] = ApiClient.convertToType(data['Context'], {'String': 'String'});
            }
            if (data.hasOwnProperty('ControllerRequired')) {
                obj['ControllerRequired'] = ApiClient.convertToType(data['ControllerRequired'], 'Boolean');
            }
            if (data.hasOwnProperty('ControllersExpected')) {
                obj['ControllersExpected'] = ApiClient.convertToType(data['ControllersExpected'], 'Number');
            }
            if (data.hasOwnProperty('ControllersHealthy')) {
                obj['ControllersHealthy'] = ApiClient.convertToType(data['ControllersHealthy'], 'Number');
            }
            if (data.hasOwnProperty('CreateIndex')) {
                obj['CreateIndex'] = ApiClient.convertToType(data['CreateIndex'], 'Number');
            }
            if (data.hasOwnProperty('ExternalID')) {
                obj['ExternalID'] = ApiClient.convertToType(data['ExternalID'], 'String');
            }
            if (data.hasOwnProperty('ID')) {
                obj['ID'] = ApiClient.convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('ModifyIndex')) {
                obj['ModifyIndex'] = ApiClient.convertToType(data['ModifyIndex'], 'Number');
            }
            if (data.hasOwnProperty('MountOptions')) {
                obj['MountOptions'] = CSIMountOptions.constructFromObject(data['MountOptions']);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('NodesExpected')) {
                obj['NodesExpected'] = ApiClient.convertToType(data['NodesExpected'], 'Number');
            }
            if (data.hasOwnProperty('NodesHealthy')) {
                obj['NodesHealthy'] = ApiClient.convertToType(data['NodesHealthy'], 'Number');
            }
            if (data.hasOwnProperty('Parameters')) {
                obj['Parameters'] = ApiClient.convertToType(data['Parameters'], {'String': 'String'});
            }
            if (data.hasOwnProperty('PluginID')) {
                obj['PluginID'] = ApiClient.convertToType(data['PluginID'], 'String');
            }
            if (data.hasOwnProperty('Provider')) {
                obj['Provider'] = ApiClient.convertToType(data['Provider'], 'String');
            }
            if (data.hasOwnProperty('ProviderVersion')) {
                obj['ProviderVersion'] = ApiClient.convertToType(data['ProviderVersion'], 'String');
            }
            if (data.hasOwnProperty('ReadAllocs')) {
                obj['ReadAllocs'] = ApiClient.convertToType(data['ReadAllocs'], {'String': Allocation});
            }
            if (data.hasOwnProperty('RequestedCapabilities')) {
                obj['RequestedCapabilities'] = ApiClient.convertToType(data['RequestedCapabilities'], [CSIVolumeCapability]);
            }
            if (data.hasOwnProperty('RequestedCapacityMax')) {
                obj['RequestedCapacityMax'] = ApiClient.convertToType(data['RequestedCapacityMax'], 'Number');
            }
            if (data.hasOwnProperty('RequestedCapacityMin')) {
                obj['RequestedCapacityMin'] = ApiClient.convertToType(data['RequestedCapacityMin'], 'Number');
            }
            if (data.hasOwnProperty('RequestedTopologies')) {
                obj['RequestedTopologies'] = CSITopologyRequest.constructFromObject(data['RequestedTopologies']);
            }
            if (data.hasOwnProperty('ResourceExhausted')) {
                obj['ResourceExhausted'] = ApiClient.convertToType(data['ResourceExhausted'], 'Date');
            }
            if (data.hasOwnProperty('Schedulable')) {
                obj['Schedulable'] = ApiClient.convertToType(data['Schedulable'], 'Boolean');
            }
            if (data.hasOwnProperty('Secrets')) {
                obj['Secrets'] = ApiClient.convertToType(data['Secrets'], {'String': 'String'});
            }
            if (data.hasOwnProperty('SnapshotID')) {
                obj['SnapshotID'] = ApiClient.convertToType(data['SnapshotID'], 'String');
            }
            if (data.hasOwnProperty('Topologies')) {
                obj['Topologies'] = ApiClient.convertToType(data['Topologies'], [CSITopology]);
            }
            if (data.hasOwnProperty('WriteAllocs')) {
                obj['WriteAllocs'] = ApiClient.convertToType(data['WriteAllocs'], {'String': Allocation});
            }
        }
        return obj;
    }


}

/**
 * @member {String} AccessMode
 */
CSIVolume.prototype['AccessMode'] = undefined;

/**
 * @member {Array.<module:model/AllocationListStub>} Allocations
 */
CSIVolume.prototype['Allocations'] = undefined;

/**
 * @member {String} AttachmentMode
 */
CSIVolume.prototype['AttachmentMode'] = undefined;

/**
 * @member {Number} Capacity
 */
CSIVolume.prototype['Capacity'] = undefined;

/**
 * @member {String} CloneID
 */
CSIVolume.prototype['CloneID'] = undefined;

/**
 * @member {Object.<String, String>} Context
 */
CSIVolume.prototype['Context'] = undefined;

/**
 * @member {Boolean} ControllerRequired
 */
CSIVolume.prototype['ControllerRequired'] = undefined;

/**
 * @member {Number} ControllersExpected
 */
CSIVolume.prototype['ControllersExpected'] = undefined;

/**
 * @member {Number} ControllersHealthy
 */
CSIVolume.prototype['ControllersHealthy'] = undefined;

/**
 * @member {Number} CreateIndex
 */
CSIVolume.prototype['CreateIndex'] = undefined;

/**
 * @member {String} ExternalID
 */
CSIVolume.prototype['ExternalID'] = undefined;

/**
 * @member {String} ID
 */
CSIVolume.prototype['ID'] = undefined;

/**
 * @member {Number} ModifyIndex
 */
CSIVolume.prototype['ModifyIndex'] = undefined;

/**
 * @member {module:model/CSIMountOptions} MountOptions
 */
CSIVolume.prototype['MountOptions'] = undefined;

/**
 * @member {String} Name
 */
CSIVolume.prototype['Name'] = undefined;

/**
 * @member {String} Namespace
 */
CSIVolume.prototype['Namespace'] = undefined;

/**
 * @member {Number} NodesExpected
 */
CSIVolume.prototype['NodesExpected'] = undefined;

/**
 * @member {Number} NodesHealthy
 */
CSIVolume.prototype['NodesHealthy'] = undefined;

/**
 * @member {Object.<String, String>} Parameters
 */
CSIVolume.prototype['Parameters'] = undefined;

/**
 * @member {String} PluginID
 */
CSIVolume.prototype['PluginID'] = undefined;

/**
 * @member {String} Provider
 */
CSIVolume.prototype['Provider'] = undefined;

/**
 * @member {String} ProviderVersion
 */
CSIVolume.prototype['ProviderVersion'] = undefined;

/**
 * @member {Object.<String, module:model/Allocation>} ReadAllocs
 */
CSIVolume.prototype['ReadAllocs'] = undefined;

/**
 * @member {Array.<module:model/CSIVolumeCapability>} RequestedCapabilities
 */
CSIVolume.prototype['RequestedCapabilities'] = undefined;

/**
 * @member {Number} RequestedCapacityMax
 */
CSIVolume.prototype['RequestedCapacityMax'] = undefined;

/**
 * @member {Number} RequestedCapacityMin
 */
CSIVolume.prototype['RequestedCapacityMin'] = undefined;

/**
 * @member {module:model/CSITopologyRequest} RequestedTopologies
 */
CSIVolume.prototype['RequestedTopologies'] = undefined;

/**
 * @member {Date} ResourceExhausted
 */
CSIVolume.prototype['ResourceExhausted'] = undefined;

/**
 * @member {Boolean} Schedulable
 */
CSIVolume.prototype['Schedulable'] = undefined;

/**
 * @member {Object.<String, String>} Secrets
 */
CSIVolume.prototype['Secrets'] = undefined;

/**
 * @member {String} SnapshotID
 */
CSIVolume.prototype['SnapshotID'] = undefined;

/**
 * @member {Array.<module:model/CSITopology>} Topologies
 */
CSIVolume.prototype['Topologies'] = undefined;

/**
 * @member {Object.<String, module:model/Allocation>} WriteAllocs
 */
CSIVolume.prototype['WriteAllocs'] = undefined;






export default CSIVolume;

