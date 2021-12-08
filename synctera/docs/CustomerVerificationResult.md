# CustomerVerificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID for this verification result. | [optional] 
**Issues** | Pointer to **[]string** | List of potential problems found. These are subject to change.  | [optional] 
**RawResponse** | Pointer to [**RawResponse**](RawResponse.md) |  | [optional] 
**Result** | **string** | The determination of this verification. | 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**VerificationTime** | **time.Time** | The date and time the verification was completed. | 
**VerificationType** | [**VerificationType**](VerificationType.md) |  | 

## Methods

### NewCustomerVerificationResult

`func NewCustomerVerificationResult(result string, verificationTime time.Time, verificationType VerificationType, ) *CustomerVerificationResult`

NewCustomerVerificationResult instantiates a new CustomerVerificationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerVerificationResultWithDefaults

`func NewCustomerVerificationResultWithDefaults() *CustomerVerificationResult`

NewCustomerVerificationResultWithDefaults instantiates a new CustomerVerificationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerVerificationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerVerificationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerVerificationResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerVerificationResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssues

`func (o *CustomerVerificationResult) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CustomerVerificationResult) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CustomerVerificationResult) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CustomerVerificationResult) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRawResponse

`func (o *CustomerVerificationResult) GetRawResponse() RawResponse`

GetRawResponse returns the RawResponse field if non-nil, zero value otherwise.

### GetRawResponseOk

`func (o *CustomerVerificationResult) GetRawResponseOk() (*RawResponse, bool)`

GetRawResponseOk returns a tuple with the RawResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponse

`func (o *CustomerVerificationResult) SetRawResponse(v RawResponse)`

SetRawResponse sets RawResponse field to given value.

### HasRawResponse

`func (o *CustomerVerificationResult) HasRawResponse() bool`

HasRawResponse returns a boolean if a field has been set.

### GetResult

`func (o *CustomerVerificationResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CustomerVerificationResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CustomerVerificationResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetVendorInfo

`func (o *CustomerVerificationResult) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *CustomerVerificationResult) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *CustomerVerificationResult) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *CustomerVerificationResult) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerificationTime

`func (o *CustomerVerificationResult) GetVerificationTime() time.Time`

GetVerificationTime returns the VerificationTime field if non-nil, zero value otherwise.

### GetVerificationTimeOk

`func (o *CustomerVerificationResult) GetVerificationTimeOk() (*time.Time, bool)`

GetVerificationTimeOk returns a tuple with the VerificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTime

`func (o *CustomerVerificationResult) SetVerificationTime(v time.Time)`

SetVerificationTime sets VerificationTime field to given value.


### GetVerificationType

`func (o *CustomerVerificationResult) GetVerificationType() VerificationType`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *CustomerVerificationResult) GetVerificationTypeOk() (*VerificationType, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *CustomerVerificationResult) SetVerificationType(v VerificationType)`

SetVerificationType sets VerificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


