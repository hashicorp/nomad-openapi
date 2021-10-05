use std::rc::Rc;

use hyper;
use super::configuration::Configuration;

pub struct APIClient {
    acl_api: Box<dyn crate::apis::ACLApi>,
    allocations_api: Box<dyn crate::apis::AllocationsApi>,
    enterprise_api: Box<dyn crate::apis::EnterpriseApi>,
    jobs_api: Box<dyn crate::apis::JobsApi>,
    metrics_api: Box<dyn crate::apis::MetricsApi>,
    namespaces_api: Box<dyn crate::apis::NamespacesApi>,
    plugins_api: Box<dyn crate::apis::PluginsApi>,
    regions_api: Box<dyn crate::apis::RegionsApi>,
    scaling_api: Box<dyn crate::apis::ScalingApi>,
    search_api: Box<dyn crate::apis::SearchApi>,
    system_api: Box<dyn crate::apis::SystemApi>,
    volumes_api: Box<dyn crate::apis::VolumesApi>,
}

impl APIClient {
    pub fn new<C: hyper::client::Connect>(configuration: Configuration<C>) -> APIClient {
        let rc = Rc::new(configuration);

        APIClient {
            acl_api: Box::new(crate::apis::ACLApiClient::new(rc.clone())),
            allocations_api: Box::new(crate::apis::AllocationsApiClient::new(rc.clone())),
            enterprise_api: Box::new(crate::apis::EnterpriseApiClient::new(rc.clone())),
            jobs_api: Box::new(crate::apis::JobsApiClient::new(rc.clone())),
            metrics_api: Box::new(crate::apis::MetricsApiClient::new(rc.clone())),
            namespaces_api: Box::new(crate::apis::NamespacesApiClient::new(rc.clone())),
            plugins_api: Box::new(crate::apis::PluginsApiClient::new(rc.clone())),
            regions_api: Box::new(crate::apis::RegionsApiClient::new(rc.clone())),
            scaling_api: Box::new(crate::apis::ScalingApiClient::new(rc.clone())),
            search_api: Box::new(crate::apis::SearchApiClient::new(rc.clone())),
            system_api: Box::new(crate::apis::SystemApiClient::new(rc.clone())),
            volumes_api: Box::new(crate::apis::VolumesApiClient::new(rc.clone())),
        }
    }

    pub fn acl_api(&self) -> &dyn crate::apis::ACLApi{
        self.acl_api.as_ref()
    }

    pub fn allocations_api(&self) -> &dyn crate::apis::AllocationsApi{
        self.allocations_api.as_ref()
    }

    pub fn enterprise_api(&self) -> &dyn crate::apis::EnterpriseApi{
        self.enterprise_api.as_ref()
    }

    pub fn jobs_api(&self) -> &dyn crate::apis::JobsApi{
        self.jobs_api.as_ref()
    }

    pub fn metrics_api(&self) -> &dyn crate::apis::MetricsApi{
        self.metrics_api.as_ref()
    }

    pub fn namespaces_api(&self) -> &dyn crate::apis::NamespacesApi{
        self.namespaces_api.as_ref()
    }

    pub fn plugins_api(&self) -> &dyn crate::apis::PluginsApi{
        self.plugins_api.as_ref()
    }

    pub fn regions_api(&self) -> &dyn crate::apis::RegionsApi{
        self.regions_api.as_ref()
    }

    pub fn scaling_api(&self) -> &dyn crate::apis::ScalingApi{
        self.scaling_api.as_ref()
    }

    pub fn search_api(&self) -> &dyn crate::apis::SearchApi{
        self.search_api.as_ref()
    }

    pub fn system_api(&self) -> &dyn crate::apis::SystemApi{
        self.system_api.as_ref()
    }

    pub fn volumes_api(&self) -> &dyn crate::apis::VolumesApi{
        self.volumes_api.as_ref()
    }

}
