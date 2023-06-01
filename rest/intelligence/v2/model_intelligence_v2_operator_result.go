/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// IntelligenceV2OperatorResult struct for IntelligenceV2OperatorResult
type IntelligenceV2OperatorResult struct {
	OperatorType *string `json:"operator_type,omitempty"`
	// The name of the applied Language Understanding.
	Name *string `json:"name,omitempty"`
	// A 34 character string that identifies this Language Understanding operator sid.
	OperatorSid *string `json:"operator_sid,omitempty"`
	// Boolean to tell if extract Language Understanding Processing model matches results.
	ExtractMatch *bool `json:"extract_match,omitempty"`
	// Percentage of 'matching' class needed to consider a sentence matches
	MatchProbability *float32 `json:"match_probability,omitempty"`
	// Normalized output of extraction stage which matches Label.
	NormalizedResult *string `json:"normalized_result,omitempty"`
	// List of mapped utterance object which matches sentences.
	UtteranceResults *[]interface{} `json:"utterance_results,omitempty"`
	// Boolean to tell if Utterance matches results.
	UtteranceMatch *bool `json:"utterance_match,omitempty"`
	// The 'matching' class. This might be available on conversation classify model outputs.
	PredictedLabel *string `json:"predicted_label,omitempty"`
	// Percentage of 'matching' class needed to consider a sentence matches.
	PredictedProbability *float32 `json:"predicted_probability,omitempty"`
	// The labels probabilities. This might be available on conversation classify model outputs.
	LabelProbabilities *interface{} `json:"label_probabilities,omitempty"`
	// List of text extraction results. This might be available on classify-extract model outputs.
	ExtractResults *interface{} `json:"extract_results,omitempty"`
	// A 34 character string that uniquely identifies this Transcript.
	TranscriptSid *string `json:"transcript_sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}

func (response *IntelligenceV2OperatorResult) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		OperatorType         *string        `json:"operator_type"`
		Name                 *string        `json:"name"`
		OperatorSid          *string        `json:"operator_sid"`
		ExtractMatch         *bool          `json:"extract_match"`
		MatchProbability     *interface{}   `json:"match_probability"`
		NormalizedResult     *string        `json:"normalized_result"`
		UtteranceResults     *[]interface{} `json:"utterance_results"`
		UtteranceMatch       *bool          `json:"utterance_match"`
		PredictedLabel       *string        `json:"predicted_label"`
		PredictedProbability *interface{}   `json:"predicted_probability"`
		LabelProbabilities   *interface{}   `json:"label_probabilities"`
		ExtractResults       *interface{}   `json:"extract_results"`
		TranscriptSid        *string        `json:"transcript_sid"`
		Url                  *string        `json:"url"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = IntelligenceV2OperatorResult{
		OperatorType:       raw.OperatorType,
		Name:               raw.Name,
		OperatorSid:        raw.OperatorSid,
		ExtractMatch:       raw.ExtractMatch,
		NormalizedResult:   raw.NormalizedResult,
		UtteranceResults:   raw.UtteranceResults,
		UtteranceMatch:     raw.UtteranceMatch,
		PredictedLabel:     raw.PredictedLabel,
		LabelProbabilities: raw.LabelProbabilities,
		ExtractResults:     raw.ExtractResults,
		TranscriptSid:      raw.TranscriptSid,
		Url:                raw.Url,
	}

	responseMatchProbability, err := client.UnmarshalFloat32(raw.MatchProbability)
	if err != nil {
		return err
	}
	response.MatchProbability = responseMatchProbability

	responsePredictedProbability, err := client.UnmarshalFloat32(raw.PredictedProbability)
	if err != nil {
		return err
	}
	response.PredictedProbability = responsePredictedProbability

	return
}
