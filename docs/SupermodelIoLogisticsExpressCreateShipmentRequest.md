# SupermodelIoLogisticsExpressCreateShipmentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedShippingDateAndTime** | **string** | Identifies the date and time the package is tendered. Both the date and time portions of the string are expected to be used. The date should not be a past date or a date more than 10 days in the future. The time is the local time of the shipment based on the shipper&#x27;s time zone. The date component must be in the format: YYYY-MM-DD; the time component must be in the format: HH:MM:SS using a 24 hour clock. The date and time parts are separated by the letter T (e.g. 2006-06-26T17:00:00 GMT+01:00). | [default to null]
**Pickup** | [***Pickup**](pickup.md) |  | [default to null]
**ProductCode** | **string** | Please enter DHL Express Global Product code | [default to null]
**LocalProductCode** | **string** | Please enter DHL Express Local Product code. Important when shipping domestic products. | [optional] [default to null]
**GetRateEstimates** | **bool** | Please advise if you want to get rate estimates for given shipment | [optional] [default to false]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) | Please enter all the DHL Express accounts and types to be used for this shipment | [default to null]
**ValueAddedServices** | [**[]SupermodelIoLogisticsExpressValueAddedServices**](supermodelIoLogisticsExpressValueAddedServices.md) | This section communicates additional shipping services, such as Insurance (or Shipment Value Protection). | [optional] [default to null]
**OutputImageProperties** | [***SupermodelIoLogisticsExpressCreateShipmentRequestOutputImageProperties**](supermodelIoLogisticsExpressCreateShipmentRequest_outputImageProperties.md) |  | [optional] [default to null]
**CustomerReferences** | [**[]SupermodelIoLogisticsExpressReference**](supermodelIoLogisticsExpressReference.md) | Here you can declare your customer references | [optional] [default to null]
**Identifiers** | [**[]SupermodelIoLogisticsExpressIdentifier**](supermodelIoLogisticsExpressIdentifier.md) | Identifiers section is on the shipment level where you can optionaly provide a DHL Express waybill number. This has to be enabled by your DHL Express IT contact. | [optional] [default to null]
**CustomerDetails** | [***SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetails**](supermodelIoLogisticsExpressCreateShipmentRequest_customerDetails.md) |  | [default to null]
**Content** | [***SupermodelIoLogisticsExpressCreateShipmentRequestContent**](supermodelIoLogisticsExpressCreateShipmentRequest_content.md) |  | [default to null]
**DocumentImages** | [***[]SupermodelIoLogisticsExpressDocumentImagesInner**](array.md) |  | [optional] [default to null]
**OnDemandDelivery** | [***SupermodelIoLogisticsExpressCreateShipmentRequestOnDemandDelivery**](supermodelIoLogisticsExpressCreateShipmentRequest_onDemandDelivery.md) |  | [optional] [default to null]
**RequestOndemandDeliveryURL** | **bool** | Determines whether to request the On Demand Delivery (ODD) link. When set to true it will provide an URL link for the specified Waybill Number, Shipper Account Number. The default value is false, no ODD link URL is provided in the response message. | [optional] [default to null]
**ShipmentNotification** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestShipmentNotification**](supermodelIoLogisticsExpressCreateShipmentRequest_shipmentNotification.md) | This is to support sending email notification once the shipment is created. The email will contain the basic information on the shipper, recipient, waybill number, and shipment information | [optional] [default to null]
**PrepaidCharges** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestPrepaidCharges**](supermodelIoLogisticsExpressCreateShipmentRequest_prepaidCharges.md) | Please provide any charges you have already paid for this shipment, like freight paid upfront. To allow using this section please contact your DHL Express representative | [optional] [default to null]
**GetTransliteratedResponse** | **bool** | If set to true, response will return transliterated text of shipper and receiver details. | [optional] [default to null]
**EstimatedDeliveryDate** | [***EstimatedDeliveryDate**](estimated delivery date.md) |  | [optional] [default to null]
**GetAdditionalInformation** | [**[]SupermodelIoLogisticsExpressCreateShipmentRequestGetAdditionalInformation**](supermodelIoLogisticsExpressCreateShipmentRequest_getAdditionalInformation.md) | Provides additional information in the response like service area details, routing code and pickup-related information | [optional] [default to null]
**ParentShipment** | [***SupermodelIoLogisticsExpressCreateShipmentRequestParentShipment**](supermodelIoLogisticsExpressCreateShipmentRequest_parentShipment.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

