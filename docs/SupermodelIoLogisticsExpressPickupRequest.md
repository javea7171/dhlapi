# SupermodelIoLogisticsExpressPickupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedPickupDateAndTime** | **string** | Identifies the date and time the package is ready for pickup Both the date and time portions of the string are expected to be used. The date should not be a past date or a date more than 10 days in the future. The time is the local time of the shipment based on the shipper&#x27;s time zone. The date component must be in the format: YYYY-MM-DD; the time component must be in the format: HH:MM:SS using a 24 hour clock. The date and time parts are separated by the letter T (e.g. 2006-06-26T17:00:00 GMT+01:00).&lt;BR&gt;                            | [default to null]
**CloseTime** | **string** | The latest time the location premises is available to dispatch the DHL Express shipment. (HH:MM)  | [optional] [default to null]
**Location** | **string** | Provides information on where the package should be picked up by DHL courier. &lt;BR&gt;                            | [optional] [default to null]
**LocationType** | **string** | Provides information on where the package should be picked up by DHL courier. &lt;BR&gt;                            | [optional] [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) |  | [default to null]
**SpecialInstructions** | [**[]SupermodelIoLogisticsExpressPickupRequestSpecialInstructions**](supermodelIoLogisticsExpressPickupRequest_specialInstructions.md) | Details special pickup instructions you may wish to send to the DHL Courier. | [optional] [default to null]
**Remark** | **string** | Please provide additional pickup remark | [optional] [default to null]
**CustomerDetails** | [***SupermodelIoLogisticsExpressPickupRequestCustomerDetails**](supermodelIoLogisticsExpressPickupRequest_customerDetails.md) |  | [default to null]
**ShipmentDetails** | [**[]SupermodelIoLogisticsExpressPickupRequestShipmentDetails**](supermodelIoLogisticsExpressPickupRequest_shipmentDetails.md) | Please provide details related to shipment you want to do the pickup for | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

