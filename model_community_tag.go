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

// checks if the CommunityTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityTag{}

// CommunityTag struct for CommunityTag
type CommunityTag struct {
	Name string `json:"name"`
	// Display name of the community
	CommunityId int32 `json:"community_id"`
}

type _CommunityTag CommunityTag

// NewCommunityTag instantiates a new CommunityTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityTag(name string, communityId int32) *CommunityTag {
	this := CommunityTag{}
	this.Name = name
	this.CommunityId = communityId
	return &this
}

// NewCommunityTagWithDefaults instantiates a new CommunityTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityTagWithDefaults() *CommunityTag {
	this := CommunityTag{}
	return &this
}

// GetName returns the Name field value
func (o *CommunityTag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommunityTag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CommunityTag) SetName(v string) {
	o.Name = v
}

// GetCommunityId returns the CommunityId field value
func (o *CommunityTag) GetCommunityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommunityId
}

// GetCommunityIdOk returns a tuple with the CommunityId field value
// and a boolean to check if the value has been set.
func (o *CommunityTag) GetCommunityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunityId, true
}

// SetCommunityId sets field value
func (o *CommunityTag) SetCommunityId(v int32) {
	o.CommunityId = v
}

func (o CommunityTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["community_id"] = o.CommunityId
	return toSerialize, nil
}

func (o *CommunityTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"community_id",
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

	varCommunityTag := _CommunityTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityTag)

	if err != nil {
		return err
	}

	*o = CommunityTag(varCommunityTag)

	return err
}

type NullableCommunityTag struct {
	value *CommunityTag
	isSet bool
}

func (v NullableCommunityTag) Get() *CommunityTag {
	return v.value
}

func (v *NullableCommunityTag) Set(val *CommunityTag) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityTag) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityTag(val *CommunityTag) *NullableCommunityTag {
	return &NullableCommunityTag{value: val, isSet: true}
}

func (v NullableCommunityTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


