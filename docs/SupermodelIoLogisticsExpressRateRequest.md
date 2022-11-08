# SupermodelIoLogisticsExpressRateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerDetails** | [***SupermodelIoLogisticsExpressRateRequestCustomerDetails**](supermodelIoLogisticsExpressRateRequest_customerDetails.md) |  | [default to null]
**Accounts** | [**[]SupermodelIoLogisticsExpressAccount**](supermodelIoLogisticsExpressAccount.md) | Please enter all the DHL Express accounts and types to be used for this shipment | [optional] [default to null]
**ProductCode** | **string** | Please enter DHL Express Global Product code | [optional] [default to null]
**LocalProductCode** | **string** | Please enter DHL Express Local Product code | [optional] [default to null]
**ValueAddedServices** | [**[]SupermodelIoLogisticsExpressValueAddedServicesRates**](supermodelIoLogisticsExpressValueAddedServicesRates.md) | Please use if you wish to filter the response by value added services | [optional] [default to null]
**ProductsAndServices** | [**[]SupermodelIoLogisticsExpressRateRequestProductsAndServices**](supermodelIoLogisticsExpressRateRequest_productsAndServices.md) | Please use if you wish to filter the response by product(s) and/or value added services | [optional] [default to null]
**PayerCountryCode** | **string** | payerCountryCode is to be provided if your profile has been enabled to view rates without an account number (this will provide DHL Express published rates for the payer country) | [optional] [default to null]
**PlannedShippingDateAndTime** | **string** | Identifies the date and time the package is tendered. Both the date and time portions of the string are expected to be used. The date should not be a past date or a date more than 10 days in the future. The time is the local time of the shipment based on the shipper&#x27;s time zone. The date component must be in the format: YYYY-MM-DD; the time component must be in the format: HH:MM:SS using a 24 hour clock. The date and time parts are separated by the letter T (e.g. 2006-06-26T17:00:00 GMT+01:00). | [default to null]
**UnitOfMeasurement** | **string** | Please enter Unit of measurement - metric,imperial | [default to null]
**IsCustomsDeclarable** | **bool** | For customs purposes please advise if your shipment is dutiable (true) or non dutiable (false) | [default to null]
**MonetaryAmount** | [**[]SupermodelIoLogisticsExpressRateRequestMonetaryAmount**](supermodelIoLogisticsExpressRateRequest_monetaryAmount.md) | Please provide monetary amount related to your shipment, for example shipment declared value | [optional] [default to null]
**RequestAllValueAddedServices** | **bool** | Legacy field and replaced by newer field getAdditionalInformation. Please set this to true to receive all value added services for each product available | [optional] [default to null]
**EstimatedDeliveryDate** | [***EstimatedDeliveryDate1**](estimated delivery date_1.md) |  | [optional] [default to null]
**GetAdditionalInformation** | [**[]SupermodelIoLogisticsExpressRateRequestGetAdditionalInformation**](supermodelIoLogisticsExpressRateRequest_getAdditionalInformation.md) | Provides additional information in the response like all value added services, and rule groups | [optional] [default to null]
**ReturnStandardProductsOnly** | **bool** | Please set this to true to filter out all products which needs DHL Express special customer agreement | [optional] [default to null]
**NextBusinessDay** | **bool** | Please set this to true in case you want to receive products which are not available on planned shipping date but next available day | [optional] [default to false]
**ProductTypeCode** | **string** | Please select which type of priducts you are interested in | [optional] [default to null]
**Packages** | [**[]SupermodelIoLogisticsExpressPackageRr**](supermodelIoLogisticsExpressPackageRR.md) | Here you can define properties per package | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

