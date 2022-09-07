/*
Nomad

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.4
Contact: support@hashicorp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// SystemApiService SystemApi service
type SystemApiService service

type ApiPutSystemGCRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

// Filters results based on the specified region.
func (r ApiPutSystemGCRequest) Region(region string) ApiPutSystemGCRequest {
	r.region = &region
	return r
}
// Filters results based on the specified namespace.
func (r ApiPutSystemGCRequest) Namespace(namespace string) ApiPutSystemGCRequest {
	r.namespace = &namespace
	return r
}
// A Nomad ACL token.
func (r ApiPutSystemGCRequest) XNomadToken(xNomadToken string) ApiPutSystemGCRequest {
	r.xNomadToken = &xNomadToken
	return r
}
// Can be used to ensure operations are only run once.
func (r ApiPutSystemGCRequest) IdempotencyToken(idempotencyToken string) ApiPutSystemGCRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiPutSystemGCRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutSystemGCExecute(r)
}

/*
PutSystemGC Method for PutSystemGC

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutSystemGCRequest
*/
func (a *SystemApiService) PutSystemGC(ctx context.Context) ApiPutSystemGCRequest {
	return ApiPutSystemGCRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SystemApiService) PutSystemGCExecute(r ApiPutSystemGCRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.PutSystemGC")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/gc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.idempotencyToken != nil {
		localVarQueryParams.Add("idempotency_token", parameterToString(*r.idempotencyToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xNomadToken != nil {
		localVarHeaderParams["X-Nomad-Token"] = parameterToString(*r.xNomadToken, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPutSystemReconcileSummariesRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

// Filters results based on the specified region.
func (r ApiPutSystemReconcileSummariesRequest) Region(region string) ApiPutSystemReconcileSummariesRequest {
	r.region = &region
	return r
}
// Filters results based on the specified namespace.
func (r ApiPutSystemReconcileSummariesRequest) Namespace(namespace string) ApiPutSystemReconcileSummariesRequest {
	r.namespace = &namespace
	return r
}
// A Nomad ACL token.
func (r ApiPutSystemReconcileSummariesRequest) XNomadToken(xNomadToken string) ApiPutSystemReconcileSummariesRequest {
	r.xNomadToken = &xNomadToken
	return r
}
// Can be used to ensure operations are only run once.
func (r ApiPutSystemReconcileSummariesRequest) IdempotencyToken(idempotencyToken string) ApiPutSystemReconcileSummariesRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiPutSystemReconcileSummariesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutSystemReconcileSummariesExecute(r)
}

/*
PutSystemReconcileSummaries Method for PutSystemReconcileSummaries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutSystemReconcileSummariesRequest
*/
func (a *SystemApiService) PutSystemReconcileSummaries(ctx context.Context) ApiPutSystemReconcileSummariesRequest {
	return ApiPutSystemReconcileSummariesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SystemApiService) PutSystemReconcileSummariesExecute(r ApiPutSystemReconcileSummariesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.PutSystemReconcileSummaries")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reconcile/summaries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.idempotencyToken != nil {
		localVarQueryParams.Add("idempotency_token", parameterToString(*r.idempotencyToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xNomadToken != nil {
		localVarHeaderParams["X-Nomad-Token"] = parameterToString(*r.xNomadToken, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
