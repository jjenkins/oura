# SessionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** | The date when the session occurred. | 
**StartDatetime** | **string** | Timestamp indicating when the Moment ended. | 
**EndDatetime** | **string** | Timestamp indicating when the Moment ended. | 
**Type** | [**MomentType**](MomentType.md) |  | 
**HeartRate** | Pointer to [**NullableSampleModel**](SampleModel.md) |  | [optional] 
**HeartRateVariability** | Pointer to [**NullableSampleModel**](SampleModel.md) |  | [optional] 
**Mood** | Pointer to [**NullableMomentMood**](MomentMood.md) |  | [optional] 
**MotionCount** | Pointer to [**NullableSampleModel**](SampleModel.md) |  | [optional] 

## Methods

### NewSessionModel

`func NewSessionModel(id string, day string, startDatetime string, endDatetime string, type_ MomentType, ) *SessionModel`

NewSessionModel instantiates a new SessionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionModelWithDefaults

`func NewSessionModelWithDefaults() *SessionModel`

NewSessionModelWithDefaults instantiates a new SessionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionModel) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *SessionModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *SessionModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *SessionModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetStartDatetime

`func (o *SessionModel) GetStartDatetime() string`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *SessionModel) GetStartDatetimeOk() (*string, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *SessionModel) SetStartDatetime(v string)`

SetStartDatetime sets StartDatetime field to given value.


### GetEndDatetime

`func (o *SessionModel) GetEndDatetime() string`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *SessionModel) GetEndDatetimeOk() (*string, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *SessionModel) SetEndDatetime(v string)`

SetEndDatetime sets EndDatetime field to given value.


### GetType

`func (o *SessionModel) GetType() MomentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionModel) GetTypeOk() (*MomentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionModel) SetType(v MomentType)`

SetType sets Type field to given value.


### GetHeartRate

`func (o *SessionModel) GetHeartRate() SampleModel`

GetHeartRate returns the HeartRate field if non-nil, zero value otherwise.

### GetHeartRateOk

`func (o *SessionModel) GetHeartRateOk() (*SampleModel, bool)`

GetHeartRateOk returns a tuple with the HeartRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartRate

`func (o *SessionModel) SetHeartRate(v SampleModel)`

SetHeartRate sets HeartRate field to given value.

### HasHeartRate

`func (o *SessionModel) HasHeartRate() bool`

HasHeartRate returns a boolean if a field has been set.

### SetHeartRateNil

`func (o *SessionModel) SetHeartRateNil(b bool)`

 SetHeartRateNil sets the value for HeartRate to be an explicit nil

### UnsetHeartRate
`func (o *SessionModel) UnsetHeartRate()`

UnsetHeartRate ensures that no value is present for HeartRate, not even an explicit nil
### GetHeartRateVariability

`func (o *SessionModel) GetHeartRateVariability() SampleModel`

GetHeartRateVariability returns the HeartRateVariability field if non-nil, zero value otherwise.

### GetHeartRateVariabilityOk

`func (o *SessionModel) GetHeartRateVariabilityOk() (*SampleModel, bool)`

GetHeartRateVariabilityOk returns a tuple with the HeartRateVariability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartRateVariability

`func (o *SessionModel) SetHeartRateVariability(v SampleModel)`

SetHeartRateVariability sets HeartRateVariability field to given value.

### HasHeartRateVariability

`func (o *SessionModel) HasHeartRateVariability() bool`

HasHeartRateVariability returns a boolean if a field has been set.

### SetHeartRateVariabilityNil

`func (o *SessionModel) SetHeartRateVariabilityNil(b bool)`

 SetHeartRateVariabilityNil sets the value for HeartRateVariability to be an explicit nil

### UnsetHeartRateVariability
`func (o *SessionModel) UnsetHeartRateVariability()`

UnsetHeartRateVariability ensures that no value is present for HeartRateVariability, not even an explicit nil
### GetMood

`func (o *SessionModel) GetMood() MomentMood`

GetMood returns the Mood field if non-nil, zero value otherwise.

### GetMoodOk

`func (o *SessionModel) GetMoodOk() (*MomentMood, bool)`

GetMoodOk returns a tuple with the Mood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMood

`func (o *SessionModel) SetMood(v MomentMood)`

SetMood sets Mood field to given value.

### HasMood

`func (o *SessionModel) HasMood() bool`

HasMood returns a boolean if a field has been set.

### SetMoodNil

`func (o *SessionModel) SetMoodNil(b bool)`

 SetMoodNil sets the value for Mood to be an explicit nil

### UnsetMood
`func (o *SessionModel) UnsetMood()`

UnsetMood ensures that no value is present for Mood, not even an explicit nil
### GetMotionCount

`func (o *SessionModel) GetMotionCount() SampleModel`

GetMotionCount returns the MotionCount field if non-nil, zero value otherwise.

### GetMotionCountOk

`func (o *SessionModel) GetMotionCountOk() (*SampleModel, bool)`

GetMotionCountOk returns a tuple with the MotionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionCount

`func (o *SessionModel) SetMotionCount(v SampleModel)`

SetMotionCount sets MotionCount field to given value.

### HasMotionCount

`func (o *SessionModel) HasMotionCount() bool`

HasMotionCount returns a boolean if a field has been set.

### SetMotionCountNil

`func (o *SessionModel) SetMotionCountNil(b bool)`

 SetMotionCountNil sets the value for MotionCount to be an explicit nil

### UnsetMotionCount
`func (o *SessionModel) UnsetMotionCount()`

UnsetMotionCount ensures that no value is present for MotionCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


