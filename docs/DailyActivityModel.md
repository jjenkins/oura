# DailyActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Class5Min** | **NullableString** |  | 
**Score** | **NullableInt32** |  | 
**ActiveCalories** | **int32** | Active calories expended (in kilocalories) | 
**AverageMetMinutes** | **float32** | Average metabolic equivalent (MET) in minutes | 
**Contributors** | [**ActivityContributors**](ActivityContributors.md) |  | 
**EquivalentWalkingDistance** | **int32** | Equivalent walking distance (in meters) of energy expenditure | 
**HighActivityMetMinutes** | **int32** | High activity metabolic equivalent (MET) in minutes | 
**HighActivityTime** | **int32** | High activity metabolic equivalent (MET) in seconds | 
**InactivityAlerts** | **int32** | Number of inactivity alerts received | 
**LowActivityMetMinutes** | **int32** | Low activity metabolic equivalent (MET) in minutes | 
**LowActivityTime** | **int32** | Low activity metabolic equivalent (MET) in seconds | 
**MediumActivityMetMinutes** | **int32** | Medium activity metabolic equivalent (MET) in minutes | 
**MediumActivityTime** | **int32** | Medium activity metabolic equivalent (MET) in seconds | 
**Met** | [**SampleModel**](SampleModel.md) |  | 
**MetersToTarget** | **int32** | Remaining meters to target (from &#x60;&#x60;&#x60;target_meters&#x60;&#x60;&#x60; | 
**NonWearTime** | **int32** | The time (in seconds) in which the ring was not worn | 
**RestingTime** | **int32** | Resting time (in seconds) | 
**SedentaryMetMinutes** | **int32** | Sedentary metabolic equivalent (MET) in minutes | 
**SedentaryTime** | **int32** | Sedentary metabolic equivalent (MET) in seconds | 
**Steps** | **int32** | Total number of steps taken | 
**TargetCalories** | **int32** | Daily activity target (in kilocalories) | 
**TargetMeters** | **int32** | Daily activity target (in meters) | 
**TotalCalories** | **int32** | Total calories expended (in kilocalories) | 
**Day** | **string** | The &#x60;&#x60;&#x60;YYYY-MM-DD&#x60;&#x60;&#x60; formatted local date indicating when the daily activity occurred | 
**Timestamp** | **string** | ISO 8601 formatted local timestamp indicating the start datetime of when the daily activity occurred | 

## Methods

### NewDailyActivityModel

`func NewDailyActivityModel(id string, class5Min NullableString, score NullableInt32, activeCalories int32, averageMetMinutes float32, contributors ActivityContributors, equivalentWalkingDistance int32, highActivityMetMinutes int32, highActivityTime int32, inactivityAlerts int32, lowActivityMetMinutes int32, lowActivityTime int32, mediumActivityMetMinutes int32, mediumActivityTime int32, met SampleModel, metersToTarget int32, nonWearTime int32, restingTime int32, sedentaryMetMinutes int32, sedentaryTime int32, steps int32, targetCalories int32, targetMeters int32, totalCalories int32, day string, timestamp string, ) *DailyActivityModel`

NewDailyActivityModel instantiates a new DailyActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyActivityModelWithDefaults

`func NewDailyActivityModelWithDefaults() *DailyActivityModel`

NewDailyActivityModelWithDefaults instantiates a new DailyActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailyActivityModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailyActivityModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailyActivityModel) SetId(v string)`

SetId sets Id field to given value.


### GetClass5Min

`func (o *DailyActivityModel) GetClass5Min() string`

GetClass5Min returns the Class5Min field if non-nil, zero value otherwise.

### GetClass5MinOk

`func (o *DailyActivityModel) GetClass5MinOk() (*string, bool)`

GetClass5MinOk returns a tuple with the Class5Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass5Min

`func (o *DailyActivityModel) SetClass5Min(v string)`

SetClass5Min sets Class5Min field to given value.


### SetClass5MinNil

`func (o *DailyActivityModel) SetClass5MinNil(b bool)`

 SetClass5MinNil sets the value for Class5Min to be an explicit nil

### UnsetClass5Min
`func (o *DailyActivityModel) UnsetClass5Min()`

UnsetClass5Min ensures that no value is present for Class5Min, not even an explicit nil
### GetScore

`func (o *DailyActivityModel) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DailyActivityModel) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DailyActivityModel) SetScore(v int32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *DailyActivityModel) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *DailyActivityModel) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetActiveCalories

