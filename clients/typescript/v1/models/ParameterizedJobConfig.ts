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

export class ParameterizedJobConfig {
    'metaOptional'?: Array<string>;
    'metaRequired'?: Array<string>;
    'payload'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "metaOptional",
            "baseName": "MetaOptional",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "metaRequired",
            "baseName": "MetaRequired",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "payload",
            "baseName": "Payload",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ParameterizedJobConfig.attributeTypeMap;
    }
    
    public constructor() {
    }
}

