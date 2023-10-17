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


// Optional parameters for the method 'CreateDeviceSecret'
type CreateDeviceSecretParams struct {
    // The secret key; up to 100 characters.
    Key *string `json:"Key,omitempty"`
    // The secret value; up to 4096 characters.
    Value *string `json:"Value,omitempty"`
}

func (params *CreateDeviceSecretParams) SetKey(Key string) (*CreateDeviceSecretParams){
    params.Key = &Key
    return params
}
func (params *CreateDeviceSecretParams) SetValue(Value string) (*CreateDeviceSecretParams){
    params.Value = &Value
    return params
}

// Create a secret for a Microvisor Device.
func (c *ApiService) CreateDeviceSecret(DeviceSid string, params *CreateDeviceSecretParams) (*MicrovisorV1DeviceSecret, error) {
    path := "/v1/Devices/{DeviceSid}/Secrets"
        path = strings.Replace(path, "{"+"DeviceSid"+"}", DeviceSid, -1)

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

    ps := &MicrovisorV1DeviceSecret{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Delete a secret for a Microvisor Device.
func (c *ApiService) DeleteDeviceSecret(DeviceSid string, Key string, ) (error) {
    path := "/v1/Devices/{DeviceSid}/Secrets/{Key}"
        path = strings.Replace(path, "{"+"DeviceSid"+"}", DeviceSid, -1)
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

// Retrieve a Secret for a Device.
func (c *ApiService) FetchDeviceSecret(DeviceSid string, Key string, ) (*MicrovisorV1DeviceSecret, error) {
    path := "/v1/Devices/{DeviceSid}/Secrets/{Key}"
        path = strings.Replace(path, "{"+"DeviceSid"+"}", DeviceSid, -1)
    path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MicrovisorV1DeviceSecret{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListDeviceSecret'
type ListDeviceSecretParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListDeviceSecretParams) SetPageSize(PageSize int) (*ListDeviceSecretParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListDeviceSecretParams) SetLimit(Limit int) (*ListDeviceSecretParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of DeviceSecret records from the API. Request is executed immediately.
func (c *ApiService) PageDeviceSecret(DeviceSid string, params *ListDeviceSecretParams, pageToken, pageNumber string) (*ListDeviceSecretResponse, error) {
    path := "/v1/Devices/{DeviceSid}/Secrets"

        path = strings.Replace(path, "{"+"DeviceSid"+"}", DeviceSid, -1)

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

    ps := &ListDeviceSecretResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists DeviceSecret records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDeviceSecret(DeviceSid string, params *ListDeviceSecretParams) ([]MicrovisorV1DeviceSecret, error) {
	response, errors := c.StreamDeviceSecret(DeviceSid, params)

	records := make([]MicrovisorV1DeviceSecret, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DeviceSecret records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDeviceSecret(DeviceSid string, params *ListDeviceSecretParams) (chan MicrovisorV1DeviceSecret, chan error) {
	if params == nil {
		params = &ListDeviceSecretParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MicrovisorV1DeviceSecret, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDeviceSecret(DeviceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDeviceSecret(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamDeviceSecret(response *ListDeviceSecretResponse, params *ListDeviceSecretParams, recordChannel chan MicrovisorV1DeviceSecret, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Secrets
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDeviceSecretResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDeviceSecretResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDeviceSecretResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListDeviceSecretResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateDeviceSecret'
type UpdateDeviceSecretParams struct {
    // The secret value; up to 4096 characters.
    Value *string `json:"Value,omitempty"`
}

func (params *UpdateDeviceSecretParams) SetValue(Value string) (*UpdateDeviceSecretParams){
    params.Value = &Value
    return params
}

// Update a secret for a Microvisor Device.
func (c *ApiService) UpdateDeviceSecret(DeviceSid string, Key string, params *UpdateDeviceSecretParams) (*MicrovisorV1DeviceSecret, error) {
    path := "/v1/Devices/{DeviceSid}/Secrets/{Key}"
        path = strings.Replace(path, "{"+"DeviceSid"+"}", DeviceSid, -1)
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

    ps := &MicrovisorV1DeviceSecret{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
