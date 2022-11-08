# SupermodelIoLogisticsExpressPackage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | **string** | Please contact your DHL Express representative if you wish to use a DHL specific package otherwise ignore this element. | [optional] [default to null]
**Weight** | **float64** | The weight of the package. | [default to null]
**Dimensions** | [***Dimensions**](dimensions.md) |  | [default to null]
**CustomerReferences** | [**[]SupermodelIoLogisticsExpressPackageReference**](supermodelIoLogisticsExpressPackageReference.md) | Here you can declare your customer references for each package | [optional] [default to null]
**Identifiers** | [**[]SupermodelIoLogisticsExpressIdentifier**](supermodelIoLogisticsExpressIdentifier.md) | Identifiers section is on the package level where you can optionaly provide a DHL Express waybill number. This has to be enabled by your DHL Express IT contact. | [optional] [default to null]
**Description** | **string** | Please enter description of content for each package | [optional] [default to null]
**LabelBarcodes** | [**[]SupermodelIoLogisticsExpressPackageLabelBarcodes**](supermodelIoLogisticsExpressPackage_labelBarcodes.md) | This allows you to define up to two bespoke barcodes on the DHL Express Tranport label. To use this feature please set outputImageProperties/imageOptions/templateName as ECOM26_84CI_003 for typeCode&#x3D;label | [optional] [default to null]
**LabelText** | [**[]SupermodelIoLogisticsExpressPackageLabelText**](supermodelIoLogisticsExpressPackage_labelText.md) | This allows you to enter up to two bespoke texts on the DHL Express Tranport label. To use this feature please set outputImageProperties/imageOptions/templateName as ECOM26_84CI_003 for typeCode&#x3D;label | [optional] [default to null]
**LabelDescription** | **string** | Please enter additional customer description | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

