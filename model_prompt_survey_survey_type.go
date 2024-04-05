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

// PromptSurveySurveyType - The survey type  * `open_ended` - open_ended * `close_ended` - close_ended
type PromptSurveySurveyType struct {
	BlankEnum *BlankEnum
	SurveyTypeEnum *SurveyTypeEnum
}

// BlankEnumAsPromptSurveySurveyType is a convenience function that returns BlankEnum wrapped in PromptSurveySurveyType
func BlankEnumAsPromptSurveySurveyType(v *BlankEnum) PromptSurveySurveyType {
	return PromptSurveySurveyType{
		BlankEnum: v,
	}
}

// SurveyTypeEnumAsPromptSurveySurveyType is a convenience function that returns SurveyTypeEnum wrapped in PromptSurveySurveyType
func SurveyTypeEnumAsPromptSurveySurveyType(v *SurveyTypeEnum) PromptSurveySurveyType {
	return PromptSurveySurveyType{
		SurveyTypeEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PromptSurveySurveyType) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into SurveyTypeEnum
	err = newStrictDecoder(data).Decode(&dst.SurveyTypeEnum)
	if err == nil {
		jsonSurveyTypeEnum, _ := json.Marshal(dst.SurveyTypeEnum)
		if string(jsonSurveyTypeEnum) == "{}" { // empty struct
			dst.SurveyTypeEnum = nil
		} else {
			match++
		}
	} else {
		dst.SurveyTypeEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.SurveyTypeEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PromptSurveySurveyType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PromptSurveySurveyType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PromptSurveySurveyType) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.SurveyTypeEnum != nil {
		return json.Marshal(&src.SurveyTypeEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PromptSurveySurveyType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.SurveyTypeEnum != nil {
		return obj.SurveyTypeEnum
	}

	// all schemas are nil
	return nil
}

type NullablePromptSurveySurveyType struct {
	value *PromptSurveySurveyType
	isSet bool
}

func (v NullablePromptSurveySurveyType) Get() *PromptSurveySurveyType {
	return v.value
}

func (v *NullablePromptSurveySurveyType) Set(val *PromptSurveySurveyType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSurveySurveyType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSurveySurveyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSurveySurveyType(val *PromptSurveySurveyType) *NullablePromptSurveySurveyType {
	return &NullablePromptSurveySurveyType{value: val, isSet: true}
}

func (v NullablePromptSurveySurveyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSurveySurveyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


