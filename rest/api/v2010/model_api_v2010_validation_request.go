/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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
// ApiV2010ValidationRequest struct for ApiV2010ValidationRequest
type ApiV2010ValidationRequest struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the Caller ID.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Caller ID is associated with.
	CallSid *string `json:"call_sid,omitempty"`
		// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"phone_number,omitempty"`
		// The 6 digit validation code that someone must enter to validate the Caller ID  when `phone_number` is called.
	ValidationCode *string `json:"validation_code,omitempty"`
}


