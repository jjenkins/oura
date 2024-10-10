# \DailyStressRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailyStressDocumentsV2UsercollectionDailyStressGet**](DailyStressRoutesAPI.md#MultipleDailyStressDocumentsV2UsercollectionDailyStressGet) | **Get** /v2/usercollection/daily_stress | Multiple Daily Stress Documents
[**SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet**](DailyStressRoutesAPI.md#SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet) | **Get** /v2/usercollection/daily_stress/{document_id} | Single Daily Stress Document



## MultipleDailyStressDocumentsV2UsercollectionDailyStressGet

> MultiDocumentResponseDailyStressModel MultipleDailyStressDocumentsV2UsercollectionDailyStressGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Stress Documents

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
	resp, r, err := apiClient.DailyStressRoutesAPI.MultipleDailyStressDocumentsV2UsercollectionDailyStressGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyStressRoutesAPI.MultipleDailyStressDocumentsV2UsercollectionDailyStressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailyStressDocumentsV2UsercollectionDailyStressGet`: MultiDocumentResponseDailyStressModel
	fmt.Fprintf(os.Stdout, "Response from `DailyStressRoutesAPI.MultipleDailyStressDocumentsV2UsercollectionDailyStressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailyStressDocumentsV2UsercollectionDailyStressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailyStressModel**](MultiDocumentResponseDailyStressModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet

> DailyStressModel SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet(ctx, documentId).Execute()

Single Daily Stress Document

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
	resp, r, err := apiClient.DailyStressRoutesAPI.SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyStressRoutesAPI.SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet`: DailyStressModel
	fmt.Fprintf(os.Stdout, "Response from `DailyStressRoutesAPI.SingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailyStressDocumentV2UsercollectionDailyStressDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailyStressModel**](DailyStressModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

