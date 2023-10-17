/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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


// 
func (c *ApiService) FetchEvent(WorkspaceSid string, Sid string, ) (*TaskrouterV1Event, error) {
    path := "/v1/Workspaces/{WorkspaceSid}/Events/{Sid}"
        path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TaskrouterV1Event{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListEvent'
type ListEventParams struct {
    // Only include Events that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
    EndDate *time.Time `json:"EndDate,omitempty"`
    // The type of Events to read. Returns only Events of the type specified.
    EventType *string `json:"EventType,omitempty"`
    // The period of events to read in minutes. Returns only Events that occurred since this many minutes in the past. The default is `15` minutes. Task Attributes for Events occuring more 43,200 minutes ago will be redacted.
    Minutes *int `json:"Minutes,omitempty"`
    // The SID of the Reservation with the Events to read. Returns only Events that pertain to the specified Reservation.
    ReservationSid *string `json:"ReservationSid,omitempty"`
    // Only include Events from on or after this date and time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Task Attributes for Events older than 30 days will be redacted.
    StartDate *time.Time `json:"StartDate,omitempty"`
    // The SID of the TaskQueue with the Events to read. Returns only the Events that pertain to the specified TaskQueue.
    TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
    // The SID of the Task with the Events to read. Returns only the Events that pertain to the specified Task.
    TaskSid *string `json:"TaskSid,omitempty"`
    // The SID of the Worker with the Events to read. Returns only the Events that pertain to the specified Worker.
    WorkerSid *string `json:"WorkerSid,omitempty"`
    // The SID of the Workflow with the Events to read. Returns only the Events that pertain to the specified Workflow.
    WorkflowSid *string `json:"WorkflowSid,omitempty"`
    // The TaskChannel with the Events to read. Returns only the Events that pertain to the specified TaskChannel.
    TaskChannel *string `json:"TaskChannel,omitempty"`
    // The SID of the Event resource to read.
    Sid *string `json:"Sid,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListEventParams) SetEndDate(EndDate time.Time) (*ListEventParams){
    params.EndDate = &EndDate
    return params
}
func (params *ListEventParams) SetEventType(EventType string) (*ListEventParams){
    params.EventType = &EventType
    return params
}
func (params *ListEventParams) SetMinutes(Minutes int) (*ListEventParams){
    params.Minutes = &Minutes
    return params
}
func (params *ListEventParams) SetReservationSid(ReservationSid string) (*ListEventParams){
    params.ReservationSid = &ReservationSid
    return params
}
func (params *ListEventParams) SetStartDate(StartDate time.Time) (*ListEventParams){
    params.StartDate = &StartDate
    return params
}
func (params *ListEventParams) SetTaskQueueSid(TaskQueueSid string) (*ListEventParams){
    params.TaskQueueSid = &TaskQueueSid
    return params
}
func (params *ListEventParams) SetTaskSid(TaskSid string) (*ListEventParams){
    params.TaskSid = &TaskSid
    return params
}
func (params *ListEventParams) SetWorkerSid(WorkerSid string) (*ListEventParams){
    params.WorkerSid = &WorkerSid
    return params
}
func (params *ListEventParams) SetWorkflowSid(WorkflowSid string) (*ListEventParams){
    params.WorkflowSid = &WorkflowSid
    return params
}
func (params *ListEventParams) SetTaskChannel(TaskChannel string) (*ListEventParams){
    params.TaskChannel = &TaskChannel
    return params
}
func (params *ListEventParams) SetSid(Sid string) (*ListEventParams){
    params.Sid = &Sid
    return params
}
func (params *ListEventParams) SetPageSize(PageSize int) (*ListEventParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListEventParams) SetLimit(Limit int) (*ListEventParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Event records from the API. Request is executed immediately.
func (c *ApiService) PageEvent(WorkspaceSid string, params *ListEventParams, pageToken, pageNumber string) (*ListEventResponse, error) {
    path := "/v1/Workspaces/{WorkspaceSid}/Events"

        path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.EndDate != nil {
    data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
}
if params != nil && params.EventType != nil {
    data.Set("EventType", *params.EventType)
}
if params != nil && params.Minutes != nil {
    data.Set("Minutes", fmt.Sprint(*params.Minutes))
}
if params != nil && params.ReservationSid != nil {
    data.Set("ReservationSid", *params.ReservationSid)
}
if params != nil && params.StartDate != nil {
    data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
}
if params != nil && params.TaskQueueSid != nil {
    data.Set("TaskQueueSid", *params.TaskQueueSid)
}
if params != nil && params.TaskSid != nil {
    data.Set("TaskSid", *params.TaskSid)
}
if params != nil && params.WorkerSid != nil {
    data.Set("WorkerSid", *params.WorkerSid)
}
if params != nil && params.WorkflowSid != nil {
    data.Set("WorkflowSid", *params.WorkflowSid)
}
if params != nil && params.TaskChannel != nil {
    data.Set("TaskChannel", *params.TaskChannel)
}
if params != nil && params.Sid != nil {
    data.Set("Sid", *params.Sid)
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

    ps := &ListEventResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Event records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEvent(WorkspaceSid string, params *ListEventParams) ([]TaskrouterV1Event, error) {
	response, errors := c.StreamEvent(WorkspaceSid, params)

	records := make([]TaskrouterV1Event, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Event records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEvent(WorkspaceSid string, params *ListEventParams) (chan TaskrouterV1Event, chan error) {
	if params == nil {
		params = &ListEventParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1Event, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEvent(WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEvent(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamEvent(response *ListEventResponse, params *ListEventParams, recordChannel chan TaskrouterV1Event, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Events
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEventResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEventResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEventResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListEventResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

