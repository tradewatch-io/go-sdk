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

// checks if the CryptoConversion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoConversion{}

// CryptoConversion struct for CryptoConversion
type CryptoConversion struct {
	Info ConversionInfo `json:"info"`
	Query CryptoConversionQuery `json:"query"`
	Result float32 `json:"result"`
}

type _CryptoConversion CryptoConversion

// NewCryptoConversion instantiates a new CryptoConversion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoConversion(info ConversionInfo, query CryptoConversionQuery, result float32) *CryptoConversion {
	this := CryptoConversion{}
	this.Info = info
	this.Query = query
	this.Result = result
	return &this
}

// NewCryptoConversionWithDefaults instantiates a new CryptoConversion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoConversionWithDefaults() *CryptoConversion {
	this := CryptoConversion{}
	return &this
}

// GetInfo returns the Info field value
func (o *CryptoConversion) GetInfo() ConversionInfo {
	if o == nil {
		var ret ConversionInfo
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *CryptoConversion) GetInfoOk() (*ConversionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *CryptoConversion) SetInfo(v ConversionInfo) {
	o.Info = v
}

// GetQuery returns the Query field value
func (o *CryptoConversion) GetQuery() CryptoConversionQuery {
	if o == nil {
		var ret CryptoConversionQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CryptoConversion) GetQueryOk() (*CryptoConversionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CryptoConversion) SetQuery(v CryptoConversionQuery) {
	o.Query = v
}

// GetResult returns the Result field value
func (o *CryptoConversion) GetResult() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CryptoConversion) GetResultOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CryptoConversion) SetResult(v float32) {
	o.Result = v
}

func (o CryptoConversion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoConversion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["info"] = o.Info
	toSerialize["query"] = o.Query
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CryptoConversion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"info",
		"query",
		"result",
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

	varCryptoConversion := _CryptoConversion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCryptoConversion)

	if err != nil {
		return err
	}

	*o = CryptoConversion(varCryptoConversion)

	return err
}

type NullableCryptoConversion struct {
	value *CryptoConversion
	isSet bool
}

func (v NullableCryptoConversion) Get() *CryptoConversion {
	return v.value
}

func (v *NullableCryptoConversion) Set(val *CryptoConversion) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoConversion) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoConversion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoConversion(val *CryptoConversion) *NullableCryptoConversion {
	return &NullableCryptoConversion{value: val, isSet: true}
}

func (v NullableCryptoConversion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoConversion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

