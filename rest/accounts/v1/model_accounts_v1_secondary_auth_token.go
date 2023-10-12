/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Accounts
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
// AccountsV1SecondaryAuthToken struct for AccountsV1SecondaryAuthToken
type AccountsV1SecondaryAuthToken struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the secondary Auth Token was created for.
	AccountSid *string `json:"account_sid,omitempty"`
		// The date and time in UTC when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in UTC when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The generated secondary Auth Token that can be used to authenticate future API requests.
	SecondaryAuthToken *string `json:"secondary_auth_token,omitempty"`
		// The URI for this resource, relative to `https://accounts.twilio.com`
	Url *string `json:"url,omitempty"`
}


