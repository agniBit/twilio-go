/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct for TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics
type TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct {
	AccountSid   *string                 `json:"AccountSid,omitempty"`
	Cumulative   *map[string]interface{} `json:"Cumulative,omitempty"`
	Url          *string                 `json:"Url,omitempty"`
	WorkerSid    *string                 `json:"WorkerSid,omitempty"`
	WorkspaceSid *string                 `json:"WorkspaceSid,omitempty"`
}