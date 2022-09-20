/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Routes
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// RoutesV2Trunks struct for RoutesV2Trunks
type RoutesV2Trunks struct {
	// The SIP Trunk
	SipTrunkDomain *string `json:"sip_trunk_domain,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// A string that uniquely identifies the Inbound Processing Region assignments for this SIP Trunk.
	Sid *string `json:"sid,omitempty"`
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// A human readable description of the Inbound Processing Region assignments for this SIP Trunk.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The Inbound Processing Region used for this SIP Trunk for voice.
	VoiceRegion *string `json:"voice_region,omitempty"`
	// The date that this SIP Trunk was assigned an Inbound Processing Region.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that the Inbound Processing Region was updated for this SIP Trunk.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}