/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Comment describing your JSON Schema
type SupermodelIoLogisticsExpressPickupResponse struct {
	// List of Dispatch Confirmation Numbers which identifies the scheduled pickup
	DispatchConfirmationNumbers []string `json:"dispatchConfirmationNumbers,omitempty"`
	ReadyByTime                 string   `json:"readyByTime,omitempty"`
	NextPickupDate              string   `json:"nextPickupDate,omitempty"`
	Warnings                    []string `json:"warnings,omitempty"`
}
