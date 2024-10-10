# \DailyResilienceRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet**](DailyResilienceRoutesAPI.md#MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet) | **Get** /v2/usercollection/daily_resilience | Multiple Daily Resilience Documents
[**SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet**](DailyResilienceRoutesAPI.md#SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet) | **Get** /v2/usercollection/daily_resilience/{document_id} | Single Daily Resilience Document



## MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet

> MultiDocumentResponseDailyResilienceModel MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Resilience Documents

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
	resp, r, err := apiClient.DailyResilienceRoutesAPI.MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyResilienceRoutesAPI.MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet`: MultiDocumentResponseDailyResilienceModel
	fmt.Fprintf(os.Stdout, "Response from `DailyResilienceRoutesAPI.MultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailyResilienceDocumentsV2UsercollectionDailyResilienceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailyResilienceModel**](MultiDocumentResponseDailyResilienceModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet

> DailyResilienceModel SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet(ctx, documentId).Execute()

Single Daily Resilience Document

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
	resp, r, err := apiClient.DailyResilienceRoutesAPI.SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyResilienceRoutesAPI.SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet`: DailyResilienceModel
	fmt.Fprintf(os.Stdout, "Response from `DailyResilienceRoutesAPI.SingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailyResilienceDocumentV2UsercollectionDailyResilienceDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailyResilienceModel**](DailyResilienceModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

