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
pub struct RescheduleEvent {
    #[serde(rename = "PrevAllocID", skip_serializing_if = "Option::is_none")]
    pub prev_alloc_id: Option<String>,
    #[serde(rename = "PrevNodeID", skip_serializing_if = "Option::is_none")]
    pub prev_node_id: Option<String>,
    #[serde(rename = "RescheduleTime", skip_serializing_if = "Option::is_none")]
    pub reschedule_time: Option<i64>,
}

impl RescheduleEvent {
    pub fn new() -> RescheduleEvent {
        RescheduleEvent {
            prev_alloc_id: None,
            prev_node_id: None,
            reschedule_time: None,
        }
    }
}


