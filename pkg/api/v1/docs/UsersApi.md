# \UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelatedGroups**](UsersApi.md#GetRelatedGroups) | **Get** /api/gphotos/v1/users/{user_id}/groups/related | 
[**GetRelatedUsers**](UsersApi.md#GetRelatedUsers) | **Get** /api/gphotos/v1/users/{user_id}/related | 
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/gphotos/v1/users | 



## GetRelatedGroups

> UserList GetRelatedGroups(ctx, userId)



Get all the groups which shared an album with specified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user | 

### Return type

[**UserList**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedUsers

> UserList GetRelatedUsers(ctx, userId)



Get all the users which shared an album with the logged user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user | 

### Return type

[**UserList**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UserList GetUsers(ctx, )



Return the list of all users except administators.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserList**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

