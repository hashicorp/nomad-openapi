use hyper;
use serde;
use serde_json;

#[derive(Debug)]
pub enum Error<T> {
    UriError(hyper::error::UriError),
    Hyper(hyper::Error),
    Serde(serde_json::Error),
    ApiError(ApiError<T>),
}

#[derive(Debug)]
pub struct ApiError<T> {
    pub code: hyper::StatusCode,
    pub content: Option<T>,
}

impl<'de, T> From<(hyper::StatusCode, &'de [u8])> for Error<T> 
    where T: serde::Deserialize<'de> {
    fn from(e: (hyper::StatusCode, &'de [u8])) -> Self {
        if e.1.len() == 0 {
            return Error::ApiError(ApiError{
                code: e.0,
                content: None,
            });
        }
        match serde_json::from_slice::<T>(e.1) {
            Ok(t) => Error::ApiError(ApiError{
                code: e.0,
                content: Some(t),
            }),
            Err(e) => {
                Error::from(e)
            }
        }
    }
}

impl<T> From<hyper::Error> for Error<T> {
    fn from(e: hyper::Error) -> Self {
        return Error::Hyper(e)
    }
}

impl<T> From<serde_json::Error> for Error<T> {
    fn from(e: serde_json::Error) -> Self {
        return Error::Serde(e)
    }
}

mod request;

mod acl_api;
pub use self::acl_api::{ ACLApi, ACLApiClient };
mod allocations_api;
pub use self::allocations_api::{ AllocationsApi, AllocationsApiClient };
mod enterprise_api;
pub use self::enterprise_api::{ EnterpriseApi, EnterpriseApiClient };
mod jobs_api;
pub use self::jobs_api::{ JobsApi, JobsApiClient };
mod metrics_api;
pub use self::metrics_api::{ MetricsApi, MetricsApiClient };
mod namespaces_api;
pub use self::namespaces_api::{ NamespacesApi, NamespacesApiClient };
mod regions_api;
pub use self::regions_api::{ RegionsApi, RegionsApiClient };
mod scaling_api;
pub use self::scaling_api::{ ScalingApi, ScalingApiClient };
mod search_api;
pub use self::search_api::{ SearchApi, SearchApiClient };
mod volumes_api;
pub use self::volumes_api::{ VolumesApi, VolumesApiClient };

pub mod configuration;
pub mod client;
