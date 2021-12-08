# CustomerVerificationResultList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verifications** | [**[]CustomerVerificationResult**](CustomerVerificationResult.md) | Array of verification results. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCustomerVerificationResultList

`func NewCustomerVerificationResultList(verifications []CustomerVerificationResult, ) *CustomerVerificationResultList`

NewCustomerVerificationResultList instantiates a new CustomerVerificationResultList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerificationResultListWithDefaults

`func NewCustomerVerificationResultListWithDefaults() *CustomerVerificationResultList`

NewCustomerVerificationResultListWithDefaults instantiates a new CustomerVerificationResultList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifications

`func (o *CustomerVerificationResultList) GetVerifications() []CustomerVerificationResult`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CustomerVerificationResultList) GetVerificationsOk() (*[]CustomerVerificationResult, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CustomerVerificationResultList) SetVerifications(v []CustomerVerificationResult)`

SetVerifications sets Verifications field to given value.


### GetNextPageToken

`func (o *CustomerVerificationResultList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CustomerVerificationResultList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CustomerVerificationResultList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CustomerVerificationResultList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


