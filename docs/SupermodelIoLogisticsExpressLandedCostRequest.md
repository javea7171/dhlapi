# SupermodelIoLogisticsExpressLandedCostRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerDetails** | [***SupermodelIoLogisticsExpressLandedCostRequestCustomerDetails**](supermodelIoLogisticsExpressLandedCostRequest_customerDetails.md) |  | [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) | Please enter all the DHL Express accounts and types to be used for this shipment | [default to null]
**ProductCode** | **string** | Please enter DHL Express Global Product code | [optional] [default to null]
**LocalProductCode** | **string** | Please enter DHL Express Local Product code | [optional] [default to null]
**UnitOfMeasurement** | **string** | Please enter Unit of measurement - metric,imperial | [default to null]
**CurrencyCode** | **string** | Currency code for the item price (the product being sold) and freight charge. The Landed Cost calculation result will be returned in this defined currency | [default to null]
**IsCustomsDeclarable** | **bool** | Set this to true is shipment contains declarable content | [default to null]
**IsDTPRequested** | **bool** | Set this to true if you want DHL EXpress product Duties and Taxes Paid outside shipment destination | [optional] [default to null]
**IsInsuranceRequested** | **bool** | Set this true if you ask for DHL Express insurance service | [optional] [default to null]
**GetCostBreakdown** | **bool** | Allowed values &#x27;true&#x27; - item cost breakdown will be returned, &#x27;false&#x27; - item cost breakdown will not be returned | [default to null]
**Charges** | [**[]SupermodelIoLogisticsExpressLandedCostRequestCharges**](supermodelIoLogisticsExpressLandedCostRequest_charges.md) | Please provide any additional charges you would like to include in total cost calculation | [optional] [default to null]
**ShipmentPurpose** | **string** | Possible values:&lt;BR&gt;      commercial: B2B&lt;BR&gt;      personal: B2C&lt;BR&gt;      commercia&#x27;: B2B&lt;BR&gt;      personal: B2C | [optional] [default to null]
**TransportationMode** | **string** |  | [optional] [default to null]
**MerchantSelectedCarrierName** | **string** | Carrier being used to ship with. Allowed values are:&lt;BR&gt;      &#x27;DHL&#x27;,&#x27;UPS&#x27;,&#x27;FEDEX&#x27;,&#x27;TNT&#x27;,&#x27;POST&#x27;,&lt;BR&gt;      &#x27;OTHERS&#x27; | [optional] [default to null]
**Packages** | [**[]SupermodelIoLogisticsExpressPackageRr**](supermodelIoLogisticsExpressPackageRR.md) | Here you can define properties per package | [default to null]
**Items** | [**[]SupermodelIoLogisticsExpressLandedCostRequestItems**](supermodelIoLogisticsExpressLandedCostRequest_items.md) |  | [default to null]
**GetTariffFormula** | **bool** | Allowed values &#x27;true&#x27; - tariff formula on item and shipment level will be returned, &#x27;false&#x27; - tariff formula on item and shipment level will not be returned | [optional] [default to null]
**GetQuotationID** | **bool** | Allowed values &#x27;true&#x27; - quotation ID on shipment level will be returned, &#x27;false&#x27; - quotation ID on shipment level will not be returned | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

