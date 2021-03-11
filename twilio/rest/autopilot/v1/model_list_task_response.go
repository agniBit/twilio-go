/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTaskResponse struct for ListTaskResponse
type ListTaskResponse struct {
	Meta  ListAssistantResponseMeta  `json:"Meta,omitempty"`
	Tasks []AutopilotV1AssistantTask `json:"Tasks,omitempty"`
}