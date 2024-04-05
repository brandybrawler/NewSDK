# \AuthAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAdminCreate**](AuthAPI.md#AuthAdminCreate) | **Post** /api/auth/admin/ | 
[**AuthAdminverificationconfirmationRetrieve**](AuthAPI.md#AuthAdminverificationconfirmationRetrieve) | **Get** /api/auth/adminverificationconfirmation/{uidb64}/ | 
[**AuthAdminverificationrequestCreate**](AuthAPI.md#AuthAdminverificationrequestCreate) | **Post** /api/auth/adminverificationrequest/ | 
[**AuthAnonymoususerCreate**](AuthAPI.md#AuthAnonymoususerCreate) | **Post** /api/auth/anonymoususer/ | 
[**AuthAnonymoususerDestroy**](AuthAPI.md#AuthAnonymoususerDestroy) | **Delete** /api/auth/anonymoususer/ | 
[**AuthAnonymoususerRetrieve**](AuthAPI.md#AuthAnonymoususerRetrieve) | **Get** /api/auth/anonymoususer/ | 
[**AuthAuthGoogleCreate**](AuthAPI.md#AuthAuthGoogleCreate) | **Post** /api/auth/auth/google/ | 
[**AuthClientCreate**](AuthAPI.md#AuthClientCreate) | **Post** /api/auth/client/ | 
[**AuthClientUsersSignupSigninCreate**](AuthAPI.md#AuthClientUsersSignupSigninCreate) | **Post** /api/auth/client_users_signup_signin/ | 
[**AuthClientverificationconfirmationRetrieve**](AuthAPI.md#AuthClientverificationconfirmationRetrieve) | **Get** /api/auth/clientverificationconfirmation/{uidb64}/ | 
[**AuthClientverificationrequestCreate**](AuthAPI.md#AuthClientverificationrequestCreate) | **Post** /api/auth/clientverificationrequest/ | 
[**AuthEmployeeCreate**](AuthAPI.md#AuthEmployeeCreate) | **Post** /api/auth/employee/ | 
[**AuthEmployeeverificationconfirmationRetrieve**](AuthAPI.md#AuthEmployeeverificationconfirmationRetrieve) | **Get** /api/auth/employeeverificationconfirmation/{uidb64}/ | 
[**AuthEmployeeverificationrequestCreate**](AuthAPI.md#AuthEmployeeverificationrequestCreate) | **Post** /api/auth/employeeverificationrequest/ | 
[**AuthInviteusersCreate**](AuthAPI.md#AuthInviteusersCreate) | **Post** /api/auth/inviteusers/ | 
[**AuthPasswordResetCompletePartialUpdate**](AuthAPI.md#AuthPasswordResetCompletePartialUpdate) | **Patch** /api/auth/password-reset-complete/ | 
[**AuthPasswordResetRetrieve**](AuthAPI.md#AuthPasswordResetRetrieve) | **Get** /api/auth/password-reset/{uidb64}/{token}/ | 
[**AuthRequestlogintokenCreate**](AuthAPI.md#AuthRequestlogintokenCreate) | **Post** /api/auth/requestlogintoken/ | 
[**AuthRequestpasswordtokenCreate**](AuthAPI.md#AuthRequestpasswordtokenCreate) | **Post** /api/auth/requestpasswordtoken/ | 
[**AuthRequestresetpassowrdCreate**](AuthAPI.md#AuthRequestresetpassowrdCreate) | **Post** /api/auth/requestresetpassowrd/ | 
[**AuthSigninCreate**](AuthAPI.md#AuthSigninCreate) | **Post** /api/auth/signin/ | 
[**AuthTokenloginCreate**](AuthAPI.md#AuthTokenloginCreate) | **Post** /api/auth/tokenlogin/ | 
[**AuthTokenpasswordchangeCreate**](AuthAPI.md#AuthTokenpasswordchangeCreate) | **Post** /api/auth/tokenpasswordchange/ | 
[**AuthTokensigninPartialUpdate**](AuthAPI.md#AuthTokensigninPartialUpdate) | **Patch** /api/auth/tokensignin/ | 
[**AuthTokensigninconfirmationCreate**](AuthAPI.md#AuthTokensigninconfirmationCreate) | **Post** /api/auth/tokensigninconfirmation/{uidb64}/ | 
[**AuthTokensigninrequestRetrieve**](AuthAPI.md#AuthTokensigninrequestRetrieve) | **Get** /api/auth/tokensigninrequest/ | 
[**AuthTokenverificationCreate**](AuthAPI.md#AuthTokenverificationCreate) | **Post** /api/auth/tokenverification/ | 
[**AuthUnverifiedusersList**](AuthAPI.md#AuthUnverifiedusersList) | **Get** /api/auth/unverifiedusers/ | 
[**AuthVerifyinvitedusersCreate**](AuthAPI.md#AuthVerifyinvitedusersCreate) | **Post** /api/auth/verifyinvitedusers/ | 



## AuthAdminCreate

> Admin AuthAdminCreate(ctx).Admin(admin).Execute()





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
	admin := *openapiclient.NewAdmin("Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example") // Admin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthAdminCreate(context.Background()).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAdminCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAdminCreate`: Admin
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAdminCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **admin** | [**Admin**](Admin.md) |  | 

### Return type

[**Admin**](Admin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAdminverificationconfirmationRetrieve

> SetNewPassword AuthAdminverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
	uidb64 := "uidb64_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthAdminverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAdminverificationconfirmationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAdminverificationconfirmationRetrieve`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAdminverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAdminverificationrequestCreate

> ResetPasswordEmailRequest AuthAdminverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthAdminverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAdminverificationrequestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAdminverificationrequestCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAdminverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAnonymoususerCreate

> AnonymousUser AuthAnonymoususerCreate(ctx).AnonymousUser(anonymousUser).Execute()





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
	anonymousUser := *openapiclient.NewAnonymousUser(int32(123), "Token_example") // AnonymousUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthAnonymoususerCreate(context.Background()).AnonymousUser(anonymousUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAnonymoususerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAnonymoususerCreate`: AnonymousUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAnonymoususerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anonymousUser** | [**AnonymousUser**](AnonymousUser.md) |  | 

### Return type

[**AnonymousUser**](AnonymousUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAnonymoususerDestroy

> AuthAnonymoususerDestroy(ctx).Execute()





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
	r, err := apiClient.AuthAPI.AuthAnonymoususerDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAnonymoususerDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerDestroyRequest struct via the builder pattern


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


## AuthAnonymoususerRetrieve

> AnonymousUser AuthAnonymoususerRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.AuthAPI.AuthAnonymoususerRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAnonymoususerRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAnonymoususerRetrieve`: AnonymousUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAnonymoususerRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerRetrieveRequest struct via the builder pattern


### Return type

[**AnonymousUser**](AnonymousUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthGoogleCreate

> UserLogin AuthAuthGoogleCreate(ctx).UserLogin(userLogin).Execute()



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
	userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthAuthGoogleCreate(context.Background()).UserLogin(userLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthAuthGoogleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAuthGoogleCreate`: UserLogin
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthAuthGoogleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthGoogleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientCreate

> Client AuthClientCreate(ctx).Client(client).Execute()





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
	client := *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example") // Client | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthClientCreate(context.Background()).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthClientCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthClientCreate`: Client
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthClientCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientUsersSignupSigninCreate

> UserLogin AuthClientUsersSignupSigninCreate(ctx).UserLogin(userLogin).Execute()



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
	userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthClientUsersSignupSigninCreate(context.Background()).UserLogin(userLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthClientUsersSignupSigninCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthClientUsersSignupSigninCreate`: UserLogin
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthClientUsersSignupSigninCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientUsersSignupSigninCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientverificationconfirmationRetrieve

> SetNewPassword AuthClientverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
	uidb64 := "uidb64_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthClientverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthClientverificationconfirmationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthClientverificationconfirmationRetrieve`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthClientverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientverificationrequestCreate

> ResetPasswordEmailRequest AuthClientverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthClientverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthClientverificationrequestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthClientverificationrequestCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthClientverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeCreate

> Employee AuthEmployeeCreate(ctx).Employee(employee).Execute()





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
	employee := *openapiclient.NewEmployee(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example") // Employee | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthEmployeeCreate(context.Background()).Employee(employee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthEmployeeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthEmployeeCreate`: Employee
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthEmployeeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employee** | [**Employee**](Employee.md) |  | 

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeverificationconfirmationRetrieve

> SetNewPassword AuthEmployeeverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
	uidb64 := "uidb64_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthEmployeeverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthEmployeeverificationconfirmationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthEmployeeverificationconfirmationRetrieve`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthEmployeeverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeverificationrequestCreate

> ResetPasswordEmailRequest AuthEmployeeverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthEmployeeverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthEmployeeverificationrequestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthEmployeeverificationrequestCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthEmployeeverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthInviteusersCreate

> InviteEmployees AuthInviteusersCreate(ctx).InviteEmployees(inviteEmployees).Execute()



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
	inviteEmployees := *openapiclient.NewInviteEmployees("Email_example", int32(123)) // InviteEmployees | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthInviteusersCreate(context.Background()).InviteEmployees(inviteEmployees).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthInviteusersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthInviteusersCreate`: InviteEmployees
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthInviteusersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthInviteusersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteEmployees** | [**InviteEmployees**](InviteEmployees.md) |  | 

### Return type

[**InviteEmployees**](InviteEmployees.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthPasswordResetCompletePartialUpdate

> SetNewPassword AuthPasswordResetCompletePartialUpdate(ctx).PatchedSetNewPassword(patchedSetNewPassword).Execute()



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
	patchedSetNewPassword := *openapiclient.NewPatchedSetNewPassword() // PatchedSetNewPassword |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthPasswordResetCompletePartialUpdate(context.Background()).PatchedSetNewPassword(patchedSetNewPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthPasswordResetCompletePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthPasswordResetCompletePartialUpdate`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthPasswordResetCompletePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthPasswordResetCompletePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedSetNewPassword** | [**PatchedSetNewPassword**](PatchedSetNewPassword.md) |  | 

### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthPasswordResetRetrieve

> SetNewPassword AuthPasswordResetRetrieve(ctx, token, uidb64).Execute()



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
	token := "token_example" // string | 
	uidb64 := "uidb64_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthPasswordResetRetrieve(context.Background(), token, uidb64).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthPasswordResetRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthPasswordResetRetrieve`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthPasswordResetRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthPasswordResetRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestlogintokenCreate

> ResetPasswordEmailRequest AuthRequestlogintokenCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthRequestlogintokenCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthRequestlogintokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRequestlogintokenCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthRequestlogintokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestlogintokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestpasswordtokenCreate

> ResetPasswordEmailRequest AuthRequestpasswordtokenCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthRequestpasswordtokenCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthRequestpasswordtokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRequestpasswordtokenCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthRequestpasswordtokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestpasswordtokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestresetpassowrdCreate

> ResetPasswordEmailRequest AuthRequestresetpassowrdCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthRequestresetpassowrdCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthRequestresetpassowrdCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRequestresetpassowrdCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthRequestresetpassowrdCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestresetpassowrdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSigninCreate

> UserLogin AuthSigninCreate(ctx).UserLogin(userLogin).Execute()



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
	userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthSigninCreate(context.Background()).UserLogin(userLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthSigninCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSigninCreate`: UserLogin
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthSigninCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSigninCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenloginCreate

> TokenLogin AuthTokenloginCreate(ctx).TokenLogin(tokenLogin).Execute()





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
	tokenLogin := *openapiclient.NewTokenLogin(int32(123), "Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", openapiclient.UserTypeFc6Enum("Admin")) // TokenLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokenloginCreate(context.Background()).TokenLogin(tokenLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokenloginCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenloginCreate`: TokenLogin
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokenloginCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenloginCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLogin** | [**TokenLogin**](TokenLogin.md) |  | 

### Return type

[**TokenLogin**](TokenLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenpasswordchangeCreate

> TokenSetNewPassword AuthTokenpasswordchangeCreate(ctx).TokenSetNewPassword(tokenSetNewPassword).Execute()





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
	tokenSetNewPassword := *openapiclient.NewTokenSetNewPassword("Password_example", int32(123)) // TokenSetNewPassword | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokenpasswordchangeCreate(context.Background()).TokenSetNewPassword(tokenSetNewPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokenpasswordchangeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenpasswordchangeCreate`: TokenSetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokenpasswordchangeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenpasswordchangeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenSetNewPassword** | [**TokenSetNewPassword**](TokenSetNewPassword.md) |  | 

### Return type

[**TokenSetNewPassword**](TokenSetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninPartialUpdate

> SetNewPassword AuthTokensigninPartialUpdate(ctx).PatchedSetNewPassword(patchedSetNewPassword).Execute()



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
	patchedSetNewPassword := *openapiclient.NewPatchedSetNewPassword() // PatchedSetNewPassword |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokensigninPartialUpdate(context.Background()).PatchedSetNewPassword(patchedSetNewPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokensigninPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokensigninPartialUpdate`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokensigninPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedSetNewPassword** | [**PatchedSetNewPassword**](PatchedSetNewPassword.md) |  | 

### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninconfirmationCreate

> ResetPasswordEmailRequest AuthTokensigninconfirmationCreate(ctx, uidb64).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
	uidb64 := "uidb64_example" // string | 
	resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokensigninconfirmationCreate(context.Background(), uidb64).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokensigninconfirmationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokensigninconfirmationCreate`: ResetPasswordEmailRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokensigninconfirmationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninconfirmationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninrequestRetrieve

> SetNewPassword AuthTokensigninrequestRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.AuthAPI.AuthTokensigninrequestRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokensigninrequestRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokensigninrequestRetrieve`: SetNewPassword
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokensigninrequestRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninrequestRetrieveRequest struct via the builder pattern


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenverificationCreate

> EmailVerification AuthTokenverificationCreate(ctx).EmailVerification(emailVerification).Execute()





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
	emailVerification := *openapiclient.NewEmailVerification(int32(123)) // EmailVerification | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokenverificationCreate(context.Background()).EmailVerification(emailVerification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokenverificationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenverificationCreate`: EmailVerification
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokenverificationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenverificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailVerification** | [**EmailVerification**](EmailVerification.md) |  | 

### Return type

[**EmailVerification**](EmailVerification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthUnverifiedusersList

> PaginatedUnverifiedUsersList AuthUnverifiedusersList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.AuthAPI.AuthUnverifiedusersList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthUnverifiedusersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUnverifiedusersList`: PaginatedUnverifiedUsersList
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthUnverifiedusersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUnverifiedusersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedUnverifiedUsersList**](PaginatedUnverifiedUsersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthVerifyinvitedusersCreate

> VerifyUsersToken AuthVerifyinvitedusersCreate(ctx).VerifyUsersToken(verifyUsersToken).Execute()





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
	verifyUsersToken := *openapiclient.NewVerifyUsersToken(int32(123)) // VerifyUsersToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthVerifyinvitedusersCreate(context.Background()).VerifyUsersToken(verifyUsersToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthVerifyinvitedusersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthVerifyinvitedusersCreate`: VerifyUsersToken
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthVerifyinvitedusersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthVerifyinvitedusersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyUsersToken** | [**VerifyUsersToken**](VerifyUsersToken.md) |  | 

### Return type

[**VerifyUsersToken**](VerifyUsersToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

