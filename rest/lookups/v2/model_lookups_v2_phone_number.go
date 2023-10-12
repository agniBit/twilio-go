/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Lookups
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
// LookupsV2PhoneNumber struct for LookupsV2PhoneNumber
type LookupsV2PhoneNumber struct {
		// International dialing prefix of the phone number defined in the E.164 standard.
	CallingCountryCode *string `json:"calling_country_code,omitempty"`
		// The phone number's [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string `json:"country_code,omitempty"`
		// The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"phone_number,omitempty"`
		// The phone number in [national format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers).
	NationalFormat *string `json:"national_format,omitempty"`
		// Boolean which indicates if the phone number is in a valid range that can be freely assigned by a carrier to a user.
	Valid *bool `json:"valid,omitempty"`
		// Contains reasons why a phone number is invalid. Possible values: TOO_SHORT, TOO_LONG, INVALID_BUT_POSSIBLE, INVALID_COUNTRY_CODE, INVALID_LENGTH, NOT_A_NUMBER.
	ValidationErrors *[]string `json:"validation_errors,omitempty"`
		// An object that contains caller name information based on [CNAM](https://support.twilio.com/hc/en-us/articles/360051670533-Getting-Started-with-CNAM-Caller-ID).
	CallerName *interface{} `json:"caller_name,omitempty"`
		// An object that contains information on the last date the subscriber identity module (SIM) was changed for a mobile phone number.
	SimSwap *interface{} `json:"sim_swap,omitempty"`
		// An object that contains information on the unconditional call forwarding status of mobile phone number.
	CallForwarding *interface{} `json:"call_forwarding,omitempty"`
		// An object that contains live activity information for a mobile phone number.
	LiveActivity *interface{} `json:"live_activity,omitempty"`
		// An object that contains line type information including the carrier name, mobile country code, and mobile network code.
	LineTypeIntelligence *interface{} `json:"line_type_intelligence,omitempty"`
		// An object that contains identity match information. The result of comparing user-provided information including name, address, date of birth, national ID, against authoritative phone-based data sources
	IdentityMatch *interface{} `json:"identity_match,omitempty"`
		// An object that contains reassigned number information. Reassigned Numbers will return a phone number's reassignment status given a phone number and date
	ReassignedNumber *interface{} `json:"reassigned_number,omitempty"`
		// An object that contains information on if a phone number has been currently or previously blocked by Verify Fraud Guard for receiving malicious SMS pumping traffic as well as other signals associated with risky carriers and low conversion rates.
	SmsPumpingRisk *interface{} `json:"sms_pumping_risk,omitempty"`
		// An object that contains information on if a mobile phone number could be a disposable or burner number.
	DisposablePhoneNumberRisk *interface{} `json:"disposable_phone_number_risk,omitempty"`
		// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}


