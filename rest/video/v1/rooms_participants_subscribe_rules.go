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


// Returns a list of Subscribe Rules for the Participant.
func (c *ApiService) FetchRoomParticipantSubscribeRule(RoomSid string, ParticipantSid string, ) (*VideoV1RoomParticipantSubscribeRule, error) {
    path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules"
        path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
    path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VideoV1RoomParticipantSubscribeRule{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateRoomParticipantSubscribeRule'
type UpdateRoomParticipantSubscribeRuleParams struct {
    // A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information.
    Rules *interface{} `json:"Rules,omitempty"`
}

func (params *UpdateRoomParticipantSubscribeRuleParams) SetRules(Rules interface{}) (*UpdateRoomParticipantSubscribeRuleParams){
    params.Rules = &Rules
    return params
}

// Update the Subscribe Rules for the Participant
func (c *ApiService) UpdateRoomParticipantSubscribeRule(RoomSid string, ParticipantSid string, params *UpdateRoomParticipantSubscribeRuleParams) (*VideoV1RoomParticipantSubscribeRule, error) {
    path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules"
        path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
    path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Rules != nil {
    v, err := json.Marshal(params.Rules)

    if err != nil {
        return nil, err
    }

    data.Set("Rules", string(v))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &VideoV1RoomParticipantSubscribeRule{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
