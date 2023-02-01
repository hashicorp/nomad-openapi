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
import CSISnapshot from './CSISnapshot';

/**
 * The CSISnapshotListResponse model module.
 * @module model/CSISnapshotListResponse
 * @version 1.1.4
 */
class CSISnapshotListResponse {
    /**
     * Constructs a new <code>CSISnapshotListResponse</code>.
     * @alias module:model/CSISnapshotListResponse
     */
    constructor() { 
        
        CSISnapshotListResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSISnapshotListResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSISnapshotListResponse} obj Optional instance to populate.
     * @return {module:model/CSISnapshotListResponse} The populated <code>CSISnapshotListResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSISnapshotListResponse();

            if (data.hasOwnProperty('KnownLeader')) {
                obj['KnownLeader'] = ApiClient.convertToType(data['KnownLeader'], 'Boolean');
            }
            if (data.hasOwnProperty('LastContact')) {
                obj['LastContact'] = ApiClient.convertToType(data['LastContact'], 'Number');
            }
            if (data.hasOwnProperty('LastIndex')) {
                obj['LastIndex'] = ApiClient.convertToType(data['LastIndex'], 'Number');
            }
            if (data.hasOwnProperty('NextToken')) {
                obj['NextToken'] = ApiClient.convertToType(data['NextToken'], 'String');
            }
            if (data.hasOwnProperty('RequestTime')) {
                obj['RequestTime'] = ApiClient.convertToType(data['RequestTime'], 'Number');
            }
            if (data.hasOwnProperty('Snapshots')) {
                obj['Snapshots'] = ApiClient.convertToType(data['Snapshots'], [CSISnapshot]);
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} KnownLeader
 */
CSISnapshotListResponse.prototype['KnownLeader'] = undefined;

/**
 * @member {Number} LastContact
 */
CSISnapshotListResponse.prototype['LastContact'] = undefined;

/**
 * @member {Number} LastIndex
 */
CSISnapshotListResponse.prototype['LastIndex'] = undefined;

/**
 * @member {String} NextToken
 */
CSISnapshotListResponse.prototype['NextToken'] = undefined;

/**
 * @member {Number} RequestTime
 */
CSISnapshotListResponse.prototype['RequestTime'] = undefined;

/**
 * @member {Array.<module:model/CSISnapshot>} Snapshots
 */
CSISnapshotListResponse.prototype['Snapshots'] = undefined;






export default CSISnapshotListResponse;

