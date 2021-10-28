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

import { ConsulMeshGateway } from './ConsulMeshGateway';
import { HttpFile } from '../http/http';

export class ConsulUpstream {
    'datacenter'?: string;
    'destinationName'?: string;
    'localBindAddress'?: string;
    'localBindPort'?: number;
    'meshGateway'?: ConsulMeshGateway;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "datacenter",
            "baseName": "Datacenter",
            "type": "string",
            "format": ""
        },
        {
            "name": "destinationName",
            "baseName": "DestinationName",
            "type": "string",
            "format": ""
        },
        {
            "name": "localBindAddress",
            "baseName": "LocalBindAddress",
            "type": "string",
            "format": ""
        },
        {
            "name": "localBindPort",
            "baseName": "LocalBindPort",
            "type": "number",
            "format": ""
        },
        {
            "name": "meshGateway",
            "baseName": "MeshGateway",
            "type": "ConsulMeshGateway",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConsulUpstream.attributeTypeMap;
    }
    
    public constructor() {
    }
}

