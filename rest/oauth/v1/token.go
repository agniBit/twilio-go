/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Oauth
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


// Optional parameters for the method 'CreateToken'
type CreateTokenParams struct {
    // Grant type is a credential representing resource owner's authorization which can be used by client to obtain access token.
    GrantType *string `json:"GrantType,omitempty"`
    // A 34 character string that uniquely identifies this OAuth App.
    ClientSid *string `json:"ClientSid,omitempty"`
    // The credential for confidential OAuth App.
    ClientSecret *string `json:"ClientSecret,omitempty"`
    // JWT token related to the authorization code grant type.
    Code *string `json:"Code,omitempty"`
    // A code which is generation cryptographically.
    CodeVerifier *string `json:"CodeVerifier,omitempty"`
    // JWT token related to the device code grant type.
    DeviceCode *string `json:"DeviceCode,omitempty"`
    // JWT token related to the refresh token grant type.
    RefreshToken *string `json:"RefreshToken,omitempty"`
    // The Id of the device associated with the token (refresh token).
    DeviceId *string `json:"DeviceId,omitempty"`
}

func (params *CreateTokenParams) SetGrantType(GrantType string) (*CreateTokenParams){
    params.GrantType = &GrantType
    return params
}
func (params *CreateTokenParams) SetClientSid(ClientSid string) (*CreateTokenParams){
    params.ClientSid = &ClientSid
    return params
}
func (params *CreateTokenParams) SetClientSecret(ClientSecret string) (*CreateTokenParams){
    params.ClientSecret = &ClientSecret
    return params
}
func (params *CreateTokenParams) SetCode(Code string) (*CreateTokenParams){
    params.Code = &Code
    return params
}
func (params *CreateTokenParams) SetCodeVerifier(CodeVerifier string) (*CreateTokenParams){
    params.CodeVerifier = &CodeVerifier
    return params
}
func (params *CreateTokenParams) SetDeviceCode(DeviceCode string) (*CreateTokenParams){
    params.DeviceCode = &DeviceCode
    return params
}
func (params *CreateTokenParams) SetRefreshToken(RefreshToken string) (*CreateTokenParams){
    params.RefreshToken = &RefreshToken
    return params
}
func (params *CreateTokenParams) SetDeviceId(DeviceId string) (*CreateTokenParams){
    params.DeviceId = &DeviceId
    return params
}

// Issues a new Access token (optionally identity_token & refresh_token) in exchange of Oauth grant
func (c *ApiService) CreateToken(params *CreateTokenParams) (*OauthV1Token, error) {
    path := "/v1/token"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.GrantType != nil {
    data.Set("GrantType", *params.GrantType)
}
if params != nil && params.ClientSid != nil {
    data.Set("ClientSid", *params.ClientSid)
}
if params != nil && params.ClientSecret != nil {
    data.Set("ClientSecret", *params.ClientSecret)
}
if params != nil && params.Code != nil {
    data.Set("Code", *params.Code)
}
if params != nil && params.CodeVerifier != nil {
    data.Set("CodeVerifier", *params.CodeVerifier)
}
if params != nil && params.DeviceCode != nil {
    data.Set("DeviceCode", *params.DeviceCode)
}
if params != nil && params.RefreshToken != nil {
    data.Set("RefreshToken", *params.RefreshToken)
}
if params != nil && params.DeviceId != nil {
    data.Set("DeviceId", *params.DeviceId)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &OauthV1Token{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
