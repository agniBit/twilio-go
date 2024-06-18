/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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
	"strings"
)

// Optional parameters for the method 'FetchMedia'
type FetchMediaParams struct {
	// Grant access to PII Redacted/Unredacted Media. If redaction is enabled, the default is `true` to access redacted media.
	Redacted *bool `json:"Redacted,omitempty"`
}

func (params *FetchMediaParams) SetRedacted(Redacted bool) *FetchMediaParams {
	params.Redacted = &Redacted
	return params
}

// Get download URLs for media if possible
func (c *ApiService) FetchMedia(Sid string, params *FetchMediaParams) (*IntelligenceV2Media, error) {
	path := "/v2/Transcripts/{Sid}/Media"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Redacted != nil {
		data.Set("Redacted", fmt.Sprint(*params.Redacted))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IntelligenceV2Media{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
