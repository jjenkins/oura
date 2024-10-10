# \DailySpo2RoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get**](DailySpo2RoutesAPI.md#MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get) | **Get** /v2/usercollection/daily_spo2 | Multiple Daily Spo2 Documents
[**SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet**](DailySpo2RoutesAPI.md#SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet) | **Get** /v2/usercollection/daily_spo2/{document_id} | Single Daily Spo2 Document



## MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get

> MultiDocumentResponseDailySpO2Model MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Daily Spo2 Documents

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
	resp, r, err := apiClient.DailySpo2RoutesAPI.MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailySpo2RoutesAPI.MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get`: MultiDocumentResponseDailySpO2Model
	fmt.Fprintf(os.Stdout, "Response from `DailySpo2RoutesAPI.MultipleDailySpo2DocumentsV2UsercollectionDailySpo2Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleDailySpo2DocumentsV2UsercollectionDailySpo2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseDailySpO2Model**](MultiDocumentResponseDailySpO2Model.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet

> DailySpO2Model SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet(ctx, documentId).Execute()

Single Daily Spo2 Document

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
	resp, r, err := apiClient.DailySpo2RoutesAPI.SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailySpo2RoutesAPI.SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet`: DailySpO2Model
	fmt.Fprintf(os.Stdout, "Response from `DailySpo2RoutesAPI.SingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleDailySpo2DocumentV2UsercollectionDailySpo2DocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailySpO2Model**](DailySpO2Model.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

