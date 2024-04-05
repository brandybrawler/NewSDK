# \PaymentsAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsFreetrialCreate**](PaymentsAPI.md#PaymentsFreetrialCreate) | **Post** /api/payments/freetrial/ | 
[**PaymentsGetpaymentsRetrieve**](PaymentsAPI.md#PaymentsGetpaymentsRetrieve) | **Get** /api/payments/getpayments/ | 
[**PaymentsSubscriptionRetrieve**](PaymentsAPI.md#PaymentsSubscriptionRetrieve) | **Get** /api/payments/subscription/ | 
[**PaymentsTabularpaymentsRetrieve**](PaymentsAPI.md#PaymentsTabularpaymentsRetrieve) | **Get** /api/payments/tabularpayments/ | 
[**PaymentsTiersList**](PaymentsAPI.md#PaymentsTiersList) | **Get** /api/payments/tiers/ | 
[**PaymentsUpdateSubscriptionCreate**](PaymentsAPI.md#PaymentsUpdateSubscriptionCreate) | **Post** /api/payments/update_subscription/{id}/ | 



## PaymentsFreetrialCreate

> FreeTrialEnroll PaymentsFreetrialCreate(ctx).FreeTrialEnroll(freeTrialEnroll).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	freeTrialEnroll := *openapiclient.NewFreeTrialEnroll("Email_example", int32(123)) // FreeTrialEnroll | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsFreetrialCreate(context.Background()).FreeTrialEnroll(freeTrialEnroll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsFreetrialCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsFreetrialCreate`: FreeTrialEnroll
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsFreetrialCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsFreetrialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeTrialEnroll** | [**FreeTrialEnroll**](FreeTrialEnroll.md) |  | 

### Return type

[**FreeTrialEnroll**](FreeTrialEnroll.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsGetpaymentsRetrieve

> Payment PaymentsGetpaymentsRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsGetpaymentsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsGetpaymentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsGetpaymentsRetrieve`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsGetpaymentsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsGetpaymentsRetrieveRequest struct via the builder pattern


### Return type

[**Payment**](Payment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsSubscriptionRetrieve

> Subscription PaymentsSubscriptionRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsSubscriptionRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsSubscriptionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsSubscriptionRetrieve`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsSubscriptionRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsSubscriptionRetrieveRequest struct via the builder pattern


### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsTabularpaymentsRetrieve

> PaymentsTabularpaymentsRetrieve(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentsAPI.PaymentsTabularpaymentsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsTabularpaymentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsTabularpaymentsRetrieveRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsTiersList

> PaginatedTierList PaymentsTiersList(ctx).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsTiersList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsTiersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsTiersList`: PaginatedTierList
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsTiersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsTiersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedTierList**](PaginatedTierList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsUpdateSubscriptionCreate

> Subscription PaymentsUpdateSubscriptionCreate(ctx, id).Subscription(subscription).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	subscription := *openapiclient.NewSubscription(int32(123), *openapiclient.NewPaymentsTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), *openapiclient.NewTier("Name_example", []openapiclient.Feature{*openapiclient.NewFeature("Name_example", "Description_example")}, []openapiclient.Price{*openapiclient.NewPrice(*openapiclient.NewQuota("Name_example", int32(123)), "Amount_example", *openapiclient.NewCompanyType(openapiclient.CompanyTypeTypeEnum("small")))}), *openapiclient.NewQuota("Name_example", int32(123)), "CreatedAt_example", time.Now(), time.Now(), time.Now()) // Subscription |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsUpdateSubscriptionCreate(context.Background(), id).Subscription(subscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsUpdateSubscriptionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsUpdateSubscriptionCreate`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsUpdateSubscriptionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsUpdateSubscriptionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscription** | [**Subscription**](Subscription.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

