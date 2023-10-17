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


// Optional parameters for the method 'FetchMember'
type FetchMemberParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMemberParams) SetPathAccountSid(PathAccountSid string) (*FetchMemberParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Fetch a specific member from the queue
func (c *ApiService) FetchMember(QueueSid string, CallSid string, params *FetchMemberParams) (*ApiV2010Member, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)
    path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Member{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListMember'
type ListMemberParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListMemberParams) SetPathAccountSid(PathAccountSid string) (*ListMemberParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListMemberParams) SetPageSize(PageSize int) (*ListMemberParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListMemberParams) SetLimit(Limit int) (*ListMemberParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Member records from the API. Request is executed immediately.
func (c *ApiService) PageMember(QueueSid string, params *ListMemberParams, pageToken, pageNumber string) (*, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)

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

// Lists Member records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMember(QueueSid string, params *ListMemberParams) (ListMember200Response, error) {
	response, errors := c.StreamMember(QueueSid, params)

	records := make(ListMember200Response, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Member records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMember(QueueSid string, params *ListMemberParams) (chan ListMember200Response, chan error) {
	if params == nil {
		params = &ListMemberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ListMember200Response, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMember(QueueSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMember(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamMember(response *, params *ListMemberParams, recordChannel chan ListMember200Response, errorChannel chan error) {
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


// Optional parameters for the method 'UpdateMember'
type UpdateMemberParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // The absolute URL of the Queue resource.
    Url *string `json:"Url,omitempty"`
    // How to pass the update request data. Can be `GET` or `POST` and the default is `POST`. `POST` sends the data as encoded form data and `GET` sends the data as query parameters.
    Method *string `json:"Method,omitempty"`
}

func (params *UpdateMemberParams) SetPathAccountSid(PathAccountSid string) (*UpdateMemberParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *UpdateMemberParams) SetUrl(Url string) (*UpdateMemberParams){
    params.Url = &Url
    return params
}
func (params *UpdateMemberParams) SetMethod(Method string) (*UpdateMemberParams){
    params.Method = &Method
    return params
}

// Dequeue a member from a queue and have the member's call begin executing the TwiML document at that URL
func (c *ApiService) UpdateMember(QueueSid string, CallSid string, params *UpdateMemberParams) (*ApiV2010Member, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)
    path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Url != nil {
    data.Set("Url", *params.Url)
}
if params != nil && params.Method != nil {
    data.Set("Method", *params.Method)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Member{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
