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
pub struct DeploymentState {
    #[serde(rename = "AutoRevert", skip_serializing_if = "Option::is_none")]
    pub auto_revert: Option<bool>,
    #[serde(rename = "DesiredCanaries", skip_serializing_if = "Option::is_none")]
    pub desired_canaries: Option<i32>,
    #[serde(rename = "DesiredTotal", skip_serializing_if = "Option::is_none")]
    pub desired_total: Option<i32>,
    #[serde(rename = "HealthyAllocs", skip_serializing_if = "Option::is_none")]
    pub healthy_allocs: Option<i32>,
    #[serde(rename = "PlacedAllocs", skip_serializing_if = "Option::is_none")]
    pub placed_allocs: Option<i32>,
    #[serde(rename = "PlacedCanaries", skip_serializing_if = "Option::is_none")]
    pub placed_canaries: Option<Vec<String>>,
    #[serde(rename = "ProgressDeadline", skip_serializing_if = "Option::is_none")]
    pub progress_deadline: Option<i64>,
    #[serde(rename = "Promoted", skip_serializing_if = "Option::is_none")]
    pub promoted: Option<bool>,
    #[serde(rename = "RequireProgressBy", skip_serializing_if = "Option::is_none")]
    pub require_progress_by: Option<String>,
    #[serde(rename = "UnhealthyAllocs", skip_serializing_if = "Option::is_none")]
    pub unhealthy_allocs: Option<i32>,
}

impl DeploymentState {
    pub fn new() -> DeploymentState {
        DeploymentState {
            auto_revert: None,
            desired_canaries: None,
            desired_total: None,
            healthy_allocs: None,
            placed_allocs: None,
            placed_canaries: None,
            progress_deadline: None,
            promoted: None,
            require_progress_by: None,
            unhealthy_allocs: None,
        }
    }
}


