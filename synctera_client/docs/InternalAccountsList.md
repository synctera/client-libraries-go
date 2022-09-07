# InternalAccountsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalAccounts** | [**[]InternalAccount**](InternalAccount.md) | Array of internal accounts | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewInternalAccountsList

`func NewInternalAccountsList(internalAccounts []InternalAccount, ) *InternalAccountsList`

NewInternalAccountsList instantiates a new InternalAccountsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalAccountsListWithDefaults

`func NewInternalAccountsListWithDefaults() *InternalAccountsList`

NewInternalAccountsListWithDefaults instantiates a new InternalAccountsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalAccounts

`func (o *InternalAccountsList) GetInternalAccounts() []InternalAccount`

GetInternalAccounts returns the InternalAccounts field if non-nil, zero value otherwise.

### GetInternalAccountsOk

`func (o *InternalAccountsList) GetInternalAccountsOk() (*[]InternalAccount, bool)`

GetInternalAccountsOk returns a tuple with the InternalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccounts

`func (o *InternalAccountsList) SetInternalAccounts(v []InternalAccount)`

SetInternalAccounts sets InternalAccounts field to given value.


### GetNextPageToken

`func (o *InternalAccountsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InternalAccountsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InternalAccountsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InternalAccountsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


