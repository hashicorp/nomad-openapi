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
 * The NodeUpdateEligibilityRequest model module.
 * @module model/NodeUpdateEligibilityRequest
 * @version 1.1.4
 */
class NodeUpdateEligibilityRequest {
    /**
     * Constructs a new <code>NodeUpdateEligibilityRequest</code>.
     * @alias module:model/NodeUpdateEligibilityRequest
     */
    constructor() { 
        
        NodeUpdateEligibilityRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NodeUpdateEligibilityRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NodeUpdateEligibilityRequest} obj Optional instance to populate.
     * @return {module:model/NodeUpdateEligibilityRequest} The populated <code>NodeUpdateEligibilityRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NodeUpdateEligibilityRequest();

            if (data.hasOwnProperty('Eligibility')) {
                obj['Eligibility'] = ApiClient.convertToType(data['Eligibility'], 'String');
            }
            if (data.hasOwnProperty('NodeID')) {
                obj['NodeID'] = ApiClient.convertToType(data['NodeID'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} Eligibility
 */
NodeUpdateEligibilityRequest.prototype['Eligibility'] = undefined;

/**
 * @member {String} NodeID
 */
NodeUpdateEligibilityRequest.prototype['NodeID'] = undefined;






export default NodeUpdateEligibilityRequest;

