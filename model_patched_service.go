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

// checks if the PatchedService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedService{}

// PatchedService struct for PatchedService
type PatchedService struct {
	ServiceId *int32 `json:"service_id,omitempty"`
	// The tenant chats ID.
	TenantId *int32 `json:"tenant_id,omitempty"`
	NameOfService *string `json:"name_of_service,omitempty"`
	ServiceDescription *string `json:"service_description,omitempty"`
	ServiceType *int32 `json:"service_type,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	ResolutionPeriod *string `json:"resolution_period,omitempty"`
	ServiceAvailability *ServiceAvailabilityEnum `json:"service_availability,omitempty"`
	AiPowered *bool `json:"ai_powered,omitempty"`
	AiSolved NullableBool `json:"ai_solved,omitempty"`
	Details interface{} `json:"details,omitempty"`
	ServiceActionPlan interface{} `json:"service_action_plan,omitempty"`
}

// NewPatchedService instantiates a new PatchedService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedService() *PatchedService {
	this := PatchedService{}
	return &this
}

// NewPatchedServiceWithDefaults instantiates a new PatchedService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedServiceWithDefaults() *PatchedService {
	this := PatchedService{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PatchedService) GetServiceId() int32 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int32
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetServiceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PatchedService) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int32 and assigns it to the ServiceId field.
func (o *PatchedService) SetServiceId(v int32) {
	o.ServiceId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PatchedService) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PatchedService) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *PatchedService) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetNameOfService returns the NameOfService field value if set, zero value otherwise.
func (o *PatchedService) GetNameOfService() string {
	if o == nil || IsNil(o.NameOfService) {
		var ret string
		return ret
	}
	return *o.NameOfService
}

// GetNameOfServiceOk returns a tuple with the NameOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetNameOfServiceOk() (*string, bool) {
	if o == nil || IsNil(o.NameOfService) {
		return nil, false
	}
	return o.NameOfService, true
}

// HasNameOfService returns a boolean if a field has been set.
func (o *PatchedService) HasNameOfService() bool {
	if o != nil && !IsNil(o.NameOfService) {
		return true
	}

	return false
}

// SetNameOfService gets a reference to the given string and assigns it to the NameOfService field.
func (o *PatchedService) SetNameOfService(v string) {
	o.NameOfService = &v
}

// GetServiceDescription returns the ServiceDescription field value if set, zero value otherwise.
func (o *PatchedService) GetServiceDescription() string {
	if o == nil || IsNil(o.ServiceDescription) {
		var ret string
		return ret
	}
	return *o.ServiceDescription
}

// GetServiceDescriptionOk returns a tuple with the ServiceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetServiceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDescription) {
		return nil, false
	}
	return o.ServiceDescription, true
}

// HasServiceDescription returns a boolean if a field has been set.
func (o *PatchedService) HasServiceDescription() bool {
	if o != nil && !IsNil(o.ServiceDescription) {
		return true
	}

	return false
}

// SetServiceDescription gets a reference to the given string and assigns it to the ServiceDescription field.
func (o *PatchedService) SetServiceDescription(v string) {
	o.ServiceDescription = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *PatchedService) GetServiceType() int32 {
	if o == nil || IsNil(o.ServiceType) {
		var ret int32
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetServiceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *PatchedService) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given int32 and assigns it to the ServiceType field.
func (o *PatchedService) SetServiceType(v int32) {
	o.ServiceType = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedService) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedService) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedService) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetResolutionPeriod returns the ResolutionPeriod field value if set, zero value otherwise.
func (o *PatchedService) GetResolutionPeriod() string {
	if o == nil || IsNil(o.ResolutionPeriod) {
		var ret string
		return ret
	}
	return *o.ResolutionPeriod
}

// GetResolutionPeriodOk returns a tuple with the ResolutionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetResolutionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ResolutionPeriod) {
		return nil, false
	}
	return o.ResolutionPeriod, true
}

// HasResolutionPeriod returns a boolean if a field has been set.
func (o *PatchedService) HasResolutionPeriod() bool {
	if o != nil && !IsNil(o.ResolutionPeriod) {
		return true
	}

	return false
}

// SetResolutionPeriod gets a reference to the given string and assigns it to the ResolutionPeriod field.
func (o *PatchedService) SetResolutionPeriod(v string) {
	o.ResolutionPeriod = &v
}

// GetServiceAvailability returns the ServiceAvailability field value if set, zero value otherwise.
func (o *PatchedService) GetServiceAvailability() ServiceAvailabilityEnum {
	if o == nil || IsNil(o.ServiceAvailability) {
		var ret ServiceAvailabilityEnum
		return ret
	}
	return *o.ServiceAvailability
}

// GetServiceAvailabilityOk returns a tuple with the ServiceAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetServiceAvailabilityOk() (*ServiceAvailabilityEnum, bool) {
	if o == nil || IsNil(o.ServiceAvailability) {
		return nil, false
	}
	return o.ServiceAvailability, true
}

// HasServiceAvailability returns a boolean if a field has been set.
func (o *PatchedService) HasServiceAvailability() bool {
	if o != nil && !IsNil(o.ServiceAvailability) {
		return true
	}

	return false
}

// SetServiceAvailability gets a reference to the given ServiceAvailabilityEnum and assigns it to the ServiceAvailability field.
func (o *PatchedService) SetServiceAvailability(v ServiceAvailabilityEnum) {
	o.ServiceAvailability = &v
}

// GetAiPowered returns the AiPowered field value if set, zero value otherwise.
func (o *PatchedService) GetAiPowered() bool {
	if o == nil || IsNil(o.AiPowered) {
		var ret bool
		return ret
	}
	return *o.AiPowered
}

// GetAiPoweredOk returns a tuple with the AiPowered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedService) GetAiPoweredOk() (*bool, bool) {
	if o == nil || IsNil(o.AiPowered) {
		return nil, false
	}
	return o.AiPowered, true
}

// HasAiPowered returns a boolean if a field has been set.
func (o *PatchedService) HasAiPowered() bool {
	if o != nil && !IsNil(o.AiPowered) {
		return true
	}

	return false
}

// SetAiPowered gets a reference to the given bool and assigns it to the AiPowered field.
func (o *PatchedService) SetAiPowered(v bool) {
	o.AiPowered = &v
}

// GetAiSolved returns the AiSolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedService) GetAiSolved() bool {
	if o == nil || IsNil(o.AiSolved.Get()) {
		var ret bool
		return ret
	}
	return *o.AiSolved.Get()
}

// GetAiSolvedOk returns a tuple with the AiSolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedService) GetAiSolvedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AiSolved.Get(), o.AiSolved.IsSet()
}

// HasAiSolved returns a boolean if a field has been set.
func (o *PatchedService) HasAiSolved() bool {
	if o != nil && o.AiSolved.IsSet() {
		return true
	}

	return false
}

// SetAiSolved gets a reference to the given NullableBool and assigns it to the AiSolved field.
func (o *PatchedService) SetAiSolved(v bool) {
	o.AiSolved.Set(&v)
}
// SetAiSolvedNil sets the value for AiSolved to be an explicit nil
func (o *PatchedService) SetAiSolvedNil() {
	o.AiSolved.Set(nil)
}

// UnsetAiSolved ensures that no value is present for AiSolved, not even an explicit nil
func (o *PatchedService) UnsetAiSolved() {
	o.AiSolved.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedService) GetDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedService) GetDetailsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PatchedService) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given interface{} and assigns it to the Details field.
func (o *PatchedService) SetDetails(v interface{}) {
	o.Details = v
}

// GetServiceActionPlan returns the ServiceActionPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedService) GetServiceActionPlan() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ServiceActionPlan
}

// GetServiceActionPlanOk returns a tuple with the ServiceActionPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedService) GetServiceActionPlanOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ServiceActionPlan) {
		return nil, false
	}
	return &o.ServiceActionPlan, true
}

// HasServiceActionPlan returns a boolean if a field has been set.
func (o *PatchedService) HasServiceActionPlan() bool {
	if o != nil && !IsNil(o.ServiceActionPlan) {
		return true
	}

	return false
}

// SetServiceActionPlan gets a reference to the given interface{} and assigns it to the ServiceActionPlan field.
func (o *PatchedService) SetServiceActionPlan(v interface{}) {
	o.ServiceActionPlan = v
}

func (o PatchedService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.NameOfService) {
		toSerialize["name_of_service"] = o.NameOfService
	}
	if !IsNil(o.ServiceDescription) {
		toSerialize["service_description"] = o.ServiceDescription
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.ResolutionPeriod) {
		toSerialize["resolution_period"] = o.ResolutionPeriod
	}
	if !IsNil(o.ServiceAvailability) {
		toSerialize["service_availability"] = o.ServiceAvailability
	}
	if !IsNil(o.AiPowered) {
		toSerialize["ai_powered"] = o.AiPowered
	}
	if o.AiSolved.IsSet() {
		toSerialize["ai_solved"] = o.AiSolved.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ServiceActionPlan != nil {
		toSerialize["service_action_plan"] = o.ServiceActionPlan
	}
	return toSerialize, nil
}

type NullablePatchedService struct {
	value *PatchedService
	isSet bool
}

func (v NullablePatchedService) Get() *PatchedService {
	return v.value
}

func (v *NullablePatchedService) Set(val *PatchedService) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedService) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedService(val *PatchedService) *NullablePatchedService {
	return &NullablePatchedService{value: val, isSet: true}
}

func (v NullablePatchedService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


