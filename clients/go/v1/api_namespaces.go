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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// NamespacesApiService NamespacesApi service
type NamespacesApiService service

type ApiCreateNamespaceRequest struct {
	ctx _context.Context
	ApiService *NamespacesApiService
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiCreateNamespaceRequest) Region(region string) ApiCreateNamespaceRequest {
	r.region = &region
	return r
}
func (r ApiCreateNamespaceRequest) Namespace(namespace string) ApiCreateNamespaceRequest {
	r.namespace = &namespace
	return r
}
func (r ApiCreateNamespaceRequest) XNomadToken(xNomadToken string) ApiCreateNamespaceRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiCreateNamespaceRequest) IdempotencyToken(idempotencyToken string) ApiCreateNamespaceRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiCreateNamespaceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateNamespaceExecute(r)
}

/*
 * CreateNamespace Method for CreateNamespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateNamespaceRequest
 */
func (a *NamespacesApiService) CreateNamespace(ctx _context.Context) ApiCreateNamespaceRequest {
	return ApiCreateNamespaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *NamespacesApiService) CreateNamespaceExecute(r ApiCreateNamespaceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespacesApiService.CreateNamespace")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespace"

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

type ApiDeleteNamespaceRequest struct {
	ctx _context.Context
	ApiService *NamespacesApiService
	namespaceName string
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiDeleteNamespaceRequest) Region(region string) ApiDeleteNamespaceRequest {
	r.region = &region
	return r
}
func (r ApiDeleteNamespaceRequest) Namespace(namespace string) ApiDeleteNamespaceRequest {
	r.namespace = &namespace
	return r
}
func (r ApiDeleteNamespaceRequest) XNomadToken(xNomadToken string) ApiDeleteNamespaceRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiDeleteNamespaceRequest) IdempotencyToken(idempotencyToken string) ApiDeleteNamespaceRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiDeleteNamespaceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNamespaceExecute(r)
}

/*
 * DeleteNamespace Method for DeleteNamespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param namespaceName The namespace identifier.
 * @return ApiDeleteNamespaceRequest
 */
func (a *NamespacesApiService) DeleteNamespace(ctx _context.Context, namespaceName string) ApiDeleteNamespaceRequest {
	return ApiDeleteNamespaceRequest{
		ApiService: a,
		ctx: ctx,
		namespaceName: namespaceName,
	}
}

/*
 * Execute executes the request
 */
func (a *NamespacesApiService) DeleteNamespaceExecute(r ApiDeleteNamespaceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespacesApiService.DeleteNamespace")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespace/{namespaceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceName"+"}", _neturl.PathEscape(parameterToString(r.namespaceName, "")), -1)

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

type ApiGetNamespaceRequest struct {
	ctx _context.Context
	ApiService *NamespacesApiService
	namespaceName string
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

func (r ApiGetNamespaceRequest) Region(region string) ApiGetNamespaceRequest {
	r.region = &region
	return r
}
func (r ApiGetNamespaceRequest) Namespace(namespace string) ApiGetNamespaceRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetNamespaceRequest) Index(index int32) ApiGetNamespaceRequest {
	r.index = &index
	return r
}
func (r ApiGetNamespaceRequest) Wait(wait string) ApiGetNamespaceRequest {
	r.wait = &wait
	return r
}
func (r ApiGetNamespaceRequest) Stale(stale string) ApiGetNamespaceRequest {
	r.stale = &stale
	return r
}
func (r ApiGetNamespaceRequest) Prefix(prefix string) ApiGetNamespaceRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetNamespaceRequest) XNomadToken(xNomadToken string) ApiGetNamespaceRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetNamespaceRequest) PerPage(perPage int32) ApiGetNamespaceRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetNamespaceRequest) NextToken(nextToken string) ApiGetNamespaceRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetNamespaceRequest) Execute() (Namespace, *_nethttp.Response, error) {
	return r.ApiService.GetNamespaceExecute(r)
}

/*
 * GetNamespace Method for GetNamespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param namespaceName The namespace identifier.
 * @return ApiGetNamespaceRequest
 */
