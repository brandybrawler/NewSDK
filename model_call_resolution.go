/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CallResolution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallResolution{}

// CallResolution struct for CallResolution
type CallResolution struct {
	// The call resolution id.
	CallResolutionId int32 `json:"call_resolution_id"`
	Call Call `json:"call"`
	Status *CallResolutionStatusEnum `json:"status,omitempty"`
	PendingActions interface{} `json:"pending_actions,omitempty"`
	FollowUpRequired *bool `json:"follow_up_required,omitempty"`
	FollowUpDate NullableTime `json:"follow_up_date,omitempty"`
}

type _CallResolution CallResolution

// NewCallResolution instantiates a new CallResolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallResolution(callResolutionId int32, call Call) *CallResolution {
	this := CallResolution{}
	this.CallResolutionId = callResolutionId
	this.Call = call
	return &this
}

// NewCallResolutionWithDefaults instantiates a new CallResolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallResolutionWithDefaults() *CallResolution {
	this := CallResolution{}
	return &this
}

// GetCallResolutionId returns the CallResolutionId field value
func (o *CallResolution) GetCallResolutionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CallResolutionId
}

// GetCallResolutionIdOk returns a tuple with the CallResolutionId field value
// and a boolean to check if the value has been set.
func (o *CallResolution) GetCallResolutionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallResolutionId, true
}

// SetCallResolutionId sets field value
func (o *CallResolution) SetCallResolutionId(v int32) {
	o.CallResolutionId = v
}

// GetCall returns the Call field value
func (o *CallResolution) GetCall() Call {
	if o == nil {
		var ret Call
		return ret
	}

	return o.Call
}

// GetCallOk returns a tuple with the Call field value
// and a boolean to check if the value has been set.
func (o *CallResolution) GetCallOk() (*Call, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Call, true
}

// SetCall sets field value
func (o *CallResolution) SetCall(v Call) {
	o.Call = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CallResolution) GetStatus() CallResolutionStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret CallResolutionStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallResolution) GetStatusOk() (*CallResolutionStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CallResolution) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CallResolutionStatusEnum and assigns it to the Status field.
func (o *CallResolution) SetStatus(v CallResolutionStatusEnum) {
	o.Status = &v
}

// GetPendingActions returns the PendingActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallResolution) GetPendingActions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PendingActions
}

// GetPendingActionsOk returns a tuple with the PendingActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallResolution) GetPendingActionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PendingActions) {
		return nil, false
	}
	return &o.PendingActions, true
}

// HasPendingActions returns a boolean if a field has been set.
func (o *CallResolution) HasPendingActions() bool {
	if o != nil && !IsNil(o.PendingActions) {
		return true
	}

	return false
}

// SetPendingActions gets a reference to the given interface{} and assigns it to the PendingActions field.
func (o *CallResolution) SetPendingActions(v interface{}) {
	o.PendingActions = v
}

// GetFollowUpRequired returns the FollowUpRequired field value if set, zero value otherwise.
func (o *CallResolution) GetFollowUpRequired() bool {
	if o == nil || IsNil(o.FollowUpRequired) {
		var ret bool
		return ret
	}
	return *o.FollowUpRequired
}

// GetFollowUpRequiredOk returns a tuple with the FollowUpRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallResolution) GetFollowUpRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowUpRequired) {
		return nil, false
	}
	return o.FollowUpRequired, true
}

// HasFollowUpRequired returns a boolean if a field has been set.
func (o *CallResolution) HasFollowUpRequired() bool {
	if o != nil && !IsNil(o.FollowUpRequired) {
		return true
	}

	return false
}

// SetFollowUpRequired gets a reference to the given bool and assigns it to the FollowUpRequired field.
func (o *CallResolution) SetFollowUpRequired(v bool) {
	o.FollowUpRequired = &v
}

// GetFollowUpDate returns the FollowUpDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallResolution) GetFollowUpDate() time.Time {
	if o == nil || IsNil(o.FollowUpDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FollowUpDate.Get()
}

// GetFollowUpDateOk returns a tuple with the FollowUpDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallResolution) GetFollowUpDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowUpDate.Get(), o.FollowUpDate.IsSet()
}

// HasFollowUpDate returns a boolean if a field has been set.
func (o *CallResolution) HasFollowUpDate() bool {
	if o != nil && o.FollowUpDate.IsSet() {
		return true
	}

	return false
}

// SetFollowUpDate gets a reference to the given NullableTime and assigns it to the FollowUpDate field.
func (o *CallResolution) SetFollowUpDate(v time.Time) {
	o.FollowUpDate.Set(&v)
}
// SetFollowUpDateNil sets the value for FollowUpDate to be an explicit nil
func (o *CallResolution) SetFollowUpDateNil() {
	o.FollowUpDate.Set(nil)
}

// UnsetFollowUpDate ensures that no value is present for FollowUpDate, not even an explicit nil
func (o *CallResolution) UnsetFollowUpDate() {
	o.FollowUpDate.Unset()
}

func (o CallResolution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallResolution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["call_resolution_id"] = o.CallResolutionId
	toSerialize["call"] = o.Call
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.PendingActions != nil {
		toSerialize["pending_actions"] = o.PendingActions
	}
	if !IsNil(o.FollowUpRequired) {
		toSerialize["follow_up_required"] = o.FollowUpRequired
	}
	if o.FollowUpDate.IsSet() {
		toSerialize["follow_up_date"] = o.FollowUpDate.Get()
	}
	return toSerialize, nil
}

func (o *CallResolution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"call_resolution_id",
		"call",
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

	varCallResolution := _CallResolution{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCallResolution)

	if err != nil {
		return err
	}

	*o = CallResolution(varCallResolution)

	return err
}

type NullableCallResolution struct {
	value *CallResolution
	isSet bool
}

func (v NullableCallResolution) Get() *CallResolution {
	return v.value
}

func (v *NullableCallResolution) Set(val *CallResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableCallResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableCallResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallResolution(val *CallResolution) *NullableCallResolution {
	return &NullableCallResolution{value: val, isSet: true}
}

func (v NullableCallResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


