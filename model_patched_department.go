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

// checks if the PatchedDepartment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDepartment{}

// PatchedDepartment struct for PatchedDepartment
type PatchedDepartment struct {
	// The department ID.
	DepartmentId *int32 `json:"department_id,omitempty"`
	// The tenant this department belongs to
	Tenant *int32 `json:"tenant,omitempty"`
	// Name of the department
	Name *string `json:"name,omitempty"`
	// Description of the department
	Description NullableString `json:"description,omitempty"`
	Employees []int32 `json:"employees,omitempty"`
	Admins []int32 `json:"admins,omitempty"`
}

// NewPatchedDepartment instantiates a new PatchedDepartment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDepartment() *PatchedDepartment {
	this := PatchedDepartment{}
	return &this
}

// NewPatchedDepartmentWithDefaults instantiates a new PatchedDepartment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDepartmentWithDefaults() *PatchedDepartment {
	this := PatchedDepartment{}
	return &this
}

// GetDepartmentId returns the DepartmentId field value if set, zero value otherwise.
func (o *PatchedDepartment) GetDepartmentId() int32 {
	if o == nil || IsNil(o.DepartmentId) {
		var ret int32
		return ret
	}
	return *o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDepartment) GetDepartmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DepartmentId) {
		return nil, false
	}
	return o.DepartmentId, true
}

// HasDepartmentId returns a boolean if a field has been set.
func (o *PatchedDepartment) HasDepartmentId() bool {
	if o != nil && !IsNil(o.DepartmentId) {
		return true
	}

	return false
}

// SetDepartmentId gets a reference to the given int32 and assigns it to the DepartmentId field.
func (o *PatchedDepartment) SetDepartmentId(v int32) {
	o.DepartmentId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PatchedDepartment) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant) {
		var ret int32
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDepartment) GetTenantOk() (*int32, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedDepartment) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given int32 and assigns it to the Tenant field.
func (o *PatchedDepartment) SetTenant(v int32) {
	o.Tenant = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDepartment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDepartment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDepartment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDepartment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDepartment) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDepartment) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDepartment) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedDepartment) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedDepartment) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedDepartment) UnsetDescription() {
	o.Description.Unset()
}

// GetEmployees returns the Employees field value if set, zero value otherwise.
func (o *PatchedDepartment) GetEmployees() []int32 {
	if o == nil || IsNil(o.Employees) {
		var ret []int32
		return ret
	}
	return o.Employees
}

// GetEmployeesOk returns a tuple with the Employees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDepartment) GetEmployeesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Employees) {
		return nil, false
	}
	return o.Employees, true
}

// HasEmployees returns a boolean if a field has been set.
func (o *PatchedDepartment) HasEmployees() bool {
	if o != nil && !IsNil(o.Employees) {
		return true
	}

	return false
}

// SetEmployees gets a reference to the given []int32 and assigns it to the Employees field.
func (o *PatchedDepartment) SetEmployees(v []int32) {
	o.Employees = v
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *PatchedDepartment) GetAdmins() []int32 {
	if o == nil || IsNil(o.Admins) {
		var ret []int32
		return ret
	}
	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDepartment) GetAdminsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Admins) {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *PatchedDepartment) HasAdmins() bool {
	if o != nil && !IsNil(o.Admins) {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given []int32 and assigns it to the Admins field.
func (o *PatchedDepartment) SetAdmins(v []int32) {
	o.Admins = v
}

func (o PatchedDepartment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDepartment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DepartmentId) {
		toSerialize["department_id"] = o.DepartmentId
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Employees) {
		toSerialize["employees"] = o.Employees
	}
	if !IsNil(o.Admins) {
		toSerialize["admins"] = o.Admins
	}
	return toSerialize, nil
}

type NullablePatchedDepartment struct {
	value *PatchedDepartment
	isSet bool
}

func (v NullablePatchedDepartment) Get() *PatchedDepartment {
	return v.value
}

func (v *NullablePatchedDepartment) Set(val *PatchedDepartment) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDepartment) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDepartment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDepartment(val *PatchedDepartment) *NullablePatchedDepartment {
	return &NullablePatchedDepartment{value: val, isSet: true}
}

func (v NullablePatchedDepartment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDepartment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


