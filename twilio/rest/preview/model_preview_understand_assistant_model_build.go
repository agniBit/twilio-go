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

import (
	"time"
)

// PreviewUnderstandAssistantModelBuild struct for PreviewUnderstandAssistantModelBuild
type PreviewUnderstandAssistantModelBuild struct {
	AccountSid    *string           `json:"AccountSid,omitempty"`
	AssistantSid  *string           `json:"AssistantSid,omitempty"`
	BuildDuration *int32            `json:"BuildDuration,omitempty"`
	DateCreated   *time.Time        `json:"DateCreated,omitempty"`
	DateUpdated   *time.Time        `json:"DateUpdated,omitempty"`
	ErrorCode     *int32            `json:"ErrorCode,omitempty"`
	Sid           *string           `json:"Sid,omitempty"`
	Status        *ModelBuildStatus `json:"Status,omitempty"`
	UniqueName    *string           `json:"UniqueName,omitempty"`
	Url           *string           `json:"Url,omitempty"`
}