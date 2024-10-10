# SampleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **float32** | Interval in seconds between the sampled items. | 
**Items** | **[]float32** | Recorded sample items. | 
**Timestamp** | **string** | Timestamp when the sample recording started. | 

## Methods

### NewSampleModel

`func NewSampleModel(interval float32, items []*float32, timestamp string, ) *SampleModel`

NewSampleModel instantiates a new SampleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleModelWithDefaults

`func NewSampleModelWithDefaults() *SampleModel`

NewSampleModelWithDefaults instantiates a new SampleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *SampleModel) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SampleModel) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SampleModel) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetItems

`func (o *SampleModel) GetItems() []*float32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SampleModel) GetItemsOk() (*[]*float32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SampleModel) SetItems(v []*float32)`

SetItems sets Items field to given value.


### GetTimestamp

`func (o *SampleModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SampleModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SampleModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


