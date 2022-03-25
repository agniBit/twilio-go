/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListInteractionChannelParticipantResponse struct for ListInteractionChannelParticipantResponse
type ListInteractionChannelParticipantResponse struct {
	Meta         ListChannelResponseMeta               `json:"meta,omitempty"`
	Participants []FlexV1InteractionChannelParticipant `json:"participants,omitempty"`
}