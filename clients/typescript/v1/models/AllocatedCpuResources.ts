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

import { HttpFile } from '../http/http';

export class AllocatedCpuResources {
    'cpuShares'?: number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "cpuShares",
            "baseName": "CpuShares",
            "type": "number",
            "format": "int64"
        }    ];

    static getAttributeTypeMap() {
        return AllocatedCpuResources.attributeTypeMap;
    }
    
    public constructor() {
    }
}

