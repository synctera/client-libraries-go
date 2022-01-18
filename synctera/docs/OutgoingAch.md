# OutgoingAch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount to transfer in ISO 4217 minor currency units | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**CustomerId** | **string** | The customer&#39;s unique identifier | 
**DcSign** | **string** | The type of transaction (debit or credit). A debit is a transfer in and a credit is a transfer out of the originating account | 
**EffectiveDate** | Pointer to **string** | Effective date transaction proccesses (is_same_day needs to be false or not present at all) | [optional] 
**ExternalData** | Pointer to **map[string]interface{}** | Additional transfer metadata structured as key-value pairs | [optional] 
**FinalCustomerId** | Pointer to **string** | ID of the international customer that receives the final remittance transfer | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsSameDay** | Pointer to **bool** | Send as same day ACH transaction (use only is_same_day without specific effective_date) | [optional] 
**Memo** | Pointer to **string** | Memo for the payment | [optional] 
**OriginatingAccountId** | **string** | The unique identifier for an originating account | 
**ReceivingAccountId** | **string** | The unique identifier for an receiving account | 
**ReferenceInfo** | Pointer to **string** | Will be sent to the ACH network and maps to Addenda record 05 - the recipient bank will receive this info | [optional] 
**Risk** | [**RiskData**](RiskData.md) |  | 

## Methods

### NewOutgoingAch

`func NewOutgoingAch(amount int32, currency string, customerId string, dcSign string, originatingAccountId string, receivingAccountId string, risk RiskData, ) *OutgoingAch`

NewOutgoingAch instantiates a new OutgoingAch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchWithDefaults

`func NewOutgoingAchWithDefaults() *OutgoingAch`

NewOutgoingAchWithDefaults instantiates a new OutgoingAch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OutgoingAch) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OutgoingAch) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OutgoingAch) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *OutgoingAch) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OutgoingAch) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OutgoingAch) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *OutgoingAch) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OutgoingAch) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OutgoingAch) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *OutgoingAch) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *OutgoingAch) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *OutgoingAch) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetEffectiveDate

`func (o *OutgoingAch) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OutgoingAch) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OutgoingAch) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *OutgoingAch) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExternalData

`func (o *OutgoingAch) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *OutgoingAch) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *OutgoingAch) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *OutgoingAch) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetFinalCustomerId

`func (o *OutgoingAch) GetFinalCustomerId() string`

GetFinalCustomerId returns the FinalCustomerId field if non-nil, zero value otherwise.

### GetFinalCustomerIdOk

`func (o *OutgoingAch) GetFinalCustomerIdOk() (*string, bool)`

GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerId

`func (o *OutgoingAch) SetFinalCustomerId(v string)`

SetFinalCustomerId sets FinalCustomerId field to given value.

### HasFinalCustomerId

`func (o *OutgoingAch) HasFinalCustomerId() bool`

HasFinalCustomerId returns a boolean if a field has been set.

### GetId

`func (o *OutgoingAch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingAch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingAch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutgoingAch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSameDay

`func (o *OutgoingAch) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *OutgoingAch) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *OutgoingAch) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *OutgoingAch) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetMemo

`func (o *OutgoingAch) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *OutgoingAch) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *OutgoingAch) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *OutgoingAch) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *OutgoingAch) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *OutgoingAch) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *OutgoingAch) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetReceivingAccountId

`func (o *OutgoingAch) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *OutgoingAch) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *OutgoingAch) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetReferenceInfo

`func (o *OutgoingAch) GetReferenceInfo() string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *OutgoingAch) GetReferenceInfoOk() (*string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *OutgoingAch) SetReferenceInfo(v string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *OutgoingAch) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetRisk

`func (o *OutgoingAch) GetRisk() RiskData`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *OutgoingAch) GetRiskOk() (*RiskData, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *OutgoingAch) SetRisk(v RiskData)`

SetRisk sets Risk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


