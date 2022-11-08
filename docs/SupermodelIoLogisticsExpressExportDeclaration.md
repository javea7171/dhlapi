# SupermodelIoLogisticsExpressExportDeclaration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItems** | [**[]SupermodelIoLogisticsExpressExportDeclarationLineItems**](supermodelIoLogisticsExpressExportDeclaration_lineItems.md) | Please enter details for each export line item | [default to null]
**Invoice** | [***SupermodelIoLogisticsExpressExportDeclarationInvoice**](supermodelIoLogisticsExpressExportDeclaration_invoice.md) |  | [default to null]
**Remarks** | [**[]SupermodelIoLogisticsExpressExportDeclarationRemarks**](supermodelIoLogisticsExpressExportDeclaration_remarks.md) | Please enter up to three remarks | [optional] [default to null]
**AdditionalCharges** | [**[]SupermodelIoLogisticsExpressExportDeclarationAdditionalCharges**](supermodelIoLogisticsExpressExportDeclaration_additionalCharges.md) | Please enter additional charge to appear on the invoice&lt;BR&gt;      admin, Administration Charge&lt;BR&gt;      delivery, Delivery Charge&lt;BR&gt;      documentation, Documentation Charge&lt;BR&gt;      expedite, Expedite Charge&lt;BR&gt;      freight, Freight Charge&lt;BR&gt;      fuel surcharge, Fuel Surcharge&lt;BR&gt;      logistic, Logistic Charge&lt;BR&gt;      other, Other Charge&lt;BR&gt;      packaging, Packaging Charge&lt;BR&gt;      pickup, Pickup Charge&lt;BR&gt;      handling, Handling Charge&lt;BR&gt;      vat, VAT Charge&lt;BR&gt;      insurance, Insurance Cost | [optional] [default to null]
**PlaceOfIncoterm** | **string** | Name of port of departure, shipment or destination as required under the applicable delivery term. | [optional] [default to null]
**RecipientReference** | **string** | Please enter recipient reference | [optional] [default to null]
**Exporter** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationExporter**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_exporter.md) |  | [optional] [default to null]
**ExportReasonType** | **string** | Please provide the reason for export | [optional] [default to null]
**ShipmentType** | **string** | Please provide the shipment was sent for Personal (Gift) or Commercial (Sale) reasons | [optional] [default to null]
**CustomsDocuments** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments1**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_customsDocuments_1.md) | Please provide the Customs Documents at invoice level | [optional] [default to null]
**Incoterm** | **string** | The Incoterms rules are a globally-recognized set of standards, used worldwide in international and domestic contracts for the delivery of goods, illustrating responsibilities between buyer and seller for costs and risk, as well as cargo insurance.&lt;BR&gt;      EXW ExWorks&lt;BR&gt;      FCA Free Carrier&lt;BR&gt;      CPT Carriage Paid To&lt;BR&gt;      CIP Carriage and Insurance Paid To&lt;BR&gt;      DPU Delivered at Place Unloaded&lt;BR&gt;      DAP Delivered at Place&lt;BR&gt;      DDP Delivered Duty Paid&lt;BR&gt;      FAS Free Alongside Ship&lt;BR&gt;      FOB Free on Board&lt;BR&gt;      CFR Cost and Freight&lt;BR&gt;      CIF Cost, Insurance and Freight&lt;BR&gt;      DAF Delivered at Frontier&lt;BR&gt;      DAT Delivered at Terminal&lt;BR&gt;      DDU Delivered Duty Unpaid&lt;BR&gt;      DEQ Delivery ex Quay&lt;BR&gt;      DES Delivered ex Ship | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

