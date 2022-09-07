# DrainMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessorID** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDrainMetadata

`func NewDrainMetadata() *DrainMetadata`

NewDrainMetadata instantiates a new DrainMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrainMetadataWithDefaults

`func NewDrainMetadataWithDefaults() *DrainMetadata`

NewDrainMetadataWithDefaults instantiates a new DrainMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessorID

`func (o *DrainMetadata) GetAccessorID() string`

GetAccessorID returns the AccessorID field if non-nil, zero value otherwise.

### GetAccessorIDOk

`func (o *DrainMetadata) GetAccessorIDOk() (*string, bool)`

GetAccessorIDOk returns a tuple with the AccessorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessorID

`func (o *DrainMetadata) SetAccessorID(v string)`

SetAccessorID sets AccessorID field to given value.

### HasAccessorID

`func (o *DrainMetadata) HasAccessorID() bool`

HasAccessorID returns a boolean if a field has been set.

### GetMeta

`func (o *DrainMetadata) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DrainMetadata) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DrainMetadata) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DrainMetadata) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStartedAt

`func (o *DrainMetadata) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DrainMetadata) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DrainMetadata) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DrainMetadata) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *DrainMetadata) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *DrainMetadata) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *DrainMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrainMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrainMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrainMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DrainMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DrainMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DrainMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DrainMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DrainMetadata) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DrainMetadata) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


