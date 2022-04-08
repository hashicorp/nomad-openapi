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

import { PreemptionConfig } from './PreemptionConfig';
import { HttpFile } from '../http/http';

export class SchedulerConfiguration {
    'createIndex'?: number;
    'memoryOversubscriptionEnabled'?: boolean;
    'modifyIndex'?: number;
    'preemptionConfig'?: PreemptionConfig;
    'rejectJobRegistration'?: boolean;
    'schedulerAlgorithm'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "createIndex",
            "baseName": "CreateIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "memoryOversubscriptionEnabled",
            "baseName": "MemoryOversubscriptionEnabled",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "modifyIndex",
            "baseName": "ModifyIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "preemptionConfig",
            "baseName": "PreemptionConfig",
            "type": "PreemptionConfig",
            "format": ""
        },
        {
            "name": "rejectJobRegistration",
            "baseName": "RejectJobRegistration",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "schedulerAlgorithm",
            "baseName": "SchedulerAlgorithm",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return SchedulerConfiguration.attributeTypeMap;
    }
    
    public constructor() {
    }
}

