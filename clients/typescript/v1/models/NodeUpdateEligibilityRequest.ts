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

export class NodeUpdateEligibilityRequest {
    'eligibility'?: string;
    'nodeID'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "eligibility",
            "baseName": "Eligibility",
            "type": "string",
            "format": ""
        },
        {
            "name": "nodeID",
            "baseName": "NodeID",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return NodeUpdateEligibilityRequest.attributeTypeMap;
    }
    
    public constructor() {
    }
}

