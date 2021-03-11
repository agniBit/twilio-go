/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ServiceNumberSelectionBehavior the model 'ServiceNumberSelectionBehavior'
type ServiceNumberSelectionBehavior string

// List of service_number_selection_behavior
const (
	SERVICENUMBERSELECTIONBEHAVIOR_AVOID_STICKY  ServiceNumberSelectionBehavior = "avoid-sticky"
	SERVICENUMBERSELECTIONBEHAVIOR_PREFER_STICKY ServiceNumberSelectionBehavior = "prefer-sticky"
)