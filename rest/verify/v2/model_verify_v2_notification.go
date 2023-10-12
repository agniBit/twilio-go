/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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
// VerifyV2Notification struct for VerifyV2Notification
type VerifyV2Notification struct {
		// A 34 character string that uniquely identifies this Notification.
	Sid *string `json:"sid,omitempty"`
		// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique SID identifier of the Service.
	ServiceSid *string `json:"service_sid,omitempty"`
		// The unique SID identifier of the Entity.
	EntitySid *string `json:"entity_sid,omitempty"`
		// Customer unique identity for the Entity owner of the Challenge. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.
	Identity *string `json:"identity,omitempty"`
		// The unique SID identifier of the Challenge.
	ChallengeSid *string `json:"challenge_sid,omitempty"`
		// The priority of the notification. For `push` Challenges it's always `high` which sends the notification immediately, and can wake up a sleeping device.
	Priority *string `json:"priority,omitempty"`
		// How long, in seconds, the notification is valid. Max: 5 minutes
	Ttl *int `json:"ttl,omitempty"`
		// The date that this Notification was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
}


