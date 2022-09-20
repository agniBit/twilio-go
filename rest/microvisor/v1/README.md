# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.35.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://microvisor.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppsApi* | [**DeleteApp**](docs/AppsApi.md#deleteapp) | **Delete** /v1/Apps/{Sid} | 
*AppsApi* | [**FetchApp**](docs/AppsApi.md#fetchapp) | **Get** /v1/Apps/{Sid} | 
*AppsApi* | [**ListApp**](docs/AppsApi.md#listapp) | **Get** /v1/Apps | 
*DevicesApi* | [**FetchDevice**](docs/DevicesApi.md#fetchdevice) | **Get** /v1/Devices/{Sid} | 
*DevicesApi* | [**ListDevice**](docs/DevicesApi.md#listdevice) | **Get** /v1/Devices | 
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Post** /v1/Devices/{Sid} | 


## Documentation For Models

 - [MicrovisorV1Device](docs/MicrovisorV1Device.md)
 - [ListDeviceResponse](docs/ListDeviceResponse.md)
 - [MicrovisorV1App](docs/MicrovisorV1App.md)
 - [ListAppResponseMeta](docs/ListAppResponseMeta.md)
 - [ListAppResponse](docs/ListAppResponse.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
