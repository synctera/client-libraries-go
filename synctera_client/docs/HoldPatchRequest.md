# HoldPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPostAt** | **time.Time** | The time the transaction will be automatically posted. | 
**EnhancedTransaction** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewHoldPatchRequest

`func NewHoldPatchRequest(autoPostAt time.Time, enhancedTransaction map[string]interface{}, expiresAt time.Time, externalData map[string]interface{}, referenceId NullableString, riskInfo map[string]interface{}, userData map[string]interface{}, ) *HoldPatchRequest`

NewHoldPatchRequest instantiates a new HoldPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldPatchRequestWithDefaults

`func NewHoldPatchRequestWithDefaults() *HoldPatchRequest`

NewHoldPatchRequestWithDefaults instantiates a new HoldPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPostAt

`func (o *HoldPatchRequest) GetAutoPostAt() time.Time`

GetAutoPostAt returns the AutoPostAt field if non-nil, zero value otherwise.

### GetAutoPostAtOk

`func (o *HoldPatchRequest) GetAutoPostAtOk() (*time.Time, bool)`

GetAutoPostAtOk returns a tuple with the AutoPostAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPostAt

`func (o *HoldPatchRequest) SetAutoPostAt(v time.Time)`

SetAutoPostAt sets AutoPostAt field to given value.


### GetEnhancedTransaction

`func (o *HoldPatchRequest) GetEnhancedTransaction() map[string]interface{}`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *HoldPatchRequest) GetEnhancedTransactionOk() (*map[string]interface{}, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *HoldPatchRequest) SetEnhancedTransaction(v map[string]interface{})`

SetEnhancedTransaction sets EnhancedTransaction field to given value.


### SetEnhancedTransactionNil

`func (o *HoldPatchRequest) SetEnhancedTransactionNil(b bool)`

 SetEnhancedTransactionNil sets the value for EnhancedTransaction to be an explicit nil

### UnsetEnhancedTransaction
`func (o *HoldPatchRequest) UnsetEnhancedTransaction()`

UnsetEnhancedTransaction ensures that no value is present for EnhancedTransaction, not even an explicit nil
### GetExpiresAt

`func (o *HoldPatchRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HoldPatchRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HoldPatchRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExternalData

`func (o *HoldPatchRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *HoldPatchRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *HoldPatchRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *HoldPatchRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *HoldPatchRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetReferenceId

`func (o *HoldPatchRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *HoldPatchRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *HoldPatchRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *HoldPatchRequest) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *HoldPatchRequest) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRiskInfo

`func (o *HoldPatchRequest) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *HoldPatchRequest) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *HoldPatchRequest) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *HoldPatchRequest) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *HoldPatchRequest) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetUserData

`func (o *HoldPatchRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *HoldPatchRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *HoldPatchRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *HoldPatchRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *HoldPatchRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


