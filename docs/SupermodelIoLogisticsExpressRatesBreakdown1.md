# SupermodelIoLogisticsExpressRatesBreakdown1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the charge | [optional] [default to null]
**ServiceCode** | **string** | Special service or extra charge code. This is the code you would have to use in the /shipment service if you wish to add an optional Service such as Saturday delivery | [optional] [default to null]
**LocalServiceCode** | **string** | Local service code | [optional] [default to null]
**TypeCode** | **string** | Charge type or category.&lt;BR&gt;                        Possible values;&lt;BR&gt;                        - DUTY&lt;BR&gt;                        - TAX&lt;BR&gt;                        - FEE | [default to null]
**ServiceTypeCode** | **string** | Special service charge code type for service. XCH type charge codes are Optional Services and should be displayed to users for selection.&lt;BR&gt;                        The possible values are;&lt;BR&gt;                        - XCH &#x3D; Extra charge&lt;BR&gt;                        - FEE &#x3D; Fee&lt;BR&gt;                        - SCH &#x3D; Surcharge&lt;BR&gt;                        - NRI &#x3D; Non Revenue Item&lt;BR&gt;                        Other charges may be automatically returned when applicable. | [optional] [default to null]
**Price** | **float64** | The charge amount of the line item charge. | [default to null]
**PriceCurrency** | **string** | This the currency of the rated shipment for the prices listed. | [optional] [default to null]
**IsCustomerAgreement** | **bool** | Customer agreement indicator for product and services, if service is offered with prior customer agreement | [optional] [default to null]
**IsMarketedService** | **bool** | Indicator if the special service is marketed service | [optional] [default to null]
**IsBillingServiceIndicator** | **bool** | Indicator if there is any discount allowed | [optional] [default to null]
**PriceBreakdown** | [**[]SupermodelIoLogisticsExpressRatesPriceBreakdown2**](supermodelIoLogisticsExpressRates_priceBreakdown_2.md) |  | [optional] [default to null]
**TariffRateFormula** | **string** | Tariff Rate Formula on Line Item Level | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

