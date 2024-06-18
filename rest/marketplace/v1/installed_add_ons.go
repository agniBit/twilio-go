/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateMarketplaceInstalledAddOn'
type CreateMarketplaceInstalledAddOnParams struct {
	// The SID of the AvaliableAddOn to install.
	AvailableAddOnSid *string `json:"AvailableAddOnSid,omitempty"`
	// Whether the Terms of Service were accepted.
	AcceptTermsOfService *bool `json:"AcceptTermsOfService,omitempty"`
	// The JSON object that represents the configuration of the new Add-on being installed.
	Configuration *interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateMarketplaceInstalledAddOnParams) SetAvailableAddOnSid(AvailableAddOnSid string) *CreateMarketplaceInstalledAddOnParams {
	params.AvailableAddOnSid = &AvailableAddOnSid
	return params
}
func (params *CreateMarketplaceInstalledAddOnParams) SetAcceptTermsOfService(AcceptTermsOfService bool) *CreateMarketplaceInstalledAddOnParams {
	params.AcceptTermsOfService = &AcceptTermsOfService
	return params
}
func (params *CreateMarketplaceInstalledAddOnParams) SetConfiguration(Configuration interface{}) *CreateMarketplaceInstalledAddOnParams {
	params.Configuration = &Configuration
	return params
}
func (params *CreateMarketplaceInstalledAddOnParams) SetUniqueName(UniqueName string) *CreateMarketplaceInstalledAddOnParams {
	params.UniqueName = &UniqueName
	return params
}

// Install an Add-on for the Account specified.
func (c *ApiService) CreateMarketplaceInstalledAddOn(params *CreateMarketplaceInstalledAddOnParams) (*MarketplaceInstalledAddOn, error) {
	path := "/v1/InstalledAddOns"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AvailableAddOnSid != nil {
		data.Set("AvailableAddOnSid", *params.AvailableAddOnSid)
	}
	if params != nil && params.AcceptTermsOfService != nil {
		data.Set("AcceptTermsOfService", fmt.Sprint(*params.AcceptTermsOfService))
	}
	if params != nil && params.Configuration != nil {
		v, err := json.Marshal(params.Configuration)

		if err != nil {
			return nil, err
		}

		data.Set("Configuration", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceInstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Add-on installation from your account
func (c *ApiService) DeleteMarketplaceInstalledAddOn(Sid string) error {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch an instance of an Add-on currently installed on this Account.
func (c *ApiService) FetchMarketplaceInstalledAddOn(Sid string) (*MarketplaceInstalledAddOn, error) {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceInstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMarketplaceInstalledAddOn'
type ListMarketplaceInstalledAddOnParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMarketplaceInstalledAddOnParams) SetPageSize(PageSize int) *ListMarketplaceInstalledAddOnParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMarketplaceInstalledAddOnParams) SetLimit(Limit int) *ListMarketplaceInstalledAddOnParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MarketplaceInstalledAddOn records from the API. Request is executed immediately.
func (c *ApiService) PageMarketplaceInstalledAddOn(params *ListMarketplaceInstalledAddOnParams, pageToken, pageNumber string) (*ListMarketplaceInstalledAddOnResponse, error) {
	path := "/v1/InstalledAddOns"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMarketplaceInstalledAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MarketplaceInstalledAddOn records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMarketplaceInstalledAddOn(params *ListMarketplaceInstalledAddOnParams) ([]MarketplaceInstalledAddOn, error) {
	response, errors := c.StreamMarketplaceInstalledAddOn(params)

	records := make([]MarketplaceInstalledAddOn, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams MarketplaceInstalledAddOn records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMarketplaceInstalledAddOn(params *ListMarketplaceInstalledAddOnParams) (chan MarketplaceInstalledAddOn, chan error) {
	if params == nil {
		params = &ListMarketplaceInstalledAddOnParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MarketplaceInstalledAddOn, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMarketplaceInstalledAddOn(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMarketplaceInstalledAddOn(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMarketplaceInstalledAddOn(response *ListMarketplaceInstalledAddOnResponse, params *ListMarketplaceInstalledAddOnParams, recordChannel chan MarketplaceInstalledAddOn, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.InstalledAddOns
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMarketplaceInstalledAddOnResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMarketplaceInstalledAddOnResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMarketplaceInstalledAddOnResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMarketplaceInstalledAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMarketplaceInstalledAddOn'
type UpdateMarketplaceInstalledAddOnParams struct {
	// Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured
	Configuration *interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateMarketplaceInstalledAddOnParams) SetConfiguration(Configuration interface{}) *UpdateMarketplaceInstalledAddOnParams {
	params.Configuration = &Configuration
	return params
}
func (params *UpdateMarketplaceInstalledAddOnParams) SetUniqueName(UniqueName string) *UpdateMarketplaceInstalledAddOnParams {
	params.UniqueName = &UniqueName
	return params
}

// Update an Add-on installation for the Account specified.
func (c *ApiService) UpdateMarketplaceInstalledAddOn(Sid string, params *UpdateMarketplaceInstalledAddOnParams) (*MarketplaceInstalledAddOn, error) {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Configuration != nil {
		v, err := json.Marshal(params.Configuration)

		if err != nil {
			return nil, err
		}

		data.Set("Configuration", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceInstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
