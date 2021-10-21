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

pub struct OperatorApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> OperatorApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> OperatorApiClient<C> {
        OperatorApiClient {
            configuration,
        }
    }
}

pub trait OperatorApi {
    fn delete_operator_raft(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>>;
    fn get_operator_autopilot_configuration(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AutopilotConfiguration, Error = Error<serde_json::Value>>>;
    fn get_operator_autopilot_health(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::OperatorHealthReply, Error = Error<serde_json::Value>>>;
    fn get_operator_raft(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::RaftServer>, Error = Error<serde_json::Value>>>;
    fn get_operator_scheduler_configuration(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SchedulerConfigurationResponse, Error = Error<serde_json::Value>>>;
    fn post_operator_scheduler_configuration(&self, scheduler_configuration: crate::models::SchedulerConfiguration, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SchedulerSetConfigurationResponse, Error = Error<serde_json::Value>>>;
    fn put_operator_autopilot_configuration(&self, autopilot_configuration: crate::models::AutopilotConfiguration, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>OperatorApi for OperatorApiClient<C> {
    fn delete_operator_raft(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Delete, "/operator/raft/".to_string())
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
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.returns_nothing();

        req.execute(self.configuration.borrow())
    }

    fn get_operator_autopilot_configuration(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AutopilotConfiguration, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/operator/autopilot/configuration".to_string())
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

    fn get_operator_autopilot_health(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::OperatorHealthReply, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/operator/autopilot/health".to_string())
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

    fn get_operator_raft(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::RaftServer>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/operator/raft/".to_string())
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

    fn get_operator_scheduler_configuration(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SchedulerConfigurationResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/operator/scheduler/configuration".to_string())
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

    fn post_operator_scheduler_configuration(&self, scheduler_configuration: crate::models::SchedulerConfiguration, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SchedulerSetConfigurationResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/operator/scheduler/configuration".to_string())
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
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(scheduler_configuration);

        req.execute(self.configuration.borrow())
    }

    fn put_operator_autopilot_configuration(&self, autopilot_configuration: crate::models::AutopilotConfiguration, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Put, "/operator/autopilot/configuration".to_string())
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
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(autopilot_configuration);
        req = req.returns_nothing();

        req.execute(self.configuration.borrow())
    }

}
