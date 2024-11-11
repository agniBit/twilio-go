/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Assistants
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

// Attach Tool to Assistant
func (c *ApiService) CreateAssistantToolAttachment(AssistantId string, Id string) error {
	path := "/v1/Assistants/{assistantId}/Tools/{id}"
	path = strings.Replace(path, "{"+"AssistantId"+"}", AssistantId, -1)
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Detach Tool to Assistant
func (c *ApiService) DeleteAssistantToolAttachment(AssistantId string, Id string) error {
	path := "/v1/Assistants/{assistantId}/Tools/{id}"
	path = strings.Replace(path, "{"+"AssistantId"+"}", AssistantId, -1)
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

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

// Optional parameters for the method 'ListToolsByAssistant'
type ListToolsByAssistantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListToolsByAssistantParams) SetPageSize(PageSize int) *ListToolsByAssistantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListToolsByAssistantParams) SetLimit(Limit int) *ListToolsByAssistantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ToolsByAssistant records from the API. Request is executed immediately.
func (c *ApiService) PageToolsByAssistant(AssistantId string, params *ListToolsByAssistantParams, pageToken, pageNumber string) (*ListToolsByAssistantResponse, error) {
	path := "/v1/Assistants/{assistantId}/Tools"

	path = strings.Replace(path, "{"+"AssistantId"+"}", AssistantId, -1)

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

	ps := &ListToolsByAssistantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ToolsByAssistant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListToolsByAssistant(AssistantId string, params *ListToolsByAssistantParams) ([]AssistantsV1Tool, error) {
	response, errors := c.StreamToolsByAssistant(AssistantId, params)

	records := make([]AssistantsV1Tool, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ToolsByAssistant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamToolsByAssistant(AssistantId string, params *ListToolsByAssistantParams) (chan AssistantsV1Tool, chan error) {
	if params == nil {
		params = &ListToolsByAssistantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AssistantsV1Tool, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageToolsByAssistant(AssistantId, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamToolsByAssistant(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamToolsByAssistant(response *ListToolsByAssistantResponse, params *ListToolsByAssistantParams, recordChannel chan AssistantsV1Tool, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Tools
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListToolsByAssistantResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListToolsByAssistantResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListToolsByAssistantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListToolsByAssistantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}