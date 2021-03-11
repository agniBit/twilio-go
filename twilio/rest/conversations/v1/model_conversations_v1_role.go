/*
 * Twilio - Conversations
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

// ConversationsV1Role struct for ConversationsV1Role
type ConversationsV1Role struct {
	AccountSid     *string       `json:"AccountSid,omitempty"`
	ChatServiceSid *string       `json:"ChatServiceSid,omitempty"`
	DateCreated    *time.Time    `json:"DateCreated,omitempty"`
	DateUpdated    *time.Time    `json:"DateUpdated,omitempty"`
	FriendlyName   *string       `json:"FriendlyName,omitempty"`
	Permissions    *[]string     `json:"Permissions,omitempty"`
	Sid            *string       `json:"Sid,omitempty"`
	Type           *RoleRoleType `json:"Type,omitempty"`
	Url            *string       `json:"Url,omitempty"`
}