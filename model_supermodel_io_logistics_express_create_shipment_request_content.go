/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Here you can define all the properties related to the content of the prospected shipment
type SupermodelIoLogisticsExpressCreateShipmentRequestContent struct {
	// Here you can define properties per package
	Packages []SupermodelIoLogisticsExpressPackage `json:"packages"`
	// For customs purposes please advise if your shipment is dutiable (true) or non dutiable (false).Note:If the shipment is dutiable, exportDeclaration element must be provided.
	IsCustomsDeclarable bool `json:"isCustomsDeclarable"`
	// For customs purposes please advise on declared value of the shipment
	DeclaredValue float64 `json:"declaredValue,omitempty"`
	// For customs purposes please advise on declared value currency code of the shipment
	DeclaredValueCurrency string                                                                     `json:"declaredValueCurrency,omitempty"`
	ExportDeclaration     *SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration `json:"exportDeclaration,omitempty"`
	// Please enter description of your shipment
	Description string `json:"description"`
	// This is used for the US AES4, FTR and ITN numbers to be printed on the Transport Label
	USFilingTypeValue string `json:"USFilingTypeValue,omitempty"`
	// The Incoterms rules are a globally-recognized set of standards, used worldwide in international and domestic contracts for the delivery of goods, illustrating responsibilities between buyer and seller for costs and risk, as well as cargo insurance.<BR>          EXW ExWorks<BR>          FCA Free Carrier<BR>          CPT Carriage Paid To<BR>          CIP Carriage and Insurance Paid To<BR>          DPU Delivered at Place Unloaded<BR>          DAP Delivered at Place<BR>          DDP Delivered Duty Paid<BR>          FAS Free Alongside Ship<BR>          FOB Free on Board<BR>          CFR Cost and Freight<BR>          CIF Cost, Insurance and Freight<BR>          DAF Delivered at Frontier<BR>          DAT Delivered at Terminal<BR>          DDU Delivered Duty Unpaid<BR>          DEQ Delivery ex Quay<BR>          DES Delivered ex Ship
	Incoterm string `json:"incoterm"`
	// Please enter Unit of measurement - metric,imperial
	UnitOfMeasurement string `json:"unitOfMeasurement"`
}
