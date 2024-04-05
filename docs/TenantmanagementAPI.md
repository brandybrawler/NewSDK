# \TenantmanagementAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantmanagementAddressCreate**](TenantmanagementAPI.md#TenantmanagementAddressCreate) | **Post** /api/tenantmanagement/address/ | 
[**TenantmanagementAddressRetrieve**](TenantmanagementAPI.md#TenantmanagementAddressRetrieve) | **Get** /api/tenantmanagement/address/ | 
[**TenantmanagementDepartmentCreate**](TenantmanagementAPI.md#TenantmanagementDepartmentCreate) | **Post** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentDestroy**](TenantmanagementAPI.md#TenantmanagementDepartmentDestroy) | **Delete** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentPartialUpdate**](TenantmanagementAPI.md#TenantmanagementDepartmentPartialUpdate) | **Patch** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentRetrieve**](TenantmanagementAPI.md#TenantmanagementDepartmentRetrieve) | **Get** /api/tenantmanagement/department | 
[**TenantmanagementMetadataCreate**](TenantmanagementAPI.md#TenantmanagementMetadataCreate) | **Post** /api/tenantmanagement/metadata/ | 
[**TenantmanagementMetadataDestroy**](TenantmanagementAPI.md#TenantmanagementMetadataDestroy) | **Delete** /api/tenantmanagement/metadata/ | 
[**TenantmanagementMetadataRetrieve**](TenantmanagementAPI.md#TenantmanagementMetadataRetrieve) | **Get** /api/tenantmanagement/metadata/ | 
[**TenantmanagementProductCreate**](TenantmanagementAPI.md#TenantmanagementProductCreate) | **Post** /api/tenantmanagement/product/ | 
[**TenantmanagementProductDestroy**](TenantmanagementAPI.md#TenantmanagementProductDestroy) | **Delete** /api/tenantmanagement/product/ | 
[**TenantmanagementProductRetrieve**](TenantmanagementAPI.md#TenantmanagementProductRetrieve) | **Get** /api/tenantmanagement/product/ | 
[**TenantmanagementTenantCreate**](TenantmanagementAPI.md#TenantmanagementTenantCreate) | **Post** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantDestroy**](TenantmanagementAPI.md#TenantmanagementTenantDestroy) | **Delete** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantRetrieve**](TenantmanagementAPI.md#TenantmanagementTenantRetrieve) | **Get** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantdetailsRetrieve**](TenantmanagementAPI.md#TenantmanagementTenantdetailsRetrieve) | **Get** /api/tenantmanagement/tenantdetails/ | 
[**TenantmanagementVirtualAssistantsCreate**](TenantmanagementAPI.md#TenantmanagementVirtualAssistantsCreate) | **Post** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualAssistantsRetrieve**](TenantmanagementAPI.md#TenantmanagementVirtualAssistantsRetrieve) | **Get** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualAssistantsUpdate**](TenantmanagementAPI.md#TenantmanagementVirtualAssistantsUpdate) | **Put** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualassistantCreate**](TenantmanagementAPI.md#TenantmanagementVirtualassistantCreate) | **Post** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassistantRetrieve**](TenantmanagementAPI.md#TenantmanagementVirtualassistantRetrieve) | **Get** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassistantUpdate**](TenantmanagementAPI.md#TenantmanagementVirtualassistantUpdate) | **Put** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassitantdocsCreate**](TenantmanagementAPI.md#TenantmanagementVirtualassitantdocsCreate) | **Post** /api/tenantmanagement/virtualassitantdocs/ | 
[**TenantmanagementVirtualassitantdocsRetrieve**](TenantmanagementAPI.md#TenantmanagementVirtualassitantdocsRetrieve) | **Get** /api/tenantmanagement/virtualassitantdocs/ | 



## TenantmanagementAddressCreate

> Address TenantmanagementAddressCreate(ctx).Address(address).Execute()





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
	address := *openapiclient.NewAddress(int32(123), int32(123), "City_example", "Country_example", "PostalCode_example", "State_example", "PaymentNumber_example") // Address | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementAddressCreate(context.Background()).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementAddressCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementAddressCreate`: Address
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementAddressCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementAddressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | [**Address**](Address.md) |  | 

### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementAddressRetrieve

> Address TenantmanagementAddressRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementAddressRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementAddressRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementAddressRetrieve`: Address
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementAddressRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementAddressRetrieveRequest struct via the builder pattern


### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentCreate

> Department TenantmanagementDepartmentCreate(ctx).Department(department).Execute()





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
	department := *openapiclient.NewDepartment(int32(123), int32(123), "Name_example", []int32{int32(123)}, []int32{int32(123)}) // Department | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementDepartmentCreate(context.Background()).Department(department).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementDepartmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementDepartmentCreate`: Department
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementDepartmentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **department** | [**Department**](Department.md) |  | 

### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentDestroy

> TenantmanagementDepartmentDestroy(ctx).Execute()





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
	r, err := apiClient.TenantmanagementAPI.TenantmanagementDepartmentDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementDepartmentDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentDestroyRequest struct via the builder pattern


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


## TenantmanagementDepartmentPartialUpdate

> Department TenantmanagementDepartmentPartialUpdate(ctx).PatchedDepartment(patchedDepartment).Execute()





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
	patchedDepartment := *openapiclient.NewPatchedDepartment() // PatchedDepartment |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementDepartmentPartialUpdate(context.Background()).PatchedDepartment(patchedDepartment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementDepartmentPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementDepartmentPartialUpdate`: Department
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementDepartmentPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedDepartment** | [**PatchedDepartment**](PatchedDepartment.md) |  | 

### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentRetrieve

> Department TenantmanagementDepartmentRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementDepartmentRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementDepartmentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementDepartmentRetrieve`: Department
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementDepartmentRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentRetrieveRequest struct via the builder pattern


### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementMetadataCreate

> Metadata TenantmanagementMetadataCreate(ctx).Metadata(metadata).Execute()





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
	metadata := *openapiclient.NewMetadata(int32(123), interface{}(123)) // Metadata | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementMetadataCreate(context.Background()).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementMetadataCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementMetadataCreate`: Metadata
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementMetadataCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**Metadata**](Metadata.md) |  | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementMetadataDestroy

> TenantmanagementMetadataDestroy(ctx).Execute()





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
	r, err := apiClient.TenantmanagementAPI.TenantmanagementMetadataDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementMetadataDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataDestroyRequest struct via the builder pattern


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


## TenantmanagementMetadataRetrieve

> Metadata TenantmanagementMetadataRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementMetadataRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementMetadataRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementMetadataRetrieve`: Metadata
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementMetadataRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataRetrieveRequest struct via the builder pattern


### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementProductCreate

> Product TenantmanagementProductCreate(ctx).Product(product).Execute()





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
	product := *openapiclient.NewProduct(int32(123), int32(123), "Name_example", "Description_example", "Price_example") // Product | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementProductCreate(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementProductCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementProductCreate`: Product
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementProductCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementProductDestroy

> TenantmanagementProductDestroy(ctx).Execute()





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
	r, err := apiClient.TenantmanagementAPI.TenantmanagementProductDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementProductDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductDestroyRequest struct via the builder pattern


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


## TenantmanagementProductRetrieve

> Product TenantmanagementProductRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementProductRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementProductRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementProductRetrieve`: Product
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementProductRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductRetrieveRequest struct via the builder pattern


### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantCreate

> Tenant TenantmanagementTenantCreate(ctx).Tenant(tenant).Execute()





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
	tenant := *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")) // Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementTenantCreate(context.Background()).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementTenantCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementTenantCreate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementTenantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantDestroy

> TenantmanagementTenantDestroy(ctx).Execute()





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
	r, err := apiClient.TenantmanagementAPI.TenantmanagementTenantDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementTenantDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantDestroyRequest struct via the builder pattern


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


## TenantmanagementTenantRetrieve

> Tenant TenantmanagementTenantRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementTenantRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementTenantRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementTenantRetrieve`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementTenantRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantRetrieveRequest struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantdetailsRetrieve

> TenantmanagementTenantdetailsRetrieve(ctx).Execute()



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
	r, err := apiClient.TenantmanagementAPI.TenantmanagementTenantdetailsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementTenantdetailsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantdetailsRetrieveRequest struct via the builder pattern


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


## TenantmanagementVirtualAssistantsCreate

> VirtualAssistant TenantmanagementVirtualAssistantsCreate(ctx, id).VirtualAssistant(virtualAssistant).Execute()



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
	id := int32(56) // int32 | 
	virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualAssistantsCreate(context.Background(), id).VirtualAssistant(virtualAssistant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualAssistantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualAssistantsCreate`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualAssistantsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualAssistantsRetrieve

> VirtualAssistant TenantmanagementVirtualAssistantsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualAssistantsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualAssistantsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualAssistantsRetrieve`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualAssistantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualAssistantsUpdate

> VirtualAssistant TenantmanagementVirtualAssistantsUpdate(ctx, id).VirtualAssistant(virtualAssistant).Execute()



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
	id := int32(56) // int32 | 
	virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualAssistantsUpdate(context.Background(), id).VirtualAssistant(virtualAssistant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualAssistantsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualAssistantsUpdate`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualAssistantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantCreate

> VirtualAssistant TenantmanagementVirtualassistantCreate(ctx).VirtualAssistant(virtualAssistant).Execute()



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
	virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualassistantCreate(context.Background()).VirtualAssistant(virtualAssistant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualassistantCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualassistantCreate`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualassistantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantRetrieve

> VirtualAssistant TenantmanagementVirtualassistantRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualassistantRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualassistantRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualassistantRetrieve`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualassistantRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantRetrieveRequest struct via the builder pattern


### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantUpdate

> VirtualAssistant TenantmanagementVirtualassistantUpdate(ctx).VirtualAssistant(virtualAssistant).Execute()



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
	virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualassistantUpdate(context.Background()).VirtualAssistant(virtualAssistant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualassistantUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualassistantUpdate`: VirtualAssistant
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualassistantUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassitantdocsCreate

> VirtualAssistantDocuments TenantmanagementVirtualassitantdocsCreate(ctx).VirtualAssistantDocuments(virtualAssistantDocuments).Execute()



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
	virtualAssistantDocuments := *openapiclient.NewVirtualAssistantDocuments(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "Pdf_example", int32(123)) // VirtualAssistantDocuments | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualassitantdocsCreate(context.Background()).VirtualAssistantDocuments(virtualAssistantDocuments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualassitantdocsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualassitantdocsCreate`: VirtualAssistantDocuments
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualassitantdocsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassitantdocsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistantDocuments** | [**VirtualAssistantDocuments**](VirtualAssistantDocuments.md) |  | 

### Return type

[**VirtualAssistantDocuments**](VirtualAssistantDocuments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassitantdocsRetrieve

> VirtualAssistantDocuments TenantmanagementVirtualassitantdocsRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.TenantmanagementAPI.TenantmanagementVirtualassitantdocsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementAPI.TenantmanagementVirtualassitantdocsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantmanagementVirtualassitantdocsRetrieve`: VirtualAssistantDocuments
	fmt.Fprintf(os.Stdout, "Response from `TenantmanagementAPI.TenantmanagementVirtualassitantdocsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassitantdocsRetrieveRequest struct via the builder pattern


### Return type

[**VirtualAssistantDocuments**](VirtualAssistantDocuments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

