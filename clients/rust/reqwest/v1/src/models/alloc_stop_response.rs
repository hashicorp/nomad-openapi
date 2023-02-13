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
pub struct AllocStopResponse {
    #[serde(rename = "EvalID", skip_serializing_if = "Option::is_none")]
    pub eval_id: Option<String>,
    #[serde(rename = "Index", skip_serializing_if = "Option::is_none")]
    pub index: Option<i32>,
}

impl AllocStopResponse {
    pub fn new() -> AllocStopResponse {
        AllocStopResponse {
            eval_id: None,
            index: None,
        }
    }
}


