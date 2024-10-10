# SleepTimeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** | Corresponding day for the sleep time. | 
**OptimalBedtime** | Pointer to [**NullableSleepTimeWindow**](SleepTimeWindow.md) |  | [optional] 
**Recommendation** | Pointer to [**NullableSleepTimeRecommendation**](SleepTimeRecommendation.md) |  | [optional] 
**Status** | Pointer to [**NullableSleepTimeStatus**](SleepTimeStatus.md) |  | [optional] 

## Methods

### NewSleepTimeModel

`func NewSleepTimeModel(id string, day string, ) *SleepTimeModel`

NewSleepTimeModel instantiates a new SleepTimeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleepTimeModelWithDefaults

`func NewSleepTimeModelWithDefaults() *SleepTimeModel`

NewSleepTimeModelWithDefaults instantiates a new SleepTimeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SleepTimeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SleepTimeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SleepTimeModel) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *SleepTimeModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *SleepTimeModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *SleepTimeModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetOptimalBedtime

`func (o *SleepTimeModel) GetOptimalBedtime() SleepTimeWindow`

GetOptimalBedtime returns the OptimalBedtime field if non-nil, zero value otherwise.

### GetOptimalBedtimeOk

`func (o *SleepTimeModel) GetOptimalBedtimeOk() (*SleepTimeWindow, bool)`

GetOptimalBedtimeOk returns a tuple with the OptimalBedtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimalBedtime

`func (o *SleepTimeModel) SetOptimalBedtime(v SleepTimeWindow)`

SetOptimalBedtime sets OptimalBedtime field to given value.

### HasOptimalBedtime

`func (o *SleepTimeModel) HasOptimalBedtime() bool`

HasOptimalBedtime returns a boolean if a field has been set.

### SetOptimalBedtimeNil

`func (o *SleepTimeModel) SetOptimalBedtimeNil(b bool)`

 SetOptimalBedtimeNil sets the value for OptimalBedtime to be an explicit nil

### UnsetOptimalBedtime
`func (o *SleepTimeModel) UnsetOptimalBedtime()`

UnsetOptimalBedtime ensures that no value is present for OptimalBedtime, not even an explicit nil
### GetRecommendation

`func (o *SleepTimeModel) GetRecommendation() SleepTimeRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *SleepTimeModel) GetRecommendationOk() (*SleepTimeRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *SleepTimeModel) SetRecommendation(v SleepTimeRecommendation)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *SleepTimeModel) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *SleepTimeModel) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *SleepTimeModel) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetStatus

`func (o *SleepTimeModel) GetStatus() SleepTimeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SleepTimeModel) GetStatusOk() (*SleepTimeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SleepTimeModel) SetStatus(v SleepTimeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SleepTimeModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SleepTimeModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SleepTimeModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


