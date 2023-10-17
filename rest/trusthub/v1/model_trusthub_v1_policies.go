/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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
// TrusthubV1Policies struct for TrusthubV1Policies
type TrusthubV1Policies struct {
		// The unique string that identifies the Policy resource.
	Sid *string `json:"sid,omitempty"`
		// A human-readable description that is assigned to describe the Policy resource. Examples can include Primary Customer profile policy
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The SID of an object that holds the policy information
	Requirements *interface{} `json:"requirements,omitempty"`
		// The absolute URL of the Policy resource.
	Url *string `json:"url,omitempty"`
}


