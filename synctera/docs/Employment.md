# Employment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerName** | **string** | Name of customer&#39;s employer | 
**EmploymentFrom** | Pointer to **time.Time** | First day of employment | [optional] 
**EmploymentHours** | Pointer to **float32** | Number of hours spent per week working for specified employment | [optional] 
**EmploymentIncome** | Pointer to **int32** | Annual income in cents | [optional] 
**EmploymentIncomeCurrency** | Pointer to **string** | The 3-letter alphabetic ISO 4217 code for the currency in which the employee was paid  | [optional] 
**EmploymentInfo** | Pointer to **map[string]interface{}** | A collection of arbitrary key-value pairs providing additional information about this employment relationship  | [optional] 
**EmploymentOccupation** | Pointer to **string** | Customer&#39;s work title, profession, or field | [optional] 
**EmploymentTo** | Pointer to **time.Time** | Last day of employment | [optional] 
**Id** | Pointer to **string** | Unique ID for this employment relationship | [optional] [readonly] 

## Methods

### NewEmployment

`func NewEmployment(employerName string, ) *Employment`

NewEmployment instantiates a new Employment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentWithDefaults

`func NewEmploymentWithDefaults() *Employment`

NewEmploymentWithDefaults instantiates a new Employment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerName

`func (o *Employment) GetEmployerName() string`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *Employment) GetEmployerNameOk() (*string, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *Employment) SetEmployerName(v string)`

SetEmployerName sets EmployerName field to given value.


### GetEmploymentFrom

`func (o *Employment) GetEmploymentFrom() time.Time`

GetEmploymentFrom returns the EmploymentFrom field if non-nil, zero value otherwise.

### GetEmploymentFromOk

`func (o *Employment) GetEmploymentFromOk() (*time.Time, bool)`

GetEmploymentFromOk returns a tuple with the EmploymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentFrom

`func (o *Employment) SetEmploymentFrom(v time.Time)`

SetEmploymentFrom sets EmploymentFrom field to given value.

### HasEmploymentFrom

`func (o *Employment) HasEmploymentFrom() bool`

HasEmploymentFrom returns a boolean if a field has been set.

### GetEmploymentHours

`func (o *Employment) GetEmploymentHours() float32`

GetEmploymentHours returns the EmploymentHours field if non-nil, zero value otherwise.

### GetEmploymentHoursOk

`func (o *Employment) GetEmploymentHoursOk() (*float32, bool)`

GetEmploymentHoursOk returns a tuple with the EmploymentHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentHours

`func (o *Employment) SetEmploymentHours(v float32)`

SetEmploymentHours sets EmploymentHours field to given value.

### HasEmploymentHours

`func (o *Employment) HasEmploymentHours() bool`

HasEmploymentHours returns a boolean if a field has been set.

### GetEmploymentIncome

`func (o *Employment) GetEmploymentIncome() int32`

GetEmploymentIncome returns the EmploymentIncome field if non-nil, zero value otherwise.

### GetEmploymentIncomeOk

`func (o *Employment) GetEmploymentIncomeOk() (*int32, bool)`

GetEmploymentIncomeOk returns a tuple with the EmploymentIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentIncome

`func (o *Employment) SetEmploymentIncome(v int32)`

SetEmploymentIncome sets EmploymentIncome field to given value.

### HasEmploymentIncome

`func (o *Employment) HasEmploymentIncome() bool`

HasEmploymentIncome returns a boolean if a field has been set.

### GetEmploymentIncomeCurrency

`func (o *Employment) GetEmploymentIncomeCurrency() string`

GetEmploymentIncomeCurrency returns the EmploymentIncomeCurrency field if non-nil, zero value otherwise.

### GetEmploymentIncomeCurrencyOk

`func (o *Employment) GetEmploymentIncomeCurrencyOk() (*string, bool)`

GetEmploymentIncomeCurrencyOk returns a tuple with the EmploymentIncomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentIncomeCurrency

`func (o *Employment) SetEmploymentIncomeCurrency(v string)`

SetEmploymentIncomeCurrency sets EmploymentIncomeCurrency field to given value.

### HasEmploymentIncomeCurrency

`func (o *Employment) HasEmploymentIncomeCurrency() bool`

HasEmploymentIncomeCurrency returns a boolean if a field has been set.

### GetEmploymentInfo

`func (o *Employment) GetEmploymentInfo() map[string]interface{}`

GetEmploymentInfo returns the EmploymentInfo field if non-nil, zero value otherwise.

### GetEmploymentInfoOk

`func (o *Employment) GetEmploymentInfoOk() (*map[string]interface{}, bool)`

GetEmploymentInfoOk returns a tuple with the EmploymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentInfo

`func (o *Employment) SetEmploymentInfo(v map[string]interface{})`

SetEmploymentInfo sets EmploymentInfo field to given value.

### HasEmploymentInfo

`func (o *Employment) HasEmploymentInfo() bool`

HasEmploymentInfo returns a boolean if a field has been set.

### GetEmploymentOccupation

`func (o *Employment) GetEmploymentOccupation() string`

GetEmploymentOccupation returns the EmploymentOccupation field if non-nil, zero value otherwise.

### GetEmploymentOccupationOk

`func (o *Employment) GetEmploymentOccupationOk() (*string, bool)`

GetEmploymentOccupationOk returns a tuple with the EmploymentOccupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentOccupation

`func (o *Employment) SetEmploymentOccupation(v string)`

SetEmploymentOccupation sets EmploymentOccupation field to given value.

### HasEmploymentOccupation

`func (o *Employment) HasEmploymentOccupation() bool`

HasEmploymentOccupation returns a boolean if a field has been set.

### GetEmploymentTo

`func (o *Employment) GetEmploymentTo() time.Time`

GetEmploymentTo returns the EmploymentTo field if non-nil, zero value otherwise.

### GetEmploymentToOk

`func (o *Employment) GetEmploymentToOk() (*time.Time, bool)`

GetEmploymentToOk returns a tuple with the EmploymentTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentTo

`func (o *Employment) SetEmploymentTo(v time.Time)`

SetEmploymentTo sets EmploymentTo field to given value.

### HasEmploymentTo

`func (o *Employment) HasEmploymentTo() bool`

HasEmploymentTo returns a boolean if a field has been set.

### GetId

`func (o *Employment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Employment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Employment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Employment) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


