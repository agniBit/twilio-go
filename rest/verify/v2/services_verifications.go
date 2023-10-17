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


// Optional parameters for the method 'CreateVerification'
type CreateVerificationParams struct {
    // The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
    To *string `json:"To,omitempty"`
    // The verification method to use. One of: [`email`](https://www.twilio.com/docs/verify/email), `sms`, `whatsapp`, `call`, `sna` or `auto`.
    Channel *string `json:"Channel,omitempty"`
    // A custom user defined friendly name that overwrites the existing one in the verification message
    CustomFriendlyName *string `json:"CustomFriendlyName,omitempty"`
    // The text of a custom message to use for the verification.
    CustomMessage *string `json:"CustomMessage,omitempty"`
    // The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits).
    SendDigits *string `json:"SendDigits,omitempty"`
    // Locale will automatically resolve based on phone number country code for SMS, WhatsApp, and call channel verifications. It will fallback to English or the template’s default translation if the selected translation is not available. This parameter will override the automatic locale resolution. [See supported languages and more information here](https://www.twilio.com/docs/verify/supported-languages).
    Locale *string `json:"Locale,omitempty"`
    // A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive.
    CustomCode *string `json:"CustomCode,omitempty"`
    // The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
    Amount *string `json:"Amount,omitempty"`
    // The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
    Payee *string `json:"Payee,omitempty"`
    // The custom key-value pairs of Programmable Rate Limits. Keys correspond to `unique_name` fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request.
    RateLimits *interface{} `json:"RateLimits,omitempty"`
    // [`email`](https://www.twilio.com/docs/verify/email) channel configuration in json format. The fields 'from' and 'from_name' are optional but if included the 'from' field must have a valid email address.
    ChannelConfiguration *interface{} `json:"ChannelConfiguration,omitempty"`
    // Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: `<#> Your AppName verification code is: 1234 He42w354ol9`.
    AppHash *string `json:"AppHash,omitempty"`
    // The message [template](https://www.twilio.com/docs/verify/api/templates). If provided, will override the default template for the Service. SMS and Voice channels only.
    TemplateSid *string `json:"TemplateSid,omitempty"`
    // A stringified JSON object in which the keys are the template's special variables and the values are the variables substitutions.
    TemplateCustomSubstitutions *string `json:"TemplateCustomSubstitutions,omitempty"`
    // Strongly encouraged if using the auto channel. The IP address of the client's device. If provided, it has to be a valid IPv4 or IPv6 address.
    DeviceIp *string `json:"DeviceIp,omitempty"`
    // 
    RiskCheck *string `json:"RiskCheck,omitempty"`
}

func (params *CreateVerificationParams) SetTo(To string) (*CreateVerificationParams){
    params.To = &To
    return params
}
func (params *CreateVerificationParams) SetChannel(Channel string) (*CreateVerificationParams){
    params.Channel = &Channel
    return params
}
func (params *CreateVerificationParams) SetCustomFriendlyName(CustomFriendlyName string) (*CreateVerificationParams){
    params.CustomFriendlyName = &CustomFriendlyName
    return params
}
func (params *CreateVerificationParams) SetCustomMessage(CustomMessage string) (*CreateVerificationParams){
    params.CustomMessage = &CustomMessage
    return params
}
func (params *CreateVerificationParams) SetSendDigits(SendDigits string) (*CreateVerificationParams){
    params.SendDigits = &SendDigits
    return params
}
func (params *CreateVerificationParams) SetLocale(Locale string) (*CreateVerificationParams){
    params.Locale = &Locale
    return params
}
func (params *CreateVerificationParams) SetCustomCode(CustomCode string) (*CreateVerificationParams){
    params.CustomCode = &CustomCode
    return params
}
func (params *CreateVerificationParams) SetAmount(Amount string) (*CreateVerificationParams){
    params.Amount = &Amount
    return params
}
func (params *CreateVerificationParams) SetPayee(Payee string) (*CreateVerificationParams){
    params.Payee = &Payee
    return params
}
func (params *CreateVerificationParams) SetRateLimits(RateLimits interface{}) (*CreateVerificationParams){
    params.RateLimits = &RateLimits
    return params
}
func (params *CreateVerificationParams) SetChannelConfiguration(ChannelConfiguration interface{}) (*CreateVerificationParams){
    params.ChannelConfiguration = &ChannelConfiguration
    return params
}
func (params *CreateVerificationParams) SetAppHash(AppHash string) (*CreateVerificationParams){
    params.AppHash = &AppHash
    return params
}
func (params *CreateVerificationParams) SetTemplateSid(TemplateSid string) (*CreateVerificationParams){
    params.TemplateSid = &TemplateSid
    return params
}
func (params *CreateVerificationParams) SetTemplateCustomSubstitutions(TemplateCustomSubstitutions string) (*CreateVerificationParams){
    params.TemplateCustomSubstitutions = &TemplateCustomSubstitutions
    return params
}
func (params *CreateVerificationParams) SetDeviceIp(DeviceIp string) (*CreateVerificationParams){
    params.DeviceIp = &DeviceIp
    return params
}
func (params *CreateVerificationParams) SetRiskCheck(RiskCheck string) (*CreateVerificationParams){
    params.RiskCheck = &RiskCheck
    return params
}

