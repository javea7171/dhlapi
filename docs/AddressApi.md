# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiAddressValidate**](AddressApi.md#ExpApiAddressValidate) | **Get** /address-validate | Validate DHL Express pickup/delivery capabilities at origin/destination

# **ExpApiAddressValidate**
> SupermodelIoLogisticsExpressAddressValidateResponse ExpApiAddressValidate(ctx, type_, countryCode, optional)
Validate DHL Express pickup/delivery capabilities at origin/destination

Validates if DHL Express has got pickup/delivery capabilities at origin/destination 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **countryCode** | **string**| A short text string code (see values defined in ISO 3166) specifying the shipment origin country. https://gs1.org/voc/Country, Alpha-2 Code | 
 **optional** | ***AddressApiExpApiAddressValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiExpApiAddressValidateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postalCode** | **optional.String**| Text specifying the postal code for an address. https://gs1.org/voc/postalCode | 
 **cityName** | **optional.String**| Text specifying the city name | 
 **countyName** | **optional.String**| Text specifying the county name | 
 **strictValidation** | **optional.Bool**| If set to true service will return no records when exact valid match not found | [default to true]
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressAddressValidateResponse**](supermodelIoLogisticsExpressAddressValidateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

