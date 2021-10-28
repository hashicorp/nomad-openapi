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

import { NodeScoreMeta } from './NodeScoreMeta';
import { Resources } from './Resources';
import { HttpFile } from '../http/http';

export class AllocationMetric {
    'allocationTime'?: number;
    'classExhausted'?: { [key: string]: number; };
    'classFiltered'?: { [key: string]: number; };
    'coalescedFailures'?: number;
    'constraintFiltered'?: { [key: string]: number; };
    'dimensionExhausted'?: { [key: string]: number; };
    'nodesAvailable'?: { [key: string]: number; };
    'nodesEvaluated'?: number;
    'nodesExhausted'?: number;
    'nodesFiltered'?: number;
    'quotaExhausted'?: Array<string>;
    'resourcesExhausted'?: { [key: string]: Resources; };
    'scoreMetaData'?: Array<NodeScoreMeta>;
    'scores'?: { [key: string]: number; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "allocationTime",
            "baseName": "AllocationTime",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "classExhausted",
            "baseName": "ClassExhausted",
            "type": "{ [key: string]: number; }",
            "format": ""
        },
        {
            "name": "classFiltered",
            "baseName": "ClassFiltered",
            "type": "{ [key: string]: number; }",
            "format": ""
        },
        {
            "name": "coalescedFailures",
            "baseName": "CoalescedFailures",
            "type": "number",
            "format": ""
        },
        {
            "name": "constraintFiltered",
            "baseName": "ConstraintFiltered",
            "type": "{ [key: string]: number; }",
            "format": ""
        },
        {
            "name": "dimensionExhausted",
            "baseName": "DimensionExhausted",
            "type": "{ [key: string]: number; }",
            "format": ""
        },
        {
            "name": "nodesAvailable",
            "baseName": "NodesAvailable",
            "type": "{ [key: string]: number; }",
            "format": ""
        },
        {
            "name": "nodesEvaluated",
            "baseName": "NodesEvaluated",
            "type": "number",
            "format": ""
        },
        {
            "name": "nodesExhausted",
            "baseName": "NodesExhausted",
            "type": "number",
            "format": ""
        },
        {
            "name": "nodesFiltered",
            "baseName": "NodesFiltered",
            "type": "number",
            "format": ""
        },
        {
            "name": "quotaExhausted",
            "baseName": "QuotaExhausted",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "resourcesExhausted",
            "baseName": "ResourcesExhausted",
            "type": "{ [key: string]: Resources; }",
            "format": ""
        },
        {
            "name": "scoreMetaData",
            "baseName": "ScoreMetaData",
            "type": "Array<NodeScoreMeta>",
            "format": ""
        },
        {
            "name": "scores",
            "baseName": "Scores",
            "type": "{ [key: string]: number; }",
            "format": "double"
        }    ];

    static getAttributeTypeMap() {
        return AllocationMetric.attributeTypeMap;
    }
    
    public constructor() {
    }
}

