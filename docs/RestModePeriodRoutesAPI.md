# \RestModePeriodRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet**](RestModePeriodRoutesAPI.md#MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet) | **Get** /v2/usercollection/rest_mode_period | Multiple Rest Mode Period Documents
[**SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet**](RestModePeriodRoutesAPI.md#SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet) | **Get** /v2/usercollection/rest_mode_period/{document_id} | Single Rest Mode Period Document



## MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet

> MultiDocumentResponseRestModePeriodModel MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Rest Mode Period Documents

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
	resp, r, err := apiClient.RestModePeriodRoutesAPI.MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestModePeriodRoutesAPI.MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet`: MultiDocumentResponseRestModePeriodModel
	fmt.Fprintf(os.Stdout, "Response from `RestModePeriodRoutesAPI.MultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleRestModePeriodDocumentsV2UsercollectionRestModePeriodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseRestModePeriodModel**](MultiDocumentResponseRestModePeriodModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet

> RestModePeriodModel SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet(ctx, documentId).Execute()

Single Rest Mode Period Document

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
	resp, r, err := apiClient.RestModePeriodRoutesAPI.SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestModePeriodRoutesAPI.SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet`: RestModePeriodModel
	fmt.Fprintf(os.Stdout, "Response from `RestModePeriodRoutesAPI.SingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleRestModePeriodDocumentV2UsercollectionRestModePeriodDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestModePeriodModel**](RestModePeriodModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

