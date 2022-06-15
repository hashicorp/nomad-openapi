# QuotaLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegionLimit** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 

## Methods

### NewQuotaLimit

`func NewQuotaLimit() *QuotaLimit`

NewQuotaLimit instantiates a new QuotaLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaLimitWithDefaults

`func NewQuotaLimitWithDefaults() *QuotaLimit`

NewQuotaLimitWithDefaults instantiates a new QuotaLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *QuotaLimit) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *QuotaLimit) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *QuotaLimit) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *QuotaLimit) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetRegion

`func (o *QuotaLimit) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *QuotaLimit) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *QuotaLimit) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *QuotaLimit) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionLimit

`func (o *QuotaLimit) GetRegionLimit() Resources`

GetRegionLimit returns the RegionLimit field if non-nil, zero value otherwise.

### GetRegionLimitOk

`func (o *QuotaLimit) GetRegionLimitOk() (*Resources, bool)`

GetRegionLimitOk returns a tuple with the RegionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLimit

`func (o *QuotaLimit) SetRegionLimit(v Resources)`

SetRegionLimit sets RegionLimit field to given value.

### HasRegionLimit

`func (o *QuotaLimit) HasRegionLimit() bool`

HasRegionLimit returns a boolean if a field has been set.

### SetRegionLimitNil

`func (o *QuotaLimit) SetRegionLimitNil(b bool)`

 SetRegionLimitNil sets the value for RegionLimit to be an explicit nil

### UnsetRegionLimit
`func (o *QuotaLimit) UnsetRegionLimit()`

UnsetRegionLimit ensures that no value is present for RegionLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


