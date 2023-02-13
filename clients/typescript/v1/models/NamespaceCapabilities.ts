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

export class NamespaceCapabilities {
    'disabledTaskDrivers'?: Array<string>;
    'enabledTaskDrivers'?: Array<string>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "disabledTaskDrivers",
            "baseName": "DisabledTaskDrivers",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "enabledTaskDrivers",
            "baseName": "EnabledTaskDrivers",
            "type": "Array<string>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return NamespaceCapabilities.attributeTypeMap;
    }
    
    public constructor() {
    }
}

