# \EnhancedTagRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet**](EnhancedTagRoutesAPI.md#MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet) | **Get** /v2/usercollection/enhanced_tag | Multiple Enhanced Tag Documents
[**SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet**](EnhancedTagRoutesAPI.md#SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet) | **Get** /v2/usercollection/enhanced_tag/{document_id} | Single Enhanced Tag Document



## MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet

> MultiDocumentResponseEnhancedTagModel MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Enhanced Tag Documents

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
	resp, r, err := apiClient.EnhancedTagRoutesAPI.MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTagRoutesAPI.MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet`: MultiDocumentResponseEnhancedTagModel
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTagRoutesAPI.MultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleEnhancedTagDocumentsV2UsercollectionEnhancedTagGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseEnhancedTagModel**](MultiDocumentResponseEnhancedTagModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet

> EnhancedTagModel SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet(ctx, documentId).Execute()

Single Enhanced Tag Document

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
	resp, r, err := apiClient.EnhancedTagRoutesAPI.SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTagRoutesAPI.SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet`: EnhancedTagModel
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTagRoutesAPI.SingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleEnhancedTagDocumentV2UsercollectionEnhancedTagDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnhancedTagModel**](EnhancedTagModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

