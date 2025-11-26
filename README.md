# TradeWatch.io Go SDK

![](https://tradewatch.io/)


<a href="https://tradewatch.io/">
  <img src="https://pub-e8bb70a6cc1844138d6a55fa4a44ba42.r2.dev/logo-purple.png" alt="TradeWatch.io logo" title="TradeWatch.io" align="right" height="60" />
</a>

Official SDK for the [TradeWatch.io API](https://tradewatch.io/docs/api-reference/introduction).

## Other SDKs
[![Python SDK](https://img.shields.io/badge/Python_SDK-3776AB?style=flat-square&logo=python&logoColor=white)](https://github.com/tradewatch-io/python-sdk)
[![TypeScript SDK](https://img.shields.io/badge/TypeScript_SDK-3178C6?style=flat-square&logo=typescript&logoColor=white)](https://github.com/tradewatch-io/typescript-sdk)
[![.NET SDK](https://img.shields.io/badge/.NET_SDK-512BD4?style=flat-square&logo=dotnet&logoColor=white)](https://github.com/tradewatch-io/dotnet-sdk)
[![PHP SDK](https://img.shields.io/badge/PHP_SDK-777BB4?style=flat-square&logo=php&logoColor=white)](https://github.com/tradewatch-io/php-sdk)
[![Java SDK](https://img.shields.io/badge/Java_SDK-ED8B00?style=flat-square&logo=openjdk&logoColor=white)](https://github.com/tradewatch-io/java-sdk)
[![Ruby SDK](https://img.shields.io/badge/Ruby_SDK-CC342D?style=flat-square&logo=ruby&logoColor=white)](https://github.com/tradewatch-io/ruby-sdk)
[![Swift SDK](https://img.shields.io/badge/Swift_SDK-FA7343?style=flat-square&logo=swift&logoColor=white)](https://github.com/tradewatch-io/swift-sdk)
[![Rust SDK](https://img.shields.io/badge/Rust_SDK-000000?style=flat-square&logo=rust&logoColor=white)](https://github.com/tradewatch-io/rust-sdk)
## What is TradeWatch.io?
TradeWatch.io is a market data platform and API for real-time and historical prices across crypto, stocks, indices, currencies, and commodities.

## Try the Interactive API Playground
Want to test endpoints without writing code first? Use the [TradeWatch Interactive API Playground](https://dash.tradewatch.io/api-explorer) to run requests directly in your browser.

[![TradeWatch Interactive API Playground](https://tradewatch.io/api-playground.png)](https://dash.tradewatch.io/api-explorer)

## Resources

- REST API reference: [https://tradewatch.io/docs/api-reference/introduction](https://tradewatch.io/docs/api-reference/introduction)
- WebSocket API reference: [https://tradewatch.io/docs/websocket-api/introduction](https://tradewatch.io/docs/websocket-api/introduction)
- Support channels: [https://tradewatch.io/docs/platform/support](https://tradewatch.io/docs/platform/support)

## Quick Start
1. Create an API key in the [TradeWatch Dashboard](https://dash.tradewatch.io/register).
2. Follow platform setup docs: [Getting started](https://tradewatch.io/docs/quickstart).


## Table of Contents

- [Documentation](#documentation)
- [Reference](#reference)
- [Usage](#usage)
- [Environments](#environments)
- [Errors](#errors)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)
  - [Explicit Null](#explicit-null)

## Documentation

API reference documentation is available [here](https://tradewatch.io/docs/api-reference/introduction).

## Reference

A full reference for this library is available [here](./reference.md).

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    client "sdk/client"
    option "sdk/option"
    sdk "sdk"
    context "context"
)

func do() {
    client := client.NewClient(
        option.WithAPIKey(
            "<value>",
        ),
    )
    request := &sdk.CryptoGetQuoteRequest{
        Symbol: "BTC-USD",
        Precision: sdk.Int(
            2,
        ),
    }
    client.Crypto.GetQuote(
        context.TODO(),
        request,
    )
}
```

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(api.Environments.Default),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. These errors are compatible
with the `errors.Is` and `errors.As` APIs, so you can access the error like so:

```go
response, err := client.Crypto.GetQuote(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, apiError) {
        // Do something with the API error ...
    }
    return err
}
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
// Specify default options applied on every request.
client := client.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.Crypto.GetQuote(
    ...,
    option.WithToken("<YOUR_API_KEY>"),
)
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call. (When the endpoint is paginated,
the raw HTTP response data will be included automatically in the Page response object.)

```go
response, err := client.Crypto.WithRawResponse.GetQuote(...)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v", response.Header)
fmt.Printf("Got status code: %d", response.StatusCode)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

If the `Retry-After` header is present in the response, the SDK will prioritize respecting its value exactly
over the default exponential backoff.

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewClient(
    option.WithMaxAttempts(1),
)

response, err := client.Crypto.GetQuote(
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

response, err := client.Crypto.GetQuote(ctx, ...)
```

### Explicit Null

If you want to send the explicit `null` JSON value through an optional parameter, you can use the setters\
that come with every object. Calling a setter method for a property will flip a bit in the `explicitFields`
bitfield for that setter's object; during serialization, any property with a flipped bit will have its
omittable status stripped, so zero or `nil` values will be sent explicitly rather than omitted altogether:

```go
type ExampleRequest struct {
    // An optional string parameter.
    Name *string `json:"name,omitempty" url:"-"`

    // Private bitmask of fields set to an explicit value and therefore not to be omitted
    explicitFields *big.Int `json:"-" url:"-"`
}

request := &ExampleRequest{}
request.SetName(nil)

response, err := client.Crypto.GetQuote(ctx, request, ...)
```

## Available Methods

### `Account`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`GetUsage()`](https://tradewatch.io/docs/api-reference/account/usage-statistics) | - | Usage statistics | Get the usage statistics of your API account |

### `Currencies`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`Convert(from, to)`](https://tradewatch.io/docs/api-reference/currencies/conversion) | from, to | Conversion | Convert one symbol to another |
| [`GetInsights()`](https://tradewatch.io/docs/api-reference/currencies/get-insights) | - | Get Insights | Get recent currencies insights. |
| [`GetQuote(symbol)`](https://tradewatch.io/docs/api-reference/currencies/last-quote) | symbol | Last Quote | Get the last quote tick for the provided symbol. |
| [`GetQuotes(symbols)`](https://tradewatch.io/docs/api-reference/currencies/last-quotes) | symbols | Last Quotes | Get the last quote tick for the provided symbols. |
| [`GetSymbols()`](https://tradewatch.io/docs/api-reference/currencies/available-symbols) | - | Available Symbols | Get list of available symbols |

### `Crypto`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`Convert(from, to)`](https://tradewatch.io/docs/api-reference/crypto/conversion) | from, to | Conversion | Convert one symbol to another |
| [`GetExchanges()`](https://tradewatch.io/docs/api-reference/crypto/available-exchanges) | - | Available Exchanges | Get list of available cryptocurrency exchanges |
| [`GetInsights()`](https://tradewatch.io/docs/api-reference/crypto/get-insights) | - | Get Insights | Get recent crypto insights. |
| [`GetQuote(symbol)`](https://tradewatch.io/docs/api-reference/crypto/last-quote) | symbol | Last Quote | Get the last quote tick for the provided symbol. |
| [`GetQuotes(symbols)`](https://tradewatch.io/docs/api-reference/crypto/last-quotes) | symbols | Last Quotes | Get the last quote tick for the provided symbols. |
| [`GetSymbols()`](https://tradewatch.io/docs/api-reference/crypto/available-symbols) | - | Available Symbols | Get list of available symbols |

### `Indices`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`GetInsights()`](https://tradewatch.io/docs/api-reference/indices/get-insights) | - | Get Insights | Get recent indices insights. |
| [`GetQuote(symbol)`](https://tradewatch.io/docs/api-reference/indices/last-quote) | symbol | Last Quote | Get the last quote tick for the provided symbol. |
| [`GetQuotes(symbols)`](https://tradewatch.io/docs/api-reference/indices/last-quotes) | symbols | Last Quotes | Get the last quote tick for the provided symbols. |
| [`GetSymbols()`](https://tradewatch.io/docs/api-reference/indices/available-symbols) | - | Available Symbols | Get list of available symbols |

### `Stocks`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`GetHistoricalOhlc(symbol, resolution, start, end)`](https://tradewatch.io/docs/api-reference/stocks/get-historical-ohlc) | symbol, resolution, start, end | Get Historical Ohlc | Get historical OHLC candles for a symbol in a selected resolution and time range. |
| [`GetHistoricalTicks(symbol, start, end)`](https://tradewatch.io/docs/api-reference/stocks/get-historical-ticks) | symbol, start, end | Get Historical Ticks | Get raw historical ticks for a symbol in a selected time range using cursor pagination. |
| [`GetInsights()`](https://tradewatch.io/docs/api-reference/stocks/get-insights) | - | Get Insights | Get recent stocks insights. |
| [`GetMarketHolidays(start, end)`](https://tradewatch.io/docs/api-reference/stocks/get-market-holidays) | start, end | Get Market Holidays | Get market holidays. It takes half-days into account. |
| [`GetMarketStatus()`](https://tradewatch.io/docs/api-reference/stocks/get-market-status) | - | Get Market Status | Get the current status (open or closed) of a market. It takes holidays and half-days into account but does not factor in circuit breakers or halts. |
| [`GetMarkets()`](https://tradewatch.io/docs/api-reference/stocks/get-markets) | - | Get Markets | Get details about the markets available in this API. |
| [`GetQuote(symbol)`](https://tradewatch.io/docs/api-reference/stocks/last-quote) | symbol | Last Quote | Get the last quote tick for the provided symbol. |
| [`GetQuotes(symbols)`](https://tradewatch.io/docs/api-reference/stocks/last-quotes) | symbols | Last Quotes | Get the last quote tick for the provided symbols. |
| [`GetStockData(symbol)`](https://tradewatch.io/docs/api-reference/stocks/get-stock-data) | symbol | Get Stock Data | Get Stock Data |
| [`GetSymbols()`](https://tradewatch.io/docs/api-reference/stocks/available-symbols) | - | Available Symbols | Get list of available symbols |
| [`GetTradingHours(start, end)`](https://tradewatch.io/docs/api-reference/stocks/get-trading-hours) | start, end | Get Trading Hours | Get trading hours. It takes half-days into account. |
| [`StockGetCountries()`](https://tradewatch.io/docs/api-reference/stocks/available-countries) | - | Available Countries | Get list of available countries |

### `Commodities`

| Method | Required Params | Summary | Description |
| --- | --- | --- | --- |
| [`GetInsights()`](https://tradewatch.io/docs/api-reference/commodities/get-insights) | - | Get Insights | Get recent commodities insights. |
| [`GetQuote(symbol)`](https://tradewatch.io/docs/api-reference/commodities/last-quote) | symbol | Last Quote | Get the last quote tick for the provided symbol. |
| [`GetQuotes(symbols)`](https://tradewatch.io/docs/api-reference/commodities/last-quotes) | symbols | Last Quotes | Get the last quote tick for the provided symbols. |
| [`GetSymbols()`](https://tradewatch.io/docs/api-reference/commodities/available-symbols) | - | Available Symbols | Get list of available symbols |
| [`GetTypes()`](https://tradewatch.io/docs/api-reference/commodities/available-types) | - | Available Types | Get list of available commodity types |
