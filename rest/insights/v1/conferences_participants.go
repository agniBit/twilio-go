/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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


// Optional parameters for the method 'FetchConferenceParticipant'
type FetchConferenceParticipantParams struct {
    // Conference events generated by application or participant activity; e.g. `hold`, `mute`, etc.
    Events *string `json:"Events,omitempty"`
    // Object. Contains participant call quality metrics.
    Metrics *string `json:"Metrics,omitempty"`
}

func (params *FetchConferenceParticipantParams) SetEvents(Events string) (*FetchConferenceParticipantParams){
    params.Events = &Events
    return params
}
func (params *FetchConferenceParticipantParams) SetMetrics(Metrics string) (*FetchConferenceParticipantParams){
    params.Metrics = &Metrics
    return params
}

// Get a specific Conference Participant Summary for a Conference.
func (c *ApiService) FetchConferenceParticipant(ConferenceSid string, ParticipantSid string, params *FetchConferenceParticipantParams) (*InsightsV1ConferenceParticipant, error) {
    path := "/v1/Conferences/{ConferenceSid}/Participants/{ParticipantSid}"
        path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
    path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.Events != nil {
        data.Set("Events", *params.Events)
    }
    if params != nil && params.Metrics != nil {
        data.Set("Metrics", *params.Metrics)
    }



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InsightsV1ConferenceParticipant{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListConferenceParticipant'
type ListConferenceParticipantParams struct {
    // The unique SID identifier of the Participant.
    ParticipantSid *string `json:"ParticipantSid,omitempty"`
    // User-specified label for a participant.
    Label *string `json:"Label,omitempty"`
    // Conference events generated by application or participant activity; e.g. `hold`, `mute`, etc.
    Events *string `json:"Events,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListConferenceParticipantParams) SetParticipantSid(ParticipantSid string) (*ListConferenceParticipantParams){
    params.ParticipantSid = &ParticipantSid
    return params
}
func (params *ListConferenceParticipantParams) SetLabel(Label string) (*ListConferenceParticipantParams){
    params.Label = &Label
    return params
}
func (params *ListConferenceParticipantParams) SetEvents(Events string) (*ListConferenceParticipantParams){
    params.Events = &Events
    return params
}
func (params *ListConferenceParticipantParams) SetPageSize(PageSize int) (*ListConferenceParticipantParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListConferenceParticipantParams) SetLimit(Limit int) (*ListConferenceParticipantParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of ConferenceParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams, pageToken, pageNumber string) (*ListConferenceParticipantResponse, error) {
    path := "/v1/Conferences/{ConferenceSid}/Participants"

        path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})

    if params != nil && params.ParticipantSid != nil {
        data.Set("ParticipantSid", *params.ParticipantSid)
    }
    if params != nil && params.Label != nil {
        data.Set("Label", *params.Label)
    }
    if params != nil && params.Events != nil {
        data.Set("Events", *params.Events)
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

    ps := &ListConferenceParticipantResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists ConferenceParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams) ([]InsightsV1ConferenceParticipant, error) {
	response, errors := c.StreamConferenceParticipant(ConferenceSid, params)

	records := make([]InsightsV1ConferenceParticipant, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConferenceParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConferenceParticipant(ConferenceSid string, params *ListConferenceParticipantParams) (chan InsightsV1ConferenceParticipant, chan error) {
	if params == nil {
		params = &ListConferenceParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan InsightsV1ConferenceParticipant, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConferenceParticipant(ConferenceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConferenceParticipant(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamConferenceParticipant(response *ListConferenceParticipantResponse, params *ListConferenceParticipantParams, recordChannel chan InsightsV1ConferenceParticipant, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Participants
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListConferenceParticipantResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConferenceParticipantResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConferenceParticipantResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListConferenceParticipantResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

