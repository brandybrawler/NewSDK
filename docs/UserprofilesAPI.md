# \UserprofilesAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserprofilesAdminPartialUpdate**](UserprofilesAPI.md#UserprofilesAdminPartialUpdate) | **Patch** /api/userprofiles/admin/ | 
[**UserprofilesAdminRetrieve**](UserprofilesAPI.md#UserprofilesAdminRetrieve) | **Get** /api/userprofiles/admin/ | 
[**UserprofilesAdminUpdate**](UserprofilesAPI.md#UserprofilesAdminUpdate) | **Put** /api/userprofiles/admin/ | 
[**UserprofilesAdminprofilesRetrieve**](UserprofilesAPI.md#UserprofilesAdminprofilesRetrieve) | **Get** /api/userprofiles/adminprofiles/{admin_id}/ | 
[**UserprofilesClientPartialUpdate**](UserprofilesAPI.md#UserprofilesClientPartialUpdate) | **Patch** /api/userprofiles/client/ | 
[**UserprofilesClientRetrieve**](UserprofilesAPI.md#UserprofilesClientRetrieve) | **Get** /api/userprofiles/client/ | 
[**UserprofilesClientUpdate**](UserprofilesAPI.md#UserprofilesClientUpdate) | **Put** /api/userprofiles/client/ | 
[**UserprofilesClientprofilesRetrieve**](UserprofilesAPI.md#UserprofilesClientprofilesRetrieve) | **Get** /api/userprofiles/clientprofiles/{client_id}/ | 
[**UserprofilesEmployeePartialUpdate**](UserprofilesAPI.md#UserprofilesEmployeePartialUpdate) | **Patch** /api/userprofiles/employee/ | 
[**UserprofilesEmployeeRetrieve**](UserprofilesAPI.md#UserprofilesEmployeeRetrieve) | **Get** /api/userprofiles/employee/ | 
[**UserprofilesEmployeeUpdate**](UserprofilesAPI.md#UserprofilesEmployeeUpdate) | **Put** /api/userprofiles/employee/ | 
[**UserprofilesEmployeeprofilesRetrieve**](UserprofilesAPI.md#UserprofilesEmployeeprofilesRetrieve) | **Get** /api/userprofiles/employeeprofiles/{employee_id}/ | 



## UserprofilesAdminPartialUpdate

> AdminUser UserprofilesAdminPartialUpdate(ctx).PatchedAdminUser(patchedAdminUser).Execute()



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
	patchedAdminUser := *openapiclient.NewPatchedAdminUser() // PatchedAdminUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesAdminPartialUpdate(context.Background()).PatchedAdminUser(patchedAdminUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesAdminPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesAdminPartialUpdate`: AdminUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesAdminPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesAdminPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedAdminUser** | [**PatchedAdminUser**](PatchedAdminUser.md) |  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesAdminRetrieve

> AdminUser UserprofilesAdminRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesAdminRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesAdminRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesAdminRetrieve`: AdminUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesAdminRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesAdminRetrieveRequest struct via the builder pattern


### Return type

[**AdminUser**](AdminUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesAdminUpdate

> AdminUser UserprofilesAdminUpdate(ctx).AdminUser(adminUser).Execute()



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
	adminUser := *openapiclient.NewAdminUser(int32(123), "FirstName_example", "LastName_example", "Username_example", "Email_example", "Password_example", "Token_example", *openapiclient.NewAdminProfile("Username_example", "ProfilePhoto_example"), "ProfilePhoto_example", "Country_example", "County_example", "City_example", "PostalCode_example", "Location_example") // AdminUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesAdminUpdate(context.Background()).AdminUser(adminUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesAdminUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesAdminUpdate`: AdminUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesAdminUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesAdminUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminUser** | [**AdminUser**](AdminUser.md) |  | 

### Return type

[**AdminUser**](AdminUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesAdminprofilesRetrieve

> AdminProfile UserprofilesAdminprofilesRetrieve(ctx, adminId).Execute()



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
	adminId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesAdminprofilesRetrieve(context.Background(), adminId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesAdminprofilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesAdminprofilesRetrieve`: AdminProfile
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesAdminprofilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adminId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesAdminprofilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminProfile**](AdminProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesClientPartialUpdate

> ClientUser UserprofilesClientPartialUpdate(ctx).PatchedClientUser(patchedClientUser).Execute()



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
	patchedClientUser := *openapiclient.NewPatchedClientUser() // PatchedClientUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesClientPartialUpdate(context.Background()).PatchedClientUser(patchedClientUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesClientPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesClientPartialUpdate`: ClientUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesClientPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesClientPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedClientUser** | [**PatchedClientUser**](PatchedClientUser.md) |  | 

### Return type

[**ClientUser**](ClientUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesClientRetrieve

> ClientUser UserprofilesClientRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesClientRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesClientRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesClientRetrieve`: ClientUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesClientRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesClientRetrieveRequest struct via the builder pattern


### Return type

[**ClientUser**](ClientUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesClientUpdate

> ClientUser UserprofilesClientUpdate(ctx).ClientUser(clientUser).Execute()



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
	clientUser := *openapiclient.NewClientUser(int32(123), "FirstName_example", "LastName_example", "Username_example", "Email_example", "Password_example", "Token_example", *openapiclient.NewClientProfile("Username_example", "ProfilePhoto_example"), "ProfilePhoto_example", "Country_example", "County_example", "City_example", "PostalCode_example", "Location_example") // ClientUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesClientUpdate(context.Background()).ClientUser(clientUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesClientUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesClientUpdate`: ClientUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesClientUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesClientUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientUser** | [**ClientUser**](ClientUser.md) |  | 

### Return type

[**ClientUser**](ClientUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesClientprofilesRetrieve

> ClientProfile UserprofilesClientprofilesRetrieve(ctx, clientId).Execute()



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
	clientId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesClientprofilesRetrieve(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesClientprofilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesClientprofilesRetrieve`: ClientProfile
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesClientprofilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesClientprofilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientProfile**](ClientProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesEmployeePartialUpdate

> EmployeeUser UserprofilesEmployeePartialUpdate(ctx).PatchedEmployeeUser(patchedEmployeeUser).Execute()



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
	patchedEmployeeUser := *openapiclient.NewPatchedEmployeeUser() // PatchedEmployeeUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesEmployeePartialUpdate(context.Background()).PatchedEmployeeUser(patchedEmployeeUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesEmployeePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesEmployeePartialUpdate`: EmployeeUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesEmployeePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesEmployeePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedEmployeeUser** | [**PatchedEmployeeUser**](PatchedEmployeeUser.md) |  | 

### Return type

[**EmployeeUser**](EmployeeUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesEmployeeRetrieve

> EmployeeUser UserprofilesEmployeeRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesEmployeeRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesEmployeeRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesEmployeeRetrieve`: EmployeeUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesEmployeeRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesEmployeeRetrieveRequest struct via the builder pattern


### Return type

[**EmployeeUser**](EmployeeUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesEmployeeUpdate

> EmployeeUser UserprofilesEmployeeUpdate(ctx).EmployeeUser(employeeUser).Execute()



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
	employeeUser := *openapiclient.NewEmployeeUser(int32(123), "FirstName_example", "LastName_example", "Username_example", "Email_example", "Password_example", "Token_example", *openapiclient.NewEmployeeProfile("Username_example", "ProfilePhoto_example"), "ProfilePhoto_example", "Country_example", "County_example", "City_example", "PostalCode_example", "Location_example") // EmployeeUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesEmployeeUpdate(context.Background()).EmployeeUser(employeeUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesEmployeeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesEmployeeUpdate`: EmployeeUser
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesEmployeeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesEmployeeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employeeUser** | [**EmployeeUser**](EmployeeUser.md) |  | 

### Return type

[**EmployeeUser**](EmployeeUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserprofilesEmployeeprofilesRetrieve

> EmployeeProfile UserprofilesEmployeeprofilesRetrieve(ctx, employeeId).Execute()



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
	employeeId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserprofilesAPI.UserprofilesEmployeeprofilesRetrieve(context.Background(), employeeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserprofilesAPI.UserprofilesEmployeeprofilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserprofilesEmployeeprofilesRetrieve`: EmployeeProfile
	fmt.Fprintf(os.Stdout, "Response from `UserprofilesAPI.UserprofilesEmployeeprofilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserprofilesEmployeeprofilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmployeeProfile**](EmployeeProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

