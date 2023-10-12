/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi
import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
)
// TaskrouterV1TaskQueueRealTimeStatistics struct for TaskrouterV1TaskQueueRealTimeStatistics
type TaskrouterV1TaskQueueRealTimeStatistics struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The number of current Workers by Activity.
	ActivityStatistics *[]interface{} `json:"activity_statistics,omitempty"`
		// The age of the longest waiting Task.
	LongestTaskWaitingAge *int `json:"longest_task_waiting_age,omitempty"`
		// The SID of the longest waiting Task.
	LongestTaskWaitingSid *string `json:"longest_task_waiting_sid,omitempty"`
		// The relative age in the TaskQueue for the longest waiting Task. Calculation is based on the time when the Task entered the TaskQueue.
	LongestRelativeTaskAgeInQueue *int `json:"longest_relative_task_age_in_queue,omitempty"`
		// The Task SID of the Task waiting in the TaskQueue the longest. Calculation is based on the time when the Task entered the TaskQueue.
	LongestRelativeTaskSidInQueue *string `json:"longest_relative_task_sid_in_queue,omitempty"`
		// The SID of the TaskQueue from which these statistics were calculated.
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
		// The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99.
	TasksByPriority *interface{} `json:"tasks_by_priority,omitempty"`
		// The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`.
	TasksByStatus *interface{} `json:"tasks_by_status,omitempty"`
		// The total number of Workers available for Tasks in the TaskQueue.
	TotalAvailableWorkers *int `json:"total_available_workers,omitempty"`
		// The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state.
	TotalEligibleWorkers *int `json:"total_eligible_workers,omitempty"`
		// The total number of Tasks.
	TotalTasks *int `json:"total_tasks,omitempty"`
		// The SID of the Workspace that contains the TaskQueue.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
		// The absolute URL of the TaskQueue statistics resource.
	Url *string `json:"url,omitempty"`
}