// Create a new Verification using a Service
func (c *ApiService) CreateVerification(ServiceSid string, params *CreateVerificationParams) (*VerifyV2Verification, error) {
    path := "/v2/Services/{ServiceSid}/Verifications"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.To != nil {
        data.Set("To", *params.To)
    }
    if params != nil && params.Channel != nil {
        data.Set("Channel", *params.Channel)
    }
    if params != nil && params.CustomFriendlyName != nil {
        data.Set("CustomFriendlyName", *params.CustomFriendlyName)
    }
    if params != nil && params.CustomMessage != nil {
        data.Set("CustomMessage", *params.CustomMessage)
    }
    if params != nil && params.SendDigits != nil {
        data.Set("SendDigits", *params.SendDigits)
    }
    if params != nil && params.Locale != nil {
        data.Set("Locale", *params.Locale)
    }
    if params != nil && params.CustomCode != nil {
        data.Set("CustomCode", *params.CustomCode)
    }
    if params != nil && params.Amount != nil {
        data.Set("Amount", *params.Amount)
    }
    if params != nil && params.Payee != nil {
        data.Set("Payee", *params.Payee)
    }
    if params != nil && params.RateLimits != nil {
        v, err := json.Marshal(params.RateLimits)

        if err != nil {
            return nil, err
        }

        data.Set("RateLimits", string(v))
    }
    if params != nil && params.ChannelConfiguration != nil {
        v, err := json.Marshal(params.ChannelConfiguration)

        if err != nil {
            return nil, err
        }

        data.Set("ChannelConfiguration", string(v))
    }
    if params != nil && params.AppHash != nil {
        data.Set("AppHash", *params.AppHash)
    }
    if params != nil && params.TemplateSid != nil {
        data.Set("TemplateSid", *params.TemplateSid)
    }
    if params != nil && params.TemplateCustomSubstitutions != nil {
        data.Set("TemplateCustomSubstitutions", *params.TemplateCustomSubstitutions)
    }
    if params != nil && params.DeviceIp != nil {
        data.Set("DeviceIp", *params.DeviceIp)
    }
    if params != nil && params.RiskCheck != nil {
        data.Set("RiskCheck", *params.RiskCheck)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VerifyV2Verification{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Fetch a specific Verification
func (c *ApiService) FetchVerification(ServiceSid string, Sid string, ) (*VerifyV2Verification, error) {
    path := "/v2/Services/{ServiceSid}/Verifications/{Sid}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VerifyV2Verification{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateVerification'
type UpdateVerificationParams struct {
    // 
    Status *string `json:"Status,omitempty"`
}

func (params *UpdateVerificationParams) SetStatus(Status string) (*UpdateVerificationParams){
    params.Status = &Status
    return params
}

// Update a Verification status
func (c *ApiService) UpdateVerification(ServiceSid string, Sid string, params *UpdateVerificationParams) (*VerifyV2Verification, error) {
    path := "/v2/Services/{ServiceSid}/Verifications/{Sid}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VerifyV2Verification{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
