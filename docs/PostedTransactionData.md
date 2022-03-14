# PostedTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**HoldId** | Pointer to **string** | The uuid of the hold (pending transaction) that this transaction originated from, if any. | [optional] 
**Lines** | [**[]TransactionLine**](TransactionLine.md) | The set of accounting entries associated with this transaction. For example, a debit to a customer account will have a corresponding credit in a general ledger account. | 
**Memo** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**OriginalTrx** | Pointer to **string** | The \&quot;original\&quot; transaction that this transaction is related to. This is only populated in the case of reversed transactions. | [optional] 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewPostedTransactionData

`func NewPostedTransactionData(externalData map[string]interface{}, lines []TransactionLine, memo string, metadata map[string]interface{}, riskInfo map[string]interface{}, userData map[string]interface{}, ) *PostedTransactionData`

NewPostedTransactionData instantiates a new PostedTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostedTransactionDataWithDefaults

`func NewPostedTransactionDataWithDefaults() *PostedTransactionData`

NewPostedTransactionDataWithDefaults instantiates a new PostedTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalData

`func (o *PostedTransactionData) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *PostedTransactionData) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *PostedTransactionData) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *PostedTransactionData) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *PostedTransactionData) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetHoldId

`func (o *PostedTransactionData) GetHoldId() string`

GetHoldId returns the HoldId field if non-nil, zero value otherwise.

### GetHoldIdOk

`func (o *PostedTransactionData) GetHoldIdOk() (*string, bool)`

GetHoldIdOk returns a tuple with the HoldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldId

`func (o *PostedTransactionData) SetHoldId(v string)`

SetHoldId sets HoldId field to given value.

### HasHoldId

`func (o *PostedTransactionData) HasHoldId() bool`

HasHoldId returns a boolean if a field has been set.

### GetLines

`func (o *PostedTransactionData) GetLines() []TransactionLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *PostedTransactionData) GetLinesOk() (*[]TransactionLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *PostedTransactionData) SetLines(v []TransactionLine)`

SetLines sets Lines field to given value.


### GetMemo

`func (o *PostedTransactionData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PostedTransactionData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PostedTransactionData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetMetadata

`func (o *PostedTransactionData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostedTransactionData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostedTransactionData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PostedTransactionData) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PostedTransactionData) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginalTrx

`func (o *PostedTransactionData) GetOriginalTrx() string`

GetOriginalTrx returns the OriginalTrx field if non-nil, zero value otherwise.

### GetOriginalTrxOk

`func (o *PostedTransactionData) GetOriginalTrxOk() (*string, bool)`

GetOriginalTrxOk returns a tuple with the OriginalTrx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTrx

`func (o *PostedTransactionData) SetOriginalTrx(v string)`

SetOriginalTrx sets OriginalTrx field to given value.

### HasOriginalTrx

`func (o *PostedTransactionData) HasOriginalTrx() bool`

HasOriginalTrx returns a boolean if a field has been set.

### GetRiskInfo

`func (o *PostedTransactionData) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *PostedTransactionData) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *PostedTransactionData) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *PostedTransactionData) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *PostedTransactionData) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetUserData

`func (o *PostedTransactionData) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PostedTransactionData) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PostedTransactionData) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *PostedTransactionData) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *PostedTransactionData) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


