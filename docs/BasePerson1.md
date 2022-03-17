# BasePerson1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Dob** | Pointer to [**oapi.Date**](oapi.Date.md) | Person&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD). | [optional] 
**Email** | Pointer to **string** | Person&#39;s email. | [optional] 
**FirstName** | Pointer to **string** | Person&#39;s first name. | [optional] 
**Id** | Pointer to **string** | Person&#39;s unique identifier. | [optional] [readonly] 
**IsCustomer** | Pointer to **bool** | Distinguish personal and business customers from non-customer parties. | [optional] 
**LastName** | Pointer to **string** | Person&#39;s last name. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**MiddleName** | Pointer to **string** | Person&#39;s middle name. | [optional] 
**PhoneNumber** | Pointer to **string** | Person&#39;s mobile phone number with country code in E.164 format. | [optional] 
**ShippingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Ssn** | Pointer to **string** | Person&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC. The response contains the last 4 digits only (e.g. 6789). | [optional] 
**Status** | Pointer to [**Status1**](Status1.md) |  | [optional] 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYC verification was last run on the person. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 

## Methods

### NewBasePerson1

`func NewBasePerson1() *BasePerson1`

NewBasePerson1 instantiates a new BasePerson1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasePerson1WithDefaults

`func NewBasePerson1WithDefaults() *BasePerson1`

NewBasePerson1WithDefaults instantiates a new BasePerson1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *BasePerson1) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BasePerson1) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BasePerson1) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BasePerson1) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDob

`func (o *BasePerson1) GetDob() oapi.Date`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *BasePerson1) GetDobOk() (*oapi.Date, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *BasePerson1) SetDob(v oapi.Date)`

SetDob sets Dob field to given value.

### HasDob

`func (o *BasePerson1) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *BasePerson1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BasePerson1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BasePerson1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BasePerson1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BasePerson1) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BasePerson1) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BasePerson1) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BasePerson1) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *BasePerson1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasePerson1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasePerson1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasePerson1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *BasePerson1) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *BasePerson1) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *BasePerson1) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.

### HasIsCustomer

`func (o *BasePerson1) HasIsCustomer() bool`

HasIsCustomer returns a boolean if a field has been set.

### GetLastName

`func (o *BasePerson1) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BasePerson1) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BasePerson1) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BasePerson1) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BasePerson1) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BasePerson1) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BasePerson1) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BasePerson1) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *BasePerson1) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *BasePerson1) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *BasePerson1) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *BasePerson1) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *BasePerson1) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BasePerson1) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BasePerson1) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BasePerson1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *BasePerson1) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BasePerson1) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BasePerson1) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BasePerson1) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BasePerson1) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BasePerson1) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BasePerson1) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BasePerson1) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *BasePerson1) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *BasePerson1) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *BasePerson1) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *BasePerson1) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *BasePerson1) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *BasePerson1) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *BasePerson1) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *BasePerson1) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetStatus

`func (o *BasePerson1) GetStatus() Status1`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BasePerson1) GetStatusOk() (*Status1, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BasePerson1) SetStatus(v Status1)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BasePerson1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *BasePerson1) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *BasePerson1) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *BasePerson1) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *BasePerson1) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *BasePerson1) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *BasePerson1) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *BasePerson1) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *BasePerson1) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


