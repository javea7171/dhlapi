/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type RatingApiService service

/*
RatingApiService Landed Cost
The Landed Cost section allows further information around products being sold to be provided. In return the duty, tax and shipping charges are calculated in real time and provides transparency about any extra costs the buyer may have to pay before they reach them.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Details about the requested shipment
 * @param optional nil or *RatingApiExpApiLandedCostOpts - Optional Parameters:
     * @param "MessageReference" (optional.String) -  Please provide message reference
     * @param "MessageReferenceDate" (optional.String) -  Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2
     * @param "PluginName" (optional.String) -  Please provide name of the plugin (applicable to 3PV only)
     * @param "PluginVersion" (optional.String) -  Please provide version of the plugin (applicable to 3PV only)
     * @param "ShippingSystemPlatformName" (optional.String) -  Please provide name of the shipping platform(applicable to 3PV only)
     * @param "ShippingSystemPlatformVersion" (optional.String) -  Please provide version of the shipping platform (applicable to 3PV only)
     * @param "WebstorePlatformName" (optional.String) -  Please provide name of the webstore platform (applicable to 3PV only)
     * @param "WebstorePlatformVersion" (optional.String) -  Please provide version of the webstore platform (applicable to 3PV only)
@return SupermodelIoLogisticsExpressRates
*/

type RatingApiExpApiLandedCostOpts struct {
	MessageReference              optional.String
	MessageReferenceDate          optional.String
	PluginName                    optional.String
	PluginVersion                 optional.String
	ShippingSystemPlatformName    optional.String
	ShippingSystemPlatformVersion optional.String
	WebstorePlatformName          optional.String
	WebstorePlatformVersion       optional.String
}

