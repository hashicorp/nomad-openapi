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
import CSIMountOptions from './CSIMountOptions';

/**
 * The VolumeRequest model module.
 * @module model/VolumeRequest
 * @version 1.1.4
 */
class VolumeRequest {
    /**
     * Constructs a new <code>VolumeRequest</code>.
     * @alias module:model/VolumeRequest
     */
    constructor() { 
        
        VolumeRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VolumeRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VolumeRequest} obj Optional instance to populate.
     * @return {module:model/VolumeRequest} The populated <code>VolumeRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VolumeRequest();

            if (data.hasOwnProperty('AccessMode')) {
                obj['AccessMode'] = ApiClient.convertToType(data['AccessMode'], 'String');
            }
            if (data.hasOwnProperty('AttachmentMode')) {
                obj['AttachmentMode'] = ApiClient.convertToType(data['AttachmentMode'], 'String');
            }
            if (data.hasOwnProperty('MountOptions')) {
                obj['MountOptions'] = CSIMountOptions.constructFromObject(data['MountOptions']);
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = ApiClient.convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('PerAlloc')) {
                obj['PerAlloc'] = ApiClient.convertToType(data['PerAlloc'], 'Boolean');
            }
            if (data.hasOwnProperty('ReadOnly')) {
                obj['ReadOnly'] = ApiClient.convertToType(data['ReadOnly'], 'Boolean');
            }
            if (data.hasOwnProperty('Source')) {
                obj['Source'] = ApiClient.convertToType(data['Source'], 'String');
            }
            if (data.hasOwnProperty('Type')) {
                obj['Type'] = ApiClient.convertToType(data['Type'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} AccessMode
 */
VolumeRequest.prototype['AccessMode'] = undefined;

/**
 * @member {String} AttachmentMode
 */
VolumeRequest.prototype['AttachmentMode'] = undefined;

/**
 * @member {module:model/CSIMountOptions} MountOptions
 */
VolumeRequest.prototype['MountOptions'] = undefined;

/**
 * @member {String} Name
 */
VolumeRequest.prototype['Name'] = undefined;

/**
 * @member {Boolean} PerAlloc
 */
VolumeRequest.prototype['PerAlloc'] = undefined;

/**
 * @member {Boolean} ReadOnly
 */
VolumeRequest.prototype['ReadOnly'] = undefined;

/**
 * @member {String} Source
 */
VolumeRequest.prototype['Source'] = undefined;

/**
 * @member {String} Type
 */
VolumeRequest.prototype['Type'] = undefined;






export default VolumeRequest;

