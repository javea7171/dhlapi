# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiShipmentsTracking**](TrackingApi.md#ExpApiShipmentsTracking) | **Get** /shipments/{shipmentTrackingNumber}/tracking | Track a single DHL Express Shipment
[**ExpApiShipmentsTrackingMulti**](TrackingApi.md#ExpApiShipmentsTrackingMulti) | **Get** /tracking | Track a single or multiple DHL Express Shipments

# **ExpApiShipmentsTracking**
> SupermodelIoLogisticsExpressTrackingResponse ExpApiShipmentsTracking(ctx, shipmentTrackingNumber, optional)
Track a single DHL Express Shipment

The Tracking service retrieves tracking statuses for a single DHL Express Shipment 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentTrackingNumber** | **string**| DHL Express shipment identification number | 
 **optional** | ***TrackingApiExpApiShipmentsTrackingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrackingApiExpApiShipmentsTrackingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackingView** | **optional.String**|  | [default to all-checkpoints]
 **levelOfDetail** | **optional.String**|  | [default to shipment]
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **acceptLanguage** | **optional.String**|  | [default to eng]
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressTrackingResponse**](supermodelIoLogisticsExpressTrackingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiShipmentsTrackingMulti**
> SupermodelIoLogisticsExpressTrackingResponse ExpApiShipmentsTrackingMulti(ctx, optional)
Track a single or multiple DHL Express Shipments

The Tracking service retrieves tracking statuses for a single or multiple DHL Express Shipments 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TrackingApiExpApiShipmentsTrackingMultiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrackingApiExpApiShipmentsTrackingMultiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentTrackingNumber** | [**optional.Interface of []string**](string.md)| DHL Express shipment identification number | 
 **pieceTrackingNumber** | [**optional.Interface of []string**](string.md)| DHL Express shipment piece tracking number | 
 **shipmentReference** | **optional.String**| Shipment reference which was provided during the shipment label creation  | 
 **shipmentReferenceType** | **optional.String**| Shipment reference type which was provided during the shipment label creation  | 
 **shipperAccountNumber** | **optional.String**| Shipper DHL Express Account number under which the shipment label was created  | 
 **dateRangeFrom** | **optional.String**| When tracking by Shipment reference you need to restrict the search by timeframe. Please provide the start of the period.  | 
 **dateRangeTo** | **optional.String**| When tracking by Shipment reference you need to restrict the search by timeframe. Please provide the end of the period.  | 
 **trackingView** | **optional.String**|  | [default to all-checkpoints]
 **levelOfDetail** | **optional.String**|  | [default to shipment]
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **acceptLanguage** | **optional.String**|  | [default to eng]
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressTrackingResponse**](supermodelIoLogisticsExpressTrackingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

