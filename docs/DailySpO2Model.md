# DailySpO2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** |  | 
**Spo2Percentage** | [**NullableDailySpO2AggregatedValuesModel**](DailySpO2AggregatedValuesModel.md) |  | 

## Methods

### NewDailySpO2Model

`func NewDailySpO2Model(id string, day string, spo2Percentage NullableDailySpO2AggregatedValuesModel, ) *DailySpO2Model`

NewDailySpO2Model instantiates a new DailySpO2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailySpO2ModelWithDefaults

`func NewDailySpO2ModelWithDefaults() *DailySpO2Model`

NewDailySpO2ModelWithDefaults instantiates a new DailySpO2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DailySpO2Model) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DailySpO2Model) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DailySpO2Model) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *DailySpO2Model) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DailySpO2Model) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DailySpO2Model) SetDay(v string)`

SetDay sets Day field to given value.


### GetSpo2Percentage

`func (o *DailySpO2Model) GetSpo2Percentage() DailySpO2AggregatedValuesModel`

GetSpo2Percentage returns the Spo2Percentage field if non-nil, zero value otherwise.

### GetSpo2PercentageOk

`func (o *DailySpO2Model) GetSpo2PercentageOk() (*DailySpO2AggregatedValuesModel, bool)`

GetSpo2PercentageOk returns a tuple with the Spo2Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpo2Percentage

`func (o *DailySpO2Model) SetSpo2Percentage(v DailySpO2AggregatedValuesModel)`

SetSpo2Percentage sets Spo2Percentage field to given value.


### SetSpo2PercentageNil

`func (o *DailySpO2Model) SetSpo2PercentageNil(b bool)`

 SetSpo2PercentageNil sets the value for Spo2Percentage to be an explicit nil

### UnsetSpo2Percentage
`func (o *DailySpO2Model) UnsetSpo2Percentage()`

UnsetSpo2Percentage ensures that no value is present for Spo2Percentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


