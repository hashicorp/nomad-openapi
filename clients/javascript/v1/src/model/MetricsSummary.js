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
import GaugeValue from './GaugeValue';
import PointValue from './PointValue';
import SampledValue from './SampledValue';

/**
 * The MetricsSummary model module.
 * @module model/MetricsSummary
 * @version 1.1.4
 */
class MetricsSummary {
    /**
     * Constructs a new <code>MetricsSummary</code>.
     * @alias module:model/MetricsSummary
     */
    constructor() { 
        
        MetricsSummary.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>MetricsSummary</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MetricsSummary} obj Optional instance to populate.
     * @return {module:model/MetricsSummary} The populated <code>MetricsSummary</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MetricsSummary();

            if (data.hasOwnProperty('Counters')) {
                obj['Counters'] = ApiClient.convertToType(data['Counters'], [SampledValue]);
            }
            if (data.hasOwnProperty('Gauges')) {
                obj['Gauges'] = ApiClient.convertToType(data['Gauges'], [GaugeValue]);
            }
            if (data.hasOwnProperty('Points')) {
                obj['Points'] = ApiClient.convertToType(data['Points'], [PointValue]);
            }
            if (data.hasOwnProperty('Samples')) {
                obj['Samples'] = ApiClient.convertToType(data['Samples'], [SampledValue]);
            }
            if (data.hasOwnProperty('Timestamp')) {
                obj['Timestamp'] = ApiClient.convertToType(data['Timestamp'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/SampledValue>} Counters
 */
MetricsSummary.prototype['Counters'] = undefined;

/**
 * @member {Array.<module:model/GaugeValue>} Gauges
 */
MetricsSummary.prototype['Gauges'] = undefined;

/**
 * @member {Array.<module:model/PointValue>} Points
 */
MetricsSummary.prototype['Points'] = undefined;

/**
 * @member {Array.<module:model/SampledValue>} Samples
 */
MetricsSummary.prototype['Samples'] = undefined;

/**
 * @member {String} Timestamp
 */
MetricsSummary.prototype['Timestamp'] = undefined;






export default MetricsSummary;

