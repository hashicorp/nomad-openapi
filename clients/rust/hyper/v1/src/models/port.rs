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
pub struct Port {
    #[serde(rename = "HostNetwork", skip_serializing_if = "Option::is_none")]
    pub host_network: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "To", skip_serializing_if = "Option::is_none")]
    pub to: Option<i32>,
    #[serde(rename = "Value", skip_serializing_if = "Option::is_none")]
    pub value: Option<i32>,
}

impl Port {
    pub fn new() -> Port {
        Port {
            host_network: None,
            label: None,
            to: None,
            value: None,
        }
    }
}


