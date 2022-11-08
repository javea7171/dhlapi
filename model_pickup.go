/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type Pickup struct {
	// Please advise if a pickup is needed for this shipment
	IsRequested bool `json:"isRequested"`
	// The latest time the location premises is available to dispatch the DHL Express shipment. (HH:MM)
	CloseTime string `json:"closeTime,omitempty"`
	// Provides information on where the package should be picked up by DHL courier.
	Location string `json:"location,omitempty"`
	// Details special pickup instructions you may wish to send to the DHL Courier.
	SpecialInstructions    []PickupSpecialInstructions   `json:"specialInstructions,omitempty"`
	PickupDetails          *PickupPickupDetails          `json:"pickupDetails,omitempty"`
	PickupRequestorDetails *PickupPickupRequestorDetails `json:"pickupRequestorDetails,omitempty"`
}