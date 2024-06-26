/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedInteractionHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedInteractionHistory{}

// PatchedInteractionHistory struct for PatchedInteractionHistory
type PatchedInteractionHistory struct {
	// The call id.
	InteractionHistoryId *int32 `json:"interaction_history_id,omitempty"`
	Call *Call `json:"call,omitempty"`
	// Type of interaction.  * `NOTE` - Note * `DECISION` - Decision * `ESCALATION` - Escalation
	InteractionType *InteractionTypeEnum `json:"interaction_type,omitempty"`
	// Details of the interaction.
	Details *string `json:"details,omitempty"`
}

// NewPatchedInteractionHistory instantiates a new PatchedInteractionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedInteractionHistory() *PatchedInteractionHistory {
	this := PatchedInteractionHistory{}
	return &this
}

// NewPatchedInteractionHistoryWithDefaults instantiates a new PatchedInteractionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedInteractionHistoryWithDefaults() *PatchedInteractionHistory {
	this := PatchedInteractionHistory{}
	return &this
}

// GetInteractionHistoryId returns the InteractionHistoryId field value if set, zero value otherwise.
func (o *PatchedInteractionHistory) GetInteractionHistoryId() int32 {
	if o == nil || IsNil(o.InteractionHistoryId) {
		var ret int32
		return ret
	}
	return *o.InteractionHistoryId
}

// GetInteractionHistoryIdOk returns a tuple with the InteractionHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInteractionHistory) GetInteractionHistoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InteractionHistoryId) {
		return nil, false
	}
	return o.InteractionHistoryId, true
}

// HasInteractionHistoryId returns a boolean if a field has been set.
func (o *PatchedInteractionHistory) HasInteractionHistoryId() bool {
	if o != nil && !IsNil(o.InteractionHistoryId) {
		return true
	}

	return false
}

// SetInteractionHistoryId gets a reference to the given int32 and assigns it to the InteractionHistoryId field.
func (o *PatchedInteractionHistory) SetInteractionHistoryId(v int32) {
	o.InteractionHistoryId = &v
}

// GetCall returns the Call field value if set, zero value otherwise.
func (o *PatchedInteractionHistory) GetCall() Call {
	if o == nil || IsNil(o.Call) {
		var ret Call
		return ret
	}
	return *o.Call
}

// GetCallOk returns a tuple with the Call field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInteractionHistory) GetCallOk() (*Call, bool) {
	if o == nil || IsNil(o.Call) {
		return nil, false
	}
	return o.Call, true
}

// HasCall returns a boolean if a field has been set.
func (o *PatchedInteractionHistory) HasCall() bool {
	if o != nil && !IsNil(o.Call) {
		return true
	}

	return false
}

// SetCall gets a reference to the given Call and assigns it to the Call field.
func (o *PatchedInteractionHistory) SetCall(v Call) {
	o.Call = &v
}

// GetInteractionType returns the InteractionType field value if set, zero value otherwise.
func (o *PatchedInteractionHistory) GetInteractionType() InteractionTypeEnum {
	if o == nil || IsNil(o.InteractionType) {
		var ret InteractionTypeEnum
		return ret
	}
	return *o.InteractionType
}

// GetInteractionTypeOk returns a tuple with the InteractionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInteractionHistory) GetInteractionTypeOk() (*InteractionTypeEnum, bool) {
	if o == nil || IsNil(o.InteractionType) {
		return nil, false
	}
	return o.InteractionType, true
}

// HasInteractionType returns a boolean if a field has been set.
func (o *PatchedInteractionHistory) HasInteractionType() bool {
	if o != nil && !IsNil(o.InteractionType) {
		return true
	}

	return false
}

// SetInteractionType gets a reference to the given InteractionTypeEnum and assigns it to the InteractionType field.
func (o *PatchedInteractionHistory) SetInteractionType(v InteractionTypeEnum) {
	o.InteractionType = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PatchedInteractionHistory) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInteractionHistory) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PatchedInteractionHistory) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *PatchedInteractionHistory) SetDetails(v string) {
	o.Details = &v
}

func (o PatchedInteractionHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedInteractionHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InteractionHistoryId) {
		toSerialize["interaction_history_id"] = o.InteractionHistoryId
	}
	if !IsNil(o.Call) {
		toSerialize["call"] = o.Call
	}
	if !IsNil(o.InteractionType) {
		toSerialize["interaction_type"] = o.InteractionType
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullablePatchedInteractionHistory struct {
	value *PatchedInteractionHistory
	isSet bool
}

func (v NullablePatchedInteractionHistory) Get() *PatchedInteractionHistory {
	return v.value
}

func (v *NullablePatchedInteractionHistory) Set(val *PatchedInteractionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedInteractionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedInteractionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedInteractionHistory(val *PatchedInteractionHistory) *NullablePatchedInteractionHistory {
	return &NullablePatchedInteractionHistory{value: val, isSet: true}
}

func (v NullablePatchedInteractionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedInteractionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


