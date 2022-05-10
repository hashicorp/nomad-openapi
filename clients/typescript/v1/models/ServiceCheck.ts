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

import { CheckRestart } from './CheckRestart';
import { HttpFile } from '../http/http';

export class ServiceCheck {
    'addressMode'?: string;
    'advertise'?: string;
    'args'?: Array<string>;
    'body'?: string;
    'checkRestart'?: CheckRestart;
    'command'?: string;
    'expose'?: boolean;
    'failuresBeforeCritical'?: number;
    'gRPCService'?: string;
    'gRPCUseTLS'?: boolean;
    'header'?: { [key: string]: Array<string>; };
    'initialStatus'?: string;
    'interval'?: number;
    'method'?: string;
    'name'?: string;
    'onUpdate'?: string;
    'path'?: string;
    'portLabel'?: string;
    'protocol'?: string;
    'successBeforePassing'?: number;
    'tLSSkipVerify'?: boolean;
    'taskName'?: string;
    'timeout'?: number;
    'type'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "addressMode",
            "baseName": "AddressMode",
            "type": "string",
            "format": ""
        },
        {
            "name": "advertise",
            "baseName": "Advertise",
            "type": "string",
            "format": ""
        },
        {
            "name": "args",
            "baseName": "Args",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "body",
            "baseName": "Body",
            "type": "string",
            "format": ""
        },
        {
            "name": "checkRestart",
            "baseName": "CheckRestart",
            "type": "CheckRestart",
            "format": ""
        },
        {
            "name": "command",
            "baseName": "Command",
            "type": "string",
            "format": ""
        },
        {
            "name": "expose",
            "baseName": "Expose",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "failuresBeforeCritical",
            "baseName": "FailuresBeforeCritical",
            "type": "number",
            "format": ""
        },
        {
            "name": "gRPCService",
            "baseName": "GRPCService",
            "type": "string",
            "format": ""
        },
        {
            "name": "gRPCUseTLS",
            "baseName": "GRPCUseTLS",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "header",
            "baseName": "Header",
            "type": "{ [key: string]: Array<string>; }",
            "format": ""
        },
        {
            "name": "initialStatus",
            "baseName": "InitialStatus",
            "type": "string",
            "format": ""
        },
        {
            "name": "interval",
            "baseName": "Interval",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "method",
            "baseName": "Method",
            "type": "string",
            "format": ""
        },
        {
            "name": "name",
            "baseName": "Name",
            "type": "string",
            "format": ""
        },
        {
            "name": "onUpdate",
            "baseName": "OnUpdate",
            "type": "string",
            "format": ""
        },
        {
            "name": "path",
            "baseName": "Path",
            "type": "string",
            "format": ""
        },
        {
            "name": "portLabel",
            "baseName": "PortLabel",
            "type": "string",
            "format": ""
        },
        {
            "name": "protocol",
            "baseName": "Protocol",
            "type": "string",
            "format": ""
        },
        {
            "name": "successBeforePassing",
            "baseName": "SuccessBeforePassing",
            "type": "number",
            "format": ""
        },
        {
            "name": "tLSSkipVerify",
            "baseName": "TLSSkipVerify",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "taskName",
            "baseName": "TaskName",
            "type": "string",
            "format": ""
        },
        {
            "name": "timeout",
            "baseName": "Timeout",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "type",
            "baseName": "Type",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ServiceCheck.attributeTypeMap;
    }
    
    public constructor() {
    }
}

