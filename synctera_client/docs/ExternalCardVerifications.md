# ExternalCardVerifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerificationResults** | Pointer to **string** | Address verification results  Status | Description --- | --- VERIFIED | Address verified NOT_VERIFIED | Address not verified ADDRESS_NO_MATCH | Postal/ZIP match, street addresses do not match or street address not included in request  | [optional] 
**CardType** | Pointer to **string** | Indicates whether the external card is credit, debit, prepaid, deferred debit, or charge. | [optional] 
**Cvv2Result** | Pointer to **string** | Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration dates verified INCORRECT | Either CVV or expiration date is incorrect NOT_SUPPORTED |  Issuer does not participate in CVV2 service  | [optional] 
**FastFundsIndicator** | Pointer to **bool** | Indicates if card is Fast Funds eligible (i.e. if the funds will settle in 30 mins or less). If not eligible, typically funds will settle within 2 business days. | [optional] 
**OnlineGamblingBlockIndicator** | Pointer to **bool** | Indicates if the card can receive push-payments for online gambling payouts. | [optional] 
**Processor** | Pointer to [**Processor**](Processor.md) |  | [optional] 
**PushFundsBlockIndicator** | Pointer to **bool** | Indicates if the associated card can receive push-to-card disbursements. | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalCardVerifications

`func NewExternalCardVerifications() *ExternalCardVerifications`

NewExternalCardVerifications instantiates a new ExternalCardVerifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardVerificationsWithDefaults

`func NewExternalCardVerificationsWithDefaults() *ExternalCardVerifications`

NewExternalCardVerificationsWithDefaults instantiates a new ExternalCardVerifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerificationResults

`func (o *ExternalCardVerifications) GetAddressVerificationResults() string`

GetAddressVerificationResults returns the AddressVerificationResults field if non-nil, zero value otherwise.

### GetAddressVerificationResultsOk

`func (o *ExternalCardVerifications) GetAddressVerificationResultsOk() (*string, bool)`

GetAddressVerificationResultsOk returns a tuple with the AddressVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResults

`func (o *ExternalCardVerifications) SetAddressVerificationResults(v string)`

SetAddressVerificationResults sets AddressVerificationResults field to given value.

### HasAddressVerificationResults

`func (o *ExternalCardVerifications) HasAddressVerificationResults() bool`

HasAddressVerificationResults returns a boolean if a field has been set.

### GetCardType

`func (o *ExternalCardVerifications) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ExternalCardVerifications) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ExternalCardVerifications) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ExternalCardVerifications) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCvv2Result

`func (o *ExternalCardVerifications) GetCvv2Result() string`

GetCvv2Result returns the Cvv2Result field if non-nil, zero value otherwise.

### GetCvv2ResultOk

`func (o *ExternalCardVerifications) GetCvv2ResultOk() (*string, bool)`

GetCvv2ResultOk returns a tuple with the Cvv2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv2Result

`func (o *ExternalCardVerifications) SetCvv2Result(v string)`

SetCvv2Result sets Cvv2Result field to given value.

### HasCvv2Result

`func (o *ExternalCardVerifications) HasCvv2Result() bool`

HasCvv2Result returns a boolean if a field has been set.

### GetFastFundsIndicator

`func (o *ExternalCardVerifications) GetFastFundsIndicator() bool`

GetFastFundsIndicator returns the FastFundsIndicator field if non-nil, zero value otherwise.

### GetFastFundsIndicatorOk

`func (o *ExternalCardVerifications) GetFastFundsIndicatorOk() (*bool, bool)`

GetFastFundsIndicatorOk returns a tuple with the FastFundsIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastFundsIndicator

`func (o *ExternalCardVerifications) SetFastFundsIndicator(v bool)`

SetFastFundsIndicator sets FastFundsIndicator field to given value.

### HasFastFundsIndicator

`func (o *ExternalCardVerifications) HasFastFundsIndicator() bool`

HasFastFundsIndicator returns a boolean if a field has been set.

### GetOnlineGamblingBlockIndicator

`func (o *ExternalCardVerifications) GetOnlineGamblingBlockIndicator() bool`

GetOnlineGamblingBlockIndicator returns the OnlineGamblingBlockIndicator field if non-nil, zero value otherwise.

### GetOnlineGamblingBlockIndicatorOk

`func (o *ExternalCardVerifications) GetOnlineGamblingBlockIndicatorOk() (*bool, bool)`

GetOnlineGamblingBlockIndicatorOk returns a tuple with the OnlineGamblingBlockIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineGamblingBlockIndicator

`func (o *ExternalCardVerifications) SetOnlineGamblingBlockIndicator(v bool)`

SetOnlineGamblingBlockIndicator sets OnlineGamblingBlockIndicator field to given value.

### HasOnlineGamblingBlockIndicator

`func (o *ExternalCardVerifications) HasOnlineGamblingBlockIndicator() bool`

HasOnlineGamblingBlockIndicator returns a boolean if a field has been set.

### GetProcessor

`func (o *ExternalCardVerifications) GetProcessor() Processor`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ExternalCardVerifications) GetProcessorOk() (*Processor, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ExternalCardVerifications) SetProcessor(v Processor)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *ExternalCardVerifications) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetPushFundsBlockIndicator

`func (o *ExternalCardVerifications) GetPushFundsBlockIndicator() bool`

GetPushFundsBlockIndicator returns the PushFundsBlockIndicator field if non-nil, zero value otherwise.

### GetPushFundsBlockIndicatorOk

`func (o *ExternalCardVerifications) GetPushFundsBlockIndicatorOk() (*bool, bool)`

GetPushFundsBlockIndicatorOk returns a tuple with the PushFundsBlockIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushFundsBlockIndicator

`func (o *ExternalCardVerifications) SetPushFundsBlockIndicator(v bool)`

SetPushFundsBlockIndicator sets PushFundsBlockIndicator field to given value.

### HasPushFundsBlockIndicator

`func (o *ExternalCardVerifications) HasPushFundsBlockIndicator() bool`

HasPushFundsBlockIndicator returns a boolean if a field has been set.

### GetState

`func (o *ExternalCardVerifications) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExternalCardVerifications) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExternalCardVerifications) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExternalCardVerifications) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


