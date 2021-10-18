# ExternalAccountUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The user&#39;s email address | [optional] 
**EmailAddressVerifiedTime** | Pointer to **time.Time** | The date and time the email address was verified | [optional] 
**PhoneNumber** | Pointer to **string** | The user&#39;s phone number | [optional] 
**PhoneNumberVerifiedTime** | Pointer to **time.Time** | The date and time the phone number was verified | [optional] 
**Ssn** | Pointer to **string** | To be provided in the format \&quot;ddd-dd-dddd\&quot;.  | [optional] 
**Status** | Pointer to **string** | The status of this alert | [optional] 
**VendorCustomerId** | **string** | Unique identifier for this user | 

## Methods

### NewExternalAccountUser

`func NewExternalAccountUser(vendorCustomerId string, ) *ExternalAccountUser`

NewExternalAccountUser instantiates a new ExternalAccountUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountUserWithDefaults

`func NewExternalAccountUserWithDefaults() *ExternalAccountUser`

NewExternalAccountUserWithDefaults instantiates a new ExternalAccountUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ExternalAccountUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ExternalAccountUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ExternalAccountUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ExternalAccountUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAddressVerifiedTime

`func (o *ExternalAccountUser) GetEmailAddressVerifiedTime() time.Time`

GetEmailAddressVerifiedTime returns the EmailAddressVerifiedTime field if non-nil, zero value otherwise.

### GetEmailAddressVerifiedTimeOk

`func (o *ExternalAccountUser) GetEmailAddressVerifiedTimeOk() (*time.Time, bool)`

GetEmailAddressVerifiedTimeOk returns a tuple with the EmailAddressVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressVerifiedTime

`func (o *ExternalAccountUser) SetEmailAddressVerifiedTime(v time.Time)`

SetEmailAddressVerifiedTime sets EmailAddressVerifiedTime field to given value.

### HasEmailAddressVerifiedTime

`func (o *ExternalAccountUser) HasEmailAddressVerifiedTime() bool`

HasEmailAddressVerifiedTime returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ExternalAccountUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ExternalAccountUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ExternalAccountUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ExternalAccountUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberVerifiedTime

`func (o *ExternalAccountUser) GetPhoneNumberVerifiedTime() time.Time`

GetPhoneNumberVerifiedTime returns the PhoneNumberVerifiedTime field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedTimeOk

`func (o *ExternalAccountUser) GetPhoneNumberVerifiedTimeOk() (*time.Time, bool)`

GetPhoneNumberVerifiedTimeOk returns a tuple with the PhoneNumberVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerifiedTime

`func (o *ExternalAccountUser) SetPhoneNumberVerifiedTime(v time.Time)`

SetPhoneNumberVerifiedTime sets PhoneNumberVerifiedTime field to given value.

### HasPhoneNumberVerifiedTime

`func (o *ExternalAccountUser) HasPhoneNumberVerifiedTime() bool`

HasPhoneNumberVerifiedTime returns a boolean if a field has been set.

### GetSsn

`func (o *ExternalAccountUser) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *ExternalAccountUser) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *ExternalAccountUser) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *ExternalAccountUser) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalAccountUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalAccountUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalAccountUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalAccountUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorCustomerId

`func (o *ExternalAccountUser) GetVendorCustomerId() string`

GetVendorCustomerId returns the VendorCustomerId field if non-nil, zero value otherwise.

### GetVendorCustomerIdOk

`func (o *ExternalAccountUser) GetVendorCustomerIdOk() (*string, bool)`

GetVendorCustomerIdOk returns a tuple with the VendorCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCustomerId

`func (o *ExternalAccountUser) SetVendorCustomerId(v string)`

SetVendorCustomerId sets VendorCustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


