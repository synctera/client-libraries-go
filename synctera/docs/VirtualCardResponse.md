# VirtualCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**AccountId** | **string** | The ID of the account to which the card will be linked | 
**Bin** | Pointer to **string** | The bin number | [optional] 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardProductId** | **string** | The card product to which the card is attached | 
**CreationTime** | **time.Time** | The timestamp representing when the card issuance request was made | [readonly] 
**CustomerId** | **string** | The ID of the customer to whom the card will be issued | 
**EmbossName** | [**EmbossName**](EmbossName.md) |  | 
**ExpirationMonth** | **string** |  | [readonly] 
**ExpirationTime** | Pointer to **time.Time** | The timestamp representing when the card would expire at | [optional] [readonly] 
**ExpirationYear** | **string** |  | [readonly] 
**Id** | **string** | Card ID | [readonly] 
**LastFour** | **string** | The last 4 digits of the card PAN | [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the card was last modified at | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**ReissueReason** | Pointer to **string** | The reason the card needs to be reissued | [optional] 
**ReissuedFromId** | Pointer to **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#39;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced.  | [optional] 
**ReissuedToId** | Pointer to **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [readonly] 
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**Type** | **string** | Indicates the type of card to be issued | 
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 

## Methods

### NewVirtualCardResponse

`func NewVirtualCardResponse(form string, accountId string, cardBrand CardBrand, cardProductId string, creationTime time.Time, customerId string, embossName EmbossName, expirationMonth string, expirationYear string, id string, lastFour string, type_ string, cardStatus CardStatus, statusReason CardStatusReasonCode, ) *VirtualCardResponse`

NewVirtualCardResponse instantiates a new VirtualCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCardResponseWithDefaults

`func NewVirtualCardResponseWithDefaults() *VirtualCardResponse`

NewVirtualCardResponseWithDefaults instantiates a new VirtualCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *VirtualCardResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *VirtualCardResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *VirtualCardResponse) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountId

`func (o *VirtualCardResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VirtualCardResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VirtualCardResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBin

`func (o *VirtualCardResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *VirtualCardResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *VirtualCardResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *VirtualCardResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardBrand

`func (o *VirtualCardResponse) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *VirtualCardResponse) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *VirtualCardResponse) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardProductId

`func (o *VirtualCardResponse) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *VirtualCardResponse) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *VirtualCardResponse) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCreationTime

`func (o *VirtualCardResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VirtualCardResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VirtualCardResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *VirtualCardResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *VirtualCardResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *VirtualCardResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEmbossName

`func (o *VirtualCardResponse) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *VirtualCardResponse) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *VirtualCardResponse) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.


### GetExpirationMonth

`func (o *VirtualCardResponse) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *VirtualCardResponse) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *VirtualCardResponse) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationTime

`func (o *VirtualCardResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VirtualCardResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VirtualCardResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VirtualCardResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpirationYear

`func (o *VirtualCardResponse) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *VirtualCardResponse) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *VirtualCardResponse) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetId

`func (o *VirtualCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastFour

`func (o *VirtualCardResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *VirtualCardResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *VirtualCardResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetLastModifiedTime

`func (o *VirtualCardResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *VirtualCardResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *VirtualCardResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *VirtualCardResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *VirtualCardResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualCardResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualCardResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VirtualCardResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReissueReason

`func (o *VirtualCardResponse) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *VirtualCardResponse) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *VirtualCardResponse) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *VirtualCardResponse) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *VirtualCardResponse) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *VirtualCardResponse) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *VirtualCardResponse) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *VirtualCardResponse) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissuedToId

`func (o *VirtualCardResponse) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *VirtualCardResponse) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *VirtualCardResponse) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *VirtualCardResponse) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetShipping

`func (o *VirtualCardResponse) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *VirtualCardResponse) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *VirtualCardResponse) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *VirtualCardResponse) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetType

`func (o *VirtualCardResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualCardResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualCardResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCardStatus

`func (o *VirtualCardResponse) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *VirtualCardResponse) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *VirtualCardResponse) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *VirtualCardResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *VirtualCardResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *VirtualCardResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *VirtualCardResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetStatusReason

`func (o *VirtualCardResponse) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *VirtualCardResponse) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *VirtualCardResponse) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


