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
pub struct ScalingRequest {
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "Error", skip_serializing_if = "Option::is_none")]
    pub error: Option<bool>,
    #[serde(rename = "Message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename = "Meta", skip_serializing_if = "Option::is_none")]
    pub meta: Option<::std::collections::HashMap<String, serde_json::Value>>,
    #[serde(rename = "Namespace", skip_serializing_if = "Option::is_none")]
    pub namespace: Option<String>,
    #[serde(rename = "PolicyOverride", skip_serializing_if = "Option::is_none")]
    pub policy_override: Option<bool>,
    #[serde(rename = "Region", skip_serializing_if = "Option::is_none")]
    pub region: Option<String>,
    #[serde(rename = "SecretID", skip_serializing_if = "Option::is_none")]
    pub secret_id: Option<String>,
    #[serde(rename = "Target", skip_serializing_if = "Option::is_none")]
    pub target: Option<::std::collections::HashMap<String, String>>,
}

impl ScalingRequest {
    pub fn new() -> ScalingRequest {
        ScalingRequest {
            count: None,
            error: None,
            message: None,
            meta: None,
            namespace: None,
            policy_override: None,
            region: None,
            secret_id: None,
            target: None,
        }
    }
}


