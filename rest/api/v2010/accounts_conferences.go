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


// Optional parameters for the method 'FetchConference'
type FetchConferenceParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchConferenceParams) SetPathAccountSid(PathAccountSid string) (*FetchConferenceParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Fetch an instance of a conference
func (c *ApiService) FetchConference(Sid string, params *FetchConferenceParams) (*ApiV2010Conference, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Conference{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListConference'
type ListConferenceParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`.
    DateCreated *string `json:"DateCreated,omitempty"`
    // The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`.
    DateUpdated *string `json:"DateUpdated,omitempty"`
    // The string that identifies the Conference resources to read.
    FriendlyName *string `json:"FriendlyName,omitempty"`
    // The status of the resources to read. Can be: `init`, `in-progress`, or `completed`.
    Status *string `json:"Status,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListConferenceParams) SetPathAccountSid(PathAccountSid string) (*ListConferenceParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListConferenceParams) SetDateCreated(DateCreated string) (*ListConferenceParams){
    params.DateCreated = &DateCreated
    return params
}
func (params *ListConferenceParams) SetDateUpdated(DateUpdated string) (*ListConferenceParams){
    params.DateUpdated = &DateUpdated
    return params
}
func (params *ListConferenceParams) SetFriendlyName(FriendlyName string) (*ListConferenceParams){
    params.FriendlyName = &FriendlyName
    return params
}
func (params *ListConferenceParams) SetStatus(Status string) (*ListConferenceParams){
    params.Status = &Status
    return params
}
func (params *ListConferenceParams) SetPageSize(PageSize int) (*ListConferenceParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListConferenceParams) SetLimit(Limit int) (*ListConferenceParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Conference records from the API. Request is executed immediately.
func (c *ApiService) PageConference(params *ListConferenceParams, pageToken, pageNumber string) (*, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Conferences.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.DateCreated != nil {
        data.Set("DateCreated", fmt.Sprint(*params.DateCreated))
    }
    if params != nil && params.DateUpdated != nil {
        data.Set("DateUpdated", fmt.Sprint(*params.DateUpdated))
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
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

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Conference records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConference(params *ListConferenceParams) (ListConference200Response, error) {
	response, errors := c.StreamConference(params)

	records := make(ListConference200Response, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Conference records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConference(params *ListConferenceParams) (chan ListConference200Response, chan error) {
	if params == nil {
		params = &ListConferenceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ListConference200Response, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConference(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConference(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamConference(response *, params *ListConferenceParams, recordChannel chan ListConference200Response, errorChannel chan error) {
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


// Optional parameters for the method 'UpdateConference'
type UpdateConferenceParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // 
    Status *string `json:"Status,omitempty"`
    // The URL we should call to announce something into the conference. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.
    AnnounceUrl *string `json:"AnnounceUrl,omitempty"`
    // The HTTP method used to call `announce_url`. Can be: `GET` or `POST` and the default is `POST`
    AnnounceMethod *string `json:"AnnounceMethod,omitempty"`
}

func (params *UpdateConferenceParams) SetPathAccountSid(PathAccountSid string) (*UpdateConferenceParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *UpdateConferenceParams) SetStatus(Status string) (*UpdateConferenceParams){
    params.Status = &Status
    return params
}
func (params *UpdateConferenceParams) SetAnnounceUrl(AnnounceUrl string) (*UpdateConferenceParams){
    params.AnnounceUrl = &AnnounceUrl
    return params
}
func (params *UpdateConferenceParams) SetAnnounceMethod(AnnounceMethod string) (*UpdateConferenceParams){
    params.AnnounceMethod = &AnnounceMethod
    return params
}

// 
func (c *ApiService) UpdateConference(Sid string, params *UpdateConferenceParams) (*ApiV2010Conference, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }
    if params != nil && params.AnnounceUrl != nil {
        data.Set("AnnounceUrl", *params.AnnounceUrl)
    }
    if params != nil && params.AnnounceMethod != nil {
        data.Set("AnnounceMethod", *params.AnnounceMethod)
    }



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010Conference{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
