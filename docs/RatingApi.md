# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiLandedCost**](RatingApi.md#ExpApiLandedCost) | **Post** /landed-cost | Landed Cost
[**ExpApiRates**](RatingApi.md#ExpApiRates) | **Get** /rates | Retrieve Rates for a one piece Shipment
[**ExpApiRatesMany**](RatingApi.md#ExpApiRatesMany) | **Post** /rates | Retrieve Rates for Multi-piece Shipments

# **ExpApiLandedCost**
> SupermodelIoLogisticsExpressRates ExpApiLandedCost(ctx, body, optional)
Landed Cost

The Landed Cost section allows further information around products being sold to be provided. In return the duty, tax and shipping charges are calculated in real time and provides transparency about any extra costs the buyer may have to pay before they reach them. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressLandedCostRequest**](SupermodelIoLogisticsExpressLandedCostRequest.md)| Details about the requested shipment | 
 **optional** | ***RatingApiExpApiLandedCostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingApiExpApiLandedCostOpts struct
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

[**SupermodelIoLogisticsExpressRates**](supermodelIoLogisticsExpressRates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiRates**
> SupermodelIoLogisticsExpressRates ExpApiRates(ctx, accountNumber, originCountryCode, originCityName, destinationCountryCode, destinationCityName, weight, length, width, height, plannedShippingDate, isCustomsDeclarable, unitOfMeasurement, optional)
Retrieve Rates for a one piece Shipment

The Rate request will return DHL's product capabilities and prices (where applicable) based on the input data. Using the shipper and receiver address as well as the dimension and weights of the pieces belonging to a shipment, this operation returns the available products including the shipping price (where applicable) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**| DHL Express customer account number | 
  **originCountryCode** | **string**| A short text string code (see values defined in ISO 3166) specifying the shipment origin country. https://gs1.org/voc/Country, Alpha-2 Code | 
  **originCityName** | **string**| Text specifying the city name | 
  **destinationCountryCode** | **string**| A short text string code (see values defined in ISO 3166) specifying the shipment destination country. https://gs1.org/voc/Country, Alpha-2 Code | 
  **destinationCityName** | **string**| Text specifying the city name | 
  **weight** | **float64**| Gross weight of the shipment including packaging. | 
  **length** | **float64**| Total length of the shipment including packaging. | 
  **width** | **float64**| Total width of the shipment including packaging. | 
  **height** | **float64**| Total height of the shipment including packaging. | 
  **plannedShippingDate** | **string**| Timestamp represents the date you plan to ship your prospected shipment  | 
  **isCustomsDeclarable** | **bool**|  | 
  **unitOfMeasurement** | **string**| The UnitOfMeasurement node conveys the unit of measurements used in the operation. This single value corresponds to the units of weight and measurement which are used throughout the message processing.  | 
 **optional** | ***RatingApiExpApiRatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingApiExpApiRatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------












 **originPostalCode** | **optional.String**| Text specifying the postal code for an address. https://gs1.org/voc/postalCode | 
 **destinationPostalCode** | **optional.String**| Text specifying the postal code for an address. https://gs1.org/voc/postalCode | 
 **nextBusinessDay** | **optional.Bool**| When set to true and there are no products available for given plannedShippingDate then products available for the next possible pickup date are returned  | 
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 
 **strictValidation** | **optional.Bool**| If set to true, indicate strict DCT validation of address details, and validation of product and service(s) combination provided in request. | [default to false]
 **getAllValueAddedServices** | **optional.Bool**| Option to return list of all value added services and its rule groups if applicable | [default to false]
 **requestEstimatedDeliveryDate** | **optional.Bool**| Option to return Estimated Delivery Date in response | [default to true]
 **estimatedDeliveryDateType** | **optional.String**| Estimated Delivery Date Type. QDDF: is the fastest &#x27;docs&#x27; transit time as quoted to the customer at booking or shipment creation. No custom clearance is considered. QDDC: constitutes DHL&#x27;s service commitment as quoted at booking or shipment creation. QDDc builds in clearance time, and potentially other special perational non-transport component(s), when relevant.  | [default to QDDF]

### Return type

[**SupermodelIoLogisticsExpressRates**](supermodelIoLogisticsExpressRates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiRatesMany**
> SupermodelIoLogisticsExpressRates ExpApiRatesMany(ctx, body, optional)
Retrieve Rates for Multi-piece Shipments

The Rate request will return DHL's product capabilities and prices (where applicable) based on the input data. Using the shipper and receiver address as well as the dimension and weights of the pieces belonging to a shipment, this operation returns the available products including the shipping price (where applicable) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressRateRequest**](SupermodelIoLogisticsExpressRateRequest.md)| Details about the requested shipment | 
 **optional** | ***RatingApiExpApiRatesManyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingApiExpApiRatesManyOpts struct
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
 **strictValidation** | **optional.**| If set to true, indicate strict DCT validation of address details, and validation of product and service(s) combination provided in request. | [default to false]

### Return type

[**SupermodelIoLogisticsExpressRates**](supermodelIoLogisticsExpressRates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

