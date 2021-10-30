# DeploymentPromoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** |  | [optional] 
**DeploymentID** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentPromoteRequest

`func NewDeploymentPromoteRequest() *DeploymentPromoteRequest`

NewDeploymentPromoteRequest instantiates a new DeploymentPromoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPromoteRequestWithDefaults

`func NewDeploymentPromoteRequestWithDefaults() *DeploymentPromoteRequest`

NewDeploymentPromoteRequestWithDefaults instantiates a new DeploymentPromoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *DeploymentPromoteRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeploymentPromoteRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeploymentPromoteRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *DeploymentPromoteRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetDeploymentID

`func (o *DeploymentPromoteRequest) GetDeploymentID() string`

GetDeploymentID returns the DeploymentID field if non-nil, zero value otherwise.

### GetDeploymentIDOk

`func (o *DeploymentPromoteRequest) GetDeploymentIDOk() (*string, bool)`

GetDeploymentIDOk returns a tuple with the DeploymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentID

`func (o *DeploymentPromoteRequest) SetDeploymentID(v string)`

SetDeploymentID sets DeploymentID field to given value.

### HasDeploymentID

`func (o *DeploymentPromoteRequest) HasDeploymentID() bool`

HasDeploymentID returns a boolean if a field has been set.

### GetGroups

`func (o *DeploymentPromoteRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DeploymentPromoteRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DeploymentPromoteRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DeploymentPromoteRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetNamespace

`func (o *DeploymentPromoteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeploymentPromoteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeploymentPromoteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeploymentPromoteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *DeploymentPromoteRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentPromoteRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentPromoteRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeploymentPromoteRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *DeploymentPromoteRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *DeploymentPromoteRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *DeploymentPromoteRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *DeploymentPromoteRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


