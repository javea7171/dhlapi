/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Landed cost request model description
type SupermodelIoLogisticsExpressLandedCostRequest struct {
	CustomerDetails *SupermodelIoLogisticsExpressLandedCostRequestCustomerDetails `json:"customerDetails"`
	// Please enter all the DHL Express accounts and types to be used for this shipment
	Accounts []SupermodelIoLogisticsExpressAccount `json:"accounts"`
	// Please enter DHL Express Global Product code
	ProductCode string `json:"productCode,omitempty"`
	// Please enter DHL Express Local Product code
	LocalProductCode string `json:"localProductCode,omitempty"`
	// Please enter Unit of measurement - metric,imperial
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	// Currency code for the item price (the product being sold) and freight charge. The Landed Cost calculation result will be returned in this defined currency
	CurrencyCode string `json:"currencyCode"`
	// Set this to true is shipment contains declarable content
	IsCustomsDeclarable bool `json:"isCustomsDeclarable"`
	// Set this to true if you want DHL EXpress product Duties and Taxes Paid outside shipment destination
	IsDTPRequested bool `json:"isDTPRequested,omitempty"`
	// Set this true if you ask for DHL Express insurance service
	IsInsuranceRequested bool `json:"isInsuranceRequested,omitempty"`
	// Allowed values 'true' - item cost breakdown will be returned, 'false' - item cost breakdown will not be returned
	GetCostBreakdown bool `json:"getCostBreakdown"`
	// Please provide any additional charges you would like to include in total cost calculation
	Charges []SupermodelIoLogisticsExpressLandedCostRequestCharges `json:"charges,omitempty"`
	// Possible values:<BR>      commercial: B2B<BR>      personal: B2C<BR>      commercia': B2B<BR>      personal: B2C
	ShipmentPurpose    string `json:"shipmentPurpose,omitempty"`
	TransportationMode string `json:"transportationMode,omitempty"`
	// Carrier being used to ship with. Allowed values are:<BR>      'DHL','UPS','FEDEX','TNT','POST',<BR>      'OTHERS'
	MerchantSelectedCarrierName string `json:"merchantSelectedCarrierName,omitempty"`
	// Here you can define properties per package
	Packages []SupermodelIoLogisticsExpressPackageRr              `json:"packages"`
	Items    []SupermodelIoLogisticsExpressLandedCostRequestItems `json:"items"`
	// Allowed values 'true' - tariff formula on item and shipment level will be returned, 'false' - tariff formula on item and shipment level will not be returned
	GetTariffFormula bool `json:"getTariffFormula,omitempty"`
	// Allowed values 'true' - quotation ID on shipment level will be returned, 'false' - quotation ID on shipment level will not be returned
	GetQuotationID bool `json:"getQuotationID,omitempty"`
}
