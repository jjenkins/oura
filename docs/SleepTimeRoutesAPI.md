# \SleepTimeRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet**](SleepTimeRoutesAPI.md#MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet) | **Get** /v2/usercollection/sleep_time | Multiple Sleep Time Documents
[**SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet**](SleepTimeRoutesAPI.md#SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet) | **Get** /v2/usercollection/sleep_time/{document_id} | Single Sleep Time Document



## MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet

> MultiDocumentResponseSleepTimeModel MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Sleep Time Documents

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
	resp, r, err := apiClient.SleepTimeRoutesAPI.MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SleepTimeRoutesAPI.MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet`: MultiDocumentResponseSleepTimeModel
	fmt.Fprintf(os.Stdout, "Response from `SleepTimeRoutesAPI.MultipleSleepTimeDocumentsV2UsercollectionSleepTimeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleSleepTimeDocumentsV2UsercollectionSleepTimeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseSleepTimeModel**](MultiDocumentResponseSleepTimeModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet

> SleepTimeModel SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet(ctx, documentId).Execute()

Single Sleep Time Document

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
	resp, r, err := apiClient.SleepTimeRoutesAPI.SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SleepTimeRoutesAPI.SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet`: SleepTimeModel
	fmt.Fprintf(os.Stdout, "Response from `SleepTimeRoutesAPI.SingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleSleepTimeDocumentV2UsercollectionSleepTimeDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SleepTimeModel**](SleepTimeModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

