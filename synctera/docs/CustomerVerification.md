# CustomerVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerConsent** | **bool** | Whether this customer has consented to a KYC check.  | 
**CustomerIpAddress** | Pointer to **string** | IP address of the customer being verified. | [optional] 
**DocumentId** | Pointer to **string** | The ID of the uploaded government-issued identification document provided by the DV API endpoint.  | [optional] 
**VerificationType** | [**[]VerificationType**](VerificationType.md) | List of possible checks to run on a customer. | 

## Methods

### NewCustomerVerification

`func NewCustomerVerification(customerConsent bool, verificationType []VerificationType, ) *CustomerVerification`

NewCustomerVerification instantiates a new CustomerVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerificationWithDefaults

`func NewCustomerVerificationWithDefaults() *CustomerVerification`

NewCustomerVerificationWithDefaults instantiates a new CustomerVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerConsent

`func (o *CustomerVerification) GetCustomerConsent() bool`

GetCustomerConsent returns the CustomerConsent field if non-nil, zero value otherwise.

### GetCustomerConsentOk

`func (o *CustomerVerification) GetCustomerConsentOk() (*bool, bool)`

GetCustomerConsentOk returns a tuple with the CustomerConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerConsent

`func (o *CustomerVerification) SetCustomerConsent(v bool)`

SetCustomerConsent sets CustomerConsent field to given value.


### GetCustomerIpAddress

`func (o *CustomerVerification) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *CustomerVerification) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *CustomerVerification) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *CustomerVerification) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.

### GetDocumentId

`func (o *CustomerVerification) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CustomerVerification) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CustomerVerification) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CustomerVerification) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetVerificationType

`func (o *CustomerVerification) GetVerificationType() []VerificationType`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *CustomerVerification) GetVerificationTypeOk() (*[]VerificationType, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *CustomerVerification) SetVerificationType(v []VerificationType)`

SetVerificationType sets VerificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


