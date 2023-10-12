/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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


// Optional parameters for the method 'ListDependentHostedNumberOrder'
type ListDependentHostedNumberOrderParams struct {
    // Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource#status-values) for more information on each of these statuses.
    Status *string `json:"Status,omitempty"`
    // An E164 formatted phone number hosted by this HostedNumberOrder.
    PhoneNumber *string `json:"PhoneNumber,omitempty"`
    // A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder.
    IncomingPhoneNumberSid *string `json:"IncomingPhoneNumberSid,omitempty"`
    // A human readable description of this resource, up to 128 characters.
    FriendlyName *string `json:"FriendlyName,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListDependentHostedNumberOrderParams) SetStatus(Status string) (*ListDependentHostedNumberOrderParams){
    params.Status = &Status
    return params
}
func (params *ListDependentHostedNumberOrderParams) SetPhoneNumber(PhoneNumber string) (*ListDependentHostedNumberOrderParams){
    params.PhoneNumber = &PhoneNumber
    return params
}
func (params *ListDependentHostedNumberOrderParams) SetIncomingPhoneNumberSid(IncomingPhoneNumberSid string) (*ListDependentHostedNumberOrderParams){
    params.IncomingPhoneNumberSid = &IncomingPhoneNumberSid
    return params
}
func (params *ListDependentHostedNumberOrderParams) SetFriendlyName(FriendlyName string) (*ListDependentHostedNumberOrderParams){
    params.FriendlyName = &FriendlyName
    return params
}
func (params *ListDependentHostedNumberOrderParams) SetPageSize(PageSize int) (*ListDependentHostedNumberOrderParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListDependentHostedNumberOrderParams) SetLimit(Limit int) (*ListDependentHostedNumberOrderParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of DependentHostedNumberOrder records from the API. Request is executed immediately.
func (c *ApiService) PageDependentHostedNumberOrder(SigningDocumentSid string, params *ListDependentHostedNumberOrderParams, pageToken, pageNumber string) (*ListDependentHostedNumberOrderResponse, error) {
    path := "/v2/HostedNumber/AuthorizationDocuments/{SigningDocumentSid}/DependentHostedNumberOrders"

        path = strings.Replace(path, "{"+"SigningDocumentSid"+"}", SigningDocumentSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Status != nil {
    data.Set("Status", *params.Status)
}
if params != nil && params.PhoneNumber != nil {
    data.Set("PhoneNumber", *params.PhoneNumber)
}
if params != nil && params.IncomingPhoneNumberSid != nil {
    data.Set("IncomingPhoneNumberSid", *params.IncomingPhoneNumberSid)
}
if params != nil && params.FriendlyName != nil {
    data.Set("FriendlyName", *params.FriendlyName)
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

    ps := &ListDependentHostedNumberOrderResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists DependentHostedNumberOrder records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDependentHostedNumberOrder(SigningDocumentSid string, params *ListDependentHostedNumberOrderParams) ([]NumbersV2DependentHostedNumberOrder, error) {
	response, errors := c.StreamDependentHostedNumberOrder(SigningDocumentSid, params)

	records := make([]NumbersV2DependentHostedNumberOrder, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DependentHostedNumberOrder records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDependentHostedNumberOrder(SigningDocumentSid string, params *ListDependentHostedNumberOrderParams) (chan NumbersV2DependentHostedNumberOrder, chan error) {
	if params == nil {
		params = &ListDependentHostedNumberOrderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2DependentHostedNumberOrder, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDependentHostedNumberOrder(SigningDocumentSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDependentHostedNumberOrder(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamDependentHostedNumberOrder(response *ListDependentHostedNumberOrderResponse, params *ListDependentHostedNumberOrderParams, recordChannel chan NumbersV2DependentHostedNumberOrder, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Items
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDependentHostedNumberOrderResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDependentHostedNumberOrderResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDependentHostedNumberOrderResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListDependentHostedNumberOrderResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

