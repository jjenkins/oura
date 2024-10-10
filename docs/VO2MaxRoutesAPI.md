# \VO2MaxRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet**](VO2MaxRoutesAPI.md#MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet) | **Get** /v2/usercollection/vO2_max | Multiple Vo2 Max Documents
[**SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet**](VO2MaxRoutesAPI.md#SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet) | **Get** /v2/usercollection/vO2_max/{document_id} | Single Vo2 Max Document



## MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet

> MultiDocumentResponseVO2MaxModel MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Vo2 Max Documents

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
	resp, r, err := apiClient.VO2MaxRoutesAPI.MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VO2MaxRoutesAPI.MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet`: MultiDocumentResponseVO2MaxModel
	fmt.Fprintf(os.Stdout, "Response from `VO2MaxRoutesAPI.MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleVO2MaxDocumentsV2UsercollectionVO2MaxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseVO2MaxModel**](MultiDocumentResponseVO2MaxModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet

> VO2MaxModel SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet(ctx, documentId).Execute()

Single Vo2 Max Document

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
	resp, r, err := apiClient.VO2MaxRoutesAPI.SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VO2MaxRoutesAPI.SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet`: VO2MaxModel
	fmt.Fprintf(os.Stdout, "Response from `VO2MaxRoutesAPI.SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VO2MaxModel**](VO2MaxModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

