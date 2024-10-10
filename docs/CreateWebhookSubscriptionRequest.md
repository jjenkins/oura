# CreateWebhookSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** |  | 
**VerificationToken** | **string** |  | 
**EventType** | [**WebhookOperation**](WebhookOperation.md) |  | 
**DataType** | [**ExtApiV2DataType**](ExtApiV2DataType.md) |  | 

## Methods

### NewCreateWebhookSubscriptionRequest

`func NewCreateWebhookSubscriptionRequest(callbackUrl string, verificationToken string, eventType WebhookOperation, dataType ExtApiV2DataType, ) *CreateWebhookSubscriptionRequest`

NewCreateWebhookSubscriptionRequest instantiates a new CreateWebhookSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookSubscriptionRequestWithDefaults

`func NewCreateWebhookSubscriptionRequestWithDefaults() *CreateWebhookSubscriptionRequest`

NewCreateWebhookSubscriptionRequestWithDefaults instantiates a new CreateWebhookSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *CreateWebhookSubscriptionRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateWebhookSubscriptionRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateWebhookSubscriptionRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetVerificationToken

`func (o *CreateWebhookSubscriptionRequest) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *CreateWebhookSubscriptionRequest) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *CreateWebhookSubscriptionRequest) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.


### GetEventType

`func (o *CreateWebhookSubscriptionRequest) GetEventType() WebhookOperation`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CreateWebhookSubscriptionRequest) GetEventTypeOk() (*WebhookOperation, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CreateWebhookSubscriptionRequest) SetEventType(v WebhookOperation)`

SetEventType sets EventType field to given value.


### GetDataType

`func (o *CreateWebhookSubscriptionRequest) GetDataType() ExtApiV2DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CreateWebhookSubscriptionRequest) GetDataTypeOk() (*ExtApiV2DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CreateWebhookSubscriptionRequest) SetDataType(v ExtApiV2DataType)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


