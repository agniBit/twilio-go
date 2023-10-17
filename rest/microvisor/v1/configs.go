/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Microvisor
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


// Optional parameters for the method 'CreateAccountConfig'
type CreateAccountConfigParams struct {
    // The config key; up to 100 characters.
    Key *string `json:"Key,omitempty"`
    // The config value; up to 4096 characters.
    Value *string `json:"Value,omitempty"`
}

func (params *CreateAccountConfigParams) SetKey(Key string) (*CreateAccountConfigParams){
    params.Key = &Key
    return params
}
func (params *CreateAccountConfigParams) SetValue(Value string) (*CreateAccountConfigParams){
    params.Value = &Value
    return params
}

// Create a config for an Account.
func (c *ApiService) CreateAccountConfig(params *CreateAccountConfigParams) (*MicrovisorV1AccountConfig, error) {
    path := "/v1/Configs"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Key != nil {
    data.Set("Key", *params.Key)
}
if params != nil && params.Value != nil {
    data.Set("Value", *params.Value)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MicrovisorV1AccountConfig{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Delete a config for an Account.
func (c *ApiService) DeleteAccountConfig(Key string, ) (error) {
    path := "/v1/Configs/{Key}"
        path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Retrieve a Config for an Account.
func (c *ApiService) FetchAccountConfig(Key string, ) (*MicrovisorV1AccountConfig, error) {
    path := "/v1/Configs/{Key}"
        path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MicrovisorV1AccountConfig{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListAccountConfig'
type ListAccountConfigParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListAccountConfigParams) SetPageSize(PageSize int) (*ListAccountConfigParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListAccountConfigParams) SetLimit(Limit int) (*ListAccountConfigParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of AccountConfig records from the API. Request is executed immediately.
func (c *ApiService) PageAccountConfig(params *ListAccountConfigParams, pageToken, pageNumber string) (*ListAccountConfigResponse, error) {
    path := "/v1/Configs"

    
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

    ps := &ListAccountConfigResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists AccountConfig records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAccountConfig(params *ListAccountConfigParams) ([]MicrovisorV1AccountConfig, error) {
	response, errors := c.StreamAccountConfig(params)

	records := make([]MicrovisorV1AccountConfig, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AccountConfig records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAccountConfig(params *ListAccountConfigParams) (chan MicrovisorV1AccountConfig, chan error) {
	if params == nil {
		params = &ListAccountConfigParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MicrovisorV1AccountConfig, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAccountConfig(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAccountConfig(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamAccountConfig(response *ListAccountConfigResponse, params *ListAccountConfigParams, recordChannel chan MicrovisorV1AccountConfig, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Configs
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAccountConfigResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAccountConfigResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAccountConfigResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListAccountConfigResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateAccountConfig'
type UpdateAccountConfigParams struct {
    // The config value; up to 4096 characters.
    Value *string `json:"Value,omitempty"`
}

func (params *UpdateAccountConfigParams) SetValue(Value string) (*UpdateAccountConfigParams){
    params.Value = &Value
    return params
}

// Update a config for an Account.
func (c *ApiService) UpdateAccountConfig(Key string, params *UpdateAccountConfigParams) (*MicrovisorV1AccountConfig, error) {
    path := "/v1/Configs/{Key}"
        path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Value != nil {
    data.Set("Value", *params.Value)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MicrovisorV1AccountConfig{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
