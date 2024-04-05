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

// checks if the SurveyReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveyReport{}

// SurveyReport struct for SurveyReport
type SurveyReport struct {
	// The survey report id
	SurveyReportId int32 `json:"survey_report_id"`
	// Display name of the survey
	SurveyId int32 `json:"survey_id"`
	// The survey subgroup name
	Conclusion string `json:"conclusion"`
	// The survey success
	SurveySuccess *bool `json:"survey_success,omitempty"`
	// Employee who made the survey review
	SurveyReporter NullableInt32 `json:"survey_reporter,omitempty"`
}

type _SurveyReport SurveyReport

// NewSurveyReport instantiates a new SurveyReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveyReport(surveyReportId int32, surveyId int32, conclusion string) *SurveyReport {
	this := SurveyReport{}
	this.SurveyReportId = surveyReportId
	this.SurveyId = surveyId
	this.Conclusion = conclusion
	return &this
}

// NewSurveyReportWithDefaults instantiates a new SurveyReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyReportWithDefaults() *SurveyReport {
	this := SurveyReport{}
	return &this
}

// GetSurveyReportId returns the SurveyReportId field value
func (o *SurveyReport) GetSurveyReportId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SurveyReportId
}

// GetSurveyReportIdOk returns a tuple with the SurveyReportId field value
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetSurveyReportIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SurveyReportId, true
}

// SetSurveyReportId sets field value
func (o *SurveyReport) SetSurveyReportId(v int32) {
	o.SurveyReportId = v
}

// GetSurveyId returns the SurveyId field value
func (o *SurveyReport) GetSurveyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SurveyId
}

// GetSurveyIdOk returns a tuple with the SurveyId field value
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SurveyId, true
}

// SetSurveyId sets field value
func (o *SurveyReport) SetSurveyId(v int32) {
	o.SurveyId = v
}

// GetConclusion returns the Conclusion field value
func (o *SurveyReport) GetConclusion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetConclusionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conclusion, true
}

// SetConclusion sets field value
func (o *SurveyReport) SetConclusion(v string) {
	o.Conclusion = v
}

// GetSurveySuccess returns the SurveySuccess field value if set, zero value otherwise.
func (o *SurveyReport) GetSurveySuccess() bool {
	if o == nil || IsNil(o.SurveySuccess) {
		var ret bool
		return ret
	}
	return *o.SurveySuccess
}

// GetSurveySuccessOk returns a tuple with the SurveySuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetSurveySuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.SurveySuccess) {
		return nil, false
	}
	return o.SurveySuccess, true
}

// HasSurveySuccess returns a boolean if a field has been set.
func (o *SurveyReport) HasSurveySuccess() bool {
	if o != nil && !IsNil(o.SurveySuccess) {
		return true
	}

	return false
}

// SetSurveySuccess gets a reference to the given bool and assigns it to the SurveySuccess field.
func (o *SurveyReport) SetSurveySuccess(v bool) {
	o.SurveySuccess = &v
}

// GetSurveyReporter returns the SurveyReporter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveyReport) GetSurveyReporter() int32 {
	if o == nil || IsNil(o.SurveyReporter.Get()) {
		var ret int32
		return ret
	}
	return *o.SurveyReporter.Get()
}

// GetSurveyReporterOk returns a tuple with the SurveyReporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveyReport) GetSurveyReporterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SurveyReporter.Get(), o.SurveyReporter.IsSet()
}

// HasSurveyReporter returns a boolean if a field has been set.
func (o *SurveyReport) HasSurveyReporter() bool {
	if o != nil && o.SurveyReporter.IsSet() {
		return true
	}

	return false
}

// SetSurveyReporter gets a reference to the given NullableInt32 and assigns it to the SurveyReporter field.
func (o *SurveyReport) SetSurveyReporter(v int32) {
	o.SurveyReporter.Set(&v)
}
// SetSurveyReporterNil sets the value for SurveyReporter to be an explicit nil
func (o *SurveyReport) SetSurveyReporterNil() {
	o.SurveyReporter.Set(nil)
}

// UnsetSurveyReporter ensures that no value is present for SurveyReporter, not even an explicit nil
func (o *SurveyReport) UnsetSurveyReporter() {
	o.SurveyReporter.Unset()
}

func (o SurveyReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveyReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["survey_report_id"] = o.SurveyReportId
	toSerialize["survey_id"] = o.SurveyId
	toSerialize["conclusion"] = o.Conclusion
	if !IsNil(o.SurveySuccess) {
		toSerialize["survey_success"] = o.SurveySuccess
	}
	if o.SurveyReporter.IsSet() {
		toSerialize["survey_reporter"] = o.SurveyReporter.Get()
	}
	return toSerialize, nil
}

func (o *SurveyReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"survey_report_id",
		"survey_id",
		"conclusion",
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

	varSurveyReport := _SurveyReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSurveyReport)

	if err != nil {
		return err
	}

	*o = SurveyReport(varSurveyReport)

	return err
}

type NullableSurveyReport struct {
	value *SurveyReport
	isSet bool
}

func (v NullableSurveyReport) Get() *SurveyReport {
	return v.value
}

func (v *NullableSurveyReport) Set(val *SurveyReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveyReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveyReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveyReport(val *SurveyReport) *NullableSurveyReport {
	return &NullableSurveyReport{value: val, isSet: true}
}

func (v NullableSurveyReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveyReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

