# \SleepRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleSleepDocumentsV2UsercollectionSleepGet**](SleepRoutesAPI.md#MultipleSleepDocumentsV2UsercollectionSleepGet) | **Get** /v2/usercollection/sleep | Multiple Sleep Documents
[**SingleSleepDocumentV2UsercollectionSleepDocumentIdGet**](SleepRoutesAPI.md#SingleSleepDocumentV2UsercollectionSleepDocumentIdGet) | **Get** /v2/usercollection/sleep/{document_id} | Single Sleep Document



## MultipleSleepDocumentsV2UsercollectionSleepGet

> MultiDocumentResponseSleepModel MultipleSleepDocumentsV2UsercollectionSleepGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Sleep Documents

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
	startDate := *openapiclient.NewStartDate() // StartDate |  (optional)
	endDate := *openapiclient.NewEndDate1() // EndDate1 |  (optional)
	nextToken := "nextToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SleepRoutesAPI.MultipleSleepDocumentsV2UsercollectionSleepGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SleepRoutesAPI.MultipleSleepDocumentsV2UsercollectionSleepGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleSleepDocumentsV2UsercollectionSleepGet`: MultiDocumentResponseSleepModel
	fmt.Fprintf(os.Stdout, "Response from `SleepRoutesAPI.MultipleSleepDocumentsV2UsercollectionSleepGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleSleepDocumentsV2UsercollectionSleepGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseSleepModel**](MultiDocumentResponseSleepModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleSleepDocumentV2UsercollectionSleepDocumentIdGet

> SleepModel SingleSleepDocumentV2UsercollectionSleepDocumentIdGet(ctx, documentId).Execute()

Single Sleep Document

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
	resp, r, err := apiClient.SleepRoutesAPI.SingleSleepDocumentV2UsercollectionSleepDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SleepRoutesAPI.SingleSleepDocumentV2UsercollectionSleepDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleSleepDocumentV2UsercollectionSleepDocumentIdGet`: SleepModel
	fmt.Fprintf(os.Stdout, "Response from `SleepRoutesAPI.SingleSleepDocumentV2UsercollectionSleepDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleSleepDocumentV2UsercollectionSleepDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SleepModel**](SleepModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

