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

pub struct ACLApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> ACLApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> ACLApiClient<C> {
        ACLApiClient {
            configuration,
        }
    }
}

pub trait ACLApi {
    fn delete_acl_policy(&self, policy_name: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>>;
    fn delete_acl_token(&self, token_accessor: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>>;
    fn get_acl_policies(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclPolicyListStub>, Error = Error<serde_json::Value>>>;
    fn get_acl_policy(&self, policy_name: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclPolicy, Error = Error<serde_json::Value>>>;
    fn get_acl_token(&self, token_accessor: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
    fn get_acl_token_self(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
    fn get_acl_tokens(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclTokenListStub>, Error = Error<serde_json::Value>>>;
    fn post_acl_bootstrap(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>>;
    fn post_acl_policy(&self, policy_name: &str, acl_policy: crate::models::AclPolicy, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>>;
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>>;
>>>>>>> b1a8ea3 (tokenAccessorParam --> aclTokenAccessorParam to stay aligned with naming)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
>>>>>>> eed23b0 (acl.go PostACLToken should return objectSchema and not arraySchema, added v1/acl.go)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>>;
>>>>>>> e073f3a (merge)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
>>>>>>> 87acc5b (acl.go PostACLToken should return objectSchema and not arraySchema, added v1/acl.go)
    fn post_acl_token_onetime(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::OneTimeToken, Error = Error<serde_json::Value>>>;
    fn post_acl_token_onetime_exchange(&self, one_time_token_exchange_request: crate::models::OneTimeTokenExchangeRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>ACLApi for ACLApiClient<C> {
    fn delete_acl_policy(&self, policy_name: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Delete, "/acl/policy/{policyName}".to_string())
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
        req = req.with_path_param("policyName".to_string(), policy_name.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.returns_nothing();

        req.execute(self.configuration.borrow())
    }

    fn delete_acl_token(&self, token_accessor: &str, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Delete, "/acl/token/{tokenAccessor}".to_string())
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
        req = req.with_path_param("tokenAccessor".to_string(), token_accessor.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.returns_nothing();

        req.execute(self.configuration.borrow())
    }

    fn get_acl_policies(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclPolicyListStub>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/acl/policies".to_string())
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

    fn get_acl_policy(&self, policy_name: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclPolicy, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/acl/policy/{policyName}".to_string())
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
        req = req.with_path_param("policyName".to_string(), policy_name.to_string());
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn get_acl_token(&self, token_accessor: &str, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/acl/token/{tokenAccessor}".to_string())
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
        req = req.with_path_param("tokenAccessor".to_string(), token_accessor.to_string());
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }

        req.execute(self.configuration.borrow())
    }

    fn get_acl_token_self(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/acl/token".to_string())
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

    fn get_acl_tokens(&self, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclTokenListStub>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/acl/tokens".to_string())
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

    fn post_acl_bootstrap(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/acl/bootstrap".to_string())
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

        req.execute(self.configuration.borrow())
    }

    fn post_acl_policy(&self, policy_name: &str, acl_policy: crate::models::AclPolicy, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = (), Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/acl/policy/{policyName}".to_string())
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
        req = req.with_path_param("policyName".to_string(), policy_name.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(acl_policy);
        req = req.returns_nothing();

        req.execute(self.configuration.borrow())
    }

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>> {
>>>>>>> b1a8ea3 (tokenAccessorParam --> aclTokenAccessorParam to stay aligned with naming)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
>>>>>>> eed23b0 (acl.go PostACLToken should return objectSchema and not arraySchema, added v1/acl.go)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = Vec<crate::models::AclToken>, Error = Error<serde_json::Value>>> {
>>>>>>> e073f3a (merge)
=======
    fn post_acl_token(&self, token_accessor: &str, acl_token: crate::models::AclToken, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
>>>>>>> 87acc5b (acl.go PostACLToken should return objectSchema and not arraySchema, added v1/acl.go)
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/acl/token/{tokenAccessor}".to_string())
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
        req = req.with_path_param("tokenAccessor".to_string(), token_accessor.to_string());
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(acl_token);

        req.execute(self.configuration.borrow())
    }

    fn post_acl_token_onetime(&self, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::OneTimeToken, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/acl/token/onetime".to_string())
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

        req.execute(self.configuration.borrow())
    }

    fn post_acl_token_onetime_exchange(&self, one_time_token_exchange_request: crate::models::OneTimeTokenExchangeRequest, region: Option<&str>, namespace: Option<&str>, x_nomad_token: Option<&str>, idempotency_token: Option<&str>) -> Box<dyn Future<Item = crate::models::AclToken, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/acl/token/onetime/exchange".to_string())
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
        req = req.with_body_param(one_time_token_exchange_request);

        req.execute(self.configuration.borrow())
    }

}
