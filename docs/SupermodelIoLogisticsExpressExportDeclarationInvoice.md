# SupermodelIoLogisticsExpressExportDeclarationInvoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Please enter commercial invoice number | [default to null]
**Date** | **string** | Please enter commercial invoice date | [default to null]
**Function** | **string** | Please provide the purpose was the document details captured and are planned to be used. Note: export and import is only applicable for approve Sale In Transit customers | [default to null]
**CustomerReferences** | [**[]SupermodelIoLogisticsExpressExportDeclarationInvoiceCustomerReferences**](supermodelIoLogisticsExpressExportDeclaration_invoice_customerReferences.md) | Please provide the customer references at invoice level.   Note: customerReference/0/value with typeCode &#x27;CU&#x27; is mandatory if using POST method and no shipmentTrackingNumber is provided in request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

