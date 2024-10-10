# \DailyActivityRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet**](DailyActivityRoutesAPI.md#MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet) | **Get** /v2/usercollection/daily_activity | Multiple Daily Activity Documents
[**SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet**](DailyActivityRoutesAPI.md#SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet) | **Get** /v2/usercollection/daily_activity/{document_id} | Single Daily Activity Document



## MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet

> MultiDocumentResponseDailyActivityModel MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Activity Documents

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
	resp, r, err := apiClient.DailyActivityRoutesAPI.MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyActivityRoutesAPI.MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet`: MultiDocumentResponseDailyActivityModel
	fmt.Fprintf(os.Stdout, "Response from `DailyActivityRoutesAPI.MultipleDailyActivityDocumentsV2UsercollectionDailyActivityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailyActivityDocumentsV2UsercollectionDailyActivityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailyActivityModel**](MultiDocumentResponseDailyActivityModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet

> DailyActivityModel SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet(ctx, documentId).Execute()

Single Daily Activity Document

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
	resp, r, err := apiClient.DailyActivityRoutesAPI.SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyActivityRoutesAPI.SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet`: DailyActivityModel
	fmt.Fprintf(os.Stdout, "Response from `DailyActivityRoutesAPI.SingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailyActivityDocumentV2UsercollectionDailyActivityDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailyActivityModel**](DailyActivityModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

