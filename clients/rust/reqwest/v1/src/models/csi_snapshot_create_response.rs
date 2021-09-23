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
pub struct CsiSnapshotCreateResponse {
    #[serde(rename = "KnownLeader", skip_serializing_if = "Option::is_none")]
    pub known_leader: Option<bool>,
    #[serde(rename = "LastContact", skip_serializing_if = "Option::is_none")]
    pub last_contact: Option<i64>,
    #[serde(rename = "LastIndex", skip_serializing_if = "Option::is_none")]
    pub last_index: Option<i32>,
    #[serde(rename = "RequestTime", skip_serializing_if = "Option::is_none")]
    pub request_time: Option<i64>,
    #[serde(rename = "Snapshots", skip_serializing_if = "Option::is_none")]
    pub snapshots: Option<Vec<crate::models::CsiSnapshot>>,
}

impl CsiSnapshotCreateResponse {
    pub fn new() -> CsiSnapshotCreateResponse {
        CsiSnapshotCreateResponse {
            known_leader: None,
            last_contact: None,
            last_index: None,
            request_time: None,
            snapshots: None,
        }
    }
}


