# \DailySleepRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailySleepDocumentsV2UsercollectionDailySleepGet**](DailySleepRoutesAPI.md#MultipleDailySleepDocumentsV2UsercollectionDailySleepGet) | **Get** /v2/usercollection/daily_sleep | Multiple Daily Sleep Documents
[**SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet**](DailySleepRoutesAPI.md#SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet) | **Get** /v2/usercollection/daily_sleep/{document_id} | Single Daily Sleep Document



## MultipleDailySleepDocumentsV2UsercollectionDailySleepGet

> MultiDocumentResponseDailySleepModel MultipleDailySleepDocumentsV2UsercollectionDailySleepGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Sleep Documents

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
	resp, r, err := apiClient.DailySleepRoutesAPI.MultipleDailySleepDocumentsV2UsercollectionDailySleepGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailySleepRoutesAPI.MultipleDailySleepDocumentsV2UsercollectionDailySleepGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailySleepDocumentsV2UsercollectionDailySleepGet`: MultiDocumentResponseDailySleepModel
	fmt.Fprintf(os.Stdout, "Response from `DailySleepRoutesAPI.MultipleDailySleepDocumentsV2UsercollectionDailySleepGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailySleepDocumentsV2UsercollectionDailySleepGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailySleepModel**](MultiDocumentResponseDailySleepModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet

> DailySleepModel SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet(ctx, documentId).Execute()

Single Daily Sleep Document

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
	resp, r, err := apiClient.DailySleepRoutesAPI.SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailySleepRoutesAPI.SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet`: DailySleepModel
	fmt.Fprintf(os.Stdout, "Response from `DailySleepRoutesAPI.SingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailySleepDocumentV2UsercollectionDailySleepDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailySleepModel**](DailySleepModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

