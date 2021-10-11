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
pub struct QuotaLimit {
    #[serde(rename = "Hash", skip_serializing_if = "Option::is_none")]
    pub hash: Option<String>,
    #[serde(rename = "Region", skip_serializing_if = "Option::is_none")]
    pub region: Option<String>,
    #[serde(rename = "RegionLimit", skip_serializing_if = "Option::is_none")]
    pub region_limit: Option<Box<crate::models::Resources>>,
}

impl QuotaLimit {
    pub fn new() -> QuotaLimit {
        QuotaLimit {
            hash: None,
            region: None,
            region_limit: None,
        }
    }
}

