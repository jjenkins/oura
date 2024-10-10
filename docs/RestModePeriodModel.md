# RestModePeriodModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EndDay** | Pointer to **NullableString** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Episodes** | [**[]RestModeEpisode**](RestModeEpisode.md) | Collection of episodes during rest mode, consisting of tags. | 
**StartDay** | **string** | Start date of rest mode. | 
**StartTime** | **string** |  | 

## Methods

### NewRestModePeriodModel

`func NewRestModePeriodModel(id string, episodes []RestModeEpisode, startDay string, startTime string, ) *RestModePeriodModel`

NewRestModePeriodModel instantiates a new RestModePeriodModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestModePeriodModelWithDefaults

`func NewRestModePeriodModelWithDefaults() *RestModePeriodModel`

NewRestModePeriodModelWithDefaults instantiates a new RestModePeriodModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestModePeriodModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestModePeriodModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestModePeriodModel) SetId(v string)`

SetId sets Id field to given value.


### GetEndDay

`func (o *RestModePeriodModel) GetEndDay() string`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *RestModePeriodModel) GetEndDayOk() (*string, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *RestModePeriodModel) SetEndDay(v string)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *RestModePeriodModel) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### SetEndDayNil

`func (o *RestModePeriodModel) SetEndDayNil(b bool)`

 SetEndDayNil sets the value for EndDay to be an explicit nil

### UnsetEndDay
`func (o *RestModePeriodModel) UnsetEndDay()`

UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
### GetEndTime

`func (o *RestModePeriodModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RestModePeriodModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RestModePeriodModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RestModePeriodModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEpisodes

`func (o *RestModePeriodModel) GetEpisodes() []RestModeEpisode`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *RestModePeriodModel) GetEpisodesOk() (*[]RestModeEpisode, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *RestModePeriodModel) SetEpisodes(v []RestModeEpisode)`

SetEpisodes sets Episodes field to given value.


### GetStartDay

`func (o *RestModePeriodModel) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *RestModePeriodModel) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *RestModePeriodModel) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.


### GetStartTime

`func (o *RestModePeriodModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RestModePeriodModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RestModePeriodModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


