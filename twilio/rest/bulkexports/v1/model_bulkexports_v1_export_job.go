/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1ExportJob struct for BulkexportsV1ExportJob
type BulkexportsV1ExportJob struct {
	Details       *map[string]interface{} `json:"Details,omitempty"`
	Email         *string                 `json:"Email,omitempty"`
	EndDay        *string                 `json:"EndDay,omitempty"`
	FriendlyName  *string                 `json:"FriendlyName,omitempty"`
	JobSid        *string                 `json:"JobSid,omitempty"`
	ResourceType  *string                 `json:"ResourceType,omitempty"`
	StartDay      *string                 `json:"StartDay,omitempty"`
	Url           *string                 `json:"Url,omitempty"`
	WebhookMethod *string                 `json:"WebhookMethod,omitempty"`
	WebhookUrl    *string                 `json:"WebhookUrl,omitempty"`
}