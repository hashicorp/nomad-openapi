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

import { CSIMountOptions } from './CSIMountOptions';
import { HttpFile } from '../http/http';

export class VolumeRequest {
    'accessMode'?: string;
    'attachmentMode'?: string;
    'mountOptions'?: CSIMountOptions;
    'name'?: string;
    'perAlloc'?: boolean;
    'readOnly'?: boolean;
    'source'?: string;
    'type'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "accessMode",
            "baseName": "AccessMode",
            "type": "string",
            "format": ""
        },
        {
            "name": "attachmentMode",
            "baseName": "AttachmentMode",
            "type": "string",
            "format": ""
        },
        {
            "name": "mountOptions",
            "baseName": "MountOptions",
            "type": "CSIMountOptions",
            "format": ""
        },
        {
            "name": "name",
            "baseName": "Name",
            "type": "string",
            "format": ""
        },
        {
            "name": "perAlloc",
            "baseName": "PerAlloc",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "readOnly",
            "baseName": "ReadOnly",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "source",
            "baseName": "Source",
            "type": "string",
            "format": ""
        },
        {
            "name": "type",
            "baseName": "Type",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return VolumeRequest.attributeTypeMap;
    }

    public constructor() {
    }
}

