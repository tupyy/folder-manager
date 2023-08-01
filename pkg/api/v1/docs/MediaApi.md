# \MediaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePhoto**](MediaApi.md#DeletePhoto) | **Delete** /api/gphotos/v1/album/{album_id}/photo/{photo_id} | 
[**GetAlbumPhotos**](MediaApi.md#GetAlbumPhotos) | **Get** /api/gphotos/v1/albums/{album_id}/photos | 
[**GetPhoto**](MediaApi.md#GetPhoto) | **Get** /api/gphotos/v1/album/{album_id}/photo/{photo_id} | 
[**UploadPhoto**](MediaApi.md#UploadPhoto) | **Post** /api/gphotos/v1/albums/{album_id}/photos | 



## DeletePhoto

> DeletePhoto(ctx, albumId, photoId)



Delete photo with specified id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
**photoId** | **string**| The ID of the resource | 

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


## GetAlbumPhotos

> PhotoList GetAlbumPhotos(ctx, albumId, optional)



Retrieve the list of photos of the album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
 **optional** | ***GetAlbumPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlbumPhotosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page number | 
 **size** | **optional.Int32**| total number of items per page | 

### Return type

[**PhotoList**](PhotoList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoto

> *os.File GetPhoto(ctx, albumId, photoId)



Get photo with specified id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
**photoId** | **string**| The ID of the resource | 

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


## UploadPhoto

> Photo UploadPhoto(ctx, albumId, optional)



Upload a photo to specified album

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
 **optional** | ***UploadPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadPhotoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**Photo**](Photo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: image/jpeg
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

