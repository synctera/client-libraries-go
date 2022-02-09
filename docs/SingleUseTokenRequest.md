# SingleUseTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The ID of the customer to whom the token will be issued | 
**AccountId** | **string** | The ID of the account to which the token will be linked | 

## Methods

### NewSingleUseTokenRequest

`func NewSingleUseTokenRequest(customerId string, accountId string, ) *SingleUseTokenRequest`

NewSingleUseTokenRequest instantiates a new SingleUseTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleUseTokenRequestWithDefaults

`func NewSingleUseTokenRequestWithDefaults() *SingleUseTokenRequest`

NewSingleUseTokenRequestWithDefaults instantiates a new SingleUseTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *SingleUseTokenRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SingleUseTokenRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SingleUseTokenRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetAccountId

`func (o *SingleUseTokenRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SingleUseTokenRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SingleUseTokenRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


