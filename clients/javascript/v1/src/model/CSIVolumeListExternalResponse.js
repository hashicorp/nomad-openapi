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
import CSIVolumeExternalStub from './CSIVolumeExternalStub';

/**
 * The CSIVolumeListExternalResponse model module.
 * @module model/CSIVolumeListExternalResponse
 * @version 1.1.4
 */
class CSIVolumeListExternalResponse {
    /**
     * Constructs a new <code>CSIVolumeListExternalResponse</code>.
     * @alias module:model/CSIVolumeListExternalResponse
     */
    constructor() { 
        
        CSIVolumeListExternalResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSIVolumeListExternalResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSIVolumeListExternalResponse} obj Optional instance to populate.
     * @return {module:model/CSIVolumeListExternalResponse} The populated <code>CSIVolumeListExternalResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSIVolumeListExternalResponse();

            if (data.hasOwnProperty('NextToken')) {
                obj['NextToken'] = ApiClient.convertToType(data['NextToken'], 'String');
            }
            if (data.hasOwnProperty('Volumes')) {
                obj['Volumes'] = ApiClient.convertToType(data['Volumes'], [CSIVolumeExternalStub]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} NextToken
 */
CSIVolumeListExternalResponse.prototype['NextToken'] = undefined;

/**
 * @member {Array.<module:model/CSIVolumeExternalStub>} Volumes
 */
CSIVolumeListExternalResponse.prototype['Volumes'] = undefined;






export default CSIVolumeListExternalResponse;