func (a *RatingApiService) ExpApiLandedCost(ctx context.Context, body SupermodelIoLogisticsExpressLandedCostRequest, localVarOptionals *RatingApiExpApiLandedCostOpts) (SupermodelIoLogisticsExpressRates, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SupermodelIoLogisticsExpressRates
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/landed-cost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.MessageReference.IsSet() {
		localVarHeaderParams["Message-Reference"] = parameterToString(localVarOptionals.MessageReference.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.MessageReferenceDate.IsSet() {
		localVarHeaderParams["Message-Reference-Date"] = parameterToString(localVarOptionals.MessageReferenceDate.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginName.IsSet() {
		localVarHeaderParams["Plugin-Name"] = parameterToString(localVarOptionals.PluginName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginVersion.IsSet() {
		localVarHeaderParams["Plugin-Version"] = parameterToString(localVarOptionals.PluginVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformName.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Name"] = parameterToString(localVarOptionals.ShippingSystemPlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformVersion.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Version"] = parameterToString(localVarOptionals.ShippingSystemPlatformVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformName.IsSet() {
		localVarHeaderParams["Webstore-Platform-Name"] = parameterToString(localVarOptionals.WebstorePlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformVersion.IsSet() {
		localVarHeaderParams["Webstore-Platform-Version"] = parameterToString(localVarOptionals.WebstorePlatformVersion.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SupermodelIoLogisticsExpressRates
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RatingApiService Retrieve Rates for a one piece Shipment
The Rate request will return DHL&#x27;s product capabilities and prices (where applicable) based on the input data. Using the shipper and receiver address as well as the dimension and weights of the pieces belonging to a shipment, this operation returns the available products including the shipping price (where applicable)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountNumber DHL Express customer account number
 * @param originCountryCode A short text string code (see values defined in ISO 3166) specifying the shipment origin country. https://gs1.org/voc/Country, Alpha-2 Code
 * @param originCityName Text specifying the city name
 * @param destinationCountryCode A short text string code (see values defined in ISO 3166) specifying the shipment destination country. https://gs1.org/voc/Country, Alpha-2 Code
 * @param destinationCityName Text specifying the city name
 * @param weight Gross weight of the shipment including packaging.
 * @param length Total length of the shipment including packaging.
 * @param width Total width of the shipment including packaging.
 * @param height Total height of the shipment including packaging.
 * @param plannedShippingDate Timestamp represents the date you plan to ship your prospected shipment
 * @param isCustomsDeclarable
 * @param unitOfMeasurement The UnitOfMeasurement node conveys the unit of measurements used in the operation. This single value corresponds to the units of weight and measurement which are used throughout the message processing.
 * @param optional nil or *RatingApiExpApiRatesOpts - Optional Parameters:
     * @param "OriginPostalCode" (optional.String) -  Text specifying the postal code for an address. https://gs1.org/voc/postalCode
     * @param "DestinationPostalCode" (optional.String) -  Text specifying the postal code for an address. https://gs1.org/voc/postalCode
     * @param "NextBusinessDay" (optional.Bool) -  When set to true and there are no products available for given plannedShippingDate then products available for the next possible pickup date are returned
     * @param "MessageReference" (optional.String) -  Please provide message reference
     * @param "MessageReferenceDate" (optional.String) -  Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2
     * @param "PluginName" (optional.String) -  Please provide name of the plugin (applicable to 3PV only)
     * @param "PluginVersion" (optional.String) -  Please provide version of the plugin (applicable to 3PV only)
     * @param "ShippingSystemPlatformName" (optional.String) -  Please provide name of the shipping platform(applicable to 3PV only)
     * @param "ShippingSystemPlatformVersion" (optional.String) -  Please provide version of the shipping platform (applicable to 3PV only)
     * @param "WebstorePlatformName" (optional.String) -  Please provide name of the webstore platform (applicable to 3PV only)
     * @param "WebstorePlatformVersion" (optional.String) -  Please provide version of the webstore platform (applicable to 3PV only)
     * @param "StrictValidation" (optional.Bool) -  If set to true, indicate strict DCT validation of address details, and validation of product and service(s) combination provided in request.
     * @param "GetAllValueAddedServices" (optional.Bool) -  Option to return list of all value added services and its rule groups if applicable
     * @param "RequestEstimatedDeliveryDate" (optional.Bool) -  Option to return Estimated Delivery Date in response
     * @param "EstimatedDeliveryDateType" (optional.String) -  Estimated Delivery Date Type. QDDF: is the fastest &#x27;docs&#x27; transit time as quoted to the customer at booking or shipment creation. No custom clearance is considered. QDDC: constitutes DHL&#x27;s service commitment as quoted at booking or shipment creation. QDDc builds in clearance time, and potentially other special perational non-transport component(s), when relevant.
@return SupermodelIoLogisticsExpressRates
*/

type RatingApiExpApiRatesOpts struct {
	OriginPostalCode              optional.String
	DestinationPostalCode         optional.String
	NextBusinessDay               optional.Bool
	MessageReference              optional.String
	MessageReferenceDate          optional.String
	PluginName                    optional.String
	PluginVersion                 optional.String
	ShippingSystemPlatformName    optional.String
	ShippingSystemPlatformVersion optional.String
	WebstorePlatformName          optional.String
	WebstorePlatformVersion       optional.String
	StrictValidation              optional.Bool
	GetAllValueAddedServices      optional.Bool
	RequestEstimatedDeliveryDate  optional.Bool
	EstimatedDeliveryDateType     optional.String
}

func (a *RatingApiService) ExpApiRates(ctx context.Context, accountNumber string, originCountryCode string, originCityName string, destinationCountryCode string, destinationCityName string, weight float64, length float64, width float64, height float64, plannedShippingDate string, isCustomsDeclarable bool, unitOfMeasurement string, localVarOptionals *RatingApiExpApiRatesOpts) (SupermodelIoLogisticsExpressRates, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SupermodelIoLogisticsExpressRates
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if weight < 0 {
		return localVarReturnValue, nil, reportError("weight must be greater than 0")
	}
	if weight > 999999999999 {
		return localVarReturnValue, nil, reportError("weight must be less than 999999999999")
	}
	if length < 1 {
		return localVarReturnValue, nil, reportError("length must be greater than 1")
	}
	if length > 9999999999 {
		return localVarReturnValue, nil, reportError("length must be less than 9999999999")
	}
	if width < 1 {
		return localVarReturnValue, nil, reportError("width must be greater than 1")
	}
	if width > 9999999999 {
		return localVarReturnValue, nil, reportError("width must be less than 9999999999")
	}
	if height < 1 {
		return localVarReturnValue, nil, reportError("height must be greater than 1")
	}
	if height > 9999999999 {
		return localVarReturnValue, nil, reportError("height must be less than 9999999999")
	}

	localVarQueryParams.Add("accountNumber", parameterToString(accountNumber, ""))
	localVarQueryParams.Add("originCountryCode", parameterToString(originCountryCode, ""))
	if localVarOptionals != nil && localVarOptionals.OriginPostalCode.IsSet() {
		localVarQueryParams.Add("originPostalCode", parameterToString(localVarOptionals.OriginPostalCode.Value(), ""))
	}
	localVarQueryParams.Add("originCityName", parameterToString(originCityName, ""))
	localVarQueryParams.Add("destinationCountryCode", parameterToString(destinationCountryCode, ""))
	if localVarOptionals != nil && localVarOptionals.DestinationPostalCode.IsSet() {
		localVarQueryParams.Add("destinationPostalCode", parameterToString(localVarOptionals.DestinationPostalCode.Value(), ""))
	}
	localVarQueryParams.Add("destinationCityName", parameterToString(destinationCityName, ""))
	localVarQueryParams.Add("weight", parameterToString(weight, ""))
	localVarQueryParams.Add("length", parameterToString(length, ""))
	localVarQueryParams.Add("width", parameterToString(width, ""))
	localVarQueryParams.Add("height", parameterToString(height, ""))
	localVarQueryParams.Add("plannedShippingDate", parameterToString(plannedShippingDate, ""))
	localVarQueryParams.Add("isCustomsDeclarable", parameterToString(isCustomsDeclarable, ""))
	localVarQueryParams.Add("unitOfMeasurement", parameterToString(unitOfMeasurement, ""))
	if localVarOptionals != nil && localVarOptionals.NextBusinessDay.IsSet() {
		localVarQueryParams.Add("nextBusinessDay", parameterToString(localVarOptionals.NextBusinessDay.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StrictValidation.IsSet() {
		localVarQueryParams.Add("strictValidation", parameterToString(localVarOptionals.StrictValidation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetAllValueAddedServices.IsSet() {
		localVarQueryParams.Add("getAllValueAddedServices", parameterToString(localVarOptionals.GetAllValueAddedServices.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RequestEstimatedDeliveryDate.IsSet() {
		localVarQueryParams.Add("requestEstimatedDeliveryDate", parameterToString(localVarOptionals.RequestEstimatedDeliveryDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EstimatedDeliveryDateType.IsSet() {
		localVarQueryParams.Add("estimatedDeliveryDateType", parameterToString(localVarOptionals.EstimatedDeliveryDateType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.MessageReference.IsSet() {
		localVarHeaderParams["Message-Reference"] = parameterToString(localVarOptionals.MessageReference.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.MessageReferenceDate.IsSet() {
		localVarHeaderParams["Message-Reference-Date"] = parameterToString(localVarOptionals.MessageReferenceDate.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginName.IsSet() {
		localVarHeaderParams["Plugin-Name"] = parameterToString(localVarOptionals.PluginName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginVersion.IsSet() {
		localVarHeaderParams["Plugin-Version"] = parameterToString(localVarOptionals.PluginVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformName.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Name"] = parameterToString(localVarOptionals.ShippingSystemPlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformVersion.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Version"] = parameterToString(localVarOptionals.ShippingSystemPlatformVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformName.IsSet() {
		localVarHeaderParams["Webstore-Platform-Name"] = parameterToString(localVarOptionals.WebstorePlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformVersion.IsSet() {
		localVarHeaderParams["Webstore-Platform-Version"] = parameterToString(localVarOptionals.WebstorePlatformVersion.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SupermodelIoLogisticsExpressRates
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RatingApiService Retrieve Rates for Multi-piece Shipments
The Rate request will return DHL&#x27;s product capabilities and prices (where applicable) based on the input data. Using the shipper and receiver address as well as the dimension and weights of the pieces belonging to a shipment, this operation returns the available products including the shipping price (where applicable)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Details about the requested shipment
 * @param optional nil or *RatingApiExpApiRatesManyOpts - Optional Parameters:
     * @param "MessageReference" (optional.String) -  Please provide message reference
     * @param "MessageReferenceDate" (optional.String) -  Optional reference date in the  HTTP-date format https://tools.ietf.org/html/rfc7231#section-7.1.1.2
     * @param "PluginName" (optional.String) -  Please provide name of the plugin (applicable to 3PV only)
     * @param "PluginVersion" (optional.String) -  Please provide version of the plugin (applicable to 3PV only)
     * @param "ShippingSystemPlatformName" (optional.String) -  Please provide name of the shipping platform(applicable to 3PV only)
     * @param "ShippingSystemPlatformVersion" (optional.String) -  Please provide version of the shipping platform (applicable to 3PV only)
     * @param "WebstorePlatformName" (optional.String) -  Please provide name of the webstore platform (applicable to 3PV only)
     * @param "WebstorePlatformVersion" (optional.String) -  Please provide version of the webstore platform (applicable to 3PV only)
     * @param "StrictValidation" (optional.Bool) -  If set to true, indicate strict DCT validation of address details, and validation of product and service(s) combination provided in request.
@return SupermodelIoLogisticsExpressRates
*/

type RatingApiExpApiRatesManyOpts struct {
	MessageReference              optional.String
	MessageReferenceDate          optional.String
	PluginName                    optional.String
	PluginVersion                 optional.String
	ShippingSystemPlatformName    optional.String
	ShippingSystemPlatformVersion optional.String
	WebstorePlatformName          optional.String
	WebstorePlatformVersion       optional.String
	StrictValidation              optional.Bool
}

func (a *RatingApiService) ExpApiRatesMany(ctx context.Context, body SupermodelIoLogisticsExpressRateRequest, localVarOptionals *RatingApiExpApiRatesManyOpts) (SupermodelIoLogisticsExpressRates, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SupermodelIoLogisticsExpressRates
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.StrictValidation.IsSet() {
		localVarQueryParams.Add("strictValidation", parameterToString(localVarOptionals.StrictValidation.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.MessageReference.IsSet() {
		localVarHeaderParams["Message-Reference"] = parameterToString(localVarOptionals.MessageReference.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.MessageReferenceDate.IsSet() {
		localVarHeaderParams["Message-Reference-Date"] = parameterToString(localVarOptionals.MessageReferenceDate.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginName.IsSet() {
		localVarHeaderParams["Plugin-Name"] = parameterToString(localVarOptionals.PluginName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.PluginVersion.IsSet() {
		localVarHeaderParams["Plugin-Version"] = parameterToString(localVarOptionals.PluginVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformName.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Name"] = parameterToString(localVarOptionals.ShippingSystemPlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ShippingSystemPlatformVersion.IsSet() {
		localVarHeaderParams["Shipping-System-Platform-Version"] = parameterToString(localVarOptionals.ShippingSystemPlatformVersion.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformName.IsSet() {
		localVarHeaderParams["Webstore-Platform-Name"] = parameterToString(localVarOptionals.WebstorePlatformName.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.WebstorePlatformVersion.IsSet() {
		localVarHeaderParams["Webstore-Platform-Version"] = parameterToString(localVarOptionals.WebstorePlatformVersion.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SupermodelIoLogisticsExpressRates
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v SupermodelIoLogisticsExpressErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}