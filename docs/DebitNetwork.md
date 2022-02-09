# DebitNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Debit Network ID | [optional] [readonly] 
**Name** | Pointer to **string** | The name describing the debit network | [optional] 
**Active** | Pointer to **bool** | indicates whether debit network is active | [optional] 
**StartDate** | Pointer to **time.Time** | The time when debit network goes live | [optional] 
**EndDate** | Pointer to **time.Time** | The time when debit network became inactive | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the debit network was created | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the debit network was last modified | [optional] [readonly] 

## Methods

### NewDebitNetwork

`func NewDebitNetwork() *DebitNetwork`

NewDebitNetwork instantiates a new DebitNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitNetworkWithDefaults

`func NewDebitNetworkWithDefaults() *DebitNetwork`

NewDebitNetworkWithDefaults instantiates a new DebitNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DebitNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebitNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebitNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DebitNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DebitNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebitNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebitNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DebitNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *DebitNetwork) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DebitNetwork) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DebitNetwork) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DebitNetwork) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartDate

`func (o *DebitNetwork) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DebitNetwork) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DebitNetwork) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DebitNetwork) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *DebitNetwork) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DebitNetwork) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DebitNetwork) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DebitNetwork) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreationTime

`func (o *DebitNetwork) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DebitNetwork) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DebitNetwork) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DebitNetwork) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DebitNetwork) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DebitNetwork) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DebitNetwork) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DebitNetwork) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


