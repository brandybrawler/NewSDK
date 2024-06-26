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

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price struct for Price
type Price struct {
	Quota Quota `json:"quota"`
	Amount float64 `json:"amount"`
	CompanyType CompanyType `json:"company_type"`
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(quota Quota, amount float64, companyType CompanyType) *Price {
	this := Price{}
	this.Quota = quota
	this.Amount = amount
	this.CompanyType = companyType
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetQuota returns the Quota field value
func (o *Price) GetQuota() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
func (o *Price) GetQuotaOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *Price) SetQuota(v Quota) {
	o.Quota = v
}

// GetAmount returns the Amount field value
func (o *Price) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Price) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Price) SetAmount(v float64) {
	o.Amount = v
}

// GetCompanyType returns the CompanyType field value
func (o *Price) GetCompanyType() CompanyType {
	if o == nil {
		var ret CompanyType
		return ret
	}

	return o.CompanyType
}

// GetCompanyTypeOk returns a tuple with the CompanyType field value
// and a boolean to check if the value has been set.
func (o *Price) GetCompanyTypeOk() (*CompanyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyType, true
}

// SetCompanyType sets field value
func (o *Price) SetCompanyType(v CompanyType) {
	o.CompanyType = v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quota"] = o.Quota
	toSerialize["amount"] = o.Amount
	toSerialize["company_type"] = o.CompanyType
	return toSerialize, nil
}

func (o *Price) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quota",
		"amount",
		"company_type",
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

	varPrice := _Price{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrice)

	if err != nil {
		return err
	}

	*o = Price(varPrice)

	return err
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


