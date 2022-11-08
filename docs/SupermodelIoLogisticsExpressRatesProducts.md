# SupermodelIoLogisticsExpressRatesProducts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** | Name of the DHL Express product | [optional] [default to null]
**ProductCode** | **string** | This is the global DHL Express product code for which the delivery is feasible respecting the input data from the request. | [optional] [default to null]
**LocalProductCode** | **string** | This is the local DHL Express product code for which the delivery is feasible respecting the input data from the request. | [optional] [default to null]
**LocalProductCountryCode** | **string** | The country code for the local service used | [optional] [default to null]
**NetworkTypeCode** | **string** | The NetworkTypeCode element indicates the product belongs to the Day Definite (DD) or Time Definite (TD) network.&lt;BR&gt;            Possible Values;&lt;BR&gt;             DD: Day Definite product&lt;BR&gt;             TD: Time Definite product | [optional] [default to null]
**IsCustomerAgreement** | **bool** | Indicator that the product only can be offered to customers with prior agreement. | [optional] [default to null]
**Weight** | [***Weight1**](weight_1.md) |  | [default to null]
**TotalPrice** | [**[]SupermodelIoLogisticsExpressRatesTotalPrice**](supermodelIoLogisticsExpressRates_totalPrice.md) |  | [default to null]
**TotalPriceBreakdown** | [**[]SupermodelIoLogisticsExpressRatesTotalPriceBreakdown**](supermodelIoLogisticsExpressRates_totalPriceBreakdown.md) |  | [optional] [default to null]
**DetailedPriceBreakdown** | [**[]SupermodelIoLogisticsExpressRatesDetailedPriceBreakdown**](supermodelIoLogisticsExpressRates_detailedPriceBreakdown.md) |  | [optional] [default to null]
**ServiceCodeMutuallyExclusiveGroups** | [**[]SupermodelIoLogisticsExpressProductsServiceCodeMutuallyExclusiveGroups**](supermodelIoLogisticsExpressProducts_serviceCodeMutuallyExclusiveGroups.md) | Group of serviceCodes that are mutually exclusive.  Only one serviceCode among the list must be applied for a shipment | [optional] [default to null]
**ServiceCodeDependencyRuleGroups** | [**[]SupermodelIoLogisticsExpressProductsServiceCodeDependencyRuleGroups**](supermodelIoLogisticsExpressProducts_serviceCodeDependencyRuleGroups.md) | Dependency rule groups for a particular serviceCode. | [optional] [default to null]
**PickupCapabilities** | [***SupermodelIoLogisticsExpressRatesPickupCapabilities**](supermodelIoLogisticsExpressRates_pickupCapabilities.md) |  | [optional] [default to null]
**DeliveryCapabilities** | [***SupermodelIoLogisticsExpressRatesDeliveryCapabilities**](supermodelIoLogisticsExpressRates_deliveryCapabilities.md) |  | [optional] [default to null]
**Items** | [**[]SupermodelIoLogisticsExpressRatesItems**](supermodelIoLogisticsExpressRates_items.md) |  | [optional] [default to null]
**PricingDate** | **string** | The date when the rates for DHL products and services is provided | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