`func (o *DailyActivityModel) GetActiveCalories() int32`

GetActiveCalories returns the ActiveCalories field if non-nil, zero value otherwise.

### GetActiveCaloriesOk

`func (o *DailyActivityModel) GetActiveCaloriesOk() (*int32, bool)`

GetActiveCaloriesOk returns a tuple with the ActiveCalories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCalories

`func (o *DailyActivityModel) SetActiveCalories(v int32)`

SetActiveCalories sets ActiveCalories field to given value.


### GetAverageMetMinutes

`func (o *DailyActivityModel) GetAverageMetMinutes() float32`

GetAverageMetMinutes returns the AverageMetMinutes field if non-nil, zero value otherwise.

### GetAverageMetMinutesOk

`func (o *DailyActivityModel) GetAverageMetMinutesOk() (*float32, bool)`

GetAverageMetMinutesOk returns a tuple with the AverageMetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMetMinutes

`func (o *DailyActivityModel) SetAverageMetMinutes(v float32)`

SetAverageMetMinutes sets AverageMetMinutes field to given value.


### GetContributors

`func (o *DailyActivityModel) GetContributors() ActivityContributors`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DailyActivityModel) GetContributorsOk() (*ActivityContributors, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DailyActivityModel) SetContributors(v ActivityContributors)`

SetContributors sets Contributors field to given value.


### GetEquivalentWalkingDistance

`func (o *DailyActivityModel) GetEquivalentWalkingDistance() int32`

GetEquivalentWalkingDistance returns the EquivalentWalkingDistance field if non-nil, zero value otherwise.

### GetEquivalentWalkingDistanceOk

`func (o *DailyActivityModel) GetEquivalentWalkingDistanceOk() (*int32, bool)`

GetEquivalentWalkingDistanceOk returns a tuple with the EquivalentWalkingDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquivalentWalkingDistance

`func (o *DailyActivityModel) SetEquivalentWalkingDistance(v int32)`

SetEquivalentWalkingDistance sets EquivalentWalkingDistance field to given value.


### GetHighActivityMetMinutes

`func (o *DailyActivityModel) GetHighActivityMetMinutes() int32`

GetHighActivityMetMinutes returns the HighActivityMetMinutes field if non-nil, zero value otherwise.

### GetHighActivityMetMinutesOk

`func (o *DailyActivityModel) GetHighActivityMetMinutesOk() (*int32, bool)`

GetHighActivityMetMinutesOk returns a tuple with the HighActivityMetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighActivityMetMinutes

`func (o *DailyActivityModel) SetHighActivityMetMinutes(v int32)`

SetHighActivityMetMinutes sets HighActivityMetMinutes field to given value.


### GetHighActivityTime

`func (o *DailyActivityModel) GetHighActivityTime() int32`

GetHighActivityTime returns the HighActivityTime field if non-nil, zero value otherwise.

### GetHighActivityTimeOk

`func (o *DailyActivityModel) GetHighActivityTimeOk() (*int32, bool)`

GetHighActivityTimeOk returns a tuple with the HighActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighActivityTime

`func (o *DailyActivityModel) SetHighActivityTime(v int32)`

SetHighActivityTime sets HighActivityTime field to given value.


### GetInactivityAlerts

`func (o *DailyActivityModel) GetInactivityAlerts() int32`

GetInactivityAlerts returns the InactivityAlerts field if non-nil, zero value otherwise.

### GetInactivityAlertsOk

`func (o *DailyActivityModel) GetInactivityAlertsOk() (*int32, bool)`

GetInactivityAlertsOk returns a tuple with the InactivityAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityAlerts

`func (o *DailyActivityModel) SetInactivityAlerts(v int32)`

SetInactivityAlerts sets InactivityAlerts field to given value.


### GetLowActivityMetMinutes

`func (o *DailyActivityModel) GetLowActivityMetMinutes() int32`

GetLowActivityMetMinutes returns the LowActivityMetMinutes field if non-nil, zero value otherwise.

### GetLowActivityMetMinutesOk

`func (o *DailyActivityModel) GetLowActivityMetMinutesOk() (*int32, bool)`

GetLowActivityMetMinutesOk returns a tuple with the LowActivityMetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowActivityMetMinutes

`func (o *DailyActivityModel) SetLowActivityMetMinutes(v int32)`

SetLowActivityMetMinutes sets LowActivityMetMinutes field to given value.


### GetLowActivityTime

`func (o *DailyActivityModel) GetLowActivityTime() int32`

GetLowActivityTime returns the LowActivityTime field if non-nil, zero value otherwise.

### GetLowActivityTimeOk

`func (o *DailyActivityModel) GetLowActivityTimeOk() (*int32, bool)`

GetLowActivityTimeOk returns a tuple with the LowActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowActivityTime

`func (o *DailyActivityModel) SetLowActivityTime(v int32)`

SetLowActivityTime sets LowActivityTime field to given value.


### GetMediumActivityMetMinutes

`func (o *DailyActivityModel) GetMediumActivityMetMinutes() int32`

GetMediumActivityMetMinutes returns the MediumActivityMetMinutes field if non-nil, zero value otherwise.

### GetMediumActivityMetMinutesOk

`func (o *DailyActivityModel) GetMediumActivityMetMinutesOk() (*int32, bool)`

GetMediumActivityMetMinutesOk returns a tuple with the MediumActivityMetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumActivityMetMinutes

`func (o *DailyActivityModel) SetMediumActivityMetMinutes(v int32)`

SetMediumActivityMetMinutes sets MediumActivityMetMinutes field to given value.


### GetMediumActivityTime

`func (o *DailyActivityModel) GetMediumActivityTime() int32`

GetMediumActivityTime returns the MediumActivityTime field if non-nil, zero value otherwise.

### GetMediumActivityTimeOk

`func (o *DailyActivityModel) GetMediumActivityTimeOk() (*int32, bool)`

GetMediumActivityTimeOk returns a tuple with the MediumActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumActivityTime

`func (o *DailyActivityModel) SetMediumActivityTime(v int32)`

SetMediumActivityTime sets MediumActivityTime field to given value.


### GetMet

`func (o *DailyActivityModel) GetMet() SampleModel`

GetMet returns the Met field if non-nil, zero value otherwise.

### GetMetOk

`func (o *DailyActivityModel) GetMetOk() (*SampleModel, bool)`

GetMetOk returns a tuple with the Met field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMet

`func (o *DailyActivityModel) SetMet(v SampleModel)`

SetMet sets Met field to given value.


### GetMetersToTarget

`func (o *DailyActivityModel) GetMetersToTarget() int32`

GetMetersToTarget returns the MetersToTarget field if non-nil, zero value otherwise.

### GetMetersToTargetOk

`func (o *DailyActivityModel) GetMetersToTargetOk() (*int32, bool)`

GetMetersToTargetOk returns a tuple with the MetersToTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetersToTarget

`func (o *DailyActivityModel) SetMetersToTarget(v int32)`

SetMetersToTarget sets MetersToTarget field to given value.


### GetNonWearTime

`func (o *DailyActivityModel) GetNonWearTime() int32`

GetNonWearTime returns the NonWearTime field if non-nil, zero value otherwise.

### GetNonWearTimeOk

`func (o *DailyActivityModel) GetNonWearTimeOk() (*int32, bool)`

GetNonWearTimeOk returns a tuple with the NonWearTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonWearTime

`func (o *DailyActivityModel) SetNonWearTime(v int32)`

SetNonWearTime sets NonWearTime field to given value.


### GetRestingTime

`func (o *DailyActivityModel) GetRestingTime() int32`

GetRestingTime returns the RestingTime field if non-nil, zero value otherwise.

### GetRestingTimeOk

`func (o *DailyActivityModel) GetRestingTimeOk() (*int32, bool)`

GetRestingTimeOk returns a tuple with the RestingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestingTime

`func (o *DailyActivityModel) SetRestingTime(v int32)`

SetRestingTime sets RestingTime field to given value.


### GetSedentaryMetMinutes

`func (o *DailyActivityModel) GetSedentaryMetMinutes() int32`

GetSedentaryMetMinutes returns the SedentaryMetMinutes field if non-nil, zero value otherwise.

### GetSedentaryMetMinutesOk

`func (o *DailyActivityModel) GetSedentaryMetMinutesOk() (*int32, bool)`

GetSedentaryMetMinutesOk returns a tuple with the SedentaryMetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedentaryMetMinutes

`func (o *DailyActivityModel) SetSedentaryMetMinutes(v int32)`

SetSedentaryMetMinutes sets SedentaryMetMinutes field to given value.


### GetSedentaryTime

`func (o *DailyActivityModel) GetSedentaryTime() int32`

GetSedentaryTime returns the SedentaryTime field if non-nil, zero value otherwise.

### GetSedentaryTimeOk

`func (o *DailyActivityModel) GetSedentaryTimeOk() (*int32, bool)`

GetSedentaryTimeOk returns a tuple with the SedentaryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedentaryTime

`func (o *DailyActivityModel) SetSedentaryTime(v int32)`

SetSedentaryTime sets SedentaryTime field to given value.


### GetSteps

`func (o *DailyActivityModel) GetSteps() int32`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DailyActivityModel) GetStepsOk() (*int32, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DailyActivityModel) SetSteps(v int32)`

