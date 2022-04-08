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

import { CSISnapshot } from './CSISnapshot';
import { HttpFile } from '../http/http';

export class CSISnapshotCreateResponse {
    'knownLeader'?: boolean;
    'lastContact'?: number;
    'lastIndex'?: number;
    'nextToken'?: string;
    'requestTime'?: number;
    'snapshots'?: Array<CSISnapshot>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "knownLeader",
            "baseName": "KnownLeader",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "lastContact",
            "baseName": "LastContact",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "lastIndex",
            "baseName": "LastIndex",
            "type": "number",
            "format": ""
        },
        {
            "name": "nextToken",
            "baseName": "NextToken",
            "type": "string",
            "format": ""
        },
        {
            "name": "requestTime",
            "baseName": "RequestTime",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "snapshots",
            "baseName": "Snapshots",
            "type": "Array<CSISnapshot>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return CSISnapshotCreateResponse.attributeTypeMap;
    }
    
    public constructor() {
    }
}

