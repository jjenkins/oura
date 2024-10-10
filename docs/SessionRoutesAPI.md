# \SessionRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleSessionDocumentsV2UsercollectionSessionGet**](SessionRoutesAPI.md#MultipleSessionDocumentsV2UsercollectionSessionGet) | **Get** /v2/usercollection/session | Multiple Session Documents
[**SingleSessionDocumentV2UsercollectionSessionDocumentIdGet**](SessionRoutesAPI.md#SingleSessionDocumentV2UsercollectionSessionDocumentIdGet) | **Get** /v2/usercollection/session/{document_id} | Single Session Document



## MultipleSessionDocumentsV2UsercollectionSessionGet

> MultiDocumentResponseSessionModel MultipleSessionDocumentsV2UsercollectionSessionGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Session Documents

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
	resp, r, err := apiClient.SessionRoutesAPI.MultipleSessionDocumentsV2UsercollectionSessionGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRoutesAPI.MultipleSessionDocumentsV2UsercollectionSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleSessionDocumentsV2UsercollectionSessionGet`: MultiDocumentResponseSessionModel
	fmt.Fprintf(os.Stdout, "Response from `SessionRoutesAPI.MultipleSessionDocumentsV2UsercollectionSessionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleSessionDocumentsV2UsercollectionSessionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseSessionModel**](MultiDocumentResponseSessionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleSessionDocumentV2UsercollectionSessionDocumentIdGet

> SessionModel SingleSessionDocumentV2UsercollectionSessionDocumentIdGet(ctx, documentId).Execute()

Single Session Document

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
	resp, r, err := apiClient.SessionRoutesAPI.SingleSessionDocumentV2UsercollectionSessionDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRoutesAPI.SingleSessionDocumentV2UsercollectionSessionDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleSessionDocumentV2UsercollectionSessionDocumentIdGet`: SessionModel
	fmt.Fprintf(os.Stdout, "Response from `SessionRoutesAPI.SingleSessionDocumentV2UsercollectionSessionDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleSessionDocumentV2UsercollectionSessionDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

