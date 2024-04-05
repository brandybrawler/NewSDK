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

// checks if the UserLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLogin{}

// UserLogin struct for UserLogin
type UserLogin struct {
	Id string `json:"id"`
	TenantId Tenant `json:"tenant_id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	UserType UserTypeFc6Enum `json:"user_type"`
}

type _UserLogin UserLogin

// NewUserLogin instantiates a new UserLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLogin(id string, tenantId Tenant, username string, firstName string, lastName string, email string, password string, token string, userType UserTypeFc6Enum) *UserLogin {
	this := UserLogin{}
	this.Id = id
	this.TenantId = tenantId
	this.Username = username
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Password = password
	this.Token = token
	this.UserType = userType
	return &this
}

// NewUserLoginWithDefaults instantiates a new UserLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginWithDefaults() *UserLogin {
	this := UserLogin{}
	return &this
}

// GetId returns the Id field value
func (o *UserLogin) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserLogin) SetId(v string) {
	o.Id = v
}

// GetTenantId returns the TenantId field value
func (o *UserLogin) GetTenantId() Tenant {
	if o == nil {
		var ret Tenant
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetTenantIdOk() (*Tenant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *UserLogin) SetTenantId(v Tenant) {
	o.TenantId = v
}

// GetUsername returns the Username field value
func (o *UserLogin) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserLogin) SetUsername(v string) {
	o.Username = v
}

// GetFirstName returns the FirstName field value
func (o *UserLogin) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserLogin) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserLogin) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserLogin) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserLogin) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserLogin) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UserLogin) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserLogin) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *UserLogin) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserLogin) SetToken(v string) {
	o.Token = v
}

// GetUserType returns the UserType field value
func (o *UserLogin) GetUserType() UserTypeFc6Enum {
	if o == nil {
		var ret UserTypeFc6Enum
		return ret
	}

	return o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetUserTypeOk() (*UserTypeFc6Enum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserType, true
}

// SetUserType sets field value
func (o *UserLogin) SetUserType(v UserTypeFc6Enum) {
	o.UserType = v
}

func (o UserLogin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["username"] = o.Username
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	toSerialize["token"] = o.Token
	toSerialize["user_type"] = o.UserType
	return toSerialize, nil
}

func (o *UserLogin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tenant_id",
		"username",
		"first_name",
		"last_name",
		"email",
		"password",
		"token",
		"user_type",
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

	varUserLogin := _UserLogin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserLogin)

	if err != nil {
		return err
	}

	*o = UserLogin(varUserLogin)

	return err
}

type NullableUserLogin struct {
	value *UserLogin
	isSet bool
}

func (v NullableUserLogin) Get() *UserLogin {
	return v.value
}

func (v *NullableUserLogin) Set(val *UserLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLogin(val *UserLogin) *NullableUserLogin {
	return &NullableUserLogin{value: val, isSet: true}
}

func (v NullableUserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

