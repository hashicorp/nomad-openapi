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

import { Resources } from './Resources';
import { HttpFile } from '../http/http';

export class QuotaLimit {
    'hash'?: string;
    'region'?: string;
    'regionLimit'?: Resources;
    'variablesLimit'?: number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "hash",
            "baseName": "Hash",
            "type": "string",
            "format": "byte"
        },
        {
            "name": "region",
            "baseName": "Region",
            "type": "string",
            "format": ""
        },
        {
            "name": "regionLimit",
            "baseName": "RegionLimit",
            "type": "Resources",
            "format": ""
        },
        {
            "name": "variablesLimit",
            "baseName": "VariablesLimit",
            "type": "number",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return QuotaLimit.attributeTypeMap;
    }
    
    public constructor() {
    }
}

