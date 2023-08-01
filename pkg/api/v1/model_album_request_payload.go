/*
 * gphotos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: #VERSION#
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// AlbumRequestPayload struct for AlbumRequestPayload
type AlbumRequestPayload struct {
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	CreatedAt        int64  `json:"created_at,omitempty"`
	Location         string `json:"location,omitempty"`
	UserPermissions  string `json:"user_permissions,omitempty"`
	GroupPermissions string `json:"group_permissions,omitempty"`
	// name of the thumbnail
	Thumbnail string `json:"thumbnail,omitempty"`
}
