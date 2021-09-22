# TransactionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** |  | [optional] 
**CardExpirationDateYymm** | Pointer to **string** |  | [optional] 
**DatabaseTransactionTimeout** | Pointer to **int32** |  | [optional] 
**EncryptionKeyId** | Pointer to **string** |  | [optional] 
**IsAsync** | Pointer to **bool** |  | [optional] [default to false]
**PreAuthTimeLimit** | Pointer to **string** |  | [optional] 
**SendExpirationDate** | Pointer to **bool** |  | [optional] [default to false]
**SendTrackData** | Pointer to **bool** |  | [optional] [default to false]
**TransactionTimeoutThresholdSeconds** | Pointer to **int64** |  | [optional] 
**TransactionToken** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionOptions

`func NewTransactionOptions() *TransactionOptions`

NewTransactionOptions instantiates a new TransactionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOptionsWithDefaults

`func NewTransactionOptionsWithDefaults() *TransactionOptions`

NewTransactionOptionsWithDefaults instantiates a new TransactionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *TransactionOptions) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *TransactionOptions) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *TransactionOptions) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *TransactionOptions) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCardExpirationDateYymm

`func (o *TransactionOptions) GetCardExpirationDateYymm() string`

GetCardExpirationDateYymm returns the CardExpirationDateYymm field if non-nil, zero value otherwise.

### GetCardExpirationDateYymmOk

`func (o *TransactionOptions) GetCardExpirationDateYymmOk() (*string, bool)`

GetCardExpirationDateYymmOk returns a tuple with the CardExpirationDateYymm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationDateYymm

`func (o *TransactionOptions) SetCardExpirationDateYymm(v string)`

SetCardExpirationDateYymm sets CardExpirationDateYymm field to given value.

### HasCardExpirationDateYymm

`func (o *TransactionOptions) HasCardExpirationDateYymm() bool`

HasCardExpirationDateYymm returns a boolean if a field has been set.

### GetDatabaseTransactionTimeout

`func (o *TransactionOptions) GetDatabaseTransactionTimeout() int32`

GetDatabaseTransactionTimeout returns the DatabaseTransactionTimeout field if non-nil, zero value otherwise.

### GetDatabaseTransactionTimeoutOk

`func (o *TransactionOptions) GetDatabaseTransactionTimeoutOk() (*int32, bool)`

GetDatabaseTransactionTimeoutOk returns a tuple with the DatabaseTransactionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseTransactionTimeout

`func (o *TransactionOptions) SetDatabaseTransactionTimeout(v int32)`

SetDatabaseTransactionTimeout sets DatabaseTransactionTimeout field to given value.

### HasDatabaseTransactionTimeout

`func (o *TransactionOptions) HasDatabaseTransactionTimeout() bool`

HasDatabaseTransactionTimeout returns a boolean if a field has been set.

### GetEncryptionKeyId

`func (o *TransactionOptions) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *TransactionOptions) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *TransactionOptions) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *TransactionOptions) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### GetIsAsync

`func (o *TransactionOptions) GetIsAsync() bool`

GetIsAsync returns the IsAsync field if non-nil, zero value otherwise.

### GetIsAsyncOk

`func (o *TransactionOptions) GetIsAsyncOk() (*bool, bool)`

GetIsAsyncOk returns a tuple with the IsAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsync

`func (o *TransactionOptions) SetIsAsync(v bool)`

SetIsAsync sets IsAsync field to given value.

### HasIsAsync

`func (o *TransactionOptions) HasIsAsync() bool`

HasIsAsync returns a boolean if a field has been set.

### GetPreAuthTimeLimit

`func (o *TransactionOptions) GetPreAuthTimeLimit() string`

GetPreAuthTimeLimit returns the PreAuthTimeLimit field if non-nil, zero value otherwise.

### GetPreAuthTimeLimitOk

`func (o *TransactionOptions) GetPreAuthTimeLimitOk() (*string, bool)`

GetPreAuthTimeLimitOk returns a tuple with the PreAuthTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthTimeLimit

`func (o *TransactionOptions) SetPreAuthTimeLimit(v string)`

SetPreAuthTimeLimit sets PreAuthTimeLimit field to given value.

### HasPreAuthTimeLimit

`func (o *TransactionOptions) HasPreAuthTimeLimit() bool`

HasPreAuthTimeLimit returns a boolean if a field has been set.

### GetSendExpirationDate

`func (o *TransactionOptions) GetSendExpirationDate() bool`

GetSendExpirationDate returns the SendExpirationDate field if non-nil, zero value otherwise.

### GetSendExpirationDateOk

`func (o *TransactionOptions) GetSendExpirationDateOk() (*bool, bool)`

GetSendExpirationDateOk returns a tuple with the SendExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendExpirationDate

`func (o *TransactionOptions) SetSendExpirationDate(v bool)`

SetSendExpirationDate sets SendExpirationDate field to given value.

### HasSendExpirationDate

`func (o *TransactionOptions) HasSendExpirationDate() bool`

HasSendExpirationDate returns a boolean if a field has been set.

### GetSendTrackData

`func (o *TransactionOptions) GetSendTrackData() bool`

GetSendTrackData returns the SendTrackData field if non-nil, zero value otherwise.

### GetSendTrackDataOk

`func (o *TransactionOptions) GetSendTrackDataOk() (*bool, bool)`

GetSendTrackDataOk returns a tuple with the SendTrackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTrackData

`func (o *TransactionOptions) SetSendTrackData(v bool)`

SetSendTrackData sets SendTrackData field to given value.

### HasSendTrackData

`func (o *TransactionOptions) HasSendTrackData() bool`

HasSendTrackData returns a boolean if a field has been set.

### GetTransactionTimeoutThresholdSeconds

`func (o *TransactionOptions) GetTransactionTimeoutThresholdSeconds() int64`

GetTransactionTimeoutThresholdSeconds returns the TransactionTimeoutThresholdSeconds field if non-nil, zero value otherwise.

### GetTransactionTimeoutThresholdSecondsOk

`func (o *TransactionOptions) GetTransactionTimeoutThresholdSecondsOk() (*int64, bool)`

GetTransactionTimeoutThresholdSecondsOk returns a tuple with the TransactionTimeoutThresholdSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimeoutThresholdSeconds

`func (o *TransactionOptions) SetTransactionTimeoutThresholdSeconds(v int64)`

SetTransactionTimeoutThresholdSeconds sets TransactionTimeoutThresholdSeconds field to given value.

### HasTransactionTimeoutThresholdSeconds

`func (o *TransactionOptions) HasTransactionTimeoutThresholdSeconds() bool`

HasTransactionTimeoutThresholdSeconds returns a boolean if a field has been set.

### GetTransactionToken

`func (o *TransactionOptions) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *TransactionOptions) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *TransactionOptions) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.

### HasTransactionToken

`func (o *TransactionOptions) HasTransactionToken() bool`

HasTransactionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


