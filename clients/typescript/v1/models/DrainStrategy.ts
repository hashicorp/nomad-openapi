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

export class DrainStrategy {
    'deadline'?: number;
    'forceDeadline'?: Date;
    'ignoreSystemJobs'?: boolean;
    'startedAt'?: Date;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "deadline",
            "baseName": "Deadline",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "forceDeadline",
            "baseName": "ForceDeadline",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "ignoreSystemJobs",
            "baseName": "IgnoreSystemJobs",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "startedAt",
            "baseName": "StartedAt",
            "type": "Date",
            "format": "date-time"
        }    ];

    static getAttributeTypeMap() {
        return DrainStrategy.attributeTypeMap;
    }
    
    public constructor() {
    }
}
