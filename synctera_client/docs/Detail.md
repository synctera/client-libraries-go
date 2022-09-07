# Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Human-readable description explaining the individual check. | [optional] 
**Label** | Pointer to **string** | Human-readable grouping describing the aspect of the customer&#39;s identity examined by this check. | [optional] 
**Result** | Pointer to **string** | The result of the individual check. One of the following: * &#x60;PASS&#x60; – the check passed contributing to a positive outcome (or accepted verification result). * &#x60;WARN&#x60; – the results of the check were inconclusive and might require review. * &#x60;FAIL&#x60; – the check failed and might result in a failing outcome (or rejected verification_result).  | [optional] 
**VendorCode** | Pointer to **string** | Machine-readable description of the individual check. This field contains vendor-specific terms and may not be populated in all cases. | [optional] 

## Methods

### NewDetail

`func NewDetail() *Detail`

NewDetail instantiates a new Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailWithDefaults

`func NewDetailWithDefaults() *Detail`

NewDetailWithDefaults instantiates a new Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Detail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Detail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Detail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Detail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *Detail) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Detail) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Detail) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Detail) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetResult

`func (o *Detail) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Detail) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Detail) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Detail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetVendorCode

`func (o *Detail) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *Detail) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *Detail) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *Detail) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


