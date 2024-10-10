# MultiDocumentResponseVO2MaxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VO2MaxModel**](VO2MaxModel.md) |  | 
**NextToken** | **NullableString** |  | 

## Methods

### NewMultiDocumentResponseVO2MaxModel

`func NewMultiDocumentResponseVO2MaxModel(data []VO2MaxModel, nextToken NullableString, ) *MultiDocumentResponseVO2MaxModel`

NewMultiDocumentResponseVO2MaxModel instantiates a new MultiDocumentResponseVO2MaxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDocumentResponseVO2MaxModelWithDefaults

`func NewMultiDocumentResponseVO2MaxModelWithDefaults() *MultiDocumentResponseVO2MaxModel`

NewMultiDocumentResponseVO2MaxModelWithDefaults instantiates a new MultiDocumentResponseVO2MaxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MultiDocumentResponseVO2MaxModel) GetData() []VO2MaxModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultiDocumentResponseVO2MaxModel) GetDataOk() (*[]VO2MaxModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultiDocumentResponseVO2MaxModel) SetData(v []VO2MaxModel)`

SetData sets Data field to given value.


### GetNextToken

`func (o *MultiDocumentResponseVO2MaxModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MultiDocumentResponseVO2MaxModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MultiDocumentResponseVO2MaxModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *MultiDocumentResponseVO2MaxModel) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *MultiDocumentResponseVO2MaxModel) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


