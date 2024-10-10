# TimeSeriesResponseHeartRateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]HeartRateModel**](HeartRateModel.md) |  | 
**NextToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTimeSeriesResponseHeartRateModel

`func NewTimeSeriesResponseHeartRateModel(data []HeartRateModel, ) *TimeSeriesResponseHeartRateModel`

NewTimeSeriesResponseHeartRateModel instantiates a new TimeSeriesResponseHeartRateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesResponseHeartRateModelWithDefaults

`func NewTimeSeriesResponseHeartRateModelWithDefaults() *TimeSeriesResponseHeartRateModel`

NewTimeSeriesResponseHeartRateModelWithDefaults instantiates a new TimeSeriesResponseHeartRateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TimeSeriesResponseHeartRateModel) GetData() []HeartRateModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TimeSeriesResponseHeartRateModel) GetDataOk() (*[]HeartRateModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TimeSeriesResponseHeartRateModel) SetData(v []HeartRateModel)`

SetData sets Data field to given value.


### GetNextToken

`func (o *TimeSeriesResponseHeartRateModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *TimeSeriesResponseHeartRateModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *TimeSeriesResponseHeartRateModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *TimeSeriesResponseHeartRateModel) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *TimeSeriesResponseHeartRateModel) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *TimeSeriesResponseHeartRateModel) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


