/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AdminGender - Enter your gender.  * `Male` - Male * `Female` - Female * `Other` - Other
type AdminGender struct {
	BlankEnum *BlankEnum
	GenderEnum *GenderEnum
}

// BlankEnumAsAdminGender is a convenience function that returns BlankEnum wrapped in AdminGender
func BlankEnumAsAdminGender(v *BlankEnum) AdminGender {
	return AdminGender{
		BlankEnum: v,
	}
}

// GenderEnumAsAdminGender is a convenience function that returns GenderEnum wrapped in AdminGender
func GenderEnumAsAdminGender(v *GenderEnum) AdminGender {
	return AdminGender{
		GenderEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AdminGender) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			match++
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into GenderEnum
	err = newStrictDecoder(data).Decode(&dst.GenderEnum)
	if err == nil {
		jsonGenderEnum, _ := json.Marshal(dst.GenderEnum)
		if string(jsonGenderEnum) == "{}" { // empty struct
			dst.GenderEnum = nil
		} else {
			match++
		}
	} else {
		dst.GenderEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.GenderEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AdminGender)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AdminGender)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AdminGender) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.GenderEnum != nil {
		return json.Marshal(&src.GenderEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AdminGender) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.GenderEnum != nil {
		return obj.GenderEnum
	}

	// all schemas are nil
	return nil
}

type NullableAdminGender struct {
	value *AdminGender
	isSet bool
}

func (v NullableAdminGender) Get() *AdminGender {
	return v.value
}

func (v *NullableAdminGender) Set(val *AdminGender) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminGender) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminGender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminGender(val *AdminGender) *NullableAdminGender {
	return &NullableAdminGender{value: val, isSet: true}
}

func (v NullableAdminGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminGender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


