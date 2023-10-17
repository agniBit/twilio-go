/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
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


// Delete a specific Sync Document Permission.
func (c *ApiService) DeleteDocumentPermission(ServiceSid string, DocumentSid string, Identity string, ) (error) {
    path := "/v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"DocumentSid"+"}", DocumentSid, -1)
    path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Fetch a specific Sync Document Permission.
func (c *ApiService) FetchDocumentPermission(ServiceSid string, DocumentSid string, Identity string, ) (*SyncV1DocumentPermission, error) {
    path := "/v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"DocumentSid"+"}", DocumentSid, -1)
    path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

    data := url.Values{}
    headers := make(map[string]interface{})



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &SyncV1DocumentPermission{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListDocumentPermission'
type ListDocumentPermissionParams struct {
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListDocumentPermissionParams) SetPageSize(PageSize int) (*ListDocumentPermissionParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListDocumentPermissionParams) SetLimit(Limit int) (*ListDocumentPermissionParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of DocumentPermission records from the API. Request is executed immediately.
func (c *ApiService) PageDocumentPermission(ServiceSid string, DocumentSid string, params *ListDocumentPermissionParams, pageToken, pageNumber string) (*ListDocumentPermissionResponse, error) {
    path := "/v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions"

        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"DocumentSid"+"}", DocumentSid, -1)

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

    ps := &ListDocumentPermissionResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists DocumentPermission records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDocumentPermission(ServiceSid string, DocumentSid string, params *ListDocumentPermissionParams) ([]SyncV1DocumentPermission, error) {
	response, errors := c.StreamDocumentPermission(ServiceSid, DocumentSid, params)

	records := make([]SyncV1DocumentPermission, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DocumentPermission records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDocumentPermission(ServiceSid string, DocumentSid string, params *ListDocumentPermissionParams) (chan SyncV1DocumentPermission, chan error) {
	if params == nil {
		params = &ListDocumentPermissionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SyncV1DocumentPermission, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDocumentPermission(ServiceSid, DocumentSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDocumentPermission(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamDocumentPermission(response *ListDocumentPermissionResponse, params *ListDocumentPermissionParams, recordChannel chan SyncV1DocumentPermission, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Permissions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDocumentPermissionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDocumentPermissionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDocumentPermissionResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListDocumentPermissionResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateDocumentPermission'
type UpdateDocumentPermissionParams struct {
    // Whether the identity can read the Sync Document. Default value is `false`.
    Read *bool `json:"Read,omitempty"`
    // Whether the identity can update the Sync Document. Default value is `false`.
    Write *bool `json:"Write,omitempty"`
    // Whether the identity can delete the Sync Document. Default value is `false`.
    Manage *bool `json:"Manage,omitempty"`
}

func (params *UpdateDocumentPermissionParams) SetRead(Read bool) (*UpdateDocumentPermissionParams){
    params.Read = &Read
    return params
}
func (params *UpdateDocumentPermissionParams) SetWrite(Write bool) (*UpdateDocumentPermissionParams){
    params.Write = &Write
    return params
}
func (params *UpdateDocumentPermissionParams) SetManage(Manage bool) (*UpdateDocumentPermissionParams){
    params.Manage = &Manage
    return params
}

// Update an identity's access to a specific Sync Document.
func (c *ApiService) UpdateDocumentPermission(ServiceSid string, DocumentSid string, Identity string, params *UpdateDocumentPermissionParams) (*SyncV1DocumentPermission, error) {
    path := "/v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity}"
        path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
    path = strings.Replace(path, "{"+"DocumentSid"+"}", DocumentSid, -1)
    path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

    data := url.Values{}
    headers := make(map[string]interface{})
if params != nil && params.Read != nil {
    data.Set("Read", fmt.Sprint(*params.Read))
}
if params != nil && params.Write != nil {
    data.Set("Write", fmt.Sprint(*params.Write))
}
if params != nil && params.Manage != nil {
    data.Set("Manage", fmt.Sprint(*params.Manage))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &SyncV1DocumentPermission{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
