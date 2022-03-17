# Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Identifier for the type of check performed. | [optional] 
**Description** | Pointer to **string** | Human-readable description explaining the result. | [optional] 
**Label** | Pointer to **string** | Human-readable version of the category. | [optional] 
**Result** | Pointer to **string** | The result of the individual check. One of the following: * &#x60;PASS&#x60; – the check passed contributing to a positive outcome (or accepted verification result). * &#x60;WARN&#x60; – the results of the check were inconclusive and might require review. * &#x60;FAIL&#x60; – the check failed and might result in a failing outcome (or rejected verification_result).  | [optional] 

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

### GetCategory

`func (o *Detail) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Detail) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Detail) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Detail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


