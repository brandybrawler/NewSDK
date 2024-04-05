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

// checks if the Knowledgebase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Knowledgebase{}

// Knowledgebase struct for Knowledgebase
type Knowledgebase struct {
	// The tenant chats ID.
	KnowledgebaseId int32 `json:"knowledgebase_id"`
	// Display name of the tenant
	TenantId int32 `json:"tenant_id"`
	// The name of the knowledgebase
	Name string `json:"name"`
	// description
	Description NullableString `json:"description,omitempty"`
	TypeOfDatasource NullableString `json:"type_of_datasource,omitempty"`
	// This will be responsible for ft=etching the information source passed
	DatasourceTarget NullableString `json:"datasource_target,omitempty"`
	ChatPrompts []ChatPrompts `json:"chat_prompts"`
	CommunityPrompts []CommunityPrompts `json:"community_prompts"`
	ServicePrompts []ServicePrompts `json:"service_prompts"`
	SurveyPrompts []SurveyPrompts `json:"survey_prompts"`
}

type _Knowledgebase Knowledgebase

// NewKnowledgebase instantiates a new Knowledgebase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnowledgebase(knowledgebaseId int32, tenantId int32, name string, chatPrompts []ChatPrompts, communityPrompts []CommunityPrompts, servicePrompts []ServicePrompts, surveyPrompts []SurveyPrompts) *Knowledgebase {
	this := Knowledgebase{}
	this.KnowledgebaseId = knowledgebaseId
	this.TenantId = tenantId
	this.Name = name
	this.ChatPrompts = chatPrompts
	this.CommunityPrompts = communityPrompts
	this.ServicePrompts = servicePrompts
	this.SurveyPrompts = surveyPrompts
	return &this
}

// NewKnowledgebaseWithDefaults instantiates a new Knowledgebase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnowledgebaseWithDefaults() *Knowledgebase {
	this := Knowledgebase{}
	return &this
}

// GetKnowledgebaseId returns the KnowledgebaseId field value
func (o *Knowledgebase) GetKnowledgebaseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KnowledgebaseId
}

// GetKnowledgebaseIdOk returns a tuple with the KnowledgebaseId field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetKnowledgebaseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnowledgebaseId, true
}

// SetKnowledgebaseId sets field value
func (o *Knowledgebase) SetKnowledgebaseId(v int32) {
	o.KnowledgebaseId = v
}

// GetTenantId returns the TenantId field value
func (o *Knowledgebase) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Knowledgebase) SetTenantId(v int32) {
	o.TenantId = v
}

// GetName returns the Name field value
func (o *Knowledgebase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Knowledgebase) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Knowledgebase) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Knowledgebase) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Knowledgebase) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Knowledgebase) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Knowledgebase) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Knowledgebase) UnsetDescription() {
	o.Description.Unset()
}

// GetTypeOfDatasource returns the TypeOfDatasource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Knowledgebase) GetTypeOfDatasource() string {
	if o == nil || IsNil(o.TypeOfDatasource.Get()) {
		var ret string
		return ret
	}
	return *o.TypeOfDatasource.Get()
}

// GetTypeOfDatasourceOk returns a tuple with the TypeOfDatasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Knowledgebase) GetTypeOfDatasourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeOfDatasource.Get(), o.TypeOfDatasource.IsSet()
}

// HasTypeOfDatasource returns a boolean if a field has been set.
func (o *Knowledgebase) HasTypeOfDatasource() bool {
	if o != nil && o.TypeOfDatasource.IsSet() {
		return true
	}

	return false
}

// SetTypeOfDatasource gets a reference to the given NullableString and assigns it to the TypeOfDatasource field.
func (o *Knowledgebase) SetTypeOfDatasource(v string) {
	o.TypeOfDatasource.Set(&v)
}
// SetTypeOfDatasourceNil sets the value for TypeOfDatasource to be an explicit nil
func (o *Knowledgebase) SetTypeOfDatasourceNil() {
	o.TypeOfDatasource.Set(nil)
}

// UnsetTypeOfDatasource ensures that no value is present for TypeOfDatasource, not even an explicit nil
func (o *Knowledgebase) UnsetTypeOfDatasource() {
	o.TypeOfDatasource.Unset()
}

// GetDatasourceTarget returns the DatasourceTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Knowledgebase) GetDatasourceTarget() string {
	if o == nil || IsNil(o.DatasourceTarget.Get()) {
		var ret string
		return ret
	}
	return *o.DatasourceTarget.Get()
}

// GetDatasourceTargetOk returns a tuple with the DatasourceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Knowledgebase) GetDatasourceTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasourceTarget.Get(), o.DatasourceTarget.IsSet()
}

