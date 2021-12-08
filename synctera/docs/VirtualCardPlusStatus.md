# VirtualCardPlusStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**AccountId** | Pointer to **string** | The ID of the account to which the card will be linked | [optional] 
**CardBrand** | Pointer to [**CardBrand**](CardBrand.md) |  | [optional] 
**CardProductId** | Pointer to **string** | The card product to which the card is attached | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the card issuance request was made | [optional] [readonly] 
**CustomerId** | Pointer to **string** | The ID of the customer to whom the card will be issued | [optional] 
**EmbossName** | Pointer to [**EmbossName**](EmbossName.md) |  | [optional] 
**ExpirationMonth** | Pointer to **string** |  | [optional] [readonly] 
**ExpirationTime** | Pointer to **time.Time** | The timestamp representing when the card would expire at | [optional] [readonly] 
**ExpirationYear** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Card ID | [optional] [readonly] 
**LastFour** | Pointer to **string** | The last 4 digits of the card PAN | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the card was last modified at | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**ReissueReason** | Pointer to **string** | The reason the card needs to be reissued | [optional] 
**ReissuedFromId** | Pointer to **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#39;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced.  | [optional] 
**ReissuedToId** | Pointer to **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [readonly] 
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**Type** | Pointer to **string** | Indicates the type of card to be issued | [optional] 
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 

## Methods

### NewVirtualCardPlusStatus

`func NewVirtualCardPlusStatus(form string, cardStatus CardStatus, statusReason CardStatusReasonCode, ) *VirtualCardPlusStatus`

NewVirtualCardPlusStatus instantiates a new VirtualCardPlusStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCardPlusStatusWithDefaults

`func NewVirtualCardPlusStatusWithDefaults() *VirtualCardPlusStatus`

NewVirtualCardPlusStatusWithDefaults instantiates a new VirtualCardPlusStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *VirtualCardPlusStatus) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *VirtualCardPlusStatus) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *VirtualCardPlusStatus) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountId

`func (o *VirtualCardPlusStatus) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VirtualCardPlusStatus) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VirtualCardPlusStatus) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VirtualCardPlusStatus) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCardBrand

`func (o *VirtualCardPlusStatus) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *VirtualCardPlusStatus) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *VirtualCardPlusStatus) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *VirtualCardPlusStatus) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardProductId

`func (o *VirtualCardPlusStatus) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *VirtualCardPlusStatus) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *VirtualCardPlusStatus) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.

### HasCardProductId

`func (o *VirtualCardPlusStatus) HasCardProductId() bool`

HasCardProductId returns a boolean if a field has been set.

### GetCreationTime

`func (o *VirtualCardPlusStatus) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VirtualCardPlusStatus) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VirtualCardPlusStatus) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VirtualCardPlusStatus) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCustomerId

`func (o *VirtualCardPlusStatus) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *VirtualCardPlusStatus) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *VirtualCardPlusStatus) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *VirtualCardPlusStatus) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmbossName

`func (o *VirtualCardPlusStatus) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *VirtualCardPlusStatus) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *VirtualCardPlusStatus) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.

### HasEmbossName

`func (o *VirtualCardPlusStatus) HasEmbossName() bool`

HasEmbossName returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *VirtualCardPlusStatus) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *VirtualCardPlusStatus) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *VirtualCardPlusStatus) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *VirtualCardPlusStatus) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationTime

`func (o *VirtualCardPlusStatus) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VirtualCardPlusStatus) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VirtualCardPlusStatus) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VirtualCardPlusStatus) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpirationYear

`func (o *VirtualCardPlusStatus) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *VirtualCardPlusStatus) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *VirtualCardPlusStatus) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *VirtualCardPlusStatus) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetId

`func (o *VirtualCardPlusStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCardPlusStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCardPlusStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCardPlusStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastFour

`func (o *VirtualCardPlusStatus) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *VirtualCardPlusStatus) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *VirtualCardPlusStatus) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *VirtualCardPlusStatus) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *VirtualCardPlusStatus) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *VirtualCardPlusStatus) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *VirtualCardPlusStatus) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *VirtualCardPlusStatus) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *VirtualCardPlusStatus) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualCardPlusStatus) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualCardPlusStatus) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VirtualCardPlusStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReissueReason

`func (o *VirtualCardPlusStatus) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *VirtualCardPlusStatus) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *VirtualCardPlusStatus) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *VirtualCardPlusStatus) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *VirtualCardPlusStatus) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *VirtualCardPlusStatus) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *VirtualCardPlusStatus) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *VirtualCardPlusStatus) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissuedToId

`func (o *VirtualCardPlusStatus) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *VirtualCardPlusStatus) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *VirtualCardPlusStatus) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *VirtualCardPlusStatus) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetShipping

`func (o *VirtualCardPlusStatus) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *VirtualCardPlusStatus) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *VirtualCardPlusStatus) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *VirtualCardPlusStatus) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetType

`func (o *VirtualCardPlusStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualCardPlusStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualCardPlusStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualCardPlusStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCardStatus

`func (o *VirtualCardPlusStatus) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *VirtualCardPlusStatus) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *VirtualCardPlusStatus) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *VirtualCardPlusStatus) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *VirtualCardPlusStatus) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *VirtualCardPlusStatus) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *VirtualCardPlusStatus) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetStatusReason

`func (o *VirtualCardPlusStatus) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *VirtualCardPlusStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *VirtualCardPlusStatus) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


