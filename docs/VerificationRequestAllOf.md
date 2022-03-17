# VerificationRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CustomerConsent** | Pointer to **bool** | Whether this customer has consented to be verified. | [optional] 
**CustomerIpAddress** | Pointer to **string** | The customer&#39;s IP address. | [optional] 
**DocumentId** | Pointer to **string** | The ID of the uploaded government-issued identification document provided by the DV API endpoint.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 

## Methods

### NewVerificationRequestAllOf

`func NewVerificationRequestAllOf() *VerificationRequestAllOf`

NewVerificationRequestAllOf instantiates a new VerificationRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRequestAllOfWithDefaults

`func NewVerificationRequestAllOfWithDefaults() *VerificationRequestAllOf`

NewVerificationRequestAllOfWithDefaults instantiates a new VerificationRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *VerificationRequestAllOf) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *VerificationRequestAllOf) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *VerificationRequestAllOf) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *VerificationRequestAllOf) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerConsent

`func (o *VerificationRequestAllOf) GetCustomerConsent() bool`

GetCustomerConsent returns the CustomerConsent field if non-nil, zero value otherwise.

### GetCustomerConsentOk

`func (o *VerificationRequestAllOf) GetCustomerConsentOk() (*bool, bool)`

GetCustomerConsentOk returns a tuple with the CustomerConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerConsent

`func (o *VerificationRequestAllOf) SetCustomerConsent(v bool)`

SetCustomerConsent sets CustomerConsent field to given value.

### HasCustomerConsent

`func (o *VerificationRequestAllOf) HasCustomerConsent() bool`

HasCustomerConsent returns a boolean if a field has been set.

### GetCustomerIpAddress

`func (o *VerificationRequestAllOf) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *VerificationRequestAllOf) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *VerificationRequestAllOf) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *VerificationRequestAllOf) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.

### GetDocumentId

`func (o *VerificationRequestAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *VerificationRequestAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *VerificationRequestAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *VerificationRequestAllOf) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetPersonId

`func (o *VerificationRequestAllOf) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *VerificationRequestAllOf) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *VerificationRequestAllOf) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *VerificationRequestAllOf) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


