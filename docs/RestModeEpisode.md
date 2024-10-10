# RestModeEpisode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** | Tags selected for the episode. | 
**Timestamp** | **string** | Timestamp indicating when the episode occurred. | 

## Methods

### NewRestModeEpisode

`func NewRestModeEpisode(tags []string, timestamp string, ) *RestModeEpisode`

NewRestModeEpisode instantiates a new RestModeEpisode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestModeEpisodeWithDefaults

`func NewRestModeEpisodeWithDefaults() *RestModeEpisode`

NewRestModeEpisodeWithDefaults instantiates a new RestModeEpisode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RestModeEpisode) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestModeEpisode) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestModeEpisode) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTimestamp

`func (o *RestModeEpisode) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestModeEpisode) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestModeEpisode) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


