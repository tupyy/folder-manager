# \TagsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /api/gphotos/v1/tags | 
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /api/gphotos/v1/tags/{tag_id} | 
[**GetTags**](TagsApi.md#GetTags) | **Get** /api/gphotos/v1/tags | 
[**RemoveTagFromAlbum**](TagsApi.md#RemoveTagFromAlbum) | **Delete** /api/gphotos/v1/albums/{album_id}/tags/{tag_id} | 
[**SetTagToAlbum**](TagsApi.md#SetTagToAlbum) | **Post** /api/gphotos/v1/albums/{album_id}/tags/{tag_id} | 
[**UpdateTag**](TagsApi.md#UpdateTag) | **Patch** /api/gphotos/v1/tags/{tag_id} | 



## CreateTag

> Tag CreateTag(ctx, optional)



Create an tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagRequestPayload** | [**optional.Interface of TagRequestPayload**](TagRequestPayload.md)| Tag data | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> Tag DeleteTag(ctx, tagId)



Delete the specified tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string**| The ID of the tag | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> TagList GetTags(ctx, optional)



Get all tags owned the current logged user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| page number | 
 **size** | **optional.Int32**| total number of items per page | 

### Return type

[**TagList**](TagList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTagFromAlbum

> RemoveTagFromAlbum(ctx, albumId, tagId)



Untag an album

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
**tagId** | **string**| The ID of the tag | 

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


## SetTagToAlbum

> Tag SetTagToAlbum(ctx, albumId, tagId)



Associate a tag with an album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The ID of the album | 
**tagId** | **string**| The ID of the tag | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> Tag UpdateTag(ctx, tagId, optional)



Update the specified tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string**| The ID of the tag | 
 **optional** | ***UpdateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagRequestPayload** | [**optional.Interface of TagRequestPayload**](TagRequestPayload.md)| Tag data | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

