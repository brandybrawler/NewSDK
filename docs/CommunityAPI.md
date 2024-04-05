# \CommunityAPI

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunityClientleftcommunityRetrieve**](CommunityAPI.md#CommunityClientleftcommunityRetrieve) | **Get** /api/community/clientleftcommunity/ | 
[**CommunityCommentCreate**](CommunityAPI.md#CommunityCommentCreate) | **Post** /api/community/comment/ | 
[**CommunityCommentDestroy**](CommunityAPI.md#CommunityCommentDestroy) | **Delete** /api/community/comment/ | 
[**CommunityCommentRetrieve**](CommunityAPI.md#CommunityCommentRetrieve) | **Get** /api/community/comment/ | 
[**CommunityCommunityRetrieve**](CommunityAPI.md#CommunityCommunityRetrieve) | **Get** /api/community/community/ | 
[**CommunityCommunityfeedRetrieve**](CommunityAPI.md#CommunityCommunityfeedRetrieve) | **Get** /api/community/communityfeed/ | 
[**CommunityCommunitymembersRetrieve**](CommunityAPI.md#CommunityCommunitymembersRetrieve) | **Get** /api/community/communitymembers/ | 
[**CommunityCommunitytagsCreate**](CommunityAPI.md#CommunityCommunitytagsCreate) | **Post** /api/community/communitytags/ | 
[**CommunityCommunitytagsRetrieve**](CommunityAPI.md#CommunityCommunitytagsRetrieve) | **Get** /api/community/communitytags/ | 
[**CommunityFavoritecommunitiesCreate**](CommunityAPI.md#CommunityFavoritecommunitiesCreate) | **Post** /api/community/favoritecommunities/ | 
[**CommunityFavoritecommunitiesRetrieve**](CommunityAPI.md#CommunityFavoritecommunitiesRetrieve) | **Get** /api/community/favoritecommunities/ | 
[**CommunityIssueCreate**](CommunityAPI.md#CommunityIssueCreate) | **Post** /api/community/issue/ | 
[**CommunityIssuePartialUpdate**](CommunityAPI.md#CommunityIssuePartialUpdate) | **Patch** /api/community/issue/ | 
[**CommunityIssueRetrieve**](CommunityAPI.md#CommunityIssueRetrieve) | **Get** /api/community/issue/ | 
[**CommunityIssueUpdate**](CommunityAPI.md#CommunityIssueUpdate) | **Put** /api/community/issue/ | 
[**CommunityJoincommunityCreate**](CommunityAPI.md#CommunityJoincommunityCreate) | **Post** /api/community/joincommunity/ | 
[**CommunityJoinedmembersRetrieve**](CommunityAPI.md#CommunityJoinedmembersRetrieve) | **Get** /api/community/joinedmembers/ | 
[**CommunityLeavecommunityCreate**](CommunityAPI.md#CommunityLeavecommunityCreate) | **Post** /api/community/leavecommunity/ | 
[**CommunityLikeordislikecommentCreate**](CommunityAPI.md#CommunityLikeordislikecommentCreate) | **Post** /api/community/likeordislikecomment/ | 
[**CommunityThreadCreate**](CommunityAPI.md#CommunityThreadCreate) | **Post** /api/community/thread/ | 
[**CommunityThreadRetrieve**](CommunityAPI.md#CommunityThreadRetrieve) | **Get** /api/community/thread/ | 



## CommunityClientleftcommunityRetrieve

> Event CommunityClientleftcommunityRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.CommunityAPI.CommunityClientleftcommunityRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityClientleftcommunityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityClientleftcommunityRetrieve`: Event
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityClientleftcommunityRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityClientleftcommunityRetrieveRequest struct via the builder pattern


### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommentCreate

> Comment CommunityCommentCreate(ctx).Comment(comment).Execute()





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
	comment := *openapiclient.NewComment(int32(123), int32(123), *openapiclient.NewClientLikes(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), []openapiclient.ClientLikes{*openapiclient.NewClientLikes(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")}) // Comment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityCommentCreate(context.Background()).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommentCreate`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | [**Comment**](Comment.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommentDestroy

> CommunityCommentDestroy(ctx).Execute()





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
	r, err := apiClient.CommunityAPI.CommunityCommentDestroy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommentDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommentDestroyRequest struct via the builder pattern


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


## CommunityCommentRetrieve

> Comment CommunityCommentRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityCommentRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommentRetrieve`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommentRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommentRetrieveRequest struct via the builder pattern


### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommunityRetrieve

> Community CommunityCommunityRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityCommunityRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommunityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommunityRetrieve`: Community
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommunityRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommunityRetrieveRequest struct via the builder pattern


### Return type

[**Community**](Community.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommunityfeedRetrieve

> CommunityFeed CommunityCommunityfeedRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityCommunityfeedRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommunityfeedRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommunityfeedRetrieve`: CommunityFeed
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommunityfeedRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommunityfeedRetrieveRequest struct via the builder pattern


### Return type

[**CommunityFeed**](CommunityFeed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommunitymembersRetrieve

> Community CommunityCommunitymembersRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.CommunityAPI.CommunityCommunitymembersRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommunitymembersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommunitymembersRetrieve`: Community
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommunitymembersRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommunitymembersRetrieveRequest struct via the builder pattern


### Return type

[**Community**](Community.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommunitytagsCreate

> CommunityTag CommunityCommunitytagsCreate(ctx).CommunityTag(communityTag).Execute()





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
	communityTag := *openapiclient.NewCommunityTag("Name_example", int32(123)) // CommunityTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityCommunitytagsCreate(context.Background()).CommunityTag(communityTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommunitytagsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommunitytagsCreate`: CommunityTag
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommunitytagsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommunitytagsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **communityTag** | [**CommunityTag**](CommunityTag.md) |  | 

### Return type

[**CommunityTag**](CommunityTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityCommunitytagsRetrieve

> CommunityTag CommunityCommunitytagsRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityCommunitytagsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityCommunitytagsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityCommunitytagsRetrieve`: CommunityTag
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityCommunitytagsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityCommunitytagsRetrieveRequest struct via the builder pattern


### Return type

[**CommunityTag**](CommunityTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityFavoritecommunitiesCreate

> FavoriteCommunitiesClient CommunityFavoritecommunitiesCreate(ctx).FavoriteCommunitiesClient(favoriteCommunitiesClient).Execute()



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
	favoriteCommunitiesClient := *openapiclient.NewFavoriteCommunitiesClient("Username_example", "Email_example", []openapiclient.FavoriteCommunity{*openapiclient.NewFavoriteCommunity(int32(123), *openapiclient.NewCommunityTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")))}) // FavoriteCommunitiesClient | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityFavoritecommunitiesCreate(context.Background()).FavoriteCommunitiesClient(favoriteCommunitiesClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityFavoritecommunitiesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityFavoritecommunitiesCreate`: FavoriteCommunitiesClient
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityFavoritecommunitiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityFavoritecommunitiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **favoriteCommunitiesClient** | [**FavoriteCommunitiesClient**](FavoriteCommunitiesClient.md) |  | 

### Return type

[**FavoriteCommunitiesClient**](FavoriteCommunitiesClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityFavoritecommunitiesRetrieve

> FavoriteCommunitiesClient CommunityFavoritecommunitiesRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.CommunityAPI.CommunityFavoritecommunitiesRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityFavoritecommunitiesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityFavoritecommunitiesRetrieve`: FavoriteCommunitiesClient
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityFavoritecommunitiesRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityFavoritecommunitiesRetrieveRequest struct via the builder pattern


### Return type

[**FavoriteCommunitiesClient**](FavoriteCommunitiesClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityIssueCreate

> Issue CommunityIssueCreate(ctx).Issue(issue).Execute()





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
	issue := *openapiclient.NewIssue(int32(123), "Issue_example", "Description_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, "CreatedAt_example", time.Now(), []openapiclient.CommunityTag{*openapiclient.NewCommunityTag("Name_example", int32(123))}, int32(123), int32(123)) // Issue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityIssueCreate(context.Background()).Issue(issue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityIssueCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityIssueCreate`: Issue
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityIssueCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityIssueCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issue** | [**Issue**](Issue.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityIssuePartialUpdate

> Issue CommunityIssuePartialUpdate(ctx).PatchedIssue(patchedIssue).Execute()





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
	patchedIssue := *openapiclient.NewPatchedIssue() // PatchedIssue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityIssuePartialUpdate(context.Background()).PatchedIssue(patchedIssue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityIssuePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityIssuePartialUpdate`: Issue
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityIssuePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityIssuePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedIssue** | [**PatchedIssue**](PatchedIssue.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityIssueRetrieve

> Issue CommunityIssueRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityIssueRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityIssueRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityIssueRetrieve`: Issue
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityIssueRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityIssueRetrieveRequest struct via the builder pattern


### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityIssueUpdate

> Issue CommunityIssueUpdate(ctx).Issue(issue).Execute()





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
	issue := *openapiclient.NewIssue(int32(123), "Issue_example", "Description_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, "CreatedAt_example", time.Now(), []openapiclient.CommunityTag{*openapiclient.NewCommunityTag("Name_example", int32(123))}, int32(123), int32(123)) // Issue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityIssueUpdate(context.Background()).Issue(issue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityIssueUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityIssueUpdate`: Issue
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityIssueUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityIssueUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issue** | [**Issue**](Issue.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityJoincommunityCreate

> JoinCommunity CommunityJoincommunityCreate(ctx).JoinCommunity(joinCommunity).Execute()



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
	joinCommunity := *openapiclient.NewJoinCommunity(int32(123), int32(123)) // JoinCommunity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityJoincommunityCreate(context.Background()).JoinCommunity(joinCommunity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityJoincommunityCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityJoincommunityCreate`: JoinCommunity
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityJoincommunityCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityJoincommunityCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **joinCommunity** | [**JoinCommunity**](JoinCommunity.md) |  | 

### Return type

[**JoinCommunity**](JoinCommunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityJoinedmembersRetrieve

> Event CommunityJoinedmembersRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.CommunityAPI.CommunityJoinedmembersRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityJoinedmembersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityJoinedmembersRetrieve`: Event
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityJoinedmembersRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityJoinedmembersRetrieveRequest struct via the builder pattern


### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityLeavecommunityCreate

> LeaveCommunity CommunityLeavecommunityCreate(ctx).LeaveCommunity(leaveCommunity).Execute()



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
	leaveCommunity := *openapiclient.NewLeaveCommunity(int32(123), int32(123)) // LeaveCommunity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityLeavecommunityCreate(context.Background()).LeaveCommunity(leaveCommunity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityLeavecommunityCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityLeavecommunityCreate`: LeaveCommunity
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityLeavecommunityCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityLeavecommunityCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaveCommunity** | [**LeaveCommunity**](LeaveCommunity.md) |  | 

### Return type

[**LeaveCommunity**](LeaveCommunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityLikeordislikecommentCreate

> CommunityLikeordislikecommentCreate(ctx).Execute()



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
	r, err := apiClient.CommunityAPI.CommunityLikeordislikecommentCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityLikeordislikecommentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityLikeordislikecommentCreateRequest struct via the builder pattern


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


## CommunityThreadCreate

> Thread CommunityThreadCreate(ctx).Thread(thread).Execute()





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
	thread := *openapiclient.NewThread(int32(123), *openapiclient.NewNested(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "Issue_example", "Description_example", int32(123))) // Thread |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunityAPI.CommunityThreadCreate(context.Background()).Thread(thread).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityThreadCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityThreadCreate`: Thread
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityThreadCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunityThreadCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thread** | [**Thread**](Thread.md) |  | 

### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunityThreadRetrieve

> Thread CommunityThreadRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CommunityAPI.CommunityThreadRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityAPI.CommunityThreadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunityThreadRetrieve`: Thread
	fmt.Fprintf(os.Stdout, "Response from `CommunityAPI.CommunityThreadRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommunityThreadRetrieveRequest struct via the builder pattern


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

