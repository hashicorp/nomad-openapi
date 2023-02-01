// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://openapi-generator.tech
 */

use std::rc::Rc;
use std::borrow::Borrow;
#[allow(unused_imports)]
use std::option::Option;

use hyper;
use serde_json;
use futures::Future;

use super::{Error, configuration};
use super::request as __internal_request;

pub struct DeploymentsApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> DeploymentsApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> DeploymentsApiClient<C> {
        DeploymentsApiClient {
            configuration,
        }
    }
}

pub trait DeploymentsApi {
    fn get_deployment(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::Deployment, Error = Error<serde_json::Value>>>;
    fn get_deployment_allocations(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AllocationListStub>, Error = Error<serde_json::Value>>>;
    fn get_deployments(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::Deployment>, Error = Error<serde_json::Value>>>;
    fn post_deployment_allocation_health(&self, deployment_id: &str, deployment_alloc_health_request: crate::models::DeploymentAllocHealthRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>>;
    fn post_deployment_fail(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>>;
    fn post_deployment_pause(&self, deployment_id: &str, deployment_pause_request: crate::models::DeploymentPauseRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>>;
    fn post_deployment_promote(&self, deployment_id: &str, deployment_promote_request: crate::models::DeploymentPromoteRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>>;
    fn post_deployment_unblock(&self, deployment_id: &str, deployment_unblock_request: crate::models::DeploymentUnblockRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>DeploymentsApi for DeploymentsApiClient<C> {
    fn get_deployment(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::Deployment, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/deployment/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = wait {
            req = req.with_query_param("wait".to_string(), s.to_string());
        }
        if let Some(ref s) = stale {
            req = req.with_query_param("stale".to_string(), s.to_string());
        }
        if let Some(ref s) = prefix {
            req = req.with_query_param("prefix".to_string(), s.to_string());
        }
        if let Some(ref s) = per_page {
            req = req.with_query_param("per_page".to_string(), s.to_string());
        }
        if let Some(ref s) = next_token {
            req = req.with_query_param("next_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn get_deployment_allocations(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AllocationListStub>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/deployment/allocations/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = wait {
            req = req.with_query_param("wait".to_string(), s.to_string());
        }
        if let Some(ref s) = stale {
            req = req.with_query_param("stale".to_string(), s.to_string());
        }
        if let Some(ref s) = prefix {
            req = req.with_query_param("prefix".to_string(), s.to_string());
        }
        if let Some(ref s) = per_page {
            req = req.with_query_param("per_page".to_string(), s.to_string());
        }
        if let Some(ref s) = next_token {
            req = req.with_query_param("next_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn get_deployments(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::Deployment>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/deployments".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = wait {
            req = req.with_query_param("wait".to_string(), s.to_string());
        }
        if let Some(ref s) = stale {
            req = req.with_query_param("stale".to_string(), s.to_string());
        }
        if let Some(ref s) = prefix {
            req = req.with_query_param("prefix".to_string(), s.to_string());
        }
        if let Some(ref s) = per_page {
            req = req.with_query_param("per_page".to_string(), s.to_string());
        }
        if let Some(ref s) = next_token {
            req = req.with_query_param("next_token".to_string(), s.to_string());
        }
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn post_deployment_allocation_health(&self, deployment_id: &str, deployment_alloc_health_request: crate::models::DeploymentAllocHealthRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/deployment/allocation-health/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = idempotency_token {
            req = req.with_query_param("idempotency_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(deployment_alloc_health_request);

        req.execute(self.configuration.borrow())
    }

    fn post_deployment_fail(&self, deployment_id: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/deployment/fail/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = idempotency_token {
            req = req.with_query_param("idempotency_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn post_deployment_pause(&self, deployment_id: &str, deployment_pause_request: crate::models::DeploymentPauseRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/deployment/pause/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = idempotency_token {
            req = req.with_query_param("idempotency_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(deployment_pause_request);

        req.execute(self.configuration.borrow())
    }

    fn post_deployment_promote(&self, deployment_id: &str, deployment_promote_request: crate::models::DeploymentPromoteRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/deployment/promote/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = idempotency_token {
            req = req.with_query_param("idempotency_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(deployment_promote_request);

        req.execute(self.configuration.borrow())
    }

    fn post_deployment_unblock(&self, deployment_id: &str, deployment_unblock_request: crate::models::DeploymentUnblockRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::DeploymentUpdateResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/deployment/unblock/{deploymentID}".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = idempotency_token {
            req = req.with_query_param("idempotency_token".to_string(), s.to_string());
        }
        req = req.with_path_param("deploymentID".to_string(), deployment_id.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(deployment_unblock_request);

        req.execute(self.configuration.borrow())
    }

}
