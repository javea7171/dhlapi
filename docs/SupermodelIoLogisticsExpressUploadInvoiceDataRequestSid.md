# SupermodelIoLogisticsExpressUploadInvoiceDataRequestSid

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentTrackingNumber** | **string** | Please provide Shipment Identification number (AWB number) | [optional] [default to null]
**PlannedShipDate** | **string** | The planned shipment date for the provided shipmentTrackingNumber.  The date must be in the format: YYYY-MM-DD | [optional] [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) | Please enter all the DHL Express accounts and types to be used for this shipment.   Note: accounts/0/number with typeCode &#x27;shipper&#x27; is mandatory if using POST method and no shipmentTrackingNumber is provided in request. | [optional] [default to null]
**Content** | [***SupermodelIoLogisticsExpressUploadInvoiceDataRequestContent**](supermodelIoLogisticsExpressUploadInvoiceDataRequest_content.md) |  | [default to null]
**OutputImageProperties** | [***SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImageProperties**](supermodelIoLogisticsExpressUploadInvoiceDataRequest_outputImageProperties.md) |  | [optional] [default to null]
**CustomerDetails** | [***SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetails**](supermodelIoLogisticsExpressUploadInvoiceDataRequest_customerDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

