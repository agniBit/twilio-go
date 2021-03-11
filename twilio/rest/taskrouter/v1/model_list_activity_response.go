/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListActivityResponse struct for ListActivityResponse
type ListActivityResponse struct {
	Activities []TaskrouterV1WorkspaceActivity `json:"Activities,omitempty"`
	Meta       ListWorkspaceResponseMeta       `json:"Meta,omitempty"`
}