/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Definition of /rates request message
type SupermodelIoLogisticsExpressRateRequest struct {
	CustomerDetails *SupermodelIoLogisticsExpressRateRequestCustomerDetails `json:"customerDetails"`
	// Please enter all the DHL Express accounts and types to be used for this shipment
	Accounts []SupermodelIoLogisticsExpressAccount `json:"accounts,omitempty"`
	// Please enter DHL Express Global Product code
	ProductCode string `json:"productCode,omitempty"`
	// Please enter DHL Express Local Product code
	LocalProductCode string `json:"localProductCode,omitempty"`
	// Please use if you wish to filter the response by value added services
	ValueAddedServices []SupermodelIoLogisticsExpressValueAddedServicesRates `json:"valueAddedServices,omitempty"`
	// Please use if you wish to filter the response by product(s) and/or value added services
	ProductsAndServices []SupermodelIoLogisticsExpressRateRequestProductsAndServices `json:"productsAndServices,omitempty"`
	// payerCountryCode is to be provided if your profile has been enabled to view rates without an account number (this will provide DHL Express published rates for the payer country)
	PayerCountryCode string `json:"payerCountryCode,omitempty"`
	// Identifies the date and time the package is tendered. Both the date and time portions of the string are expected to be used. The date should not be a past date or a date more than 10 days in the future. The time is the local time of the shipment based on the shipper's time zone. The date component must be in the format: YYYY-MM-DD; the time component must be in the format: HH:MM:SS using a 24 hour clock. The date and time parts are separated by the letter T (e.g. 2006-06-26T17:00:00 GMT+01:00).
	PlannedShippingDateAndTime string `json:"plannedShippingDateAndTime"`
	// Please enter Unit of measurement - metric,imperial
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	// For customs purposes please advise if your shipment is dutiable (true) or non dutiable (false)
	IsCustomsDeclarable bool `json:"isCustomsDeclarable"`
	// Please provide monetary amount related to your shipment, for example shipment declared value
	MonetaryAmount []SupermodelIoLogisticsExpressRateRequestMonetaryAmount `json:"monetaryAmount,omitempty"`
	// Legacy field and replaced by newer field getAdditionalInformation. Please set this to true to receive all value added services for each product available
	RequestAllValueAddedServices bool                    `json:"requestAllValueAddedServices,omitempty"`
	EstimatedDeliveryDate        *EstimatedDeliveryDate1 `json:"estimatedDeliveryDate,omitempty"`
	// Provides additional information in the response like all value added services, and rule groups
	GetAdditionalInformation []SupermodelIoLogisticsExpressRateRequestGetAdditionalInformation `json:"getAdditionalInformation,omitempty"`
	// Please set this to true to filter out all products which needs DHL Express special customer agreement
	ReturnStandardProductsOnly bool `json:"returnStandardProductsOnly,omitempty"`
	// Please set this to true in case you want to receive products which are not available on planned shipping date but next available day
	NextBusinessDay bool `json:"nextBusinessDay,omitempty"`
	// Please select which type of priducts you are interested in
	ProductTypeCode string `json:"productTypeCode,omitempty"`
	// Here you can define properties per package
	Packages []SupermodelIoLogisticsExpressPackageRr `json:"packages"`
}
