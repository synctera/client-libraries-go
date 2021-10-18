# CustomerVerifyResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycStatus** | [**CustomerKycStatus**](CustomerKycStatus.md) |  | 
**Verifications** | [**[]CustomerVerificationResult**](CustomerVerificationResult.md) | Array of Verification results | 

## Methods

### NewCustomerVerifyResponseAllOf

`func NewCustomerVerifyResponseAllOf(kycStatus CustomerKycStatus, verifications []CustomerVerificationResult, ) *CustomerVerifyResponseAllOf`

NewCustomerVerifyResponseAllOf instantiates a new CustomerVerifyResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerifyResponseAllOfWithDefaults

`func NewCustomerVerifyResponseAllOfWithDefaults() *CustomerVerifyResponseAllOf`

NewCustomerVerifyResponseAllOfWithDefaults instantiates a new CustomerVerifyResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycStatus

`func (o *CustomerVerifyResponseAllOf) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CustomerVerifyResponseAllOf) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CustomerVerifyResponseAllOf) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.


### GetVerifications

`func (o *CustomerVerifyResponseAllOf) GetVerifications() []CustomerVerificationResult`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CustomerVerifyResponseAllOf) GetVerificationsOk() (*[]CustomerVerificationResult, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CustomerVerifyResponseAllOf) SetVerifications(v []CustomerVerificationResult)`

SetVerifications sets Verifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


