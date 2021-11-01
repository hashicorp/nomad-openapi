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

export class Attribute {
    'bool'?: boolean;
    '_float'?: number;
    '_int'?: number;
    'string'?: string;
    'unit'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "bool",
            "baseName": "Bool",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "_float",
            "baseName": "Float",
            "type": "number",
            "format": "double"
        },
        {
            "name": "_int",
            "baseName": "Int",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "string",
            "baseName": "String",
            "type": "string",
            "format": ""
        },
        {
            "name": "unit",
            "baseName": "Unit",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Attribute.attributeTypeMap;
    }
    
    public constructor() {
    }
}

