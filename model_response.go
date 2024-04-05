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

// checks if the Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Response{}

// Response struct for Response
type Response struct {
	// The lilling details id ID UUID
	ResponseId int32 `json:"response_id"`
	// Display name of the tenant
	SurveyId NullableInt32 `json:"survey_id,omitempty"`
	// Display name of the tenant
	PromptSurveyId NullableInt32 `json:"prompt_survey_id,omitempty"`
	// Display name of the tenant
	QuestionsInducedSurveyId NullableInt32 `json:"questions_induced_survey_id,omitempty"`
	Client TargetAudience `json:"client"`
	// The survey response
	SurveyResponse interface{} `json:"survey_response,omitempty"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	CreatedAt NullableString `json:"created_at"`
}

type _Response Response

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse(responseId int32, client TargetAudience, timestamp NullableTime, createdAt NullableString) *Response {
	this := Response{}
	this.ResponseId = responseId
	this.Client = client
	this.Timestamp = timestamp
	this.CreatedAt = createdAt
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetResponseId returns the ResponseId field value
func (o *Response) GetResponseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value
// and a boolean to check if the value has been set.
func (o *Response) GetResponseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseId, true
}

// SetResponseId sets field value
func (o *Response) SetResponseId(v int32) {
	o.ResponseId = v
}

// GetSurveyId returns the SurveyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetSurveyId() int32 {
	if o == nil || IsNil(o.SurveyId.Get()) {
		var ret int32
		return ret
	}
	return *o.SurveyId.Get()
}

// GetSurveyIdOk returns a tuple with the SurveyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SurveyId.Get(), o.SurveyId.IsSet()
}

// HasSurveyId returns a boolean if a field has been set.
func (o *Response) HasSurveyId() bool {
	if o != nil && o.SurveyId.IsSet() {
		return true
	}

	return false
}

// SetSurveyId gets a reference to the given NullableInt32 and assigns it to the SurveyId field.
func (o *Response) SetSurveyId(v int32) {
	o.SurveyId.Set(&v)
}
// SetSurveyIdNil sets the value for SurveyId to be an explicit nil
func (o *Response) SetSurveyIdNil() {
	o.SurveyId.Set(nil)
}

// UnsetSurveyId ensures that no value is present for SurveyId, not even an explicit nil
func (o *Response) UnsetSurveyId() {
	o.SurveyId.Unset()
}

// GetPromptSurveyId returns the PromptSurveyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetPromptSurveyId() int32 {
	if o == nil || IsNil(o.PromptSurveyId.Get()) {
		var ret int32
		return ret
	}
	return *o.PromptSurveyId.Get()
}

// GetPromptSurveyIdOk returns a tuple with the PromptSurveyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetPromptSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromptSurveyId.Get(), o.PromptSurveyId.IsSet()
}

// HasPromptSurveyId returns a boolean if a field has been set.
func (o *Response) HasPromptSurveyId() bool {
	if o != nil && o.PromptSurveyId.IsSet() {
		return true
	}

	return false
}

// SetPromptSurveyId gets a reference to the given NullableInt32 and assigns it to the PromptSurveyId field.
func (o *Response) SetPromptSurveyId(v int32) {
	o.PromptSurveyId.Set(&v)
}
// SetPromptSurveyIdNil sets the value for PromptSurveyId to be an explicit nil
func (o *Response) SetPromptSurveyIdNil() {
	o.PromptSurveyId.Set(nil)
}

// UnsetPromptSurveyId ensures that no value is present for PromptSurveyId, not even an explicit nil
func (o *Response) UnsetPromptSurveyId() {
	o.PromptSurveyId.Unset()
}

// GetQuestionsInducedSurveyId returns the QuestionsInducedSurveyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetQuestionsInducedSurveyId() int32 {
	if o == nil || IsNil(o.QuestionsInducedSurveyId.Get()) {
		var ret int32
		return ret
	}
	return *o.QuestionsInducedSurveyId.Get()
}

// GetQuestionsInducedSurveyIdOk returns a tuple with the QuestionsInducedSurveyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetQuestionsInducedSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuestionsInducedSurveyId.Get(), o.QuestionsInducedSurveyId.IsSet()
}

// HasQuestionsInducedSurveyId returns a boolean if a field has been set.
func (o *Response) HasQuestionsInducedSurveyId() bool {
	if o != nil && o.QuestionsInducedSurveyId.IsSet() {
		return true
	}

	return false
}

// SetQuestionsInducedSurveyId gets a reference to the given NullableInt32 and assigns it to the QuestionsInducedSurveyId field.
func (o *Response) SetQuestionsInducedSurveyId(v int32) {
	o.QuestionsInducedSurveyId.Set(&v)
}
// SetQuestionsInducedSurveyIdNil sets the value for QuestionsInducedSurveyId to be an explicit nil
func (o *Response) SetQuestionsInducedSurveyIdNil() {
	o.QuestionsInducedSurveyId.Set(nil)
}

// UnsetQuestionsInducedSurveyId ensures that no value is present for QuestionsInducedSurveyId, not even an explicit nil
func (o *Response) UnsetQuestionsInducedSurveyId() {
	o.QuestionsInducedSurveyId.Unset()
}

// GetClient returns the Client field value
func (o *Response) GetClient() TargetAudience {
	if o == nil {
		var ret TargetAudience
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *Response) GetClientOk() (*TargetAudience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *Response) SetClient(v TargetAudience) {
	o.Client = v
}

// GetSurveyResponse returns the SurveyResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Response) GetSurveyResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SurveyResponse
}

// GetSurveyResponseOk returns a tuple with the SurveyResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetSurveyResponseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SurveyResponse) {
		return nil, false
	}
	return &o.SurveyResponse, true
}

// HasSurveyResponse returns a boolean if a field has been set.
func (o *Response) HasSurveyResponse() bool {
	if o != nil && !IsNil(o.SurveyResponse) {
		return true
	}

	return false
}

// SetSurveyResponse gets a reference to the given interface{} and assigns it to the SurveyResponse field.
func (o *Response) SetSurveyResponse(v interface{}) {
	o.SurveyResponse = v
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Response) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *Response) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Response) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Response) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response_id"] = o.ResponseId
	if o.SurveyId.IsSet() {
		toSerialize["survey_id"] = o.SurveyId.Get()
	}
	if o.PromptSurveyId.IsSet() {
		toSerialize["prompt_survey_id"] = o.PromptSurveyId.Get()
	}
	if o.QuestionsInducedSurveyId.IsSet() {
		toSerialize["questions_induced_survey_id"] = o.QuestionsInducedSurveyId.Get()
	}
	toSerialize["client"] = o.Client
	if o.SurveyResponse != nil {
		toSerialize["survey_response"] = o.SurveyResponse
	}
	toSerialize["timestamp"] = o.Timestamp.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()
	return toSerialize, nil
}

func (o *Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"response_id",
		"client",
		"timestamp",
		"created_at",
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

	varResponse := _Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponse)

	if err != nil {
		return err
	}

	*o = Response(varResponse)

	return err
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


