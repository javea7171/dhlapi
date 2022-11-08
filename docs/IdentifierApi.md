# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiIdentifiers**](IdentifierApi.md#ExpApiIdentifiers) | **Get** /identifiers | Service to allocate identifiers upfront for DHL Express Breakbulk or Loose Break Bulk shipments

# **ExpApiIdentifiers**
> SupermodelIoLogisticsExpressIdentifierResponse ExpApiIdentifiers(ctx, accountNumber, type_, size, optional)
Service to allocate identifiers upfront for DHL Express Breakbulk or Loose Break Bulk shipments

Service to allocate identifiers upfront for DHL Express Breakbulk or Loose Break Bulk shipments. Requires authorization to use this service from DHL Express. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**| DHL Express customer account number | 
  **type_** | **string**| Type of DHL Express identifier to retrieve | 
  **size** | **string**| Number of identifiers to be retrieved | 
 **optional** | ***IdentifierApiExpApiIdentifiersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentifierApiExpApiIdentifiersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressIdentifierResponse**](supermodelIoLogisticsExpressIdentifierResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

