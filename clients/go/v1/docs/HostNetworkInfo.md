# HostNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CIDR** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReservedPorts** | Pointer to **string** |  | [optional] 

## Methods

### NewHostNetworkInfo

`func NewHostNetworkInfo() *HostNetworkInfo`

NewHostNetworkInfo instantiates a new HostNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostNetworkInfoWithDefaults

`func NewHostNetworkInfoWithDefaults() *HostNetworkInfo`

NewHostNetworkInfoWithDefaults instantiates a new HostNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCIDR

`func (o *HostNetworkInfo) GetCIDR() string`

GetCIDR returns the CIDR field if non-nil, zero value otherwise.

### GetCIDROk

`func (o *HostNetworkInfo) GetCIDROk() (*string, bool)`

GetCIDROk returns a tuple with the CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIDR

`func (o *HostNetworkInfo) SetCIDR(v string)`

SetCIDR sets CIDR field to given value.

### HasCIDR

`func (o *HostNetworkInfo) HasCIDR() bool`

HasCIDR returns a boolean if a field has been set.

### GetInterface

`func (o *HostNetworkInfo) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *HostNetworkInfo) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *HostNetworkInfo) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *HostNetworkInfo) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetName

`func (o *HostNetworkInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostNetworkInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostNetworkInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostNetworkInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReservedPorts

`func (o *HostNetworkInfo) GetReservedPorts() string`

GetReservedPorts returns the ReservedPorts field if non-nil, zero value otherwise.

### GetReservedPortsOk

`func (o *HostNetworkInfo) GetReservedPortsOk() (*string, bool)`

GetReservedPortsOk returns a tuple with the ReservedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPorts

`func (o *HostNetworkInfo) SetReservedPorts(v string)`

SetReservedPorts sets ReservedPorts field to given value.

### HasReservedPorts

`func (o *HostNetworkInfo) HasReservedPorts() bool`

HasReservedPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