func (a *NamespacesApiService) GetNamespace(ctx _context.Context, namespaceName string) ApiGetNamespaceRequest {
	return ApiGetNamespaceRequest{
		ApiService: a,
		ctx: ctx,
		namespaceName: namespaceName,
	}
}

/*
 * Execute executes the request
 * @return Namespace
 */
func (a *NamespacesApiService) GetNamespaceExecute(r ApiGetNamespaceRequest) (Namespace, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Namespace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespacesApiService.GetNamespace")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespace/{namespaceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceName"+"}", _neturl.PathEscape(parameterToString(r.namespaceName, "")), -1)

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

type ApiGetNamespacesRequest struct {
	ctx _context.Context
	ApiService *NamespacesApiService
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

func (r ApiGetNamespacesRequest) Region(region string) ApiGetNamespacesRequest {
	r.region = &region
	return r
}
func (r ApiGetNamespacesRequest) Namespace(namespace string) ApiGetNamespacesRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetNamespacesRequest) Index(index int32) ApiGetNamespacesRequest {
	r.index = &index
	return r
}
func (r ApiGetNamespacesRequest) Wait(wait string) ApiGetNamespacesRequest {
	r.wait = &wait
	return r
}
func (r ApiGetNamespacesRequest) Stale(stale string) ApiGetNamespacesRequest {
	r.stale = &stale
	return r
}
func (r ApiGetNamespacesRequest) Prefix(prefix string) ApiGetNamespacesRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetNamespacesRequest) XNomadToken(xNomadToken string) ApiGetNamespacesRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetNamespacesRequest) PerPage(perPage int32) ApiGetNamespacesRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetNamespacesRequest) NextToken(nextToken string) ApiGetNamespacesRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetNamespacesRequest) Execute() ([]Namespace, *_nethttp.Response, error) {
	return r.ApiService.GetNamespacesExecute(r)
}

/*
 * GetNamespaces Method for GetNamespaces
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetNamespacesRequest
 */
func (a *NamespacesApiService) GetNamespaces(ctx _context.Context) ApiGetNamespacesRequest {
	return ApiGetNamespacesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Namespace
 */
func (a *NamespacesApiService) GetNamespacesExecute(r ApiGetNamespacesRequest) ([]Namespace, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Namespace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespacesApiService.GetNamespaces")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespaces"

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

type ApiPostNamespaceRequest struct {
	ctx _context.Context
	ApiService *NamespacesApiService
	namespaceName string
	namespace2 *Namespace
	region *string
	namespace *string
	xNomadToken *string
	idempotencyToken *string
}

func (r ApiPostNamespaceRequest) Namespace2(namespace2 Namespace) ApiPostNamespaceRequest {
	r.namespace2 = &namespace2
	return r
}
func (r ApiPostNamespaceRequest) Region(region string) ApiPostNamespaceRequest {
	r.region = &region
	return r
}
func (r ApiPostNamespaceRequest) Namespace(namespace string) ApiPostNamespaceRequest {
	r.namespace = &namespace
	return r
}
func (r ApiPostNamespaceRequest) XNomadToken(xNomadToken string) ApiPostNamespaceRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiPostNamespaceRequest) IdempotencyToken(idempotencyToken string) ApiPostNamespaceRequest {
	r.idempotencyToken = &idempotencyToken
	return r
}

func (r ApiPostNamespaceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PostNamespaceExecute(r)
}

/*
 * PostNamespace Method for PostNamespace
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param namespaceName The namespace identifier.
 * @return ApiPostNamespaceRequest
 */
func (a *NamespacesApiService) PostNamespace(ctx _context.Context, namespaceName string) ApiPostNamespaceRequest {
	return ApiPostNamespaceRequest{
		ApiService: a,
		ctx: ctx,
		namespaceName: namespaceName,
	}
}

/*
 * Execute executes the request
 */
func (a *NamespacesApiService) PostNamespaceExecute(r ApiPostNamespaceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespacesApiService.PostNamespace")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespace/{namespaceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceName"+"}", _neturl.PathEscape(parameterToString(r.namespaceName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.namespace2 == nil {
		return nil, reportError("namespace2 is required and must be specified")
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
	localVarPostBody = r.namespace2
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
