# VO2MaxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Day** | **string** | Day that the estimate belongs to. | 
**Timestamp** | **string** | Timestamp indicating when the estimate was created. | 
**Vo2Max** | **NullableFloat32** |  | 

## Methods

### NewVO2MaxModel

`func NewVO2MaxModel(id string, day string, timestamp string, vo2Max NullableFloat32, ) *VO2MaxModel`

NewVO2MaxModel instantiates a new VO2MaxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVO2MaxModelWithDefaults

`func NewVO2MaxModelWithDefaults() *VO2MaxModel`

NewVO2MaxModelWithDefaults instantiates a new VO2MaxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VO2MaxModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VO2MaxModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VO2MaxModel) SetId(v string)`

SetId sets Id field to given value.


### GetDay

`func (o *VO2MaxModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *VO2MaxModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *VO2MaxModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetTimestamp

`func (o *VO2MaxModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VO2MaxModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VO2MaxModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetVo2Max

`func (o *VO2MaxModel) GetVo2Max() float32`

GetVo2Max returns the Vo2Max field if non-nil, zero value otherwise.

### GetVo2MaxOk

`func (o *VO2MaxModel) GetVo2MaxOk() (*float32, bool)`

GetVo2MaxOk returns a tuple with the Vo2Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVo2Max

`func (o *VO2MaxModel) SetVo2Max(v float32)`

SetVo2Max sets Vo2Max field to given value.


### SetVo2MaxNil

`func (o *VO2MaxModel) SetVo2MaxNil(b bool)`

 SetVo2MaxNil sets the value for Vo2Max to be an explicit nil

### UnsetVo2Max
`func (o *VO2MaxModel) UnsetVo2Max()`

UnsetVo2Max ensures that no value is present for Vo2Max, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


