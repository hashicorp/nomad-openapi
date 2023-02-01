// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CsiTopologyRequest {
    #[serde(rename = "Preferred", skip_serializing_if = "Option::is_none")]
    pub preferred: Option<Vec<crate::models::CsiTopology>>,
    #[serde(rename = "Required", skip_serializing_if = "Option::is_none")]
    pub required: Option<Vec<crate::models::CsiTopology>>,
}

impl CsiTopologyRequest {
    pub fn new() -> CsiTopologyRequest {
        CsiTopologyRequest {
            preferred: None,
            required: None,
        }
    }
}


