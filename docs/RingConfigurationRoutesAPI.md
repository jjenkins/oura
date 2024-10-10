# \RingConfigurationRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet**](RingConfigurationRoutesAPI.md#MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet) | **Get** /v2/usercollection/ring_configuration | Multiple Ring Configuration Documents
[**SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet**](RingConfigurationRoutesAPI.md#SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet) | **Get** /v2/usercollection/ring_configuration/{document_id} | Single Ring Configuration Document



## MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet

> MultiDocumentResponseRingConfigurationModel MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet(ctx).NextToken(nextToken).Execute()

Multiple Ring Configuration Documents

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
	nextToken := "nextToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RingConfigurationRoutesAPI.MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet(context.Background()).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RingConfigurationRoutesAPI.MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet`: MultiDocumentResponseRingConfigurationModel
	fmt.Fprintf(os.Stdout, "Response from `RingConfigurationRoutesAPI.MultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleRingConfigurationDocumentsV2UsercollectionRingConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseRingConfigurationModel**](MultiDocumentResponseRingConfigurationModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet

> RingConfigurationModel SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet(ctx, documentId).Execute()

Single Ring Configuration Document

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
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RingConfigurationRoutesAPI.SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RingConfigurationRoutesAPI.SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet`: RingConfigurationModel
	fmt.Fprintf(os.Stdout, "Response from `RingConfigurationRoutesAPI.SingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleRingConfigurationDocumentV2UsercollectionRingConfigurationDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RingConfigurationModel**](RingConfigurationModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

