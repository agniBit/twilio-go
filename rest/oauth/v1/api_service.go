/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Oauth
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
    twilio "github.com/twilio/twilio-go/client"
)

type ApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewApiService(requestHandler *twilio.RequestHandler) *ApiService {
	return &ApiService {
		requestHandler: requestHandler,
		baseURL:        "https://oauth.twilio.com",
	}
}

func NewApiServiceWithClient(client twilio.BaseClient) *ApiService {
    return NewApiService(twilio.NewRequestHandler(client))
}
