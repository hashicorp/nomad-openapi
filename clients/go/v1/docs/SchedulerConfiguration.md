# SchedulerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIndex** | Pointer to **int32** |  | [optional] 
**MemoryOversubscriptionEnabled** | Pointer to **bool** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**PreemptionConfig** | Pointer to [**PreemptionConfig**](PreemptionConfig.md) |  | [optional] 
**RejectJobRegistration** | Pointer to **bool** |  | [optional] 
**SchedulerAlgorithm** | Pointer to **string** |  | [optional] 

## Methods

### NewSchedulerConfiguration

`func NewSchedulerConfiguration() *SchedulerConfiguration`

NewSchedulerConfiguration instantiates a new SchedulerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerConfigurationWithDefaults

`func NewSchedulerConfigurationWithDefaults() *SchedulerConfiguration`

NewSchedulerConfigurationWithDefaults instantiates a new SchedulerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIndex

`func (o *SchedulerConfiguration) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *SchedulerConfiguration) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *SchedulerConfiguration) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *SchedulerConfiguration) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetMemoryOversubscriptionEnabled

`func (o *SchedulerConfiguration) GetMemoryOversubscriptionEnabled() bool`

GetMemoryOversubscriptionEnabled returns the MemoryOversubscriptionEnabled field if non-nil, zero value otherwise.

### GetMemoryOversubscriptionEnabledOk

`func (o *SchedulerConfiguration) GetMemoryOversubscriptionEnabledOk() (*bool, bool)`

GetMemoryOversubscriptionEnabledOk returns a tuple with the MemoryOversubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOversubscriptionEnabled

`func (o *SchedulerConfiguration) SetMemoryOversubscriptionEnabled(v bool)`

SetMemoryOversubscriptionEnabled sets MemoryOversubscriptionEnabled field to given value.

### HasMemoryOversubscriptionEnabled

`func (o *SchedulerConfiguration) HasMemoryOversubscriptionEnabled() bool`

HasMemoryOversubscriptionEnabled returns a boolean if a field has been set.

### GetModifyIndex

`func (o *SchedulerConfiguration) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *SchedulerConfiguration) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *SchedulerConfiguration) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *SchedulerConfiguration) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetPreemptionConfig

`func (o *SchedulerConfiguration) GetPreemptionConfig() PreemptionConfig`

GetPreemptionConfig returns the PreemptionConfig field if non-nil, zero value otherwise.

### GetPreemptionConfigOk

`func (o *SchedulerConfiguration) GetPreemptionConfigOk() (*PreemptionConfig, bool)`

GetPreemptionConfigOk returns a tuple with the PreemptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptionConfig

`func (o *SchedulerConfiguration) SetPreemptionConfig(v PreemptionConfig)`

SetPreemptionConfig sets PreemptionConfig field to given value.

### HasPreemptionConfig

`func (o *SchedulerConfiguration) HasPreemptionConfig() bool`

HasPreemptionConfig returns a boolean if a field has been set.

### GetRejectJobRegistration

`func (o *SchedulerConfiguration) GetRejectJobRegistration() bool`

GetRejectJobRegistration returns the RejectJobRegistration field if non-nil, zero value otherwise.

### GetRejectJobRegistrationOk

`func (o *SchedulerConfiguration) GetRejectJobRegistrationOk() (*bool, bool)`

GetRejectJobRegistrationOk returns a tuple with the RejectJobRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectJobRegistration

`func (o *SchedulerConfiguration) SetRejectJobRegistration(v bool)`

SetRejectJobRegistration sets RejectJobRegistration field to given value.

### HasRejectJobRegistration

`func (o *SchedulerConfiguration) HasRejectJobRegistration() bool`

HasRejectJobRegistration returns a boolean if a field has been set.

### GetSchedulerAlgorithm

`func (o *SchedulerConfiguration) GetSchedulerAlgorithm() string`

GetSchedulerAlgorithm returns the SchedulerAlgorithm field if non-nil, zero value otherwise.

### GetSchedulerAlgorithmOk

`func (o *SchedulerConfiguration) GetSchedulerAlgorithmOk() (*string, bool)`

GetSchedulerAlgorithmOk returns a tuple with the SchedulerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerAlgorithm

`func (o *SchedulerConfiguration) SetSchedulerAlgorithm(v string)`

SetSchedulerAlgorithm sets SchedulerAlgorithm field to given value.

### HasSchedulerAlgorithm

`func (o *SchedulerConfiguration) HasSchedulerAlgorithm() bool`

HasSchedulerAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


