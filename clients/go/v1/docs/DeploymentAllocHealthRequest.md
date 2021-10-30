# DeploymentAllocHealthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentID** | Pointer to **string** |  | [optional] 
**HealthyAllocationIDs** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**UnhealthyAllocationIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeploymentAllocHealthRequest

`func NewDeploymentAllocHealthRequest() *DeploymentAllocHealthRequest`

NewDeploymentAllocHealthRequest instantiates a new DeploymentAllocHealthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentAllocHealthRequestWithDefaults

`func NewDeploymentAllocHealthRequestWithDefaults() *DeploymentAllocHealthRequest`

NewDeploymentAllocHealthRequestWithDefaults instantiates a new DeploymentAllocHealthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentID

`func (o *DeploymentAllocHealthRequest) GetDeploymentID() string`

GetDeploymentID returns the DeploymentID field if non-nil, zero value otherwise.

### GetDeploymentIDOk

`func (o *DeploymentAllocHealthRequest) GetDeploymentIDOk() (*string, bool)`

GetDeploymentIDOk returns a tuple with the DeploymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentID

`func (o *DeploymentAllocHealthRequest) SetDeploymentID(v string)`

SetDeploymentID sets DeploymentID field to given value.

### HasDeploymentID

`func (o *DeploymentAllocHealthRequest) HasDeploymentID() bool`

HasDeploymentID returns a boolean if a field has been set.

### GetHealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) GetHealthyAllocationIDs() []string`

GetHealthyAllocationIDs returns the HealthyAllocationIDs field if non-nil, zero value otherwise.

### GetHealthyAllocationIDsOk

`func (o *DeploymentAllocHealthRequest) GetHealthyAllocationIDsOk() (*[]string, bool)`

GetHealthyAllocationIDsOk returns a tuple with the HealthyAllocationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) SetHealthyAllocationIDs(v []string)`

SetHealthyAllocationIDs sets HealthyAllocationIDs field to given value.

### HasHealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) HasHealthyAllocationIDs() bool`

HasHealthyAllocationIDs returns a boolean if a field has been set.

### GetNamespace

`func (o *DeploymentAllocHealthRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeploymentAllocHealthRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeploymentAllocHealthRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeploymentAllocHealthRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *DeploymentAllocHealthRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentAllocHealthRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentAllocHealthRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeploymentAllocHealthRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *DeploymentAllocHealthRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *DeploymentAllocHealthRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *DeploymentAllocHealthRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *DeploymentAllocHealthRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetUnhealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) GetUnhealthyAllocationIDs() []string`

GetUnhealthyAllocationIDs returns the UnhealthyAllocationIDs field if non-nil, zero value otherwise.

### GetUnhealthyAllocationIDsOk

`func (o *DeploymentAllocHealthRequest) GetUnhealthyAllocationIDsOk() (*[]string, bool)`

GetUnhealthyAllocationIDsOk returns a tuple with the UnhealthyAllocationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) SetUnhealthyAllocationIDs(v []string)`

SetUnhealthyAllocationIDs sets UnhealthyAllocationIDs field to given value.

### HasUnhealthyAllocationIDs

`func (o *DeploymentAllocHealthRequest) HasUnhealthyAllocationIDs() bool`

HasUnhealthyAllocationIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


