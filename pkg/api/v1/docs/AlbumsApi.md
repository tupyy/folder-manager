# \AlbumsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlbum**](AlbumsApi.md#CreateAlbum) | **Post** /api/gphotos/v1/albums | 
[**DeleteAlbum**](AlbumsApi.md#DeleteAlbum) | **Delete** /api/gphotos/v1/albums/{album_id} | 
[**GetAlbumByID**](AlbumsApi.md#GetAlbumByID) | **Get** /api/gphotos/v1/albums/{album_id} | 
[**GetAlbumThumbnail**](AlbumsApi.md#GetAlbumThumbnail) | **Get** /api/gphotos/v1/albums/{album_id}/thumbnail | 
[**GetAlbums**](AlbumsApi.md#GetAlbums) | **Get** /api/gphotos/v1/albums | 
[**GetAlbumsByGroup**](AlbumsApi.md#GetAlbumsByGroup) | **Get** /api/gphotos/v1/albums/groups/{group_id} | 
[**GetAlbumsByUser**](AlbumsApi.md#GetAlbumsByUser) | **Get** /api/gphotos/v1/albums/users/{user_id} | 
[**UpdateAlbum**](AlbumsApi.md#UpdateAlbum) | **Patch** /api/gphotos/v1/albums/{album_id} | 



## CreateAlbum

> Album CreateAlbum(ctx, optional)



Create an album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAlbumOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAlbumOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumRequestPayload** | [**optional.Interface of AlbumRequestPayload**](AlbumRequestPayload.md)| album data | 

### Return type

[**Album**](Album.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlbum

> Album DeleteAlbum(ctx, albumId)



Delete the album with specified id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 

### Return type

[**Album**](Album.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbumByID

> Album GetAlbumByID(ctx, albumId)



Get album by specified id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 

### Return type

[**Album**](Album.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbumThumbnail

> *os.File GetAlbumThumbnail(ctx, albumId)



Get the thumbnail of the specified album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbums

> AlbumList GetAlbums(ctx, optional)



Get all albums owned by or shared with the current logged user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlbumsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlbumsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort the list of albums. | 
 **personal** | **optional.Bool**| Fetch personal albums. | 
 **shared** | **optional.Bool**| Fetch shared albums. | 
 **search** | **optional.String**| search expression | 
 **page** | **optional.Int32**| page number | 
 **size** | **optional.Int32**| total number of items per page | 

### Return type

[**AlbumList**](AlbumList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbumsByGroup

> AlbumList GetAlbumsByGroup(ctx, groupId, optional)



Get all groups's album shared with the current logger user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| The ID of the group | 
 **optional** | ***GetAlbumsByGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlbumsByGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page number | 
 **size** | **optional.Int32**| total number of items per page | 

### Return type

[**AlbumList**](AlbumList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlbumsByUser

> AlbumList GetAlbumsByUser(ctx, userId, optional)



Get all user's album shared with the current logger user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user | 
 **optional** | ***GetAlbumsByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlbumsByUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page number | 
 **size** | **optional.Int32**| total number of items per page | 

### Return type

[**AlbumList**](AlbumList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlbum

> Album UpdateAlbum(ctx, albumId, optional)



Update an album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
 **optional** | ***UpdateAlbumOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAlbumOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **albumRequestPayload** | [**optional.Interface of AlbumRequestPayload**](AlbumRequestPayload.md)| Album data | 

### Return type

[**Album**](Album.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

