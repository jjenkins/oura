# WebhookSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CallbackUrl** | **string** |  | 
**EventType** | [**WebhookOperation**](WebhookOperation.md) |  | 
**DataType** | [**ExtApiV2DataType**](ExtApiV2DataType.md) |  | 
**ExpirationTime** | **time.Time** |  | 

## Methods

### NewWebhookSubscriptionModel

`func NewWebhookSubscriptionModel(id string, callbackUrl string, eventType WebhookOperation, dataType ExtApiV2DataType, expirationTime time.Time, ) *WebhookSubscriptionModel`

NewWebhookSubscriptionModel instantiates a new WebhookSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionModelWithDefaults

`func NewWebhookSubscriptionModelWithDefaults() *WebhookSubscriptionModel`

NewWebhookSubscriptionModelWithDefaults instantiates a new WebhookSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookSubscriptionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSubscriptionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSubscriptionModel) SetId(v string)`

SetId sets Id field to given value.


### GetCallbackUrl

`func (o *WebhookSubscriptionModel) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *WebhookSubscriptionModel) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *WebhookSubscriptionModel) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetEventType

`func (o *WebhookSubscriptionModel) GetEventType() WebhookOperation`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookSubscriptionModel) GetEventTypeOk() (*WebhookOperation, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookSubscriptionModel) SetEventType(v WebhookOperation)`

SetEventType sets EventType field to given value.


### GetDataType

`func (o *WebhookSubscriptionModel) GetDataType() ExtApiV2DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WebhookSubscriptionModel) GetDataTypeOk() (*ExtApiV2DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WebhookSubscriptionModel) SetDataType(v ExtApiV2DataType)`

SetDataType sets DataType field to given value.


### GetExpirationTime

`func (o *WebhookSubscriptionModel) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *WebhookSubscriptionModel) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *WebhookSubscriptionModel) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


