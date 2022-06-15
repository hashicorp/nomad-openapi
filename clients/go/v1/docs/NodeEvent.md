# NodeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Subsystem** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewNodeEvent

`func NewNodeEvent() *NodeEvent`

NewNodeEvent instantiates a new NodeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeEventWithDefaults

`func NewNodeEventWithDefaults() *NodeEvent`

NewNodeEventWithDefaults instantiates a new NodeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIndex

`func (o *NodeEvent) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *NodeEvent) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *NodeEvent) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *NodeEvent) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDetails

`func (o *NodeEvent) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NodeEvent) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NodeEvent) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NodeEvent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *NodeEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubsystem

`func (o *NodeEvent) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *NodeEvent) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *NodeEvent) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.

### HasSubsystem

`func (o *NodeEvent) HasSubsystem() bool`

HasSubsystem returns a boolean if a field has been set.

### GetTimestamp

`func (o *NodeEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NodeEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NodeEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NodeEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *NodeEvent) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *NodeEvent) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


