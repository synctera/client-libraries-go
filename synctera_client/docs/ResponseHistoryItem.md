# ResponseHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Response code from the request | [optional] 
**ResponseBody** | Pointer to **string** | Response body from the request(Length more than 1024 will be trimmed) | [optional] 
**ResponseTime** | Pointer to **time.Time** | Timestamp that the response is received | [optional] 
**SentTime** | Pointer to **time.Time** | Timestamp that the request is sent | [optional] 

## Methods

### NewResponseHistoryItem

`func NewResponseHistoryItem() *ResponseHistoryItem`

NewResponseHistoryItem instantiates a new ResponseHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseHistoryItemWithDefaults

`func NewResponseHistoryItemWithDefaults() *ResponseHistoryItem`

NewResponseHistoryItemWithDefaults instantiates a new ResponseHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponseHistoryItem) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseHistoryItem) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseHistoryItem) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponseHistoryItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *ResponseHistoryItem) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ResponseHistoryItem) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ResponseHistoryItem) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *ResponseHistoryItem) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseTime

`func (o *ResponseHistoryItem) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *ResponseHistoryItem) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *ResponseHistoryItem) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *ResponseHistoryItem) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSentTime

`func (o *ResponseHistoryItem) GetSentTime() time.Time`

GetSentTime returns the SentTime field if non-nil, zero value otherwise.

### GetSentTimeOk

`func (o *ResponseHistoryItem) GetSentTimeOk() (*time.Time, bool)`

GetSentTimeOk returns a tuple with the SentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTime

`func (o *ResponseHistoryItem) SetSentTime(v time.Time)`

SetSentTime sets SentTime field to given value.

### HasSentTime

`func (o *ResponseHistoryItem) HasSentTime() bool`

HasSentTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


