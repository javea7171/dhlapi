# Go API client for swagger

Welcome to the official DHL Express APIs (MyDHL API) below are the published API Documentation to fulfill your shipping needs with DHL Express.       Please follow the process described [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--get-access) to request access to the DHL Express - MyDHL API services    In case you already have DHL Express - MyDHL API Service credentials please ensure to use the endpoints/environments listed  [here](https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/user-guide--environments) 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.7.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://developer.dhl.com/help-center](https://developer.dhl.com/help-center)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/javea7171/dhlapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api-mock.dhl.com/mydhlapi*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressApi* | [**ExpApiAddressValidate**](docs/AddressApi.md#expapiaddressvalidate) | **Get** /address-validate | Validate DHL Express pickup/delivery capabilities at origin/destination
*IdentifierApi* | [**ExpApiIdentifiers**](docs/IdentifierApi.md#expapiidentifiers) | **Get** /identifiers | Service to allocate identifiers upfront for DHL Express Breakbulk or Loose Break Bulk shipments
*InvoiceApi* | [**ExpApiShipmentsInvoiceData**](docs/InvoiceApi.md#expapishipmentsinvoicedata) | **Post** /invoices/upload-invoice-data | Upload Commercial invoice data
*PickupApi* | [**ExpApiPickups**](docs/PickupApi.md#expapipickups) | **Post** /pickups | Create a DHL Express pickup booking request
*PickupApi* | [**ExpApiPickupsCancel**](docs/PickupApi.md#expapipickupscancel) | **Delete** /pickups/{dispatchConfirmationNumber} | Cancel a DHL Express pickup booking request
*PickupApi* | [**ExpApiPickupsUpdate**](docs/PickupApi.md#expapipickupsupdate) | **Patch** /pickups/{dispatchConfirmationNumber} | Update pickup information for an existing DHL Express pickup booking request
*ProductApi* | [**ExpApiProducts**](docs/ProductApi.md#expapiproducts) | **Get** /products | Retrieve available DHL Express products for a one piece Shipment
*RatingApi* | [**ExpApiLandedCost**](docs/RatingApi.md#expapilandedcost) | **Post** /landed-cost | Landed Cost
*RatingApi* | [**ExpApiRates**](docs/RatingApi.md#expapirates) | **Get** /rates | Retrieve Rates for a one piece Shipment
*RatingApi* | [**ExpApiRatesMany**](docs/RatingApi.md#expapiratesmany) | **Post** /rates | Retrieve Rates for Multi-piece Shipments
*ShipmentApi* | [**ExpApiShipments**](docs/ShipmentApi.md#expapishipments) | **Post** /shipments | Create Shipment
*ShipmentApi* | [**ExpApiShipmentsDocumentimage**](docs/ShipmentApi.md#expapishipmentsdocumentimage) | **Get** /shipments/{shipmentTrackingNumber}/get-image | Get Image
*ShipmentApi* | [**ExpApiShipmentsEpod**](docs/ShipmentApi.md#expapishipmentsepod) | **Get** /shipments/{shipmentTrackingNumber}/proof-of-delivery | Electronic Proof of Delivery
*ShipmentApi* | [**ExpApiShipmentsImgUpload**](docs/ShipmentApi.md#expapishipmentsimgupload) | **Patch** /shipments/{shipmentTrackingNumber}/upload-image | Upload Paperless Trade shipment (PLT) images of previously created shipment.
*ShipmentApi* | [**ExpApiShipmentsInvoiceDataAwb**](docs/ShipmentApi.md#expapishipmentsinvoicedataawb) | **Patch** /shipments/{shipmentTrackingNumber}/upload-invoice-data | Upload Commercial Invoice data for your DHL Express shipment
*TrackingApi* | [**ExpApiShipmentsTracking**](docs/TrackingApi.md#expapishipmentstracking) | **Get** /shipments/{shipmentTrackingNumber}/tracking | Track a single DHL Express Shipment
*TrackingApi* | [**ExpApiShipmentsTrackingMulti**](docs/TrackingApi.md#expapishipmentstrackingmulti) | **Get** /tracking | Track a single or multiple DHL Express Shipments

## Documentation For Models

 - [Dimensions](docs/Dimensions.md)
 - [Dimensions1](docs/Dimensions1.md)
 - [Dimensions2](docs/Dimensions2.md)
 - [Dimensions3](docs/Dimensions3.md)
 - [EstimatedDeliveryDate](docs/EstimatedDeliveryDate.md)
 - [EstimatedDeliveryDate1](docs/EstimatedDeliveryDate1.md)
 - [Pickup](docs/Pickup.md)
 - [PickupPickupDetails](docs/PickupPickupDetails.md)
 - [PickupPickupRequestorDetails](docs/PickupPickupRequestorDetails.md)
 - [PickupSpecialInstructions](docs/PickupSpecialInstructions.md)
 - [SupermodelIoLogisticsExpressAccount](docs/SupermodelIoLogisticsExpressAccount.md)
 - [SupermodelIoLogisticsExpressAddress](docs/SupermodelIoLogisticsExpressAddress.md)
 - [SupermodelIoLogisticsExpressAddressCreateShipmentRequest](docs/SupermodelIoLogisticsExpressAddressCreateShipmentRequest.md)
 - [SupermodelIoLogisticsExpressAddressCreateShipmentResponse](docs/SupermodelIoLogisticsExpressAddressCreateShipmentResponse.md)
 - [SupermodelIoLogisticsExpressAddressRatesRequest](docs/SupermodelIoLogisticsExpressAddressRatesRequest.md)
 - [SupermodelIoLogisticsExpressAddressValidateResponse](docs/SupermodelIoLogisticsExpressAddressValidateResponse.md)
 - [SupermodelIoLogisticsExpressAddressValidateResponseAddress](docs/SupermodelIoLogisticsExpressAddressValidateResponseAddress.md)
 - [SupermodelIoLogisticsExpressAddressValidateResponseServiceArea](docs/SupermodelIoLogisticsExpressAddressValidateResponseServiceArea.md)
 - [SupermodelIoLogisticsExpressBankDetailsInner](docs/SupermodelIoLogisticsExpressBankDetailsInner.md)
 - [SupermodelIoLogisticsExpressContact](docs/SupermodelIoLogisticsExpressContact.md)
 - [SupermodelIoLogisticsExpressContactBuyer](docs/SupermodelIoLogisticsExpressContactBuyer.md)
 - [SupermodelIoLogisticsExpressContactCreateShipmentResponse](docs/SupermodelIoLogisticsExpressContactCreateShipmentResponse.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequest](docs/SupermodelIoLogisticsExpressCreateShipmentRequest.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContent](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContent.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclaration.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationAdditionalCharges](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationAdditionalCharges.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCommodityCodes](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCommodityCodes.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomerReferences](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomerReferences.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments1](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationCustomsDocuments1.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationDeclarationNotes](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationDeclarationNotes.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationExporter](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationExporter.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoice](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoice.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceCustomerReferences](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceCustomerReferences.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceIndicativeCustomsValues](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationInvoiceIndicativeCustomsValues.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLicenses](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLicenses.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLineItems](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationLineItems.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationQuantity](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationQuantity.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationRemarks](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationRemarks.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationWeight](docs/SupermodelIoLogisticsExpressCreateShipmentRequestContentExportDeclarationWeight.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsBuyerDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsBuyerDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsExporterDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsExporterDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsImporterDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsImporterDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsPayerDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsPayerDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsReceiverDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsReceiverDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsSellerDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsSellerDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsShipperDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsShipperDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsUltimateConsigneeDetails](docs/SupermodelIoLogisticsExpressCreateShipmentRequestCustomerDetailsUltimateConsigneeDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestGetAdditionalInformation](docs/SupermodelIoLogisticsExpressCreateShipmentRequestGetAdditionalInformation.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestOnDemandDelivery](docs/SupermodelIoLogisticsExpressCreateShipmentRequestOnDemandDelivery.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestOutputImageProperties](docs/SupermodelIoLogisticsExpressCreateShipmentRequestOutputImageProperties.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesCustomerBarcodes](docs/SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesCustomerBarcodes.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesCustomerLogos](docs/SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesCustomerLogos.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesImageOptions](docs/SupermodelIoLogisticsExpressCreateShipmentRequestOutputImagePropertiesImageOptions.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestParentShipment](docs/SupermodelIoLogisticsExpressCreateShipmentRequestParentShipment.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestPrepaidCharges](docs/SupermodelIoLogisticsExpressCreateShipmentRequestPrepaidCharges.md)
 - [SupermodelIoLogisticsExpressCreateShipmentRequestShipmentNotification](docs/SupermodelIoLogisticsExpressCreateShipmentRequestShipmentNotification.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponse](docs/SupermodelIoLogisticsExpressCreateShipmentResponse.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseBarcodeInfo](docs/SupermodelIoLogisticsExpressCreateShipmentResponseBarcodeInfo.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseBarcodeInfoTrackingNumberBarcodes](docs/SupermodelIoLogisticsExpressCreateShipmentResponseBarcodeInfoTrackingNumberBarcodes.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseCustomerDetails](docs/SupermodelIoLogisticsExpressCreateShipmentResponseCustomerDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseCustomerDetailsShipperDetails](docs/SupermodelIoLogisticsExpressCreateShipmentResponseCustomerDetailsShipperDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseDestinationServiceArea](docs/SupermodelIoLogisticsExpressCreateShipmentResponseDestinationServiceArea.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseDocuments](docs/SupermodelIoLogisticsExpressCreateShipmentResponseDocuments.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseDocuments1](docs/SupermodelIoLogisticsExpressCreateShipmentResponseDocuments1.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseEstimatedDeliveryDate](docs/SupermodelIoLogisticsExpressCreateShipmentResponseEstimatedDeliveryDate.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseOriginServiceArea](docs/SupermodelIoLogisticsExpressCreateShipmentResponseOriginServiceArea.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponsePackages](docs/SupermodelIoLogisticsExpressCreateShipmentResponsePackages.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponsePickupDetails](docs/SupermodelIoLogisticsExpressCreateShipmentResponsePickupDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseServiceBreakdown](docs/SupermodelIoLogisticsExpressCreateShipmentResponseServiceBreakdown.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseShipmentCharges](docs/SupermodelIoLogisticsExpressCreateShipmentResponseShipmentCharges.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseShipmentDetails](docs/SupermodelIoLogisticsExpressCreateShipmentResponseShipmentDetails.md)
 - [SupermodelIoLogisticsExpressCreateShipmentResponseValueAddedServices](docs/SupermodelIoLogisticsExpressCreateShipmentResponseValueAddedServices.md)
 - [SupermodelIoLogisticsExpressDocumentImageResponse](docs/SupermodelIoLogisticsExpressDocumentImageResponse.md)
 - [SupermodelIoLogisticsExpressDocumentImageResponseDocuments](docs/SupermodelIoLogisticsExpressDocumentImageResponseDocuments.md)
 - [SupermodelIoLogisticsExpressDocumentImagesInner](docs/SupermodelIoLogisticsExpressDocumentImagesInner.md)
 - [SupermodelIoLogisticsExpressEpodResponse](docs/SupermodelIoLogisticsExpressEpodResponse.md)
 - [SupermodelIoLogisticsExpressEpodResponseDocuments](docs/SupermodelIoLogisticsExpressEpodResponseDocuments.md)
 - [SupermodelIoLogisticsExpressErrorResponse](docs/SupermodelIoLogisticsExpressErrorResponse.md)
 - [SupermodelIoLogisticsExpressExportDeclaration](docs/SupermodelIoLogisticsExpressExportDeclaration.md)
 - [SupermodelIoLogisticsExpressExportDeclarationAdditionalCharges](docs/SupermodelIoLogisticsExpressExportDeclarationAdditionalCharges.md)
 - [SupermodelIoLogisticsExpressExportDeclarationCustomerReferences](docs/SupermodelIoLogisticsExpressExportDeclarationCustomerReferences.md)
 - [SupermodelIoLogisticsExpressExportDeclarationInvoice](docs/SupermodelIoLogisticsExpressExportDeclarationInvoice.md)
 - [SupermodelIoLogisticsExpressExportDeclarationInvoiceCustomerReferences](docs/SupermodelIoLogisticsExpressExportDeclarationInvoiceCustomerReferences.md)
 - [SupermodelIoLogisticsExpressExportDeclarationLineItems](docs/SupermodelIoLogisticsExpressExportDeclarationLineItems.md)
 - [SupermodelIoLogisticsExpressExportDeclarationQuantity](docs/SupermodelIoLogisticsExpressExportDeclarationQuantity.md)
 - [SupermodelIoLogisticsExpressExportDeclarationRemarks](docs/SupermodelIoLogisticsExpressExportDeclarationRemarks.md)
 - [SupermodelIoLogisticsExpressExportDeclarationWeight](docs/SupermodelIoLogisticsExpressExportDeclarationWeight.md)
 - [SupermodelIoLogisticsExpressIdentifier](docs/SupermodelIoLogisticsExpressIdentifier.md)
 - [SupermodelIoLogisticsExpressIdentifierResponse](docs/SupermodelIoLogisticsExpressIdentifierResponse.md)
 - [SupermodelIoLogisticsExpressIdentifierResponseIdentifiers](docs/SupermodelIoLogisticsExpressIdentifierResponseIdentifiers.md)
 - [SupermodelIoLogisticsExpressImageUploadRequest](docs/SupermodelIoLogisticsExpressImageUploadRequest.md)
 - [SupermodelIoLogisticsExpressLandedCostRequest](docs/SupermodelIoLogisticsExpressLandedCostRequest.md)
 - [SupermodelIoLogisticsExpressLandedCostRequestAdditionalQuantityDefinitions](docs/SupermodelIoLogisticsExpressLandedCostRequestAdditionalQuantityDefinitions.md)
 - [SupermodelIoLogisticsExpressLandedCostRequestCharges](docs/SupermodelIoLogisticsExpressLandedCostRequestCharges.md)
 - [SupermodelIoLogisticsExpressLandedCostRequestCustomerDetails](docs/SupermodelIoLogisticsExpressLandedCostRequestCustomerDetails.md)
 - [SupermodelIoLogisticsExpressLandedCostRequestGoodsCharacteristics](docs/SupermodelIoLogisticsExpressLandedCostRequestGoodsCharacteristics.md)
 - [SupermodelIoLogisticsExpressLandedCostRequestItems](docs/SupermodelIoLogisticsExpressLandedCostRequestItems.md)
 - [SupermodelIoLogisticsExpressPackage](docs/SupermodelIoLogisticsExpressPackage.md)
 - [SupermodelIoLogisticsExpressPackageLabelBarcodes](docs/SupermodelIoLogisticsExpressPackageLabelBarcodes.md)
 - [SupermodelIoLogisticsExpressPackageLabelText](docs/SupermodelIoLogisticsExpressPackageLabelText.md)
 - [SupermodelIoLogisticsExpressPackageReference](docs/SupermodelIoLogisticsExpressPackageReference.md)
 - [SupermodelIoLogisticsExpressPackageRr](docs/SupermodelIoLogisticsExpressPackageRr.md)
 - [SupermodelIoLogisticsExpressPickupRequest](docs/SupermodelIoLogisticsExpressPickupRequest.md)
 - [SupermodelIoLogisticsExpressPickupRequestCustomerDetails](docs/SupermodelIoLogisticsExpressPickupRequestCustomerDetails.md)
 - [SupermodelIoLogisticsExpressPickupRequestCustomerDetailsBookingRequestorDetails](docs/SupermodelIoLogisticsExpressPickupRequestCustomerDetailsBookingRequestorDetails.md)
 - [SupermodelIoLogisticsExpressPickupRequestCustomerDetailsShipperDetails](docs/SupermodelIoLogisticsExpressPickupRequestCustomerDetailsShipperDetails.md)
 - [SupermodelIoLogisticsExpressPickupRequestShipmentDetails](docs/SupermodelIoLogisticsExpressPickupRequestShipmentDetails.md)
 - [SupermodelIoLogisticsExpressPickupRequestSpecialInstructions](docs/SupermodelIoLogisticsExpressPickupRequestSpecialInstructions.md)
 - [SupermodelIoLogisticsExpressPickupResponse](docs/SupermodelIoLogisticsExpressPickupResponse.md)
 - [SupermodelIoLogisticsExpressProducts](docs/SupermodelIoLogisticsExpressProducts.md)
 - [SupermodelIoLogisticsExpressProductsBreakdown](docs/SupermodelIoLogisticsExpressProductsBreakdown.md)
 - [SupermodelIoLogisticsExpressProductsDeliveryCapabilities](docs/SupermodelIoLogisticsExpressProductsDeliveryCapabilities.md)
 - [SupermodelIoLogisticsExpressProductsDependencyRuleGroup](docs/SupermodelIoLogisticsExpressProductsDependencyRuleGroup.md)
 - [SupermodelIoLogisticsExpressProductsPickupCapabilities](docs/SupermodelIoLogisticsExpressProductsPickupCapabilities.md)
 - [SupermodelIoLogisticsExpressProductsProducts](docs/SupermodelIoLogisticsExpressProductsProducts.md)
 - [SupermodelIoLogisticsExpressProductsRequiredServiceCodes](docs/SupermodelIoLogisticsExpressProductsRequiredServiceCodes.md)
 - [SupermodelIoLogisticsExpressProductsServiceCodeDependencyRuleGroups](docs/SupermodelIoLogisticsExpressProductsServiceCodeDependencyRuleGroups.md)
 - [SupermodelIoLogisticsExpressProductsServiceCodeMutuallyExclusiveGroups](docs/SupermodelIoLogisticsExpressProductsServiceCodeMutuallyExclusiveGroups.md)
 - [SupermodelIoLogisticsExpressProductsServiceCodes](docs/SupermodelIoLogisticsExpressProductsServiceCodes.md)
 - [SupermodelIoLogisticsExpressRateRequest](docs/SupermodelIoLogisticsExpressRateRequest.md)
 - [SupermodelIoLogisticsExpressRateRequestCustomerDetails](docs/SupermodelIoLogisticsExpressRateRequestCustomerDetails.md)
 - [SupermodelIoLogisticsExpressRateRequestGetAdditionalInformation](docs/SupermodelIoLogisticsExpressRateRequestGetAdditionalInformation.md)
 - [SupermodelIoLogisticsExpressRateRequestMonetaryAmount](docs/SupermodelIoLogisticsExpressRateRequestMonetaryAmount.md)
 - [SupermodelIoLogisticsExpressRateRequestProductsAndServices](docs/SupermodelIoLogisticsExpressRateRequestProductsAndServices.md)
 - [SupermodelIoLogisticsExpressRates](docs/SupermodelIoLogisticsExpressRates.md)
 - [SupermodelIoLogisticsExpressRatesBreakdown](docs/SupermodelIoLogisticsExpressRatesBreakdown.md)
 - [SupermodelIoLogisticsExpressRatesBreakdown1](docs/SupermodelIoLogisticsExpressRatesBreakdown1.md)
 - [SupermodelIoLogisticsExpressRatesDeliveryCapabilities](docs/SupermodelIoLogisticsExpressRatesDeliveryCapabilities.md)
 - [SupermodelIoLogisticsExpressRatesDetailedPriceBreakdown](docs/SupermodelIoLogisticsExpressRatesDetailedPriceBreakdown.md)
 - [SupermodelIoLogisticsExpressRatesExchangeRates](docs/SupermodelIoLogisticsExpressRatesExchangeRates.md)
 - [SupermodelIoLogisticsExpressRatesItems](docs/SupermodelIoLogisticsExpressRatesItems.md)
 - [SupermodelIoLogisticsExpressRatesPickupCapabilities](docs/SupermodelIoLogisticsExpressRatesPickupCapabilities.md)
 - [SupermodelIoLogisticsExpressRatesPriceBreakdown](docs/SupermodelIoLogisticsExpressRatesPriceBreakdown.md)
 - [SupermodelIoLogisticsExpressRatesPriceBreakdown1](docs/SupermodelIoLogisticsExpressRatesPriceBreakdown1.md)
 - [SupermodelIoLogisticsExpressRatesPriceBreakdown2](docs/SupermodelIoLogisticsExpressRatesPriceBreakdown2.md)
 - [SupermodelIoLogisticsExpressRatesProducts](docs/SupermodelIoLogisticsExpressRatesProducts.md)
 - [SupermodelIoLogisticsExpressRatesTotalPrice](docs/SupermodelIoLogisticsExpressRatesTotalPrice.md)
 - [SupermodelIoLogisticsExpressRatesTotalPriceBreakdown](docs/SupermodelIoLogisticsExpressRatesTotalPriceBreakdown.md)
 - [SupermodelIoLogisticsExpressReference](docs/SupermodelIoLogisticsExpressReference.md)
 - [SupermodelIoLogisticsExpressRegistrationNumbers](docs/SupermodelIoLogisticsExpressRegistrationNumbers.md)
 - [SupermodelIoLogisticsExpressTrackingResponse](docs/SupermodelIoLogisticsExpressTrackingResponse.md)
 - [SupermodelIoLogisticsExpressTrackingResponseEvents](docs/SupermodelIoLogisticsExpressTrackingResponseEvents.md)
 - [SupermodelIoLogisticsExpressTrackingResponseEvents1](docs/SupermodelIoLogisticsExpressTrackingResponseEvents1.md)
 - [SupermodelIoLogisticsExpressTrackingResponsePieces](docs/SupermodelIoLogisticsExpressTrackingResponsePieces.md)
 - [SupermodelIoLogisticsExpressTrackingResponseReceiverDetails](docs/SupermodelIoLogisticsExpressTrackingResponseReceiverDetails.md)
 - [SupermodelIoLogisticsExpressTrackingResponseReceiverDetailsPostalAddress](docs/SupermodelIoLogisticsExpressTrackingResponseReceiverDetailsPostalAddress.md)
 - [SupermodelIoLogisticsExpressTrackingResponseReceiverDetailsServiceArea](docs/SupermodelIoLogisticsExpressTrackingResponseReceiverDetailsServiceArea.md)
 - [SupermodelIoLogisticsExpressTrackingResponseServiceArea](docs/SupermodelIoLogisticsExpressTrackingResponseServiceArea.md)
 - [SupermodelIoLogisticsExpressTrackingResponseServiceArea1](docs/SupermodelIoLogisticsExpressTrackingResponseServiceArea1.md)
 - [SupermodelIoLogisticsExpressTrackingResponseShipments](docs/SupermodelIoLogisticsExpressTrackingResponseShipments.md)
 - [SupermodelIoLogisticsExpressTrackingResponseShipperDetails](docs/SupermodelIoLogisticsExpressTrackingResponseShipperDetails.md)
 - [SupermodelIoLogisticsExpressTrackingResponseShipperDetailsPostalAddress](docs/SupermodelIoLogisticsExpressTrackingResponseShipperDetailsPostalAddress.md)
 - [SupermodelIoLogisticsExpressTrackingResponseShipperDetailsServiceArea](docs/SupermodelIoLogisticsExpressTrackingResponseShipperDetailsServiceArea.md)
 - [SupermodelIoLogisticsExpressUpdatePickupRequest](docs/SupermodelIoLogisticsExpressUpdatePickupRequest.md)
 - [SupermodelIoLogisticsExpressUpdatePickupRequestShipmentDetails](docs/SupermodelIoLogisticsExpressUpdatePickupRequestShipmentDetails.md)
 - [SupermodelIoLogisticsExpressUpdatePickupResponse](docs/SupermodelIoLogisticsExpressUpdatePickupResponse.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequest](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequest.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestContent](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestContent.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsBuyerDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsBuyerDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsExporterDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsExporterDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsImporterDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsImporterDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsSellerDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsSellerDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsUltimateConsigneeDetails](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestCustomerDetailsUltimateConsigneeDetails.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImageProperties](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImageProperties.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImagePropertiesImageOptions](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestOutputImagePropertiesImageOptions.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataRequestSid](docs/SupermodelIoLogisticsExpressUploadInvoiceDataRequestSid.md)
 - [SupermodelIoLogisticsExpressUploadInvoiceDataResponse](docs/SupermodelIoLogisticsExpressUploadInvoiceDataResponse.md)
 - [SupermodelIoLogisticsExpressValueAddedServices](docs/SupermodelIoLogisticsExpressValueAddedServices.md)
 - [SupermodelIoLogisticsExpressValueAddedServicesDangerousGoods](docs/SupermodelIoLogisticsExpressValueAddedServicesDangerousGoods.md)
 - [SupermodelIoLogisticsExpressValueAddedServicesRates](docs/SupermodelIoLogisticsExpressValueAddedServicesRates.md)
 - [Weight](docs/Weight.md)
 - [Weight1](docs/Weight1.md)

## Documentation For Authorization

## basicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author


