# PointValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Points** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewPointValue

`func NewPointValue() *PointValue`

NewPointValue instantiates a new PointValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointValueWithDefaults

`func NewPointValueWithDefaults() *PointValue`

NewPointValueWithDefaults instantiates a new PointValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PointValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PointValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PointValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PointValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoints

`func (o *PointValue) GetPoints() []float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *PointValue) GetPointsOk() (*[]float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *PointValue) SetPoints(v []float32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *PointValue) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


