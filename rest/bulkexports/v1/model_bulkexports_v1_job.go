/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
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
// BulkexportsV1Job struct for BulkexportsV1Job
type BulkexportsV1Job struct {
		// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
		// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The details of a job which is an object that contains an array of status grouped by `status` state.  Each `status` object has a `status` string, a count which is the number of days in that `status`, and list of days in that `status`. The day strings are in the format yyyy-MM-dd. As an example, a currently running job may have a status object for COMPLETED and a `status` object for SUBMITTED each with its own count and list of days.
	Details *interface{} `json:"details,omitempty"`
		// The start time for the export specified when creating the job
	StartDay *string `json:"start_day,omitempty"`
		// The end time for the export specified when creating the job
	EndDay *string `json:"end_day,omitempty"`
		// The job_sid returned when the export was created
	JobSid *string `json:"job_sid,omitempty"`
		// The optional webhook url called on completion
	WebhookUrl *string `json:"webhook_url,omitempty"`
		// This is the method used to call the webhook
	WebhookMethod *string `json:"webhook_method,omitempty"`
		// The optional email to send the completion notification to
	Email *string `json:"email,omitempty"`
	Url *string `json:"url,omitempty"`
		// This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease
	JobQueuePosition *string `json:"job_queue_position,omitempty"`
		// this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position
	EstimatedCompletionTime *string `json:"estimated_completion_time,omitempty"`
}


