# SupermodelIoLogisticsExpressUpdatePickupRequestShipmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | Please provide DHL Express Global product code of the shipment | [default to null]
**LocalProductCode** | **string** | Please provide DHL Express Local product code of the shipment | [optional] [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) |  | [optional] [default to null]
**ValueAddedServices** | [**[]SupermodelIoLogisticsExpressValueAddedServicesRates**](supermodelIoLogisticsExpressValueAddedServicesRates.md) |  | [optional] [default to null]
**IsCustomsDeclarable** | **bool** | For customs purposes please advise if your shipment is dutiable (true) or non dutiable (false) | [default to null]
**DeclaredValue** | **float64** | For customs purposes please advise on declared value of the shipment | [optional] [default to null]
**DeclaredValueCurrency** | **string** | For customs purposes please advise on declared value currency code of the shipment | [optional] [default to null]
**UnitOfMeasurement** | **string** | Please enter Unit of measurement - metric,imperial | [default to null]
**ShipmentTrackingNumber** | **string** | Please provide Shipment Identification number (AWB number) | [optional] [default to null]
**Packages** | [**[]SupermodelIoLogisticsExpressPackageRr**](supermodelIoLogisticsExpressPackageRR.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

