# ExternalAccountVendorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumberMask** | Pointer to **string** | The last alphanumeric characters of an account&#39;s official account number. Note that the mask may be non-unique between accounts, and it may also not match the mask that the bank displays to the user.  | [optional] 
**InstitutionId** | Pointer to **string** | The ID of the institution external account belongs | [optional] 

## Methods

### NewExternalAccountVendorData

`func NewExternalAccountVendorData() *ExternalAccountVendorData`

NewExternalAccountVendorData instantiates a new ExternalAccountVendorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountVendorDataWithDefaults

`func NewExternalAccountVendorDataWithDefaults() *ExternalAccountVendorData`

NewExternalAccountVendorDataWithDefaults instantiates a new ExternalAccountVendorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumberMask

`func (o *ExternalAccountVendorData) GetAccountNumberMask() string`

GetAccountNumberMask returns the AccountNumberMask field if non-nil, zero value otherwise.

### GetAccountNumberMaskOk

`func (o *ExternalAccountVendorData) GetAccountNumberMaskOk() (*string, bool)`

GetAccountNumberMaskOk returns a tuple with the AccountNumberMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMask

`func (o *ExternalAccountVendorData) SetAccountNumberMask(v string)`

SetAccountNumberMask sets AccountNumberMask field to given value.

### HasAccountNumberMask

`func (o *ExternalAccountVendorData) HasAccountNumberMask() bool`

HasAccountNumberMask returns a boolean if a field has been set.

### GetInstitutionId

`func (o *ExternalAccountVendorData) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *ExternalAccountVendorData) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *ExternalAccountVendorData) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *ExternalAccountVendorData) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


