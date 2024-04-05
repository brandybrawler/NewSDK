# \SurveyAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SurveyEndsurveyCreate**](SurveyAPI.md#SurveyEndsurveyCreate) | **Post** /api/survey/endsurvey/ | 
[**SurveyPromptsurveyCreate**](SurveyAPI.md#SurveyPromptsurveyCreate) | **Post** /api/survey/promptsurvey/ | 
[**SurveyPromptsurveyRetrieve**](SurveyAPI.md#SurveyPromptsurveyRetrieve) | **Get** /api/survey/promptsurvey/ | 
[**SurveyQuestioninducedCreate**](SurveyAPI.md#SurveyQuestioninducedCreate) | **Post** /api/survey/questioninduced/ | 
[**SurveyQuestioninducedRetrieve**](SurveyAPI.md#SurveyQuestioninducedRetrieve) | **Get** /api/survey/questioninduced/ | 
[**SurveyResponseCreate**](SurveyAPI.md#SurveyResponseCreate) | **Post** /api/survey/response/ | 
[**SurveyResponseRetrieve**](SurveyAPI.md#SurveyResponseRetrieve) | **Get** /api/survey/response/ | 
[**SurveySurveyCreate**](SurveyAPI.md#SurveySurveyCreate) | **Post** /api/survey/survey/ | 
[**SurveySurveyRetrieve**](SurveyAPI.md#SurveySurveyRetrieve) | **Get** /api/survey/survey/ | 
[**SurveySurveyreportviewCreate**](SurveyAPI.md#SurveySurveyreportviewCreate) | **Post** /api/survey/surveyreportview/ | 
[**SurveySurveyreportviewRetrieve**](SurveyAPI.md#SurveySurveyreportviewRetrieve) | **Get** /api/survey/surveyreportview/ | 
[**SurveySurveysubgroupCreate**](SurveyAPI.md#SurveySurveysubgroupCreate) | **Post** /api/survey/surveysubgroup/ | 
[**SurveySurveysubgroupRetrieve**](SurveyAPI.md#SurveySurveysubgroupRetrieve) | **Get** /api/survey/surveysubgroup/ | 



## SurveyEndsurveyCreate

> SurveyEndsurveyCreate(ctx).Execute()





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
	r, err := apiClient.SurveyAPI.SurveyEndsurveyCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyEndsurveyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyEndsurveyCreateRequest struct via the builder pattern


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


## SurveyPromptsurveyCreate

> PromptSurvey SurveyPromptsurveyCreate(ctx).PromptSurvey(promptSurvey).Execute()



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
	promptSurvey := *openapiclient.NewPromptSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", "SurveyPrompt_example", *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example")) // PromptSurvey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveyPromptsurveyCreate(context.Background()).PromptSurvey(promptSurvey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyPromptsurveyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyPromptsurveyCreate`: PromptSurvey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyPromptsurveyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyPromptsurveyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptSurvey** | [**PromptSurvey**](PromptSurvey.md) |  | 

### Return type

[**PromptSurvey**](PromptSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyPromptsurveyRetrieve

> PromptSurvey SurveyPromptsurveyRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveyPromptsurveyRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyPromptsurveyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyPromptsurveyRetrieve`: PromptSurvey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyPromptsurveyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyPromptsurveyRetrieveRequest struct via the builder pattern


### Return type

[**PromptSurvey**](PromptSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyQuestioninducedCreate

> QuestionInducedSurvey SurveyQuestioninducedCreate(ctx).QuestionInducedSurvey(questionInducedSurvey).Execute()



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
	questionInducedSurvey := *openapiclient.NewQuestionInducedSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example")) // QuestionInducedSurvey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveyQuestioninducedCreate(context.Background()).QuestionInducedSurvey(questionInducedSurvey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyQuestioninducedCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyQuestioninducedCreate`: QuestionInducedSurvey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyQuestioninducedCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyQuestioninducedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionInducedSurvey** | [**QuestionInducedSurvey**](QuestionInducedSurvey.md) |  | 

### Return type

[**QuestionInducedSurvey**](QuestionInducedSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyQuestioninducedRetrieve

> QuestionInducedSurvey SurveyQuestioninducedRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveyQuestioninducedRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyQuestioninducedRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyQuestioninducedRetrieve`: QuestionInducedSurvey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyQuestioninducedRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyQuestioninducedRetrieveRequest struct via the builder pattern


### Return type

[**QuestionInducedSurvey**](QuestionInducedSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyResponseCreate

> Response SurveyResponseCreate(ctx).Response(response).Execute()





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
	response := *openapiclient.NewResponse(int32(123), *openapiclient.NewTargetAudience(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), time.Now(), "CreatedAt_example") // Response | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveyResponseCreate(context.Background()).Response(response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyResponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyResponseCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyResponseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyResponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **response** | [**Response**](Response.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyResponseRetrieve

> Response SurveyResponseRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveyResponseRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveyResponseRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveyResponseRetrieve`: Response
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveyResponseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyResponseRetrieveRequest struct via the builder pattern


### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyCreate

> Survey SurveySurveyCreate(ctx).Survey(survey).Execute()





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
	survey := *openapiclient.NewSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", "SurveyContext_example", time.Now(), int32(123), int32(123), float64(123), int32(123)) // Survey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveySurveyCreate(context.Background()).Survey(survey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveyCreate`: Survey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **survey** | [**Survey**](Survey.md) |  | 

### Return type

[**Survey**](Survey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyRetrieve

> Survey SurveySurveyRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveySurveyRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveyRetrieve`: Survey
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyRetrieveRequest struct via the builder pattern


### Return type

[**Survey**](Survey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyreportviewCreate

> SurveyReport SurveySurveyreportviewCreate(ctx).SurveyReport(surveyReport).Execute()





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
	surveyReport := *openapiclient.NewSurveyReport(int32(123), int32(123), "Conclusion_example") // SurveyReport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveySurveyreportviewCreate(context.Background()).SurveyReport(surveyReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveyreportviewCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveyreportviewCreate`: SurveyReport
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveyreportviewCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyreportviewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surveyReport** | [**SurveyReport**](SurveyReport.md) |  | 

### Return type

[**SurveyReport**](SurveyReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyreportviewRetrieve

> SurveyReport SurveySurveyreportviewRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveySurveyreportviewRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveyreportviewRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveyreportviewRetrieve`: SurveyReport
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveyreportviewRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyreportviewRetrieveRequest struct via the builder pattern


### Return type

[**SurveyReport**](SurveyReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveysubgroupCreate

> SurveySubGroups SurveySurveysubgroupCreate(ctx).SurveySubGroups(surveySubGroups).Execute()





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
	surveySubGroups := *openapiclient.NewSurveySubGroups(int32(123), "SubgroupName_example", "SubgroupDescription_example", []openapiclient.TargetAudience{*openapiclient.NewTargetAudience(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")}) // SurveySubGroups | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyAPI.SurveySurveysubgroupCreate(context.Background()).SurveySubGroups(surveySubGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveysubgroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveysubgroupCreate`: SurveySubGroups
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveysubgroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveysubgroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surveySubGroups** | [**SurveySubGroups**](SurveySubGroups.md) |  | 

### Return type

[**SurveySubGroups**](SurveySubGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveysubgroupRetrieve

> SurveySubGroups SurveySurveysubgroupRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.SurveyAPI.SurveySurveysubgroupRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyAPI.SurveySurveysubgroupRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveySurveysubgroupRetrieve`: SurveySubGroups
	fmt.Fprintf(os.Stdout, "Response from `SurveyAPI.SurveySurveysubgroupRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveysubgroupRetrieveRequest struct via the builder pattern


### Return type

[**SurveySubGroups**](SurveySubGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

