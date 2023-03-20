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

export class DriverInfo {
    'attributes'?: { [key: string]: string; };
    'detected'?: boolean;
    'healthDescription'?: string;
    'healthy'?: boolean;
    'updateTime'?: Date;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "attributes",
            "baseName": "Attributes",
            "type": "{ [key: string]: string; }",
            "format": ""
        },
        {
            "name": "detected",
            "baseName": "Detected",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "healthDescription",
            "baseName": "HealthDescription",
            "type": "string",
            "format": ""
        },
        {
            "name": "healthy",
            "baseName": "Healthy",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "updateTime",
            "baseName": "UpdateTime",
            "type": "Date",
            "format": "date-time"
        }    ];

    static getAttributeTypeMap() {
        return DriverInfo.attributeTypeMap;
    }
    
    public constructor() {
    }
}

