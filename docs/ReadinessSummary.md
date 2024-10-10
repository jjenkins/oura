# ReadinessSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contributors** | [**ReadinessContributors**](ReadinessContributors.md) |  | 
**Score** | Pointer to **NullableInt32** |  | [optional] 
**TemperatureDeviation** | Pointer to **NullableFloat32** |  | [optional] 
**TemperatureTrendDeviation** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewReadinessSummary

`func NewReadinessSummary(contributors ReadinessContributors, ) *ReadinessSummary`

NewReadinessSummary instantiates a new ReadinessSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadinessSummaryWithDefaults

`func NewReadinessSummaryWithDefaults() *ReadinessSummary`

NewReadinessSummaryWithDefaults instantiates a new ReadinessSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributors

`func (o *ReadinessSummary) GetContributors() ReadinessContributors`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ReadinessSummary) GetContributorsOk() (*ReadinessContributors, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ReadinessSummary) SetContributors(v ReadinessContributors)`

SetContributors sets Contributors field to given value.


### GetScore

`func (o *ReadinessSummary) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ReadinessSummary) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ReadinessSummary) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ReadinessSummary) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ReadinessSummary) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ReadinessSummary) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetTemperatureDeviation

`func (o *ReadinessSummary) GetTemperatureDeviation() float32`

GetTemperatureDeviation returns the TemperatureDeviation field if non-nil, zero value otherwise.

### GetTemperatureDeviationOk

`func (o *ReadinessSummary) GetTemperatureDeviationOk() (*float32, bool)`

GetTemperatureDeviationOk returns a tuple with the TemperatureDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureDeviation

`func (o *ReadinessSummary) SetTemperatureDeviation(v float32)`

SetTemperatureDeviation sets TemperatureDeviation field to given value.

### HasTemperatureDeviation

`func (o *ReadinessSummary) HasTemperatureDeviation() bool`

HasTemperatureDeviation returns a boolean if a field has been set.

### SetTemperatureDeviationNil

`func (o *ReadinessSummary) SetTemperatureDeviationNil(b bool)`

 SetTemperatureDeviationNil sets the value for TemperatureDeviation to be an explicit nil

### UnsetTemperatureDeviation
`func (o *ReadinessSummary) UnsetTemperatureDeviation()`

UnsetTemperatureDeviation ensures that no value is present for TemperatureDeviation, not even an explicit nil
### GetTemperatureTrendDeviation

`func (o *ReadinessSummary) GetTemperatureTrendDeviation() float32`

GetTemperatureTrendDeviation returns the TemperatureTrendDeviation field if non-nil, zero value otherwise.

### GetTemperatureTrendDeviationOk

`func (o *ReadinessSummary) GetTemperatureTrendDeviationOk() (*float32, bool)`

GetTemperatureTrendDeviationOk returns a tuple with the TemperatureTrendDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureTrendDeviation

`func (o *ReadinessSummary) SetTemperatureTrendDeviation(v float32)`

SetTemperatureTrendDeviation sets TemperatureTrendDeviation field to given value.

### HasTemperatureTrendDeviation

`func (o *ReadinessSummary) HasTemperatureTrendDeviation() bool`

HasTemperatureTrendDeviation returns a boolean if a field has been set.

### SetTemperatureTrendDeviationNil

`func (o *ReadinessSummary) SetTemperatureTrendDeviationNil(b bool)`

 SetTemperatureTrendDeviationNil sets the value for TemperatureTrendDeviation to be an explicit nil

### UnsetTemperatureTrendDeviation
`func (o *ReadinessSummary) UnsetTemperatureTrendDeviation()`

UnsetTemperatureTrendDeviation ensures that no value is present for TemperatureTrendDeviation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


