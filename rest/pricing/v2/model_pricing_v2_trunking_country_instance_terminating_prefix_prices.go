/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
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
// PricingV2TrunkingCountryInstanceTerminatingPrefixPrices struct for PricingV2TrunkingCountryInstanceTerminatingPrefixPrices
type PricingV2TrunkingCountryInstanceTerminatingPrefixPrices struct {
	OriginationPrefixes []string `json:"origination_prefixes,omitempty"`
	DestinationPrefixes []string `json:"destination_prefixes,omitempty"`
	BasePrice float32 `json:"base_price,omitempty"`
	CurrentPrice float32 `json:"current_price,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
}

func (response *PricingV2TrunkingCountryInstanceTerminatingPrefixPrices) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		OriginationPrefixes []string `json:"origination_prefixes"`
		DestinationPrefixes []string `json:"destination_prefixes"`
		BasePrice interface{} `json:"base_price"`
		CurrentPrice interface{} `json:"current_price"`
		FriendlyName string `json:"friendly_name"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = PricingV2TrunkingCountryInstanceTerminatingPrefixPrices{
		OriginationPrefixes: raw.OriginationPrefixes,
		DestinationPrefixes: raw.DestinationPrefixes,
		FriendlyName: raw.FriendlyName,
	}

	responseBasePrice, err := client.UnmarshalFloat32(&raw.BasePrice)
	if err != nil {
		return err
	}
	response.BasePrice = *responseBasePrice

	responseCurrentPrice, err := client.UnmarshalFloat32(&raw.CurrentPrice)
	if err != nil {
		return err
	}
	response.CurrentPrice = *responseCurrentPrice

	return
}

