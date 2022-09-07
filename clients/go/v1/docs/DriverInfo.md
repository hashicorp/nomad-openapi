# DriverInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Detected** | Pointer to **bool** |  | [optional] 
**HealthDescription** | Pointer to **string** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDriverInfo

`func NewDriverInfo() *DriverInfo`

NewDriverInfo instantiates a new DriverInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverInfoWithDefaults

`func NewDriverInfoWithDefaults() *DriverInfo`

NewDriverInfoWithDefaults instantiates a new DriverInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *DriverInfo) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DriverInfo) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DriverInfo) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DriverInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDetected

`func (o *DriverInfo) GetDetected() bool`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *DriverInfo) GetDetectedOk() (*bool, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *DriverInfo) SetDetected(v bool)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *DriverInfo) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### GetHealthDescription

`func (o *DriverInfo) GetHealthDescription() string`

GetHealthDescription returns the HealthDescription field if non-nil, zero value otherwise.

### GetHealthDescriptionOk

`func (o *DriverInfo) GetHealthDescriptionOk() (*string, bool)`

GetHealthDescriptionOk returns a tuple with the HealthDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthDescription

`func (o *DriverInfo) SetHealthDescription(v string)`

SetHealthDescription sets HealthDescription field to given value.

### HasHealthDescription

`func (o *DriverInfo) HasHealthDescription() bool`

HasHealthDescription returns a boolean if a field has been set.

### GetHealthy

`func (o *DriverInfo) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *DriverInfo) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *DriverInfo) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *DriverInfo) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetUpdateTime

`func (o *DriverInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *DriverInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *DriverInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *DriverInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### SetUpdateTimeNil

`func (o *DriverInfo) SetUpdateTimeNil(b bool)`

 SetUpdateTimeNil sets the value for UpdateTime to be an explicit nil

### UnsetUpdateTime
`func (o *DriverInfo) UnsetUpdateTime()`

UnsetUpdateTime ensures that no value is present for UpdateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


