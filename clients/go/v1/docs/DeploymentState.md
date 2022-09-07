# DeploymentState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRevert** | Pointer to **bool** |  | [optional] 
**DesiredCanaries** | Pointer to **int32** |  | [optional] 
**DesiredTotal** | Pointer to **int32** |  | [optional] 
**HealthyAllocs** | Pointer to **int32** |  | [optional] 
**PlacedAllocs** | Pointer to **int32** |  | [optional] 
**PlacedCanaries** | Pointer to **[]string** |  | [optional] 
**ProgressDeadline** | Pointer to **int64** |  | [optional] 
**Promoted** | Pointer to **bool** |  | [optional] 
**RequireProgressBy** | Pointer to **NullableTime** |  | [optional] 
**UnhealthyAllocs** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeploymentState

`func NewDeploymentState() *DeploymentState`

NewDeploymentState instantiates a new DeploymentState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateWithDefaults

`func NewDeploymentStateWithDefaults() *DeploymentState`

NewDeploymentStateWithDefaults instantiates a new DeploymentState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRevert

`func (o *DeploymentState) GetAutoRevert() bool`

GetAutoRevert returns the AutoRevert field if non-nil, zero value otherwise.

### GetAutoRevertOk

`func (o *DeploymentState) GetAutoRevertOk() (*bool, bool)`

GetAutoRevertOk returns a tuple with the AutoRevert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRevert

`func (o *DeploymentState) SetAutoRevert(v bool)`

SetAutoRevert sets AutoRevert field to given value.

### HasAutoRevert

`func (o *DeploymentState) HasAutoRevert() bool`

HasAutoRevert returns a boolean if a field has been set.

### GetDesiredCanaries

`func (o *DeploymentState) GetDesiredCanaries() int32`

GetDesiredCanaries returns the DesiredCanaries field if non-nil, zero value otherwise.

### GetDesiredCanariesOk

`func (o *DeploymentState) GetDesiredCanariesOk() (*int32, bool)`

GetDesiredCanariesOk returns a tuple with the DesiredCanaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCanaries

`func (o *DeploymentState) SetDesiredCanaries(v int32)`

SetDesiredCanaries sets DesiredCanaries field to given value.

### HasDesiredCanaries

`func (o *DeploymentState) HasDesiredCanaries() bool`

HasDesiredCanaries returns a boolean if a field has been set.

### GetDesiredTotal

`func (o *DeploymentState) GetDesiredTotal() int32`

GetDesiredTotal returns the DesiredTotal field if non-nil, zero value otherwise.

### GetDesiredTotalOk

`func (o *DeploymentState) GetDesiredTotalOk() (*int32, bool)`

GetDesiredTotalOk returns a tuple with the DesiredTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTotal

`func (o *DeploymentState) SetDesiredTotal(v int32)`

SetDesiredTotal sets DesiredTotal field to given value.

### HasDesiredTotal

`func (o *DeploymentState) HasDesiredTotal() bool`

HasDesiredTotal returns a boolean if a field has been set.

### GetHealthyAllocs

`func (o *DeploymentState) GetHealthyAllocs() int32`

GetHealthyAllocs returns the HealthyAllocs field if non-nil, zero value otherwise.

### GetHealthyAllocsOk

`func (o *DeploymentState) GetHealthyAllocsOk() (*int32, bool)`

GetHealthyAllocsOk returns a tuple with the HealthyAllocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyAllocs

`func (o *DeploymentState) SetHealthyAllocs(v int32)`

SetHealthyAllocs sets HealthyAllocs field to given value.

### HasHealthyAllocs

`func (o *DeploymentState) HasHealthyAllocs() bool`

HasHealthyAllocs returns a boolean if a field has been set.

### GetPlacedAllocs

`func (o *DeploymentState) GetPlacedAllocs() int32`

GetPlacedAllocs returns the PlacedAllocs field if non-nil, zero value otherwise.

### GetPlacedAllocsOk

`func (o *DeploymentState) GetPlacedAllocsOk() (*int32, bool)`

GetPlacedAllocsOk returns a tuple with the PlacedAllocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAllocs

