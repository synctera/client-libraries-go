# AchInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**OutgoingAchRequest**](OutgoingAchRequest.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAchInstruction

`func NewAchInstruction(request OutgoingAchRequest, type_ string, ) *AchInstruction`

NewAchInstruction instantiates a new AchInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchInstructionWithDefaults

`func NewAchInstructionWithDefaults() *AchInstruction`

NewAchInstructionWithDefaults instantiates a new AchInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *AchInstruction) GetRequest() OutgoingAchRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AchInstruction) GetRequestOk() (*OutgoingAchRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AchInstruction) SetRequest(v OutgoingAchRequest)`

SetRequest sets Request field to given value.


### GetType

`func (o *AchInstruction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AchInstruction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AchInstruction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


