/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressCreateShipmentRequestShipmentNotification struct {
	// Please enter channel type to send the notification by. At this moment only email is supported
	TypeCode string `json:"typeCode"`
	// Please enter notification receiver email address
	ReceiverId string `json:"receiverId"`
	// Please enter three letter lanuage code in which you wish to receive the notification in
	LanguageCode string `json:"languageCode,omitempty"`
	// Please enter two letter language country code
	LanguageCountryCode string `json:"languageCountryCode,omitempty"`
	// Please enter your message which will be added to the DHL Express notification email
	BespokeMessage string `json:"bespokeMessage,omitempty"`
}
