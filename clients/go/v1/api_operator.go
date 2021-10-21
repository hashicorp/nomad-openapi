/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.4
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// OperatorApiService OperatorApi service
type OperatorApiService service

type ApiDeleteOperatorRaftRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiDeleteOperatorRaftRequest) Region(region string) ApiDeleteOperatorRaftRequest {
	r.region = &region
	return r
}
func (r ApiDeleteOperatorRaftRequest) Namespace(namespace string) ApiDeleteOperatorRaftRequest {
	r.namespace = &namespace
	return r
}
func (r ApiDeleteOperatorRaftRequest) XNomadToken(xNomadToken string) ApiDeleteOperatorRaftRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiDeleteOperatorRaftRequest) IdempotencyToken(idempotencyToken string) ApiDeleteOperatorRaftRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiDeleteOperatorRaftRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteOperatorRaftExecute(r)
}

/*
 * DeleteOperatorRaft Method for DeleteOperatorRaft
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeleteOperatorRaftRequest
 */
func (a *OperatorApiService) DeleteOperatorRaft(ctx _context.Context) ApiDeleteOperatorRaftRequest {
	return ApiDeleteOperatorRaftRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *OperatorApiService) DeleteOperatorRaftExecute(r ApiDeleteOperatorRaftRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.DeleteOperatorRaft")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/raft/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOperatorAutopilotConfigurationRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	region *string
	namespace *string
	index *int32
	wait *string
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiGetOperatorAutopilotConfigurationRequest) Region(region string) ApiGetOperatorAutopilotConfigurationRequest {
	r.region = &region
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) Namespace(namespace string) ApiGetOperatorAutopilotConfigurationRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) Index(index int32) ApiGetOperatorAutopilotConfigurationRequest {
	r.index = &index
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) Wait(wait string) ApiGetOperatorAutopilotConfigurationRequest {
	r.wait = &wait
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) Stale(stale string) ApiGetOperatorAutopilotConfigurationRequest {
	r.stale = &stale
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) Prefix(prefix string) ApiGetOperatorAutopilotConfigurationRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) XNomadToken(xNomadToken string) ApiGetOperatorAutopilotConfigurationRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) PerPage(perPage int32) ApiGetOperatorAutopilotConfigurationRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetOperatorAutopilotConfigurationRequest) NextToken(nextToken string) ApiGetOperatorAutopilotConfigurationRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOperatorAutopilotConfigurationRequest) Execute() (AutopilotConfiguration, *_nethttp.Response, error) {
	return r.ApiService.GetOperatorAutopilotConfigurationExecute(r)
}

/*
 * GetOperatorAutopilotConfiguration Method for GetOperatorAutopilotConfiguration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOperatorAutopilotConfigurationRequest
 */
func (a *OperatorApiService) GetOperatorAutopilotConfiguration(ctx _context.Context) ApiGetOperatorAutopilotConfigurationRequest {
	return ApiGetOperatorAutopilotConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AutopilotConfiguration
 */
func (a *OperatorApiService) GetOperatorAutopilotConfigurationExecute(r ApiGetOperatorAutopilotConfigurationRequest) (AutopilotConfiguration, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AutopilotConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.GetOperatorAutopilotConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/autopilot/configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOperatorAutopilotHealthRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	region *string
	namespace *string
	index *int32
	wait *string
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiGetOperatorAutopilotHealthRequest) Region(region string) ApiGetOperatorAutopilotHealthRequest {
	r.region = &region
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) Namespace(namespace string) ApiGetOperatorAutopilotHealthRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) Index(index int32) ApiGetOperatorAutopilotHealthRequest {
	r.index = &index
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) Wait(wait string) ApiGetOperatorAutopilotHealthRequest {
	r.wait = &wait
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) Stale(stale string) ApiGetOperatorAutopilotHealthRequest {
	r.stale = &stale
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) Prefix(prefix string) ApiGetOperatorAutopilotHealthRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) XNomadToken(xNomadToken string) ApiGetOperatorAutopilotHealthRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) PerPage(perPage int32) ApiGetOperatorAutopilotHealthRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetOperatorAutopilotHealthRequest) NextToken(nextToken string) ApiGetOperatorAutopilotHealthRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOperatorAutopilotHealthRequest) Execute() (OperatorHealthReply, *_nethttp.Response, error) {
	return r.ApiService.GetOperatorAutopilotHealthExecute(r)
}

