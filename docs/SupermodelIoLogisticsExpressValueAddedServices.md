# SupermodelIoLogisticsExpressValueAddedServices

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCode** | **string** | Please enter DHL Express value added service code. For detailed list of all available service codes for your prospect shipment please invoke GET /products or GET /rates | [default to null]
**Value** | **float64** | Please enter monetary value of service (e.g. Insured Value) | [optional] [default to null]
**Currency** | **string** | Please enter currency code (e.g. Insured Value currency code) | [optional] [default to null]
**Method** | **string** | Payment method code (e.g. Cash) | [optional] [default to null]
**DangerousGoods** | [**[]SupermodelIoLogisticsExpressValueAddedServicesDangerousGoods**](supermodelIoLogisticsExpressValueAddedServices_dangerousGoods.md) | The DangerousGoods section indicates if there is dangerous good content within the shipment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

