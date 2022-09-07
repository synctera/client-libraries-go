# Prospect1AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**Email** | **string** | Email | 
**EmailValidation** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Prospect ID | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Most recent updated time | [optional] [readonly] 
**Points** | Pointer to **int32** | Number of referral points | [optional] 
**RecaptchaToken** | Pointer to **string** | A token used to verify the user is not a robot | [optional] 
**ReferralCode** | Pointer to **string** | Referral code | [optional] [readonly] 
**ReferredBy** | Pointer to **string** | UUID of the referrer | [optional] [readonly] 
**ReferredByCode** | Pointer to **string** | Referral code of the referrer | [optional] 
**ReferredProspects** | Pointer to **int32** | Number of people this prospect referred | [optional] [readonly] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**VerificationToken** | Pointer to **string** | Verification token sent to prospect | [optional] [readonly] 
**WaitlistId** | Pointer to **string** | Waitlist ID | [optional] [readonly] 
**WaitlistPosition** | Pointer to **NullableInt32** | Prospect&#39;s number in the waitlist | [optional] [readonly] 

## Methods

### NewProspect1AllOf

`func NewProspect1AllOf(email string, ) *Prospect1AllOf`

NewProspect1AllOf instantiates a new Prospect1AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProspect1AllOfWithDefaults

`func NewProspect1AllOfWithDefaults() *Prospect1AllOf`

NewProspect1AllOfWithDefaults instantiates a new Prospect1AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Prospect1AllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Prospect1AllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Prospect1AllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Prospect1AllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *Prospect1AllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Prospect1AllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Prospect1AllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailValidation

`func (o *Prospect1AllOf) GetEmailValidation() string`

GetEmailValidation returns the EmailValidation field if non-nil, zero value otherwise.

### GetEmailValidationOk

`func (o *Prospect1AllOf) GetEmailValidationOk() (*string, bool)`

GetEmailValidationOk returns a tuple with the EmailValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidation

`func (o *Prospect1AllOf) SetEmailValidation(v string)`

SetEmailValidation sets EmailValidation field to given value.

### HasEmailValidation

`func (o *Prospect1AllOf) HasEmailValidation() bool`

HasEmailValidation returns a boolean if a field has been set.

### GetId

`func (o *Prospect1AllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prospect1AllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prospect1AllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Prospect1AllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Prospect1AllOf) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Prospect1AllOf) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Prospect1AllOf) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Prospect1AllOf) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPoints

`func (o *Prospect1AllOf) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Prospect1AllOf) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Prospect1AllOf) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *Prospect1AllOf) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetRecaptchaToken

`func (o *Prospect1AllOf) GetRecaptchaToken() string`

GetRecaptchaToken returns the RecaptchaToken field if non-nil, zero value otherwise.

### GetRecaptchaTokenOk

`func (o *Prospect1AllOf) GetRecaptchaTokenOk() (*string, bool)`

GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaToken

`func (o *Prospect1AllOf) SetRecaptchaToken(v string)`

SetRecaptchaToken sets RecaptchaToken field to given value.

### HasRecaptchaToken

`func (o *Prospect1AllOf) HasRecaptchaToken() bool`

HasRecaptchaToken returns a boolean if a field has been set.

### GetReferralCode

`func (o *Prospect1AllOf) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *Prospect1AllOf) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *Prospect1AllOf) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *Prospect1AllOf) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### GetReferredBy

`func (o *Prospect1AllOf) GetReferredBy() string`

GetReferredBy returns the ReferredBy field if non-nil, zero value otherwise.

### GetReferredByOk

`func (o *Prospect1AllOf) GetReferredByOk() (*string, bool)`

GetReferredByOk returns a tuple with the ReferredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredBy

`func (o *Prospect1AllOf) SetReferredBy(v string)`

SetReferredBy sets ReferredBy field to given value.

### HasReferredBy

`func (o *Prospect1AllOf) HasReferredBy() bool`

HasReferredBy returns a boolean if a field has been set.

### GetReferredByCode

`func (o *Prospect1AllOf) GetReferredByCode() string`

GetReferredByCode returns the ReferredByCode field if non-nil, zero value otherwise.

### GetReferredByCodeOk

`func (o *Prospect1AllOf) GetReferredByCodeOk() (*string, bool)`

GetReferredByCodeOk returns a tuple with the ReferredByCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredByCode

`func (o *Prospect1AllOf) SetReferredByCode(v string)`

SetReferredByCode sets ReferredByCode field to given value.

### HasReferredByCode

`func (o *Prospect1AllOf) HasReferredByCode() bool`

HasReferredByCode returns a boolean if a field has been set.

### GetReferredProspects

`func (o *Prospect1AllOf) GetReferredProspects() int32`

GetReferredProspects returns the ReferredProspects field if non-nil, zero value otherwise.

### GetReferredProspectsOk

`func (o *Prospect1AllOf) GetReferredProspectsOk() (*int32, bool)`

GetReferredProspectsOk returns a tuple with the ReferredProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredProspects

`func (o *Prospect1AllOf) SetReferredProspects(v int32)`

SetReferredProspects sets ReferredProspects field to given value.

### HasReferredProspects

`func (o *Prospect1AllOf) HasReferredProspects() bool`

HasReferredProspects returns a boolean if a field has been set.

### GetVendorInfo

`func (o *Prospect1AllOf) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Prospect1AllOf) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Prospect1AllOf) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *Prospect1AllOf) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerificationToken

`func (o *Prospect1AllOf) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *Prospect1AllOf) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *Prospect1AllOf) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.

### HasVerificationToken

`func (o *Prospect1AllOf) HasVerificationToken() bool`

HasVerificationToken returns a boolean if a field has been set.

### GetWaitlistId

`func (o *Prospect1AllOf) GetWaitlistId() string`

GetWaitlistId returns the WaitlistId field if non-nil, zero value otherwise.

### GetWaitlistIdOk

`func (o *Prospect1AllOf) GetWaitlistIdOk() (*string, bool)`

GetWaitlistIdOk returns a tuple with the WaitlistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistId

`func (o *Prospect1AllOf) SetWaitlistId(v string)`

SetWaitlistId sets WaitlistId field to given value.

### HasWaitlistId

`func (o *Prospect1AllOf) HasWaitlistId() bool`

HasWaitlistId returns a boolean if a field has been set.

### GetWaitlistPosition

`func (o *Prospect1AllOf) GetWaitlistPosition() int32`

GetWaitlistPosition returns the WaitlistPosition field if non-nil, zero value otherwise.

### GetWaitlistPositionOk

`func (o *Prospect1AllOf) GetWaitlistPositionOk() (*int32, bool)`

GetWaitlistPositionOk returns a tuple with the WaitlistPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistPosition

`func (o *Prospect1AllOf) SetWaitlistPosition(v int32)`

SetWaitlistPosition sets WaitlistPosition field to given value.

### HasWaitlistPosition

`func (o *Prospect1AllOf) HasWaitlistPosition() bool`

HasWaitlistPosition returns a boolean if a field has been set.

### SetWaitlistPositionNil

`func (o *Prospect1AllOf) SetWaitlistPositionNil(b bool)`

 SetWaitlistPositionNil sets the value for WaitlistPosition to be an explicit nil

### UnsetWaitlistPosition
`func (o *Prospect1AllOf) UnsetWaitlistPosition()`

UnsetWaitlistPosition ensures that no value is present for WaitlistPosition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


