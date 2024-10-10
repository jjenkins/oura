# \TagRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleTagDocumentsV2UsercollectionTagGet**](TagRoutesAPI.md#MultipleTagDocumentsV2UsercollectionTagGet) | **Get** /v2/usercollection/tag | Multiple Tag Documents
[**SingleTagDocumentV2UsercollectionTagDocumentIdGet**](TagRoutesAPI.md#SingleTagDocumentV2UsercollectionTagDocumentIdGet) | **Get** /v2/usercollection/tag/{document_id} | Single Tag Document



## MultipleTagDocumentsV2UsercollectionTagGet

> MultiDocumentResponseTagModel MultipleTagDocumentsV2UsercollectionTagGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Tag Documents

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
	endDate := *openapiclient.NewEndDate() // EndDate |  (optional)
	nextToken := "nextToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagRoutesAPI.MultipleTagDocumentsV2UsercollectionTagGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRoutesAPI.MultipleTagDocumentsV2UsercollectionTagGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleTagDocumentsV2UsercollectionTagGet`: MultiDocumentResponseTagModel
	fmt.Fprintf(os.Stdout, "Response from `TagRoutesAPI.MultipleTagDocumentsV2UsercollectionTagGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleTagDocumentsV2UsercollectionTagGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate**](EndDate.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseTagModel**](MultiDocumentResponseTagModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleTagDocumentV2UsercollectionTagDocumentIdGet

> TagModel SingleTagDocumentV2UsercollectionTagDocumentIdGet(ctx, documentId).Execute()

Single Tag Document

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
	resp, r, err := apiClient.TagRoutesAPI.SingleTagDocumentV2UsercollectionTagDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagRoutesAPI.SingleTagDocumentV2UsercollectionTagDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleTagDocumentV2UsercollectionTagDocumentIdGet`: TagModel
	fmt.Fprintf(os.Stdout, "Response from `TagRoutesAPI.SingleTagDocumentV2UsercollectionTagDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleTagDocumentV2UsercollectionTagDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagModel**](TagModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

