# Scan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number | [readonly] 
**BackImageId** | **string** | ID of the image of the back of the check | 
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | [readonly] 
**CheckCurrency** | **string** | ISO 4217 currency code for the check amount | [readonly] 
**DateScanned** | **string** | Date the check image was scanned, in RFC 3339 format | [readonly] 
**FrontImageId** | **string** | ID of the image of the front of the check | 
**Id** | **string** | RDC Scan ID | [readonly] 
**RoutingNumber** | **string** | Bank routing number | [readonly] 

## Methods

### NewScan

`func NewScan(accountNumber string, backImageId string, checkAmount int32, checkCurrency string, dateScanned string, frontImageId string, id string, routingNumber string, ) *Scan`

NewScan instantiates a new Scan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanWithDefaults

`func NewScanWithDefaults() *Scan`

NewScanWithDefaults instantiates a new Scan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Scan) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Scan) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Scan) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBackImageId

`func (o *Scan) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *Scan) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *Scan) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.


### GetCheckAmount

`func (o *Scan) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *Scan) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *Scan) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.


### GetCheckCurrency

`func (o *Scan) GetCheckCurrency() string`

GetCheckCurrency returns the CheckCurrency field if non-nil, zero value otherwise.

### GetCheckCurrencyOk

`func (o *Scan) GetCheckCurrencyOk() (*string, bool)`

GetCheckCurrencyOk returns a tuple with the CheckCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCurrency

`func (o *Scan) SetCheckCurrency(v string)`

SetCheckCurrency sets CheckCurrency field to given value.


### GetDateScanned

`func (o *Scan) GetDateScanned() string`

GetDateScanned returns the DateScanned field if non-nil, zero value otherwise.

### GetDateScannedOk

`func (o *Scan) GetDateScannedOk() (*string, bool)`

GetDateScannedOk returns a tuple with the DateScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateScanned

`func (o *Scan) SetDateScanned(v string)`

SetDateScanned sets DateScanned field to given value.


### GetFrontImageId

`func (o *Scan) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *Scan) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *Scan) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.


### GetId

`func (o *Scan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Scan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Scan) SetId(v string)`

SetId sets Id field to given value.


### GetRoutingNumber

`func (o *Scan) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *Scan) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *Scan) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


