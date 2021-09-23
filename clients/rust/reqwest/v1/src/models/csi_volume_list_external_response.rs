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
pub struct CsiVolumeListExternalResponse {
    #[serde(rename = "NextToken", skip_serializing_if = "Option::is_none")]
    pub next_token: Option<String>,
    #[serde(rename = "Volumes", skip_serializing_if = "Option::is_none")]
    pub volumes: Option<Vec<crate::models::CsiVolumeExternalStub>>,
}

impl CsiVolumeListExternalResponse {
    pub fn new() -> CsiVolumeListExternalResponse {
        CsiVolumeListExternalResponse {
            next_token: None,
            volumes: None,
        }
    }
}


