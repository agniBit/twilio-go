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
	"time"
)
// TrusthubV1CustomerProfileEntityAssignment struct for TrusthubV1CustomerProfileEntityAssignment
type TrusthubV1CustomerProfileEntityAssignment struct {
		// The unique string that we created to identify the Item Assignment resource.
	Sid *string `json:"sid,omitempty"`
		// The unique string that we created to identify the CustomerProfile resource.
	CustomerProfileSid *string `json:"customer_profile_sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Item Assignment resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"object_sid,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The absolute URL of the Identity resource.
	Url *string `json:"url,omitempty"`
}


