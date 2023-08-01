/*
 * gphotos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: #VERSION#
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// TagAllOf struct for TagAllOf
type TagAllOf struct {
	Albums []ObjectReference `json:"albums"`
	User   ObjectReference   `json:"user"`
	// name of the tag
	Name string `json:"name"`
	// color of the tag in hex format
	Color string `json:"color,omitempty"`
}
