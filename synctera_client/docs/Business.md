# Business

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Ein** | Pointer to **string** | U.S. Employer Identification Number (EIN) for this business, in the format xx-xxxxxxx. | [optional] 
**Email** | Pointer to **string** | Business&#39;s email. | [optional] 
**EntityName** | Pointer to **string** | Business&#39;s legal name. | [optional] 
**FormationDate** | Pointer to **string** | Date the business was legally registered in RFC 3339 full-date format (YYYY-MM-DD). | [optional] 
**FormationState** | Pointer to **string** | U.S. state where the business is legally registered (2-letter abbreviation). | [optional] 
**Id** | Pointer to **string** | Business&#39;s unique identifier. | [optional] [readonly] 
**IsCustomer** | **bool** | True for personal and business customers with a direct relationship with the fintech or bank. | 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PhoneNumber** | Pointer to **string** | Business&#39;s phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] 
**Status** | **string** | Status of the business. One of the following: * &#x60;PROSPECT&#x60; – a potential customer, used for information-gathering and disclosures. * &#x60;ACTIVE&#x60; –  is an integrator defined status.  Integrators should set a business to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the business is eligible for specific actions such as initiating transactions or issuing a card. * &#x60;FROZEN&#x60; – business&#39;s actions are blocked for security, legal, or other reasons. * &#x60;SANCTION&#x60; – business is on a sanctions list and should be carefully monitored. * &#x60;DISSOLVED&#x60; – an inactive status indicating a business entity has filed articles of dissolution or a certificate of termination to terminate its existence. * &#x60;CANCELLED&#x60; – an inactive status indicating that a business entity has filed a cancellation or has failed to file its periodic report after notice of forfeiture of its rights to do business. * &#x60;SUSPENDED&#x60; – an inactive status indicating that the business entity has lost the right to operate in it&#39;s registered jurisdiction. * &#x60;MERGED&#x60; – an inactive status indicating that the business entity has terminated existence by merging into another entity. * &#x60;INACTIVE&#x60; – an inactive status indicating that the business entity is no longer active. * &#x60;CONVERTED&#x60; – An inactive status indicating that the business entity has been converted to another type of business entity in the same jurisdiction.  | 
**Structure** | Pointer to **string** | Business&#39;s legal structure. | [optional] 
**TradeNames** | Pointer to **[]string** | Other names by which this business is known. | [optional] 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYB verification was last run on the business. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**Website** | Pointer to **string** | Business&#39;s website. | [optional] 

## Methods

### NewBusiness

`func NewBusiness(isCustomer bool, status string, ) *Business`

NewBusiness instantiates a new Business object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessWithDefaults

`func NewBusinessWithDefaults() *Business`

NewBusinessWithDefaults instantiates a new Business object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Business) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Business) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Business) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Business) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEin

`func (o *Business) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *Business) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *Business) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *Business) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetEmail

`func (o *Business) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Business) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Business) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Business) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntityName

`func (o *Business) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *Business) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *Business) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *Business) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetFormationDate

`func (o *Business) GetFormationDate() string`

GetFormationDate returns the FormationDate field if non-nil, zero value otherwise.

### GetFormationDateOk

`func (o *Business) GetFormationDateOk() (*string, bool)`

GetFormationDateOk returns a tuple with the FormationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationDate

`func (o *Business) SetFormationDate(v string)`

SetFormationDate sets FormationDate field to given value.

### HasFormationDate

`func (o *Business) HasFormationDate() bool`

HasFormationDate returns a boolean if a field has been set.

### GetFormationState

`func (o *Business) GetFormationState() string`

GetFormationState returns the FormationState field if non-nil, zero value otherwise.

### GetFormationStateOk

`func (o *Business) GetFormationStateOk() (*string, bool)`

GetFormationStateOk returns a tuple with the FormationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationState

`func (o *Business) SetFormationState(v string)`

SetFormationState sets FormationState field to given value.

### HasFormationState

`func (o *Business) HasFormationState() bool`

HasFormationState returns a boolean if a field has been set.

### GetId

`func (o *Business) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Business) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Business) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Business) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *Business) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *Business) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *Business) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.


### GetLastUpdatedTime

`func (o *Business) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Business) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Business) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Business) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Business) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Business) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Business) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *Business) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *Business) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Business) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Business) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Business) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Business) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Business) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Business) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Business) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Business) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Business) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Business) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStructure

`func (o *Business) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *Business) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *Business) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *Business) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetTradeNames

`func (o *Business) GetTradeNames() []string`

GetTradeNames returns the TradeNames field if non-nil, zero value otherwise.

### GetTradeNamesOk

`func (o *Business) GetTradeNamesOk() (*[]string, bool)`

GetTradeNamesOk returns a tuple with the TradeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNames

`func (o *Business) SetTradeNames(v []string)`

SetTradeNames sets TradeNames field to given value.

### HasTradeNames

`func (o *Business) HasTradeNames() bool`

HasTradeNames returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *Business) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *Business) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *Business) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *Business) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *Business) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *Business) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *Business) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *Business) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetWebsite

`func (o *Business) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Business) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Business) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Business) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


