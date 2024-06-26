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

// checks if the Thread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Thread{}

// Thread struct for Thread
type Thread struct {
	// The thread ID UUID for an instance of a thread.
	ThreadId int32 `json:"thread_id"`
	Issue Nested `json:"issue"`
}

type _Thread Thread

// NewThread instantiates a new Thread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThread(threadId int32, issue Nested) *Thread {
	this := Thread{}
	this.ThreadId = threadId
	this.Issue = issue
	return &this
}

// NewThreadWithDefaults instantiates a new Thread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadWithDefaults() *Thread {
	this := Thread{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *Thread) GetThreadId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *Thread) GetThreadIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *Thread) SetThreadId(v int32) {
	o.ThreadId = v
}

// GetIssue returns the Issue field value
func (o *Thread) GetIssue() Nested {
	if o == nil {
		var ret Nested
		return ret
	}

	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value
// and a boolean to check if the value has been set.
func (o *Thread) GetIssueOk() (*Nested, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issue, true
}

// SetIssue sets field value
func (o *Thread) SetIssue(v Nested) {
	o.Issue = v
}

func (o Thread) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Thread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thread_id"] = o.ThreadId
	toSerialize["issue"] = o.Issue
	return toSerialize, nil
}

func (o *Thread) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"thread_id",
		"issue",
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

	varThread := _Thread{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThread)

	if err != nil {
		return err
	}

	*o = Thread(varThread)

	return err
}

type NullableThread struct {
	value *Thread
	isSet bool
}

func (v NullableThread) Get() *Thread {
	return v.value
}

func (v *NullableThread) Set(val *Thread) {
	v.value = val
	v.isSet = true
}

func (v NullableThread) IsSet() bool {
	return v.isSet
}

func (v *NullableThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThread(val *Thread) *NullableThread {
	return &NullableThread{value: val, isSet: true}
}

func (v NullableThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


