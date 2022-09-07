# Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**CareOfLine** | Pointer to **string** | The name of the person to send in care of | [optional] 
**IsExpeditedFulfillment** | Pointer to **bool** | Is the shipment expedited | [optional] [default to false]
**Method** | Pointer to **string** | The shipping method | [optional] [default to "LOCAL_MAIL"]
**PhoneNumber** | Pointer to **string** | The phone number of the recipient | [optional] 
**RecipientName** | Pointer to [**RecipientName**](RecipientName.md) |  | [optional] 

## Methods

### NewShipping

`func NewShipping() *Shipping`

NewShipping instantiates a new Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingWithDefaults

`func NewShippingWithDefaults() *Shipping`

NewShippingWithDefaults instantiates a new Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Shipping) GetAddress() Address1`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Shipping) GetAddressOk() (*Address1, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Shipping) SetAddress(v Address1)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Shipping) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCareOfLine

`func (o *Shipping) GetCareOfLine() string`

GetCareOfLine returns the CareOfLine field if non-nil, zero value otherwise.

### GetCareOfLineOk

`func (o *Shipping) GetCareOfLineOk() (*string, bool)`

GetCareOfLineOk returns a tuple with the CareOfLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOfLine

`func (o *Shipping) SetCareOfLine(v string)`

SetCareOfLine sets CareOfLine field to given value.

### HasCareOfLine

`func (o *Shipping) HasCareOfLine() bool`

HasCareOfLine returns a boolean if a field has been set.

### GetIsExpeditedFulfillment

`func (o *Shipping) GetIsExpeditedFulfillment() bool`

GetIsExpeditedFulfillment returns the IsExpeditedFulfillment field if non-nil, zero value otherwise.

### GetIsExpeditedFulfillmentOk

`func (o *Shipping) GetIsExpeditedFulfillmentOk() (*bool, bool)`

GetIsExpeditedFulfillmentOk returns a tuple with the IsExpeditedFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpeditedFulfillment

`func (o *Shipping) SetIsExpeditedFulfillment(v bool)`

SetIsExpeditedFulfillment sets IsExpeditedFulfillment field to given value.

### HasIsExpeditedFulfillment

`func (o *Shipping) HasIsExpeditedFulfillment() bool`

HasIsExpeditedFulfillment returns a boolean if a field has been set.

### GetMethod

`func (o *Shipping) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Shipping) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Shipping) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Shipping) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Shipping) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Shipping) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Shipping) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Shipping) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRecipientName

`func (o *Shipping) GetRecipientName() RecipientName`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *Shipping) GetRecipientNameOk() (*RecipientName, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *Shipping) SetRecipientName(v RecipientName)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *Shipping) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


