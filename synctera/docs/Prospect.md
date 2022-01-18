# Prospect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | Pointer to **string** | Customer&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | [optional] 
**FirstName** | Pointer to **string** | Customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Customer&#39;s last name | [optional] 
**Status** | **string** | Customer&#39;s status | 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Email** | Pointer to **string** | Customer&#39;s email | [optional] 
**Id** | Pointer to **string** | Customer unique identifier | [optional] [readonly] 
**KycExempt** | Pointer to **bool** | Customer&#39;s KYC exemption | [optional] [readonly] 
**KycLastRun** | Pointer to **time.Time** | Date and time KYC was last run on the customer | [optional] [readonly] 
**KycStatus** | Pointer to [**CustomerKycStatus**](CustomerKycStatus.md) |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied metadata. Do not use to store PII. | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**PhoneNumber** | Pointer to **string** | Customer&#39;s mobile phone number with country code in E.164 format | [optional] 
**RelatedCustomers** | Pointer to [**[]Relationship1**](Relationship1.md) | Customer&#39;s relationships with other accounts eg. guardian | [optional] 
**ShippingAddress** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**Ssn** | Pointer to **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789). | [optional] 

## Methods

### NewProspect

`func NewProspect(status string, ) *Prospect`

NewProspect instantiates a new Prospect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProspectWithDefaults

`func NewProspectWithDefaults() *Prospect`

NewProspectWithDefaults instantiates a new Prospect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDob

`func (o *Prospect) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Prospect) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Prospect) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Prospect) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetFirstName

`func (o *Prospect) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Prospect) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Prospect) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Prospect) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Prospect) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Prospect) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Prospect) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Prospect) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetStatus

`func (o *Prospect) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Prospect) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Prospect) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreationTime

`func (o *Prospect) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Prospect) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Prospect) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Prospect) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *Prospect) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Prospect) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Prospect) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Prospect) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Prospect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prospect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prospect) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Prospect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKycExempt

`func (o *Prospect) GetKycExempt() bool`

GetKycExempt returns the KycExempt field if non-nil, zero value otherwise.

### GetKycExemptOk

`func (o *Prospect) GetKycExemptOk() (*bool, bool)`

GetKycExemptOk returns a tuple with the KycExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycExempt

`func (o *Prospect) SetKycExempt(v bool)`

SetKycExempt sets KycExempt field to given value.

### HasKycExempt

`func (o *Prospect) HasKycExempt() bool`

HasKycExempt returns a boolean if a field has been set.

### GetKycLastRun

`func (o *Prospect) GetKycLastRun() time.Time`

GetKycLastRun returns the KycLastRun field if non-nil, zero value otherwise.

### GetKycLastRunOk

`func (o *Prospect) GetKycLastRunOk() (*time.Time, bool)`

GetKycLastRunOk returns a tuple with the KycLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLastRun

`func (o *Prospect) SetKycLastRun(v time.Time)`

SetKycLastRun sets KycLastRun field to given value.

### HasKycLastRun

`func (o *Prospect) HasKycLastRun() bool`

HasKycLastRun returns a boolean if a field has been set.

### GetKycStatus

`func (o *Prospect) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *Prospect) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *Prospect) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *Prospect) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Prospect) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Prospect) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Prospect) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Prospect) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Prospect) GetLegalAddress() Address1`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Prospect) GetLegalAddressOk() (*Address1, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Prospect) SetLegalAddress(v Address1)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *Prospect) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *Prospect) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Prospect) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Prospect) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Prospect) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *Prospect) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Prospect) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Prospect) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Prospect) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Prospect) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Prospect) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Prospect) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Prospect) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRelatedCustomers

`func (o *Prospect) GetRelatedCustomers() []Relationship1`

GetRelatedCustomers returns the RelatedCustomers field if non-nil, zero value otherwise.

### GetRelatedCustomersOk

`func (o *Prospect) GetRelatedCustomersOk() (*[]Relationship1, bool)`

GetRelatedCustomersOk returns a tuple with the RelatedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCustomers

`func (o *Prospect) SetRelatedCustomers(v []Relationship1)`

SetRelatedCustomers sets RelatedCustomers field to given value.

### HasRelatedCustomers

`func (o *Prospect) HasRelatedCustomers() bool`

HasRelatedCustomers returns a boolean if a field has been set.

### GetShippingAddress

`func (o *Prospect) GetShippingAddress() Address1`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Prospect) GetShippingAddressOk() (*Address1, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Prospect) SetShippingAddress(v Address1)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Prospect) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *Prospect) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *Prospect) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *Prospect) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *Prospect) HasSsn() bool`

HasSsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


