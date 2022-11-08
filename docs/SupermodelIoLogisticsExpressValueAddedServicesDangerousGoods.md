# SupermodelIoLogisticsExpressValueAddedServicesDangerousGoods

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **string** | Please enter valid DHL Express Dangerous good content id (please contact your DHL Express IT representative for the relevant content ID code if you are shipping Dan | [default to null]
**DryIceTotalNetWeight** | **float64** | Please enter dry ice total net weight when shipping dry ice | [optional] [default to null]
**UnCode** | **string** | Please enter comma separated UN code(s) | [optional] [default to null]
**CustomDescription** | **string** | The customDescription node contains the customized Dangerous Goods statement to declare contents accurately. The customDescription node value will be displayed in the Transport Label and Waybill Document, replacing the default IATA Dangerous Goods statement constructed based on contentId node.&lt;BR&gt;            Multiple customDescription nodes from multiple dangerousGoods nodes will be concatenated using comma separator with combined maximum limit of 200 characters.&lt;BR&gt;            &lt;BR&gt;            It is recommended to use customDescription for entire shipment for each dangerousGoods to fully utilize customDescription printout in Transport Label and Waybill Document.&lt;BR&gt;            &lt;BR&gt;            Note: For &#x27;customDescription&#x27; usage, ensure all &#x27;dangerousGoods&#x27; segments are including the &#x27;customDescription&#x27; field value. Any of the dangerousGoods does not provide with customDescription field value will be ignored from printing in Transport Label and Waybill Document.   | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

