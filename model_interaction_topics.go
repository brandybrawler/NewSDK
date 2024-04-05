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

// checks if the InteractionTopics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionTopics{}

// InteractionTopics struct for InteractionTopics
type InteractionTopics struct {
	// The chat ID UUID for an instance of a chat.
	InteractionTopicsId int32 `json:"interaction_topics_id"`
	// Name of the topic
	Name string `json:"name"`
	// Description of the topic
	Description string `json:"description"`
	// The topic category
	Category int32 `json:"category"`
}

type _InteractionTopics InteractionTopics

// NewInteractionTopics instantiates a new InteractionTopics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionTopics(interactionTopicsId int32, name string, description string, category int32) *InteractionTopics {
	this := InteractionTopics{}
	this.InteractionTopicsId = interactionTopicsId
	this.Name = name
	this.Description = description
	this.Category = category
	return &this
}

// NewInteractionTopicsWithDefaults instantiates a new InteractionTopics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionTopicsWithDefaults() *InteractionTopics {
	this := InteractionTopics{}
	return &this
}

// GetInteractionTopicsId returns the InteractionTopicsId field value
func (o *InteractionTopics) GetInteractionTopicsId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InteractionTopicsId
}

// GetInteractionTopicsIdOk returns a tuple with the InteractionTopicsId field value
// and a boolean to check if the value has been set.
func (o *InteractionTopics) GetInteractionTopicsIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionTopicsId, true
}

// SetInteractionTopicsId sets field value
func (o *InteractionTopics) SetInteractionTopicsId(v int32) {
	o.InteractionTopicsId = v
}

// GetName returns the Name field value
func (o *InteractionTopics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InteractionTopics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InteractionTopics) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *InteractionTopics) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InteractionTopics) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InteractionTopics) SetDescription(v string) {
	o.Description = v
}

// GetCategory returns the Category field value
func (o *InteractionTopics) GetCategory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *InteractionTopics) GetCategoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *InteractionTopics) SetCategory(v int32) {
	o.Category = v
}

func (o InteractionTopics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionTopics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interaction_topics_id"] = o.InteractionTopicsId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["category"] = o.Category
	return toSerialize, nil
}

func (o *InteractionTopics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interaction_topics_id",
		"name",
		"description",
		"category",
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

	varInteractionTopics := _InteractionTopics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInteractionTopics)

	if err != nil {
		return err
	}

	*o = InteractionTopics(varInteractionTopics)

	return err
}

type NullableInteractionTopics struct {
	value *InteractionTopics
	isSet bool
}

func (v NullableInteractionTopics) Get() *InteractionTopics {
	return v.value
}

func (v *NullableInteractionTopics) Set(val *InteractionTopics) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionTopics) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionTopics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionTopics(val *InteractionTopics) *NullableInteractionTopics {
	return &NullableInteractionTopics{value: val, isSet: true}
}

func (v NullableInteractionTopics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionTopics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

