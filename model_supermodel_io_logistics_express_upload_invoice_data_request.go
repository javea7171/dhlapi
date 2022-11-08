/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// MyDHL API REST UploadInvoiceData request JSON Schema
type SupermodelIoLogisticsExpressUploadInvoiceDataRequest struct {
	// The planned shipment date for the provided shipmentTrackingNumber.  The date must be in the format: YYYY-MM-DD
	PlannedShipDate string `json:"plannedShipDate,omitempty"`
	// Please enter all the DHL Express accounts and types to be used for this shipment.   Note: accounts/0/number with typeCode 'shipper' is mandatory if using POST method and no shipmentTrackingNumber is provided in request.
	Accounts              []SupermodelIoLogisticsExpressAccount                                      `json:"accounts,omitempty"`
	Content               *SupermodelIoLogisticsExpressUploadInvoiceDataRequestContent               `json:"content"`
	OutputImageProperties *SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImageProperties `json:"outputImageProperties,omitempty"`
	CustomerDetails       *SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetails       `json:"customerDetails,omitempty"`
}
