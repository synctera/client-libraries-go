# WaitlistEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Waitlist Name | [optional] 
**Increment** | Pointer to **int32** | Number of points earned per verified referral | [optional] 
**MaxProspects** | Pointer to **int32** | Max number of prospects allowed on this waitlist. Default is 10,000,000. | [optional] 
**WaitlistName** | Pointer to **string** | Waitlist Name | [optional] 

## Methods

### NewWaitlistEditable

`func NewWaitlistEditable() *WaitlistEditable`

NewWaitlistEditable instantiates a new WaitlistEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistEditableWithDefaults

`func NewWaitlistEditableWithDefaults() *WaitlistEditable`

NewWaitlistEditableWithDefaults instantiates a new WaitlistEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WaitlistEditable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WaitlistEditable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WaitlistEditable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WaitlistEditable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncrement

`func (o *WaitlistEditable) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *WaitlistEditable) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *WaitlistEditable) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *WaitlistEditable) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### GetMaxProspects

`func (o *WaitlistEditable) GetMaxProspects() int32`

GetMaxProspects returns the MaxProspects field if non-nil, zero value otherwise.

### GetMaxProspectsOk

`func (o *WaitlistEditable) GetMaxProspectsOk() (*int32, bool)`

GetMaxProspectsOk returns a tuple with the MaxProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProspects

`func (o *WaitlistEditable) SetMaxProspects(v int32)`

SetMaxProspects sets MaxProspects field to given value.

### HasMaxProspects

`func (o *WaitlistEditable) HasMaxProspects() bool`

HasMaxProspects returns a boolean if a field has been set.

### GetWaitlistName

`func (o *WaitlistEditable) GetWaitlistName() string`

GetWaitlistName returns the WaitlistName field if non-nil, zero value otherwise.

### GetWaitlistNameOk

`func (o *WaitlistEditable) GetWaitlistNameOk() (*string, bool)`

GetWaitlistNameOk returns a tuple with the WaitlistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistName

`func (o *WaitlistEditable) SetWaitlistName(v string)`

SetWaitlistName sets WaitlistName field to given value.

### HasWaitlistName

`func (o *WaitlistEditable) HasWaitlistName() bool`

HasWaitlistName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


