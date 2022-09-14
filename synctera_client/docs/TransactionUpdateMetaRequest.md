# TransactionUpdateMetaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewTransactionUpdateMetaRequest

`func NewTransactionUpdateMetaRequest(externalData map[string]interface{}, userData map[string]interface{}, ) *TransactionUpdateMetaRequest`

NewTransactionUpdateMetaRequest instantiates a new TransactionUpdateMetaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUpdateMetaRequestWithDefaults

`func NewTransactionUpdateMetaRequestWithDefaults() *TransactionUpdateMetaRequest`

NewTransactionUpdateMetaRequestWithDefaults instantiates a new TransactionUpdateMetaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalData

`func (o *TransactionUpdateMetaRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *TransactionUpdateMetaRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *TransactionUpdateMetaRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *TransactionUpdateMetaRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *TransactionUpdateMetaRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetUserData

`func (o *TransactionUpdateMetaRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TransactionUpdateMetaRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TransactionUpdateMetaRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *TransactionUpdateMetaRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *TransactionUpdateMetaRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


