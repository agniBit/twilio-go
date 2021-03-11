# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlphaSender**](DefaultApi.md#CreateAlphaSender) | **Post** /v1/Services/{ServiceSid}/AlphaSenders | 
[**CreateBrandRegistrations**](DefaultApi.md#CreateBrandRegistrations) | **Post** /v1/a2p/BrandRegistrations | 
[**CreatePhoneNumber**](DefaultApi.md#CreatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateShortCode**](DefaultApi.md#CreateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes | 
[**DeleteAlphaSender**](DefaultApi.md#DeleteAlphaSender) | **Delete** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**DeletePhoneNumber**](DefaultApi.md#DeletePhoneNumber) | **Delete** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteShortCode**](DefaultApi.md#DeleteShortCode) | **Delete** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**FetchAlphaSender**](DefaultApi.md#FetchAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**FetchBrandRegistrations**](DefaultApi.md#FetchBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations/{Sid} | 
[**FetchDeactivation**](DefaultApi.md#FetchDeactivation) | **Get** /v1/Deactivations | 
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchShortCode**](DefaultApi.md#FetchShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**ListAlphaSender**](DefaultApi.md#ListAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders | 
[**ListBrandRegistrations**](DefaultApi.md#ListBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations | 
[**ListPhoneNumber**](DefaultApi.md#ListPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListShortCode**](DefaultApi.md#ListShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



## CreateAlphaSender

> MessagingV1ServiceAlphaSender CreateAlphaSender(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateAlphaSenderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AlphaSender** | **string** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, and hyphen &#x60;-&#x60;. This value cannot contain only numbers. | 

### Return type

[**MessagingV1ServiceAlphaSender**](MessagingV1ServiceAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBrandRegistrations

> MessagingV1BrandRegistrations CreateBrandRegistrations(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandRegistrationsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**A2pProfileBundleSid** | **string** | A2P Messaging Profile Bundle Sid. | 
**CustomerProfileBundleSid** | **string** | Customer Profile Bundle Sid. | 

### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> MessagingV1ServicePhoneNumber CreatePhoneNumber(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PhoneNumberSid** | **string** | The SID of the Phone Number being added to the Service. | 

### Return type

[**MessagingV1ServicePhoneNumber**](MessagingV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> MessagingV1Service CreateService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AreaCodeGeomatch** | **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. | 
**FallbackMethod** | **string** | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**FallbackToLongCode** | **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. | 
**FallbackUrl** | **string** | The URL that we should call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**InboundMethod** | **string** | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
**InboundRequestUrl** | **string** | The URL we should call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled. | 
**MmsConverter** | **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. | 
**ScanMessageContent** | **string** | Reserved. | 
**SmartEncoding** | **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. | 
**StatusCallback** | **string** | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. | 
**StickySender** | **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. | 
**SynchronousValidation** | **bool** | Reserved. | 
**ValidityPeriod** | **int32** | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;. | 

### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShortCode

> MessagingV1ServiceShortCode CreateShortCode(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ShortCodeSid** | **string** | The SID of the ShortCode resource being added to the Service. | 

### Return type

[**MessagingV1ServiceShortCode**](MessagingV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlphaSender

> DeleteAlphaSender(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from. | 
**Sid** | **string** | The SID of the AlphaSender resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteAlphaSenderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneNumber

> DeletePhoneNumber(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from. | 
**Sid** | **string** | The SID of the PhoneNumber resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShortCode

> DeleteShortCode(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from. | 
**Sid** | **string** | The SID of the ShortCode resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteShortCodeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAlphaSender

> MessagingV1ServiceAlphaSender FetchAlphaSender(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from. | 
**Sid** | **string** | The SID of the AlphaSender resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAlphaSenderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MessagingV1ServiceAlphaSender**](MessagingV1ServiceAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandRegistrations

> MessagingV1BrandRegistrations FetchBrandRegistrations(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Brand Registration resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandRegistrationsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeactivation

> FetchDeactivation(ctx, optional)



Fetch a list of all United States numbers that have been deactivated on a specific date.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeactivationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Date** | **time.Time** | The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> MessagingV1ServicePhoneNumber FetchPhoneNumber(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from. | 
**Sid** | **string** | The SID of the PhoneNumber resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MessagingV1ServicePhoneNumber**](MessagingV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> MessagingV1Service FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> MessagingV1ServiceShortCode FetchShortCode(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from. | 
**Sid** | **string** | The SID of the ShortCode resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MessagingV1ServiceShortCode**](MessagingV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlphaSender

> ListAlphaSenderResponse ListAlphaSender(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListAlphaSenderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAlphaSenderResponse**](ListAlphaSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandRegistrations

> ListBrandRegistrationsResponse ListBrandRegistrations(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBrandRegistrationsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListBrandRegistrationsResponse**](ListBrandRegistrationsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListPhoneNumberResponse**](ListPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> ListShortCodeResponse ListShortCode(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListShortCodeResponse**](ListShortCodeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> MessagingV1Service UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AreaCodeGeomatch** | **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. | 
**FallbackMethod** | **string** | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**FallbackToLongCode** | **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. | 
**FallbackUrl** | **string** | The URL that we should call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**InboundMethod** | **string** | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
**InboundRequestUrl** | **string** | The URL we should call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled. | 
**MmsConverter** | **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. | 
**ScanMessageContent** | **string** | Reserved. | 
**SmartEncoding** | **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. | 
**StatusCallback** | **string** | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. | 
**StickySender** | **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. | 
**SynchronousValidation** | **bool** | Reserved. | 
**ValidityPeriod** | **int32** | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;. | 

### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
