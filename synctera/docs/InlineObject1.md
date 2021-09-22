# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceType** | Pointer to [**BalanceType**](BalanceType.md) |  | [optional] 
**PostingDate** | Pointer to **time.Time** | Posting date of the balance. Default is today&#39;s date | [optional] 

## Methods

### NewInlineObject1

`func NewInlineObject1() *InlineObject1`

NewInlineObject1 instantiates a new InlineObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1WithDefaults

`func NewInlineObject1WithDefaults() *InlineObject1`

NewInlineObject1WithDefaults instantiates a new InlineObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceType

`func (o *InlineObject1) GetBalanceType() BalanceType`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *InlineObject1) GetBalanceTypeOk() (*BalanceType, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *InlineObject1) SetBalanceType(v BalanceType)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *InlineObject1) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetPostingDate

`func (o *InlineObject1) GetPostingDate() time.Time`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *InlineObject1) GetPostingDateOk() (*time.Time, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *InlineObject1) SetPostingDate(v time.Time)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *InlineObject1) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