// HasDatasourceTarget returns a boolean if a field has been set.
func (o *Knowledgebase) HasDatasourceTarget() bool {
	if o != nil && o.DatasourceTarget.IsSet() {
		return true
	}

	return false
}

// SetDatasourceTarget gets a reference to the given NullableString and assigns it to the DatasourceTarget field.
func (o *Knowledgebase) SetDatasourceTarget(v string) {
	o.DatasourceTarget.Set(&v)
}
// SetDatasourceTargetNil sets the value for DatasourceTarget to be an explicit nil
func (o *Knowledgebase) SetDatasourceTargetNil() {
	o.DatasourceTarget.Set(nil)
}

// UnsetDatasourceTarget ensures that no value is present for DatasourceTarget, not even an explicit nil
func (o *Knowledgebase) UnsetDatasourceTarget() {
	o.DatasourceTarget.Unset()
}

// GetChatPrompts returns the ChatPrompts field value
func (o *Knowledgebase) GetChatPrompts() []ChatPrompts {
	if o == nil {
		var ret []ChatPrompts
		return ret
	}

	return o.ChatPrompts
}

// GetChatPromptsOk returns a tuple with the ChatPrompts field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetChatPromptsOk() ([]ChatPrompts, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChatPrompts, true
}

// SetChatPrompts sets field value
func (o *Knowledgebase) SetChatPrompts(v []ChatPrompts) {
	o.ChatPrompts = v
}

// GetCommunityPrompts returns the CommunityPrompts field value
func (o *Knowledgebase) GetCommunityPrompts() []CommunityPrompts {
	if o == nil {
		var ret []CommunityPrompts
		return ret
	}

	return o.CommunityPrompts
}

// GetCommunityPromptsOk returns a tuple with the CommunityPrompts field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetCommunityPromptsOk() ([]CommunityPrompts, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommunityPrompts, true
}

// SetCommunityPrompts sets field value
func (o *Knowledgebase) SetCommunityPrompts(v []CommunityPrompts) {
	o.CommunityPrompts = v
}

// GetServicePrompts returns the ServicePrompts field value
func (o *Knowledgebase) GetServicePrompts() []ServicePrompts {
	if o == nil {
		var ret []ServicePrompts
		return ret
	}

	return o.ServicePrompts
}

// GetServicePromptsOk returns a tuple with the ServicePrompts field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetServicePromptsOk() ([]ServicePrompts, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServicePrompts, true
}

// SetServicePrompts sets field value
func (o *Knowledgebase) SetServicePrompts(v []ServicePrompts) {
	o.ServicePrompts = v
}

// GetSurveyPrompts returns the SurveyPrompts field value
func (o *Knowledgebase) GetSurveyPrompts() []SurveyPrompts {
	if o == nil {
		var ret []SurveyPrompts
		return ret
	}

	return o.SurveyPrompts
}

// GetSurveyPromptsOk returns a tuple with the SurveyPrompts field value
// and a boolean to check if the value has been set.
func (o *Knowledgebase) GetSurveyPromptsOk() ([]SurveyPrompts, bool) {
	if o == nil {
		return nil, false
	}
	return o.SurveyPrompts, true
}

// SetSurveyPrompts sets field value
func (o *Knowledgebase) SetSurveyPrompts(v []SurveyPrompts) {
	o.SurveyPrompts = v
}

func (o Knowledgebase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Knowledgebase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["knowledgebase_id"] = o.KnowledgebaseId
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.TypeOfDatasource.IsSet() {
		toSerialize["type_of_datasource"] = o.TypeOfDatasource.Get()
	}
	if o.DatasourceTarget.IsSet() {
		toSerialize["datasource_target"] = o.DatasourceTarget.Get()
	}
	toSerialize["chat_prompts"] = o.ChatPrompts
	toSerialize["community_prompts"] = o.CommunityPrompts
	toSerialize["service_prompts"] = o.ServicePrompts
	toSerialize["survey_prompts"] = o.SurveyPrompts
	return toSerialize, nil
}

func (o *Knowledgebase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"knowledgebase_id",
		"tenant_id",
		"name",
		"chat_prompts",
		"community_prompts",
		"service_prompts",
		"survey_prompts",
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

	varKnowledgebase := _Knowledgebase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKnowledgebase)

	if err != nil {
		return err
	}

	*o = Knowledgebase(varKnowledgebase)

	return err
}

type NullableKnowledgebase struct {
	value *Knowledgebase
	isSet bool
}

func (v NullableKnowledgebase) Get() *Knowledgebase {
	return v.value
}

func (v *NullableKnowledgebase) Set(val *Knowledgebase) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgebase) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgebase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgebase(val *Knowledgebase) *NullableKnowledgebase {
	return &NullableKnowledgebase{value: val, isSet: true}
}

func (v NullableKnowledgebase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgebase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


