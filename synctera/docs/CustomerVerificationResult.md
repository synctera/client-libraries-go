# CustomerVerificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Verification object ID | [optional] 
**Issues** | Pointer to [**[]CustomerVerificationReasonCodes**](CustomerVerificationReasonCodes.md) | List of found issues | [optional] 
**VerificationsRun** | [**[]VerificationType**](VerificationType.md) | List of the verifications run | 
**RawResponse** | Pointer to [**RawResponse**](RawResponse.md) |  | [optional] 
**Result** | **string** | The determination of this KYC run | 
**VerificationDate** | **string** | The date on which the KYC run was completed | 

## Methods

### NewCustomerVerificationResult

`func NewCustomerVerificationResult(verificationsRun []VerificationType, result string, verificationDate string, ) *CustomerVerificationResult`

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

`func (o *CustomerVerificationResult) GetIssues() []CustomerVerificationReasonCodes`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CustomerVerificationResult) GetIssuesOk() (*[]CustomerVerificationReasonCodes, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CustomerVerificationResult) SetIssues(v []CustomerVerificationReasonCodes)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CustomerVerificationResult) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetVerificationsRun

`func (o *CustomerVerificationResult) GetVerificationsRun() []VerificationType`

GetVerificationsRun returns the VerificationsRun field if non-nil, zero value otherwise.

### GetVerificationsRunOk

`func (o *CustomerVerificationResult) GetVerificationsRunOk() (*[]VerificationType, bool)`

GetVerificationsRunOk returns a tuple with the VerificationsRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationsRun

`func (o *CustomerVerificationResult) SetVerificationsRun(v []VerificationType)`

SetVerificationsRun sets VerificationsRun field to given value.


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


### GetVerificationDate

`func (o *CustomerVerificationResult) GetVerificationDate() string`

GetVerificationDate returns the VerificationDate field if non-nil, zero value otherwise.

### GetVerificationDateOk

`func (o *CustomerVerificationResult) GetVerificationDateOk() (*string, bool)`

GetVerificationDateOk returns a tuple with the VerificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDate

`func (o *CustomerVerificationResult) SetVerificationDate(v string)`

SetVerificationDate sets VerificationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


