# OutgoingAchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount to transfer in ISO 4217 minor currency units | 
**CompanyEntryDescription** | Pointer to **string** | Company entry description ACH field. Originator inserts this field&#39;s value to provide the Receiver with a description of the entry&#39;s purpose. | [optional] 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**CustomerId** | **string** | The customer&#39;s unique identifier | 
**DcSign** | **string** | The type of transaction (debit or credit). A debit is a transfer in and a credit is a transfer out of the originating account | 
**EffectiveDate** | Pointer to **string** | Effective date transaction proccesses (is_same_day needs to be false or not present at all) | [optional] 
**ExternalData** | Pointer to **map[string]interface{}** | Additional transfer metadata structured as key-value pairs | [optional] 
**FinalCustomerId** | Pointer to **string** | ID of the international customer that receives the final remittance transfer (required for OFAC enabled payments) | [optional] 
**Hold** | Pointer to [**AchRequestHoldData**](AchRequestHoldData.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsSameDay** | Pointer to **bool** | Send as same day ACH transaction (use only is_same_day without specific effective_date) | [optional] 
**Memo** | Pointer to **string** | Memo for the payment | [optional] 
**OriginatingAccountId** | **string** | The unique identifier for an originating account | 
**ReceivingAccountId** | **string** | The unique identifier for an receiving account | 
**ReferenceInfo** | Pointer to **string** | Will be sent to the ACH network and maps to Addenda record 05 - the recipient bank will receive this info | [optional] 
**Risk** | Pointer to [**RiskData**](RiskData.md) |  | [optional] 

## Methods

### NewOutgoingAchRequest

`func NewOutgoingAchRequest(amount int32, currency string, customerId string, dcSign string, originatingAccountId string, receivingAccountId string, ) *OutgoingAchRequest`

NewOutgoingAchRequest instantiates a new OutgoingAchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchRequestWithDefaults

`func NewOutgoingAchRequestWithDefaults() *OutgoingAchRequest`

NewOutgoingAchRequestWithDefaults instantiates a new OutgoingAchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OutgoingAchRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OutgoingAchRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OutgoingAchRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCompanyEntryDescription

`func (o *OutgoingAchRequest) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *OutgoingAchRequest) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *OutgoingAchRequest) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *OutgoingAchRequest) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *OutgoingAchRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OutgoingAchRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OutgoingAchRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *OutgoingAchRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OutgoingAchRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OutgoingAchRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *OutgoingAchRequest) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *OutgoingAchRequest) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *OutgoingAchRequest) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetEffectiveDate

`func (o *OutgoingAchRequest) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OutgoingAchRequest) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OutgoingAchRequest) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *OutgoingAchRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExternalData

`func (o *OutgoingAchRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *OutgoingAchRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *OutgoingAchRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *OutgoingAchRequest) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetFinalCustomerId

`func (o *OutgoingAchRequest) GetFinalCustomerId() string`

GetFinalCustomerId returns the FinalCustomerId field if non-nil, zero value otherwise.

### GetFinalCustomerIdOk

`func (o *OutgoingAchRequest) GetFinalCustomerIdOk() (*string, bool)`

GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerId

`func (o *OutgoingAchRequest) SetFinalCustomerId(v string)`

SetFinalCustomerId sets FinalCustomerId field to given value.

### HasFinalCustomerId

`func (o *OutgoingAchRequest) HasFinalCustomerId() bool`

HasFinalCustomerId returns a boolean if a field has been set.

### GetHold

`func (o *OutgoingAchRequest) GetHold() AchRequestHoldData`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *OutgoingAchRequest) GetHoldOk() (*AchRequestHoldData, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *OutgoingAchRequest) SetHold(v AchRequestHoldData)`

SetHold sets Hold field to given value.

### HasHold

`func (o *OutgoingAchRequest) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetId

`func (o *OutgoingAchRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingAchRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingAchRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutgoingAchRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSameDay

`func (o *OutgoingAchRequest) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *OutgoingAchRequest) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *OutgoingAchRequest) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *OutgoingAchRequest) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetMemo

`func (o *OutgoingAchRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *OutgoingAchRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *OutgoingAchRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *OutgoingAchRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *OutgoingAchRequest) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *OutgoingAchRequest) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *OutgoingAchRequest) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetReceivingAccountId

`func (o *OutgoingAchRequest) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *OutgoingAchRequest) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *OutgoingAchRequest) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetReferenceInfo

`func (o *OutgoingAchRequest) GetReferenceInfo() string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *OutgoingAchRequest) GetReferenceInfoOk() (*string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *OutgoingAchRequest) SetReferenceInfo(v string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *OutgoingAchRequest) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetRisk

`func (o *OutgoingAchRequest) GetRisk() RiskData`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *OutgoingAchRequest) GetRiskOk() (*RiskData, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *OutgoingAchRequest) SetRisk(v RiskData)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *OutgoingAchRequest) HasRisk() bool`

HasRisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


