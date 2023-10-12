/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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


// Optional parameters for the method 'ListAvailablePhoneNumberTollFree'
type ListAvailablePhoneNumberTollFreeParams struct {
    // The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
    AreaCode *int `json:"AreaCode,omitempty"`
    // The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
    Contains *string `json:"Contains,omitempty"`
    // Whether the phone numbers can receive text messages. Can be: `true` or `false`.
    SmsEnabled *bool `json:"SmsEnabled,omitempty"`
    // Whether the phone numbers can receive MMS messages. Can be: `true` or `false`.
    MmsEnabled *bool `json:"MmsEnabled,omitempty"`
    // Whether the phone numbers can receive calls. Can be: `true` or `false`.
    VoiceEnabled *bool `json:"VoiceEnabled,omitempty"`
    // Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
    ExcludeAllAddressRequired *bool `json:"ExcludeAllAddressRequired,omitempty"`
    // Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
    ExcludeLocalAddressRequired *bool `json:"ExcludeLocalAddressRequired,omitempty"`
    // Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
    ExcludeForeignAddressRequired *bool `json:"ExcludeForeignAddressRequired,omitempty"`
    // Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
    Beta *bool `json:"Beta,omitempty"`
    // Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
    NearNumber *string `json:"NearNumber,omitempty"`
    // Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada.
    NearLatLong *string `json:"NearLatLong,omitempty"`
    // The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada.
    Distance *int `json:"Distance,omitempty"`
    // Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
    InPostalCode *string `json:"InPostalCode,omitempty"`
    // Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
    InRegion *string `json:"InRegion,omitempty"`
    // Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada.
    InRateCenter *string `json:"InRateCenter,omitempty"`
    // Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
    InLata *string `json:"InLata,omitempty"`
    // Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
    InLocality *string `json:"InLocality,omitempty"`
    // Whether the phone numbers can receive faxes. Can be: `true` or `false`.
    FaxEnabled *bool `json:"FaxEnabled,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListAvailablePhoneNumberTollFreeParams) SetPathAccountSid(PathAccountSid string) (*ListAvailablePhoneNumberTollFreeParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetAreaCode(AreaCode int) (*ListAvailablePhoneNumberTollFreeParams){
    params.AreaCode = &AreaCode
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetContains(Contains string) (*ListAvailablePhoneNumberTollFreeParams){
    params.Contains = &Contains
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetSmsEnabled(SmsEnabled bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.SmsEnabled = &SmsEnabled
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetMmsEnabled(MmsEnabled bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.MmsEnabled = &MmsEnabled
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetVoiceEnabled(VoiceEnabled bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.VoiceEnabled = &VoiceEnabled
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetExcludeAllAddressRequired(ExcludeAllAddressRequired bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.ExcludeAllAddressRequired = &ExcludeAllAddressRequired
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetExcludeLocalAddressRequired(ExcludeLocalAddressRequired bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.ExcludeLocalAddressRequired = &ExcludeLocalAddressRequired
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetExcludeForeignAddressRequired(ExcludeForeignAddressRequired bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.ExcludeForeignAddressRequired = &ExcludeForeignAddressRequired
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetBeta(Beta bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.Beta = &Beta
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetNearNumber(NearNumber string) (*ListAvailablePhoneNumberTollFreeParams){
    params.NearNumber = &NearNumber
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetNearLatLong(NearLatLong string) (*ListAvailablePhoneNumberTollFreeParams){
    params.NearLatLong = &NearLatLong
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetDistance(Distance int) (*ListAvailablePhoneNumberTollFreeParams){
    params.Distance = &Distance
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetInPostalCode(InPostalCode string) (*ListAvailablePhoneNumberTollFreeParams){
    params.InPostalCode = &InPostalCode
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetInRegion(InRegion string) (*ListAvailablePhoneNumberTollFreeParams){
    params.InRegion = &InRegion
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetInRateCenter(InRateCenter string) (*ListAvailablePhoneNumberTollFreeParams){
    params.InRateCenter = &InRateCenter
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetInLata(InLata string) (*ListAvailablePhoneNumberTollFreeParams){
    params.InLata = &InLata
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetInLocality(InLocality string) (*ListAvailablePhoneNumberTollFreeParams){
    params.InLocality = &InLocality
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetFaxEnabled(FaxEnabled bool) (*ListAvailablePhoneNumberTollFreeParams){
    params.FaxEnabled = &FaxEnabled
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetPageSize(PageSize int) (*ListAvailablePhoneNumberTollFreeParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListAvailablePhoneNumberTollFreeParams) SetLimit(Limit int) (*ListAvailablePhoneNumberTollFreeParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of AvailablePhoneNumberTollFree records from the API. Request is executed immediately.
func (c *ApiService) PageAvailablePhoneNumberTollFree(CountryCode string, params *ListAvailablePhoneNumberTollFreeParams, pageToken, pageNumber string) (*, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/TollFree.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.AreaCode != nil {
    data.Set("AreaCode", fmt.Sprint(*params.AreaCode))
}
if params != nil && params.Contains != nil {
    data.Set("Contains", *params.Contains)
}
if params != nil && params.SmsEnabled != nil {
    data.Set("SmsEnabled", fmt.Sprint(*params.SmsEnabled))
}
if params != nil && params.MmsEnabled != nil {
    data.Set("MmsEnabled", fmt.Sprint(*params.MmsEnabled))
}
if params != nil && params.VoiceEnabled != nil {
    data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
}
if params != nil && params.ExcludeAllAddressRequired != nil {
    data.Set("ExcludeAllAddressRequired", fmt.Sprint(*params.ExcludeAllAddressRequired))
}
if params != nil && params.ExcludeLocalAddressRequired != nil {
    data.Set("ExcludeLocalAddressRequired", fmt.Sprint(*params.ExcludeLocalAddressRequired))
}
if params != nil && params.ExcludeForeignAddressRequired != nil {
    data.Set("ExcludeForeignAddressRequired", fmt.Sprint(*params.ExcludeForeignAddressRequired))
}
if params != nil && params.Beta != nil {
    data.Set("Beta", fmt.Sprint(*params.Beta))
}
if params != nil && params.NearNumber != nil {
    data.Set("NearNumber", *params.NearNumber)
}
if params != nil && params.NearLatLong != nil {
    data.Set("NearLatLong", *params.NearLatLong)
}
if params != nil && params.Distance != nil {
    data.Set("Distance", fmt.Sprint(*params.Distance))
}
if params != nil && params.InPostalCode != nil {
    data.Set("InPostalCode", *params.InPostalCode)
}
if params != nil && params.InRegion != nil {
    data.Set("InRegion", *params.InRegion)
}
if params != nil && params.InRateCenter != nil {
    data.Set("InRateCenter", *params.InRateCenter)
}
if params != nil && params.InLata != nil {
    data.Set("InLata", *params.InLata)
}
if params != nil && params.InLocality != nil {
    data.Set("InLocality", *params.InLocality)
}
if params != nil && params.FaxEnabled != nil {
    data.Set("FaxEnabled", fmt.Sprint(*params.FaxEnabled))
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

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists AvailablePhoneNumberTollFree records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAvailablePhoneNumberTollFree(CountryCode string, params *ListAvailablePhoneNumberTollFreeParams) (ListAvailablePhoneNumberTollFree200Response, error) {
	response, errors := c.StreamAvailablePhoneNumberTollFree(CountryCode, params)

	records := make(ListAvailablePhoneNumberTollFree200Response, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AvailablePhoneNumberTollFree records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAvailablePhoneNumberTollFree(CountryCode string, params *ListAvailablePhoneNumberTollFreeParams) (chan ListAvailablePhoneNumberTollFree200Response, chan error) {
	if params == nil {
		params = &ListAvailablePhoneNumberTollFreeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ListAvailablePhoneNumberTollFree200Response, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAvailablePhoneNumberTollFree(CountryCode, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAvailablePhoneNumberTollFree(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamAvailablePhoneNumberTollFree(response *, params *ListAvailablePhoneNumberTollFreeParams, recordChannel chan ListAvailablePhoneNumberTollFree200Response, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNext)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNext(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

