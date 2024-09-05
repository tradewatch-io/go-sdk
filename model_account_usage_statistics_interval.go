/*
tradewatch.io

Financial market data for Developers

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tradewatch

import (
	"encoding/json"
	"fmt"
)

// AccountUsageStatisticsInterval Data interval
type AccountUsageStatisticsInterval string

// List of AccountUsageStatisticsInterval
const (
	_1H AccountUsageStatisticsInterval = "1h"
	_1D AccountUsageStatisticsInterval = "1d"
)

// All allowed values of AccountUsageStatisticsInterval enum
var AllowedAccountUsageStatisticsIntervalEnumValues = []AccountUsageStatisticsInterval{
	"1h",
	"1d",
}

func (v *AccountUsageStatisticsInterval) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountUsageStatisticsInterval(value)
	for _, existing := range AllowedAccountUsageStatisticsIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountUsageStatisticsInterval", value)
}

// NewAccountUsageStatisticsIntervalFromValue returns a pointer to a valid AccountUsageStatisticsInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountUsageStatisticsIntervalFromValue(v string) (*AccountUsageStatisticsInterval, error) {
	ev := AccountUsageStatisticsInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountUsageStatisticsInterval: valid values are %v", v, AllowedAccountUsageStatisticsIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountUsageStatisticsInterval) IsValid() bool {
	for _, existing := range AllowedAccountUsageStatisticsIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountUsageStatisticsInterval value
func (v AccountUsageStatisticsInterval) Ptr() *AccountUsageStatisticsInterval {
	return &v
}

type NullableAccountUsageStatisticsInterval struct {
	value *AccountUsageStatisticsInterval
	isSet bool
}

func (v NullableAccountUsageStatisticsInterval) Get() *AccountUsageStatisticsInterval {
	return v.value
}

func (v *NullableAccountUsageStatisticsInterval) Set(val *AccountUsageStatisticsInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUsageStatisticsInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUsageStatisticsInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUsageStatisticsInterval(val *AccountUsageStatisticsInterval) *NullableAccountUsageStatisticsInterval {
	return &NullableAccountUsageStatisticsInterval{value: val, isSet: true}
}

func (v NullableAccountUsageStatisticsInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUsageStatisticsInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

