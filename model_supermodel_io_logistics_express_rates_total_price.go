/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressRatesTotalPrice struct {
	// Possible Values :<BR>                  'BILLC', billing currency<BR>                  'PULCL', country public rates currency<BR>                  'BASEC', base currency
	CurrencyType string `json:"currencyType,omitempty"`
	// This the currency of the rated shipment for the prices listed.
	PriceCurrency string `json:"priceCurrency,omitempty"`
	// This is the total prize of the rated shipment for the product listed.
	Price float64 `json:"price"`
}