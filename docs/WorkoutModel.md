# WorkoutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Activity** | **string** | Type of the workout activity. | 
**Calories** | **NullableFloat32** |  | 
**Day** | **string** | Day when the workout occurred. | 
**Distance** | **NullableFloat32** |  | 
**EndDatetime** | **string** | Timestamp indicating when the workout ended. | 
**Intensity** | [**WorkoutIntensity**](WorkoutIntensity.md) |  | 
**Label** | **NullableString** |  | 
**Source** | [**WorkoutSource**](WorkoutSource.md) | Possible workout sources. | 
**StartDatetime** | **string** | Timestamp indicating when the workout started. | 

## Methods

### NewWorkoutModel

`func NewWorkoutModel(id string, activity string, calories NullableFloat32, day string, distance NullableFloat32, endDatetime string, intensity WorkoutIntensity, label NullableString, source WorkoutSource, startDatetime string, ) *WorkoutModel`

NewWorkoutModel instantiates a new WorkoutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkoutModelWithDefaults

`func NewWorkoutModelWithDefaults() *WorkoutModel`

NewWorkoutModelWithDefaults instantiates a new WorkoutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkoutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkoutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkoutModel) SetId(v string)`

SetId sets Id field to given value.


### GetActivity

`func (o *WorkoutModel) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *WorkoutModel) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *WorkoutModel) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetCalories

`func (o *WorkoutModel) GetCalories() float32`

GetCalories returns the Calories field if non-nil, zero value otherwise.

### GetCaloriesOk

`func (o *WorkoutModel) GetCaloriesOk() (*float32, bool)`

GetCaloriesOk returns a tuple with the Calories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalories

`func (o *WorkoutModel) SetCalories(v float32)`

SetCalories sets Calories field to given value.


### SetCaloriesNil

`func (o *WorkoutModel) SetCaloriesNil(b bool)`

 SetCaloriesNil sets the value for Calories to be an explicit nil

### UnsetCalories
`func (o *WorkoutModel) UnsetCalories()`

UnsetCalories ensures that no value is present for Calories, not even an explicit nil
### GetDay

`func (o *WorkoutModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *WorkoutModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *WorkoutModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetDistance

`func (o *WorkoutModel) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *WorkoutModel) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *WorkoutModel) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### SetDistanceNil

`func (o *WorkoutModel) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *WorkoutModel) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetEndDatetime

`func (o *WorkoutModel) GetEndDatetime() string`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *WorkoutModel) GetEndDatetimeOk() (*string, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *WorkoutModel) SetEndDatetime(v string)`

SetEndDatetime sets EndDatetime field to given value.


### GetIntensity

`func (o *WorkoutModel) GetIntensity() WorkoutIntensity`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *WorkoutModel) GetIntensityOk() (*WorkoutIntensity, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *WorkoutModel) SetIntensity(v WorkoutIntensity)`

SetIntensity sets Intensity field to given value.


### GetLabel

`func (o *WorkoutModel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkoutModel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkoutModel) SetLabel(v string)`

SetLabel sets Label field to given value.


### SetLabelNil

`func (o *WorkoutModel) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *WorkoutModel) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetSource

`func (o *WorkoutModel) GetSource() WorkoutSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WorkoutModel) GetSourceOk() (*WorkoutSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WorkoutModel) SetSource(v WorkoutSource)`

SetSource sets Source field to given value.


### GetStartDatetime

`func (o *WorkoutModel) GetStartDatetime() string`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *WorkoutModel) GetStartDatetimeOk() (*string, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *WorkoutModel) SetStartDatetime(v string)`

SetStartDatetime sets StartDatetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


