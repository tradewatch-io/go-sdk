# \AccountAPI

All URIs are relative to *https://api.tradewatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getUsage**](AccountAPI.md#getUsage) | **Get** /account/usage | Usage statistics



## getUsage

> map[string]interface{} getUsage(ctx).Limit(limit).Interval(interval).Execute()

Usage statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tradewatch-io/go-sdk"
)

func main() {
	limit := int32(56) // int32 |  (optional)
	interval := openapiclient.AccountUsageStatisticsInterval("1h") // AccountUsageStatisticsInterval |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.getUsage(context.Background()).Limit(limit).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.getUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `getUsage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.getUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apigetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **interval** | [**AccountUsageStatisticsInterval**](AccountUsageStatisticsInterval.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key_query](../README.md#api_key_query), [api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

