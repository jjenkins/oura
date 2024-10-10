# MultiDocumentResponseRestModePeriodModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RestModePeriodModel**](RestModePeriodModel.md) |  | 
**NextToken** | **NullableString** |  | 

## Methods

### NewMultiDocumentResponseRestModePeriodModel

`func NewMultiDocumentResponseRestModePeriodModel(data []RestModePeriodModel, nextToken NullableString, ) *MultiDocumentResponseRestModePeriodModel`

NewMultiDocumentResponseRestModePeriodModel instantiates a new MultiDocumentResponseRestModePeriodModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDocumentResponseRestModePeriodModelWithDefaults

`func NewMultiDocumentResponseRestModePeriodModelWithDefaults() *MultiDocumentResponseRestModePeriodModel`

NewMultiDocumentResponseRestModePeriodModelWithDefaults instantiates a new MultiDocumentResponseRestModePeriodModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MultiDocumentResponseRestModePeriodModel) GetData() []RestModePeriodModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultiDocumentResponseRestModePeriodModel) GetDataOk() (*[]RestModePeriodModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultiDocumentResponseRestModePeriodModel) SetData(v []RestModePeriodModel)`

SetData sets Data field to given value.


### GetNextToken

`func (o *MultiDocumentResponseRestModePeriodModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MultiDocumentResponseRestModePeriodModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MultiDocumentResponseRestModePeriodModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *MultiDocumentResponseRestModePeriodModel) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *MultiDocumentResponseRestModePeriodModel) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


