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

import { FieldDiff } from './FieldDiff';
import { HttpFile } from '../http/http';

export class ObjectDiff {
    'fields'?: Array<FieldDiff>;
    'name'?: string;
    'objects'?: Array<ObjectDiff>;
    'type'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "fields",
            "baseName": "Fields",
            "type": "Array<FieldDiff>",
            "format": ""
        },
        {
            "name": "name",
            "baseName": "Name",
            "type": "string",
            "format": ""
        },
        {
            "name": "objects",
            "baseName": "Objects",
            "type": "Array<ObjectDiff>",
            "format": ""
        },
        {
            "name": "type",
            "baseName": "Type",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ObjectDiff.attributeTypeMap;
    }
    
    public constructor() {
    }
}

