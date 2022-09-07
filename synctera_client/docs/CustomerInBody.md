# CustomerInBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | Pointer to **string** | Customer&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | [optional] 
**FirstName** | Pointer to **string** | Customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Customer&#39;s last name | [optional] 
**Status** | **string** | Customer&#39;s status | 
**BanStatus** | Pointer to [**BanStatus**](BanStatus.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Email** | Pointer to **string** | Customer&#39;s email | [optional] 
**Id** | Pointer to **string** | Customer unique identifier | [optional] [readonly] 
**KycExempt** | Pointer to **bool** | Customer&#39;s KYC exemption | [optional] [readonly] 
**KycLastRun** | Pointer to **time.Time** | Date and time KYC was last run on the customer | [optional] [readonly] 
**KycStatus** | Pointer to [**CustomerKycStatus**](CustomerKycStatus.md) |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied metadata. Do not use to store PII. | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**PhoneNumber** | Pointer to **string** | Customer&#39;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] 
**RelatedCustomers** | Pointer to [**[]Relationship1**](Relationship1.md) | Customer&#39;s relationships with other accounts eg. guardian | [optional] 
**ShippingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Ssn** | Pointer to **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789). | [optional] 
**SsnSource** | Pointer to [**SsnSource**](SsnSource.md) |  | [optional] 

## Methods

### NewCustomerInBody

`func NewCustomerInBody(status string, ) *CustomerInBody`

NewCustomerInBody instantiates a new CustomerInBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInBodyWithDefaults

`func NewCustomerInBodyWithDefaults() *CustomerInBody`

NewCustomerInBodyWithDefaults instantiates a new CustomerInBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDob

`func (o *CustomerInBody) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *CustomerInBody) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *CustomerInBody) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *CustomerInBody) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerInBody) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerInBody) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerInBody) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerInBody) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerInBody) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerInBody) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerInBody) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerInBody) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerInBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerInBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerInBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBanStatus

`func (o *CustomerInBody) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *CustomerInBody) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *CustomerInBody) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *CustomerInBody) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetCreationTime

`func (o *CustomerInBody) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CustomerInBody) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CustomerInBody) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CustomerInBody) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerInBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerInBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *CustomerInBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerInBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerInBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerInBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKycExempt

`func (o *CustomerInBody) GetKycExempt() bool`

GetKycExempt returns the KycExempt field if non-nil, zero value otherwise.

### GetKycExemptOk

`func (o *CustomerInBody) GetKycExemptOk() (*bool, bool)`

GetKycExemptOk returns a tuple with the KycExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycExempt

`func (o *CustomerInBody) SetKycExempt(v bool)`

SetKycExempt sets KycExempt field to given value.

### HasKycExempt

`func (o *CustomerInBody) HasKycExempt() bool`

HasKycExempt returns a boolean if a field has been set.

### GetKycLastRun

`func (o *CustomerInBody) GetKycLastRun() time.Time`

GetKycLastRun returns the KycLastRun field if non-nil, zero value otherwise.

### GetKycLastRunOk

`func (o *CustomerInBody) GetKycLastRunOk() (*time.Time, bool)`

GetKycLastRunOk returns a tuple with the KycLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLastRun

`func (o *CustomerInBody) SetKycLastRun(v time.Time)`

SetKycLastRun sets KycLastRun field to given value.

### HasKycLastRun

`func (o *CustomerInBody) HasKycLastRun() bool`

HasKycLastRun returns a boolean if a field has been set.

### GetKycStatus

`func (o *CustomerInBody) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CustomerInBody) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CustomerInBody) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *CustomerInBody) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CustomerInBody) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CustomerInBody) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CustomerInBody) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CustomerInBody) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *CustomerInBody) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *CustomerInBody) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *CustomerInBody) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *CustomerInBody) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerInBody) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerInBody) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerInBody) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerInBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *CustomerInBody) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CustomerInBody) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CustomerInBody) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CustomerInBody) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerInBody) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerInBody) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerInBody) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerInBody) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRelatedCustomers

`func (o *CustomerInBody) GetRelatedCustomers() []Relationship1`

GetRelatedCustomers returns the RelatedCustomers field if non-nil, zero value otherwise.

### GetRelatedCustomersOk

`func (o *CustomerInBody) GetRelatedCustomersOk() (*[]Relationship1, bool)`

GetRelatedCustomersOk returns a tuple with the RelatedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCustomers

`func (o *CustomerInBody) SetRelatedCustomers(v []Relationship1)`

SetRelatedCustomers sets RelatedCustomers field to given value.

### HasRelatedCustomers

`func (o *CustomerInBody) HasRelatedCustomers() bool`

HasRelatedCustomers returns a boolean if a field has been set.

### GetShippingAddress

`func (o *CustomerInBody) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CustomerInBody) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CustomerInBody) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CustomerInBody) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *CustomerInBody) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *CustomerInBody) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *CustomerInBody) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *CustomerInBody) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnSource

`func (o *CustomerInBody) GetSsnSource() SsnSource`

GetSsnSource returns the SsnSource field if non-nil, zero value otherwise.

### GetSsnSourceOk

`func (o *CustomerInBody) GetSsnSourceOk() (*SsnSource, bool)`

GetSsnSourceOk returns a tuple with the SsnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnSource

`func (o *CustomerInBody) SetSsnSource(v SsnSource)`

SetSsnSource sets SsnSource field to given value.

### HasSsnSource

`func (o *CustomerInBody) HasSsnSource() bool`

HasSsnSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


