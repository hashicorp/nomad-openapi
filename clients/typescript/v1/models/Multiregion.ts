/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { MultiregionRegion } from './MultiregionRegion';
import { MultiregionStrategy } from './MultiregionStrategy';
import { HttpFile } from '../http/http';

export class Multiregion {
    'regions'?: Array<MultiregionRegion>;
    'strategy'?: MultiregionStrategy;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "regions",
            "baseName": "Regions",
            "type": "Array<MultiregionRegion>",
            "format": ""
        },
        {
            "name": "strategy",
            "baseName": "Strategy",
            "type": "MultiregionStrategy",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Multiregion.attributeTypeMap;
    }
    
    public constructor() {
    }
}

