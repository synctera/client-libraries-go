# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | 
**BackImageId** | **string** | ID of the uploaded image of the back of the check | 
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**DateCaptured** | **string** | Date the deposit was captured, in RFC 3339 format | [readonly] 
**DateProcessed** | **string** | Date the deposit was processed, in RFC 3339 format | [readonly] 
**DepositAmount** | **int32** | Amount deposited in ISO 4217 minor currency units | [readonly] 
**DepositCurrency** | **string** | ISO 4217 currency code for the deposit amount | 
**FrontImageId** | **string** | ID of the uploaded image of the front of the check | 
**Id** | **string** | RDC Deposit ID | [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**VendorInfo** | [**VendorInfo**](VendorInfo.md) |  | 

## Methods

### NewDeposit

`func NewDeposit(accountId string, backImageId string, checkAmount int32, dateCaptured string, dateProcessed string, depositAmount int32, depositCurrency string, frontImageId string, id string, vendorInfo VendorInfo, ) *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Deposit) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Deposit) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Deposit) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


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


### GetCreationTime

`func (o *Deposit) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Deposit) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Deposit) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Deposit) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDateCaptured

`func (o *Deposit) GetDateCaptured() string`

GetDateCaptured returns the DateCaptured field if non-nil, zero value otherwise.

### GetDateCapturedOk

`func (o *Deposit) GetDateCapturedOk() (*string, bool)`

GetDateCapturedOk returns a tuple with the DateCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCaptured

`func (o *Deposit) SetDateCaptured(v string)`

SetDateCaptured sets DateCaptured field to given value.


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


### GetLastUpdatedTime

`func (o *Deposit) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Deposit) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Deposit) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Deposit) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *Deposit) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Deposit) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Deposit) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Deposit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVendorInfo

`func (o *Deposit) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Deposit) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Deposit) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


