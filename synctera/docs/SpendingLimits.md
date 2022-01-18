# SpendingLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to [**SpendingLimitWithTime**](SpendingLimitWithTime.md) |  | [optional] 
**Lifetime** | Pointer to [**SpendingLimitWithTime**](SpendingLimitWithTime.md) |  | [optional] 
**Month** | Pointer to [**SpendingLimitWithTime**](SpendingLimitWithTime.md) |  | [optional] 
**Transaction** | Pointer to [**SpendingLimitsTransaction**](SpendingLimitsTransaction.md) |  | [optional] 
**Week** | Pointer to [**SpendingLimitWithTime**](SpendingLimitWithTime.md) |  | [optional] 

## Methods

### NewSpendingLimits

`func NewSpendingLimits() *SpendingLimits`

NewSpendingLimits instantiates a new SpendingLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingLimitsWithDefaults

`func NewSpendingLimitsWithDefaults() *SpendingLimits`

NewSpendingLimitsWithDefaults instantiates a new SpendingLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *SpendingLimits) GetDay() SpendingLimitWithTime`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *SpendingLimits) GetDayOk() (*SpendingLimitWithTime, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *SpendingLimits) SetDay(v SpendingLimitWithTime)`

SetDay sets Day field to given value.

### HasDay

`func (o *SpendingLimits) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetLifetime

`func (o *SpendingLimits) GetLifetime() SpendingLimitWithTime`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *SpendingLimits) GetLifetimeOk() (*SpendingLimitWithTime, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *SpendingLimits) SetLifetime(v SpendingLimitWithTime)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *SpendingLimits) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetMonth

`func (o *SpendingLimits) GetMonth() SpendingLimitWithTime`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *SpendingLimits) GetMonthOk() (*SpendingLimitWithTime, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *SpendingLimits) SetMonth(v SpendingLimitWithTime)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *SpendingLimits) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetTransaction

`func (o *SpendingLimits) GetTransaction() SpendingLimitsTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *SpendingLimits) GetTransactionOk() (*SpendingLimitsTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *SpendingLimits) SetTransaction(v SpendingLimitsTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *SpendingLimits) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetWeek

`func (o *SpendingLimits) GetWeek() SpendingLimitWithTime`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *SpendingLimits) GetWeekOk() (*SpendingLimitWithTime, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *SpendingLimits) SetWeek(v SpendingLimitWithTime)`

SetWeek sets Week field to given value.

### HasWeek

`func (o *SpendingLimits) HasWeek() bool`

HasWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


