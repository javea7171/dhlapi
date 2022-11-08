# SupermodelIoLogisticsExpressRatesPickupCapabilities

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextBusinessDay** | **bool** | This indicator has values of Y or N, and tells the consumer if the service in the response has a pickup date on the same day as the requested shipment date (per the request). | [optional] [default to null]
**LocalCutoffDateAndTime** | **string** | This is the cutoff time for the service&lt;BR&gt;                offered in the response. This represents the latest time (local to origin) which the shipment can be tendered to the courier for that service on that day. | [optional] [default to null]
**GMTCutoffTime** | **string** | Pickup cut off time in GMT | [optional] [default to null]
**PickupEarliest** | **string** | The DHL earliest time possible for pickup | [optional] [default to null]
**PickupLatest** | **string** | The DHL latest time possible for pickup | [optional] [default to null]
**OriginServiceAreaCode** | **string** | The DHL Service Area Code for the origin of the Shipment | [optional] [default to null]
**OriginFacilityAreaCode** | **string** | The DHL Facility Code for the Origin | [optional] [default to null]
**PickupAdditionalDays** | **float64** | This is additional transit delays (in days) for shipment picked up from the mentioned city or postal area to arrival at the service area. | [optional] [default to null]
**PickupDayOfWeek** | **float64** | Pickup day of the week number | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

