# DeploymentPauseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentID** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Pause** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentPauseRequest

`func NewDeploymentPauseRequest() *DeploymentPauseRequest`

NewDeploymentPauseRequest instantiates a new DeploymentPauseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPauseRequestWithDefaults

`func NewDeploymentPauseRequestWithDefaults() *DeploymentPauseRequest`

NewDeploymentPauseRequestWithDefaults instantiates a new DeploymentPauseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentID

`func (o *DeploymentPauseRequest) GetDeploymentID() string`

GetDeploymentID returns the DeploymentID field if non-nil, zero value otherwise.

### GetDeploymentIDOk

`func (o *DeploymentPauseRequest) GetDeploymentIDOk() (*string, bool)`

GetDeploymentIDOk returns a tuple with the DeploymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentID

`func (o *DeploymentPauseRequest) SetDeploymentID(v string)`

SetDeploymentID sets DeploymentID field to given value.

### HasDeploymentID

`func (o *DeploymentPauseRequest) HasDeploymentID() bool`

HasDeploymentID returns a boolean if a field has been set.

### GetNamespace

`func (o *DeploymentPauseRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeploymentPauseRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeploymentPauseRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeploymentPauseRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPause

`func (o *DeploymentPauseRequest) GetPause() bool`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *DeploymentPauseRequest) GetPauseOk() (*bool, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *DeploymentPauseRequest) SetPause(v bool)`

SetPause sets Pause field to given value.

### HasPause

`func (o *DeploymentPauseRequest) HasPause() bool`

HasPause returns a boolean if a field has been set.

### GetRegion

`func (o *DeploymentPauseRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentPauseRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentPauseRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeploymentPauseRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *DeploymentPauseRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *DeploymentPauseRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *DeploymentPauseRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *DeploymentPauseRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


