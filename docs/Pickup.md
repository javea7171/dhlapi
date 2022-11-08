# Pickup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRequested** | **bool** | Please advise if a pickup is needed for this shipment | [default to false]
**CloseTime** | **string** | The latest time the location premises is available to dispatch the DHL Express shipment. (HH:MM)  | [optional] [default to null]
**Location** | **string** | Provides information on where the package should be picked up by DHL courier. | [optional] [default to null]
**SpecialInstructions** | [**[]PickupSpecialInstructions**](pickup_specialInstructions.md) | Details special pickup instructions you may wish to send to the DHL Courier. | [optional] [default to null]
**PickupDetails** | [***PickupPickupDetails**](pickup_pickupDetails.md) |  | [optional] [default to null]
**PickupRequestorDetails** | [***PickupPickupRequestorDetails**](pickup_pickupRequestorDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

