# ReconciliationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reconciliations** | [**[]Reconciliation**](Reconciliation.md) | Array of reconciliations | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewReconciliationList

`func NewReconciliationList(reconciliations []Reconciliation, ) *ReconciliationList`

NewReconciliationList instantiates a new ReconciliationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationListWithDefaults

`func NewReconciliationListWithDefaults() *ReconciliationList`

NewReconciliationListWithDefaults instantiates a new ReconciliationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReconciliations

`func (o *ReconciliationList) GetReconciliations() []Reconciliation`

GetReconciliations returns the Reconciliations field if non-nil, zero value otherwise.

### GetReconciliationsOk

`func (o *ReconciliationList) GetReconciliationsOk() (*[]Reconciliation, bool)`

GetReconciliationsOk returns a tuple with the Reconciliations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliations

`func (o *ReconciliationList) SetReconciliations(v []Reconciliation)`

SetReconciliations sets Reconciliations field to given value.


### GetNextPageToken

`func (o *ReconciliationList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReconciliationList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ReconciliationList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ReconciliationList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


