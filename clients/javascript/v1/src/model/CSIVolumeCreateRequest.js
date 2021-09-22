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
import CSIVolume from './CSIVolume';

/**
 * The CSIVolumeCreateRequest model module.
 * @module model/CSIVolumeCreateRequest
 * @version 1.1.4
 */
class CSIVolumeCreateRequest {
    /**
     * Constructs a new <code>CSIVolumeCreateRequest</code>.
     * @alias module:model/CSIVolumeCreateRequest
     */
    constructor() { 
        
        CSIVolumeCreateRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSIVolumeCreateRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSIVolumeCreateRequest} obj Optional instance to populate.
     * @return {module:model/CSIVolumeCreateRequest} The populated <code>CSIVolumeCreateRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSIVolumeCreateRequest();

            if (data.hasOwnProperty('Namespace')) {
                obj['Namespace'] = ApiClient.convertToType(data['Namespace'], 'String');
            }
            if (data.hasOwnProperty('Region')) {
                obj['Region'] = ApiClient.convertToType(data['Region'], 'String');
            }
            if (data.hasOwnProperty('SecretID')) {
                obj['SecretID'] = ApiClient.convertToType(data['SecretID'], 'String');
            }
            if (data.hasOwnProperty('Volumes')) {
                obj['Volumes'] = ApiClient.convertToType(data['Volumes'], [CSIVolume]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} Namespace
 */
CSIVolumeCreateRequest.prototype['Namespace'] = undefined;

/**
 * @member {String} Region
 */
CSIVolumeCreateRequest.prototype['Region'] = undefined;

/**
 * @member {String} SecretID
 */
CSIVolumeCreateRequest.prototype['SecretID'] = undefined;

/**
 * @member {Array.<module:model/CSIVolume>} Volumes
 */
CSIVolumeCreateRequest.prototype['Volumes'] = undefined;






export default CSIVolumeCreateRequest;
