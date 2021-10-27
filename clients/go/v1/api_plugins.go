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

// PluginsApiService PluginsApi service
type PluginsApiService service

type ApiGetPluginCSIRequest struct {
	ctx _context.Context
	ApiService *PluginsApiService
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	pluginID string
=======
	pluginId string
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
	pluginID string
>>>>>>> 13a3eee (added tests for plugins)
=======
	pluginID string
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
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

func (r ApiGetPluginCSIRequest) Region(region string) ApiGetPluginCSIRequest {
	r.region = &region
	return r
}
func (r ApiGetPluginCSIRequest) Namespace(namespace string) ApiGetPluginCSIRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetPluginCSIRequest) Index(index int32) ApiGetPluginCSIRequest {
	r.index = &index
	return r
}
func (r ApiGetPluginCSIRequest) Wait(wait string) ApiGetPluginCSIRequest {
	r.wait = &wait
	return r
}
func (r ApiGetPluginCSIRequest) Stale(stale string) ApiGetPluginCSIRequest {
	r.stale = &stale
	return r
}
func (r ApiGetPluginCSIRequest) Prefix(prefix string) ApiGetPluginCSIRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetPluginCSIRequest) XNomadToken(xNomadToken string) ApiGetPluginCSIRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetPluginCSIRequest) PerPage(perPage int32) ApiGetPluginCSIRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetPluginCSIRequest) NextToken(nextToken string) ApiGetPluginCSIRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetPluginCSIRequest) Execute() ([]CSIPlugin, *_nethttp.Response, error) {
	return r.ApiService.GetPluginCSIExecute(r)
}

/*
 * GetPluginCSI Method for GetPluginCSI
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
 * @param pluginID The CSI plugin identifier.
 * @return ApiGetPluginCSIRequest
 */
func (a *PluginsApiService) GetPluginCSI(ctx _context.Context, pluginID string) ApiGetPluginCSIRequest {
	return ApiGetPluginCSIRequest{
		ApiService: a,
		ctx: ctx,
		pluginID: pluginID,
=======
 * @param pluginId The CSI plugin identifier.
=======
 * @param pluginID The CSI plugin identifier.
>>>>>>> 13a3eee (added tests for plugins)
=======
 * @param pluginID The CSI plugin identifier.
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
 * @return ApiGetPluginCSIRequest
 */
func (a *PluginsApiService) GetPluginCSI(ctx _context.Context, pluginID string) ApiGetPluginCSIRequest {
	return ApiGetPluginCSIRequest{
		ApiService: a,
		ctx: ctx,
<<<<<<< HEAD
<<<<<<< HEAD
		pluginId: pluginId,
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
		pluginID: pluginID,
>>>>>>> 13a3eee (added tests for plugins)
=======
		pluginID: pluginID,
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
	}
}

/*
 * Execute executes the request
 * @return []CSIPlugin
 */
func (a *PluginsApiService) GetPluginCSIExecute(r ApiGetPluginCSIRequest) ([]CSIPlugin, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CSIPlugin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsApiService.GetPluginCSI")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/csi/{pluginID}"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	localVarPath = strings.Replace(localVarPath, "{"+"pluginID"+"}", _neturl.PathEscape(parameterToString(r.pluginID, "")), -1)
=======
	localVarPath = strings.Replace(localVarPath, "{"+"pluginId"+"}", _neturl.PathEscape(parameterToString(r.pluginId, "")), -1)
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
	localVarPath = strings.Replace(localVarPath, "{"+"pluginID"+"}", _neturl.PathEscape(parameterToString(r.pluginID, "")), -1)
>>>>>>> 13a3eee (added tests for plugins)
=======
	localVarPath = strings.Replace(localVarPath, "{"+"pluginID"+"}", _neturl.PathEscape(parameterToString(r.pluginID, "")), -1)
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b

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

type ApiGetPluginsRequest struct {
	ctx _context.Context
	ApiService *PluginsApiService
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

func (r ApiGetPluginsRequest) Region(region string) ApiGetPluginsRequest {
	r.region = &region
	return r
}
func (r ApiGetPluginsRequest) Namespace(namespace string) ApiGetPluginsRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetPluginsRequest) Index(index int32) ApiGetPluginsRequest {
	r.index = &index
	return r
}
func (r ApiGetPluginsRequest) Wait(wait string) ApiGetPluginsRequest {
	r.wait = &wait
	return r
}
func (r ApiGetPluginsRequest) Stale(stale string) ApiGetPluginsRequest {
	r.stale = &stale
	return r
}
func (r ApiGetPluginsRequest) Prefix(prefix string) ApiGetPluginsRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetPluginsRequest) XNomadToken(xNomadToken string) ApiGetPluginsRequest {
	r.xNomadToken = &xNomadToken
	return r
}
func (r ApiGetPluginsRequest) PerPage(perPage int32) ApiGetPluginsRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetPluginsRequest) NextToken(nextToken string) ApiGetPluginsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetPluginsRequest) Execute() ([]CSIPluginListStub, *_nethttp.Response, error) {
	return r.ApiService.GetPluginsExecute(r)
}

/*
 * GetPlugins Method for GetPlugins
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetPluginsRequest
 */
func (a *PluginsApiService) GetPlugins(ctx _context.Context) ApiGetPluginsRequest {
	return ApiGetPluginsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []CSIPluginListStub
 */
func (a *PluginsApiService) GetPluginsExecute(r ApiGetPluginsRequest) ([]CSIPluginListStub, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CSIPluginListStub
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsApiService.GetPlugins")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugins"

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
