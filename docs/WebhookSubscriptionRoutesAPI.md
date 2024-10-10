# \WebhookSubscriptionRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookSubscriptionV2WebhookSubscriptionPost**](WebhookSubscriptionRoutesAPI.md#CreateWebhookSubscriptionV2WebhookSubscriptionPost) | **Post** /v2/webhook/subscription | Create Webhook Subscription
[**DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete**](WebhookSubscriptionRoutesAPI.md#DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete) | **Delete** /v2/webhook/subscription/{id} | Delete Webhook Subscription
[**GetWebhookSubscriptionV2WebhookSubscriptionIdGet**](WebhookSubscriptionRoutesAPI.md#GetWebhookSubscriptionV2WebhookSubscriptionIdGet) | **Get** /v2/webhook/subscription/{id} | Get Webhook Subscription
[**ListWebhookSubscriptionsV2WebhookSubscriptionGet**](WebhookSubscriptionRoutesAPI.md#ListWebhookSubscriptionsV2WebhookSubscriptionGet) | **Get** /v2/webhook/subscription | List Webhook Subscriptions
[**RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut**](WebhookSubscriptionRoutesAPI.md#RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut) | **Put** /v2/webhook/subscription/renew/{id} | Renew Webhook Subscription
[**UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut**](WebhookSubscriptionRoutesAPI.md#UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut) | **Put** /v2/webhook/subscription/{id} | Update Webhook Subscription



## CreateWebhookSubscriptionV2WebhookSubscriptionPost

> WebhookSubscriptionModel CreateWebhookSubscriptionV2WebhookSubscriptionPost(ctx).CreateWebhookSubscriptionRequest(createWebhookSubscriptionRequest).Execute()

Create Webhook Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createWebhookSubscriptionRequest := *openapiclient.NewCreateWebhookSubscriptionRequest("CallbackUrl_example", "VerificationToken_example", openapiclient.WebhookOperation("create"), openapiclient.ExtApiV2DataType("tag")) // CreateWebhookSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionRoutesAPI.CreateWebhookSubscriptionV2WebhookSubscriptionPost(context.Background()).CreateWebhookSubscriptionRequest(createWebhookSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.CreateWebhookSubscriptionV2WebhookSubscriptionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookSubscriptionV2WebhookSubscriptionPost`: WebhookSubscriptionModel
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionRoutesAPI.CreateWebhookSubscriptionV2WebhookSubscriptionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookSubscriptionV2WebhookSubscriptionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookSubscriptionRequest** | [**CreateWebhookSubscriptionRequest**](CreateWebhookSubscriptionRequest.md) |  | 

### Return type

[**WebhookSubscriptionModel**](WebhookSubscriptionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete

> DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete(ctx, id).Execute()

Delete Webhook Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookSubscriptionRoutesAPI.DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.DeleteWebhookSubscriptionV2WebhookSubscriptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSubscriptionV2WebhookSubscriptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSubscriptionV2WebhookSubscriptionIdGet

> WebhookSubscriptionModel GetWebhookSubscriptionV2WebhookSubscriptionIdGet(ctx, id).Execute()

Get Webhook Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionRoutesAPI.GetWebhookSubscriptionV2WebhookSubscriptionIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.GetWebhookSubscriptionV2WebhookSubscriptionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSubscriptionV2WebhookSubscriptionIdGet`: WebhookSubscriptionModel
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionRoutesAPI.GetWebhookSubscriptionV2WebhookSubscriptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionV2WebhookSubscriptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSubscriptionModel**](WebhookSubscriptionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSubscriptionsV2WebhookSubscriptionGet

> []WebhookSubscriptionModel ListWebhookSubscriptionsV2WebhookSubscriptionGet(ctx).Execute()

List Webhook Subscriptions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionRoutesAPI.ListWebhookSubscriptionsV2WebhookSubscriptionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.ListWebhookSubscriptionsV2WebhookSubscriptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookSubscriptionsV2WebhookSubscriptionGet`: []WebhookSubscriptionModel
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionRoutesAPI.ListWebhookSubscriptionsV2WebhookSubscriptionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookSubscriptionsV2WebhookSubscriptionGetRequest struct via the builder pattern


### Return type

[**[]WebhookSubscriptionModel**](WebhookSubscriptionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut

> WebhookSubscriptionModel RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut(ctx, id).Execute()

Renew Webhook Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionRoutesAPI.RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut`: WebhookSubscriptionModel
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionRoutesAPI.RenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewWebhookSubscriptionV2WebhookSubscriptionRenewIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSubscriptionModel**](WebhookSubscriptionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut

> WebhookSubscriptionModel UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut(ctx, id).UpdateWebhookSubscriptionRequest(updateWebhookSubscriptionRequest).Execute()

Update Webhook Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	updateWebhookSubscriptionRequest := *openapiclient.NewUpdateWebhookSubscriptionRequest("VerificationToken_example") // UpdateWebhookSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSubscriptionRoutesAPI.UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut(context.Background(), id).UpdateWebhookSubscriptionRequest(updateWebhookSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionRoutesAPI.UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut`: WebhookSubscriptionModel
	fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionRoutesAPI.UpdateWebhookSubscriptionV2WebhookSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSubscriptionV2WebhookSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookSubscriptionRequest** | [**UpdateWebhookSubscriptionRequest**](UpdateWebhookSubscriptionRequest.md) |  | 

### Return type

[**WebhookSubscriptionModel**](WebhookSubscriptionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

