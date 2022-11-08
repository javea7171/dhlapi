/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressRatesPriceBreakdown1 struct {
	// If a breakdown is provided, details can either be; 'TAX',<BR>                              'DISCOUNT'
	PriceType string `json:"priceType,omitempty"`
	// Discount or tax type codes as provided by DHL Express. Example values:<BR>                              For discount;<BR>                              P: promotional<BR>                              S: special
	TypeCode string `json:"typeCode,omitempty"`
	// The actual amount of the discount/tax
	Price float64 `json:"price,omitempty"`
	// Percentage of the discount/tax
	Rate float64 `json:"rate,omitempty"`
	// The base amount of the service charge
	BasePrice float64 `json:"basePrice,omitempty"`
}
