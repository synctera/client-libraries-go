# CustomerInPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | [**oapi.Date**](oapi.Date.md) | Customer&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | 
**FirstName** | **string** | Customer&#39;s first name | 
**LastName** | **string** | Customer&#39;s last name | 
**Status** | **string** | Customer&#39;s status | 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Email** | Pointer to **string** | Customer&#39;s email | [optional] 
**Id** | Pointer to **string** | Customer unique identifier | [optional] [readonly] 
**KycExempt** | Pointer to **bool** | Customer&#39;s KYC exemption | [optional] [readonly] 
**KycLastRun** | Pointer to **time.Time** | Date and time KYC was last run on the customer | [optional] [readonly] 
**KycStatus** | Pointer to [**CustomerKycStatus**](CustomerKycStatus.md) |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied metadata. Do not use to store PII. | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**PhoneNumber** | Pointer to **string** | Customer&#39;s mobile phone number with country code in E.164 format | [optional] 
**RelatedCustomers** | Pointer to [**[]Relationship1**](Relationship1.md) | Customer&#39;s relationships with other accounts eg. guardian | [optional] 
**ShippingAddress** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**Ssn** | Pointer to **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789). | [optional] 

## Methods

### NewCustomerInPath

`func NewCustomerInPath(dob oapi.Date, firstName string, lastName string, status string, ) *CustomerInPath`

NewCustomerInPath instantiates a new CustomerInPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInPathWithDefaults

`func NewCustomerInPathWithDefaults() *CustomerInPath`

NewCustomerInPathWithDefaults instantiates a new CustomerInPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDob

`func (o *CustomerInPath) GetDob() oapi.Date`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *CustomerInPath) GetDobOk() (*oapi.Date, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *CustomerInPath) SetDob(v oapi.Date)`

SetDob sets Dob field to given value.


### GetFirstName

`func (o *CustomerInPath) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerInPath) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerInPath) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CustomerInPath) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerInPath) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerInPath) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetStatus

`func (o *CustomerInPath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerInPath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerInPath) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreationTime

`func (o *CustomerInPath) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CustomerInPath) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CustomerInPath) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CustomerInPath) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerInPath) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInPath) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInPath) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerInPath) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *CustomerInPath) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerInPath) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerInPath) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerInPath) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKycExempt

`func (o *CustomerInPath) GetKycExempt() bool`

GetKycExempt returns the KycExempt field if non-nil, zero value otherwise.

### GetKycExemptOk

`func (o *CustomerInPath) GetKycExemptOk() (*bool, bool)`

GetKycExemptOk returns a tuple with the KycExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycExempt

`func (o *CustomerInPath) SetKycExempt(v bool)`

SetKycExempt sets KycExempt field to given value.

### HasKycExempt

`func (o *CustomerInPath) HasKycExempt() bool`

HasKycExempt returns a boolean if a field has been set.

### GetKycLastRun

`func (o *CustomerInPath) GetKycLastRun() time.Time`

GetKycLastRun returns the KycLastRun field if non-nil, zero value otherwise.

### GetKycLastRunOk

`func (o *CustomerInPath) GetKycLastRunOk() (*time.Time, bool)`

GetKycLastRunOk returns a tuple with the KycLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLastRun

`func (o *CustomerInPath) SetKycLastRun(v time.Time)`

SetKycLastRun sets KycLastRun field to given value.

### HasKycLastRun

`func (o *CustomerInPath) HasKycLastRun() bool`

HasKycLastRun returns a boolean if a field has been set.

### GetKycStatus

`func (o *CustomerInPath) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CustomerInPath) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CustomerInPath) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *CustomerInPath) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CustomerInPath) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CustomerInPath) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CustomerInPath) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CustomerInPath) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *CustomerInPath) GetLegalAddress() Address1`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *CustomerInPath) GetLegalAddressOk() (*Address1, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *CustomerInPath) SetLegalAddress(v Address1)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *CustomerInPath) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerInPath) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerInPath) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerInPath) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerInPath) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *CustomerInPath) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CustomerInPath) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CustomerInPath) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CustomerInPath) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerInPath) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerInPath) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerInPath) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerInPath) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRelatedCustomers

`func (o *CustomerInPath) GetRelatedCustomers() []Relationship1`

GetRelatedCustomers returns the RelatedCustomers field if non-nil, zero value otherwise.

### GetRelatedCustomersOk

`func (o *CustomerInPath) GetRelatedCustomersOk() (*[]Relationship1, bool)`

GetRelatedCustomersOk returns a tuple with the RelatedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCustomers

`func (o *CustomerInPath) SetRelatedCustomers(v []Relationship1)`

SetRelatedCustomers sets RelatedCustomers field to given value.

### HasRelatedCustomers

`func (o *CustomerInPath) HasRelatedCustomers() bool`

HasRelatedCustomers returns a boolean if a field has been set.

### GetShippingAddress

`func (o *CustomerInPath) GetShippingAddress() Address1`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CustomerInPath) GetShippingAddressOk() (*Address1, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CustomerInPath) SetShippingAddress(v Address1)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CustomerInPath) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *CustomerInPath) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *CustomerInPath) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *CustomerInPath) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *CustomerInPath) HasSsn() bool`

HasSsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


