/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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


// Fetch a specific End-User Type Instance.
func (c *ApiService) FetchEndUserType(Sid string, ) (*TrusthubV1EndUserType, error) {
    path := "/v1/EndUserTypes/{Sid}"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TrusthubV1EndUserType{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListEndUserType'
type ListEndUserTypeParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListEndUserTypeParams) SetPageSize(PageSize int) (*ListEndUserTypeParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListEndUserTypeParams) SetLimit(Limit int) (*ListEndUserTypeParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of EndUserType records from the API. Request is executed immediately.
func (c *ApiService) PageEndUserType(params *ListEndUserTypeParams, pageToken, pageNumber string) (*ListEndUserTypeResponse, error) {
    path := "/v1/EndUserTypes"

    
data := url.Values{}
headers := make(map[string]interface{})

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

    ps := &ListEndUserTypeResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists EndUserType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEndUserType(params *ListEndUserTypeParams) ([]TrusthubV1EndUserType, error) {
	response, errors := c.StreamEndUserType(params)

	records := make([]TrusthubV1EndUserType, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams EndUserType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEndUserType(params *ListEndUserTypeParams) (chan TrusthubV1EndUserType, chan error) {
	if params == nil {
		params = &ListEndUserTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1EndUserType, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEndUserType(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEndUserType(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamEndUserType(response *ListEndUserTypeResponse, params *ListEndUserTypeParams, recordChannel chan TrusthubV1EndUserType, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.EndUserTypes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEndUserTypeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEndUserTypeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEndUserTypeResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListEndUserTypeResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

