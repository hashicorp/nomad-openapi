# NodeCpuResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuShares** | Pointer to **int64** |  | [optional] 
**ReservableCpuCores** | Pointer to **[]int32** |  | [optional] 
**TotalCpuCores** | Pointer to **int32** |  | [optional] 

## Methods

### NewNodeCpuResources

`func NewNodeCpuResources() *NodeCpuResources`

NewNodeCpuResources instantiates a new NodeCpuResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCpuResourcesWithDefaults

`func NewNodeCpuResourcesWithDefaults() *NodeCpuResources`

NewNodeCpuResourcesWithDefaults instantiates a new NodeCpuResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuShares

`func (o *NodeCpuResources) GetCpuShares() int64`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *NodeCpuResources) GetCpuSharesOk() (*int64, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *NodeCpuResources) SetCpuShares(v int64)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *NodeCpuResources) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.

### GetReservableCpuCores

`func (o *NodeCpuResources) GetReservableCpuCores() []int32`

GetReservableCpuCores returns the ReservableCpuCores field if non-nil, zero value otherwise.

### GetReservableCpuCoresOk

`func (o *NodeCpuResources) GetReservableCpuCoresOk() (*[]int32, bool)`

GetReservableCpuCoresOk returns a tuple with the ReservableCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservableCpuCores

`func (o *NodeCpuResources) SetReservableCpuCores(v []int32)`

SetReservableCpuCores sets ReservableCpuCores field to given value.

### HasReservableCpuCores

`func (o *NodeCpuResources) HasReservableCpuCores() bool`

HasReservableCpuCores returns a boolean if a field has been set.

### GetTotalCpuCores

`func (o *NodeCpuResources) GetTotalCpuCores() int32`

GetTotalCpuCores returns the TotalCpuCores field if non-nil, zero value otherwise.

### GetTotalCpuCoresOk

`func (o *NodeCpuResources) GetTotalCpuCoresOk() (*int32, bool)`

GetTotalCpuCoresOk returns a tuple with the TotalCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuCores

`func (o *NodeCpuResources) SetTotalCpuCores(v int32)`

SetTotalCpuCores sets TotalCpuCores field to given value.

### HasTotalCpuCores

`func (o *NodeCpuResources) HasTotalCpuCores() bool`

HasTotalCpuCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


