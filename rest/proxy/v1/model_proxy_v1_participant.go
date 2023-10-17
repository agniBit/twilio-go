/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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
// ProxyV1Participant struct for ProxyV1Participant
type ProxyV1Participant struct {
		// The unique string that we created to identify the Participant resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
	SessionSid *string `json:"session_sid,omitempty"`
		// The SID of the resource's parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
	ServiceSid *string `json:"service_sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The string that you assigned to describe the participant. This value must be 255 characters or fewer. Supports UTF-8 characters. **This value should not have PII.**
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The phone number or channel identifier of the Participant. This value must be 191 characters or fewer. Supports UTF-8 characters.
	Identifier *string `json:"identifier,omitempty"`
		// The phone number or short code (masked number) of the participant's partner. The participant will call or message the partner participant at this number.
	ProxyIdentifier *string `json:"proxy_identifier,omitempty"`
		// The SID of the Proxy Identifier assigned to the Participant.
	ProxyIdentifierSid *string `json:"proxy_identifier_sid,omitempty"`
		// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Participant was removed from the session.
	DateDeleted *time.Time `json:"date_deleted,omitempty"`
		// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The absolute URL of the Participant resource.
	Url *string `json:"url,omitempty"`
		// The URLs to resources related the participant.
	Links *map[string]interface{} `json:"links,omitempty"`
}


