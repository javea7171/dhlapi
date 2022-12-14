/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressDocumentImageResponseDocuments struct {
	// Shipment Tracking Number
	ShipmentTrackingNumber string `json:"shipmentTrackingNumber"`
	// Identifies type of the document like commercial invoice or waybill, or archived zip documents
	TypeCode string `json:"typeCode"`
	// Clearance code or document function whether for import, export or both.  Returned only for customs-entry
	Function string `json:"function,omitempty"`
	// Identifies image format the document is created in, like PDF, TIFF, or ZIP
	EncodingFormat string `json:"encodingFormat"`
	// Contains base64 encoded document image or archived zip
	Content string `json:"content"`
}
