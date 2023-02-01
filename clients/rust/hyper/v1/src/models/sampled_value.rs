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
pub struct SampledValue {
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i32>,
    #[serde(rename = "Labels", skip_serializing_if = "Option::is_none")]
    pub labels: Option<::std::collections::HashMap<String, String>>,
    #[serde(rename = "Max", skip_serializing_if = "Option::is_none")]
    pub max: Option<f64>,
    #[serde(rename = "Mean", skip_serializing_if = "Option::is_none")]
    pub mean: Option<f64>,
    #[serde(rename = "Min", skip_serializing_if = "Option::is_none")]
    pub min: Option<f64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Rate", skip_serializing_if = "Option::is_none")]
    pub rate: Option<f64>,
    #[serde(rename = "Stddev", skip_serializing_if = "Option::is_none")]
    pub stddev: Option<f64>,
    #[serde(rename = "Sum", skip_serializing_if = "Option::is_none")]
    pub sum: Option<f64>,
}

impl SampledValue {
    pub fn new() -> SampledValue {
        SampledValue {
            count: None,
            labels: None,
            max: None,
            mean: None,
            min: None,
            name: None,
            rate: None,
            stddev: None,
            sum: None,
        }
    }
}


