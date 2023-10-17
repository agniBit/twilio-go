/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
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


// Fetch a specific Export Configuration.
func (c *ApiService) FetchExportConfiguration(ResourceType string, ) (*BulkexportsV1ExportConfiguration, error) {
    path := "/v1/Exports/{ResourceType}/Configuration"
        path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &BulkexportsV1ExportConfiguration{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateExportConfiguration'
type UpdateExportConfigurationParams struct {
    // If true, Twilio will automatically generate every day's file when the day is over.
    Enabled *bool `json:"Enabled,omitempty"`
    // Stores the URL destination for the method specified in webhook_method.
    WebhookUrl *string `json:"WebhookUrl,omitempty"`
    // Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url
    WebhookMethod *string `json:"WebhookMethod,omitempty"`
}

func (params *UpdateExportConfigurationParams) SetEnabled(Enabled bool) (*UpdateExportConfigurationParams){
    params.Enabled = &Enabled
    return params
}
func (params *UpdateExportConfigurationParams) SetWebhookUrl(WebhookUrl string) (*UpdateExportConfigurationParams){
    params.WebhookUrl = &WebhookUrl
    return params
}
func (params *UpdateExportConfigurationParams) SetWebhookMethod(WebhookMethod string) (*UpdateExportConfigurationParams){
    params.WebhookMethod = &WebhookMethod
    return params
}

// Update a specific Export Configuration.
func (c *ApiService) UpdateExportConfiguration(ResourceType string, params *UpdateExportConfigurationParams) (*BulkexportsV1ExportConfiguration, error) {
    path := "/v1/Exports/{ResourceType}/Configuration"
        path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Enabled != nil {
    data.Set("Enabled", fmt.Sprint(*params.Enabled))
}
if params != nil && params.WebhookUrl != nil {
    data.Set("WebhookUrl", *params.WebhookUrl)
}
if params != nil && params.WebhookMethod != nil {
    data.Set("WebhookMethod", *params.WebhookMethod)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &BulkexportsV1ExportConfiguration{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
