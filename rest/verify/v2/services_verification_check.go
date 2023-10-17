/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'CreateVerificationCheck'
type CreateVerificationCheckParams struct {
    // The 4-10 character string being verified.
    Code *string `json:"Code,omitempty"`
    // The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the `verification_sid` must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
    To *string `json:"To,omitempty"`
    // A SID that uniquely identifies the Verification Check. Either this parameter or the `to` phone number/[email](https://www.twilio.com/docs/verify/email) must be specified.
    VerificationSid *string `json:"VerificationSid,omitempty"`
    // The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
    Amount *string `json:"Amount,omitempty"`
    // The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
    Payee *string `json:"Payee,omitempty"`
}

func (params *CreateVerificationCheckParams) SetCode(Code string) (*CreateVerificationCheckParams){
    params.Code = &Code
    return params
}
func (params *CreateVerificationCheckParams) SetTo(To string) (*CreateVerificationCheckParams){
    params.To = &To
    return params
}
func (params *CreateVerificationCheckParams) SetVerificationSid(VerificationSid string) (*CreateVerificationCheckParams){
    params.VerificationSid = &VerificationSid
    return params
}
func (params *CreateVerificationCheckParams) SetAmount(Amount string) (*CreateVerificationCheckParams){
    params.Amount = &Amount
    return params
}
func (params *CreateVerificationCheckParams) SetPayee(Payee string) (*CreateVerificationCheckParams){
    params.Payee = &Payee
    return params
}

// challenge a specific Verification Check.
func (c *ApiService) CreateVerificationCheck(ServiceSid string, params *CreateVerificationCheckParams) (*VerifyV2VerificationCheck, error) {
    path := "/v2/Services/{ServiceSid}/VerificationCheck"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.Code != nil {
        data.Set("Code", *params.Code)
    }
    if params != nil && params.To != nil {
        data.Set("To", *params.To)
    }
    if params != nil && params.VerificationSid != nil {
        data.Set("VerificationSid", *params.VerificationSid)
    }
    if params != nil && params.Amount != nil {
        data.Set("Amount", *params.Amount)
    }
    if params != nil && params.Payee != nil {
        data.Set("Payee", *params.Payee)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VerifyV2VerificationCheck{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
