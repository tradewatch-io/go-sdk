/*
tradewatch.io

Financial market data for Developers

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tradewatch

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LastQuote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastQuote{}

// LastQuote struct for LastQuote
type LastQuote struct {
	// Symbol name
	Symbol string `json:"symbol"`
	// The ask price.
	Ask float32 `json:"ask"`
	// The bid price.
	Bid float32 `json:"bid"`
	// The mid price.
	Mid float32 `json:"mid"`
	Timestamp int32 `json:"timestamp"`
}

type _LastQuote LastQuote

// NewLastQuote instantiates a new LastQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastQuote(symbol string, ask float32, bid float32, mid float32, timestamp int32) *LastQuote {
	this := LastQuote{}
	this.Symbol = symbol
	this.Ask = ask
	this.Bid = bid
	this.Mid = mid
	this.Timestamp = timestamp
	return &this
}

// NewLastQuoteWithDefaults instantiates a new LastQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastQuoteWithDefaults() *LastQuote {
	this := LastQuote{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *LastQuote) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *LastQuote) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *LastQuote) SetSymbol(v string) {
	o.Symbol = v
}

// GetAsk returns the Ask field value
func (o *LastQuote) GetAsk() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ask
}

// GetAskOk returns a tuple with the Ask field value
// and a boolean to check if the value has been set.
func (o *LastQuote) GetAskOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ask, true
}

// SetAsk sets field value
func (o *LastQuote) SetAsk(v float32) {
	o.Ask = v
}

// GetBid returns the Bid field value
func (o *LastQuote) GetBid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *LastQuote) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *LastQuote) SetBid(v float32) {
	o.Bid = v
}

// GetMid returns the Mid field value
func (o *LastQuote) GetMid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *LastQuote) GetMidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *LastQuote) SetMid(v float32) {
	o.Mid = v
}

// GetTimestamp returns the Timestamp field value
func (o *LastQuote) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LastQuote) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LastQuote) SetTimestamp(v int32) {
	o.Timestamp = v
}

func (o LastQuote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastQuote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["ask"] = o.Ask
	toSerialize["bid"] = o.Bid
	toSerialize["mid"] = o.Mid
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *LastQuote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"ask",
		"bid",
		"mid",
		"timestamp",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLastQuote := _LastQuote{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLastQuote)

	if err != nil {
		return err
	}

	*o = LastQuote(varLastQuote)

	return err
}

type NullableLastQuote struct {
	value *LastQuote
	isSet bool
}

func (v NullableLastQuote) Get() *LastQuote {
	return v.value
}

func (v *NullableLastQuote) Set(val *LastQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableLastQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableLastQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastQuote(val *LastQuote) *NullableLastQuote {
	return &NullableLastQuote{value: val, isSet: true}
}

func (v NullableLastQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


