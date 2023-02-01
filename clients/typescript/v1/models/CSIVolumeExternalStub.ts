/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

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

export class CSIVolumeExternalStub {
    'capacityBytes'?: number;
    'cloneID'?: string;
    'externalID'?: string;
    'isAbnormal'?: boolean;
    'publishedExternalNodeIDs'?: Array<string>;
    'snapshotID'?: string;
    'status'?: string;
    'volumeContext'?: { [key: string]: string; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "capacityBytes",
            "baseName": "CapacityBytes",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "cloneID",
            "baseName": "CloneID",
            "type": "string",
            "format": ""
        },
        {
            "name": "externalID",
            "baseName": "ExternalID",
            "type": "string",
            "format": ""
        },
        {
            "name": "isAbnormal",
            "baseName": "IsAbnormal",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "publishedExternalNodeIDs",
            "baseName": "PublishedExternalNodeIDs",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "snapshotID",
            "baseName": "SnapshotID",
            "type": "string",
            "format": ""
        },
        {
            "name": "status",
            "baseName": "Status",
            "type": "string",
            "format": ""
        },
        {
            "name": "volumeContext",
            "baseName": "VolumeContext",
            "type": "{ [key: string]: string; }",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return CSIVolumeExternalStub.attributeTypeMap;
    }
    
    public constructor() {
    }
}

