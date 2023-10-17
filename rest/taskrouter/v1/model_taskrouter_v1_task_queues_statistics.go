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
// TaskrouterV1TaskQueuesStatistics struct for TaskrouterV1TaskQueuesStatistics
type TaskrouterV1TaskQueuesStatistics struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// An object that contains the cumulative statistics for the TaskQueues.
	Cumulative *interface{} `json:"cumulative,omitempty"`
		// An object that contains the real-time statistics for the TaskQueues.
	Realtime *interface{} `json:"realtime,omitempty"`
		// The SID of the TaskQueue from which these statistics were calculated.
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
		// The SID of the Workspace that contains the TaskQueues.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}


