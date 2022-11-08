# SupermodelIoLogisticsExpressExportDeclarationLineItems

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Please provide line item number | [default to null]
**Description** | **string** | Please provide description of the line item | [default to null]
**Price** | **float64** | Please provide monetary value of the line item | [default to null]
**Quantity** | [***SupermodelIoLogisticsExpressExportDeclarationQuantity**](supermodelIoLogisticsExpressExportDeclaration_quantity.md) |  | [default to null]
**CommodityCodes** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCommodityCodes**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_commodityCodes.md) | Please provide Commodity codes for the shipment at item line level | [optional] [default to null]
**ExportReasonType** | **string** | Please provide the reason for export | [optional] [default to null]
**ManufacturerCountry** | **string** | Please enter two letter ISO manufacturer country code | [default to null]
**Weight** | [***SupermodelIoLogisticsExpressExportDeclarationWeight**](supermodelIoLogisticsExpressExportDeclaration_weight.md) |  | [default to null]
**IsTaxesPaid** | **bool** | Please provide if the Taxes is paid for the line item | [optional] [default to null]
**CustomerReferences** | [**[]SupermodelIoLogisticsExpressExportDeclarationCustomerReferences**](supermodelIoLogisticsExpressExportDeclaration_customerReferences.md) | Please provide the Customer References for the line item | [optional] [default to null]
**CustomsDocuments** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration_customsDocuments.md) | Please provide the customs documents details | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

