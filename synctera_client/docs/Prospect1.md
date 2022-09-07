# Prospect1

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
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Metadata** | Pointer to **string** | Client supplied json metadata to be stored with the prospect | [optional] 
**Source** | Pointer to **string** | Source of prospect | [optional] 
**Status** | Pointer to [**ProspectStatus**](ProspectStatus.md) |  | [optional] 

## Methods

### NewProspect1

`func NewProspect1(email string, ) *Prospect1`

NewProspect1 instantiates a new Prospect1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProspect1WithDefaults

`func NewProspect1WithDefaults() *Prospect1`

NewProspect1WithDefaults instantiates a new Prospect1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Prospect1) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Prospect1) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Prospect1) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Prospect1) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *Prospect1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Prospect1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Prospect1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailValidation

`func (o *Prospect1) GetEmailValidation() string`

GetEmailValidation returns the EmailValidation field if non-nil, zero value otherwise.

### GetEmailValidationOk

`func (o *Prospect1) GetEmailValidationOk() (*string, bool)`

GetEmailValidationOk returns a tuple with the EmailValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidation

`func (o *Prospect1) SetEmailValidation(v string)`

SetEmailValidation sets EmailValidation field to given value.

### HasEmailValidation

`func (o *Prospect1) HasEmailValidation() bool`

HasEmailValidation returns a boolean if a field has been set.

### GetId

`func (o *Prospect1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prospect1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prospect1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Prospect1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Prospect1) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Prospect1) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Prospect1) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Prospect1) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPoints

`func (o *Prospect1) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Prospect1) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Prospect1) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *Prospect1) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetRecaptchaToken

`func (o *Prospect1) GetRecaptchaToken() string`

GetRecaptchaToken returns the RecaptchaToken field if non-nil, zero value otherwise.

### GetRecaptchaTokenOk

`func (o *Prospect1) GetRecaptchaTokenOk() (*string, bool)`

GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaToken

`func (o *Prospect1) SetRecaptchaToken(v string)`

SetRecaptchaToken sets RecaptchaToken field to given value.

### HasRecaptchaToken

`func (o *Prospect1) HasRecaptchaToken() bool`

HasRecaptchaToken returns a boolean if a field has been set.

### GetReferralCode

`func (o *Prospect1) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *Prospect1) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *Prospect1) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *Prospect1) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### GetReferredBy

`func (o *Prospect1) GetReferredBy() string`

GetReferredBy returns the ReferredBy field if non-nil, zero value otherwise.

### GetReferredByOk

`func (o *Prospect1) GetReferredByOk() (*string, bool)`

GetReferredByOk returns a tuple with the ReferredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredBy

`func (o *Prospect1) SetReferredBy(v string)`

SetReferredBy sets ReferredBy field to given value.

### HasReferredBy

`func (o *Prospect1) HasReferredBy() bool`

HasReferredBy returns a boolean if a field has been set.

### GetReferredByCode

`func (o *Prospect1) GetReferredByCode() string`

GetReferredByCode returns the ReferredByCode field if non-nil, zero value otherwise.

### GetReferredByCodeOk

`func (o *Prospect1) GetReferredByCodeOk() (*string, bool)`

GetReferredByCodeOk returns a tuple with the ReferredByCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredByCode

`func (o *Prospect1) SetReferredByCode(v string)`

SetReferredByCode sets ReferredByCode field to given value.

### HasReferredByCode

`func (o *Prospect1) HasReferredByCode() bool`

HasReferredByCode returns a boolean if a field has been set.

### GetReferredProspects

`func (o *Prospect1) GetReferredProspects() int32`

GetReferredProspects returns the ReferredProspects field if non-nil, zero value otherwise.

### GetReferredProspectsOk

`func (o *Prospect1) GetReferredProspectsOk() (*int32, bool)`

GetReferredProspectsOk returns a tuple with the ReferredProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredProspects

`func (o *Prospect1) SetReferredProspects(v int32)`

SetReferredProspects sets ReferredProspects field to given value.

### HasReferredProspects

`func (o *Prospect1) HasReferredProspects() bool`

HasReferredProspects returns a boolean if a field has been set.

### GetVendorInfo

`func (o *Prospect1) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Prospect1) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Prospect1) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *Prospect1) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerificationToken

`func (o *Prospect1) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *Prospect1) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *Prospect1) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.

### HasVerificationToken

`func (o *Prospect1) HasVerificationToken() bool`

HasVerificationToken returns a boolean if a field has been set.

### GetWaitlistId

`func (o *Prospect1) GetWaitlistId() string`

GetWaitlistId returns the WaitlistId field if non-nil, zero value otherwise.

### GetWaitlistIdOk

`func (o *Prospect1) GetWaitlistIdOk() (*string, bool)`

GetWaitlistIdOk returns a tuple with the WaitlistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistId

`func (o *Prospect1) SetWaitlistId(v string)`

SetWaitlistId sets WaitlistId field to given value.

### HasWaitlistId

`func (o *Prospect1) HasWaitlistId() bool`

HasWaitlistId returns a boolean if a field has been set.

### GetWaitlistPosition

`func (o *Prospect1) GetWaitlistPosition() int32`

GetWaitlistPosition returns the WaitlistPosition field if non-nil, zero value otherwise.

### GetWaitlistPositionOk

`func (o *Prospect1) GetWaitlistPositionOk() (*int32, bool)`

GetWaitlistPositionOk returns a tuple with the WaitlistPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistPosition

`func (o *Prospect1) SetWaitlistPosition(v int32)`

SetWaitlistPosition sets WaitlistPosition field to given value.

### HasWaitlistPosition

`func (o *Prospect1) HasWaitlistPosition() bool`

HasWaitlistPosition returns a boolean if a field has been set.

### SetWaitlistPositionNil

`func (o *Prospect1) SetWaitlistPositionNil(b bool)`

 SetWaitlistPositionNil sets the value for WaitlistPosition to be an explicit nil

### UnsetWaitlistPosition
`func (o *Prospect1) UnsetWaitlistPosition()`

UnsetWaitlistPosition ensures that no value is present for WaitlistPosition, not even an explicit nil
### GetFirstName

`func (o *Prospect1) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Prospect1) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Prospect1) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Prospect1) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Prospect1) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Prospect1) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Prospect1) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Prospect1) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *Prospect1) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Prospect1) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Prospect1) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Prospect1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSource

`func (o *Prospect1) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Prospect1) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Prospect1) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Prospect1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *Prospect1) GetStatus() ProspectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Prospect1) GetStatusOk() (*ProspectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Prospect1) SetStatus(v ProspectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Prospect1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


