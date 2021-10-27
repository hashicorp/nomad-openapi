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

pub struct PluginsApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> PluginsApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> PluginsApiClient<C> {
        PluginsApiClient {
            configuration,
        }
    }
}

pub trait PluginsApi {
    fn get_plugin_csi(&self, plugin_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::CsiPlugin>, Error = Error<serde_json::Value>>>;
    fn get_plugins(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::CsiPluginListStub>, Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>PluginsApi for PluginsApiClient<C> {
    fn get_plugin_csi(&self, plugin_id: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::CsiPlugin>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/plugin/csi/{pluginID}".to_string())
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
        req = req.with_path_param("pluginID".to_string(), plugin_id.to_string());
=======
        req = req.with_path_param("pluginId".to_string(), plugin_id.to_string());
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
        req = req.with_path_param("pluginID".to_string(), plugin_id.to_string());
>>>>>>> 13a3eee (added tests for plugins)
=======
        req = req.with_path_param("pluginID".to_string(), plugin_id.to_string());
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn get_plugins(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::CsiPluginListStub>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/plugins".to_string())
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

}
