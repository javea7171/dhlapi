# SupermodelIoLogisticsExpressRatesDeliveryCapabilities

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryTypeCode** | **string** | Delivery Date capabilities considering customs clearance days.&lt;BR&gt;                QDDF: is the fastest (\&quot;docs\&quot;) transit time as quoted to the customer at booking or shipment creation. No custom clearance is considered.&lt;BR&gt;                QDDC: constitutes DHL&#x27;s service commitment as quoted at booking/shipment creation. QDDc builds in clearance time, and potentially other special operational non-transport component(s), when relevant. | [optional] [default to null]
**EstimatedDeliveryDateAndTime** | **string** | This is the estimated date/time the shipment will be delivered by for the rated shipment and product listed | [optional] [default to null]
**DestinationServiceAreaCode** | **string** | The DHL Service Area Code for the destination of the Shipment | [optional] [default to null]
**DestinationFacilityAreaCode** | **string** | The DHL Facility Code for the Destination | [optional] [default to null]
**DeliveryAdditionalDays** | **float64** | This is additional transit delays (in days) for shipment delivered to the&lt;BR&gt;                mentioned city or postal area following arrival at the service area. | [optional] [default to null]
**DeliveryDayOfWeek** | **float64** | Destination day of the week number | [optional] [default to null]
**TotalTransitDays** | **float64** | The number of transit days | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

