/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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


// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
    // 
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) (*CreateServiceParams){
    params.FriendlyName = &FriendlyName
    return params
}

// 
func (c *ApiService) CreateService(params *CreateServiceParams) (*IpMessagingV2Service, error) {
    path := "/v2/Services"
    
    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.FriendlyName != nil {
    data.Set("FriendlyName", *params.FriendlyName)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &IpMessagingV2Service{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// 
func (c *ApiService) DeleteService(Sid string, ) (error) {
    path := "/v2/Services/{Sid}"
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

// 
func (c *ApiService) FetchService(Sid string, ) (*IpMessagingV2Service, error) {
    path := "/v2/Services/{Sid}"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &IpMessagingV2Service{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) (*ListServiceParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListServiceParams) SetLimit(Limit int) (*ListServiceParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
    path := "/v2/Services"

    
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

    ps := &ListServiceResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]IpMessagingV2Service, error) {
	response, errors := c.StreamService(params)

	records := make([]IpMessagingV2Service, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan IpMessagingV2Service, chan error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IpMessagingV2Service, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageService(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamService(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamService(response *ListServiceResponse, params *ListServiceParams, recordChannel chan IpMessagingV2Service, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Services
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListServiceResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
    // 
    FriendlyName *string `json:"FriendlyName,omitempty"`
    // 
    DefaultServiceRoleSid *string `json:"DefaultServiceRoleSid,omitempty"`
    // 
    DefaultChannelRoleSid *string `json:"DefaultChannelRoleSid,omitempty"`
    // 
    DefaultChannelCreatorRoleSid *string `json:"DefaultChannelCreatorRoleSid,omitempty"`
    // 
    ReadStatusEnabled *bool `json:"ReadStatusEnabled,omitempty"`
    // 
    ReachabilityEnabled *bool `json:"ReachabilityEnabled,omitempty"`
    // 
    TypingIndicatorTimeout *int `json:"TypingIndicatorTimeout,omitempty"`
    // 
    ConsumptionReportInterval *int `json:"ConsumptionReportInterval,omitempty"`
    // 
    NotificationsNewMessageEnabled *bool `json:"Notifications.NewMessage.Enabled,omitempty"`
    // 
    NotificationsNewMessageTemplate *string `json:"Notifications.NewMessage.Template,omitempty"`
    // 
    NotificationsNewMessageSound *string `json:"Notifications.NewMessage.Sound,omitempty"`
    // 
    NotificationsNewMessageBadgeCountEnabled *bool `json:"Notifications.NewMessage.BadgeCountEnabled,omitempty"`
    // 
    NotificationsAddedToChannelEnabled *bool `json:"Notifications.AddedToChannel.Enabled,omitempty"`
    // 
    NotificationsAddedToChannelTemplate *string `json:"Notifications.AddedToChannel.Template,omitempty"`
    // 
    NotificationsAddedToChannelSound *string `json:"Notifications.AddedToChannel.Sound,omitempty"`
    // 
    NotificationsRemovedFromChannelEnabled *bool `json:"Notifications.RemovedFromChannel.Enabled,omitempty"`
    // 
    NotificationsRemovedFromChannelTemplate *string `json:"Notifications.RemovedFromChannel.Template,omitempty"`
    // 
    NotificationsRemovedFromChannelSound *string `json:"Notifications.RemovedFromChannel.Sound,omitempty"`
    // 
    NotificationsInvitedToChannelEnabled *bool `json:"Notifications.InvitedToChannel.Enabled,omitempty"`
    // 
    NotificationsInvitedToChannelTemplate *string `json:"Notifications.InvitedToChannel.Template,omitempty"`
    // 
    NotificationsInvitedToChannelSound *string `json:"Notifications.InvitedToChannel.Sound,omitempty"`
    // 
    PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
    // 
    PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
    // 
    WebhookMethod *string `json:"WebhookMethod,omitempty"`
    // 
    WebhookFilters *[]string `json:"WebhookFilters,omitempty"`
    // 
    LimitsChannelMembers *int `json:"Limits.ChannelMembers,omitempty"`
    // 
    LimitsUserChannels *int `json:"Limits.UserChannels,omitempty"`
    // 
    MediaCompatibilityMessage *string `json:"Media.CompatibilityMessage,omitempty"`
    // 
    PreWebhookRetryCount *int `json:"PreWebhookRetryCount,omitempty"`
    // 
    PostWebhookRetryCount *int `json:"PostWebhookRetryCount,omitempty"`
    // 
    NotificationsLogEnabled *bool `json:"Notifications.LogEnabled,omitempty"`
}

func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) (*UpdateServiceParams){
    params.FriendlyName = &FriendlyName
    return params
}
func (params *UpdateServiceParams) SetDefaultServiceRoleSid(DefaultServiceRoleSid string) (*UpdateServiceParams){
    params.DefaultServiceRoleSid = &DefaultServiceRoleSid
    return params
}
func (params *UpdateServiceParams) SetDefaultChannelRoleSid(DefaultChannelRoleSid string) (*UpdateServiceParams){
    params.DefaultChannelRoleSid = &DefaultChannelRoleSid
    return params
}
func (params *UpdateServiceParams) SetDefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid string) (*UpdateServiceParams){
    params.DefaultChannelCreatorRoleSid = &DefaultChannelCreatorRoleSid
    return params
}
func (params *UpdateServiceParams) SetReadStatusEnabled(ReadStatusEnabled bool) (*UpdateServiceParams){
    params.ReadStatusEnabled = &ReadStatusEnabled
    return params
}
func (params *UpdateServiceParams) SetReachabilityEnabled(ReachabilityEnabled bool) (*UpdateServiceParams){
    params.ReachabilityEnabled = &ReachabilityEnabled
    return params
}
func (params *UpdateServiceParams) SetTypingIndicatorTimeout(TypingIndicatorTimeout int) (*UpdateServiceParams){
    params.TypingIndicatorTimeout = &TypingIndicatorTimeout
    return params
}
func (params *UpdateServiceParams) SetConsumptionReportInterval(ConsumptionReportInterval int) (*UpdateServiceParams){
    params.ConsumptionReportInterval = &ConsumptionReportInterval
    return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageEnabled(NotificationsNewMessageEnabled bool) (*UpdateServiceParams){
    params.NotificationsNewMessageEnabled = &NotificationsNewMessageEnabled
    return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageTemplate(NotificationsNewMessageTemplate string) (*UpdateServiceParams){
    params.NotificationsNewMessageTemplate = &NotificationsNewMessageTemplate
    return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageSound(NotificationsNewMessageSound string) (*UpdateServiceParams){
    params.NotificationsNewMessageSound = &NotificationsNewMessageSound
    return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageBadgeCountEnabled(NotificationsNewMessageBadgeCountEnabled bool) (*UpdateServiceParams){
    params.NotificationsNewMessageBadgeCountEnabled = &NotificationsNewMessageBadgeCountEnabled
    return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled bool) (*UpdateServiceParams){
    params.NotificationsAddedToChannelEnabled = &NotificationsAddedToChannelEnabled
    return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate string) (*UpdateServiceParams){
    params.NotificationsAddedToChannelTemplate = &NotificationsAddedToChannelTemplate
    return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelSound(NotificationsAddedToChannelSound string) (*UpdateServiceParams){
    params.NotificationsAddedToChannelSound = &NotificationsAddedToChannelSound
    return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled bool) (*UpdateServiceParams){
    params.NotificationsRemovedFromChannelEnabled = &NotificationsRemovedFromChannelEnabled
    return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate string) (*UpdateServiceParams){
    params.NotificationsRemovedFromChannelTemplate = &NotificationsRemovedFromChannelTemplate
    return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelSound(NotificationsRemovedFromChannelSound string) (*UpdateServiceParams){
    params.NotificationsRemovedFromChannelSound = &NotificationsRemovedFromChannelSound
    return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled bool) (*UpdateServiceParams){
    params.NotificationsInvitedToChannelEnabled = &NotificationsInvitedToChannelEnabled
    return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate string) (*UpdateServiceParams){
    params.NotificationsInvitedToChannelTemplate = &NotificationsInvitedToChannelTemplate
    return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelSound(NotificationsInvitedToChannelSound string) (*UpdateServiceParams){
    params.NotificationsInvitedToChannelSound = &NotificationsInvitedToChannelSound
    return params
}
func (params *UpdateServiceParams) SetPreWebhookUrl(PreWebhookUrl string) (*UpdateServiceParams){
    params.PreWebhookUrl = &PreWebhookUrl
    return params
}
func (params *UpdateServiceParams) SetPostWebhookUrl(PostWebhookUrl string) (*UpdateServiceParams){
    params.PostWebhookUrl = &PostWebhookUrl
    return params
}
func (params *UpdateServiceParams) SetWebhookMethod(WebhookMethod string) (*UpdateServiceParams){
    params.WebhookMethod = &WebhookMethod
    return params
}
func (params *UpdateServiceParams) SetWebhookFilters(WebhookFilters []string) (*UpdateServiceParams){
    params.WebhookFilters = &WebhookFilters
    return params
}
func (params *UpdateServiceParams) SetLimitsChannelMembers(LimitsChannelMembers int) (*UpdateServiceParams){
    params.LimitsChannelMembers = &LimitsChannelMembers
    return params
}
func (params *UpdateServiceParams) SetLimitsUserChannels(LimitsUserChannels int) (*UpdateServiceParams){
    params.LimitsUserChannels = &LimitsUserChannels
    return params
}
func (params *UpdateServiceParams) SetMediaCompatibilityMessage(MediaCompatibilityMessage string) (*UpdateServiceParams){
    params.MediaCompatibilityMessage = &MediaCompatibilityMessage
    return params
}
func (params *UpdateServiceParams) SetPreWebhookRetryCount(PreWebhookRetryCount int) (*UpdateServiceParams){
    params.PreWebhookRetryCount = &PreWebhookRetryCount
    return params
}
func (params *UpdateServiceParams) SetPostWebhookRetryCount(PostWebhookRetryCount int) (*UpdateServiceParams){
    params.PostWebhookRetryCount = &PostWebhookRetryCount
    return params
}
func (params *UpdateServiceParams) SetNotificationsLogEnabled(NotificationsLogEnabled bool) (*UpdateServiceParams){
    params.NotificationsLogEnabled = &NotificationsLogEnabled
    return params
}

// 
func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*IpMessagingV2Service, error) {
    path := "/v2/Services/{Sid}"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.FriendlyName != nil {
    data.Set("FriendlyName", *params.FriendlyName)
}
if params != nil && params.DefaultServiceRoleSid != nil {
    data.Set("DefaultServiceRoleSid", *params.DefaultServiceRoleSid)
}
if params != nil && params.DefaultChannelRoleSid != nil {
    data.Set("DefaultChannelRoleSid", *params.DefaultChannelRoleSid)
}
if params != nil && params.DefaultChannelCreatorRoleSid != nil {
    data.Set("DefaultChannelCreatorRoleSid", *params.DefaultChannelCreatorRoleSid)
}
if params != nil && params.ReadStatusEnabled != nil {
    data.Set("ReadStatusEnabled", fmt.Sprint(*params.ReadStatusEnabled))
}
if params != nil && params.ReachabilityEnabled != nil {
    data.Set("ReachabilityEnabled", fmt.Sprint(*params.ReachabilityEnabled))
}
if params != nil && params.TypingIndicatorTimeout != nil {
    data.Set("TypingIndicatorTimeout", fmt.Sprint(*params.TypingIndicatorTimeout))
}
if params != nil && params.ConsumptionReportInterval != nil {
    data.Set("ConsumptionReportInterval", fmt.Sprint(*params.ConsumptionReportInterval))
}
if params != nil && params.NotificationsNewMessageEnabled != nil {
    data.Set("Notifications.NewMessage.Enabled", fmt.Sprint(*params.NotificationsNewMessageEnabled))
}
if params != nil && params.NotificationsNewMessageTemplate != nil {
    data.Set("Notifications.NewMessage.Template", *params.NotificationsNewMessageTemplate)
}
if params != nil && params.NotificationsNewMessageSound != nil {
    data.Set("Notifications.NewMessage.Sound", *params.NotificationsNewMessageSound)
}
if params != nil && params.NotificationsNewMessageBadgeCountEnabled != nil {
    data.Set("Notifications.NewMessage.BadgeCountEnabled", fmt.Sprint(*params.NotificationsNewMessageBadgeCountEnabled))
}
if params != nil && params.NotificationsAddedToChannelEnabled != nil {
    data.Set("Notifications.AddedToChannel.Enabled", fmt.Sprint(*params.NotificationsAddedToChannelEnabled))
}
if params != nil && params.NotificationsAddedToChannelTemplate != nil {
    data.Set("Notifications.AddedToChannel.Template", *params.NotificationsAddedToChannelTemplate)
}
if params != nil && params.NotificationsAddedToChannelSound != nil {
    data.Set("Notifications.AddedToChannel.Sound", *params.NotificationsAddedToChannelSound)
}
if params != nil && params.NotificationsRemovedFromChannelEnabled != nil {
    data.Set("Notifications.RemovedFromChannel.Enabled", fmt.Sprint(*params.NotificationsRemovedFromChannelEnabled))
}
if params != nil && params.NotificationsRemovedFromChannelTemplate != nil {
    data.Set("Notifications.RemovedFromChannel.Template", *params.NotificationsRemovedFromChannelTemplate)
}
if params != nil && params.NotificationsRemovedFromChannelSound != nil {
    data.Set("Notifications.RemovedFromChannel.Sound", *params.NotificationsRemovedFromChannelSound)
}
if params != nil && params.NotificationsInvitedToChannelEnabled != nil {
    data.Set("Notifications.InvitedToChannel.Enabled", fmt.Sprint(*params.NotificationsInvitedToChannelEnabled))
}
if params != nil && params.NotificationsInvitedToChannelTemplate != nil {
    data.Set("Notifications.InvitedToChannel.Template", *params.NotificationsInvitedToChannelTemplate)
}
if params != nil && params.NotificationsInvitedToChannelSound != nil {
    data.Set("Notifications.InvitedToChannel.Sound", *params.NotificationsInvitedToChannelSound)
}
if params != nil && params.PreWebhookUrl != nil {
    data.Set("PreWebhookUrl", *params.PreWebhookUrl)
}
if params != nil && params.PostWebhookUrl != nil {
    data.Set("PostWebhookUrl", *params.PostWebhookUrl)
}
if params != nil && params.WebhookMethod != nil {
    data.Set("WebhookMethod", *params.WebhookMethod)
}
if params != nil && params.WebhookFilters != nil {
    for _, item  := range *params.WebhookFilters {
        data.Add("WebhookFilters", item)
    }
}
if params != nil && params.LimitsChannelMembers != nil {
    data.Set("Limits.ChannelMembers", fmt.Sprint(*params.LimitsChannelMembers))
}
if params != nil && params.LimitsUserChannels != nil {
    data.Set("Limits.UserChannels", fmt.Sprint(*params.LimitsUserChannels))
}
if params != nil && params.MediaCompatibilityMessage != nil {
    data.Set("Media.CompatibilityMessage", *params.MediaCompatibilityMessage)
}
if params != nil && params.PreWebhookRetryCount != nil {
    data.Set("PreWebhookRetryCount", fmt.Sprint(*params.PreWebhookRetryCount))
}
if params != nil && params.PostWebhookRetryCount != nil {
    data.Set("PostWebhookRetryCount", fmt.Sprint(*params.PostWebhookRetryCount))
}
if params != nil && params.NotificationsLogEnabled != nil {
    data.Set("Notifications.LogEnabled", fmt.Sprint(*params.NotificationsLogEnabled))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &IpMessagingV2Service{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
