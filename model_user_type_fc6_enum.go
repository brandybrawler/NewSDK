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

// UserTypeFc6Enum * `Admin` - Admin * `Employee` - Employee * `Client` - Client
type UserTypeFc6Enum string

// List of UserTypeFc6Enum
const (
	ADMIN UserTypeFc6Enum = "Admin"
	EMPLOYEE UserTypeFc6Enum = "Employee"
	CLIENT UserTypeFc6Enum = "Client"
)

// All allowed values of UserTypeFc6Enum enum
var AllowedUserTypeFc6EnumEnumValues = []UserTypeFc6Enum{
	"Admin",
	"Employee",
	"Client",
}

func (v *UserTypeFc6Enum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserTypeFc6Enum(value)
	for _, existing := range AllowedUserTypeFc6EnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserTypeFc6Enum", value)
}

// NewUserTypeFc6EnumFromValue returns a pointer to a valid UserTypeFc6Enum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserTypeFc6EnumFromValue(v string) (*UserTypeFc6Enum, error) {
	ev := UserTypeFc6Enum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserTypeFc6Enum: valid values are %v", v, AllowedUserTypeFc6EnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserTypeFc6Enum) IsValid() bool {
	for _, existing := range AllowedUserTypeFc6EnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTypeFc6Enum value
func (v UserTypeFc6Enum) Ptr() *UserTypeFc6Enum {
	return &v
}

type NullableUserTypeFc6Enum struct {
	value *UserTypeFc6Enum
	isSet bool
}

func (v NullableUserTypeFc6Enum) Get() *UserTypeFc6Enum {
	return v.value
}

func (v *NullableUserTypeFc6Enum) Set(val *UserTypeFc6Enum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypeFc6Enum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypeFc6Enum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypeFc6Enum(val *UserTypeFc6Enum) *NullableUserTypeFc6Enum {
	return &NullableUserTypeFc6Enum{value: val, isSet: true}
}

func (v NullableUserTypeFc6Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypeFc6Enum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

