# TransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalData** | Pointer to **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | [optional] 
**Lines** | [**[]TransactionLine**](TransactionLine.md) | The set of accounting entries associated with this transaction. For example, a debit to a customer account will have a corresponding credit in a general ledger account. | 
**Memo** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTransactionData

`func NewTransactionData(lines []TransactionLine, memo string, ) *TransactionData`

NewTransactionData instantiates a new TransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataWithDefaults

`func NewTransactionDataWithDefaults() *TransactionData`

NewTransactionDataWithDefaults instantiates a new TransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalData

`func (o *TransactionData) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *TransactionData) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *TransactionData) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *TransactionData) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetLines

`func (o *TransactionData) GetLines() []TransactionLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *TransactionData) GetLinesOk() (*[]TransactionLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *TransactionData) SetLines(v []TransactionLine)`

SetLines sets Lines field to given value.


### GetMemo

`func (o *TransactionData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetMetadata

`func (o *TransactionData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


