/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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


// Optional parameters for the method 'DeleteRecordingAddOnResult'
type DeleteRecordingAddOnResultParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingAddOnResultParams) SetPathAccountSid(PathAccountSid string) (*DeleteRecordingAddOnResultParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Delete a result and purge all associated Payloads
func (c *ApiService) DeleteRecordingAddOnResult(ReferenceSid string, Sid string, params *DeleteRecordingAddOnResultParams) (error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
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

// Optional parameters for the method 'FetchRecordingAddOnResult'
type FetchRecordingAddOnResultParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingAddOnResultParams) SetPathAccountSid(PathAccountSid string) (*FetchRecordingAddOnResultParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Fetch an instance of an AddOnResult
func (c *ApiService) FetchRecordingAddOnResult(ReferenceSid string, Sid string, params *FetchRecordingAddOnResultParams) (*ApiV2010RecordingAddOnResult, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010RecordingAddOnResult{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListRecordingAddOnResult'
type ListRecordingAddOnResultParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListRecordingAddOnResultParams) SetPathAccountSid(PathAccountSid string) (*ListRecordingAddOnResultParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListRecordingAddOnResultParams) SetPageSize(PageSize int) (*ListRecordingAddOnResultParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListRecordingAddOnResultParams) SetLimit(Limit int) (*ListRecordingAddOnResultParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of RecordingAddOnResult records from the API. Request is executed immediately.
func (c *ApiService) PageRecordingAddOnResult(ReferenceSid string, params *ListRecordingAddOnResultParams, pageToken, pageNumber string) (*, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)

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

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists RecordingAddOnResult records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRecordingAddOnResult(ReferenceSid string, params *ListRecordingAddOnResultParams) (ListRecordingAddOnResult200Response, error) {
	response, errors := c.StreamRecordingAddOnResult(ReferenceSid, params)

	records := make(ListRecordingAddOnResult200Response, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RecordingAddOnResult records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRecordingAddOnResult(ReferenceSid string, params *ListRecordingAddOnResultParams) (chan ListRecordingAddOnResult200Response, chan error) {
	if params == nil {
		params = &ListRecordingAddOnResultParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ListRecordingAddOnResult200Response, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRecordingAddOnResult(ReferenceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRecordingAddOnResult(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamRecordingAddOnResult(response *, params *ListRecordingAddOnResultParams, recordChannel chan ListRecordingAddOnResult200Response, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNext)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNext(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

