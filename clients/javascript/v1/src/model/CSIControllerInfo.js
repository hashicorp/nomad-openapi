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
 * The CSIControllerInfo model module.
 * @module model/CSIControllerInfo
 * @version 1.1.4
 */
class CSIControllerInfo {
    /**
     * Constructs a new <code>CSIControllerInfo</code>.
     * @alias module:model/CSIControllerInfo
     */
    constructor() { 
        
        CSIControllerInfo.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CSIControllerInfo</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CSIControllerInfo} obj Optional instance to populate.
     * @return {module:model/CSIControllerInfo} The populated <code>CSIControllerInfo</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CSIControllerInfo();

            if (data.hasOwnProperty('SupportsAttachDetach')) {
                obj['SupportsAttachDetach'] = ApiClient.convertToType(data['SupportsAttachDetach'], 'Boolean');
            }
            if (data.hasOwnProperty('SupportsListVolumes')) {
                obj['SupportsListVolumes'] = ApiClient.convertToType(data['SupportsListVolumes'], 'Boolean');
            }
            if (data.hasOwnProperty('SupportsListVolumesAttachedNodes')) {
                obj['SupportsListVolumesAttachedNodes'] = ApiClient.convertToType(data['SupportsListVolumesAttachedNodes'], 'Boolean');
            }
            if (data.hasOwnProperty('SupportsReadOnlyAttach')) {
                obj['SupportsReadOnlyAttach'] = ApiClient.convertToType(data['SupportsReadOnlyAttach'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} SupportsAttachDetach
 */
CSIControllerInfo.prototype['SupportsAttachDetach'] = undefined;

/**
 * @member {Boolean} SupportsListVolumes
 */
CSIControllerInfo.prototype['SupportsListVolumes'] = undefined;

/**
 * @member {Boolean} SupportsListVolumesAttachedNodes
 */
CSIControllerInfo.prototype['SupportsListVolumesAttachedNodes'] = undefined;

/**
 * @member {Boolean} SupportsReadOnlyAttach
 */
CSIControllerInfo.prototype['SupportsReadOnlyAttach'] = undefined;






export default CSIControllerInfo;

