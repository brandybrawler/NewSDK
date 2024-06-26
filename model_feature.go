/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Feature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Feature{}

// Feature struct for Feature
type Feature struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ExtraFeatures interface{} `json:"extra_features,omitempty"`
}

type _Feature Feature

// NewFeature instantiates a new Feature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeature(name string, description string) *Feature {
	this := Feature{}
	this.Name = name
	this.Description = description
	return &this
}

// NewFeatureWithDefaults instantiates a new Feature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureWithDefaults() *Feature {
	this := Feature{}
	return &this
}

// GetName returns the Name field value
func (o *Feature) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Feature) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Feature) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Feature) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Feature) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Feature) SetDescription(v string) {
	o.Description = v
}

// GetExtraFeatures returns the ExtraFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Feature) GetExtraFeatures() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraFeatures
}

// GetExtraFeaturesOk returns a tuple with the ExtraFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Feature) GetExtraFeaturesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraFeatures) {
		return nil, false
	}
	return &o.ExtraFeatures, true
}

// HasExtraFeatures returns a boolean if a field has been set.
func (o *Feature) HasExtraFeatures() bool {
	if o != nil && !IsNil(o.ExtraFeatures) {
		return true
	}

	return false
}

// SetExtraFeatures gets a reference to the given interface{} and assigns it to the ExtraFeatures field.
func (o *Feature) SetExtraFeatures(v interface{}) {
	o.ExtraFeatures = v
}

func (o Feature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Feature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.ExtraFeatures != nil {
		toSerialize["extra_features"] = o.ExtraFeatures
	}
	return toSerialize, nil
}

func (o *Feature) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
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

	varFeature := _Feature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeature)

	if err != nil {
		return err
	}

	*o = Feature(varFeature)

	return err
}

type NullableFeature struct {
	value *Feature
	isSet bool
}

func (v NullableFeature) Get() *Feature {
	return v.value
}

func (v *NullableFeature) Set(val *Feature) {
	v.value = val
	v.isSet = true
}

func (v NullableFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeature(val *Feature) *NullableFeature {
	return &NullableFeature{value: val, isSet: true}
}

func (v NullableFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


