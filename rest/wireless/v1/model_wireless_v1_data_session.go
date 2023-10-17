/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Wireless
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
// WirelessV1DataSession struct for WirelessV1DataSession
type WirelessV1DataSession struct {
		// The unique string that we created to identify the DataSession resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the [Sim resource](https://www.twilio.com/docs/iot/wireless/api/sim-resource) that the Data Session is for.
	SimSid *string `json:"sim_sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DataSession resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The generation of wireless technology that the device was using.
	RadioLink *string `json:"radio_link,omitempty"`
		// The 'mobile country code' is the unique ID of the home country where the Data Session took place. See: [MCC/MNC lookup](http://mcc-mnc.com/).
	OperatorMcc *string `json:"operator_mcc,omitempty"`
		// The 'mobile network code' is the unique ID specific to the mobile operator network where the Data Session took place.
	OperatorMnc *string `json:"operator_mnc,omitempty"`
		// The three letter country code representing where the device's Data Session took place. This is determined by looking up the `operator_mcc`.
	OperatorCountry *string `json:"operator_country,omitempty"`
		// The friendly name of the mobile operator network that the [SIM](https://www.twilio.com/docs/iot/wireless/api/sim-resource)-connected device is attached to. This is determined by looking up the `operator_mnc`.
	OperatorName *string `json:"operator_name,omitempty"`
		// The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated.
	CellId *string `json:"cell_id,omitempty"`
		// An object that describes the estimated location in latitude and longitude where the device's Data Session took place. The location is derived from the `cell_id` when the Data Session was last updated. See [Cell Location Estimate Object](https://www.twilio.com/docs/iot/wireless/api/datasession-resource#cell-location-estimate-object). 
	CellLocationEstimate *interface{} `json:"cell_location_estimate,omitempty"`
		// The number of packets uploaded by the device between the `start` time and when the Data Session was last updated.
	PacketsUploaded *int `json:"packets_uploaded,omitempty"`
		// The number of packets downloaded by the device between the `start` time and when the Data Session was last updated.
	PacketsDownloaded *int `json:"packets_downloaded,omitempty"`
		// The date that the resource was last updated, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
		// The date that the Data Session started, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	Start *time.Time `json:"start,omitempty"`
		// The date that the record ended, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	End *time.Time `json:"end,omitempty"`
		// The 'international mobile equipment identity' is the unique ID of the device using the SIM to connect. An IMEI is a 15-digit string: 14 digits for the device identifier plus a check digit calculated using the Luhn formula.
	Imei *string `json:"imei,omitempty"`
}


