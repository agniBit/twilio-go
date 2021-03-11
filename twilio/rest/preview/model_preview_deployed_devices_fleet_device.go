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

// PreviewDeployedDevicesFleetDevice struct for PreviewDeployedDevicesFleetDevice
type PreviewDeployedDevicesFleetDevice struct {
	AccountSid        *string    `json:"AccountSid,omitempty"`
	DateAuthenticated *time.Time `json:"DateAuthenticated,omitempty"`
	DateCreated       *time.Time `json:"DateCreated,omitempty"`
	DateUpdated       *time.Time `json:"DateUpdated,omitempty"`
	DeploymentSid     *string    `json:"DeploymentSid,omitempty"`
	Enabled           *bool      `json:"Enabled,omitempty"`
	FleetSid          *string    `json:"FleetSid,omitempty"`
	FriendlyName      *string    `json:"FriendlyName,omitempty"`
	Identity          *string    `json:"Identity,omitempty"`
	Sid               *string    `json:"Sid,omitempty"`
	UniqueName        *string    `json:"UniqueName,omitempty"`
	Url               *string    `json:"Url,omitempty"`
}