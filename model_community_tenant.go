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

// checks if the CommunityTenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityTenant{}

// CommunityTenant struct for CommunityTenant
type CommunityTenant struct {
	// The tenant chats ID.
	TenantId int32 `json:"tenant_id"`
	// name of the company
	TenantName string `json:"tenant_name"`
	// Industry of the company  * `IT` - Information Technology * `Finance` - Finance * `Healthcare` - Healthcare * `Education` - Education * `Retail` - Retail * `Manufacturing` - Manufacturing * `Automotive` - Automotive * `Hospitality` - Hospitality * `RealEstate` - Real Estate * `Media` - Media * `Telecommunications` - Telecommunications * `Energy` - Energy * `Transportation` - Transportation * `Agriculture` - Agriculture
	Industry IndustryEnum `json:"industry"`
	// Token for the tenant
	Token NullableString `json:"token,omitempty"`
	SubCategory NullableString `json:"sub_category,omitempty"`
	// The business registration number
	RegistrationNumber NullableString `json:"registration_number,omitempty"`
}

type _CommunityTenant CommunityTenant

// NewCommunityTenant instantiates a new CommunityTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityTenant(tenantId int32, tenantName string, industry IndustryEnum) *CommunityTenant {
	this := CommunityTenant{}
	this.TenantId = tenantId
	this.TenantName = tenantName
	this.Industry = industry
	return &this
}

// NewCommunityTenantWithDefaults instantiates a new CommunityTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityTenantWithDefaults() *CommunityTenant {
	this := CommunityTenant{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *CommunityTenant) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CommunityTenant) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CommunityTenant) SetTenantId(v int32) {
	o.TenantId = v
}

// GetTenantName returns the TenantName field value
func (o *CommunityTenant) GetTenantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value
// and a boolean to check if the value has been set.
func (o *CommunityTenant) GetTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantName, true
}

// SetTenantName sets field value
func (o *CommunityTenant) SetTenantName(v string) {
	o.TenantName = v
}

// GetIndustry returns the Industry field value
func (o *CommunityTenant) GetIndustry() IndustryEnum {
	if o == nil {
		var ret IndustryEnum
		return ret
	}

	return o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value
// and a boolean to check if the value has been set.
func (o *CommunityTenant) GetIndustryOk() (*IndustryEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Industry, true
}

// SetIndustry sets field value
func (o *CommunityTenant) SetIndustry(v IndustryEnum) {
	o.Industry = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityTenant) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityTenant) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *CommunityTenant) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *CommunityTenant) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *CommunityTenant) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *CommunityTenant) UnsetToken() {
	o.Token.Unset()
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityTenant) GetSubCategory() string {
	if o == nil || IsNil(o.SubCategory.Get()) {
		var ret string
		return ret
	}
	return *o.SubCategory.Get()
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityTenant) GetSubCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubCategory.Get(), o.SubCategory.IsSet()
}

// HasSubCategory returns a boolean if a field has been set.
func (o *CommunityTenant) HasSubCategory() bool {
	if o != nil && o.SubCategory.IsSet() {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given NullableString and assigns it to the SubCategory field.
func (o *CommunityTenant) SetSubCategory(v string) {
	o.SubCategory.Set(&v)
}
// SetSubCategoryNil sets the value for SubCategory to be an explicit nil
func (o *CommunityTenant) SetSubCategoryNil() {
	o.SubCategory.Set(nil)
}

// UnsetSubCategory ensures that no value is present for SubCategory, not even an explicit nil
func (o *CommunityTenant) UnsetSubCategory() {
	o.SubCategory.Unset()
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityTenant) GetRegistrationNumber() string {
	if o == nil || IsNil(o.RegistrationNumber.Get()) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber.Get()
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityTenant) GetRegistrationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationNumber.Get(), o.RegistrationNumber.IsSet()
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *CommunityTenant) HasRegistrationNumber() bool {
	if o != nil && o.RegistrationNumber.IsSet() {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given NullableString and assigns it to the RegistrationNumber field.
func (o *CommunityTenant) SetRegistrationNumber(v string) {
	o.RegistrationNumber.Set(&v)
}
// SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil
func (o *CommunityTenant) SetRegistrationNumberNil() {
	o.RegistrationNumber.Set(nil)
}

// UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil
func (o *CommunityTenant) UnsetRegistrationNumber() {
	o.RegistrationNumber.Unset()
}

func (o CommunityTenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityTenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["tenant_name"] = o.TenantName
	toSerialize["industry"] = o.Industry
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.SubCategory.IsSet() {
		toSerialize["sub_category"] = o.SubCategory.Get()
	}
	if o.RegistrationNumber.IsSet() {
		toSerialize["registration_number"] = o.RegistrationNumber.Get()
	}
	return toSerialize, nil
}

func (o *CommunityTenant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenant_id",
		"tenant_name",
		"industry",
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

	varCommunityTenant := _CommunityTenant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityTenant)

	if err != nil {
		return err
	}

	*o = CommunityTenant(varCommunityTenant)

	return err
}

type NullableCommunityTenant struct {
	value *CommunityTenant
	isSet bool
}

func (v NullableCommunityTenant) Get() *CommunityTenant {
	return v.value
}

func (v *NullableCommunityTenant) Set(val *CommunityTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityTenant(val *CommunityTenant) *NullableCommunityTenant {
	return &NullableCommunityTenant{value: val, isSet: true}
}

func (v NullableCommunityTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


