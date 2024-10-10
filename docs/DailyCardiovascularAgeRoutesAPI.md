# \DailyCardiovascularAgeRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet**](DailyCardiovascularAgeRoutesAPI.md#MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet) | **Get** /v2/usercollection/daily_cardiovascular_age | Multiple Daily Cardiovascular Age Documents
[**SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet**](DailyCardiovascularAgeRoutesAPI.md#SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet) | **Get** /v2/usercollection/daily_cardiovascular_age/{document_id} | Single Daily Cardiovascular Age Document



## MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet

> MultiDocumentResponseDailyCardiovascularAgeModel MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Cardiovascular Age Documents

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
	resp, r, err := apiClient.DailyCardiovascularAgeRoutesAPI.MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyCardiovascularAgeRoutesAPI.MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet`: MultiDocumentResponseDailyCardiovascularAgeModel
	fmt.Fprintf(os.Stdout, "Response from `DailyCardiovascularAgeRoutesAPI.MultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailyCardiovascularAgeDocumentsV2UsercollectionDailyCardiovascularAgeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailyCardiovascularAgeModel**](MultiDocumentResponseDailyCardiovascularAgeModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet

> DailyCardiovascularAgeModel SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet(ctx, documentId).Execute()

Single Daily Cardiovascular Age Document

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
	resp, r, err := apiClient.DailyCardiovascularAgeRoutesAPI.SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyCardiovascularAgeRoutesAPI.SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet`: DailyCardiovascularAgeModel
	fmt.Fprintf(os.Stdout, "Response from `DailyCardiovascularAgeRoutesAPI.SingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailyCardiovascularAgeDocumentV2UsercollectionDailyCardiovascularAgeDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailyCardiovascularAgeModel**](DailyCardiovascularAgeModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

