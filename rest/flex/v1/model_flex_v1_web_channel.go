/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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
// FlexV1WebChannel struct for FlexV1WebChannel
type FlexV1WebChannel struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the WebChannel resource and owns this Workflow.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the Flex Flow.
	FlexFlowSid *string `json:"flex_flow_sid,omitempty"`
		// The unique string that we created to identify the WebChannel resource.
	Sid *string `json:"sid,omitempty"`
		// The absolute URL of the WebChannel resource.
	Url *string `json:"url,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}


