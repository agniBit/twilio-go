/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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


// Returns a single Track resource represented by TrackName or SID.
func (c *ApiService) FetchRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, Sid string, ) (*VideoV1RoomParticipantPublishedTrack, error) {
    path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid}"
        path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
    path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VideoV1RoomParticipantPublishedTrack{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListRoomParticipantPublishedTrack'
type ListRoomParticipantPublishedTrackParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantPublishedTrackParams) SetPageSize(PageSize int) (*ListRoomParticipantPublishedTrackParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListRoomParticipantPublishedTrackParams) SetLimit(Limit int) (*ListRoomParticipantPublishedTrackParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of RoomParticipantPublishedTrack records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams, pageToken, pageNumber string) (*ListRoomParticipantPublishedTrackResponse, error) {
    path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks"

        path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
    path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

    ps := &ListRoomParticipantPublishedTrackResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists RoomParticipantPublishedTrack records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams) ([]VideoV1RoomParticipantPublishedTrack, error) {
	response, errors := c.StreamRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params)

	records := make([]VideoV1RoomParticipantPublishedTrack, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RoomParticipantPublishedTrack records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams) (chan VideoV1RoomParticipantPublishedTrack, chan error) {
	if params == nil {
		params = &ListRoomParticipantPublishedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1RoomParticipantPublishedTrack, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRoomParticipantPublishedTrack(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamRoomParticipantPublishedTrack(response *ListRoomParticipantPublishedTrackResponse, params *ListRoomParticipantPublishedTrackParams, recordChannel chan VideoV1RoomParticipantPublishedTrack, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.PublishedTracks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRoomParticipantPublishedTrackResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoomParticipantPublishedTrackResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoomParticipantPublishedTrackResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListRoomParticipantPublishedTrackResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

