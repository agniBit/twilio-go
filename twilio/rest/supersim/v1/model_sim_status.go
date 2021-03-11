/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SimStatus the model 'SimStatus'
type SimStatus string

// List of sim_status
const (
	SIMSTATUS_NEW       SimStatus = "new"
	SIMSTATUS_READY     SimStatus = "ready"
	SIMSTATUS_ACTIVE    SimStatus = "active"
	SIMSTATUS_INACTIVE  SimStatus = "inactive"
	SIMSTATUS_SCHEDULED SimStatus = "scheduled"
)