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

// checks if the TenantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantInfo{}

// TenantInfo struct for TenantInfo
type TenantInfo struct {
	// The tenant chats ID.
	TenantId int32 `json:"tenant_id"`
	// name of the company
	TenantName string `json:"tenant_name"`
}

type _TenantInfo TenantInfo

// NewTenantInfo instantiates a new TenantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantInfo(tenantId int32, tenantName string) *TenantInfo {
	this := TenantInfo{}
	this.TenantId = tenantId
	this.TenantName = tenantName
	return &this
}

// NewTenantInfoWithDefaults instantiates a new TenantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantInfoWithDefaults() *TenantInfo {
	this := TenantInfo{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *TenantInfo) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *TenantInfo) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *TenantInfo) SetTenantId(v int32) {
	o.TenantId = v
}

// GetTenantName returns the TenantName field value
func (o *TenantInfo) GetTenantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value
// and a boolean to check if the value has been set.
func (o *TenantInfo) GetTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantName, true
}

// SetTenantName sets field value
func (o *TenantInfo) SetTenantName(v string) {
	o.TenantName = v
}

func (o TenantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["tenant_name"] = o.TenantName
	return toSerialize, nil
}

func (o *TenantInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenant_id",
		"tenant_name",
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

	varTenantInfo := _TenantInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenantInfo)

	if err != nil {
		return err
	}

	*o = TenantInfo(varTenantInfo)

	return err
}

type NullableTenantInfo struct {
	value *TenantInfo
	isSet bool
}

func (v NullableTenantInfo) Get() *TenantInfo {
	return v.value
}

func (v *NullableTenantInfo) Set(val *TenantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantInfo(val *TenantInfo) *NullableTenantInfo {
	return &NullableTenantInfo{value: val, isSet: true}
}

func (v NullableTenantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


