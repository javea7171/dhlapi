# SupermodelIoLogisticsExpressCreateShipmentRequestOnDemandDelivery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryOption** | **string** | Please choose from one of the delivery options | [default to null]
**Location** | **string** | If delivery option is signatureDelivery please specify location where to leave the shipment | [optional] [default to null]
**SpecialInstructions** | **string** | Please enter additional information that might be useful for the DHL Express courier | [optional] [default to null]
**GateCode** | **string** | Please provide entry code to gain access to an apartment complex or gate | [optional] [default to null]
**WhereToLeave** | **string** | In ase your deliveryOption is &#x27;neighbour&#x27; please specify where to leave the package  | [optional] [default to null]
**NeighbourName** | **string** | In case you wish to leave the package with neighbour please provide the neighbour&#x27;s name | [optional] [default to null]
**NeighbourHouseNumber** | **string** | In case you wish to leave the package with neighbour please provide the neighbour&#x27;s house number | [optional] [default to null]
**AuthorizerName** | **string** | In case your delivery option is &#x27;signatureRelease&#x27; please provide name of the person who is authorized to sign and receive the package | [optional] [default to null]
**ServicePointId** | **string** | In case your delivery option is &#x27;servicepoint&#x27; please provide unique DHL Express Service point location ID of where the parcel should be delieverd (please contact your local DHL Express Account Manager to obtain the list of the servicepoint IDs) | [optional] [default to null]
**RequestedDeliveryDate** | **string** | for future use | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

