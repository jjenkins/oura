# HeartRateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bpm** | **int32** |  | 
**Source** | [**HeartRateSource**](HeartRateSource.md) |  | 
**Timestamp** | **string** |  | 

## Methods

### NewHeartRateModel

`func NewHeartRateModel(bpm int32, source HeartRateSource, timestamp string, ) *HeartRateModel`

NewHeartRateModel instantiates a new HeartRateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartRateModelWithDefaults

`func NewHeartRateModelWithDefaults() *HeartRateModel`

NewHeartRateModelWithDefaults instantiates a new HeartRateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBpm

`func (o *HeartRateModel) GetBpm() int32`

GetBpm returns the Bpm field if non-nil, zero value otherwise.

### GetBpmOk

`func (o *HeartRateModel) GetBpmOk() (*int32, bool)`

GetBpmOk returns a tuple with the Bpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpm

`func (o *HeartRateModel) SetBpm(v int32)`

SetBpm sets Bpm field to given value.


### GetSource

`func (o *HeartRateModel) GetSource() HeartRateSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HeartRateModel) GetSourceOk() (*HeartRateSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HeartRateModel) SetSource(v HeartRateSource)`

SetSource sets Source field to given value.


### GetTimestamp

`func (o *HeartRateModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HeartRateModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HeartRateModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


