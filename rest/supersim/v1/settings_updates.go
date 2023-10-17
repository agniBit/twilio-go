/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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


// Optional parameters for the method 'ListSettingsUpdate'
type ListSettingsUpdateParams struct {
    // Filter the Settings Updates by a Super SIM's SID or UniqueName.
    Sim *string `json:"Sim,omitempty"`
    // Filter the Settings Updates by status. Can be `scheduled`, `in-progress`, `successful`, or `failed`.
    Status *string `json:"Status,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListSettingsUpdateParams) SetSim(Sim string) (*ListSettingsUpdateParams){
    params.Sim = &Sim
    return params
}
func (params *ListSettingsUpdateParams) SetStatus(Status string) (*ListSettingsUpdateParams){
    params.Status = &Status
    return params
}
func (params *ListSettingsUpdateParams) SetPageSize(PageSize int) (*ListSettingsUpdateParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListSettingsUpdateParams) SetLimit(Limit int) (*ListSettingsUpdateParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of SettingsUpdate records from the API. Request is executed immediately.
func (c *ApiService) PageSettingsUpdate(params *ListSettingsUpdateParams, pageToken, pageNumber string) (*ListSettingsUpdateResponse, error) {
    path := "/v1/SettingsUpdates"

    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Sim != nil {
    data.Set("Sim", *params.Sim)
}
if params != nil && params.Status != nil {
    data.Set("Status", *params.Status)
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

    ps := &ListSettingsUpdateResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists SettingsUpdate records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSettingsUpdate(params *ListSettingsUpdateParams) ([]SupersimV1SettingsUpdate, error) {
	response, errors := c.StreamSettingsUpdate(params)

	records := make([]SupersimV1SettingsUpdate, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SettingsUpdate records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSettingsUpdate(params *ListSettingsUpdateParams) (chan SupersimV1SettingsUpdate, chan error) {
	if params == nil {
		params = &ListSettingsUpdateParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1SettingsUpdate, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSettingsUpdate(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSettingsUpdate(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamSettingsUpdate(response *ListSettingsUpdateResponse, params *ListSettingsUpdateParams, recordChannel chan SupersimV1SettingsUpdate, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.SettingsUpdates
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSettingsUpdateResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSettingsUpdateResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSettingsUpdateResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListSettingsUpdateResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

