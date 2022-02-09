# PhysicalCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**Id** | **string** | Card ID | [readonly] 
**Bin** | Pointer to **string** | The bin number | [optional] 
**CustomerId** | **string** | The ID of the customer to whom the card will be issued | 
**AccountId** | **string** | The ID of the account to which the card will be linked | 
**Type** | **string** | Indicates the type of card to be issued | 
**EmbossName** | [**EmbossName**](EmbossName.md) |  | 
**LastFour** | **string** | The last 4 digits of the card PAN | [readonly] 
**CardProductId** | **string** | The card product to which the card is attached | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**ExpirationYear** | **string** |  | [readonly] 
**ExpirationMonth** | **string** |  | [readonly] 
**ExpirationTime** | Pointer to **time.Time** | The timestamp representing when the card would expire at | [optional] [readonly] 
**CreationTime** | **time.Time** | The timestamp representing when the card issuance request was made | [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the card was last modified at | [optional] [readonly] 
**ReissuedToId** | Pointer to **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [readonly] 
**ReissuedFromId** | Pointer to **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#39;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced.  | [optional] 
**ReissueReason** | Pointer to **string** | The reason the card needs to be reissued | [optional] 
**Shipping** | [**Shipping**](Shipping.md) |  | 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**IsPinSet** | Pointer to **bool** | indicates whether a pin has been set on the card | [optional] [readonly] [default to false]
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**CardFulfillmentStatus** | [**CardFulfillmentStatus**](CardFulfillmentStatus.md) |  | 
**TrackingNumber** | Pointer to **string** | The tracking number | [optional] [readonly] 
**PhysicalCardFormat** | [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | 

## Methods

### NewPhysicalCardResponse

`func NewPhysicalCardResponse(form string, id string, customerId string, accountId string, type_ string, embossName EmbossName, lastFour string, cardProductId string, cardBrand CardBrand, expirationYear string, expirationMonth string, creationTime time.Time, shipping Shipping, cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, physicalCardFormat PhysicalCardFormat, ) *PhysicalCardResponse`

NewPhysicalCardResponse instantiates a new PhysicalCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardResponseWithDefaults

`func NewPhysicalCardResponseWithDefaults() *PhysicalCardResponse`

NewPhysicalCardResponseWithDefaults instantiates a new PhysicalCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *PhysicalCardResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PhysicalCardResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PhysicalCardResponse) SetForm(v string)`

SetForm sets Form field to given value.


### GetId

`func (o *PhysicalCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBin

`func (o *PhysicalCardResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PhysicalCardResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PhysicalCardResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *PhysicalCardResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCustomerId

`func (o *PhysicalCardResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PhysicalCardResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PhysicalCardResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetAccountId

`func (o *PhysicalCardResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PhysicalCardResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PhysicalCardResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *PhysicalCardResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhysicalCardResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhysicalCardResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEmbossName

`func (o *PhysicalCardResponse) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *PhysicalCardResponse) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *PhysicalCardResponse) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.


### GetLastFour

`func (o *PhysicalCardResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *PhysicalCardResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *PhysicalCardResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetCardProductId

`func (o *PhysicalCardResponse) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *PhysicalCardResponse) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *PhysicalCardResponse) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCardBrand

`func (o *PhysicalCardResponse) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PhysicalCardResponse) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PhysicalCardResponse) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetExpirationYear

`func (o *PhysicalCardResponse) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *PhysicalCardResponse) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *PhysicalCardResponse) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetExpirationMonth

`func (o *PhysicalCardResponse) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *PhysicalCardResponse) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *PhysicalCardResponse) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationTime

`func (o *PhysicalCardResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *PhysicalCardResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *PhysicalCardResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *PhysicalCardResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *PhysicalCardResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PhysicalCardResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PhysicalCardResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetLastModifiedTime

`func (o *PhysicalCardResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *PhysicalCardResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *PhysicalCardResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *PhysicalCardResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetReissuedToId

`func (o *PhysicalCardResponse) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *PhysicalCardResponse) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *PhysicalCardResponse) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *PhysicalCardResponse) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *PhysicalCardResponse) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *PhysicalCardResponse) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *PhysicalCardResponse) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *PhysicalCardResponse) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissueReason

`func (o *PhysicalCardResponse) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *PhysicalCardResponse) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *PhysicalCardResponse) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *PhysicalCardResponse) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetShipping

`func (o *PhysicalCardResponse) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *PhysicalCardResponse) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *PhysicalCardResponse) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.


### GetMetadata

`func (o *PhysicalCardResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PhysicalCardResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PhysicalCardResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PhysicalCardResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIsPinSet

`func (o *PhysicalCardResponse) GetIsPinSet() bool`

GetIsPinSet returns the IsPinSet field if non-nil, zero value otherwise.

### GetIsPinSetOk

`func (o *PhysicalCardResponse) GetIsPinSetOk() (*bool, bool)`

GetIsPinSetOk returns a tuple with the IsPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinSet

`func (o *PhysicalCardResponse) SetIsPinSet(v bool)`

SetIsPinSet sets IsPinSet field to given value.

### HasIsPinSet

`func (o *PhysicalCardResponse) HasIsPinSet() bool`

HasIsPinSet returns a boolean if a field has been set.

### GetCardStatus

`func (o *PhysicalCardResponse) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *PhysicalCardResponse) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *PhysicalCardResponse) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetStatusReason

`func (o *PhysicalCardResponse) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PhysicalCardResponse) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PhysicalCardResponse) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.


### GetMemo

`func (o *PhysicalCardResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PhysicalCardResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PhysicalCardResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PhysicalCardResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCardFulfillmentStatus

`func (o *PhysicalCardResponse) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *PhysicalCardResponse) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *PhysicalCardResponse) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.


### GetTrackingNumber

`func (o *PhysicalCardResponse) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *PhysicalCardResponse) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *PhysicalCardResponse) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *PhysicalCardResponse) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetPhysicalCardFormat

`func (o *PhysicalCardResponse) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *PhysicalCardResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *PhysicalCardResponse) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


