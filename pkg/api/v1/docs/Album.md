# Album

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**Name** | **string** | name of the album | 
**CreatedAt** | [**time.Time**](time.Time.md) | creation date in unix timestamp | 
**Description** | **string** | description of the album | [optional] 
**Location** | **string** | location of the album | [optional] 
**Bucket** | **string** | path of the bucket where media is stored | 
**Thumbnail** | **string** | url of the thumbnail of the album | [optional] 
**Owner** | [**ObjectReference**](ObjectReference.md) |  | 
**Photos** | [**ObjectReference**](ObjectReference.md) |  | 
**Tags** | [**[]Tag**](Tag.md) |  | [optional] 
**Permissions** | [**ObjectReference**](ObjectReference.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


