# PrefillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsnFilled** | Pointer to **bool** | If true, the person&#39;s SSN was successfully populated. | [optional] [readonly] 
**SsnLast4** | Pointer to **string** | Last four digits of person&#39;s Social Security number (SSN). | [optional] 

## Methods

### NewPrefillRequest

`func NewPrefillRequest() *PrefillRequest`

NewPrefillRequest instantiates a new PrefillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefillRequestWithDefaults

`func NewPrefillRequestWithDefaults() *PrefillRequest`

NewPrefillRequestWithDefaults instantiates a new PrefillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsnFilled

`func (o *PrefillRequest) GetSsnFilled() bool`

GetSsnFilled returns the SsnFilled field if non-nil, zero value otherwise.

### GetSsnFilledOk

`func (o *PrefillRequest) GetSsnFilledOk() (*bool, bool)`

GetSsnFilledOk returns a tuple with the SsnFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnFilled

`func (o *PrefillRequest) SetSsnFilled(v bool)`

SetSsnFilled sets SsnFilled field to given value.

### HasSsnFilled

`func (o *PrefillRequest) HasSsnFilled() bool`

HasSsnFilled returns a boolean if a field has been set.

### GetSsnLast4

`func (o *PrefillRequest) GetSsnLast4() string`

GetSsnLast4 returns the SsnLast4 field if non-nil, zero value otherwise.

### GetSsnLast4Ok

`func (o *PrefillRequest) GetSsnLast4Ok() (*string, bool)`

GetSsnLast4Ok returns a tuple with the SsnLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnLast4

`func (o *PrefillRequest) SetSsnLast4(v string)`

SetSsnLast4 sets SsnLast4 field to given value.

### HasSsnLast4

`func (o *PrefillRequest) HasSsnLast4() bool`

HasSsnLast4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


