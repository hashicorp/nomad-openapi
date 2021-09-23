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
pub struct AllocDeploymentStatus {
    #[serde(rename = "Canary", skip_serializing_if = "Option::is_none")]
    pub canary: Option<bool>,
    #[serde(rename = "Healthy", skip_serializing_if = "Option::is_none")]
    pub healthy: Option<bool>,
    #[serde(rename = "ModifyIndex", skip_serializing_if = "Option::is_none")]
    pub modify_index: Option<i32>,
    #[serde(rename = "Timestamp", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<String>,
}

impl AllocDeploymentStatus {
    pub fn new() -> AllocDeploymentStatus {
        AllocDeploymentStatus {
            canary: None,
            healthy: None,
            modify_index: None,
            timestamp: None,
        }
    }
}


