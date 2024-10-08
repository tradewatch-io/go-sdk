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

// checks if the SymbolsOutFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SymbolsOutFull{}

// SymbolsOutFull struct for SymbolsOutFull
type SymbolsOutFull struct {
	Symbol string `json:"symbol"`
	Name string `json:"name"`
}

type _SymbolsOutFull SymbolsOutFull

// NewSymbolsOutFull instantiates a new SymbolsOutFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolsOutFull(symbol string, name string) *SymbolsOutFull {
	this := SymbolsOutFull{}
	this.Symbol = symbol
	this.Name = name
	return &this
}

// NewSymbolsOutFullWithDefaults instantiates a new SymbolsOutFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolsOutFullWithDefaults() *SymbolsOutFull {
	this := SymbolsOutFull{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SymbolsOutFull) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SymbolsOutFull) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SymbolsOutFull) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *SymbolsOutFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SymbolsOutFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SymbolsOutFull) SetName(v string) {
	o.Name = v
}

func (o SymbolsOutFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolsOutFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *SymbolsOutFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
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

	varSymbolsOutFull := _SymbolsOutFull{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSymbolsOutFull)

	if err != nil {
		return err
	}

	*o = SymbolsOutFull(varSymbolsOutFull)

	return err
}

type NullableSymbolsOutFull struct {
	value *SymbolsOutFull
	isSet bool
}

func (v NullableSymbolsOutFull) Get() *SymbolsOutFull {
	return v.value
}

func (v *NullableSymbolsOutFull) Set(val *SymbolsOutFull) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolsOutFull) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolsOutFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolsOutFull(val *SymbolsOutFull) *NullableSymbolsOutFull {
	return &NullableSymbolsOutFull{value: val, isSet: true}
}

func (v NullableSymbolsOutFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolsOutFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


