# \HeartRateRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipleHeartRateDocumentsV2UsercollectionHeartrateGet**](HeartRateRoutesAPI.md#MultipleHeartRateDocumentsV2UsercollectionHeartrateGet) | **Get** /v2/usercollection/heartrate | Multiple Heart Rate Documents



## MultipleHeartRateDocumentsV2UsercollectionHeartrateGet

> TimeSeriesResponseHeartRateModel MultipleHeartRateDocumentsV2UsercollectionHeartrateGet(ctx).StartDatetime(startDatetime).EndDatetime(endDatetime).NextToken(nextToken).Execute()

Multiple Heart Rate Documents

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startDatetime := time.Now() // time.Time |  (optional)
	endDatetime := time.Now() // time.Time |  (optional)
	nextToken := "nextToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeartRateRoutesAPI.MultipleHeartRateDocumentsV2UsercollectionHeartrateGet(context.Background()).StartDatetime(startDatetime).EndDatetime(endDatetime).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeartRateRoutesAPI.MultipleHeartRateDocumentsV2UsercollectionHeartrateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipleHeartRateDocumentsV2UsercollectionHeartrateGet`: TimeSeriesResponseHeartRateModel
	fmt.Fprintf(os.Stdout, "Response from `HeartRateRoutesAPI.MultipleHeartRateDocumentsV2UsercollectionHeartrateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipleHeartRateDocumentsV2UsercollectionHeartrateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDatetime** | **time.Time** |  | 
 **endDatetime** | **time.Time** |  | 
 **nextToken** | **string** |  | 

### Return type

[**TimeSeriesResponseHeartRateModel**](TimeSeriesResponseHeartRateModel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

