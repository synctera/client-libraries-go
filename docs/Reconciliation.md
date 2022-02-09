# Reconciliation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Reconciliation ID | [readonly] 
**FileName** | **string** | Filename of the data to be reconciled | [readonly] 
**IngestionStatus** | [**IngestionStatus**](IngestionStatus.md) |  | 

## Methods

### NewReconciliation

`func NewReconciliation(id string, fileName string, ingestionStatus IngestionStatus, ) *Reconciliation`

NewReconciliation instantiates a new Reconciliation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationWithDefaults

`func NewReconciliationWithDefaults() *Reconciliation`

NewReconciliationWithDefaults instantiates a new Reconciliation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reconciliation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reconciliation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reconciliation) SetId(v string)`

SetId sets Id field to given value.


### GetFileName

`func (o *Reconciliation) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Reconciliation) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Reconciliation) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetIngestionStatus

`func (o *Reconciliation) GetIngestionStatus() IngestionStatus`

GetIngestionStatus returns the IngestionStatus field if non-nil, zero value otherwise.

### GetIngestionStatusOk

`func (o *Reconciliation) GetIngestionStatusOk() (*IngestionStatus, bool)`

GetIngestionStatusOk returns a tuple with the IngestionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionStatus

`func (o *Reconciliation) SetIngestionStatus(v IngestionStatus)`

SetIngestionStatus sets IngestionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


