# EnhancedTagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TagTypeCode** | Pointer to **NullableString** |  | [optional] 
**StartTime** | **string** | Timestamp of the tag (if no duration) or the start time of the tag (with duration). | 
**EndTime** | Pointer to **string** |  | [optional] 
**StartDay** | **string** | Day of the tag (if no duration) or the start day of the tag (with duration). | 
**EndDay** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**CustomName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnhancedTagModel

`func NewEnhancedTagModel(id string, startTime string, startDay string, ) *EnhancedTagModel`

NewEnhancedTagModel instantiates a new EnhancedTagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedTagModelWithDefaults

`func NewEnhancedTagModelWithDefaults() *EnhancedTagModel`

NewEnhancedTagModelWithDefaults instantiates a new EnhancedTagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnhancedTagModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedTagModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedTagModel) SetId(v string)`

SetId sets Id field to given value.


### GetTagTypeCode

`func (o *EnhancedTagModel) GetTagTypeCode() string`

GetTagTypeCode returns the TagTypeCode field if non-nil, zero value otherwise.

### GetTagTypeCodeOk

`func (o *EnhancedTagModel) GetTagTypeCodeOk() (*string, bool)`

GetTagTypeCodeOk returns a tuple with the TagTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypeCode

`func (o *EnhancedTagModel) SetTagTypeCode(v string)`

SetTagTypeCode sets TagTypeCode field to given value.

### HasTagTypeCode

`func (o *EnhancedTagModel) HasTagTypeCode() bool`

HasTagTypeCode returns a boolean if a field has been set.

### SetTagTypeCodeNil

`func (o *EnhancedTagModel) SetTagTypeCodeNil(b bool)`

 SetTagTypeCodeNil sets the value for TagTypeCode to be an explicit nil

### UnsetTagTypeCode
`func (o *EnhancedTagModel) UnsetTagTypeCode()`

UnsetTagTypeCode ensures that no value is present for TagTypeCode, not even an explicit nil
### GetStartTime

`func (o *EnhancedTagModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnhancedTagModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnhancedTagModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *EnhancedTagModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EnhancedTagModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EnhancedTagModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EnhancedTagModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDay

`func (o *EnhancedTagModel) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *EnhancedTagModel) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *EnhancedTagModel) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.


### GetEndDay

`func (o *EnhancedTagModel) GetEndDay() string`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *EnhancedTagModel) GetEndDayOk() (*string, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *EnhancedTagModel) SetEndDay(v string)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *EnhancedTagModel) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### SetEndDayNil

`func (o *EnhancedTagModel) SetEndDayNil(b bool)`

 SetEndDayNil sets the value for EndDay to be an explicit nil

### UnsetEndDay
`func (o *EnhancedTagModel) UnsetEndDay()`

UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
### GetComment

`func (o *EnhancedTagModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EnhancedTagModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EnhancedTagModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *EnhancedTagModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *EnhancedTagModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *EnhancedTagModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCustomName

`func (o *EnhancedTagModel) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *EnhancedTagModel) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *EnhancedTagModel) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *EnhancedTagModel) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *EnhancedTagModel) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *EnhancedTagModel) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


