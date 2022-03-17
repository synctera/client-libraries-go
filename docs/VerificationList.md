# VerificationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verifications** | [**[]Verification**](Verification.md) | Array of verification results. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewVerificationList

`func NewVerificationList(verifications []Verification, ) *VerificationList`

NewVerificationList instantiates a new VerificationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationListWithDefaults

`func NewVerificationListWithDefaults() *VerificationList`

NewVerificationListWithDefaults instantiates a new VerificationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifications

`func (o *VerificationList) GetVerifications() []Verification`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *VerificationList) GetVerificationsOk() (*[]Verification, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *VerificationList) SetVerifications(v []Verification)`

SetVerifications sets Verifications field to given value.


### GetNextPageToken

`func (o *VerificationList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *VerificationList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *VerificationList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *VerificationList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


