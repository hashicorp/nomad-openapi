# ConsulGatewayTLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherSuites** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TLSMaxVersion** | Pointer to **string** |  | [optional] 
**TLSMinVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewConsulGatewayTLSConfig

`func NewConsulGatewayTLSConfig() *ConsulGatewayTLSConfig`

NewConsulGatewayTLSConfig instantiates a new ConsulGatewayTLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsulGatewayTLSConfigWithDefaults

`func NewConsulGatewayTLSConfigWithDefaults() *ConsulGatewayTLSConfig`

NewConsulGatewayTLSConfigWithDefaults instantiates a new ConsulGatewayTLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherSuites

`func (o *ConsulGatewayTLSConfig) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *ConsulGatewayTLSConfig) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *ConsulGatewayTLSConfig) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *ConsulGatewayTLSConfig) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetEnabled

`func (o *ConsulGatewayTLSConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConsulGatewayTLSConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConsulGatewayTLSConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConsulGatewayTLSConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTLSMaxVersion

`func (o *ConsulGatewayTLSConfig) GetTLSMaxVersion() string`

GetTLSMaxVersion returns the TLSMaxVersion field if non-nil, zero value otherwise.

### GetTLSMaxVersionOk

`func (o *ConsulGatewayTLSConfig) GetTLSMaxVersionOk() (*string, bool)`

GetTLSMaxVersionOk returns a tuple with the TLSMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSMaxVersion

`func (o *ConsulGatewayTLSConfig) SetTLSMaxVersion(v string)`

SetTLSMaxVersion sets TLSMaxVersion field to given value.

### HasTLSMaxVersion

`func (o *ConsulGatewayTLSConfig) HasTLSMaxVersion() bool`

HasTLSMaxVersion returns a boolean if a field has been set.

### GetTLSMinVersion

`func (o *ConsulGatewayTLSConfig) GetTLSMinVersion() string`

GetTLSMinVersion returns the TLSMinVersion field if non-nil, zero value otherwise.

### GetTLSMinVersionOk

`func (o *ConsulGatewayTLSConfig) GetTLSMinVersionOk() (*string, bool)`

GetTLSMinVersionOk returns a tuple with the TLSMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSMinVersion

`func (o *ConsulGatewayTLSConfig) SetTLSMinVersion(v string)`

SetTLSMinVersion sets TLSMinVersion field to given value.

### HasTLSMinVersion

`func (o *ConsulGatewayTLSConfig) HasTLSMinVersion() bool`

HasTLSMinVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


