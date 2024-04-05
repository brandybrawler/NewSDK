# \CallAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallAnalyticsCallResolutionAnalyticsRetrieve**](CallAPI.md#CallAnalyticsCallResolutionAnalyticsRetrieve) | **Get** /api/call/analytics/call-resolution-analytics/ | 
[**CallAnalyticsCallTagAnalyticsRetrieve**](CallAPI.md#CallAnalyticsCallTagAnalyticsRetrieve) | **Get** /api/call/analytics/call-tag-analytics/ | 
[**CallAnalyticsCallVolumeAnalyticsRetrieve**](CallAPI.md#CallAnalyticsCallVolumeAnalyticsRetrieve) | **Get** /api/call/analytics/call-volume-analytics/ | 
[**CallAnalyticsCustomerSatisfactionAnalyticsRetrieve**](CallAPI.md#CallAnalyticsCustomerSatisfactionAnalyticsRetrieve) | **Get** /api/call/analytics/customer-satisfaction-analytics/ | 
[**CallAnalyticsFeedbackCategoryDistributionRetrieve**](CallAPI.md#CallAnalyticsFeedbackCategoryDistributionRetrieve) | **Get** /api/call/analytics/feedback-category-distribution/ | 
[**CallAnalyticsInteractionTypeAnalyticsRetrieve**](CallAPI.md#CallAnalyticsInteractionTypeAnalyticsRetrieve) | **Get** /api/call/analytics/interaction-type-analytics/ | 
[**CallCallAgentPerformanceCreate**](CallAPI.md#CallCallAgentPerformanceCreate) | **Post** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformanceDestroy**](CallAPI.md#CallCallAgentPerformanceDestroy) | **Delete** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformancePartialUpdate**](CallAPI.md#CallCallAgentPerformancePartialUpdate) | **Patch** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformanceRetrieve**](CallAPI.md#CallCallAgentPerformanceRetrieve) | **Get** /api/call/call/agent-performance/ | 
[**CallCallCallMetricsCreate**](CallAPI.md#CallCallCallMetricsCreate) | **Post** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsDestroy**](CallAPI.md#CallCallCallMetricsDestroy) | **Delete** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsPartialUpdate**](CallAPI.md#CallCallCallMetricsPartialUpdate) | **Patch** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsRetrieve**](CallAPI.md#CallCallCallMetricsRetrieve) | **Get** /api/call/call/call-metrics/ | 
[**CallCallCallResolutionCreate**](CallAPI.md#CallCallCallResolutionCreate) | **Post** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionDestroy**](CallAPI.md#CallCallCallResolutionDestroy) | **Delete** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionPartialUpdate**](CallAPI.md#CallCallCallResolutionPartialUpdate) | **Patch** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionRetrieve**](CallAPI.md#CallCallCallResolutionRetrieve) | **Get** /api/call/call/call-resolution/ | 
[**CallCallCallTagCreate**](CallAPI.md#CallCallCallTagCreate) | **Post** /api/call/call/call-tag/ | 
[**CallCallCallTagDestroy**](CallAPI.md#CallCallCallTagDestroy) | **Delete** /api/call/call/call-tag/ | 
[**CallCallCallTagMappingCreate**](CallAPI.md#CallCallCallTagMappingCreate) | **Post** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagMappingDestroy**](CallAPI.md#CallCallCallTagMappingDestroy) | **Delete** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagMappingRetrieve**](CallAPI.md#CallCallCallTagMappingRetrieve) | **Get** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagRetrieve**](CallAPI.md#CallCallCallTagRetrieve) | **Get** /api/call/call/call-tag/ | 
[**CallCallCallsCreate**](CallAPI.md#CallCallCallsCreate) | **Post** /api/call/call/calls/ | 
[**CallCallCallsDestroy**](CallAPI.md#CallCallCallsDestroy) | **Delete** /api/call/call/calls/ | 
[**CallCallCallsPartialUpdate**](CallAPI.md#CallCallCallsPartialUpdate) | **Patch** /api/call/call/calls/ | 
[**CallCallCallsRetrieve**](CallAPI.md#CallCallCallsRetrieve) | **Get** /api/call/call/calls/ | 
[**CallCallCustomerFeedbackCreate**](CallAPI.md#CallCallCustomerFeedbackCreate) | **Post** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackDestroy**](CallAPI.md#CallCallCustomerFeedbackDestroy) | **Delete** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackPartialUpdate**](CallAPI.md#CallCallCustomerFeedbackPartialUpdate) | **Patch** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackRetrieve**](CallAPI.md#CallCallCustomerFeedbackRetrieve) | **Get** /api/call/call/customer-feedback/ | 
[**CallCallInteractionHistoryCreate**](CallAPI.md#CallCallInteractionHistoryCreate) | **Post** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryDestroy**](CallAPI.md#CallCallInteractionHistoryDestroy) | **Delete** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryPartialUpdate**](CallAPI.md#CallCallInteractionHistoryPartialUpdate) | **Patch** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryRetrieve**](CallAPI.md#CallCallInteractionHistoryRetrieve) | **Get** /api/call/call/interaction-history/ | 



## CallAnalyticsCallResolutionAnalyticsRetrieve

> CallAnalyticsCallResolutionAnalyticsRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsCallResolutionAnalyticsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsCallResolutionAnalyticsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallResolutionAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCallTagAnalyticsRetrieve

> CallAnalyticsCallTagAnalyticsRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsCallTagAnalyticsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsCallTagAnalyticsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallTagAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCallVolumeAnalyticsRetrieve

> CallAnalyticsCallVolumeAnalyticsRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsCallVolumeAnalyticsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsCallVolumeAnalyticsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallVolumeAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCustomerSatisfactionAnalyticsRetrieve

> CallAnalyticsCustomerSatisfactionAnalyticsRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsCustomerSatisfactionAnalyticsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsCustomerSatisfactionAnalyticsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCustomerSatisfactionAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsFeedbackCategoryDistributionRetrieve

> CallAnalyticsFeedbackCategoryDistributionRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsFeedbackCategoryDistributionRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsFeedbackCategoryDistributionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsFeedbackCategoryDistributionRetrieveRequest struct via the builder pattern


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


## CallAnalyticsInteractionTypeAnalyticsRetrieve

> CallAnalyticsInteractionTypeAnalyticsRetrieve(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallAnalyticsInteractionTypeAnalyticsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallAnalyticsInteractionTypeAnalyticsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsInteractionTypeAnalyticsRetrieveRequest struct via the builder pattern


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


## CallCallAgentPerformanceCreate

> AgentPerformance CallCallAgentPerformanceCreate(ctx).AgentPerformance(agentPerformance).Execute()





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
	agentPerformance := *openapiclient.NewAgentPerformance(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "AverageCallDuration_example", "AverageSatisfactionScore_example", "ResolutionRate_example", int32(123)) // AgentPerformance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallAgentPerformanceCreate(context.Background()).AgentPerformance(agentPerformance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallAgentPerformanceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallAgentPerformanceCreate`: AgentPerformance
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallAgentPerformanceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentPerformance** | [**AgentPerformance**](AgentPerformance.md) |  | 

### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallAgentPerformanceDestroy

> CallCallAgentPerformanceDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallAgentPerformanceDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallAgentPerformanceDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceDestroyRequest struct via the builder pattern


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


## CallCallAgentPerformancePartialUpdate

> AgentPerformance CallCallAgentPerformancePartialUpdate(ctx).PatchedAgentPerformance(patchedAgentPerformance).Execute()





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
	patchedAgentPerformance := *openapiclient.NewPatchedAgentPerformance() // PatchedAgentPerformance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallAgentPerformancePartialUpdate(context.Background()).PatchedAgentPerformance(patchedAgentPerformance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallAgentPerformancePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallAgentPerformancePartialUpdate`: AgentPerformance
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallAgentPerformancePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformancePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedAgentPerformance** | [**PatchedAgentPerformance**](PatchedAgentPerformance.md) |  | 

### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallAgentPerformanceRetrieve

> AgentPerformance CallCallAgentPerformanceRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallAgentPerformanceRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallAgentPerformanceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallAgentPerformanceRetrieve`: AgentPerformance
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallAgentPerformanceRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceRetrieveRequest struct via the builder pattern


### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsCreate

> CallMetrics CallCallCallMetricsCreate(ctx).CallMetrics(callMetrics).Execute()





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
	callMetrics := *openapiclient.NewCallMetrics(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), interface{}(123)) // CallMetrics | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallMetricsCreate(context.Background()).CallMetrics(callMetrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallMetricsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallMetricsCreate`: CallMetrics
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallMetricsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callMetrics** | [**CallMetrics**](CallMetrics.md) |  | 

### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsDestroy

> CallCallCallMetricsDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCallMetricsDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallMetricsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsDestroyRequest struct via the builder pattern


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


## CallCallCallMetricsPartialUpdate

> CallMetrics CallCallCallMetricsPartialUpdate(ctx).PatchedCallMetrics(patchedCallMetrics).Execute()





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
	patchedCallMetrics := *openapiclient.NewPatchedCallMetrics() // PatchedCallMetrics |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallMetricsPartialUpdate(context.Background()).PatchedCallMetrics(patchedCallMetrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallMetricsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallMetricsPartialUpdate`: CallMetrics
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallMetricsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCallMetrics** | [**PatchedCallMetrics**](PatchedCallMetrics.md) |  | 

### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsRetrieve

> CallMetrics CallCallCallMetricsRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCallMetricsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallMetricsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallMetricsRetrieve`: CallMetrics
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallMetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsRetrieveRequest struct via the builder pattern


### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionCreate

> CallResolution CallCallCallResolutionCreate(ctx).CallResolution(callResolution).Execute()





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
	callResolution := *openapiclient.NewCallResolution(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"))) // CallResolution |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallResolutionCreate(context.Background()).CallResolution(callResolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallResolutionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallResolutionCreate`: CallResolution
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallResolutionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callResolution** | [**CallResolution**](CallResolution.md) |  | 

### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionDestroy

> CallCallCallResolutionDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCallResolutionDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallResolutionDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionDestroyRequest struct via the builder pattern


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


## CallCallCallResolutionPartialUpdate

> CallResolution CallCallCallResolutionPartialUpdate(ctx).PatchedCallResolution(patchedCallResolution).Execute()





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
	patchedCallResolution := *openapiclient.NewPatchedCallResolution() // PatchedCallResolution |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallResolutionPartialUpdate(context.Background()).PatchedCallResolution(patchedCallResolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallResolutionPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallResolutionPartialUpdate`: CallResolution
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallResolutionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCallResolution** | [**PatchedCallResolution**](PatchedCallResolution.md) |  | 

### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionRetrieve

> CallResolution CallCallCallResolutionRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCallResolutionRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallResolutionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallResolutionRetrieve`: CallResolution
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallResolutionRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionRetrieveRequest struct via the builder pattern


### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagCreate

> CallTag CallCallCallTagCreate(ctx).CallTag(callTag).Execute()





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
	callTag := *openapiclient.NewCallTag(int32(123), "Name_example") // CallTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallTagCreate(context.Background()).CallTag(callTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallTagCreate`: CallTag
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallTagCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callTag** | [**CallTag**](CallTag.md) |  | 

### Return type

[**CallTag**](CallTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagDestroy

> CallCallCallTagDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCallTagDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagDestroyRequest struct via the builder pattern


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


## CallCallCallTagMappingCreate

> CallTagMapping CallCallCallTagMappingCreate(ctx).CallTagMapping(callTagMapping).Execute()





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
	callTagMapping := *openapiclient.NewCallTagMapping(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), *openapiclient.NewCallTag(int32(123), "Name_example")) // CallTagMapping |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallTagMappingCreate(context.Background()).CallTagMapping(callTagMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagMappingCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallTagMappingCreate`: CallTagMapping
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallTagMappingCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callTagMapping** | [**CallTagMapping**](CallTagMapping.md) |  | 

### Return type

[**CallTagMapping**](CallTagMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagMappingDestroy

> CallCallCallTagMappingDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCallTagMappingDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagMappingDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingDestroyRequest struct via the builder pattern


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


## CallCallCallTagMappingRetrieve

> CallTagMapping CallCallCallTagMappingRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCallTagMappingRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagMappingRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallTagMappingRetrieve`: CallTagMapping
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallTagMappingRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingRetrieveRequest struct via the builder pattern


### Return type

[**CallTagMapping**](CallTagMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagRetrieve

> CallTag CallCallCallTagRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCallTagRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallTagRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallTagRetrieve`: CallTag
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallTagRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagRetrieveRequest struct via the builder pattern


### Return type

[**CallTag**](CallTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsCreate

> Call CallCallCallsCreate(ctx).Call(call).Execute()





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
	call := *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")) // Call | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallsCreate(context.Background()).Call(call).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallsCreate`: Call
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **call** | [**Call**](Call.md) |  | 

### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsDestroy

> CallCallCallsDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCallsDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsDestroyRequest struct via the builder pattern


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


## CallCallCallsPartialUpdate

> Call CallCallCallsPartialUpdate(ctx).PatchedCall(patchedCall).Execute()





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
	patchedCall := *openapiclient.NewPatchedCall() // PatchedCall |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCallsPartialUpdate(context.Background()).PatchedCall(patchedCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallsPartialUpdate`: Call
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCall** | [**PatchedCall**](PatchedCall.md) |  | 

### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsRetrieve

> Call CallCallCallsRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCallsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCallsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCallsRetrieve`: Call
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCallsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsRetrieveRequest struct via the builder pattern


### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackCreate

> CustomerFeedback CallCallCustomerFeedbackCreate(ctx).CustomerFeedback(customerFeedback).Execute()





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
	customerFeedback := *openapiclient.NewCustomerFeedback(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), int32(123), openapiclient.FeedbackCategoryEnum("SERVICE")) // CustomerFeedback | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCustomerFeedbackCreate(context.Background()).CustomerFeedback(customerFeedback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCustomerFeedbackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCustomerFeedbackCreate`: CustomerFeedback
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCustomerFeedbackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerFeedback** | [**CustomerFeedback**](CustomerFeedback.md) |  | 

### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackDestroy

> CallCallCustomerFeedbackDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallCustomerFeedbackDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCustomerFeedbackDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackDestroyRequest struct via the builder pattern


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


## CallCallCustomerFeedbackPartialUpdate

> CustomerFeedback CallCallCustomerFeedbackPartialUpdate(ctx).PatchedCustomerFeedback(patchedCustomerFeedback).Execute()





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
	patchedCustomerFeedback := *openapiclient.NewPatchedCustomerFeedback() // PatchedCustomerFeedback |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallCustomerFeedbackPartialUpdate(context.Background()).PatchedCustomerFeedback(patchedCustomerFeedback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCustomerFeedbackPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCustomerFeedbackPartialUpdate`: CustomerFeedback
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCustomerFeedbackPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCustomerFeedback** | [**PatchedCustomerFeedback**](PatchedCustomerFeedback.md) |  | 

### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackRetrieve

> CustomerFeedback CallCallCustomerFeedbackRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallCustomerFeedbackRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallCustomerFeedbackRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallCustomerFeedbackRetrieve`: CustomerFeedback
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallCustomerFeedbackRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackRetrieveRequest struct via the builder pattern


### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryCreate

> InteractionHistory CallCallInteractionHistoryCreate(ctx).InteractionHistory(interactionHistory).Execute()





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
	interactionHistory := *openapiclient.NewInteractionHistory(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), openapiclient.InteractionTypeEnum("NOTE"), "Details_example") // InteractionHistory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallInteractionHistoryCreate(context.Background()).InteractionHistory(interactionHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallInteractionHistoryCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallInteractionHistoryCreate`: InteractionHistory
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallInteractionHistoryCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactionHistory** | [**InteractionHistory**](InteractionHistory.md) |  | 

### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryDestroy

> CallCallInteractionHistoryDestroy(ctx).Execute()





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
	r, err := apiClient.CallAPI.CallCallInteractionHistoryDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallInteractionHistoryDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryDestroyRequest struct via the builder pattern


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


## CallCallInteractionHistoryPartialUpdate

> InteractionHistory CallCallInteractionHistoryPartialUpdate(ctx).PatchedInteractionHistory(patchedInteractionHistory).Execute()





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
	patchedInteractionHistory := *openapiclient.NewPatchedInteractionHistory() // PatchedInteractionHistory |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallAPI.CallCallInteractionHistoryPartialUpdate(context.Background()).PatchedInteractionHistory(patchedInteractionHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallInteractionHistoryPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallInteractionHistoryPartialUpdate`: InteractionHistory
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallInteractionHistoryPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedInteractionHistory** | [**PatchedInteractionHistory**](PatchedInteractionHistory.md) |  | 

### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryRetrieve

> InteractionHistory CallCallInteractionHistoryRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CallAPI.CallCallInteractionHistoryRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallAPI.CallCallInteractionHistoryRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallCallInteractionHistoryRetrieve`: InteractionHistory
	fmt.Fprintf(os.Stdout, "Response from `CallAPI.CallCallInteractionHistoryRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryRetrieveRequest struct via the builder pattern


### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

