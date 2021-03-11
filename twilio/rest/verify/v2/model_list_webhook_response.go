/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWebhookResponse struct for ListWebhookResponse
type ListWebhookResponse struct {
	Meta     ListVerificationAttemptResponseMeta `json:"Meta,omitempty"`
	Webhooks []VerifyV2ServiceWebhook            `json:"Webhooks,omitempty"`
}