/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListDocumentPermissionResponse struct for ListDocumentPermissionResponse
type ListDocumentPermissionResponse struct {
	Meta        ListDayResponseMeta                            `json:"Meta,omitempty"`
	Permissions []PreviewSyncServiceDocumentDocumentPermission `json:"Permissions,omitempty"`
}