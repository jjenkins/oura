# UpdateWebhookSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationToken** | **string** |  | 
**CallbackUrl** | Pointer to **NullableString** |  | [optional] 
**EventType** | Pointer to [**NullableWebhookOperation**](WebhookOperation.md) |  | [optional] 
**DataType** | Pointer to [**NullableExtApiV2DataType**](ExtApiV2DataType.md) |  | [optional] 

## Methods

### NewUpdateWebhookSubscriptionRequest

`func NewUpdateWebhookSubscriptionRequest(verificationToken string, ) *UpdateWebhookSubscriptionRequest`

NewUpdateWebhookSubscriptionRequest instantiates a new UpdateWebhookSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookSubscriptionRequestWithDefaults

`func NewUpdateWebhookSubscriptionRequestWithDefaults() *UpdateWebhookSubscriptionRequest`

NewUpdateWebhookSubscriptionRequestWithDefaults instantiates a new UpdateWebhookSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationToken

`func (o *UpdateWebhookSubscriptionRequest) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *UpdateWebhookSubscriptionRequest) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *UpdateWebhookSubscriptionRequest) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.


### GetCallbackUrl

`func (o *UpdateWebhookSubscriptionRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *UpdateWebhookSubscriptionRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *UpdateWebhookSubscriptionRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *UpdateWebhookSubscriptionRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *UpdateWebhookSubscriptionRequest) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *UpdateWebhookSubscriptionRequest) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetEventType

`func (o *UpdateWebhookSubscriptionRequest) GetEventType() WebhookOperation`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UpdateWebhookSubscriptionRequest) GetEventTypeOk() (*WebhookOperation, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UpdateWebhookSubscriptionRequest) SetEventType(v WebhookOperation)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *UpdateWebhookSubscriptionRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *UpdateWebhookSubscriptionRequest) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *UpdateWebhookSubscriptionRequest) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetDataType

`func (o *UpdateWebhookSubscriptionRequest) GetDataType() ExtApiV2DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UpdateWebhookSubscriptionRequest) GetDataTypeOk() (*ExtApiV2DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UpdateWebhookSubscriptionRequest) SetDataType(v ExtApiV2DataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *UpdateWebhookSubscriptionRequest) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *UpdateWebhookSubscriptionRequest) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *UpdateWebhookSubscriptionRequest) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


