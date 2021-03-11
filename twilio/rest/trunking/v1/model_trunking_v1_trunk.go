/*
 * Twilio - Trunking
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

// TrunkingV1Trunk struct for TrunkingV1Trunk
type TrunkingV1Trunk struct {
	AccountSid             *string                 `json:"AccountSid,omitempty"`
	AuthType               *string                 `json:"AuthType,omitempty"`
	AuthTypeSet            *[]string               `json:"AuthTypeSet,omitempty"`
	CnamLookupEnabled      *bool                   `json:"CnamLookupEnabled,omitempty"`
	DateCreated            *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated            *time.Time              `json:"DateUpdated,omitempty"`
	DisasterRecoveryMethod *HttpMethod             `json:"DisasterRecoveryMethod,omitempty"`
	DisasterRecoveryUrl    *string                 `json:"DisasterRecoveryUrl,omitempty"`
	DomainName             *string                 `json:"DomainName,omitempty"`
	FriendlyName           *string                 `json:"FriendlyName,omitempty"`
	Links                  *map[string]interface{} `json:"Links,omitempty"`
	Recording              *map[string]interface{} `json:"Recording,omitempty"`
	Secure                 *bool                   `json:"Secure,omitempty"`
	Sid                    *string                 `json:"Sid,omitempty"`
	TransferMode           *TrunkTransferSetting   `json:"TransferMode,omitempty"`
	Url                    *string                 `json:"Url,omitempty"`
}