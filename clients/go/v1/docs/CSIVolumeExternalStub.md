# CSIVolumeExternalStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityBytes** | Pointer to **int64** |  | [optional] 
**CloneID** | Pointer to **string** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**IsAbnormal** | Pointer to **bool** |  | [optional] 
**PublishedExternalNodeIDs** | Pointer to **[]string** |  | [optional] 
**SnapshotID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VolumeContext** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSIVolumeExternalStub

`func NewCSIVolumeExternalStub() *CSIVolumeExternalStub`

NewCSIVolumeExternalStub instantiates a new CSIVolumeExternalStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeExternalStubWithDefaults

`func NewCSIVolumeExternalStubWithDefaults() *CSIVolumeExternalStub`

NewCSIVolumeExternalStubWithDefaults instantiates a new CSIVolumeExternalStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityBytes

`func (o *CSIVolumeExternalStub) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *CSIVolumeExternalStub) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *CSIVolumeExternalStub) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *CSIVolumeExternalStub) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### GetCloneID

`func (o *CSIVolumeExternalStub) GetCloneID() string`

GetCloneID returns the CloneID field if non-nil, zero value otherwise.

### GetCloneIDOk

`func (o *CSIVolumeExternalStub) GetCloneIDOk() (*string, bool)`

GetCloneIDOk returns a tuple with the CloneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneID

`func (o *CSIVolumeExternalStub) SetCloneID(v string)`

SetCloneID sets CloneID field to given value.

### HasCloneID

`func (o *CSIVolumeExternalStub) HasCloneID() bool`

HasCloneID returns a boolean if a field has been set.

### GetExternalID

`func (o *CSIVolumeExternalStub) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *CSIVolumeExternalStub) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *CSIVolumeExternalStub) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *CSIVolumeExternalStub) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetIsAbnormal

`func (o *CSIVolumeExternalStub) GetIsAbnormal() bool`

GetIsAbnormal returns the IsAbnormal field if non-nil, zero value otherwise.

### GetIsAbnormalOk

`func (o *CSIVolumeExternalStub) GetIsAbnormalOk() (*bool, bool)`

GetIsAbnormalOk returns a tuple with the IsAbnormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbnormal

`func (o *CSIVolumeExternalStub) SetIsAbnormal(v bool)`

SetIsAbnormal sets IsAbnormal field to given value.

### HasIsAbnormal

`func (o *CSIVolumeExternalStub) HasIsAbnormal() bool`

HasIsAbnormal returns a boolean if a field has been set.

### GetPublishedExternalNodeIDs

`func (o *CSIVolumeExternalStub) GetPublishedExternalNodeIDs() []string`

GetPublishedExternalNodeIDs returns the PublishedExternalNodeIDs field if non-nil, zero value otherwise.

### GetPublishedExternalNodeIDsOk

`func (o *CSIVolumeExternalStub) GetPublishedExternalNodeIDsOk() (*[]string, bool)`

GetPublishedExternalNodeIDsOk returns a tuple with the PublishedExternalNodeIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedExternalNodeIDs

`func (o *CSIVolumeExternalStub) SetPublishedExternalNodeIDs(v []string)`

SetPublishedExternalNodeIDs sets PublishedExternalNodeIDs field to given value.

### HasPublishedExternalNodeIDs

`func (o *CSIVolumeExternalStub) HasPublishedExternalNodeIDs() bool`

HasPublishedExternalNodeIDs returns a boolean if a field has been set.

### GetSnapshotID

`func (o *CSIVolumeExternalStub) GetSnapshotID() string`

GetSnapshotID returns the SnapshotID field if non-nil, zero value otherwise.

### GetSnapshotIDOk

`func (o *CSIVolumeExternalStub) GetSnapshotIDOk() (*string, bool)`

GetSnapshotIDOk returns a tuple with the SnapshotID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotID

`func (o *CSIVolumeExternalStub) SetSnapshotID(v string)`

SetSnapshotID sets SnapshotID field to given value.

### HasSnapshotID

`func (o *CSIVolumeExternalStub) HasSnapshotID() bool`

HasSnapshotID returns a boolean if a field has been set.

### GetStatus

`func (o *CSIVolumeExternalStub) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CSIVolumeExternalStub) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CSIVolumeExternalStub) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CSIVolumeExternalStub) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeContext

`func (o *CSIVolumeExternalStub) GetVolumeContext() map[string]string`

GetVolumeContext returns the VolumeContext field if non-nil, zero value otherwise.

### GetVolumeContextOk

`func (o *CSIVolumeExternalStub) GetVolumeContextOk() (*map[string]string, bool)`

GetVolumeContextOk returns a tuple with the VolumeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeContext

`func (o *CSIVolumeExternalStub) SetVolumeContext(v map[string]string)`

SetVolumeContext sets VolumeContext field to given value.

### HasVolumeContext

`func (o *CSIVolumeExternalStub) HasVolumeContext() bool`

HasVolumeContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


