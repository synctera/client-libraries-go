# ProspectEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Metadata** | Pointer to **string** | Client supplied json metadata to be stored with the prospect | [optional] 
**Source** | Pointer to **string** | Source of prospect | [optional] 
**Status** | Pointer to [**ProspectStatus**](ProspectStatus.md) |  | [optional] 

## Methods

### NewProspectEditable

`func NewProspectEditable() *ProspectEditable`

NewProspectEditable instantiates a new ProspectEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProspectEditableWithDefaults

`func NewProspectEditableWithDefaults() *ProspectEditable`

NewProspectEditableWithDefaults instantiates a new ProspectEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ProspectEditable) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ProspectEditable) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ProspectEditable) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ProspectEditable) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ProspectEditable) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ProspectEditable) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ProspectEditable) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ProspectEditable) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *ProspectEditable) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProspectEditable) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProspectEditable) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProspectEditable) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSource

`func (o *ProspectEditable) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProspectEditable) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProspectEditable) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProspectEditable) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ProspectEditable) GetStatus() ProspectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProspectEditable) GetStatusOk() (*ProspectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProspectEditable) SetStatus(v ProspectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProspectEditable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


