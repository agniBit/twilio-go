/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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


// Optional parameters for the method 'ListInsightsSegments'
type ListInsightsSegmentsParams struct {
    // The Authorization HTTP request header
    Authorization *string `json:"Authorization,omitempty"`
    // To unique id of the segment
    SegmentId *string `json:"SegmentId,omitempty"`
    // The list of reservation Ids
    ReservationId *[]string `json:"ReservationId,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListInsightsSegmentsParams) SetAuthorization(Authorization string) (*ListInsightsSegmentsParams){
    params.Authorization = &Authorization
    return params
}
func (params *ListInsightsSegmentsParams) SetSegmentId(SegmentId string) (*ListInsightsSegmentsParams){
    params.SegmentId = &SegmentId
    return params
}
func (params *ListInsightsSegmentsParams) SetReservationId(ReservationId []string) (*ListInsightsSegmentsParams){
    params.ReservationId = &ReservationId
    return params
}
func (params *ListInsightsSegmentsParams) SetPageSize(PageSize int) (*ListInsightsSegmentsParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListInsightsSegmentsParams) SetLimit(Limit int) (*ListInsightsSegmentsParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of InsightsSegments records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsSegments(params *ListInsightsSegmentsParams, pageToken, pageNumber string) (*ListInsightsSegmentsResponse, error) {
    path := "/v1/Insights/Segments"

    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.SegmentId != nil {
    data.Set("SegmentId", *params.SegmentId)
}
if params != nil && params.ReservationId != nil {
    for _, item  := range *params.ReservationId {
        data.Add("ReservationId", item)
    }
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

    ps := &ListInsightsSegmentsResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists InsightsSegments records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsSegments(params *ListInsightsSegmentsParams) ([]FlexV1InsightsSegments, error) {
	response, errors := c.StreamInsightsSegments(params)

	records := make([]FlexV1InsightsSegments, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsSegments records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsSegments(params *ListInsightsSegmentsParams) (chan FlexV1InsightsSegments, chan error) {
	if params == nil {
		params = &ListInsightsSegmentsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsSegments, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsSegments(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsSegments(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamInsightsSegments(response *ListInsightsSegmentsResponse, params *ListInsightsSegmentsParams, recordChannel chan FlexV1InsightsSegments, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Segments
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsSegmentsResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsSegmentsResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsSegmentsResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListInsightsSegmentsResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

