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

// checks if the JoinCommunity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoinCommunity{}

// JoinCommunity struct for JoinCommunity
type JoinCommunity struct {
	CommunityId int32 `json:"community_id"`
	ClientId int32 `json:"client_id"`
}

type _JoinCommunity JoinCommunity

// NewJoinCommunity instantiates a new JoinCommunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinCommunity(communityId int32, clientId int32) *JoinCommunity {
	this := JoinCommunity{}
	this.CommunityId = communityId
	this.ClientId = clientId
	return &this
}

// NewJoinCommunityWithDefaults instantiates a new JoinCommunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinCommunityWithDefaults() *JoinCommunity {
	this := JoinCommunity{}
	return &this
}

// GetCommunityId returns the CommunityId field value
func (o *JoinCommunity) GetCommunityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommunityId
}

// GetCommunityIdOk returns a tuple with the CommunityId field value
// and a boolean to check if the value has been set.
func (o *JoinCommunity) GetCommunityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunityId, true
}

// SetCommunityId sets field value
func (o *JoinCommunity) SetCommunityId(v int32) {
	o.CommunityId = v
}

// GetClientId returns the ClientId field value
func (o *JoinCommunity) GetClientId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *JoinCommunity) GetClientIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *JoinCommunity) SetClientId(v int32) {
	o.ClientId = v
}

func (o JoinCommunity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoinCommunity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["community_id"] = o.CommunityId
	toSerialize["client_id"] = o.ClientId
	return toSerialize, nil
}

func (o *JoinCommunity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"community_id",
		"client_id",
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

	varJoinCommunity := _JoinCommunity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJoinCommunity)

	if err != nil {
		return err
	}

	*o = JoinCommunity(varJoinCommunity)

	return err
}

type NullableJoinCommunity struct {
	value *JoinCommunity
	isSet bool
}

func (v NullableJoinCommunity) Get() *JoinCommunity {
	return v.value
}

func (v *NullableJoinCommunity) Set(val *JoinCommunity) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinCommunity) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinCommunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinCommunity(val *JoinCommunity) *NullableJoinCommunity {
	return &NullableJoinCommunity{value: val, isSet: true}
}

func (v NullableJoinCommunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinCommunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


