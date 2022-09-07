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

import { CSIVolume } from './CSIVolume';
import { HttpFile } from '../http/http';

export class CSIVolumeCreateRequest {
    'namespace'?: string;
    'region'?: string;
    'secretID'?: string;
    'volumes'?: Array<CSIVolume>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "namespace",
            "baseName": "Namespace",
            "type": "string",
            "format": ""
        },
        {
            "name": "region",
            "baseName": "Region",
            "type": "string",
            "format": ""
        },
        {
            "name": "secretID",
            "baseName": "SecretID",
            "type": "string",
            "format": ""
        },
        {
            "name": "volumes",
            "baseName": "Volumes",
            "type": "Array<CSIVolume>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return CSIVolumeCreateRequest.attributeTypeMap;
    }

    public constructor() {
    }
}

