/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
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
// ChatV1Credential struct for ChatV1Credential
type ChatV1Credential struct {
		// The unique string that we created to identify the Credential resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Credential resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Type *string `json:"type,omitempty"`
		// [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
	Sandbox *string `json:"sandbox,omitempty"`
		// The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The absolute URL of the Credential resource.
	Url *string `json:"url,omitempty"`
}


