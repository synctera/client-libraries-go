# ReconciliationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByteData** | **string** | Base64url encoded image | 
**FileName** | **string** | Filename of the data to be reconciled | 

## Methods

### NewReconciliationInput

`func NewReconciliationInput(byteData string, fileName string, ) *ReconciliationInput`

NewReconciliationInput instantiates a new ReconciliationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationInputWithDefaults

`func NewReconciliationInputWithDefaults() *ReconciliationInput`

NewReconciliationInputWithDefaults instantiates a new ReconciliationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByteData

`func (o *ReconciliationInput) GetByteData() string`

GetByteData returns the ByteData field if non-nil, zero value otherwise.

### GetByteDataOk

`func (o *ReconciliationInput) GetByteDataOk() (*string, bool)`

GetByteDataOk returns a tuple with the ByteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteData

`func (o *ReconciliationInput) SetByteData(v string)`

SetByteData sets ByteData field to given value.


### GetFileName

`func (o *ReconciliationInput) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ReconciliationInput) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ReconciliationInput) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


