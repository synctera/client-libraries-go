# VerifyResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationStatus** | [**VerificationStatus**](VerificationStatus.md) |  | 
**Verifications** | [**[]Verification**](Verification.md) | Array of verification results. | 

## Methods

### NewVerifyResponseAllOf

`func NewVerifyResponseAllOf(verificationStatus VerificationStatus, verifications []Verification, ) *VerifyResponseAllOf`

NewVerifyResponseAllOf instantiates a new VerifyResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyResponseAllOfWithDefaults

`func NewVerifyResponseAllOfWithDefaults() *VerifyResponseAllOf`

NewVerifyResponseAllOfWithDefaults instantiates a new VerifyResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationStatus

`func (o *VerifyResponseAllOf) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerifyResponseAllOf) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerifyResponseAllOf) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetVerifications

`func (o *VerifyResponseAllOf) GetVerifications() []Verification`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *VerifyResponseAllOf) GetVerificationsOk() (*[]Verification, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *VerifyResponseAllOf) SetVerifications(v []Verification)`

SetVerifications sets Verifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


