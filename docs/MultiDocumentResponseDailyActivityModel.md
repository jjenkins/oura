# MultiDocumentResponseDailyActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DailyActivityModel**](DailyActivityModel.md) |  | 
**NextToken** | **NullableString** |  | 

## Methods

### NewMultiDocumentResponseDailyActivityModel

`func NewMultiDocumentResponseDailyActivityModel(data []DailyActivityModel, nextToken NullableString, ) *MultiDocumentResponseDailyActivityModel`

NewMultiDocumentResponseDailyActivityModel instantiates a new MultiDocumentResponseDailyActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDocumentResponseDailyActivityModelWithDefaults

`func NewMultiDocumentResponseDailyActivityModelWithDefaults() *MultiDocumentResponseDailyActivityModel`

NewMultiDocumentResponseDailyActivityModelWithDefaults instantiates a new MultiDocumentResponseDailyActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MultiDocumentResponseDailyActivityModel) GetData() []DailyActivityModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultiDocumentResponseDailyActivityModel) GetDataOk() (*[]DailyActivityModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultiDocumentResponseDailyActivityModel) SetData(v []DailyActivityModel)`

SetData sets Data field to given value.


### GetNextToken

`func (o *MultiDocumentResponseDailyActivityModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MultiDocumentResponseDailyActivityModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MultiDocumentResponseDailyActivityModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *MultiDocumentResponseDailyActivityModel) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *MultiDocumentResponseDailyActivityModel) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


