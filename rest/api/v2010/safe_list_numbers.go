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
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'CreateSafelist'
type CreateSafelistParams struct {
    // The phone number to be added in SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
    PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

func (params *CreateSafelistParams) SetPhoneNumber(PhoneNumber string) (*CreateSafelistParams){
    params.PhoneNumber = &PhoneNumber
    return params
}

// Add a new phone number to SafeList.
func (c *ApiService) CreateSafelist(params *CreateSafelistParams) (*ApiV2010Safelist, error) {
    path := "/2010-04-01/SafeList/Numbers.json"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.PhoneNumber != nil {
    data.Set("PhoneNumber", *params.PhoneNumber)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Safelist{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'DeleteSafelist'
type DeleteSafelistParams struct {
    // The phone number to be removed from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
    PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

func (params *DeleteSafelistParams) SetPhoneNumber(PhoneNumber string) (*DeleteSafelistParams){
    params.PhoneNumber = &PhoneNumber
    return params
}

// Remove a phone number from SafeList.
func (c *ApiService) DeleteSafelist(params *DeleteSafelistParams) (error) {
    path := "/2010-04-01/SafeList/Numbers.json"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.PhoneNumber != nil {
    data.Set("PhoneNumber", *params.PhoneNumber)
}



    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Optional parameters for the method 'FetchSafelist'
type FetchSafelistParams struct {
    // The phone number to be fetched from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
    PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

func (params *FetchSafelistParams) SetPhoneNumber(PhoneNumber string) (*FetchSafelistParams){
    params.PhoneNumber = &PhoneNumber
    return params
}

// Check if a phone number exists in SafeList.
func (c *ApiService) FetchSafelist(params *FetchSafelistParams) (*ApiV2010Safelist, error) {
    path := "/2010-04-01/SafeList/Numbers.json"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.PhoneNumber != nil {
    data.Set("PhoneNumber", *params.PhoneNumber)
}



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Safelist{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
