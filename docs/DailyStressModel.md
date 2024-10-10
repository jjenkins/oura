# DailyStressModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** | Day that the daily stress belongs to. | 
**StressHigh** | **NullableInt32** |  | 
**RecoveryHigh** | **NullableInt32** |  | 
**DaySummary** | Pointer to [**NullableDailyStressSummary**](DailyStressSummary.md) |  | [optional] 

## Methods

### NewDailyStressModel

`func NewDailyStressModel(id string, day string, stressHigh NullableInt32, recoveryHigh NullableInt32, ) *DailyStressModel`

NewDailyStressModel instantiates a new DailyStressModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyStressModelWithDefaults

`func NewDailyStressModelWithDefaults() *DailyStressModel`

NewDailyStressModelWithDefaults instantiates a new DailyStressModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailyStressModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailyStressModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailyStressModel) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *DailyStressModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailyStressModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailyStressModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetStressHigh

`func (o *DailyStressModel) GetStressHigh() int32`

GetStressHigh returns the StressHigh field if non-nil, zero value otherwise.

### GetStressHighOk

`func (o *DailyStressModel) GetStressHighOk() (*int32, bool)`

GetStressHighOk returns a tuple with the StressHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStressHigh

`func (o *DailyStressModel) SetStressHigh(v int32)`

SetStressHigh sets StressHigh field to given value.


### SetStressHighNil

`func (o *DailyStressModel) SetStressHighNil(b bool)`

 SetStressHighNil sets the value for StressHigh to be an explicit nil

### UnsetStressHigh
`func (o *DailyStressModel) UnsetStressHigh()`

UnsetStressHigh ensures that no value is present for StressHigh, not even an explicit nil
### GetRecoveryHigh

`func (o *DailyStressModel) GetRecoveryHigh() int32`

GetRecoveryHigh returns the RecoveryHigh field if non-nil, zero value otherwise.

### GetRecoveryHighOk

`func (o *DailyStressModel) GetRecoveryHighOk() (*int32, bool)`

GetRecoveryHighOk returns a tuple with the RecoveryHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryHigh

`func (o *DailyStressModel) SetRecoveryHigh(v int32)`

SetRecoveryHigh sets RecoveryHigh field to given value.


### SetRecoveryHighNil

`func (o *DailyStressModel) SetRecoveryHighNil(b bool)`

 SetRecoveryHighNil sets the value for RecoveryHigh to be an explicit nil

### UnsetRecoveryHigh
`func (o *DailyStressModel) UnsetRecoveryHigh()`

UnsetRecoveryHigh ensures that no value is present for RecoveryHigh, not even an explicit nil
### GetDaySummary

`func (o *DailyStressModel) GetDaySummary() DailyStressSummary`

GetDaySummary returns the DaySummary field if non-nil, zero value otherwise.

### GetDaySummaryOk

`func (o *DailyStressModel) GetDaySummaryOk() (*DailyStressSummary, bool)`

GetDaySummaryOk returns a tuple with the DaySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySummary

`func (o *DailyStressModel) SetDaySummary(v DailyStressSummary)`

SetDaySummary sets DaySummary field to given value.

### HasDaySummary

`func (o *DailyStressModel) HasDaySummary() bool`

HasDaySummary returns a boolean if a field has been set.

### SetDaySummaryNil

`func (o *DailyStressModel) SetDaySummaryNil(b bool)`

 SetDaySummaryNil sets the value for DaySummary to be an explicit nil

### UnsetDaySummary
`func (o *DailyStressModel) UnsetDaySummary()`

UnsetDaySummary ensures that no value is present for DaySummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


