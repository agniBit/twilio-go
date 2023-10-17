/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trunking
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
// TrunkingV1PhoneNumber struct for TrunkingV1PhoneNumber
type TrunkingV1PhoneNumber struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the PhoneNumber resource.
	AccountSid *string `json:"account_sid,omitempty"`
	AddressRequirements *string `json:"address_requirements,omitempty"`
		// The API version used to start a new TwiML session.
	ApiVersion *string `json:"api_version,omitempty"`
		// Whether the phone number is new to the Twilio platform. Can be: `true` or `false`.
	Beta *bool `json:"beta,omitempty"`
		// The set of Boolean properties that indicate whether a phone number can receive calls or messages.  Capabilities are  `Voice`, `SMS`, and `MMS` and each capability can be: `true` or `false`.
	Capabilities *map[string]interface{} `json:"capabilities,omitempty"`
		// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
		// The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"phone_number,omitempty"`
		// The unique string that we created to identify the PhoneNumber resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the application that handles SMS messages sent to the phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application.
	SmsApplicationSid *string `json:"sms_application_sid,omitempty"`
		// The HTTP method we use to call `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod *string `json:"sms_fallback_method,omitempty"`
		// The URL that we call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl *string `json:"sms_fallback_url,omitempty"`
		// The HTTP method we use to call `sms_url`. Can be: `GET` or `POST`.
	SmsMethod *string `json:"sms_method,omitempty"`
		// The URL we call using the `sms_method` when the phone number receives an incoming SMS message.
	SmsUrl *string `json:"sms_url,omitempty"`
		// The URL we call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"status_callback,omitempty"`
		// The HTTP method we use to call `status_callback`. Can be: `GET` or `POST`.
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
		// The SID of the Trunk that handles calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice URLs and voice applications and use those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
	TrunkSid *string `json:"trunk_sid,omitempty"`
		// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
		// The SID of the application that handles calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice URLs and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
	VoiceApplicationSid *string `json:"voice_application_sid,omitempty"`
		// Whether we look up the caller's caller-ID name from the CNAM database ($0.01 per look up). Can be: `true` or `false`.
	VoiceCallerIdLookup *bool `json:"voice_caller_id_lookup,omitempty"`
		// The HTTP method that we use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
		// The URL that we call using the `voice_fallback_method` when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
		// The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"voice_method,omitempty"`
		// The URL we call using the `voice_method` when the phone number receives a call. The `voice_url` is not be used if a `voice_application_sid` or a `trunk_sid` is set.
	VoiceUrl *string `json:"voice_url,omitempty"`
}


