# MultiDocumentResponseDailyStressModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DailyStressModel**](DailyStressModel.md) |  | 
**NextToken** | **NullableString** |  | 

## Methods

### NewMultiDocumentResponseDailyStressModel

`func NewMultiDocumentResponseDailyStressModel(data []DailyStressModel, nextToken NullableString, ) *MultiDocumentResponseDailyStressModel`

NewMultiDocumentResponseDailyStressModel instantiates a new MultiDocumentResponseDailyStressModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDocumentResponseDailyStressModelWithDefaults

`func NewMultiDocumentResponseDailyStressModelWithDefaults() *MultiDocumentResponseDailyStressModel`

NewMultiDocumentResponseDailyStressModelWithDefaults instantiates a new MultiDocumentResponseDailyStressModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MultiDocumentResponseDailyStressModel) GetData() []DailyStressModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultiDocumentResponseDailyStressModel) GetDataOk() (*[]DailyStressModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultiDocumentResponseDailyStressModel) SetData(v []DailyStressModel)`

SetData sets Data field to given value.


### GetNextToken

`func (o *MultiDocumentResponseDailyStressModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MultiDocumentResponseDailyStressModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MultiDocumentResponseDailyStressModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *MultiDocumentResponseDailyStressModel) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *MultiDocumentResponseDailyStressModel) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


