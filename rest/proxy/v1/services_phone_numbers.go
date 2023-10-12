/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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


// Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
    // The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.
    Sid *string `json:"Sid,omitempty"`
    // The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
    PhoneNumber *string `json:"PhoneNumber,omitempty"`
    // Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
    IsReserved *bool `json:"IsReserved,omitempty"`
}

func (params *CreatePhoneNumberParams) SetSid(Sid string) (*CreatePhoneNumberParams){
    params.Sid = &Sid
    return params
}
func (params *CreatePhoneNumberParams) SetPhoneNumber(PhoneNumber string) (*CreatePhoneNumberParams){
    params.PhoneNumber = &PhoneNumber
    return params
}
func (params *CreatePhoneNumberParams) SetIsReserved(IsReserved bool) (*CreatePhoneNumberParams){
    params.IsReserved = &IsReserved
    return params
}

// Add a Phone Number to a Service's Proxy Number Pool.
func (c *ApiService) CreatePhoneNumber(ServiceSid string, params *CreatePhoneNumberParams) (*ProxyV1PhoneNumber, error) {
    path := "/v1/Services/{ServiceSid}/PhoneNumbers"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Sid != nil {
    data.Set("Sid", *params.Sid)
}
if params != nil && params.PhoneNumber != nil {
    data.Set("PhoneNumber", *params.PhoneNumber)
}
if params != nil && params.IsReserved != nil {
    data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ProxyV1PhoneNumber{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Delete a specific Phone Number from a Service.
func (c *ApiService) DeletePhoneNumber(ServiceSid string, Sid string, ) (error) {
    path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Fetch a specific Phone Number.
func (c *ApiService) FetchPhoneNumber(ServiceSid string, Sid string, ) (*ProxyV1PhoneNumber, error) {
    path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ProxyV1PhoneNumber{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListPhoneNumberParams) SetPageSize(PageSize int) (*ListPhoneNumberParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListPhoneNumberParams) SetLimit(Limit int) (*ListPhoneNumberParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of PhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PagePhoneNumber(ServiceSid string, params *ListPhoneNumberParams, pageToken, pageNumber string) (*ListPhoneNumberResponse, error) {
    path := "/v1/Services/{ServiceSid}/PhoneNumbers"

        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

    ps := &ListPhoneNumberResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists PhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) ([]ProxyV1PhoneNumber, error) {
	response, errors := c.StreamPhoneNumber(ServiceSid, params)

	records := make([]ProxyV1PhoneNumber, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams PhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) (chan ProxyV1PhoneNumber, chan error) {
	if params == nil {
		params = &ListPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ProxyV1PhoneNumber, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePhoneNumber(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPhoneNumber(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamPhoneNumber(response *ListPhoneNumberResponse, params *ListPhoneNumberParams, recordChannel chan ProxyV1PhoneNumber, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.PhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPhoneNumberResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPhoneNumberResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPhoneNumberResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListPhoneNumberResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdatePhoneNumber'
type UpdatePhoneNumberParams struct {
    // Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
    IsReserved *bool `json:"IsReserved,omitempty"`
}

func (params *UpdatePhoneNumberParams) SetIsReserved(IsReserved bool) (*UpdatePhoneNumberParams){
    params.IsReserved = &IsReserved
    return params
}

// Update a specific Proxy Number.
func (c *ApiService) UpdatePhoneNumber(ServiceSid string, Sid string, params *UpdatePhoneNumberParams) (*ProxyV1PhoneNumber, error) {
    path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.IsReserved != nil {
    data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ProxyV1PhoneNumber{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
