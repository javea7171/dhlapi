/*
 * DHL Express APIs (MyDHL API)
 *
 * Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments)
 *
 * API version: 2.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dhlapi

// Please enter information about quantity for this line item
type SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationQuantity struct {
	// Please enter number of pieces in the line item
	Value int32 `json:"value"`
	// Please provide correct unit of measurement<BR>                        <BR>                        Possible values;<BR>                        BOX Boxes<BR>                        2GM Centigram<BR>                        2M Centimeters<BR>                        2M3 Cubic Centimeters<BR>                        3M3 Cubic Feet<BR>                        M3 Cubic Meters<BR>                        DPR Dozen Pairs<BR>                        DOZ Dozen<BR>                        2NO Each<BR>                        PCS Pieces<BR>                        GM Grams<BR>                        GRS Gross<BR>                        KG Kilograms<BR>                        L Liters<BR>                        M Meters<BR>                        3GM Milligrams<BR>                        3L Milliliters<BR>                        X No Unit Required<BR>                        NO Number<BR>                        2KG Ounces<BR>                        PRS Pairs<BR>                        2L Gallons<BR>                        3KG Pounds<BR>                        CM2 Square Centimeters<BR>                        2M2 Square Feet<BR>                        3M2 Square Inches<BR>                        M2 Square Meters<BR>                        4M2 Square Yards<BR>                        3M Yards<BR>                        CM Centimeters<BR>                        CONE Cone<BR>                        CT Carat<BR>                        EA Each<BR>                        LBS Pounds<BR>                        RILL Rill<BR>                        ROLL Roll<BR>                        SET Set<BR>                        TU Time Unit<BR>                        YDS Yard
	UnitOfMeasurement string `json:"unitOfMeasurement"`
}
