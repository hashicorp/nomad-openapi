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
pub struct CsiNodeInfo {
    #[serde(rename = "AccessibleTopology", skip_serializing_if = "Option::is_none")]
    pub accessible_topology: Option<Box<crate::models::CsiTopology>>,
    #[serde(rename = "ID", skip_serializing_if = "Option::is_none")]
    pub ID: Option<String>,
    #[serde(rename = "MaxVolumes", skip_serializing_if = "Option::is_none")]
    pub max_volumes: Option<i64>,
    #[serde(rename = "RequiresNodeStageVolume", skip_serializing_if = "Option::is_none")]
    pub requires_node_stage_volume: Option<bool>,
    #[serde(rename = "SupportsCondition", skip_serializing_if = "Option::is_none")]
    pub supports_condition: Option<bool>,
    #[serde(rename = "SupportsExpand", skip_serializing_if = "Option::is_none")]
    pub supports_expand: Option<bool>,
    #[serde(rename = "SupportsStats", skip_serializing_if = "Option::is_none")]
    pub supports_stats: Option<bool>,
}

impl CsiNodeInfo {
    pub fn new() -> CsiNodeInfo {
        CsiNodeInfo {
            accessible_topology: None,
            ID: None,
            max_volumes: None,
            requires_node_stage_volume: None,
            supports_condition: None,
            supports_expand: None,
            supports_stats: None,
        }
    }
}


