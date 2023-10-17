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


// Optional parameters for the method 'FetchCallNotification'
type FetchCallNotificationParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallNotificationParams) SetPathAccountSid(PathAccountSid string) (*FetchCallNotificationParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// 
func (c *ApiService) FetchCallNotification(CallSid string, Sid string, params *FetchCallNotificationParams) (*ApiV2010CallNotificationInstance, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010CallNotificationInstance{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListCallNotification'
type ListCallNotificationParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read.
    Log *int `json:"Log,omitempty"`
    // Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
    MessageDate *string `json:"MessageDate,omitempty"`
    // Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
    MessageDateBefore *string `json:"MessageDate&lt;,omitempty"`
    // Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date.
    MessageDateAfter *string `json:"MessageDate&gt;,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListCallNotificationParams) SetPathAccountSid(PathAccountSid string) (*ListCallNotificationParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListCallNotificationParams) SetLog(Log int) (*ListCallNotificationParams){
    params.Log = &Log
    return params
}
func (params *ListCallNotificationParams) SetMessageDate(MessageDate string) (*ListCallNotificationParams){
    params.MessageDate = &MessageDate
    return params
}
func (params *ListCallNotificationParams) SetMessageDateBefore(MessageDateBefore string) (*ListCallNotificationParams){
    params.MessageDateBefore = &MessageDateBefore
    return params
}
func (params *ListCallNotificationParams) SetMessageDateAfter(MessageDateAfter string) (*ListCallNotificationParams){
    params.MessageDateAfter = &MessageDateAfter
    return params
}
func (params *ListCallNotificationParams) SetPageSize(PageSize int) (*ListCallNotificationParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListCallNotificationParams) SetLimit(Limit int) (*ListCallNotificationParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of CallNotification records from the API. Request is executed immediately.
func (c *ApiService) PageCallNotification(CallSid string, params *ListCallNotificationParams, pageToken, pageNumber string) (*, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Log != nil {
    data.Set("Log", fmt.Sprint(*params.Log))
}
if params != nil && params.MessageDate != nil {
    data.Set("MessageDate", fmt.Sprint(*params.MessageDate))
}
if params != nil && params.MessageDateBefore != nil {
    data.Set("MessageDate<", fmt.Sprint(*params.MessageDateBefore))
}
if params != nil && params.MessageDateAfter != nil {
    data.Set("MessageDate>", fmt.Sprint(*params.MessageDateAfter))
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

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists CallNotification records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallNotification(CallSid string, params *ListCallNotificationParams) (ListCallNotification200Response, error) {
	response, errors := c.StreamCallNotification(CallSid, params)

	records := make(ListCallNotification200Response, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CallNotification records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallNotification(CallSid string, params *ListCallNotificationParams) (chan ListCallNotification200Response, chan error) {
	if params == nil {
		params = &ListCallNotificationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ListCallNotification200Response, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCallNotification(CallSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCallNotification(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamCallNotification(response *, params *ListCallNotificationParams, recordChannel chan ListCallNotification200Response, errorChannel chan error) {
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

