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
pub struct TaskCsiPluginConfig {
    #[serde(rename = "ID", skip_serializing_if = "Option::is_none")]
    pub ID: Option<String>,
    #[serde(rename = "MountDir", skip_serializing_if = "Option::is_none")]
    pub mount_dir: Option<String>,
    #[serde(rename = "Type", skip_serializing_if = "Option::is_none")]
    pub _type: Option<String>,
}

impl TaskCsiPluginConfig {
    pub fn new() -> TaskCsiPluginConfig {
        TaskCsiPluginConfig {
            ID: None,
            mount_dir: None,
            _type: None,
        }
    }
}


