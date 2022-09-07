# BasePerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewBasePerson

`func NewBasePerson() *BasePerson`

NewBasePerson instantiates a new BasePerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasePersonWithDefaults

`func NewBasePersonWithDefaults() *BasePerson`

NewBasePersonWithDefaults instantiates a new BasePerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanStatus

`func (o *BasePerson) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *BasePerson) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *BasePerson) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *BasePerson) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetCreationTime

`func (o *BasePerson) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BasePerson) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BasePerson) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BasePerson) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *BasePerson) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BasePerson) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BasePerson) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BasePerson) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *BasePerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasePerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasePerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasePerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKycExempt

`func (o *BasePerson) GetKycExempt() bool`

GetKycExempt returns the KycExempt field if non-nil, zero value otherwise.

### GetKycExemptOk

`func (o *BasePerson) GetKycExemptOk() (*bool, bool)`

GetKycExemptOk returns a tuple with the KycExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycExempt

`func (o *BasePerson) SetKycExempt(v bool)`

SetKycExempt sets KycExempt field to given value.

### HasKycExempt

`func (o *BasePerson) HasKycExempt() bool`

HasKycExempt returns a boolean if a field has been set.

### GetKycLastRun

`func (o *BasePerson) GetKycLastRun() time.Time`

GetKycLastRun returns the KycLastRun field if non-nil, zero value otherwise.

### GetKycLastRunOk

`func (o *BasePerson) GetKycLastRunOk() (*time.Time, bool)`

GetKycLastRunOk returns a tuple with the KycLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLastRun

`func (o *BasePerson) SetKycLastRun(v time.Time)`

SetKycLastRun sets KycLastRun field to given value.

### HasKycLastRun

`func (o *BasePerson) HasKycLastRun() bool`

HasKycLastRun returns a boolean if a field has been set.

### GetKycStatus

`func (o *BasePerson) GetKycStatus() CustomerKycStatus`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *BasePerson) GetKycStatusOk() (*CustomerKycStatus, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *BasePerson) SetKycStatus(v CustomerKycStatus)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *BasePerson) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BasePerson) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BasePerson) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BasePerson) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BasePerson) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *BasePerson) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *BasePerson) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *BasePerson) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *BasePerson) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *BasePerson) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BasePerson) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BasePerson) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BasePerson) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *BasePerson) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BasePerson) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BasePerson) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BasePerson) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BasePerson) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BasePerson) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BasePerson) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BasePerson) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRelatedCustomers

`func (o *BasePerson) GetRelatedCustomers() []Relationship1`

GetRelatedCustomers returns the RelatedCustomers field if non-nil, zero value otherwise.

### GetRelatedCustomersOk

`func (o *BasePerson) GetRelatedCustomersOk() (*[]Relationship1, bool)`

GetRelatedCustomersOk returns a tuple with the RelatedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCustomers

`func (o *BasePerson) SetRelatedCustomers(v []Relationship1)`

SetRelatedCustomers sets RelatedCustomers field to given value.

### HasRelatedCustomers

`func (o *BasePerson) HasRelatedCustomers() bool`

HasRelatedCustomers returns a boolean if a field has been set.

### GetShippingAddress

`func (o *BasePerson) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *BasePerson) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *BasePerson) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *BasePerson) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *BasePerson) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *BasePerson) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *BasePerson) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *BasePerson) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnSource

`func (o *BasePerson) GetSsnSource() SsnSource`

GetSsnSource returns the SsnSource field if non-nil, zero value otherwise.

### GetSsnSourceOk

`func (o *BasePerson) GetSsnSourceOk() (*SsnSource, bool)`

GetSsnSourceOk returns a tuple with the SsnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnSource

`func (o *BasePerson) SetSsnSource(v SsnSource)`

SetSsnSource sets SsnSource field to given value.

### HasSsnSource

`func (o *BasePerson) HasSsnSource() bool`

HasSsnSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


