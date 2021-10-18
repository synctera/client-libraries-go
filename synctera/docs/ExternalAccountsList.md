# ExternalAccountsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAccounts** | [**[]ExternalAccount**](ExternalAccount.md) | Array of external accounts | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewExternalAccountsList

`func NewExternalAccountsList(externalAccounts []ExternalAccount, ) *ExternalAccountsList`

NewExternalAccountsList instantiates a new ExternalAccountsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountsListWithDefaults

`func NewExternalAccountsListWithDefaults() *ExternalAccountsList`

NewExternalAccountsListWithDefaults instantiates a new ExternalAccountsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAccounts

`func (o *ExternalAccountsList) GetExternalAccounts() []ExternalAccount`

GetExternalAccounts returns the ExternalAccounts field if non-nil, zero value otherwise.

### GetExternalAccountsOk

`func (o *ExternalAccountsList) GetExternalAccountsOk() (*[]ExternalAccount, bool)`

GetExternalAccountsOk returns a tuple with the ExternalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccounts

`func (o *ExternalAccountsList) SetExternalAccounts(v []ExternalAccount)`

SetExternalAccounts sets ExternalAccounts field to given value.


### GetNextPageToken

`func (o *ExternalAccountsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ExternalAccountsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ExternalAccountsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ExternalAccountsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


