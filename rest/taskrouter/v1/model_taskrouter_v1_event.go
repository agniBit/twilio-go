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
	"time"
)
// TaskrouterV1Event struct for TaskrouterV1Event
type TaskrouterV1Event struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Event resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the resource that triggered the event.
	ActorSid *string `json:"actor_sid,omitempty"`
		// The type of resource that triggered the event.
	ActorType *string `json:"actor_type,omitempty"`
		// The absolute URL of the resource that triggered the event.
	ActorUrl *string `json:"actor_url,omitempty"`
		// A description of the event.
	Description *string `json:"description,omitempty"`
		// Data about the event. For more information, see [Event types](https://www.twilio.com/docs/taskrouter/api/event#event-types).
	EventData *interface{} `json:"event_data,omitempty"`
		// The time the event was sent, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EventDate *time.Time `json:"event_date,omitempty"`
		// The time the event was sent in milliseconds.
	EventDateMs *int64 `json:"event_date_ms,omitempty"`
		// The identifier for the event.
	EventType *string `json:"event_type,omitempty"`
		// The SID of the object the event is most relevant to, such as a TaskSid, ReservationSid, or a  WorkerSid.
	ResourceSid *string `json:"resource_sid,omitempty"`
		// The type of object the event is most relevant to, such as a Task, Reservation, or a Worker).
	ResourceType *string `json:"resource_type,omitempty"`
		// The URL of the resource the event is most relevant to.
	ResourceUrl *string `json:"resource_url,omitempty"`
		// The unique string that we created to identify the Event resource.
	Sid *string `json:"sid,omitempty"`
		// Where the Event originated.
	Source *string `json:"source,omitempty"`
		// The IP from which the Event originated.
	SourceIpAddress *string `json:"source_ip_address,omitempty"`
		// The absolute URL of the Event resource.
	Url *string `json:"url,omitempty"`
		// The SID of the Workspace that contains the Event.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}


