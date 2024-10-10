# \PersonalInfoRoutesAPI

All URIs are relative to *https://api.None*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet**](PersonalInfoRoutesAPI.md#SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet) | **Get** /v2/usercollection/personal_info | Single Personal Info Document



## SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet

> PersonalInfoResponse SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet(ctx).Execute()

Single Personal Info Document

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalInfoRoutesAPI.SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalInfoRoutesAPI.SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet`: PersonalInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `PersonalInfoRoutesAPI.SinglePersonalInfoDocumentV2UsercollectionPersonalInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSinglePersonalInfoDocumentV2UsercollectionPersonalInfoGetRequest struct via the builder pattern


### Return type

[**PersonalInfoResponse**](PersonalInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

