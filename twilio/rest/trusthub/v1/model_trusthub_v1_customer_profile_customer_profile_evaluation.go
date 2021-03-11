/*
 * Twilio - Trusthub
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

// TrusthubV1CustomerProfileCustomerProfileEvaluation struct for TrusthubV1CustomerProfileCustomerProfileEvaluation
type TrusthubV1CustomerProfileCustomerProfileEvaluation struct {
	AccountSid         *string                          `json:"AccountSid,omitempty"`
	CustomerProfileSid *string                          `json:"CustomerProfileSid,omitempty"`
	DateCreated        *time.Time                       `json:"DateCreated,omitempty"`
	PolicySid          *string                          `json:"PolicySid,omitempty"`
	Results            *[]map[string]interface{}        `json:"Results,omitempty"`
	Sid                *string                          `json:"Sid,omitempty"`
	Status             *CustomerProfileEvaluationStatus `json:"Status,omitempty"`
	Url                *string                          `json:"Url,omitempty"`
}