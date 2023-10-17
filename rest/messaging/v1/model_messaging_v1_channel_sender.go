/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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
// MessagingV1ChannelSender struct for MessagingV1ChannelSender
type MessagingV1ChannelSender struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ChannelSender resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the [Service](https://www.twilio.com/docs/messaging/services) the resource is associated with.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
		// The unique string that we created to identify the ChannelSender resource.
	Sid *string `json:"sid,omitempty"`
		// The unique string that identifies the sender e.g whatsapp:+123456XXXX.
	Sender *string `json:"sender,omitempty"`
		// A string value that identifies the sender type e.g WhatsApp, Messenger.
	SenderType *string `json:"sender_type,omitempty"`
		// The 2-character [ISO Country Code](https://www.iso.org/iso-3166-country-codes.html) of the number.
	CountryCode *string `json:"country_code,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Url *string `json:"url,omitempty"`
}


