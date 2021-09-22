# RiskRatingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskRatings** | [**[]RiskRating**](RiskRating.md) | Array of customer risk ratings | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewRiskRatingList

`func NewRiskRatingList(riskRatings []RiskRating, ) *RiskRatingList`

NewRiskRatingList instantiates a new RiskRatingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRatingListWithDefaults

`func NewRiskRatingListWithDefaults() *RiskRatingList`

NewRiskRatingListWithDefaults instantiates a new RiskRatingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskRatings

`func (o *RiskRatingList) GetRiskRatings() []RiskRating`

GetRiskRatings returns the RiskRatings field if non-nil, zero value otherwise.

### GetRiskRatingsOk

`func (o *RiskRatingList) GetRiskRatingsOk() (*[]RiskRating, bool)`

GetRiskRatingsOk returns a tuple with the RiskRatings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRatings

`func (o *RiskRatingList) SetRiskRatings(v []RiskRating)`

SetRiskRatings sets RiskRatings field to given value.


### GetNextPageToken

`func (o *RiskRatingList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RiskRatingList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RiskRatingList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *RiskRatingList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


