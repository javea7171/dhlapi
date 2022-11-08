# SupermodelIoLogisticsExpressUpdatePickupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DispatchConfirmationNumber** | **string** | Please enter Dispatch confirmation number which identifies the already scheduled pickup | [default to null]
**OriginalShipperAccountNumber** | **string** | Please enter the shipper account number which was used during the scheduled pickup creation | [default to null]
**PlannedPickupDateAndTime** | **string** | Identifies the date and time the package is ready for pickup Both the date and time portions of the string are expected to be used. The date should not be a past date or a date more than 10 days in the future. The time is the local time of the shipment based on the shipper&#x27;s time zone. The date component must be in the format: YYYY-MM-DD; the time component must be in the format: HH:MM:SS using a 24 hour clock. The date and time parts are separated by the letter T (e.g. 2006-06-26T17:00:00 GMT+01:00).&lt;BR&gt;                            | [default to null]
**CloseTime** | **string** | The latest time the location premises is available to dispatch the DHL Express shipment. (HH:MM)  | [optional] [default to null]
**Location** | **string** | Provides information on where the package should be picked up by DHL courier. &lt;BR&gt;                            | [optional] [default to null]
**LocationType** | **string** | Provides information on where the package should be picked up by DHL courier. &lt;BR&gt;                            | [optional] [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) |  | [default to null]
**SpecialInstructions** | [**[]SupermodelIoLogisticsExpressPickupRequestSpecialInstructions**](supermodelIoLogisticsExpressPickupRequest_specialInstructions.md) | Details special pickup instructions you may wish to send to the DHL Courier. | [optional] [default to null]
**Remark** | **string** | Please provide additional pickup remark | [optional] [default to null]
**CustomerDetails** | [***SupermodelIoLogisticsExpressPickupRequestCustomerDetails**](supermodelIoLogisticsExpressPickupRequest_customerDetails.md) |  | [default to null]
**ShipmentDetails** | [**[]SupermodelIoLogisticsExpressUpdatePickupRequestShipmentDetails**](supermodelIoLogisticsExpressUpdatePickupRequest_shipmentDetails.md) | Please provide updated details related to shipment you want update the pickup for | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

