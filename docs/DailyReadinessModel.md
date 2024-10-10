# DailyReadinessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Contributors** | [**ReadinessContributors**](ReadinessContributors.md) | Contributors of the daily readiness score. | 
**Day** | **string** | Day that the daily readiness belongs to. | 
**Score** | **NullableInt32** |  | 
**TemperatureDeviation** | **NullableFloat32** |  | 
**TemperatureTrendDeviation** | **NullableFloat32** |  | 
**Timestamp** | **string** | Timestamp of the daily readiness. | 

## Methods

### NewDailyReadinessModel

`func NewDailyReadinessModel(id string, contributors ReadinessContributors, day string, score NullableInt32, temperatureDeviation NullableFloat32, temperatureTrendDeviation NullableFloat32, timestamp string, ) *DailyReadinessModel`

NewDailyReadinessModel instantiates a new DailyReadinessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyReadinessModelWithDefaults

`func NewDailyReadinessModelWithDefaults() *DailyReadinessModel`

NewDailyReadinessModelWithDefaults instantiates a new DailyReadinessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailyReadinessModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailyReadinessModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailyReadinessModel) SetId(v string)`

SetId sets Id field to given value.


### GetContributors

`func (o *DailyReadinessModel) GetContributors() ReadinessContributors`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DailyReadinessModel) GetContributorsOk() (*ReadinessContributors, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DailyReadinessModel) SetContributors(v ReadinessContributors)`

SetContributors sets Contributors field to given value.


### GetDay

`func (o *DailyReadinessModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailyReadinessModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailyReadinessModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetScore

`func (o *DailyReadinessModel) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DailyReadinessModel) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DailyReadinessModel) SetScore(v int32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *DailyReadinessModel) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *DailyReadinessModel) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetTemperatureDeviation

`func (o *DailyReadinessModel) GetTemperatureDeviation() float32`

GetTemperatureDeviation returns the TemperatureDeviation field if non-nil, zero value otherwise.

### GetTemperatureDeviationOk

`func (o *DailyReadinessModel) GetTemperatureDeviationOk() (*float32, bool)`

GetTemperatureDeviationOk returns a tuple with the TemperatureDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureDeviation

`func (o *DailyReadinessModel) SetTemperatureDeviation(v float32)`

SetTemperatureDeviation sets TemperatureDeviation field to given value.


### SetTemperatureDeviationNil

`func (o *DailyReadinessModel) SetTemperatureDeviationNil(b bool)`

 SetTemperatureDeviationNil sets the value for TemperatureDeviation to be an explicit nil

### UnsetTemperatureDeviation
`func (o *DailyReadinessModel) UnsetTemperatureDeviation()`

UnsetTemperatureDeviation ensures that no value is present for TemperatureDeviation, not even an explicit nil
### GetTemperatureTrendDeviation

`func (o *DailyReadinessModel) GetTemperatureTrendDeviation() float32`

GetTemperatureTrendDeviation returns the TemperatureTrendDeviation field if non-nil, zero value otherwise.

### GetTemperatureTrendDeviationOk

`func (o *DailyReadinessModel) GetTemperatureTrendDeviationOk() (*float32, bool)`

GetTemperatureTrendDeviationOk returns a tuple with the TemperatureTrendDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureTrendDeviation

`func (o *DailyReadinessModel) SetTemperatureTrendDeviation(v float32)`

SetTemperatureTrendDeviation sets TemperatureTrendDeviation field to given value.


### SetTemperatureTrendDeviationNil

`func (o *DailyReadinessModel) SetTemperatureTrendDeviationNil(b bool)`

 SetTemperatureTrendDeviationNil sets the value for TemperatureTrendDeviation to be an explicit nil

### UnsetTemperatureTrendDeviation
`func (o *DailyReadinessModel) UnsetTemperatureTrendDeviation()`

UnsetTemperatureTrendDeviation ensures that no value is present for TemperatureTrendDeviation, not even an explicit nil
### GetTimestamp

`func (o *DailyReadinessModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DailyReadinessModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DailyReadinessModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


