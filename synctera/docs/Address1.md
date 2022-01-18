# Address1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street address line 1 | 
**AddressLine2** | Pointer to **string** | String address line 2 | [optional] 
**City** | Pointer to **string** | City | [optional] 
**CountryCode** | **string** | ISO-3166-1 Alpha-2 country code | 
**PostalCode** | Pointer to **string** | Postal code | [optional] 
**State** | Pointer to **string** | State, region, province, or prefecture | [optional] 

## Methods

### NewAddress1

`func NewAddress1(addressLine1 string, countryCode string, ) *Address1`

NewAddress1 instantiates a new Address1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddress1WithDefaults

`func NewAddress1WithDefaults() *Address1`

NewAddress1WithDefaults instantiates a new Address1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *Address1) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address1) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address1) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *Address1) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address1) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address1) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Address1) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *Address1) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address1) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address1) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address1) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *Address1) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Address1) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Address1) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPostalCode

`func (o *Address1) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address1) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address1) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address1) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *Address1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address1) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


