# SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Please enter commercial invoice number | [default to null]
**Date** | **string** | Please enter commercial invoice date | [default to null]
**SignatureName** | **string** | Please enter who has signed the invoce | [optional] [default to null]
**SignatureTitle** | **string** | Please provide title of person who has signed the invoice | [optional] [default to null]
**SignatureImage** | **string** | Please provide the signature image | [optional] [default to null]
**Instructions** | **[]string** | Shipment instructions for customs invoice printing purposes. Printed only when using Customs Invoice template COMMERCIAL_INVOICE_04. If using Customs Invoice template    COMMERCIAL_INVOICE_04, recommended max length is 120 characters. | [optional] [default to null]
**CustomerDataTextEntries** | **[]string** | Customer data text to be printed in&lt;BR&gt;                  customs invoice.&lt;BR&gt;                  Printed only when using Customs&lt;BR&gt;                  Invoice template&lt;BR&gt;                  COMMERCIAL_INVOICE_04. | [optional] [default to null]
**TotalNetWeight** | **float64** | Please provide the total net weight | [optional] [default to null]
**TotalGrossWeight** | **float64** | Please provide the total gross weight | [optional] [default to null]
**CustomerReferences** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceCustomerReferences**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_invoice_customerReferences.md) | Please provide the customer references at invoice level | [optional] [default to null]
**TermsOfPayment** | **string** | Please provide the terms of payment | [optional] [default to null]
**IndicativeCustomsValues** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceIndicativeCustomsValues**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_invoice_indicativeCustomsValues.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