/*
 * GetOperatorAutopilotHealth Method for GetOperatorAutopilotHealth
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOperatorAutopilotHealthRequest
 */
func (a *OperatorApiService) GetOperatorAutopilotHealth(ctx _context.Context) ApiGetOperatorAutopilotHealthRequest {
	return ApiGetOperatorAutopilotHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return OperatorHealthReply
 */
func (a *OperatorApiService) GetOperatorAutopilotHealthExecute(r ApiGetOperatorAutopilotHealthRequest) (OperatorHealthReply, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OperatorHealthReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.GetOperatorAutopilotHealth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/autopilot/health"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOperatorRaftRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	region *string
	namespace *string
	index *int32
	wait *string
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiGetOperatorRaftRequest) Region(region string) ApiGetOperatorRaftRequest {
	r.region = &region
	return r
}
func (r ApiGetOperatorRaftRequest) Namespace(namespace string) ApiGetOperatorRaftRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetOperatorRaftRequest) Index(index int32) ApiGetOperatorRaftRequest {
	r.index = &index
	return r
}
func (r ApiGetOperatorRaftRequest) Wait(wait string) ApiGetOperatorRaftRequest {
	r.wait = &wait
	return r
}
func (r ApiGetOperatorRaftRequest) Stale(stale string) ApiGetOperatorRaftRequest {
	r.stale = &stale
	return r
}
func (r ApiGetOperatorRaftRequest) Prefix(prefix string) ApiGetOperatorRaftRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetOperatorRaftRequest) XNomadToken(xNomadToken string) ApiGetOperatorRaftRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetOperatorRaftRequest) PerPage(perPage int32) ApiGetOperatorRaftRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetOperatorRaftRequest) NextToken(nextToken string) ApiGetOperatorRaftRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOperatorRaftRequest) Execute() ([]RaftServer, *_nethttp.Response, error) {
	return r.ApiService.GetOperatorRaftExecute(r)
}

/*
 * GetOperatorRaft Method for GetOperatorRaft
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOperatorRaftRequest
 */
func (a *OperatorApiService) GetOperatorRaft(ctx _context.Context) ApiGetOperatorRaftRequest {
	return ApiGetOperatorRaftRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []RaftServer
 */
func (a *OperatorApiService) GetOperatorRaftExecute(r ApiGetOperatorRaftRequest) ([]RaftServer, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RaftServer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.GetOperatorRaft")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/raft/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOperatorSchedulerConfigurationRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	region *string
	namespace *string
	index *int32
	wait *string
	stale *string
	prefix *string
	xNomadToken *string
	perPage *int32
	nextToken *string
}

func (r ApiGetOperatorSchedulerConfigurationRequest) Region(region string) ApiGetOperatorSchedulerConfigurationRequest {
	r.region = &region
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) Namespace(namespace string) ApiGetOperatorSchedulerConfigurationRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) Index(index int32) ApiGetOperatorSchedulerConfigurationRequest {
	r.index = &index
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) Wait(wait string) ApiGetOperatorSchedulerConfigurationRequest {
	r.wait = &wait
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) Stale(stale string) ApiGetOperatorSchedulerConfigurationRequest {
	r.stale = &stale
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) Prefix(prefix string) ApiGetOperatorSchedulerConfigurationRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) XNomadToken(xNomadToken string) ApiGetOperatorSchedulerConfigurationRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) PerPage(perPage int32) ApiGetOperatorSchedulerConfigurationRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetOperatorSchedulerConfigurationRequest) NextToken(nextToken string) ApiGetOperatorSchedulerConfigurationRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOperatorSchedulerConfigurationRequest) Execute() (SchedulerConfigurationResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOperatorSchedulerConfigurationExecute(r)
}

/*
 * GetOperatorSchedulerConfiguration Method for GetOperatorSchedulerConfiguration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOperatorSchedulerConfigurationRequest
 */
