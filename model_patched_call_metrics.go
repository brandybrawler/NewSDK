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

// checks if the PatchedCallMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCallMetrics{}

// PatchedCallMetrics struct for PatchedCallMetrics
type PatchedCallMetrics struct {
	// The call metrics id.
	CallMetricsId *int32 `json:"call_metrics_id,omitempty"`
	Call *Call `json:"call,omitempty"`
	Transcripts interface{} `json:"transcripts,omitempty"`
	CallSatisfaction *bool `json:"call_satisfaction,omitempty"`
	IssueResolution *bool `json:"issue_resolution,omitempty"`
}

// NewPatchedCallMetrics instantiates a new PatchedCallMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCallMetrics() *PatchedCallMetrics {
	this := PatchedCallMetrics{}
	return &this
}

// NewPatchedCallMetricsWithDefaults instantiates a new PatchedCallMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCallMetricsWithDefaults() *PatchedCallMetrics {
	this := PatchedCallMetrics{}
	return &this
}

// GetCallMetricsId returns the CallMetricsId field value if set, zero value otherwise.
func (o *PatchedCallMetrics) GetCallMetricsId() int32 {
	if o == nil || IsNil(o.CallMetricsId) {
		var ret int32
		return ret
	}
	return *o.CallMetricsId
}

// GetCallMetricsIdOk returns a tuple with the CallMetricsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCallMetrics) GetCallMetricsIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CallMetricsId) {
		return nil, false
	}
	return o.CallMetricsId, true
}

// HasCallMetricsId returns a boolean if a field has been set.
func (o *PatchedCallMetrics) HasCallMetricsId() bool {
	if o != nil && !IsNil(o.CallMetricsId) {
		return true
	}

	return false
}

// SetCallMetricsId gets a reference to the given int32 and assigns it to the CallMetricsId field.
func (o *PatchedCallMetrics) SetCallMetricsId(v int32) {
	o.CallMetricsId = &v
}

// GetCall returns the Call field value if set, zero value otherwise.
func (o *PatchedCallMetrics) GetCall() Call {
	if o == nil || IsNil(o.Call) {
		var ret Call
		return ret
	}
	return *o.Call
}

// GetCallOk returns a tuple with the Call field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCallMetrics) GetCallOk() (*Call, bool) {
	if o == nil || IsNil(o.Call) {
		return nil, false
	}
	return o.Call, true
}

// HasCall returns a boolean if a field has been set.
func (o *PatchedCallMetrics) HasCall() bool {
	if o != nil && !IsNil(o.Call) {
		return true
	}

	return false
}

// SetCall gets a reference to the given Call and assigns it to the Call field.
func (o *PatchedCallMetrics) SetCall(v Call) {
	o.Call = &v
}

// GetTranscripts returns the Transcripts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCallMetrics) GetTranscripts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Transcripts
}

// GetTranscriptsOk returns a tuple with the Transcripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCallMetrics) GetTranscriptsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Transcripts) {
		return nil, false
	}
	return &o.Transcripts, true
}

// HasTranscripts returns a boolean if a field has been set.
func (o *PatchedCallMetrics) HasTranscripts() bool {
	if o != nil && !IsNil(o.Transcripts) {
		return true
	}

	return false
}

// SetTranscripts gets a reference to the given interface{} and assigns it to the Transcripts field.
func (o *PatchedCallMetrics) SetTranscripts(v interface{}) {
	o.Transcripts = v
}

// GetCallSatisfaction returns the CallSatisfaction field value if set, zero value otherwise.
func (o *PatchedCallMetrics) GetCallSatisfaction() bool {
	if o == nil || IsNil(o.CallSatisfaction) {
		var ret bool
		return ret
	}
	return *o.CallSatisfaction
}

// GetCallSatisfactionOk returns a tuple with the CallSatisfaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCallMetrics) GetCallSatisfactionOk() (*bool, bool) {
	if o == nil || IsNil(o.CallSatisfaction) {
		return nil, false
	}
	return o.CallSatisfaction, true
}

// HasCallSatisfaction returns a boolean if a field has been set.
func (o *PatchedCallMetrics) HasCallSatisfaction() bool {
	if o != nil && !IsNil(o.CallSatisfaction) {
		return true
	}

	return false
}

// SetCallSatisfaction gets a reference to the given bool and assigns it to the CallSatisfaction field.
func (o *PatchedCallMetrics) SetCallSatisfaction(v bool) {
	o.CallSatisfaction = &v
}

// GetIssueResolution returns the IssueResolution field value if set, zero value otherwise.
func (o *PatchedCallMetrics) GetIssueResolution() bool {
	if o == nil || IsNil(o.IssueResolution) {
		var ret bool
		return ret
	}
	return *o.IssueResolution
}

// GetIssueResolutionOk returns a tuple with the IssueResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCallMetrics) GetIssueResolutionOk() (*bool, bool) {
	if o == nil || IsNil(o.IssueResolution) {
		return nil, false
	}
	return o.IssueResolution, true
}

// HasIssueResolution returns a boolean if a field has been set.
func (o *PatchedCallMetrics) HasIssueResolution() bool {
	if o != nil && !IsNil(o.IssueResolution) {
		return true
	}

	return false
}

// SetIssueResolution gets a reference to the given bool and assigns it to the IssueResolution field.
func (o *PatchedCallMetrics) SetIssueResolution(v bool) {
	o.IssueResolution = &v
}

func (o PatchedCallMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCallMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallMetricsId) {
		toSerialize["call_metrics_id"] = o.CallMetricsId
	}
	if !IsNil(o.Call) {
		toSerialize["call"] = o.Call
	}
	if o.Transcripts != nil {
		toSerialize["transcripts"] = o.Transcripts
	}
	if !IsNil(o.CallSatisfaction) {
		toSerialize["call_satisfaction"] = o.CallSatisfaction
	}
	if !IsNil(o.IssueResolution) {
		toSerialize["issue_resolution"] = o.IssueResolution
	}
	return toSerialize, nil
}

type NullablePatchedCallMetrics struct {
	value *PatchedCallMetrics
	isSet bool
}

func (v NullablePatchedCallMetrics) Get() *PatchedCallMetrics {
	return v.value
}

func (v *NullablePatchedCallMetrics) Set(val *PatchedCallMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCallMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCallMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCallMetrics(val *PatchedCallMetrics) *NullablePatchedCallMetrics {
	return &NullablePatchedCallMetrics{value: val, isSet: true}
}

func (v NullablePatchedCallMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCallMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

