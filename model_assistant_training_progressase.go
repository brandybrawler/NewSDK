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

// checks if the AssistantTrainingProgressase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantTrainingProgressase{}

// AssistantTrainingProgressase struct for AssistantTrainingProgressase
type AssistantTrainingProgressase struct {
	// The tenant chats ID.
	AssistantTrainingId int32 `json:"assistant_training_id"`
	Status NullableAssistantTrainingProgressaseStatus `json:"status,omitempty"`
	// Display name of the tenant
	TenantId int32 `json:"tenant_id"`
	// Know the progress of updating an assistant
	Succesful *bool `json:"succesful,omitempty"`
	// Updating
	Message NullableString `json:"message,omitempty"`
}

type _AssistantTrainingProgressase AssistantTrainingProgressase

// NewAssistantTrainingProgressase instantiates a new AssistantTrainingProgressase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantTrainingProgressase(assistantTrainingId int32, tenantId int32) *AssistantTrainingProgressase {
	this := AssistantTrainingProgressase{}
	this.AssistantTrainingId = assistantTrainingId
	this.TenantId = tenantId
	return &this
}

// NewAssistantTrainingProgressaseWithDefaults instantiates a new AssistantTrainingProgressase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantTrainingProgressaseWithDefaults() *AssistantTrainingProgressase {
	this := AssistantTrainingProgressase{}
	return &this
}

// GetAssistantTrainingId returns the AssistantTrainingId field value
func (o *AssistantTrainingProgressase) GetAssistantTrainingId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssistantTrainingId
}

// GetAssistantTrainingIdOk returns a tuple with the AssistantTrainingId field value
// and a boolean to check if the value has been set.
func (o *AssistantTrainingProgressase) GetAssistantTrainingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssistantTrainingId, true
}

// SetAssistantTrainingId sets field value
func (o *AssistantTrainingProgressase) SetAssistantTrainingId(v int32) {
	o.AssistantTrainingId = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssistantTrainingProgressase) GetStatus() AssistantTrainingProgressaseStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret AssistantTrainingProgressaseStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssistantTrainingProgressase) GetStatusOk() (*AssistantTrainingProgressaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AssistantTrainingProgressase) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAssistantTrainingProgressaseStatus and assigns it to the Status field.
func (o *AssistantTrainingProgressase) SetStatus(v AssistantTrainingProgressaseStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AssistantTrainingProgressase) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AssistantTrainingProgressase) UnsetStatus() {
	o.Status.Unset()
}

// GetTenantId returns the TenantId field value
func (o *AssistantTrainingProgressase) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AssistantTrainingProgressase) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AssistantTrainingProgressase) SetTenantId(v int32) {
	o.TenantId = v
}

// GetSuccesful returns the Succesful field value if set, zero value otherwise.
func (o *AssistantTrainingProgressase) GetSuccesful() bool {
	if o == nil || IsNil(o.Succesful) {
		var ret bool
		return ret
	}
	return *o.Succesful
}

// GetSuccesfulOk returns a tuple with the Succesful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantTrainingProgressase) GetSuccesfulOk() (*bool, bool) {
	if o == nil || IsNil(o.Succesful) {
		return nil, false
	}
	return o.Succesful, true
}

// HasSuccesful returns a boolean if a field has been set.
func (o *AssistantTrainingProgressase) HasSuccesful() bool {
	if o != nil && !IsNil(o.Succesful) {
		return true
	}

	return false
}

// SetSuccesful gets a reference to the given bool and assigns it to the Succesful field.
func (o *AssistantTrainingProgressase) SetSuccesful(v bool) {
	o.Succesful = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssistantTrainingProgressase) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssistantTrainingProgressase) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *AssistantTrainingProgressase) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *AssistantTrainingProgressase) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *AssistantTrainingProgressase) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *AssistantTrainingProgressase) UnsetMessage() {
	o.Message.Unset()
}

func (o AssistantTrainingProgressase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantTrainingProgressase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assistant_training_id"] = o.AssistantTrainingId
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["tenant_id"] = o.TenantId
	if !IsNil(o.Succesful) {
		toSerialize["succesful"] = o.Succesful
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

func (o *AssistantTrainingProgressase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assistant_training_id",
		"tenant_id",
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

	varAssistantTrainingProgressase := _AssistantTrainingProgressase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssistantTrainingProgressase)

	if err != nil {
		return err
	}

	*o = AssistantTrainingProgressase(varAssistantTrainingProgressase)

	return err
}

type NullableAssistantTrainingProgressase struct {
	value *AssistantTrainingProgressase
	isSet bool
}

func (v NullableAssistantTrainingProgressase) Get() *AssistantTrainingProgressase {
	return v.value
}

func (v *NullableAssistantTrainingProgressase) Set(val *AssistantTrainingProgressase) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantTrainingProgressase) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantTrainingProgressase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantTrainingProgressase(val *AssistantTrainingProgressase) *NullableAssistantTrainingProgressase {
	return &NullableAssistantTrainingProgressase{value: val, isSet: true}
}

func (v NullableAssistantTrainingProgressase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantTrainingProgressase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


