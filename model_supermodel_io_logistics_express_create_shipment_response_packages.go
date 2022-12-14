/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressCreateShipmentResponsePackages struct {
	// Piece serial number
	ReferenceNumber float64 `json:"referenceNumber,omitempty"`
	// Here is provided each piece its Identification number
	TrackingNumber string `json:"trackingNumber"`
	// You can use ths URL to track your shipment by Piece Identification Number
	TrackingUrl string `json:"trackingUrl,omitempty"`
	// Here is provided each piece volumetric/ dimensional weight
	VolumetricWeight float64 `json:"volumetricWeight,omitempty"`
	// Here you can find all documents created for the piece's QRcode
	Documents []SupermodelIoLogisticsExpressCreateShipmentResponseDocuments `json:"documents,omitempty"`
}
