# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostingDate** | Pointer to **time.Time** | Posting date of the balance. Default is today&#39;s date | [optional] 
**BalanceType** | Pointer to [**BalanceType**](BalanceType.md) |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostingDate

`func (o *InlineObject) GetPostingDate() time.Time`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *InlineObject) GetPostingDateOk() (*time.Time, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *InlineObject) SetPostingDate(v time.Time)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *InlineObject) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### GetBalanceType

`func (o *InlineObject) GetBalanceType() BalanceType`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *InlineObject) GetBalanceTypeOk() (*BalanceType, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *InlineObject) SetBalanceType(v BalanceType)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *InlineObject) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


