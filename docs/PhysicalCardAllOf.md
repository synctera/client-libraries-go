# PhysicalCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardImageId** | Pointer to **string** | The ID of the custom card image used for this card | [optional] 
**IsPinSet** | Pointer to **bool** | indicates whether a pin has been set on the card | [optional] [readonly] [default to false]

## Methods

### NewPhysicalCardAllOf

`func NewPhysicalCardAllOf() *PhysicalCardAllOf`

NewPhysicalCardAllOf instantiates a new PhysicalCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardAllOfWithDefaults

`func NewPhysicalCardAllOfWithDefaults() *PhysicalCardAllOf`

NewPhysicalCardAllOfWithDefaults instantiates a new PhysicalCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardImageId

`func (o *PhysicalCardAllOf) GetCardImageId() string`

GetCardImageId returns the CardImageId field if non-nil, zero value otherwise.

### GetCardImageIdOk

`func (o *PhysicalCardAllOf) GetCardImageIdOk() (*string, bool)`

GetCardImageIdOk returns a tuple with the CardImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardImageId

`func (o *PhysicalCardAllOf) SetCardImageId(v string)`

SetCardImageId sets CardImageId field to given value.

### HasCardImageId

`func (o *PhysicalCardAllOf) HasCardImageId() bool`

HasCardImageId returns a boolean if a field has been set.

### GetIsPinSet

`func (o *PhysicalCardAllOf) GetIsPinSet() bool`

GetIsPinSet returns the IsPinSet field if non-nil, zero value otherwise.

### GetIsPinSetOk

`func (o *PhysicalCardAllOf) GetIsPinSetOk() (*bool, bool)`

GetIsPinSetOk returns a tuple with the IsPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinSet

`func (o *PhysicalCardAllOf) SetIsPinSet(v bool)`

SetIsPinSet sets IsPinSet field to given value.

### HasIsPinSet

`func (o *PhysicalCardAllOf) HasIsPinSet() bool`

HasIsPinSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


