/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Here you can find all details related to export declaration
type SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration struct {
	// Please enter details for each export line item
	LineItems []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLineItems `json:"lineItems"`
	Invoice   *SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoice    `json:"invoice,omitempty"`
	// Please enter up to three remarks. <BR>              If using Customs Invoice template COMMERCIAL_INVOICE_04, the invoice can only print the first remarks field. The recommended max length is 20 characters. <BR>              If using Customs Invoice template COMMERCIAL_INVOICE_L_10 or COMMERCIAL_INVOICE_P_10, the invoice can print all three remraks fields.  The recommended max length is 45 characters.
	Remarks []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationRemarks `json:"remarks,omitempty"`
	// Please enter additional charge to appear on the invoice<BR>              admin, Administration Charge<BR>              delivery, Delivery Charge<BR>              documentation, Documentation Charge<BR>              expedite, Expedite Charge<BR>              export, Export Charge<BR>              freight, Freight Charge<BR>              fuel_surcharge, Fuel Surcharge<BR>              logistic, Logistic Charge<BR>              other, Other Charge<BR>              packaging, Packaging Charge<BR>              pickup, Pickup Charge<BR>              handling, Handling Charge<BR>              vat, VAT Charge<BR>              insurance, Insurance Cost<BR>              reverse_charge, Reverse Charge
	AdditionalCharges []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationAdditionalCharges `json:"additionalCharges,omitempty"`
	// Please provide destination port details
	DestinationPortName string `json:"destinationPortName,omitempty"`
	// Name of port of departure, shipment or destination as required under the applicable delivery term.
	PlaceOfIncoterm string `json:"placeOfIncoterm,omitempty"`
	// Please provide Payer VAT number
	PayerVATNumber string `json:"payerVATNumber,omitempty"`
	// Please enter recipient reference
	RecipientReference string                                                                             `json:"recipientReference,omitempty"`
	Exporter           *SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationExporter `json:"exporter,omitempty"`
	// Please enter package marks
	PackageMarks string `json:"packageMarks,omitempty"`
	// Please provide up to three dcelaration notes
	DeclarationNotes []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationDeclarationNotes `json:"declarationNotes,omitempty"`
	// Please enter export reference
	ExportReference string `json:"exportReference,omitempty"`
	// Please enter export reason
	ExportReason string `json:"exportReason,omitempty"`
	// Please provide the reason for export
	ExportReasonType string `json:"exportReasonType,omitempty"`
	// Please provide details about export and import licenses
	Licenses []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLicenses `json:"licenses,omitempty"`
	// Please provide the shipment was sent for Personal (Gift) or Commercial (Sale) reasons
	ShipmentType string `json:"shipmentType,omitempty"`
	// Please provide the Customs Documents at invoice level
	CustomsDocuments []SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments1 `json:"customsDocuments,omitempty"`
}
