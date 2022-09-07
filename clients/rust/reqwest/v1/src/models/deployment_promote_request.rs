/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct DeploymentPromoteRequest {
    #[serde(rename = "All", skip_serializing_if = "Option::is_none")]
    pub all: Option<bool>,
    #[serde(rename = "DeploymentID", skip_serializing_if = "Option::is_none")]
    pub deployment_id: Option<String>,
    #[serde(rename = "Groups", skip_serializing_if = "Option::is_none")]
    pub groups: Option<Vec<String>>,
    #[serde(rename = "Namespace", skip_serializing_if = "Option::is_none")]
    pub namespace: Option<String>,
    #[serde(rename = "Region", skip_serializing_if = "Option::is_none")]
    pub region: Option<String>,
    #[serde(rename = "SecretID", skip_serializing_if = "Option::is_none")]
    pub secret_id: Option<String>,
}

impl DeploymentPromoteRequest {
    pub fn new() -> DeploymentPromoteRequest {
        DeploymentPromoteRequest {
            all: None,
            deployment_id: None,
            groups: None,
            namespace: None,
            region: None,
            secret_id: None,
        }
    }
}


