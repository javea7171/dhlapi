# SupermodelIoLogisticsExpressCreateShipmentRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Packages** | [**[]SupermodelIoLogisticsExpressPackage**](supermodelIoLogisticsExpressPackage.md) | Here you can define properties per package | [default to null]
**IsCustomsDeclarable** | **bool** | For customs purposes please advise if your shipment is dutiable (true) or non dutiable (false).Note:If the shipment is dutiable, exportDeclaration element must be provided. | [default to null]
**DeclaredValue** | **float64** | For customs purposes please advise on declared value of the shipment | [optional] [default to null]
**DeclaredValueCurrency** | **string** | For customs purposes please advise on declared value currency code of the shipment | [optional] [default to null]
**ExportDeclaration** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration**](supermodelIoLogisticsExpressCreateShipmentRequest_content_exportDeclaration.md) |  | [optional] [default to null]
**Description** | **string** | Please enter description of your shipment | [default to null]
**USFilingTypeValue** | **string** | This is used for the US AES4, FTR and ITN numbers to be printed on the Transport Label | [optional] [default to null]
**Incoterm** | **string** | The Incoterms rules are a globally-recognized set of standards, used worldwide in international and domestic contracts for the delivery of goods, illustrating responsibilities between buyer and seller for costs and risk, as well as cargo insurance.&lt;BR&gt;          EXW ExWorks&lt;BR&gt;          FCA Free Carrier&lt;BR&gt;          CPT Carriage Paid To&lt;BR&gt;          CIP Carriage and Insurance Paid To&lt;BR&gt;          DPU Delivered at Place Unloaded&lt;BR&gt;          DAP Delivered at Place&lt;BR&gt;          DDP Delivered Duty Paid&lt;BR&gt;          FAS Free Alongside Ship&lt;BR&gt;          FOB Free on Board&lt;BR&gt;          CFR Cost and Freight&lt;BR&gt;          CIF Cost, Insurance and Freight&lt;BR&gt;          DAF Delivered at Frontier&lt;BR&gt;          DAT Delivered at Terminal&lt;BR&gt;          DDU Delivered Duty Unpaid&lt;BR&gt;          DEQ Delivery ex Quay&lt;BR&gt;          DES Delivered ex Ship | [default to null]
**UnitOfMeasurement** | **string** | Please enter Unit of measurement - metric,imperial | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

