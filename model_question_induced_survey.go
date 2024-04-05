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

// checks if the QuestionInducedSurvey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionInducedSurvey{}

// QuestionInducedSurvey struct for QuestionInducedSurvey
type QuestionInducedSurvey struct {
	// The lilling details id ID UUID
	QuestionInducedSurveyId int32 `json:"question_induced_survey_id"`
	TenantId int32 `json:"tenant_id"`
	// The survey topic
	SurveyTopic string `json:"survey_topic"`
	// The survey description
	SurveyDescription string `json:"survey_description"`
	// The survey questions
	QuestionsInduced interface{} `json:"questions_induced,omitempty"`
	// The survey questions
	SurveyQuestions interface{} `json:"survey_questions,omitempty"`
	Client Client `json:"client"`
	UnauthenticatedClient AnonymousUser `json:"unauthenticated_client"`
	SurveyType NullablePromptSurveySurveyType `json:"survey_type,omitempty"`
}

type _QuestionInducedSurvey QuestionInducedSurvey

// NewQuestionInducedSurvey instantiates a new QuestionInducedSurvey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionInducedSurvey(questionInducedSurveyId int32, tenantId int32, surveyTopic string, surveyDescription string, client Client, unauthenticatedClient AnonymousUser) *QuestionInducedSurvey {
	this := QuestionInducedSurvey{}
	this.QuestionInducedSurveyId = questionInducedSurveyId
	this.TenantId = tenantId
	this.SurveyTopic = surveyTopic
	this.SurveyDescription = surveyDescription
	this.Client = client
	this.UnauthenticatedClient = unauthenticatedClient
	return &this
}

// NewQuestionInducedSurveyWithDefaults instantiates a new QuestionInducedSurvey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionInducedSurveyWithDefaults() *QuestionInducedSurvey {
	this := QuestionInducedSurvey{}
	return &this
}

// GetQuestionInducedSurveyId returns the QuestionInducedSurveyId field value
func (o *QuestionInducedSurvey) GetQuestionInducedSurveyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuestionInducedSurveyId
}

// GetQuestionInducedSurveyIdOk returns a tuple with the QuestionInducedSurveyId field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetQuestionInducedSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestionInducedSurveyId, true
}

// SetQuestionInducedSurveyId sets field value
func (o *QuestionInducedSurvey) SetQuestionInducedSurveyId(v int32) {
	o.QuestionInducedSurveyId = v
}

// GetTenantId returns the TenantId field value
func (o *QuestionInducedSurvey) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *QuestionInducedSurvey) SetTenantId(v int32) {
	o.TenantId = v
}

// GetSurveyTopic returns the SurveyTopic field value
func (o *QuestionInducedSurvey) GetSurveyTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SurveyTopic
}

// GetSurveyTopicOk returns a tuple with the SurveyTopic field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetSurveyTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SurveyTopic, true
}

// SetSurveyTopic sets field value
func (o *QuestionInducedSurvey) SetSurveyTopic(v string) {
	o.SurveyTopic = v
}

// GetSurveyDescription returns the SurveyDescription field value
func (o *QuestionInducedSurvey) GetSurveyDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SurveyDescription
}

// GetSurveyDescriptionOk returns a tuple with the SurveyDescription field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetSurveyDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SurveyDescription, true
}

// SetSurveyDescription sets field value
func (o *QuestionInducedSurvey) SetSurveyDescription(v string) {
	o.SurveyDescription = v
}

// GetQuestionsInduced returns the QuestionsInduced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuestionInducedSurvey) GetQuestionsInduced() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.QuestionsInduced
}

// GetQuestionsInducedOk returns a tuple with the QuestionsInduced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuestionInducedSurvey) GetQuestionsInducedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QuestionsInduced) {
		return nil, false
	}
	return &o.QuestionsInduced, true
}

// HasQuestionsInduced returns a boolean if a field has been set.
func (o *QuestionInducedSurvey) HasQuestionsInduced() bool {
	if o != nil && !IsNil(o.QuestionsInduced) {
		return true
	}

	return false
}

// SetQuestionsInduced gets a reference to the given interface{} and assigns it to the QuestionsInduced field.
func (o *QuestionInducedSurvey) SetQuestionsInduced(v interface{}) {
	o.QuestionsInduced = v
}

// GetSurveyQuestions returns the SurveyQuestions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuestionInducedSurvey) GetSurveyQuestions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SurveyQuestions
}

// GetSurveyQuestionsOk returns a tuple with the SurveyQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuestionInducedSurvey) GetSurveyQuestionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SurveyQuestions) {
		return nil, false
	}
	return &o.SurveyQuestions, true
}

