# Go API client for tradewatch

Financial market data for Developers

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import tradewatch "github.com/tradewatch-io/go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `tradewatch.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), tradewatch.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `tradewatch.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), tradewatch.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `tradewatch.ContextOperationServerIndices` and `tradewatch.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), tradewatch.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), tradewatch.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.tradewatch.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**getUsage**](docs/AccountAPI.md#getusage) | **Get** /account/usage | Usage statistics
*CommoditiesAPI* | [**getQuote**](docs/CommoditiesAPI.md#getquote) | **Get** /commodities/symbols/{symbol} | Last Quote
*CommoditiesAPI* | [**getSymbols**](docs/CommoditiesAPI.md#getsymbols) | **Get** /commodities/symbols | Available Symbols
*CommoditiesAPI* | [**getTypes**](docs/CommoditiesAPI.md#gettypes) | **Get** /commodities/types | Available Types
*CryptoAPI* | [**convert**](docs/CryptoAPI.md#convert) | **Get** /crypto/convert/{from}/{to} | Conversion
*CryptoAPI* | [**getExchanges**](docs/CryptoAPI.md#getexchanges) | **Get** /crypto/exchanges | Available Exchanges
*CryptoAPI* | [**getQuote**](docs/CryptoAPI.md#getquote) | **Get** /crypto/symbols/{symbol} | Last Quote
*CryptoAPI* | [**getSymbols**](docs/CryptoAPI.md#getsymbols) | **Get** /crypto/symbols | Available Symbols
*CurrenciesAPI* | [**convert**](docs/CurrenciesAPI.md#convert) | **Get** /currencies/convert/{from}/{to} | Conversion
*CurrenciesAPI* | [**getQuote**](docs/CurrenciesAPI.md#getquote) | **Get** /currencies/symbols/{symbol} | Last Quote
*CurrenciesAPI* | [**getSymbols**](docs/CurrenciesAPI.md#getsymbols) | **Get** /currencies/symbols | Available Symbols
*IndicesAPI* | [**getQuote**](docs/IndicesAPI.md#getquote) | **Get** /indices/symbols/{symbol} | Last Quote
*IndicesAPI* | [**getSymbols**](docs/IndicesAPI.md#getsymbols) | **Get** /indices/symbols | Available Symbols
*StocksAPI* | [**getCountries**](docs/StocksAPI.md#getcountries) | **Get** /stocks/countries | Available Countries
*StocksAPI* | [**getQuote**](docs/StocksAPI.md#getquote) | **Get** /stocks/symbols/{symbol} | Last Quote
*StocksAPI* | [**getSymbols**](docs/StocksAPI.md#getsymbols) | **Get** /stocks/symbols | Available Symbols


## Documentation For Models

 - [AccountUsageStatisticsInterval](docs/AccountUsageStatisticsInterval.md)
 - [ApiUsageDataTransfer](docs/ApiUsageDataTransfer.md)
 - [ApiUsageEntry](docs/ApiUsageEntry.md)
 - [CommodityType](docs/CommodityType.md)
 - [CommodityTypeObj](docs/CommodityTypeObj.md)
 - [CommodityTypesList](docs/CommodityTypesList.md)
 - [Conversion](docs/Conversion.md)
 - [ConversionInfo](docs/ConversionInfo.md)
 - [ConversionQuery](docs/ConversionQuery.md)
 - [CountriesList](docs/CountriesList.md)
 - [Country](docs/Country.md)
 - [CountryObj](docs/CountryObj.md)
 - [CryptoConversion](docs/CryptoConversion.md)
 - [CryptoConversionQuery](docs/CryptoConversionQuery.md)
 - [CryptoExchangeItem](docs/CryptoExchangeItem.md)
 - [CryptoExchangesList](docs/CryptoExchangesList.md)
 - [CursorPageTCustomizedSymbolsOutFull](docs/CursorPageTCustomizedSymbolsOutFull.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [ErrorResponseBody](docs/ErrorResponseBody.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [LastQuote](docs/LastQuote.md)
 - [SymbolsListMode](docs/SymbolsListMode.md)
 - [SymbolsOutFull](docs/SymbolsOutFull.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_key_header

- **Type**: API key
- **API key parameter name**: api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		tradewatch.ContextAPIKeys,
		map[string]tradewatch.APIKey{
			"api-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### api_key_query

- **Type**: API key
- **API key parameter name**: api-key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		tradewatch.ContextAPIKeys,
		map[string]tradewatch.APIKey{
			"api-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



