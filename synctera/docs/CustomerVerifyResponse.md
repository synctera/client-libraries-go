# CustomerVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycStatus** | [**CustomerKycStatus**](CustomerKycStatus.md) |  | 
**Verifications** | [**[]CustomerVerificationResult**](CustomerVerificationResult.md) | Array of Verification results | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCustomerVerifyResponse

`func NewCustomerVerifyResponse(kycStatus CustomerKycStatus, verifications []CustomerVerificationResult, ) *CustomerVerifyResponse`

NewCustomerVerifyResponse instantiates a new CustomerVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerifyResponseWithDefaults

`func NewCustomerVerifyResponseWithDefaults() *CustomerVerifyResponse`

NewCustomerVerifyResponseWithDefaults instantiates a new CustomerVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycStatus

`func (o *CustomerVerifyResponse) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CustomerVerifyResponse) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CustomerVerifyResponse) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.


### GetVerifications

`func (o *CustomerVerifyResponse) GetVerifications() []CustomerVerificationResult`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CustomerVerifyResponse) GetVerificationsOk() (*[]CustomerVerificationResult, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CustomerVerifyResponse) SetVerifications(v []CustomerVerificationResult)`

SetVerifications sets Verifications field to given value.


### GetNextPageToken

`func (o *CustomerVerifyResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CustomerVerifyResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CustomerVerifyResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CustomerVerifyResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


