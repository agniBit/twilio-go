/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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
func (c *ApiService) FetchChannelSender(MessagingServiceSid string, Sid string, ) (*MessagingV1ChannelSender, error) {
    path := "/v1/Services/{MessagingServiceSid}/ChannelSenders/{Sid}"
        path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MessagingV1ChannelSender{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListChannelSender'
type ListChannelSenderParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListChannelSenderParams) SetPageSize(PageSize int) (*ListChannelSenderParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListChannelSenderParams) SetLimit(Limit int) (*ListChannelSenderParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of ChannelSender records from the API. Request is executed immediately.
func (c *ApiService) PageChannelSender(MessagingServiceSid string, params *ListChannelSenderParams, pageToken, pageNumber string) (*ListChannelSenderResponse, error) {
    path := "/v1/Services/{MessagingServiceSid}/ChannelSenders"

        path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

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

    ps := &ListChannelSenderResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists ChannelSender records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannelSender(MessagingServiceSid string, params *ListChannelSenderParams) ([]MessagingV1ChannelSender, error) {
	response, errors := c.StreamChannelSender(MessagingServiceSid, params)

	records := make([]MessagingV1ChannelSender, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ChannelSender records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannelSender(MessagingServiceSid string, params *ListChannelSenderParams) (chan MessagingV1ChannelSender, chan error) {
	if params == nil {
		params = &ListChannelSenderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1ChannelSender, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageChannelSender(MessagingServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamChannelSender(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamChannelSender(response *ListChannelSenderResponse, params *ListChannelSenderParams, recordChannel chan MessagingV1ChannelSender, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Senders
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListChannelSenderResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListChannelSenderResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListChannelSenderResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListChannelSenderResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

