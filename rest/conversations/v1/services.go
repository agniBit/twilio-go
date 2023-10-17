/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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


// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
    // The human-readable name of this service, limited to 256 characters. Optional.
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) (*CreateServiceParams){
    params.FriendlyName = &FriendlyName
    return params
}

// Create a new conversation service on your account
func (c *ApiService) CreateService(params *CreateServiceParams) (*ConversationsV1Service, error) {
    path := "/v1/Services"
    
    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1Service{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Remove a conversation service with all its nested resources from your account
func (c *ApiService) DeleteService(Sid string, ) (error) {
    path := "/v1/Services/{Sid}"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})




    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Fetch a conversation service from your account
func (c *ApiService) FetchService(Sid string, ) (*ConversationsV1Service, error) {
    path := "/v1/Services/{Sid}"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1Service{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) (*ListServiceParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListServiceParams) SetLimit(Limit int) (*ListServiceParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
    path := "/v1/Services"

    
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

    ps := &ListServiceResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]ConversationsV1Service, error) {
	response, errors := c.StreamService(params)

	records := make([]ConversationsV1Service, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan ConversationsV1Service, chan error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1Service, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageService(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamService(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamService(response *ListServiceResponse, params *ListServiceParams, recordChannel chan ConversationsV1Service, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Services
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListServiceResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

