/*
 * Twilio - Sync
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

// SyncV1ServiceSyncListSyncListItem struct for SyncV1ServiceSyncListSyncListItem
type SyncV1ServiceSyncListSyncListItem struct {
	AccountSid  *string                 `json:"AccountSid,omitempty"`
	CreatedBy   *string                 `json:"CreatedBy,omitempty"`
	Data        *map[string]interface{} `json:"Data,omitempty"`
	DateCreated *time.Time              `json:"DateCreated,omitempty"`
	DateExpires *time.Time              `json:"DateExpires,omitempty"`
	DateUpdated *time.Time              `json:"DateUpdated,omitempty"`
	Index       *int32                  `json:"Index,omitempty"`
	ListSid     *string                 `json:"ListSid,omitempty"`
	Revision    *string                 `json:"Revision,omitempty"`
	ServiceSid  *string                 `json:"ServiceSid,omitempty"`
	Url         *string                 `json:"Url,omitempty"`
}