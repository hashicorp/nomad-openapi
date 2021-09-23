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
pub struct DnsConfig {
    #[serde(rename = "Options", skip_serializing_if = "Option::is_none")]
    pub options: Option<Vec<String>>,
    #[serde(rename = "Searches", skip_serializing_if = "Option::is_none")]
    pub searches: Option<Vec<String>>,
    #[serde(rename = "Servers", skip_serializing_if = "Option::is_none")]
    pub servers: Option<Vec<String>>,
}

impl DnsConfig {
    pub fn new() -> DnsConfig {
        DnsConfig {
            options: None,
            searches: None,
            servers: None,
        }
    }
}


