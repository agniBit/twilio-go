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
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// 
func (c *ApiService) FetchSipDomain(SipDomain string, ) (*RoutesV2SipDomain, error) {
    path := "/v2/SipDomains/{SipDomain}"
        path = strings.Replace(path, "{"+"SipDomain"+"}", SipDomain, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &RoutesV2SipDomain{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateSipDomain'
type UpdateSipDomainParams struct {
    // 
    VoiceRegion *string `json:"VoiceRegion,omitempty"`
    // 
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSipDomainParams) SetVoiceRegion(VoiceRegion string) (*UpdateSipDomainParams){
    params.VoiceRegion = &VoiceRegion
    return params
}
func (params *UpdateSipDomainParams) SetFriendlyName(FriendlyName string) (*UpdateSipDomainParams){
    params.FriendlyName = &FriendlyName
    return params
}

// 
func (c *ApiService) UpdateSipDomain(SipDomain string, params *UpdateSipDomainParams) (*RoutesV2SipDomain, error) {
    path := "/v2/SipDomains/{SipDomain}"
        path = strings.Replace(path, "{"+"SipDomain"+"}", SipDomain, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.VoiceRegion != nil {
    data.Set("VoiceRegion", *params.VoiceRegion)
}
if params != nil && params.FriendlyName != nil {
    data.Set("FriendlyName", *params.FriendlyName)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &RoutesV2SipDomain{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
