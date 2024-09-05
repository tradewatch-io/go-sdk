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

// checks if the CommodityTypeObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommodityTypeObj{}

// CommodityTypeObj struct for CommodityTypeObj
type CommodityTypeObj struct {
	// Commodity type name
	Symbol string `json:"symbol"`
}

type _CommodityTypeObj CommodityTypeObj

// NewCommodityTypeObj instantiates a new CommodityTypeObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommodityTypeObj(symbol string) *CommodityTypeObj {
	this := CommodityTypeObj{}
	this.Symbol = symbol
	return &this
}

// NewCommodityTypeObjWithDefaults instantiates a new CommodityTypeObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommodityTypeObjWithDefaults() *CommodityTypeObj {
	this := CommodityTypeObj{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *CommodityTypeObj) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CommodityTypeObj) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CommodityTypeObj) SetSymbol(v string) {
	o.Symbol = v
}

func (o CommodityTypeObj) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommodityTypeObj) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	return toSerialize, nil
}

func (o *CommodityTypeObj) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
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

	varCommodityTypeObj := _CommodityTypeObj{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommodityTypeObj)

	if err != nil {
		return err
	}

	*o = CommodityTypeObj(varCommodityTypeObj)

	return err
}

type NullableCommodityTypeObj struct {
	value *CommodityTypeObj
	isSet bool
}

func (v NullableCommodityTypeObj) Get() *CommodityTypeObj {
	return v.value
}

func (v *NullableCommodityTypeObj) Set(val *CommodityTypeObj) {
	v.value = val
	v.isSet = true
}

func (v NullableCommodityTypeObj) IsSet() bool {
	return v.isSet
}

func (v *NullableCommodityTypeObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommodityTypeObj(val *CommodityTypeObj) *NullableCommodityTypeObj {
	return &NullableCommodityTypeObj{value: val, isSet: true}
}

func (v NullableCommodityTypeObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommodityTypeObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

