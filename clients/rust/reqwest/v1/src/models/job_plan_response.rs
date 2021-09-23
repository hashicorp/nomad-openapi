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
pub struct JobPlanResponse {
    #[serde(rename = "Annotations", skip_serializing_if = "Option::is_none")]
    pub annotations: Option<Box<crate::models::PlanAnnotations>>,
    #[serde(rename = "CreatedEvals", skip_serializing_if = "Option::is_none")]
    pub created_evals: Option<Vec<crate::models::Evaluation>>,
    #[serde(rename = "Diff", skip_serializing_if = "Option::is_none")]
    pub diff: Option<Box<crate::models::JobDiff>>,
    #[serde(rename = "FailedTGAllocs", skip_serializing_if = "Option::is_none")]
    pub failed_tg_allocs: Option<::std::collections::HashMap<String, crate::models::AllocationMetric>>,
    #[serde(rename = "JobModifyIndex", skip_serializing_if = "Option::is_none")]
    pub job_modify_index: Option<i32>,
    #[serde(rename = "NextPeriodicLaunch", skip_serializing_if = "Option::is_none")]
    pub next_periodic_launch: Option<String>,
    #[serde(rename = "Warnings", skip_serializing_if = "Option::is_none")]
    pub warnings: Option<String>,
}

impl JobPlanResponse {
    pub fn new() -> JobPlanResponse {
        JobPlanResponse {
            annotations: None,
            created_evals: None,
            diff: None,
            failed_tg_allocs: None,
            job_modify_index: None,
            next_periodic_launch: None,
            warnings: None,
        }
    }
}


