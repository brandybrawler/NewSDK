# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | [**Quota**](Quota.md) |  | [readonly] 
**Amount** | **float64** |  | 
**CompanyType** | [**CompanyType**](CompanyType.md) |  | [readonly] 

## Methods

### NewPrice

`func NewPrice(quota Quota, amount float64, companyType CompanyType, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *Price) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Price) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Price) SetQuota(v Quota)`

SetQuota sets Quota field to given value.


### GetAmount

`func (o *Price) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Price) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Price) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCompanyType

`func (o *Price) GetCompanyType() CompanyType`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *Price) GetCompanyTypeOk() (*CompanyType, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *Price) SetCompanyType(v CompanyType)`

SetCompanyType sets CompanyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