SetSteps sets Steps field to given value.


### GetTargetCalories

`func (o *DailyActivityModel) GetTargetCalories() int32`

GetTargetCalories returns the TargetCalories field if non-nil, zero value otherwise.

### GetTargetCaloriesOk

`func (o *DailyActivityModel) GetTargetCaloriesOk() (*int32, bool)`

GetTargetCaloriesOk returns a tuple with the TargetCalories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCalories

`func (o *DailyActivityModel) SetTargetCalories(v int32)`

SetTargetCalories sets TargetCalories field to given value.


### GetTargetMeters

`func (o *DailyActivityModel) GetTargetMeters() int32`

GetTargetMeters returns the TargetMeters field if non-nil, zero value otherwise.

### GetTargetMetersOk

`func (o *DailyActivityModel) GetTargetMetersOk() (*int32, bool)`

GetTargetMetersOk returns a tuple with the TargetMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMeters

`func (o *DailyActivityModel) SetTargetMeters(v int32)`

SetTargetMeters sets TargetMeters field to given value.


### GetTotalCalories

`func (o *DailyActivityModel) GetTotalCalories() int32`

GetTotalCalories returns the TotalCalories field if non-nil, zero value otherwise.

### GetTotalCaloriesOk

`func (o *DailyActivityModel) GetTotalCaloriesOk() (*int32, bool)`

GetTotalCaloriesOk returns a tuple with the TotalCalories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCalories

`func (o *DailyActivityModel) SetTotalCalories(v int32)`

SetTotalCalories sets TotalCalories field to given value.


### GetDay

`func (o *DailyActivityModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailyActivityModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailyActivityModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetTimestamp

`func (o *DailyActivityModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DailyActivityModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DailyActivityModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


