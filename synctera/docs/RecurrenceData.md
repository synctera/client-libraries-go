# RecurrenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Recurrence ID | [optional] 
**RecurrenceFrequency** | Pointer to **int32** | the frequency at which to recur | [optional] 

## Methods

### NewRecurrenceData

`func NewRecurrenceData() *RecurrenceData`

NewRecurrenceData instantiates a new RecurrenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceDataWithDefaults

`func NewRecurrenceDataWithDefaults() *RecurrenceData`

NewRecurrenceDataWithDefaults instantiates a new RecurrenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurrenceData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurrenceData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurrenceData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecurrenceData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecurrenceFrequency

`func (o *RecurrenceData) GetRecurrenceFrequency() int32`

GetRecurrenceFrequency returns the RecurrenceFrequency field if non-nil, zero value otherwise.

### GetRecurrenceFrequencyOk

`func (o *RecurrenceData) GetRecurrenceFrequencyOk() (*int32, bool)`

GetRecurrenceFrequencyOk returns a tuple with the RecurrenceFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceFrequency

`func (o *RecurrenceData) SetRecurrenceFrequency(v int32)`

SetRecurrenceFrequency sets RecurrenceFrequency field to given value.

### HasRecurrenceFrequency

`func (o *RecurrenceData) HasRecurrenceFrequency() bool`

HasRecurrenceFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


