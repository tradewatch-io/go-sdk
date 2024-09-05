# \IndicesAPI

All URIs are relative to *https://api.tradewatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getQuote**](IndicesAPI.md#getQuote) | **Get** /indices/symbols/{symbol} | Last Quote
[**getSymbols**](IndicesAPI.md#getSymbols) | **Get** /indices/symbols | Available Symbols



## getQuote

> LastQuote getQuote(ctx, symbol).Precision(precision).Execute()

Last Quote



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
	symbol := "symbol_example" // string | 
	precision := int32(56) // int32 |  (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicesAPI.getQuote(context.Background(), symbol).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicesAPI.getQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `getQuote`: LastQuote
	fmt.Fprintf(os.Stdout, "Response from `IndicesAPI.getQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apigetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **precision** | **int32** |  | [default to 5]

### Return type

[**LastQuote**](LastQuote.md)

### Authorization

[api_key_query](../README.md#api_key_query), [api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## getSymbols

> CursorPageTCustomizedSymbolsOutFull getSymbols(ctx).Mode(mode).Size(size).Page(page).Cursor(cursor).Execute()

Available Symbols



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
	mode := openapiclient.SymbolsListMode("full") // SymbolsListMode | Listing mode
	size := int32(56) // int32 | Page offset (optional) (default to 50)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	cursor := "cursor_example" // string | Cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicesAPI.getSymbols(context.Background()).Mode(mode).Size(size).Page(page).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicesAPI.getSymbols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `getSymbols`: CursorPageTCustomizedSymbolsOutFull
	fmt.Fprintf(os.Stdout, "Response from `IndicesAPI.getSymbols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apigetSymbolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | [**SymbolsListMode**](SymbolsListMode.md) | Listing mode | 
 **size** | **int32** | Page offset | [default to 50]
 **page** | **int32** | Page number | [default to 1]
 **cursor** | **string** | Cursor for the next page | 

### Return type

[**CursorPageTCustomizedSymbolsOutFull**](CursorPageTCustomizedSymbolsOutFull.md)

### Authorization

[api_key_query](../README.md#api_key_query), [api_key_header](../README.md#api_key_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
