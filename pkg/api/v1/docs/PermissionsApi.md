# \PermissionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlbumPermissions**](PermissionsApi.md#GetAlbumPermissions) | **Get** /api/gphotos/v1/albums/{album_id}/permissions | 
[**RemoveAlbumPermissions**](PermissionsApi.md#RemoveAlbumPermissions) | **Delete** /api/gphotos/v1/albums/{album_id}/permissions | 
[**SetAlbumPermissions**](PermissionsApi.md#SetAlbumPermissions) | **Post** /api/gphotos/v1/albums/{album_id}/permissions | 



## GetAlbumPermissions

> AlbumPermissions GetAlbumPermissions(ctx, albumId)



Retrieve the list of user and group permissions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 

### Return type

[**AlbumPermissions**](AlbumPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAlbumPermissions

> RemoveAlbumPermissions(ctx, albumId)



Delete the permission of an album

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAlbumPermissions

> AlbumPermissions SetAlbumPermissions(ctx, albumId, optional)



Set permission on an album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
 **optional** | ***SetAlbumPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAlbumPermissionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of []map[string]interface{}**](map[string]interface{}.md)| Permissions to be applied on the album. | 

### Return type

[**AlbumPermissions**](AlbumPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