func (a *OperatorApiService) GetOperatorSchedulerConfiguration(ctx _context.Context) ApiGetOperatorSchedulerConfigurationRequest {
	return ApiGetOperatorSchedulerConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SchedulerConfigurationResponse
 */
func (a *OperatorApiService) GetOperatorSchedulerConfigurationExecute(r ApiGetOperatorSchedulerConfigurationRequest) (SchedulerConfigurationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SchedulerConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.GetOperatorSchedulerConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/scheduler/configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.index != nil {
		localVarHeaderParams["index"] = parameterToString(*r.index, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostOperatorSchedulerConfigurationRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	schedulerConfiguration *SchedulerConfiguration
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiPostOperatorSchedulerConfigurationRequest) SchedulerConfiguration(schedulerConfiguration SchedulerConfiguration) ApiPostOperatorSchedulerConfigurationRequest {
	r.schedulerConfiguration = &schedulerConfiguration
	return r
}
func (r ApiPostOperatorSchedulerConfigurationRequest) Region(region string) ApiPostOperatorSchedulerConfigurationRequest {
	r.region = &region
	return r
}
func (r ApiPostOperatorSchedulerConfigurationRequest) Namespace(namespace string) ApiPostOperatorSchedulerConfigurationRequest {
	r.namespace = &namespace
	return r
}
func (r ApiPostOperatorSchedulerConfigurationRequest) XNomadToken(xNomadToken string) ApiPostOperatorSchedulerConfigurationRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiPostOperatorSchedulerConfigurationRequest) IdempotencyToken(idempotencyToken string) ApiPostOperatorSchedulerConfigurationRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiPostOperatorSchedulerConfigurationRequest) Execute() (SchedulerSetConfigurationResponse, *_nethttp.Response, error) {
	return r.ApiService.PostOperatorSchedulerConfigurationExecute(r)
}

/*
 * PostOperatorSchedulerConfiguration Method for PostOperatorSchedulerConfiguration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostOperatorSchedulerConfigurationRequest
 */
func (a *OperatorApiService) PostOperatorSchedulerConfiguration(ctx _context.Context) ApiPostOperatorSchedulerConfigurationRequest {
	return ApiPostOperatorSchedulerConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SchedulerSetConfigurationResponse
 */
func (a *OperatorApiService) PostOperatorSchedulerConfigurationExecute(r ApiPostOperatorSchedulerConfigurationRequest) (SchedulerSetConfigurationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SchedulerSetConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.PostOperatorSchedulerConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/scheduler/configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.schedulerConfiguration == nil {
		return localVarReturnValue, nil, reportError("schedulerConfiguration is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xNomadToken != nil {
		localVarHeaderParams["X-Nomad-Token"] = parameterToString(*r.xNomadToken, "")
	}
	// body params
	localVarPostBody = r.schedulerConfiguration
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutOperatorAutopilotConfigurationRequest struct {
	ctx _context.Context
	ApiService *OperatorApiService
	autopilotConfiguration *AutopilotConfiguration
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiPutOperatorAutopilotConfigurationRequest) AutopilotConfiguration(autopilotConfiguration AutopilotConfiguration) ApiPutOperatorAutopilotConfigurationRequest {
	r.autopilotConfiguration = &autopilotConfiguration
	return r
}
func (r ApiPutOperatorAutopilotConfigurationRequest) Region(region string) ApiPutOperatorAutopilotConfigurationRequest {
	r.region = &region
	return r
}
func (r ApiPutOperatorAutopilotConfigurationRequest) Namespace(namespace string) ApiPutOperatorAutopilotConfigurationRequest {
	r.namespace = &namespace
	return r
}
func (r ApiPutOperatorAutopilotConfigurationRequest) XNomadToken(xNomadToken string) ApiPutOperatorAutopilotConfigurationRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiPutOperatorAutopilotConfigurationRequest) IdempotencyToken(idempotencyToken string) ApiPutOperatorAutopilotConfigurationRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiPutOperatorAutopilotConfigurationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PutOperatorAutopilotConfigurationExecute(r)
}

/*
 * PutOperatorAutopilotConfiguration Method for PutOperatorAutopilotConfiguration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPutOperatorAutopilotConfigurationRequest
 */
func (a *OperatorApiService) PutOperatorAutopilotConfiguration(ctx _context.Context) ApiPutOperatorAutopilotConfigurationRequest {
	return ApiPutOperatorAutopilotConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *OperatorApiService) PutOperatorAutopilotConfigurationExecute(r ApiPutOperatorAutopilotConfigurationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorApiService.PutOperatorAutopilotConfiguration")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operator/autopilot/configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.autopilotConfiguration == nil {
		return nil, reportError("autopilotConfiguration is required and must be specified")
	}

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
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.autopilotConfiguration
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
