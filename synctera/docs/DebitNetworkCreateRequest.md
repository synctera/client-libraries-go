# DebitNetworkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | indicates whether debit network is active | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the debit network was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when debit network became inactive | [optional] 
**Id** | Pointer to **string** | Debit Network ID | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the debit network was last modified | [optional] [readonly] 
**Name** | **string** | The name describing the debit network | 
**StartDate** | Pointer to **time.Time** | The time when debit network goes live | [optional] 

## Methods

### NewDebitNetworkCreateRequest

`func NewDebitNetworkCreateRequest(name string, ) *DebitNetworkCreateRequest`

NewDebitNetworkCreateRequest instantiates a new DebitNetworkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitNetworkCreateRequestWithDefaults

`func NewDebitNetworkCreateRequestWithDefaults() *DebitNetworkCreateRequest`

NewDebitNetworkCreateRequestWithDefaults instantiates a new DebitNetworkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DebitNetworkCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DebitNetworkCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DebitNetworkCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DebitNetworkCreateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreationTime

`func (o *DebitNetworkCreateRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DebitNetworkCreateRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DebitNetworkCreateRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DebitNetworkCreateRequest) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *DebitNetworkCreateRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DebitNetworkCreateRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DebitNetworkCreateRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DebitNetworkCreateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *DebitNetworkCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebitNetworkCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebitNetworkCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DebitNetworkCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DebitNetworkCreateRequest) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DebitNetworkCreateRequest) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DebitNetworkCreateRequest) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DebitNetworkCreateRequest) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *DebitNetworkCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebitNetworkCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebitNetworkCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *DebitNetworkCreateRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DebitNetworkCreateRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DebitNetworkCreateRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DebitNetworkCreateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


