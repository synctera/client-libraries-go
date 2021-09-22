# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number | [readonly] 
**BackImageId** | **string** | ID of the image of the back of the check | [readonly] 
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | [readonly] 
**CheckCurrency** | **string** | ISO 4217 currency code for the check amount | [readonly] 
**DateProcessed** | **string** | Date the deposit was processed, in RFC 3339 format | [readonly] 
**DateSubmitted** | **string** | Date the deposit was submitted, in RFC 3339 format | [readonly] 
**DepositAmount** | **int32** | Amount deposited in ISO 4217 minor currency units | [readonly] 
**DepositCurrency** | **string** | ISO 4217 currency code for the deposit amount | [readonly] 
**FrontImageId** | **string** | ID of the image of the front of the check | [readonly] 
**Id** | **string** | RDC Deposit ID | [readonly] 
**RoutingNumber** | **string** | Bank routing number | [readonly] 
**ScanId** | **string** | ID of the OCR scan of the check image | 

## Methods

### NewDeposit

`func NewDeposit(accountNumber string, backImageId string, checkAmount int32, checkCurrency string, dateProcessed string, dateSubmitted string, depositAmount int32, depositCurrency string, frontImageId string, id string, routingNumber string, scanId string, ) *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Deposit) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Deposit) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Deposit) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBackImageId

`func (o *Deposit) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *Deposit) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *Deposit) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.


### GetCheckAmount

`func (o *Deposit) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *Deposit) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *Deposit) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.


### GetCheckCurrency

`func (o *Deposit) GetCheckCurrency() string`

GetCheckCurrency returns the CheckCurrency field if non-nil, zero value otherwise.

### GetCheckCurrencyOk

`func (o *Deposit) GetCheckCurrencyOk() (*string, bool)`

GetCheckCurrencyOk returns a tuple with the CheckCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCurrency

`func (o *Deposit) SetCheckCurrency(v string)`

SetCheckCurrency sets CheckCurrency field to given value.


### GetDateProcessed

`func (o *Deposit) GetDateProcessed() string`

GetDateProcessed returns the DateProcessed field if non-nil, zero value otherwise.

### GetDateProcessedOk

`func (o *Deposit) GetDateProcessedOk() (*string, bool)`

GetDateProcessedOk returns a tuple with the DateProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProcessed

`func (o *Deposit) SetDateProcessed(v string)`

SetDateProcessed sets DateProcessed field to given value.


### GetDateSubmitted

`func (o *Deposit) GetDateSubmitted() string`

GetDateSubmitted returns the DateSubmitted field if non-nil, zero value otherwise.

### GetDateSubmittedOk

`func (o *Deposit) GetDateSubmittedOk() (*string, bool)`

GetDateSubmittedOk returns a tuple with the DateSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSubmitted

`func (o *Deposit) SetDateSubmitted(v string)`

SetDateSubmitted sets DateSubmitted field to given value.


### GetDepositAmount

`func (o *Deposit) GetDepositAmount() int32`

GetDepositAmount returns the DepositAmount field if non-nil, zero value otherwise.

### GetDepositAmountOk

`func (o *Deposit) GetDepositAmountOk() (*int32, bool)`

GetDepositAmountOk returns a tuple with the DepositAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAmount

`func (o *Deposit) SetDepositAmount(v int32)`

SetDepositAmount sets DepositAmount field to given value.


### GetDepositCurrency

`func (o *Deposit) GetDepositCurrency() string`

GetDepositCurrency returns the DepositCurrency field if non-nil, zero value otherwise.

### GetDepositCurrencyOk

`func (o *Deposit) GetDepositCurrencyOk() (*string, bool)`

GetDepositCurrencyOk returns a tuple with the DepositCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCurrency

`func (o *Deposit) SetDepositCurrency(v string)`

SetDepositCurrency sets DepositCurrency field to given value.


### GetFrontImageId

`func (o *Deposit) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *Deposit) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *Deposit) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.


### GetId

`func (o *Deposit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deposit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deposit) SetId(v string)`

SetId sets Id field to given value.


### GetRoutingNumber

`func (o *Deposit) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *Deposit) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *Deposit) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetScanId

`func (o *Deposit) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *Deposit) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *Deposit) SetScanId(v string)`

SetScanId sets ScanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


