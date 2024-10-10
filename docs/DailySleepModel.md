# DailySleepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Contributors** | [**SleepContributors**](SleepContributors.md) | Contributors for the daily sleep score. | 
**Day** | **string** | Day that the daily sleep belongs to. | 
**Score** | **NullableInt32** |  | 
**Timestamp** | **string** | Timestamp of the daily sleep. | 

## Methods

### NewDailySleepModel

`func NewDailySleepModel(id string, contributors SleepContributors, day string, score NullableInt32, timestamp string, ) *DailySleepModel`

NewDailySleepModel instantiates a new DailySleepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailySleepModelWithDefaults

`func NewDailySleepModelWithDefaults() *DailySleepModel`

NewDailySleepModelWithDefaults instantiates a new DailySleepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailySleepModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailySleepModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailySleepModel) SetId(v string)`

SetId sets Id field to given value.


### GetContributors

`func (o *DailySleepModel) GetContributors() SleepContributors`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DailySleepModel) GetContributorsOk() (*SleepContributors, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DailySleepModel) SetContributors(v SleepContributors)`

SetContributors sets Contributors field to given value.


### GetDay

`func (o *DailySleepModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailySleepModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailySleepModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetScore

`func (o *DailySleepModel) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DailySleepModel) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DailySleepModel) SetScore(v int32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *DailySleepModel) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *DailySleepModel) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetTimestamp

`func (o *DailySleepModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DailySleepModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DailySleepModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


