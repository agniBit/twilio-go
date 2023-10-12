/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
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
// AutopilotV1Field struct for AutopilotV1Field
type AutopilotV1Field struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Field resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The Field Type of the field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the unique_name, or the SID of a custom Field Type.
	FieldType *string `json:"field_type,omitempty"`
		// The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with this Field.
	TaskSid *string `json:"task_sid,omitempty"`
		// The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource.
	AssistantSid *string `json:"assistant_sid,omitempty"`
		// The unique string that we created to identify the Field resource.
	Sid *string `json:"sid,omitempty"`
		// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"unique_name,omitempty"`
		// The absolute URL of the Field resource.
	Url *string `json:"url,omitempty"`
}


