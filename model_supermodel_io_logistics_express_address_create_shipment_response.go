/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

type SupermodelIoLogisticsExpressAddressCreateShipmentResponse struct {
	// Postal code
	PostalCode string `json:"postalCode"`
	// City name
	CityName string `json:"cityName"`
	// Country code
	CountryCode string `json:"countryCode"`
	// Province or state code
	ProvinceCode string `json:"provinceCode,omitempty"`
	// Address line 1
	AddressLine1 string `json:"addressLine1"`
	// Address line 2
	AddressLine2 string `json:"addressLine2,omitempty"`
	// Address line 3
	AddressLine3 string `json:"addressLine3,omitempty"`
	// Suburb or county name
	CityDistrictName string `json:"cityDistrictName,omitempty"`
	// Please enter your state or province name
	ProvinceName string `json:"provinceName,omitempty"`
	// Please enter your country name
	CountryName string `json:"countryName,omitempty"`
}
