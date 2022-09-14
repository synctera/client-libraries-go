# HoldDeclineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | **string** | The account number associated with the hold | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**Reason** | **string** | The reason for the cancellation | 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewHoldDeclineRequest

`func NewHoldDeclineRequest(accountNo string, externalData map[string]interface{}, reason string, referenceId NullableString, riskInfo map[string]interface{}, userData map[string]interface{}, ) *HoldDeclineRequest`

NewHoldDeclineRequest instantiates a new HoldDeclineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldDeclineRequestWithDefaults

`func NewHoldDeclineRequestWithDefaults() *HoldDeclineRequest`

NewHoldDeclineRequestWithDefaults instantiates a new HoldDeclineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *HoldDeclineRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *HoldDeclineRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *HoldDeclineRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetExternalData

`func (o *HoldDeclineRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *HoldDeclineRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *HoldDeclineRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *HoldDeclineRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *HoldDeclineRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetReason

`func (o *HoldDeclineRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HoldDeclineRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HoldDeclineRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReferenceId

`func (o *HoldDeclineRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *HoldDeclineRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *HoldDeclineRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *HoldDeclineRequest) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *HoldDeclineRequest) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRiskInfo

`func (o *HoldDeclineRequest) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *HoldDeclineRequest) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *HoldDeclineRequest) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *HoldDeclineRequest) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *HoldDeclineRequest) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetUserData

`func (o *HoldDeclineRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *HoldDeclineRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *HoldDeclineRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *HoldDeclineRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *HoldDeclineRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


