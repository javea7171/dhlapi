# SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItems** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLineItems**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_lineItems.md) | Please enter details for each export line item | [default to null]
**Invoice** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoice**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_invoice.md) |  | [optional] [default to null]
**Remarks** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationRemarks**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_remarks.md) | Please enter up to three remarks. &lt;BR&gt;              If using Customs Invoice template COMMERCIAL_INVOICE_04, the invoice can only print the first remarks field. The recommended max length is 20 characters. &lt;BR&gt;              If using Customs Invoice template COMMERCIAL_INVOICE_L_10 or COMMERCIAL_INVOICE_P_10, the invoice can print all three remraks fields.  The recommended max length is 45 characters. | [optional] [default to null]
**AdditionalCharges** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationAdditionalCharges**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_additionalCharges.md) | Please enter additional charge to appear on the invoice&lt;BR&gt;              admin, Administration Charge&lt;BR&gt;              delivery, Delivery Charge&lt;BR&gt;              documentation, Documentation Charge&lt;BR&gt;              expedite, Expedite Charge&lt;BR&gt;              export, Export Charge&lt;BR&gt;              freight, Freight Charge&lt;BR&gt;              fuel_surcharge, Fuel Surcharge&lt;BR&gt;              logistic, Logistic Charge&lt;BR&gt;              other, Other Charge&lt;BR&gt;              packaging, Packaging Charge&lt;BR&gt;              pickup, Pickup Charge&lt;BR&gt;              handling, Handling Charge&lt;BR&gt;              vat, VAT Charge&lt;BR&gt;              insurance, Insurance Cost&lt;BR&gt;              reverse_charge, Reverse Charge | [optional] [default to null]
**DestinationPortName** | **string** | Please provide destination port details | [optional] [default to null]
**PlaceOfIncoterm** | **string** | Name of port of departure, shipment or destination as required under the applicable delivery term. | [optional] [default to null]
**PayerVATNumber** | **string** | Please provide Payer VAT number | [optional] [default to null]
**RecipientReference** | **string** | Please enter recipient reference | [optional] [default to null]
**Exporter** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationExporter**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_exporter.md) |  | [optional] [default to null]
**PackageMarks** | **string** | Please enter package marks | [optional] [default to null]
**DeclarationNotes** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationDeclarationNotes**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_declarationNotes.md) | Please provide up to three dcelaration notes | [optional] [default to null]
**ExportReference** | **string** | Please enter export reference | [optional] [default to null]
**ExportReason** | **string** | Please enter export reason | [optional] [default to null]
**ExportReasonType** | **string** | Please provide the reason for export | [optional] [default to null]
**Licenses** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLicenses**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_licenses.md) | Please provide details about export and import licenses | [optional] [default to null]
**ShipmentType** | **string** | Please provide the shipment was sent for Personal (Gift) or Commercial (Sale) reasons | [optional] [default to null]
**CustomsDocuments** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments1**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_customsDocuments_1.md) | Please provide the Customs Documents at invoice level | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