// HasSurveyQuestions returns a boolean if a field has been set.
func (o *QuestionInducedSurvey) HasSurveyQuestions() bool {
	if o != nil && !IsNil(o.SurveyQuestions) {
		return true
	}

	return false
}

// SetSurveyQuestions gets a reference to the given interface{} and assigns it to the SurveyQuestions field.
func (o *QuestionInducedSurvey) SetSurveyQuestions(v interface{}) {
	o.SurveyQuestions = v
}

// GetClient returns the Client field value
func (o *QuestionInducedSurvey) GetClient() Client {
	if o == nil {
		var ret Client
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetClientOk() (*Client, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *QuestionInducedSurvey) SetClient(v Client) {
	o.Client = v
}

// GetUnauthenticatedClient returns the UnauthenticatedClient field value
func (o *QuestionInducedSurvey) GetUnauthenticatedClient() AnonymousUser {
	if o == nil {
		var ret AnonymousUser
		return ret
	}

	return o.UnauthenticatedClient
}

// GetUnauthenticatedClientOk returns a tuple with the UnauthenticatedClient field value
// and a boolean to check if the value has been set.
func (o *QuestionInducedSurvey) GetUnauthenticatedClientOk() (*AnonymousUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnauthenticatedClient, true
}

// SetUnauthenticatedClient sets field value
func (o *QuestionInducedSurvey) SetUnauthenticatedClient(v AnonymousUser) {
	o.UnauthenticatedClient = v
}

// GetSurveyType returns the SurveyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuestionInducedSurvey) GetSurveyType() PromptSurveySurveyType {
	if o == nil || IsNil(o.SurveyType.Get()) {
		var ret PromptSurveySurveyType
		return ret
	}
	return *o.SurveyType.Get()
}

// GetSurveyTypeOk returns a tuple with the SurveyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuestionInducedSurvey) GetSurveyTypeOk() (*PromptSurveySurveyType, bool) {
	if o == nil {
		return nil, false
	}
	return o.SurveyType.Get(), o.SurveyType.IsSet()
}

// HasSurveyType returns a boolean if a field has been set.
func (o *QuestionInducedSurvey) HasSurveyType() bool {
	if o != nil && o.SurveyType.IsSet() {
		return true
	}

	return false
}

// SetSurveyType gets a reference to the given NullablePromptSurveySurveyType and assigns it to the SurveyType field.
func (o *QuestionInducedSurvey) SetSurveyType(v PromptSurveySurveyType) {
	o.SurveyType.Set(&v)
}
// SetSurveyTypeNil sets the value for SurveyType to be an explicit nil
func (o *QuestionInducedSurvey) SetSurveyTypeNil() {
	o.SurveyType.Set(nil)
}

// UnsetSurveyType ensures that no value is present for SurveyType, not even an explicit nil
func (o *QuestionInducedSurvey) UnsetSurveyType() {
	o.SurveyType.Unset()
}

func (o QuestionInducedSurvey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionInducedSurvey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["question_induced_survey_id"] = o.QuestionInducedSurveyId
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["survey_topic"] = o.SurveyTopic
	toSerialize["survey_description"] = o.SurveyDescription
	if o.QuestionsInduced != nil {
		toSerialize["questions_induced"] = o.QuestionsInduced
	}
	if o.SurveyQuestions != nil {
		toSerialize["survey_questions"] = o.SurveyQuestions
	}
	toSerialize["client"] = o.Client
	toSerialize["unauthenticated_client"] = o.UnauthenticatedClient
	if o.SurveyType.IsSet() {
		toSerialize["survey_type"] = o.SurveyType.Get()
	}
	return toSerialize, nil
}

func (o *QuestionInducedSurvey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"question_induced_survey_id",
		"tenant_id",
		"survey_topic",
		"survey_description",
		"client",
		"unauthenticated_client",
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

	varQuestionInducedSurvey := _QuestionInducedSurvey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestionInducedSurvey)

	if err != nil {
		return err
	}

	*o = QuestionInducedSurvey(varQuestionInducedSurvey)

	return err
}

type NullableQuestionInducedSurvey struct {
	value *QuestionInducedSurvey
	isSet bool
}

func (v NullableQuestionInducedSurvey) Get() *QuestionInducedSurvey {
	return v.value
}

func (v *NullableQuestionInducedSurvey) Set(val *QuestionInducedSurvey) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionInducedSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionInducedSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionInducedSurvey(val *QuestionInducedSurvey) *NullableQuestionInducedSurvey {
	return &NullableQuestionInducedSurvey{value: val, isSet: true}
}

func (v NullableQuestionInducedSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionInducedSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