`func (o *DeploymentState) SetPlacedAllocs(v int32)`

SetPlacedAllocs sets PlacedAllocs field to given value.

### HasPlacedAllocs

`func (o *DeploymentState) HasPlacedAllocs() bool`

HasPlacedAllocs returns a boolean if a field has been set.

### GetPlacedCanaries

`func (o *DeploymentState) GetPlacedCanaries() []string`

GetPlacedCanaries returns the PlacedCanaries field if non-nil, zero value otherwise.

### GetPlacedCanariesOk

`func (o *DeploymentState) GetPlacedCanariesOk() (*[]string, bool)`

GetPlacedCanariesOk returns a tuple with the PlacedCanaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedCanaries

`func (o *DeploymentState) SetPlacedCanaries(v []string)`

SetPlacedCanaries sets PlacedCanaries field to given value.

### HasPlacedCanaries

`func (o *DeploymentState) HasPlacedCanaries() bool`

HasPlacedCanaries returns a boolean if a field has been set.

### GetProgressDeadline

`func (o *DeploymentState) GetProgressDeadline() int64`

GetProgressDeadline returns the ProgressDeadline field if non-nil, zero value otherwise.

### GetProgressDeadlineOk

`func (o *DeploymentState) GetProgressDeadlineOk() (*int64, bool)`

GetProgressDeadlineOk returns a tuple with the ProgressDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadline

`func (o *DeploymentState) SetProgressDeadline(v int64)`

SetProgressDeadline sets ProgressDeadline field to given value.

### HasProgressDeadline

`func (o *DeploymentState) HasProgressDeadline() bool`

HasProgressDeadline returns a boolean if a field has been set.

### GetPromoted

`func (o *DeploymentState) GetPromoted() bool`

GetPromoted returns the Promoted field if non-nil, zero value otherwise.

### GetPromotedOk

`func (o *DeploymentState) GetPromotedOk() (*bool, bool)`

GetPromotedOk returns a tuple with the Promoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoted

`func (o *DeploymentState) SetPromoted(v bool)`

SetPromoted sets Promoted field to given value.

### HasPromoted

`func (o *DeploymentState) HasPromoted() bool`

HasPromoted returns a boolean if a field has been set.

### GetRequireProgressBy

`func (o *DeploymentState) GetRequireProgressBy() time.Time`

GetRequireProgressBy returns the RequireProgressBy field if non-nil, zero value otherwise.

### GetRequireProgressByOk

`func (o *DeploymentState) GetRequireProgressByOk() (*time.Time, bool)`

GetRequireProgressByOk returns a tuple with the RequireProgressBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProgressBy

`func (o *DeploymentState) SetRequireProgressBy(v time.Time)`

SetRequireProgressBy sets RequireProgressBy field to given value.

### HasRequireProgressBy

`func (o *DeploymentState) HasRequireProgressBy() bool`

HasRequireProgressBy returns a boolean if a field has been set.

### SetRequireProgressByNil

`func (o *DeploymentState) SetRequireProgressByNil(b bool)`

 SetRequireProgressByNil sets the value for RequireProgressBy to be an explicit nil

### UnsetRequireProgressBy
`func (o *DeploymentState) UnsetRequireProgressBy()`

UnsetRequireProgressBy ensures that no value is present for RequireProgressBy, not even an explicit nil
### GetUnhealthyAllocs

`func (o *DeploymentState) GetUnhealthyAllocs() int32`

GetUnhealthyAllocs returns the UnhealthyAllocs field if non-nil, zero value otherwise.

### GetUnhealthyAllocsOk

`func (o *DeploymentState) GetUnhealthyAllocsOk() (*int32, bool)`

GetUnhealthyAllocsOk returns a tuple with the UnhealthyAllocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyAllocs

`func (o *DeploymentState) SetUnhealthyAllocs(v int32)`

SetUnhealthyAllocs sets UnhealthyAllocs field to given value.

### HasUnhealthyAllocs

`func (o *DeploymentState) HasUnhealthyAllocs() bool`

HasUnhealthyAllocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


