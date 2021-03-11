/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VideoRoomSummaryEndReason the model 'VideoRoomSummaryEndReason'
type VideoRoomSummaryEndReason string

// List of video_room_summary_end_reason
const (
	VIDEOROOMSUMMARYENDREASON_ROOM_ENDED_VIA_API VideoRoomSummaryEndReason = "room_ended_via_api"
	VIDEOROOMSUMMARYENDREASON_TIMEOUT            VideoRoomSummaryEndReason = "timeout"
)