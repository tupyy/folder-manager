/*
 * gphotos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: #VERSION#
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// Permissions struct for Permissions
type Permissions struct {
	Owner       ObjectReference `json:"owner"`
	Permissions []string        `json:"permissions"`
}
