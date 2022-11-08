# SupermodelIoLogisticsExpressCreateShipmentResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL where the request has been sent to | [optional] [default to null]
**ShipmentTrackingNumber** | **string** | Here you will receive Shipment Identification Number of your package | [optional] [default to null]
**CancelPickupUrl** | **string** | If you requested pickup for your shipment you can use this URL to cancel the pickup | [optional] [default to null]
**TrackingUrl** | **string** | You can use ths URL to track your shipment | [optional] [default to null]
**DispatchConfirmationNumber** | **string** | If you asked for pickup service here you will find Dispach Confirmation Number which identifies your pickup booking | [optional] [default to null]
**Packages** | [**[]SupermodelIoLogisticsExpressCreateShipmentResponsePackages**](supermodelIoLogisticsExpressCreateShipmentResponse_packages.md) | Here you can find information for all pieces your shipment is having like Piece Identification Number  | [optional] [default to null]
**Documents** | [**[]SupermodelIoLogisticsExpressCreateShipmentResponseDocuments1**](supermodelIoLogisticsExpressCreateShipmentResponse_documents_1.md) | Here you can find all documents created for the shipment like Transport and WaybillDoc labels, Invoice, Receipt | [optional] [default to null]
**OnDemandDeliveryURL** | **string** | In this field you will find the On Demand Delivery (ODD) URL link if requested | [optional] [default to null]
**ShipmentDetails** | [**[]SupermodelIoLogisticsExpressCreateShipmentResponseShipmentDetails**](supermodelIoLogisticsExpressCreateShipmentResponse_shipmentDetails.md) | Here you can find additional information related to your shipment when you ask for it in the request | [optional] [default to null]
**ShipmentCharges** | [**[]SupermodelIoLogisticsExpressCreateShipmentResponseShipmentCharges**](supermodelIoLogisticsExpressCreateShipmentResponse_shipmentCharges.md) | Here you can find rates related to your shipment | [optional] [default to null]
**BarcodeInfo** | [***SupermodelIoLogisticsExpressCreateShipmentResponseBarcodeInfo**](supermodelIoLogisticsExpressCreateShipmentResponse_barcodeInfo.md) |  | [optional] [default to null]
**EstimatedDeliveryDate** | [***SupermodelIoLogisticsExpressCreateShipmentResponseEstimatedDeliveryDate**](supermodelIoLogisticsExpressCreateShipmentResponse_estimatedDeliveryDate.md) |  | [optional] [default to null]
**Warnings** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

