# SleepTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayTz** | **int32** | Timezone offset in second from GMT of the day | 
**EndOffset** | **int32** | End offset from midnight in second | 
**StartOffset** | **int32** | Start offset from midnight in second | 

## Methods

### NewSleepTimeWindow

`func NewSleepTimeWindow(dayTz int32, endOffset int32, startOffset int32, ) *SleepTimeWindow`

NewSleepTimeWindow instantiates a new SleepTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleepTimeWindowWithDefaults

`func NewSleepTimeWindowWithDefaults() *SleepTimeWindow`

NewSleepTimeWindowWithDefaults instantiates a new SleepTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayTz

`func (o *SleepTimeWindow) GetDayTz() int32`

GetDayTz returns the DayTz field if non-nil, zero value otherwise.

### GetDayTzOk

`func (o *SleepTimeWindow) GetDayTzOk() (*int32, bool)`

GetDayTzOk returns a tuple with the DayTz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTz

`func (o *SleepTimeWindow) SetDayTz(v int32)`

SetDayTz sets DayTz field to given value.


### GetEndOffset

`func (o *SleepTimeWindow) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *SleepTimeWindow) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *SleepTimeWindow) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.


### GetStartOffset

`func (o *SleepTimeWindow) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *SleepTimeWindow) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *SleepTimeWindow) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


