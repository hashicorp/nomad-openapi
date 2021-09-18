# MetricsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**[]SampledValue**](SampledValue.md) |  | [optional] 
**Gauges** | Pointer to [**[]GaugeValue**](GaugeValue.md) |  | [optional] 
**Points** | Pointer to [**[]PointValue**](PointValue.md) |  | [optional] 
**Samples** | Pointer to [**[]SampledValue**](SampledValue.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricsSummary

`func NewMetricsSummary() *MetricsSummary`

NewMetricsSummary instantiates a new MetricsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSummaryWithDefaults

`func NewMetricsSummaryWithDefaults() *MetricsSummary`

NewMetricsSummaryWithDefaults instantiates a new MetricsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *MetricsSummary) GetCounters() []SampledValue`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *MetricsSummary) GetCountersOk() (*[]SampledValue, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *MetricsSummary) SetCounters(v []SampledValue)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *MetricsSummary) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetGauges

`func (o *MetricsSummary) GetGauges() []GaugeValue`

GetGauges returns the Gauges field if non-nil, zero value otherwise.

### GetGaugesOk

`func (o *MetricsSummary) GetGaugesOk() (*[]GaugeValue, bool)`

GetGaugesOk returns a tuple with the Gauges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGauges

`func (o *MetricsSummary) SetGauges(v []GaugeValue)`

SetGauges sets Gauges field to given value.

### HasGauges

`func (o *MetricsSummary) HasGauges() bool`

HasGauges returns a boolean if a field has been set.

### GetPoints

`func (o *MetricsSummary) GetPoints() []PointValue`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *MetricsSummary) GetPointsOk() (*[]PointValue, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *MetricsSummary) SetPoints(v []PointValue)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *MetricsSummary) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetSamples

`func (o *MetricsSummary) GetSamples() []SampledValue`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricsSummary) GetSamplesOk() (*[]SampledValue, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricsSummary) SetSamples(v []SampledValue)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MetricsSummary) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetTimestamp

`func (o *MetricsSummary) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricsSummary) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricsSummary) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricsSummary) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


