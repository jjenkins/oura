# \WorkoutRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleWorkoutDocumentsV2UsercollectionWorkoutGet**](WorkoutRoutesAPI.md#MultipleWorkoutDocumentsV2UsercollectionWorkoutGet) | **Get** /v2/usercollection/workout | Multiple Workout Documents
[**SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet**](WorkoutRoutesAPI.md#SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet) | **Get** /v2/usercollection/workout/{document_id} | Single Workout Document



## MultipleWorkoutDocumentsV2UsercollectionWorkoutGet

> MultiDocumentResponseWorkoutModel MultipleWorkoutDocumentsV2UsercollectionWorkoutGet(ctx).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()

Multiple Workout Documents

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
	resp, r, err := apiClient.WorkoutRoutesAPI.MultipleWorkoutDocumentsV2UsercollectionWorkoutGet(context.Background()).StartDate(startDate).EndDate(endDate).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkoutRoutesAPI.MultipleWorkoutDocumentsV2UsercollectionWorkoutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleWorkoutDocumentsV2UsercollectionWorkoutGet`: MultiDocumentResponseWorkoutModel
	fmt.Fprintf(os.Stdout, "Response from `WorkoutRoutesAPI.MultipleWorkoutDocumentsV2UsercollectionWorkoutGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleWorkoutDocumentsV2UsercollectionWorkoutGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | [**StartDate**](StartDate.md) |  | 
 **endDate** | [**EndDate1**](EndDate1.md) |  | 
 **nextToken** | **string** |  | 

### Return type

[**MultiDocumentResponseWorkoutModel**](MultiDocumentResponseWorkoutModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet

> WorkoutModel SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet(ctx, documentId).Execute()

Single Workout Document

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
	resp, r, err := apiClient.WorkoutRoutesAPI.SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkoutRoutesAPI.SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet`: WorkoutModel
	fmt.Fprintf(os.Stdout, "Response from `WorkoutRoutesAPI.SingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSingleWorkoutDocumentV2UsercollectionWorkoutDocumentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkoutModel**](WorkoutModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

