# QuickstartT10Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | [**[]CardProductResponse**](CardProductResponse.md) | Array of card products created | 
**CardProgram** | [**CardProgramResponse**](CardProgramResponse.md) |  | 

## Methods

### NewQuickstartT10Response

`func NewQuickstartT10Response(cardProducts []CardProductResponse, cardProgram CardProgramResponse, ) *QuickstartT10Response`

NewQuickstartT10Response instantiates a new QuickstartT10Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickstartT10ResponseWithDefaults

`func NewQuickstartT10ResponseWithDefaults() *QuickstartT10Response`

NewQuickstartT10ResponseWithDefaults instantiates a new QuickstartT10Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *QuickstartT10Response) GetCardProducts() []CardProductResponse`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *QuickstartT10Response) GetCardProductsOk() (*[]CardProductResponse, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *QuickstartT10Response) SetCardProducts(v []CardProductResponse)`

SetCardProducts sets CardProducts field to given value.


### GetCardProgram

`func (o *QuickstartT10Response) GetCardProgram() CardProgramResponse`

GetCardProgram returns the CardProgram field if non-nil, zero value otherwise.

### GetCardProgramOk

`func (o *QuickstartT10Response) GetCardProgramOk() (*CardProgramResponse, bool)`

GetCardProgramOk returns a tuple with the CardProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgram

`func (o *QuickstartT10Response) SetCardProgram(v CardProgramResponse)`

SetCardProgram sets CardProgram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


