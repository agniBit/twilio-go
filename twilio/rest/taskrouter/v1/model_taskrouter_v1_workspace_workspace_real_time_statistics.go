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

// TaskrouterV1WorkspaceWorkspaceRealTimeStatistics struct for TaskrouterV1WorkspaceWorkspaceRealTimeStatistics
type TaskrouterV1WorkspaceWorkspaceRealTimeStatistics struct {
	AccountSid            *string                   `json:"AccountSid,omitempty"`
	ActivityStatistics    *[]map[string]interface{} `json:"ActivityStatistics,omitempty"`
	LongestTaskWaitingAge *int32                    `json:"LongestTaskWaitingAge,omitempty"`
	LongestTaskWaitingSid *string                   `json:"LongestTaskWaitingSid,omitempty"`
	TasksByPriority       *map[string]interface{}   `json:"TasksByPriority,omitempty"`
	TasksByStatus         *map[string]interface{}   `json:"TasksByStatus,omitempty"`
	TotalTasks            *int32                    `json:"TotalTasks,omitempty"`
	TotalWorkers          *int32                    `json:"TotalWorkers,omitempty"`
	Url                   *string                   `json:"Url,omitempty"`
	WorkspaceSid          *string                   `json:"WorkspaceSid,omitempty"`
}