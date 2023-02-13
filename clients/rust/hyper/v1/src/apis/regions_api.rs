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

pub struct RegionsApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> RegionsApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> RegionsApiClient<C> {
        RegionsApiClient {
            configuration,
        }
    }
}

pub trait RegionsApi {
    fn get_regions(&self, ) -> Box<dyn Future<Item = Vec<String>, Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>RegionsApi for RegionsApiClient<C> {
    fn get_regions(&self, ) -> Box<dyn Future<Item = Vec<String>, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Get, "/regions".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;

        req.execute(self.configuration.borrow())
    }

}
