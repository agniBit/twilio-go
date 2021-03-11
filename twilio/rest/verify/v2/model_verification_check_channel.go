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

// VerificationCheckChannel the model 'VerificationCheckChannel'
type VerificationCheckChannel string

// List of verification_check_channel
const (
	VERIFICATIONCHECKCHANNEL_SMS   VerificationCheckChannel = "sms"
	VERIFICATIONCHECKCHANNEL_CALL  VerificationCheckChannel = "call"
	VERIFICATIONCHECKCHANNEL_EMAIL VerificationCheckChannel = "email"
)