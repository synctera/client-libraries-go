# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BanStatus** | Pointer to [**BanStatus**](BanStatus.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Dob** | Pointer to **string** | Person&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD). | [optional] 
**Email** | Pointer to **string** | Person&#39;s email. | [optional] 
**FirstName** | Pointer to **string** | Person&#39;s first name. | [optional] 
**Id** | Pointer to **string** | Person&#39;s unique identifier. | [optional] [readonly] 
**IsCustomer** | **bool** | True for personal and business customers with a direct relationship with the fintech or bank. | 
**LastName** | Pointer to **string** | Person&#39;s last name. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**MiddleName** | Pointer to **string** | Person&#39;s middle name. | [optional] 
**PhoneNumber** | Pointer to **string** | Person&#39;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated | [optional] 
**ShippingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Ssn** | Pointer to **string** | Person&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC. The response contains the last 4 digits only (e.g. 6789). | [optional] 
**SsnSource** | Pointer to [**SsnSource**](SsnSource.md) |  | [optional] 
**Status** | [**Status1**](Status1.md) |  | 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYC verification was last run on the person. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 

## Methods

### NewPerson

`func NewPerson(isCustomer bool, status Status1, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanStatus

`func (o *Person) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *Person) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *Person) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *Person) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetCreationTime

`func (o *Person) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Person) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Person) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Person) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDob

`func (o *Person) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Person) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Person) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Person) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *Person) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Person) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Person) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Person) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *Person) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Person) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Person) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Person) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *Person) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Person) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Person) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Person) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *Person) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *Person) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *Person) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.


### GetLastName

`func (o *Person) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Person) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Person) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Person) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Person) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Person) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Person) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Person) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Person) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Person) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Person) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *Person) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *Person) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Person) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Person) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Person) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *Person) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Person) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Person) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Person) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Person) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Person) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Person) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Person) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *Person) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Person) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Person) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Person) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *Person) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *Person) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *Person) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *Person) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnSource

`func (o *Person) GetSsnSource() SsnSource`

GetSsnSource returns the SsnSource field if non-nil, zero value otherwise.

### GetSsnSourceOk

`func (o *Person) GetSsnSourceOk() (*SsnSource, bool)`

GetSsnSourceOk returns a tuple with the SsnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnSource

`func (o *Person) SetSsnSource(v SsnSource)`

SetSsnSource sets SsnSource field to given value.

### HasSsnSource

`func (o *Person) HasSsnSource() bool`

HasSsnSource returns a boolean if a field has been set.

### GetStatus

`func (o *Person) GetStatus() Status1`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Person) GetStatusOk() (*Status1, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Person) SetStatus(v Status1)`

SetStatus sets Status field to given value.


### GetVerificationLastRun

`func (o *Person) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *Person) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *Person) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *Person) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *Person) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *Person) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *Person) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *Person) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


