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

import { AllocatedCpuResources } from './AllocatedCpuResources';
import { AllocatedDeviceResource } from './AllocatedDeviceResource';
import { AllocatedMemoryResources } from './AllocatedMemoryResources';
import { NetworkResource } from './NetworkResource';
import { HttpFile } from '../http/http';

export class AllocatedTaskResources {
    'cpu'?: AllocatedCpuResources;
    'devices'?: Array<AllocatedDeviceResource>;
    'memory'?: AllocatedMemoryResources;
    'networks'?: Array<NetworkResource>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "cpu",
            "baseName": "Cpu",
            "type": "AllocatedCpuResources",
            "format": ""
        },
        {
            "name": "devices",
            "baseName": "Devices",
            "type": "Array<AllocatedDeviceResource>",
            "format": ""
        },
        {
            "name": "memory",
            "baseName": "Memory",
            "type": "AllocatedMemoryResources",
            "format": ""
        },
        {
            "name": "networks",
            "baseName": "Networks",
            "type": "Array<NetworkResource>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return AllocatedTaskResources.attributeTypeMap;
    }

    public constructor() {
    }
}

