# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiPickups**](PickupApi.md#ExpApiPickups) | **Post** /pickups | Create a DHL Express pickup booking request
[**ExpApiPickupsCancel**](PickupApi.md#ExpApiPickupsCancel) | **Delete** /pickups/{dispatchConfirmationNumber} | Cancel a DHL Express pickup booking request
[**ExpApiPickupsUpdate**](PickupApi.md#ExpApiPickupsUpdate) | **Patch** /pickups/{dispatchConfirmationNumber} | Update pickup information for an existing DHL Express pickup booking request

# **ExpApiPickups**
> SupermodelIoLogisticsExpressPickupResponse ExpApiPickups(ctx, body, optional)
Create a DHL Express pickup booking request

The Pickup service creates a DHL Express pickup booking request 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressPickupRequest**](SupermodelIoLogisticsExpressPickupRequest.md)| Details about the requested pickup | 
 **optional** | ***PickupApiExpApiPickupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickupApiExpApiPickupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageReference** | **optional.**| Please provide message reference  | 
 **messageReferenceDate** | **optional.**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressPickupResponse**](supermodelIoLogisticsExpressPickupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiPickupsCancel**
> ExpApiPickupsCancel(ctx, dispatchConfirmationNumber, requestorName, reason, optional)
Cancel a DHL Express pickup booking request

The Cancel Pickup service can be used to cancel a DHL Express pickup booking request 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dispatchConfirmationNumber** | **string**| Shipment pickup confirmation number for example &#x60;PRG999126012345&#x60; | 
  **requestorName** | **string**| Name of the person requesting to cancel the scheduled pickup  | 
  **reason** | **string**| Provide why scheduled pickup is being cancelled  | 
 **optional** | ***PickupApiExpApiPickupsCancelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickupApiExpApiPickupsCancelOpts struct
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

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiPickupsUpdate**
> SupermodelIoLogisticsExpressUpdatePickupResponse ExpApiPickupsUpdate(ctx, body, dispatchConfirmationNumber, optional)
Update pickup information for an existing DHL Express pickup booking request

The Update Pickup service can be used to update pickup information for an existing DHL Express pickup booking request 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressUpdatePickupRequest**](SupermodelIoLogisticsExpressUpdatePickupRequest.md)| Details about the requested pickup updates | 
  **dispatchConfirmationNumber** | **string**| Shipment pickup confirmation number for example &#x60;PRG999126012345&#x60; | 
 **optional** | ***PickupApiExpApiPickupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickupApiExpApiPickupsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messageReference** | **optional.**| Please provide message reference  | 
 **messageReferenceDate** | **optional.**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressUpdatePickupResponse**](supermodelIoLogisticsExpressUpdatePickupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

