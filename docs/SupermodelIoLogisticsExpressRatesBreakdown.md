# SupermodelIoLogisticsExpressRatesBreakdown

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | When landed-cost is requested then following items name (Charge Types) might be returned: &lt;BR&gt;                        Charge Type : Description &lt;BR&gt;                        STDIS : Quoted shipment total discount &lt;BR&gt;                        SCUSV : Shipment Customs value &lt;BR&gt;                        SINSV : Insured value &lt;BR&gt;                        SPRQD : Shipment product quote discount&lt;BR&gt;                        SPRQN : The price quoted to the Customer by DHL at the time of the booking. This quote covers the weight price including discounts and without taxes. &lt;BR&gt;                        STSCH : The total of service charges quoted to customer for DHL Express value added services, the amount is after discounts and doesn&#x27;t include tax amounts. &lt;BR&gt;                        MACHG : The total of service charges as provided by Merchant for the purpose of landed cost calculation. &lt;BR&gt;                        MFCHG : The freight charge as provided by Merchant for the purpose of landed cost calculation. | [optional] [default to null]
**ServiceCode** | **string** | Special service or extra charge code. This is the code you would have to use in the /shipment service if you wish to add an optional Service such as Saturday delivery | [optional] [default to null]
**LocalServiceCode** | **string** | Local service code | [optional] [default to null]
**TypeCode** | **string** | Price breakdown type code | [optional] [default to null]
**ServiceTypeCode** | **string** | Special service charge code type for service. | [optional] [default to null]
**Price** | **float64** | Price breakdown value | [optional] [default to null]
**PriceCurrency** | **string** | This the currency of the rated shipment for the prices listed. | [optional] [default to null]
**IsCustomerAgreement** | **bool** | Customer agreement indicator for product and services, if service is offered with prior customer agreement | [optional] [default to null]
**IsMarketedService** | **bool** | Indicator if the special service is marketed service | [optional] [default to null]
**IsBillingServiceIndicator** | **bool** | Indicator if there is any discount allowed | [optional] [default to null]
**PriceBreakdown** | [**[]SupermodelIoLogisticsExpressRatesPriceBreakdown1**](supermodelIoLogisticsExpressRates_priceBreakdown_1.md) |  | [optional] [default to null]
**TariffRateFormula** | **string** | Tariff Rate Formula on Shipment Level | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

