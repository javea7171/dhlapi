/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressRatesProducts struct {
	// Name of the DHL Express product
	ProductName string `json:"productName,omitempty"`
	// This is the global DHL Express product code for which the delivery is feasible respecting the input data from the request.
	ProductCode string `json:"productCode,omitempty"`
	// This is the local DHL Express product code for which the delivery is feasible respecting the input data from the request.
	LocalProductCode string `json:"localProductCode,omitempty"`
	// The country code for the local service used
	LocalProductCountryCode string `json:"localProductCountryCode,omitempty"`
	// The NetworkTypeCode element indicates the product belongs to the Day Definite (DD) or Time Definite (TD) network.<BR>            Possible Values;<BR>             DD: Day Definite product<BR>             TD: Time Definite product
	NetworkTypeCode string `json:"networkTypeCode,omitempty"`
	// Indicator that the product only can be offered to customers with prior agreement.
	IsCustomerAgreement    bool                                                      `json:"isCustomerAgreement,omitempty"`
	Weight                 *Weight1                                                  `json:"weight"`
	TotalPrice             []SupermodelIoLogisticsExpressRatesTotalPrice             `json:"totalPrice"`
	TotalPriceBreakdown    []SupermodelIoLogisticsExpressRatesTotalPriceBreakdown    `json:"totalPriceBreakdown,omitempty"`
	DetailedPriceBreakdown []SupermodelIoLogisticsExpressRatesDetailedPriceBreakdown `json:"detailedPriceBreakdown,omitempty"`
	// Group of serviceCodes that are mutually exclusive.  Only one serviceCode among the list must be applied for a shipment
	ServiceCodeMutuallyExclusiveGroups []SupermodelIoLogisticsExpressProductsServiceCodeMutuallyExclusiveGroups `json:"serviceCodeMutuallyExclusiveGroups,omitempty"`
	// Dependency rule groups for a particular serviceCode.
	ServiceCodeDependencyRuleGroups []SupermodelIoLogisticsExpressProductsServiceCodeDependencyRuleGroups `json:"serviceCodeDependencyRuleGroups,omitempty"`
	PickupCapabilities              *SupermodelIoLogisticsExpressRatesPickupCapabilities                  `json:"pickupCapabilities,omitempty"`
	DeliveryCapabilities            *SupermodelIoLogisticsExpressRatesDeliveryCapabilities                `json:"deliveryCapabilities,omitempty"`
	Items                           []SupermodelIoLogisticsExpressRatesItems                              `json:"items,omitempty"`
	// The date when the rates for DHL products and services is provided
	PricingDate string `json:"pricingDate,omitempty"`
}
