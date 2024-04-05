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

// checks if the Employee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Employee{}

// Employee Serializers registration requests and creates a new user.
type Employee struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phonenumber NullableString `json:"phonenumber,omitempty"`
	Gender NullableAdminGender `json:"gender,omitempty"`
	DOB NullableString `json:"DOB,omitempty"`
	UserType NullableUserType8ceEnum `json:"user_type,omitempty"`
	// Display name of the tenant
	TenantId int32 `json:"tenant_id"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token string `json:"token"`
}

type _Employee Employee

// NewEmployee instantiates a new Employee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee(id int32, username string, email string, firstName string, lastName string, tenantId int32, password string, confirmPassword string, token string) *Employee {
	this := Employee{}
	this.Id = id
	this.Username = username
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.TenantId = tenantId
	this.Password = password
	this.ConfirmPassword = confirmPassword
	this.Token = token
	return &this
}

// NewEmployeeWithDefaults instantiates a new Employee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeWithDefaults() *Employee {
	this := Employee{}
	return &this
}

// GetId returns the Id field value
func (o *Employee) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Employee) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Employee) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *Employee) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Employee) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Employee) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *Employee) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Employee) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Employee) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *Employee) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *Employee) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *Employee) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *Employee) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *Employee) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *Employee) SetLastName(v string) {
	o.LastName = v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetPhonenumber() string {
	if o == nil || IsNil(o.Phonenumber.Get()) {
		var ret string
		return ret
	}
	return *o.Phonenumber.Get()
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetPhonenumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phonenumber.Get(), o.Phonenumber.IsSet()
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *Employee) HasPhonenumber() bool {
	if o != nil && o.Phonenumber.IsSet() {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given NullableString and assigns it to the Phonenumber field.
func (o *Employee) SetPhonenumber(v string) {
	o.Phonenumber.Set(&v)
}
// SetPhonenumberNil sets the value for Phonenumber to be an explicit nil
func (o *Employee) SetPhonenumberNil() {
	o.Phonenumber.Set(nil)
}

// UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
func (o *Employee) UnsetPhonenumber() {
	o.Phonenumber.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetGender() AdminGender {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret AdminGender
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetGenderOk() (*AdminGender, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *Employee) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableAdminGender and assigns it to the Gender field.
func (o *Employee) SetGender(v AdminGender) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *Employee) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *Employee) UnsetGender() {
	o.Gender.Unset()
}

// GetDOB returns the DOB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetDOB() string {
	if o == nil || IsNil(o.DOB.Get()) {
		var ret string
		return ret
	}
	return *o.DOB.Get()
}

// GetDOBOk returns a tuple with the DOB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetDOBOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DOB.Get(), o.DOB.IsSet()
}

// HasDOB returns a boolean if a field has been set.
func (o *Employee) HasDOB() bool {
	if o != nil && o.DOB.IsSet() {
		return true
	}

	return false
}

// SetDOB gets a reference to the given NullableString and assigns it to the DOB field.
func (o *Employee) SetDOB(v string) {
	o.DOB.Set(&v)
}
// SetDOBNil sets the value for DOB to be an explicit nil
func (o *Employee) SetDOBNil() {
	o.DOB.Set(nil)
}

// UnsetDOB ensures that no value is present for DOB, not even an explicit nil
func (o *Employee) UnsetDOB() {
	o.DOB.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetUserType() UserType8ceEnum {
	if o == nil || IsNil(o.UserType.Get()) {
		var ret UserType8ceEnum
		return ret
	}
	return *o.UserType.Get()
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetUserTypeOk() (*UserType8ceEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserType.Get(), o.UserType.IsSet()
}

// HasUserType returns a boolean if a field has been set.
func (o *Employee) HasUserType() bool {
	if o != nil && o.UserType.IsSet() {
		return true
	}

	return false
}

// SetUserType gets a reference to the given NullableUserType8ceEnum and assigns it to the UserType field.
func (o *Employee) SetUserType(v UserType8ceEnum) {
	o.UserType.Set(&v)
}
// SetUserTypeNil sets the value for UserType to be an explicit nil
func (o *Employee) SetUserTypeNil() {
	o.UserType.Set(nil)
}

// UnsetUserType ensures that no value is present for UserType, not even an explicit nil
func (o *Employee) UnsetUserType() {
	o.UserType.Unset()
}

// GetTenantId returns the TenantId field value
func (o *Employee) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Employee) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Employee) SetTenantId(v int32) {
	o.TenantId = v
}

// GetPassword returns the Password field value
func (o *Employee) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Employee) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *Employee) SetPassword(v string) {
	o.Password = v
}

// GetConfirmPassword returns the ConfirmPassword field value
func (o *Employee) GetConfirmPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmPassword
}

// GetConfirmPasswordOk returns a tuple with the ConfirmPassword field value
// and a boolean to check if the value has been set.
func (o *Employee) GetConfirmPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmPassword, true
}

// SetConfirmPassword sets field value
func (o *Employee) SetConfirmPassword(v string) {
	o.ConfirmPassword = v
}

// GetToken returns the Token field value
func (o *Employee) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Employee) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *Employee) SetToken(v string) {
	o.Token = v
}

func (o Employee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Employee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	if o.Phonenumber.IsSet() {
		toSerialize["phonenumber"] = o.Phonenumber.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.DOB.IsSet() {
		toSerialize["DOB"] = o.DOB.Get()
	}
	if o.UserType.IsSet() {
		toSerialize["user_type"] = o.UserType.Get()
	}
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["password"] = o.Password
	toSerialize["confirm_password"] = o.ConfirmPassword
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *Employee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"email",
		"first_name",
		"last_name",
		"tenant_id",
		"password",
		"confirm_password",
		"token",
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

	varEmployee := _Employee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployee)

	if err != nil {
		return err
	}

	*o = Employee(varEmployee)

	return err
}

type NullableEmployee struct {
	value *Employee
	isSet bool
}

func (v NullableEmployee) Get() *Employee {
	return v.value
}

func (v *NullableEmployee) Set(val *Employee) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee(val *Employee) *NullableEmployee {
	return &NullableEmployee{value: val, isSet: true}
}

func (v NullableEmployee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


