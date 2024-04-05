# \ServicesAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesServiceformsCreate**](ServicesAPI.md#ServicesServiceformsCreate) | **Post** /api/services/serviceforms/ | 
[**ServicesServiceformsDestroy**](ServicesAPI.md#ServicesServiceformsDestroy) | **Delete** /api/services/serviceforms/ | 
[**ServicesServiceformsPartialUpdate**](ServicesAPI.md#ServicesServiceformsPartialUpdate) | **Patch** /api/services/serviceforms/ | 
[**ServicesServiceformsRetrieve**](ServicesAPI.md#ServicesServiceformsRetrieve) | **Get** /api/services/serviceforms/ | 
[**ServicesServiceformsresponseCreate**](ServicesAPI.md#ServicesServiceformsresponseCreate) | **Post** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponseDestroy**](ServicesAPI.md#ServicesServiceformsresponseDestroy) | **Delete** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponsePartialUpdate**](ServicesAPI.md#ServicesServiceformsresponsePartialUpdate) | **Patch** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponseRetrieve**](ServicesAPI.md#ServicesServiceformsresponseRetrieve) | **Get** /api/services/serviceformsresponse/ | 
[**ServicesServicerequestchatCreate**](ServicesAPI.md#ServicesServicerequestchatCreate) | **Post** /api/services/servicerequestchat/ | 
[**ServicesServicerequestchatRetrieve**](ServicesAPI.md#ServicesServicerequestchatRetrieve) | **Get** /api/services/servicerequestchat/ | 
[**ServicesServicerequestchatmessageCreate**](ServicesAPI.md#ServicesServicerequestchatmessageCreate) | **Post** /api/services/servicerequestchatmessage/ | 
[**ServicesServicerequestchatmessageRetrieve**](ServicesAPI.md#ServicesServicerequestchatmessageRetrieve) | **Get** /api/services/servicerequestchatmessage/ | 
[**ServicesServiceresquestCreate**](ServicesAPI.md#ServicesServiceresquestCreate) | **Post** /api/services/serviceresquest/ | 
[**ServicesServiceresquestDestroy**](ServicesAPI.md#ServicesServiceresquestDestroy) | **Delete** /api/services/serviceresquest/ | 
[**ServicesServiceresquestPartialUpdate**](ServicesAPI.md#ServicesServiceresquestPartialUpdate) | **Patch** /api/services/serviceresquest/ | 
[**ServicesServiceresquestRetrieve**](ServicesAPI.md#ServicesServiceresquestRetrieve) | **Get** /api/services/serviceresquest/ | 
[**ServicesServicesCreate**](ServicesAPI.md#ServicesServicesCreate) | **Post** /api/services/services/ | 
[**ServicesServicesDestroy**](ServicesAPI.md#ServicesServicesDestroy) | **Delete** /api/services/services/ | 
[**ServicesServicesPartialUpdate**](ServicesAPI.md#ServicesServicesPartialUpdate) | **Patch** /api/services/services/ | 
[**ServicesServicesRetrieve**](ServicesAPI.md#ServicesServicesRetrieve) | **Get** /api/services/services/ | 
[**ServicesServicetypeCreate**](ServicesAPI.md#ServicesServicetypeCreate) | **Post** /api/services/servicetype/ | 
[**ServicesServicetypeDestroy**](ServicesAPI.md#ServicesServicetypeDestroy) | **Delete** /api/services/servicetype/ | 
[**ServicesServicetypeRetrieve**](ServicesAPI.md#ServicesServicetypeRetrieve) | **Get** /api/services/servicetype/ | 



## ServicesServiceformsCreate

> ServiceForms ServicesServiceformsCreate(ctx).ServiceForms(serviceForms).Execute()





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
	serviceForms := *openapiclient.NewServiceForms(int32(123), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), "NameOfForm_example", interface{}(123)) // ServiceForms | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsCreate(context.Background()).ServiceForms(serviceForms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsCreate`: ServiceForms
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceForms** | [**ServiceForms**](ServiceForms.md) |  | 

### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsDestroy

> ServicesServiceformsDestroy(ctx).Execute()





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
	r, err := apiClient.ServicesAPI.ServicesServiceformsDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsDestroyRequest struct via the builder pattern


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


## ServicesServiceformsPartialUpdate

> ServiceForms ServicesServiceformsPartialUpdate(ctx).PatchedServiceForms(patchedServiceForms).Execute()





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
	patchedServiceForms := *openapiclient.NewPatchedServiceForms() // PatchedServiceForms |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsPartialUpdate(context.Background()).PatchedServiceForms(patchedServiceForms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsPartialUpdate`: ServiceForms
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceForms** | [**PatchedServiceForms**](PatchedServiceForms.md) |  | 

### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsRetrieve

> ServiceForms ServicesServiceformsRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsRetrieve`: ServiceForms
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsRetrieveRequest struct via the builder pattern


### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseCreate

> ServiceFormsResponse ServicesServiceformsresponseCreate(ctx).ServiceFormsResponse(serviceFormsResponse).Execute()





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
	serviceFormsResponse := *openapiclient.NewServiceFormsResponse("CreatedAt_example", int32(123), *openapiclient.NewServiceForms(int32(123), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), "NameOfForm_example", interface{}(123)), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), interface{}(123)) // ServiceFormsResponse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsresponseCreate(context.Background()).ServiceFormsResponse(serviceFormsResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsresponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsresponseCreate`: ServiceFormsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsresponseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceFormsResponse** | [**ServiceFormsResponse**](ServiceFormsResponse.md) |  | 

### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseDestroy

> ServicesServiceformsresponseDestroy(ctx).Execute()





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
	r, err := apiClient.ServicesAPI.ServicesServiceformsresponseDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsresponseDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseDestroyRequest struct via the builder pattern


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


## ServicesServiceformsresponsePartialUpdate

> ServiceFormsResponse ServicesServiceformsresponsePartialUpdate(ctx).PatchedServiceFormsResponse(patchedServiceFormsResponse).Execute()





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
	patchedServiceFormsResponse := *openapiclient.NewPatchedServiceFormsResponse() // PatchedServiceFormsResponse |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsresponsePartialUpdate(context.Background()).PatchedServiceFormsResponse(patchedServiceFormsResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsresponsePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsresponsePartialUpdate`: ServiceFormsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsresponsePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponsePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceFormsResponse** | [**PatchedServiceFormsResponse**](PatchedServiceFormsResponse.md) |  | 

### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseRetrieve

> ServiceFormsResponse ServicesServiceformsresponseRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServiceformsresponseRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceformsresponseRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceformsresponseRetrieve`: ServiceFormsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceformsresponseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseRetrieveRequest struct via the builder pattern


### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatCreate

> ServiceMessageChat ServicesServicerequestchatCreate(ctx).ServiceMessageChat(serviceMessageChat).Execute()





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
	serviceMessageChat := *openapiclient.NewServiceMessageChat(int32(123), int32(123), "MessageContent_example", openapiclient.MessageSenderEnum("client")) // ServiceMessageChat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServicerequestchatCreate(context.Background()).ServiceMessageChat(serviceMessageChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicerequestchatCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicerequestchatCreate`: ServiceMessageChat
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicerequestchatCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceMessageChat** | [**ServiceMessageChat**](ServiceMessageChat.md) |  | 

### Return type

[**ServiceMessageChat**](ServiceMessageChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatRetrieve

> ServiceMessageChat ServicesServicerequestchatRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServicerequestchatRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicerequestchatRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicerequestchatRetrieve`: ServiceMessageChat
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicerequestchatRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatRetrieveRequest struct via the builder pattern


### Return type

[**ServiceMessageChat**](ServiceMessageChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatmessageCreate

> ServiceRequestChat ServicesServicerequestchatmessageCreate(ctx).ServiceRequestChat(serviceRequestChat).Execute()





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
	serviceRequestChat := *openapiclient.NewServiceRequestChat(int32(123), *openapiclient.NewServiceRequest("CreatedAt_example", int32(123), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), interface{}(123), interface{}(123), interface{}(123)), time.Now(), time.Now()) // ServiceRequestChat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServicerequestchatmessageCreate(context.Background()).ServiceRequestChat(serviceRequestChat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicerequestchatmessageCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicerequestchatmessageCreate`: ServiceRequestChat
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicerequestchatmessageCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatmessageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRequestChat** | [**ServiceRequestChat**](ServiceRequestChat.md) |  | 

### Return type

[**ServiceRequestChat**](ServiceRequestChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatmessageRetrieve

> ServiceRequestChat ServicesServicerequestchatmessageRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServicerequestchatmessageRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicerequestchatmessageRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicerequestchatmessageRetrieve`: ServiceRequestChat
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicerequestchatmessageRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatmessageRetrieveRequest struct via the builder pattern


### Return type

[**ServiceRequestChat**](ServiceRequestChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestCreate

> ServiceRequest ServicesServiceresquestCreate(ctx).ServiceRequest(serviceRequest).Execute()





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
	serviceRequest := *openapiclient.NewServiceRequest("CreatedAt_example", int32(123), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), interface{}(123), interface{}(123), interface{}(123)) // ServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceresquestCreate(context.Background()).ServiceRequest(serviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceresquestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceresquestCreate`: ServiceRequest
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceresquestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRequest** | [**ServiceRequest**](ServiceRequest.md) |  | 

### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestDestroy

> ServicesServiceresquestDestroy(ctx).Execute()





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
	r, err := apiClient.ServicesAPI.ServicesServiceresquestDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceresquestDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestDestroyRequest struct via the builder pattern


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


## ServicesServiceresquestPartialUpdate

> ServiceRequest ServicesServiceresquestPartialUpdate(ctx).PatchedServiceRequest(patchedServiceRequest).Execute()





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
	patchedServiceRequest := *openapiclient.NewPatchedServiceRequest() // PatchedServiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServiceresquestPartialUpdate(context.Background()).PatchedServiceRequest(patchedServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceresquestPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceresquestPartialUpdate`: ServiceRequest
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceresquestPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceRequest** | [**PatchedServiceRequest**](PatchedServiceRequest.md) |  | 

### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestRetrieve

> ServiceRequest ServicesServiceresquestRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServiceresquestRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServiceresquestRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServiceresquestRetrieve`: ServiceRequest
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServiceresquestRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestRetrieveRequest struct via the builder pattern


### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesCreate

> Service ServicesServicesCreate(ctx).Service(service).Execute()





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
	service := *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)) // Service | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServicesCreate(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicesCreate`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesDestroy

> ServicesServicesDestroy(ctx).Execute()





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
	r, err := apiClient.ServicesAPI.ServicesServicesDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesDestroyRequest struct via the builder pattern


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


## ServicesServicesPartialUpdate

> Service ServicesServicesPartialUpdate(ctx).PatchedService(patchedService).Execute()





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
	patchedService := *openapiclient.NewPatchedService() // PatchedService |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServicesPartialUpdate(context.Background()).PatchedService(patchedService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicesPartialUpdate`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedService** | [**PatchedService**](PatchedService.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesRetrieve

> Service ServicesServicesRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServicesRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicesRetrieve`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicesRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesRetrieveRequest struct via the builder pattern


### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicetypeCreate

> ServiceType ServicesServicetypeCreate(ctx).ServiceType(serviceType).Execute()





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
	serviceType := *openapiclient.NewServiceType(int32(123), int32(123), openapiclient.ServiceTypeTypeEnum("problem_resolution"), "Description_example") // ServiceType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesServicetypeCreate(context.Background()).ServiceType(serviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicetypeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicetypeCreate`: ServiceType
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicetypeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | [**ServiceType**](ServiceType.md) |  | 

### Return type

[**ServiceType**](ServiceType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicetypeDestroy

> ServicesServicetypeDestroy(ctx).Execute()





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
	r, err := apiClient.ServicesAPI.ServicesServicetypeDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicetypeDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeDestroyRequest struct via the builder pattern


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


## ServicesServicetypeRetrieve

> ServiceType ServicesServicetypeRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.ServicesAPI.ServicesServicetypeRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesServicetypeRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesServicetypeRetrieve`: ServiceType
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesServicetypeRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeRetrieveRequest struct via the builder pattern


### Return type

[**ServiceType**](ServiceType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

