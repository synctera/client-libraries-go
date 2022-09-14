# HoldModifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | **string** | The account number associated with the hold | 
**AllowPartial** | **bool** |  | 
**Amount** | **int64** | The amount of the hold. | 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewHoldModifyRequest

`func NewHoldModifyRequest(accountNo string, allowPartial bool, amount int64, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, riskInfo map[string]interface{}, subtype string, userData map[string]interface{}, ) *HoldModifyRequest`

NewHoldModifyRequest instantiates a new HoldModifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldModifyRequestWithDefaults

`func NewHoldModifyRequestWithDefaults() *HoldModifyRequest`

NewHoldModifyRequestWithDefaults instantiates a new HoldModifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *HoldModifyRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *HoldModifyRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *HoldModifyRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAllowPartial

`func (o *HoldModifyRequest) GetAllowPartial() bool`

GetAllowPartial returns the AllowPartial field if non-nil, zero value otherwise.

### GetAllowPartialOk

`func (o *HoldModifyRequest) GetAllowPartialOk() (*bool, bool)`

GetAllowPartialOk returns a tuple with the AllowPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartial

`func (o *HoldModifyRequest) SetAllowPartial(v bool)`

SetAllowPartial sets AllowPartial field to given value.


### GetAmount

`func (o *HoldModifyRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HoldModifyRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HoldModifyRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetEffectiveDate

`func (o *HoldModifyRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *HoldModifyRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *HoldModifyRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExpiresAt

`func (o *HoldModifyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HoldModifyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HoldModifyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExternalData

`func (o *HoldModifyRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *HoldModifyRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *HoldModifyRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *HoldModifyRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *HoldModifyRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetForcePost

`func (o *HoldModifyRequest) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *HoldModifyRequest) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *HoldModifyRequest) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetRiskInfo

`func (o *HoldModifyRequest) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *HoldModifyRequest) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *HoldModifyRequest) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *HoldModifyRequest) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *HoldModifyRequest) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetSubtype

`func (o *HoldModifyRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *HoldModifyRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *HoldModifyRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetUserData

`func (o *HoldModifyRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *HoldModifyRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *HoldModifyRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *HoldModifyRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *HoldModifyRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


