/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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


// Optional parameters for the method 'FetchAccountSettings'
type FetchAccountSettingsParams struct {
    // The unique SID identifier of the Subaccount.
    SubaccountSid *string `json:"SubaccountSid,omitempty"`
}

func (params *FetchAccountSettingsParams) SetSubaccountSid(SubaccountSid string) (*FetchAccountSettingsParams){
    params.SubaccountSid = &SubaccountSid
    return params
}

// Get the Voice Insights Settings.
func (c *ApiService) FetchAccountSettings(params *FetchAccountSettingsParams) (*InsightsV1AccountSettings, error) {
    path := "/v1/Voice/Settings"
    
    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.SubaccountSid != nil {
        data.Set("SubaccountSid", *params.SubaccountSid)
    }



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InsightsV1AccountSettings{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateAccountSettings'
type UpdateAccountSettingsParams struct {
    // A boolean flag to enable Advanced Features for Voice Insights.
    AdvancedFeatures *bool `json:"AdvancedFeatures,omitempty"`
    // A boolean flag to enable Voice Trace.
    VoiceTrace *bool `json:"VoiceTrace,omitempty"`
    // The unique SID identifier of the Subaccount.
    SubaccountSid *string `json:"SubaccountSid,omitempty"`
}

func (params *UpdateAccountSettingsParams) SetAdvancedFeatures(AdvancedFeatures bool) (*UpdateAccountSettingsParams){
    params.AdvancedFeatures = &AdvancedFeatures
    return params
}
func (params *UpdateAccountSettingsParams) SetVoiceTrace(VoiceTrace bool) (*UpdateAccountSettingsParams){
    params.VoiceTrace = &VoiceTrace
    return params
}
func (params *UpdateAccountSettingsParams) SetSubaccountSid(SubaccountSid string) (*UpdateAccountSettingsParams){
    params.SubaccountSid = &SubaccountSid
    return params
}

// Update a specific Voice Insights Setting.
func (c *ApiService) UpdateAccountSettings(params *UpdateAccountSettingsParams) (*InsightsV1AccountSettings, error) {
    path := "/v1/Voice/Settings"
    
    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.AdvancedFeatures != nil {
        data.Set("AdvancedFeatures", fmt.Sprint(*params.AdvancedFeatures))
    }
    if params != nil && params.VoiceTrace != nil {
        data.Set("VoiceTrace", fmt.Sprint(*params.VoiceTrace))
    }
    if params != nil && params.SubaccountSid != nil {
        data.Set("SubaccountSid", *params.SubaccountSid)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InsightsV1AccountSettings{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
