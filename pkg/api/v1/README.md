# Go API client for v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: #VERSION#
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./v1"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlbumsApi* | [**CreateAlbum**](docs/AlbumsApi.md#createalbum) | **Post** /api/gphotos/v1/albums | 
*AlbumsApi* | [**DeleteAlbum**](docs/AlbumsApi.md#deletealbum) | **Delete** /api/gphotos/v1/albums/{album_id} | 
*AlbumsApi* | [**GetAlbumByID**](docs/AlbumsApi.md#getalbumbyid) | **Get** /api/gphotos/v1/albums/{album_id} | 
*AlbumsApi* | [**GetAlbumThumbnail**](docs/AlbumsApi.md#getalbumthumbnail) | **Get** /api/gphotos/v1/albums/{album_id}/thumbnail | 
*AlbumsApi* | [**GetAlbums**](docs/AlbumsApi.md#getalbums) | **Get** /api/gphotos/v1/albums | 
*AlbumsApi* | [**GetAlbumsByGroup**](docs/AlbumsApi.md#getalbumsbygroup) | **Get** /api/gphotos/v1/albums/groups/{group_id} | 
*AlbumsApi* | [**GetAlbumsByUser**](docs/AlbumsApi.md#getalbumsbyuser) | **Get** /api/gphotos/v1/albums/users/{user_id} | 
*AlbumsApi* | [**UpdateAlbum**](docs/AlbumsApi.md#updatealbum) | **Patch** /api/gphotos/v1/albums/{album_id} | 
*GroupsApi* | [**GetGroups**](docs/GroupsApi.md#getgroups) | **Get** /api/gphotos/v1/groups | 
*MediaApi* | [**DeletePhoto**](docs/MediaApi.md#deletephoto) | **Delete** /api/gphotos/v1/album/{album_id}/photo/{photo_id} | 
*MediaApi* | [**GetAlbumPhotos**](docs/MediaApi.md#getalbumphotos) | **Get** /api/gphotos/v1/albums/{album_id}/photos | 
*MediaApi* | [**GetPhoto**](docs/MediaApi.md#getphoto) | **Get** /api/gphotos/v1/album/{album_id}/photo/{photo_id} | 
*MediaApi* | [**UploadPhoto**](docs/MediaApi.md#uploadphoto) | **Post** /api/gphotos/v1/albums/{album_id}/photos | 
*PermissionsApi* | [**GetAlbumPermissions**](docs/PermissionsApi.md#getalbumpermissions) | **Get** /api/gphotos/v1/albums/{album_id}/permissions | 
*PermissionsApi* | [**RemoveAlbumPermissions**](docs/PermissionsApi.md#removealbumpermissions) | **Delete** /api/gphotos/v1/albums/{album_id}/permissions | 
*PermissionsApi* | [**SetAlbumPermissions**](docs/PermissionsApi.md#setalbumpermissions) | **Post** /api/gphotos/v1/albums/{album_id}/permissions | 
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /api/gphotos/v1/tags | 
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /api/gphotos/v1/tags/{tag_id} | 
*TagsApi* | [**GetTags**](docs/TagsApi.md#gettags) | **Get** /api/gphotos/v1/tags | 
*TagsApi* | [**RemoveTagFromAlbum**](docs/TagsApi.md#removetagfromalbum) | **Delete** /api/gphotos/v1/albums/{album_id}/tags/{tag_id} | 
*TagsApi* | [**SetTagToAlbum**](docs/TagsApi.md#settagtoalbum) | **Post** /api/gphotos/v1/albums/{album_id}/tags/{tag_id} | 
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Patch** /api/gphotos/v1/tags/{tag_id} | 
*UsersApi* | [**GetRelatedGroups**](docs/UsersApi.md#getrelatedgroups) | **Get** /api/gphotos/v1/users/{user_id}/groups/related | 
*UsersApi* | [**GetRelatedUsers**](docs/UsersApi.md#getrelatedusers) | **Get** /api/gphotos/v1/users/{user_id}/related | 
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /api/gphotos/v1/users | 
*VersionsApi* | [**GetVersionMetadata**](docs/VersionsApi.md#getversionmetadata) | **Get** /api/gphotos/v1 | 


## Documentation For Models

 - [Album](docs/Album.md)
 - [AlbumAllOf](docs/AlbumAllOf.md)
 - [AlbumList](docs/AlbumList.md)
 - [AlbumListAllOf](docs/AlbumListAllOf.md)
 - [AlbumPermissions](docs/AlbumPermissions.md)
 - [AlbumPermissionsAllOf](docs/AlbumPermissionsAllOf.md)
 - [AlbumRequestPayload](docs/AlbumRequestPayload.md)
 - [Error](docs/Error.md)
 - [ErrorAllOf](docs/ErrorAllOf.md)
 - [Group](docs/Group.md)
 - [GroupAllOf](docs/GroupAllOf.md)
 - [GroupList](docs/GroupList.md)
 - [GroupListAllOf](docs/GroupListAllOf.md)
 - [List](docs/List.md)
 - [ObjectReference](docs/ObjectReference.md)
 - [Permissions](docs/Permissions.md)
 - [Photo](docs/Photo.md)
 - [PhotoAllOf](docs/PhotoAllOf.md)
 - [PhotoList](docs/PhotoList.md)
 - [PhotoListAllOf](docs/PhotoListAllOf.md)
 - [Tag](docs/Tag.md)
 - [TagAllOf](docs/TagAllOf.md)
 - [TagList](docs/TagList.md)
 - [TagListAllOf](docs/TagListAllOf.md)
 - [TagRequestPayload](docs/TagRequestPayload.md)
 - [User](docs/User.md)
 - [UserAllOf](docs/UserAllOf.md)
 - [UserList](docs/UserList.md)
 - [UserListAllOf](docs/UserListAllOf.md)
 - [VersionMetadata](docs/VersionMetadata.md)
 - [VersionMetadataAllOf](docs/VersionMetadataAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author


