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
// ApiV2010IncomingPhoneNumberAssignedAddOnExtension struct for ApiV2010IncomingPhoneNumberAssignedAddOnExtension
type ApiV2010IncomingPhoneNumberAssignedAddOnExtension struct {
		// The unique string that that we created to identify the resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the Phone Number to which the Add-on is assigned.
	ResourceSid *string `json:"resource_sid,omitempty"`
		// The SID that uniquely identifies the assigned Add-on installation.
	AssignedAddOnSid *string `json:"assigned_add_on_sid,omitempty"`
		// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
		// A string that you assigned to describe the Product this Extension is used within.
	ProductName *string `json:"product_name,omitempty"`
		// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"unique_name,omitempty"`
		// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
		// Whether the Extension will be invoked.
	Enabled *bool `json:"enabled,omitempty"`
}


