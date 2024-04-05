# \TenantivaAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantivaCreateagentCreate**](TenantivaAPI.md#TenantivaCreateagentCreate) | **Post** /api/tenantiva/createagent/ | 
[**TenantivaSampletoolsRetrieve**](TenantivaAPI.md#TenantivaSampletoolsRetrieve) | **Get** /api/tenantiva/sampletools/ | 
[**TenantivaTenantknowlegebaseDestroy**](TenantivaAPI.md#TenantivaTenantknowlegebaseDestroy) | **Delete** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTenantknowlegebasePartialUpdate**](TenantivaAPI.md#TenantivaTenantknowlegebasePartialUpdate) | **Patch** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTenantknowlegebaseRetrieve**](TenantivaAPI.md#TenantivaTenantknowlegebaseRetrieve) | **Get** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTrainingprogressRetrieve**](TenantivaAPI.md#TenantivaTrainingprogressRetrieve) | **Get** /api/tenantiva/trainingprogress/ | 



## TenantivaCreateagentCreate

> CreateAgent TenantivaCreateagentCreate(ctx).CreateAgent(createAgent).Execute()



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
	createAgent := *openapiclient.NewCreateAgent("TenantName_example", "ToolName_example") // CreateAgent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantivaAPI.TenantivaCreateagentCreate(context.Background()).CreateAgent(createAgent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaCreateagentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantivaCreateagentCreate`: CreateAgent
	fmt.Fprintf(os.Stdout, "Response from `TenantivaAPI.TenantivaCreateagentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaCreateagentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgent** | [**CreateAgent**](CreateAgent.md) |  | 

### Return type

[**CreateAgent**](CreateAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaSampletoolsRetrieve

> TenantivaSampletoolsRetrieve(ctx).Execute()



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
	r, err := apiClient.TenantivaAPI.TenantivaSampletoolsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaSampletoolsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaSampletoolsRetrieveRequest struct via the builder pattern


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


## TenantivaTenantknowlegebaseDestroy

> TenantivaTenantknowlegebaseDestroy(ctx).Execute()





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
	r, err := apiClient.TenantivaAPI.TenantivaTenantknowlegebaseDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaTenantknowlegebaseDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebaseDestroyRequest struct via the builder pattern


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


## TenantivaTenantknowlegebasePartialUpdate

> Knowledgebase TenantivaTenantknowlegebasePartialUpdate(ctx).PatchedKnowledgebase(patchedKnowledgebase).Execute()





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
	patchedKnowledgebase := *openapiclient.NewPatchedKnowledgebase() // PatchedKnowledgebase |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantivaAPI.TenantivaTenantknowlegebasePartialUpdate(context.Background()).PatchedKnowledgebase(patchedKnowledgebase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaTenantknowlegebasePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantivaTenantknowlegebasePartialUpdate`: Knowledgebase
	fmt.Fprintf(os.Stdout, "Response from `TenantivaAPI.TenantivaTenantknowlegebasePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebasePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedKnowledgebase** | [**PatchedKnowledgebase**](PatchedKnowledgebase.md) |  | 

### Return type

[**Knowledgebase**](Knowledgebase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTenantknowlegebaseRetrieve

> Knowledgebase TenantivaTenantknowlegebaseRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantivaAPI.TenantivaTenantknowlegebaseRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaTenantknowlegebaseRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantivaTenantknowlegebaseRetrieve`: Knowledgebase
	fmt.Fprintf(os.Stdout, "Response from `TenantivaAPI.TenantivaTenantknowlegebaseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebaseRetrieveRequest struct via the builder pattern


### Return type

[**Knowledgebase**](Knowledgebase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTrainingprogressRetrieve

> AssistantTrainingProgressase TenantivaTrainingprogressRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.TenantivaAPI.TenantivaTrainingprogressRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantivaAPI.TenantivaTrainingprogressRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantivaTrainingprogressRetrieve`: AssistantTrainingProgressase
	fmt.Fprintf(os.Stdout, "Response from `TenantivaAPI.TenantivaTrainingprogressRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTrainingprogressRetrieveRequest struct via the builder pattern


### Return type

[**AssistantTrainingProgressase**](AssistantTrainingProgressase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

