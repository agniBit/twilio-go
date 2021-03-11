/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUsageRecordYearlyResponse struct for ListUsageRecordYearlyResponse
type ListUsageRecordYearlyResponse struct {
	End             int32                                              `json:"End,omitempty"`
	FirstPageUri    string                                             `json:"FirstPageUri,omitempty"`
	NextPageUri     string                                             `json:"NextPageUri,omitempty"`
	Page            int32                                              `json:"Page,omitempty"`
	PageSize        int32                                              `json:"PageSize,omitempty"`
	PreviousPageUri string                                             `json:"PreviousPageUri,omitempty"`
	Start           int32                                              `json:"Start,omitempty"`
	Uri             string                                             `json:"Uri,omitempty"`
	UsageRecords    []ApiV2010AccountUsageUsageRecordUsageRecordYearly `json:"UsageRecords,omitempty"`
}