/*
 * gphotos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: #VERSION#
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"time"
)

// AlbumAllOf struct for AlbumAllOf
type AlbumAllOf struct {
	// name of the album
	Name string `json:"name"`
	// creation date in unix timestamp
	CreatedAt time.Time `json:"created_at"`
	// description of the album
	Description string `json:"description,omitempty"`
	// location of the album
	Location string `json:"location,omitempty"`
	// path of the bucket where media is stored
	Bucket string `json:"bucket"`
	// url of the thumbnail of the album
	Thumbnail   string          `json:"thumbnail,omitempty"`
	Owner       ObjectReference `json:"owner"`
	Photos      ObjectReference `json:"photos"`
	Tags        []Tag           `json:"tags,omitempty"`
	Permissions ObjectReference `json:"permissions"`
}
