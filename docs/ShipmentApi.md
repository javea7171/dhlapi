# {{classname}}

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpApiShipments**](ShipmentApi.md#ExpApiShipments) | **Post** /shipments | Create Shipment
[**ExpApiShipmentsDocumentimage**](ShipmentApi.md#ExpApiShipmentsDocumentimage) | **Get** /shipments/{shipmentTrackingNumber}/get-image | Get Image
[**ExpApiShipmentsEpod**](ShipmentApi.md#ExpApiShipmentsEpod) | **Get** /shipments/{shipmentTrackingNumber}/proof-of-delivery | Electronic Proof of Delivery
[**ExpApiShipmentsImgUpload**](ShipmentApi.md#ExpApiShipmentsImgUpload) | **Patch** /shipments/{shipmentTrackingNumber}/upload-image | Upload Paperless Trade shipment (PLT) images of previously created shipment.
[**ExpApiShipmentsInvoiceDataAwb**](ShipmentApi.md#ExpApiShipmentsInvoiceDataAwb) | **Patch** /shipments/{shipmentTrackingNumber}/upload-invoice-data | Upload Commercial Invoice data for your DHL Express shipment

# **ExpApiShipments**
> SupermodelIoLogisticsExpressCreateShipmentResponse ExpApiShipments(ctx, body, optional)
Create Shipment

## Create Shipment The ShipmentRequest Operation will allow you to generate an AWB number and piece IDs, generate a shipping label, transmit manifest shipment detail to DHL, and optionally book a courier for the pickup of a shipment. The key elements in the response of the Shipment Request will be a base64 encoded PDF label and the Shipment and Piece identification numbers, which you can use for tracking on the DHL web site. While the RateRequest and ShipmentRequest services can be used independently, DHL recommends the use of RateRequest to first validate the products available for the shipper/receiver. The global product codes which are output during the RateResponse can be used directly as input into the Shipment Request, as both perform similar validations in terms of service capability. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressCreateShipmentRequest**](SupermodelIoLogisticsExpressCreateShipmentRequest.md)| Details about the shipment to be created | 
 **optional** | ***ShipmentApiExpApiShipmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiExpApiShipmentsOpts struct
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

[**SupermodelIoLogisticsExpressCreateShipmentResponse**](supermodelIoLogisticsExpressCreateShipmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiShipmentsDocumentimage**
> SupermodelIoLogisticsExpressDocumentImageResponse ExpApiShipmentsDocumentimage(ctx, shipmentTrackingNumber, shipperAccountNumber, typeCode, pickupYearAndMonth, optional)
Get Image

The Get Image service can be used to retrieve customer's own uploaded Commercial Invoice, Waybill Document  or supporting documents that uploaded during shipment creation.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentTrackingNumber** | **string**| DHL Express shipment identification number | 
  **shipperAccountNumber** | **string**| DHL Express customer shipper account number | 
  **typeCode** | **string**| Please provide correct document type. | 
  **pickupYearAndMonth** | **string**| Please provide the pickup&#x27;s date in YYYY-MM format  | 
 **optional** | ***ShipmentApiExpApiShipmentsDocumentimageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiExpApiShipmentsDocumentimageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **encodingFormat** | **optional.String**| Please provide the document image encoding format in pdf or tiff format  | 
 **allInOnePDF** | **optional.Bool**| Option to return all the document images in a single PDF file  | 
 **compressedPackage** | **optional.Bool**| Option to return all the document images in a compressed package  | 
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressDocumentImageResponse**](supermodelIoLogisticsExpressDocumentImageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiShipmentsEpod**
> SupermodelIoLogisticsExpressEpodResponse ExpApiShipmentsEpod(ctx, shipmentTrackingNumber, optional)
Electronic Proof of Delivery

The electronic proof of delivery service can be used to retrieve proof of delivery for certain delivered DHL Express shipments 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentTrackingNumber** | **string**| DHL Express shipment identification number | 
 **optional** | ***ShipmentApiExpApiShipmentsEpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiExpApiShipmentsEpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipperAccountNumber** | **optional.String**| DHL Express customer shipper account number | 
 **content** | **optional.String**|  | [default to epod-summary]
 **messageReference** | **optional.String**| Please provide message reference  | 
 **messageReferenceDate** | **optional.String**| Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2 | 
 **pluginName** | **optional.String**| Please provide name of the plugin (applicable to 3PV only)  | 
 **pluginVersion** | **optional.String**| Please provide version of the plugin (applicable to 3PV only)  | 
 **shippingSystemPlatformName** | **optional.String**| Please provide name of the shipping platform(applicable to 3PV only)  | 
 **shippingSystemPlatformVersion** | **optional.String**| Please provide version of the shipping platform (applicable to 3PV only)  | 
 **webstorePlatformName** | **optional.String**| Please provide name of the webstore platform (applicable to 3PV only)  | 
 **webstorePlatformVersion** | **optional.String**| Please provide version of the webstore platform (applicable to 3PV only)  | 

### Return type

[**SupermodelIoLogisticsExpressEpodResponse**](supermodelIoLogisticsExpressEPODResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiShipmentsImgUpload**
> ExpApiShipmentsImgUpload(ctx, body, shipmentTrackingNumber)
Upload Paperless Trade shipment (PLT) images of previously created shipment.

The upload-image service can be used to upload PLT images to a previously created shipment.  The PLT images for the shipment can be uploaded before the shipment has been physically  collected by DHL courier. However, the original shipment must contain WY as the special service otherwise,  an error will be returned when the customer wants to use the reupload function in this upload-image service.    IMPORTANT: Please note that at least 10mins must be given between the initial createShipment request and then  the upload-image request (including subsequent upload-image request).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressImageUploadRequest**](SupermodelIoLogisticsExpressImageUploadRequest.md)| Details about the shipment images | 
  **shipmentTrackingNumber** | **string**| DHL Express shipment identification number | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpApiShipmentsInvoiceDataAwb**
> ExpApiShipmentsInvoiceDataAwb(ctx, body, shipmentTrackingNumber, optional)
Upload Commercial Invoice data for your DHL Express shipment

## Upload Invoice Data with Shipment ID The upload invoice data service can be used to upload Commerical Invoice data with Shipment Identification Number for your DHL Express shipment.Customer can provide Commercial Invoice data before Shipment Data via Create Shipment flow or vice versa.  Important Note: UploadInvoiceData service is not enabled by default and must be requested per customer. Use of this service is only enabled on exceptional basis and DHL Express recommends to submit shipment requests together with a commercial invoice data.To enable use of UploadInvoiceData service, please contact your DHL Express IT representative. To use UploadInvoiceData service, it is required that \"PM\" service code is provided in MyDHL API Create Shipment request. \"PM\" service code is not enabled by  default for the customers, and needs to be enabled upon request.  When Shipment is created via MyDHL API Create Shipment service before uploading the Commercial Invoice (CIN) data,it is mandatory to provide the Shipment Identification Number as received in MyDHL API Create Shipment service Response. When Commercial Invoice (CIN) data is uploaded prior to creating a shipment via MyDHL API Create Shipment service, it is mandatory to provide Invoice Reference Number with Invoice Reference Type value \"CU\" and Shipper Account Number.   These elements are mandatory to facilitate an effective data merge of the Commercial Invoice (CIN) data with Shipment Data. As an output customer will receive Notification element value '0' on successful upload of Commercial Invoice (CIN) data.  DHL backend application performs the subsequent data merging process of the Shipment Data and Commercial Invoice data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SupermodelIoLogisticsExpressUploadInvoiceDataRequest**](SupermodelIoLogisticsExpressUploadInvoiceDataRequest.md)| Details about the invoice data | 
  **shipmentTrackingNumber** | **string**| DHL Express shipment identification number | 
 **optional** | ***ShipmentApiExpApiShipmentsInvoiceDataAwbOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiExpApiShipmentsInvoiceDataAwbOpts struct
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

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

