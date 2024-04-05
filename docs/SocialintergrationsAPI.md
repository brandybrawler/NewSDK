# \SocialintergrationsAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SocialintergrationsFacebookintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsFacebookintergrationCreate) | **Post** /api/socialintergrations/facebookintergration/ | 
[**SocialintergrationsFacebookintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsFacebookintergrationRetrieve) | **Get** /api/socialintergrations/facebookintergration/ | 
[**SocialintergrationsInstagramintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsInstagramintergrationCreate) | **Post** /api/socialintergrations/instagramintergration/ | 
[**SocialintergrationsInstagramintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsInstagramintergrationRetrieve) | **Get** /api/socialintergrations/instagramintergration/ | 
[**SocialintergrationsProximawhatsappintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsProximawhatsappintergrationCreate) | **Post** /api/socialintergrations/proximawhatsappintergration/ | 
[**SocialintergrationsProximawhatsappintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsProximawhatsappintergrationRetrieve) | **Get** /api/socialintergrations/proximawhatsappintergration/ | 
[**SocialintergrationsSavefacebookintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsSavefacebookintergrationCreate) | **Post** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSavefacebookintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsSavefacebookintergrationRetrieve) | **Get** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSavefacebookintergrationUpdate**](SocialintergrationsAPI.md#SocialintergrationsSavefacebookintergrationUpdate) | **Put** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSaveinstagramintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsSaveinstagramintergrationCreate) | **Post** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSaveinstagramintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsSaveinstagramintergrationRetrieve) | **Get** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSaveinstagramintergrationUpdate**](SocialintergrationsAPI.md#SocialintergrationsSaveinstagramintergrationUpdate) | **Put** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSavewhatsappintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsSavewhatsappintergrationCreate) | **Post** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsSavewhatsappintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsSavewhatsappintergrationRetrieve) | **Get** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsSavewhatsappintergrationUpdate**](SocialintergrationsAPI.md#SocialintergrationsSavewhatsappintergrationUpdate) | **Put** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsWhatsappintergrationCreate**](SocialintergrationsAPI.md#SocialintergrationsWhatsappintergrationCreate) | **Post** /api/socialintergrations/whatsappintergration/ | 
[**SocialintergrationsWhatsappintergrationRetrieve**](SocialintergrationsAPI.md#SocialintergrationsWhatsappintergrationRetrieve) | **Get** /api/socialintergrations/whatsappintergration/ | 



## SocialintergrationsFacebookintergrationCreate

> SocialintergrationsFacebookintergrationCreate(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsFacebookintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsFacebookintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsFacebookintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsFacebookintergrationRetrieve

> SocialintergrationsFacebookintergrationRetrieve(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsFacebookintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsFacebookintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsFacebookintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsInstagramintergrationCreate

> SocialintergrationsInstagramintergrationCreate(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsInstagramintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsInstagramintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsInstagramintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsInstagramintergrationRetrieve

> SocialintergrationsInstagramintergrationRetrieve(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsInstagramintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsInstagramintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsInstagramintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsProximawhatsappintergrationCreate

> SocialintergrationsProximawhatsappintergrationCreate(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsProximawhatsappintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsProximawhatsappintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsProximawhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsProximawhatsappintergrationRetrieve

> SocialintergrationsProximawhatsappintergrationRetrieve(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsProximawhatsappintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsProximawhatsappintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsProximawhatsappintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationCreate

> SocialintergrationsSavefacebookintergrationCreate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationRetrieve

> SocialintergrationsSavefacebookintergrationRetrieve(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationUpdate

> SocialintergrationsSavefacebookintergrationUpdate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavefacebookintergrationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationCreate

> SocialintergrationsSaveinstagramintergrationCreate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationRetrieve

> SocialintergrationsSaveinstagramintergrationRetrieve(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationUpdate

> SocialintergrationsSaveinstagramintergrationUpdate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSaveinstagramintergrationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationCreate

> SocialintergrationsSavewhatsappintergrationCreate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationRetrieve

> SocialintergrationsSavewhatsappintergrationRetrieve(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationUpdate

> SocialintergrationsSavewhatsappintergrationUpdate(ctx).Execute()





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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsSavewhatsappintergrationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsWhatsappintergrationCreate

> SocialintergrationsWhatsappintergrationCreate(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsWhatsappintergrationCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsWhatsappintergrationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsWhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsWhatsappintergrationRetrieve

> SocialintergrationsWhatsappintergrationRetrieve(ctx).Execute()



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
	r, err := apiClient.SocialintergrationsAPI.SocialintergrationsWhatsappintergrationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsAPI.SocialintergrationsWhatsappintergrationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsWhatsappintergrationRetrieveRequest struct via the builder pattern


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

