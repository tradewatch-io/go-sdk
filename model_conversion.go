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

// checks if the Conversion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conversion{}

// Conversion struct for Conversion
type Conversion struct {
	Info ConversionInfo `json:"info"`
	Query ConversionQuery `json:"query"`
	Result float32 `json:"result"`
}

type _Conversion Conversion

// NewConversion instantiates a new Conversion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversion(info ConversionInfo, query ConversionQuery, result float32) *Conversion {
	this := Conversion{}
	this.Info = info
	this.Query = query
	this.Result = result
	return &this
}

// NewConversionWithDefaults instantiates a new Conversion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionWithDefaults() *Conversion {
	this := Conversion{}
	return &this
}

// GetInfo returns the Info field value
func (o *Conversion) GetInfo() ConversionInfo {
	if o == nil {
		var ret ConversionInfo
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *Conversion) GetInfoOk() (*ConversionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *Conversion) SetInfo(v ConversionInfo) {
	o.Info = v
}

// GetQuery returns the Query field value
func (o *Conversion) GetQuery() ConversionQuery {
	if o == nil {
		var ret ConversionQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Conversion) GetQueryOk() (*ConversionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Conversion) SetQuery(v ConversionQuery) {
	o.Query = v
}

// GetResult returns the Result field value
func (o *Conversion) GetResult() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Conversion) GetResultOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Conversion) SetResult(v float32) {
	o.Result = v
}

func (o Conversion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conversion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["info"] = o.Info
	toSerialize["query"] = o.Query
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *Conversion) UnmarshalJSON(data []byte) (err error) {
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

	varConversion := _Conversion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConversion)

	if err != nil {
		return err
	}

	*o = Conversion(varConversion)

	return err
}

type NullableConversion struct {
	value *Conversion
	isSet bool
}

func (v NullableConversion) Get() *Conversion {
	return v.value
}

func (v *NullableConversion) Set(val *Conversion) {
	v.value = val
	v.isSet = true
}

func (v NullableConversion) IsSet() bool {
	return v.isSet
}

func (v *NullableConversion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversion(val *Conversion) *NullableConversion {
	return &NullableConversion{value: val, isSet: true}
}

func (v NullableConversion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


