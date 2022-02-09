# TokenListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DigitalWalletTokens** | Pointer to [**[]DigitalWalletTokenResponse**](DigitalWalletTokenResponse.md) | Array of digital wallet token information of a card | [optional] 

## Methods

### NewTokenListResponse

`func NewTokenListResponse() *TokenListResponse`

NewTokenListResponse instantiates a new TokenListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListResponseWithDefaults

`func NewTokenListResponseWithDefaults() *TokenListResponse`

NewTokenListResponseWithDefaults instantiates a new TokenListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigitalWalletTokens

`func (o *TokenListResponse) GetDigitalWalletTokens() []DigitalWalletTokenResponse`

GetDigitalWalletTokens returns the DigitalWalletTokens field if non-nil, zero value otherwise.

### GetDigitalWalletTokensOk

`func (o *TokenListResponse) GetDigitalWalletTokensOk() (*[]DigitalWalletTokenResponse, bool)`

GetDigitalWalletTokensOk returns a tuple with the DigitalWalletTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokens

`func (o *TokenListResponse) SetDigitalWalletTokens(v []DigitalWalletTokenResponse)`

SetDigitalWalletTokens sets DigitalWalletTokens field to given value.

### HasDigitalWalletTokens

`func (o *TokenListResponse) HasDigitalWalletTokens() bool`

HasDigitalWalletTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


