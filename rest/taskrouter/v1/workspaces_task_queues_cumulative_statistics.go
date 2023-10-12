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


// Optional parameters for the method 'FetchTaskQueueCumulativeStatistics'
type FetchTaskQueueCumulativeStatisticsParams struct {
    // Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
    EndDate *time.Time `json:"EndDate,omitempty"`
    // Only calculate statistics since this many minutes in the past. The default is 15 minutes.
    Minutes *int `json:"Minutes,omitempty"`
    // Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
    StartDate *time.Time `json:"StartDate,omitempty"`
    // Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
    TaskChannel *string `json:"TaskChannel,omitempty"`
    // A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. TaskRouter will calculate statistics on up to 10,000 Tasks/Reservations for any given threshold.
    SplitByWaitTime *string `json:"SplitByWaitTime,omitempty"`
}

func (params *FetchTaskQueueCumulativeStatisticsParams) SetEndDate(EndDate time.Time) (*FetchTaskQueueCumulativeStatisticsParams){
    params.EndDate = &EndDate
    return params
}
func (params *FetchTaskQueueCumulativeStatisticsParams) SetMinutes(Minutes int) (*FetchTaskQueueCumulativeStatisticsParams){
    params.Minutes = &Minutes
    return params
}
func (params *FetchTaskQueueCumulativeStatisticsParams) SetStartDate(StartDate time.Time) (*FetchTaskQueueCumulativeStatisticsParams){
    params.StartDate = &StartDate
    return params
}
func (params *FetchTaskQueueCumulativeStatisticsParams) SetTaskChannel(TaskChannel string) (*FetchTaskQueueCumulativeStatisticsParams){
    params.TaskChannel = &TaskChannel
    return params
}
func (params *FetchTaskQueueCumulativeStatisticsParams) SetSplitByWaitTime(SplitByWaitTime string) (*FetchTaskQueueCumulativeStatisticsParams){
    params.SplitByWaitTime = &SplitByWaitTime
    return params
}

// 
func (c *ApiService) FetchTaskQueueCumulativeStatistics(WorkspaceSid string, TaskQueueSid string, params *FetchTaskQueueCumulativeStatisticsParams) (*TaskrouterV1TaskQueueCumulativeStatistics, error) {
    path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/CumulativeStatistics"
        path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
    path = strings.Replace(path, "{"+"TaskQueueSid"+"}", TaskQueueSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.EndDate != nil {
    data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
}
if params != nil && params.Minutes != nil {
    data.Set("Minutes", fmt.Sprint(*params.Minutes))
}
if params != nil && params.StartDate != nil {
    data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
}
if params != nil && params.TaskChannel != nil {
    data.Set("TaskChannel", *params.TaskChannel)
}
if params != nil && params.SplitByWaitTime != nil {
    data.Set("SplitByWaitTime", *params.SplitByWaitTime)
}



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TaskrouterV1TaskQueueCumulativeStatistics{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
