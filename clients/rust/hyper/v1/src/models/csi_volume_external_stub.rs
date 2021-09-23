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
pub struct CsiVolumeExternalStub {
    #[serde(rename = "CapacityBytes", skip_serializing_if = "Option::is_none")]
    pub capacity_bytes: Option<i64>,
    #[serde(rename = "CloneID", skip_serializing_if = "Option::is_none")]
    pub clone_id: Option<String>,
    #[serde(rename = "ExternalID", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<String>,
    #[serde(rename = "IsAbnormal", skip_serializing_if = "Option::is_none")]
    pub is_abnormal: Option<bool>,
    #[serde(rename = "PublishedExternalNodeIDs", skip_serializing_if = "Option::is_none")]
    pub published_external_node_ids: Option<Vec<String>>,
    #[serde(rename = "SnapshotID", skip_serializing_if = "Option::is_none")]
    pub snapshot_id: Option<String>,
    #[serde(rename = "Status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    #[serde(rename = "VolumeContext", skip_serializing_if = "Option::is_none")]
    pub volume_context: Option<::std::collections::HashMap<String, String>>,
}

impl CsiVolumeExternalStub {
    pub fn new() -> CsiVolumeExternalStub {
        CsiVolumeExternalStub {
            capacity_bytes: None,
            clone_id: None,
            external_id: None,
            is_abnormal: None,
            published_external_node_ids: None,
            snapshot_id: None,
            status: None,
            volume_context: None,
        }
    }
}


