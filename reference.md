# Reference
## account
<details><summary><code>client.Account.GetUsage() -> sdk.AccountUsageStatisticsInput</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the usage statistics of your API account
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.AccountGetUsageRequest{}
client.Account.GetUsage(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — The number of data points to return (max 168).
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*sdk.AccountUsageStatisticsInterval` — The time interval for the usage statistics.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## currencies
<details><summary><code>client.Currencies.Convert(From, To) -> *sdk.Conversion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert one symbol to another
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrenciesConvertRequest{
        From: "EUR",
        To: "USD",
        Amount: sdk.Float64(
            1000,
        ),
        Precision: sdk.Int(
            2,
        ),
    }
client.Currencies.Convert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `string` — The symbol you want to convert from.
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — The symbol you want to convert to.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — The amount to be converted.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Currencies.GetQuotes() -> *sdk.LastQuotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbols.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrenciesGetQuotesRequest{
        Symbols: "symbols",
    }
client.Currencies.GetQuotes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `string` — Comma separated list of symbols.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Currencies.GetQuote() -> *sdk.LastQuote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrenciesGetQuoteRequest{
        Symbol: "EUR-USD",
        Precision: sdk.Int(
            4,
        ),
    }
client.Currencies.GetQuote(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` — The symbol to get the quote for.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Currencies.GetSymbols() -> *sdk.CursorPageTCustomizedSymbolsOutFull</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available symbols
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrenciesGetSymbolsRequest{}
client.Currencies.GetSymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*string` — The mode of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Type of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Currencies.GetInsights() -> *sdk.CursorPageTCustomizedNewsOut</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent currencies insights.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrenciesGetInsightsRequest{}
client.Currencies.GetInsights(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## crypto
<details><summary><code>client.Crypto.Convert(From, To) -> *sdk.CryptoConversion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert one symbol to another
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CryptoConvertRequest{
        From: "from",
        To: "to",
    }
client.Crypto.Convert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `string` — The symbol you want to convert from.
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — The symbol you want to convert to.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — The amount to be converted.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crypto.GetExchanges() -> *sdk.CryptoExchangesList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available cryptocurrency exchanges
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crypto.GetExchanges(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crypto.GetQuotes() -> *sdk.LastQuotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbols.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CryptoGetQuotesRequest{
        Symbols: "BTC-USD,ETH-USD",
        Precision: sdk.Int(
            2,
        ),
    }
client.Crypto.GetQuotes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `string` — Comma separated list of symbols.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crypto.GetQuote() -> *sdk.LastQuote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
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
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` — The symbol to get the quote for.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crypto.GetSymbols() -> *sdk.CursorPageTCustomizedSymbolsOutFull</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available symbols
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CryptoGetSymbolsRequest{}
client.Crypto.GetSymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*string` — The mode of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Type of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crypto.GetInsights() -> *sdk.CursorPageTCustomizedNewsOut</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent crypto insights.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CryptoGetInsightsRequest{}
client.Crypto.GetInsights(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## indices
<details><summary><code>client.Indices.GetQuotes() -> *sdk.LastQuotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbols.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IndicesGetQuotesRequest{
        Symbols: "symbols",
    }
client.Indices.GetQuotes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `string` — Comma separated list of symbols.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Indices.GetQuote() -> *sdk.LastQuote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IndicesGetQuoteRequest{
        Symbol: "DJI",
        Precision: sdk.Int(
            2,
        ),
    }
client.Indices.GetQuote(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` — The symbol to get the quote for.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Indices.GetSymbols() -> *sdk.CursorPageTCustomizedSymbolsOutFull</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available symbols
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IndicesGetSymbolsRequest{}
client.Indices.GetSymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*string` — The mode of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Type of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Indices.GetInsights() -> *sdk.CursorPageTCustomizedNewsOut</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent indices insights.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IndicesGetInsightsRequest{}
client.Indices.GetInsights(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## stocks
<details><summary><code>client.Stocks.GetQuotes() -> *sdk.LastQuotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbols.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetQuotesRequest{
        Symbols: "symbols",
    }
client.Stocks.GetQuotes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `string` — Comma separated list of symbols.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetQuote() -> *sdk.LastQuote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetQuoteRequest{
        Symbol: "symbol",
    }
client.Stocks.GetQuote(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` — The symbol to get the quote for.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetSymbols() -> *sdk.CursorPageTCustomizedSymbolsOutFullStocks</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available symbols
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetSymbolsRequest{}
client.Stocks.GetSymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*string` — The mode of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Type of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetInsights() -> *sdk.CursorPageTCustomizedNewsOut</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent stocks insights.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetInsightsRequest{}
client.Stocks.GetInsights(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.StockGetCountries() -> *sdk.CountriesList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available countries
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Stocks.StockGetCountries(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetMarkets() -> []*sdk.MarketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details about the markets available in this API.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetMarketsRequest{
        Mic: sdk.String(
            "XNYS",
        ),
    }
client.Stocks.GetMarkets(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mic:** `*string` — Optional list of comma separated MIC codes for which market to show data for. All market will be included if MIC code is not specified.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetMarketStatus() -> []*sdk.MarketStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the current status (open or closed) of a market. It takes holidays and half-days into account but does not factor in circuit breakers or halts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetMarketStatusRequest{
        Mic: sdk.String(
            "XNYS",
        ),
    }
client.Stocks.GetMarketStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mic:** `*string` — Optional list of comma separated MIC codes for which market to show data for. All market will be included if MIC code is not specified.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetTradingHours() -> []*sdk.TradingHoursResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get trading hours. It takes half-days into account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetTradingHoursRequest{
        Mic: sdk.String(
            "XNAS",
        ),
        Start: sdk.MustParseDate(
            "2025-01-01",
        ),
        End: sdk.MustParseDate(
            "2025-01-31",
        ),
    }
client.Stocks.GetTradingHours(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mic:** `*string` — Optional list of comma separated MIC codes for which market to show data for. All market will be included if MIC code is not specified.
    
</dd>
</dl>

<dl>
<dd>

**start:** `time.Time` — Show holidays starting at this date.
    
</dd>
</dl>

<dl>
<dd>

**end:** `time.Time` — Show holidays until this date.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetMarketHolidays() -> []*sdk.MarketHolidayResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get market holidays. It takes half-days into account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetMarketHolidaysRequest{
        Mic: sdk.String(
            "XNYS",
        ),
        Start: sdk.MustParseDate(
            "2026-02-17",
        ),
        End: sdk.MustParseDate(
            "2026-02-24",
        ),
    }
client.Stocks.GetMarketHolidays(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**mic:** `*string` — Specify comma separated list of MIC codes for which market to show data for.
    
</dd>
</dl>

<dl>
<dd>

**start:** `time.Time` — Show holidays starting at this date.
    
</dd>
</dl>

<dl>
<dd>

**end:** `time.Time` — Show holidays until this date.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetStockData(Symbol) -> *sdk.StockDataFlatResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetStockDataRequest{
        Symbol: "AAPL",
    }
client.Stocks.GetStockData(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetHistoricalOhlc(Symbol) -> *sdk.HistoricalOhlcResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get historical OHLC candles for a symbol in a selected resolution and time range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetHistoricalOhlcRequest{
        Symbol: "symbol",
        Resolution: sdk.HistoricalDataResolutionValue5,
        Start: 1,
        End: 1,
    }
client.Stocks.GetHistoricalOhlc(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*sdk.HistoricalDataResolution` — Resolution in seconds.
    
</dd>
</dl>

<dl>
<dd>

**start:** `int` — Unix timestamp (inclusive).
    
</dd>
</dl>

<dl>
<dd>

**end:** `int` — Unix timestamp (exclusive).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stocks.GetHistoricalTicks(Symbol) -> *sdk.CursorPageTCustomizedHistoricalRawTick</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get raw historical ticks for a symbol in a selected time range using cursor pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.StocksGetHistoricalTicksRequest{
        Symbol: "symbol",
        Start: 1,
        End: 1,
    }
client.Stocks.GetHistoricalTicks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `int` — Unix timestamp (inclusive).
    
</dd>
</dl>

<dl>
<dd>

**end:** `int` — Unix timestamp (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of ticks per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## commodities
<details><summary><code>client.Commodities.GetQuotes() -> *sdk.LastQuotes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbols.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommoditiesGetQuotesRequest{
        Symbols: "symbols",
    }
client.Commodities.GetQuotes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `string` — Comma separated list of symbols.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Commodities.GetQuote() -> *sdk.LastQuote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the last quote tick for the provided symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommoditiesGetQuoteRequest{
        Symbol: "GOLD",
        Precision: sdk.Int(
            2,
        ),
    }
client.Commodities.GetQuote(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbol:** `string` — The symbol to get the quote for.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*int` — The decimal precision of the result.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Commodities.GetSymbols() -> *sdk.CursorPageTCustomizedSymbolsOutFull</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available symbols
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommoditiesGetSymbolsRequest{}
client.Commodities.GetSymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*string` — The mode of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Type of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country of the instrument.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Commodities.GetInsights() -> *sdk.CursorPageTCustomizedNewsOut</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get recent commodities insights.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommoditiesGetInsightsRequest{}
client.Commodities.GetInsights(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for the next page
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — The number of items per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Commodities.GetTypes() -> *sdk.CommodityTypesList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of available commodity types
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Commodities.GetTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
