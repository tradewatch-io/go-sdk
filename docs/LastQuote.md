# LastQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol name | 
**Ask** | **float32** | The ask price. | 
**Bid** | **float32** | The bid price. | 
**Mid** | **float32** | The mid price. | 
**Timestamp** | **int32** |  | 

## Methods

### NewLastQuote

`func NewLastQuote(symbol string, ask float32, bid float32, mid float32, timestamp int32, ) *LastQuote`

NewLastQuote instantiates a new LastQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastQuoteWithDefaults

`func NewLastQuoteWithDefaults() *LastQuote`

NewLastQuoteWithDefaults instantiates a new LastQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LastQuote) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LastQuote) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LastQuote) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAsk

`func (o *LastQuote) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *LastQuote) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *LastQuote) SetAsk(v float32)`

SetAsk sets Ask field to given value.


### GetBid

`func (o *LastQuote) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *LastQuote) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *LastQuote) SetBid(v float32)`

SetBid sets Bid field to given value.


### GetMid

`func (o *LastQuote) GetMid() float32`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *LastQuote) GetMidOk() (*float32, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *LastQuote) SetMid(v float32)`

SetMid sets Mid field to given value.


### GetTimestamp

`func (o *LastQuote) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LastQuote) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LastQuote) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


