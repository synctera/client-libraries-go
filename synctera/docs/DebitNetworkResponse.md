# DebitNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | indicates whether debit network is active | 
**CreationTime** | **time.Time** | The timestamp representing when the debit network was created | [readonly] 
**EndDate** | Pointer to **time.Time** | The time when debit network became inactive | [optional] 
**Id** | **string** | Debit Network ID | [readonly] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the debit network was last modified | [readonly] 
**Name** | **string** | The name describing the debit network | 
**StartDate** | Pointer to **time.Time** | The time when debit network goes live | [optional] 

## Methods

### NewDebitNetworkResponse

`func NewDebitNetworkResponse(active bool, creationTime time.Time, id string, lastModifiedTime time.Time, name string, ) *DebitNetworkResponse`

NewDebitNetworkResponse instantiates a new DebitNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitNetworkResponseWithDefaults

`func NewDebitNetworkResponseWithDefaults() *DebitNetworkResponse`

NewDebitNetworkResponseWithDefaults instantiates a new DebitNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DebitNetworkResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DebitNetworkResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DebitNetworkResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreationTime

`func (o *DebitNetworkResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DebitNetworkResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DebitNetworkResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetEndDate

`func (o *DebitNetworkResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DebitNetworkResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DebitNetworkResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DebitNetworkResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *DebitNetworkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebitNetworkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebitNetworkResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedTime

`func (o *DebitNetworkResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DebitNetworkResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DebitNetworkResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *DebitNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebitNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebitNetworkResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *DebitNetworkResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DebitNetworkResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DebitNetworkResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DebitNetworkResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


