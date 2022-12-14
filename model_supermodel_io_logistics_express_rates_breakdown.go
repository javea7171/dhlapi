/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressRatesBreakdown struct {
	// When landed-cost is requested then following items name (Charge Types) might be returned: <BR>                        Charge Type : Description <BR>                        STDIS : Quoted shipment total discount <BR>                        SCUSV : Shipment Customs value <BR>                        SINSV : Insured value <BR>                        SPRQD : Shipment product quote discount<BR>                        SPRQN : The price quoted to the Customer by DHL at the time of the booking. This quote covers the weight price including discounts and without taxes. <BR>                        STSCH : The total of service charges quoted to customer for DHL Express value added services, the amount is after discounts and doesn't include tax amounts. <BR>                        MACHG : The total of service charges as provided by Merchant for the purpose of landed cost calculation. <BR>                        MFCHG : The freight charge as provided by Merchant for the purpose of landed cost calculation.
	Name string `json:"name,omitempty"`
	// Special service or extra charge code. This is the code you would have to use in the /shipment service if you wish to add an optional Service such as Saturday delivery
	ServiceCode string `json:"serviceCode,omitempty"`
	// Local service code
	LocalServiceCode string `json:"localServiceCode,omitempty"`
	// Price breakdown type code
	TypeCode string `json:"typeCode,omitempty"`
	// Special service charge code type for service.
	ServiceTypeCode string `json:"serviceTypeCode,omitempty"`
	// Price breakdown value
	Price float64 `json:"price,omitempty"`
	// This the currency of the rated shipment for the prices listed.
	PriceCurrency string `json:"priceCurrency,omitempty"`
	// Customer agreement indicator for product and services, if service is offered with prior customer agreement
	IsCustomerAgreement bool `json:"isCustomerAgreement,omitempty"`
	// Indicator if the special service is marketed service
	IsMarketedService bool `json:"isMarketedService,omitempty"`
	// Indicator if there is any discount allowed
	IsBillingServiceIndicator bool                                               `json:"isBillingServiceIndicator,omitempty"`
	PriceBreakdown            []SupermodelIoLogisticsExpressRatesPriceBreakdown1 `json:"priceBreakdown,omitempty"`
	// Tariff Rate Formula on Shipment Level
	TariffRateFormula string `json:"tariffRateFormula,omitempty"`
}
