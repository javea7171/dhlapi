/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Please enter address and contact details related to receiver
type SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsReceiverDetails struct {
	PostalAddress       *SupermodelIoLogisticsExpressAddressCreateShipmentRequest `json:"postalAddress"`
	ContactInformation  *SupermodelIoLogisticsExpressContact                      `json:"contactInformation"`
	RegistrationNumbers []SupermodelIoLogisticsExpressRegistrationNumbers         `json:"registrationNumbers,omitempty"`
	BankDetails         *[]SupermodelIoLogisticsExpressBankDetailsInner           `json:"bankDetails,omitempty"`
	// Please enter the business party type of the receiver
	TypeCode string `json:"typeCode,omitempty"`
}
