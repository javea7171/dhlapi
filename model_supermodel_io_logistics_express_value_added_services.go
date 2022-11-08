/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressValueAddedServices struct {
	// Please enter DHL Express value added service code. For detailed list of all available service codes for your prospect shipment please invoke GET /products or GET /rates
	ServiceCode string `json:"serviceCode"`
	// Please enter monetary value of service (e.g. Insured Value)
	Value float64 `json:"value,omitempty"`
	// Please enter currency code (e.g. Insured Value currency code)
	Currency string `json:"currency,omitempty"`
	// Payment method code (e.g. Cash)
	Method string `json:"method,omitempty"`
	// The DangerousGoods section indicates if there is dangerous good content within the shipment
	DangerousGoods []SupermodelIoLogisticsExpressValueAddedServicesDangerousGoods `json:"dangerousGoods,omitempty"`
}
