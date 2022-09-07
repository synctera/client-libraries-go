# RiskRating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | Pointer to **string** | The risk configuration id used in risk score calculation | [optional] [readonly] 
**Id** | Pointer to **string** | Risk rating ID | [optional] [readonly] 
**NextReview** | Pointer to **time.Time** | The next review date where customer risk will be calculated | [optional] 
**RiskLevel** | Pointer to **string** | A textual representation of the customer risk score | [optional] 
**RiskReview** | Pointer to **time.Time** | The date the customer risk rating was calculated | [optional] [readonly] 
**RiskScore** | Pointer to **float32** | The cumulative score of all risk rating fields | [optional] 

## Methods

### NewRiskRating

`func NewRiskRating() *RiskRating`

NewRiskRating instantiates a new RiskRating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRatingWithDefaults

`func NewRiskRatingWithDefaults() *RiskRating`

NewRiskRatingWithDefaults instantiates a new RiskRating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *RiskRating) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *RiskRating) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *RiskRating) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *RiskRating) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetId

`func (o *RiskRating) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRating) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRating) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRating) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNextReview

`func (o *RiskRating) GetNextReview() time.Time`

GetNextReview returns the NextReview field if non-nil, zero value otherwise.

### GetNextReviewOk

`func (o *RiskRating) GetNextReviewOk() (*time.Time, bool)`

GetNextReviewOk returns a tuple with the NextReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReview

`func (o *RiskRating) SetNextReview(v time.Time)`

SetNextReview sets NextReview field to given value.

### HasNextReview

`func (o *RiskRating) HasNextReview() bool`

HasNextReview returns a boolean if a field has been set.

### GetRiskLevel

`func (o *RiskRating) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *RiskRating) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *RiskRating) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *RiskRating) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetRiskReview

`func (o *RiskRating) GetRiskReview() time.Time`

GetRiskReview returns the RiskReview field if non-nil, zero value otherwise.

### GetRiskReviewOk

`func (o *RiskRating) GetRiskReviewOk() (*time.Time, bool)`

GetRiskReviewOk returns a tuple with the RiskReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskReview

`func (o *RiskRating) SetRiskReview(v time.Time)`

SetRiskReview sets RiskReview field to given value.

### HasRiskReview

`func (o *RiskRating) HasRiskReview() bool`

HasRiskReview returns a boolean if a field has been set.

### GetRiskScore

`func (o *RiskRating) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RiskRating) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RiskRating) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *RiskRating) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


