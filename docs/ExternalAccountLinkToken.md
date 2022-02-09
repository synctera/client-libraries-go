# ExternalAccountLinkToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The identifier for the customer associated with this account | 
**VendorInstitutionId** | Pointer to **string** | The ID of the institution the access token is requested for. If present the link token will be created in an update mode.  | [optional] 
**Type** | **string** | The type of the link token. DEPOSITORY for checking and savings accounts, CREDIT for credit card type accounts. | 
**ClientName** | **string** | The name of your application, as it should be displayed in Link. Maximum length of 30 characters. | 
**Language** | **string** | The language that corresponds to the link token. For Plaid, see their [documentation](https://plaid.com/docs/api/tokens/#link-token-create-request-language) for a list of allowed values.  | 
**CountryCodes** | **[]string** | Country codes in the ISO-3166-1 alpha-2 country code standard. | 
**LinkCustomizationName** | Pointer to **string** | The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the default customization will be used. When using a Link customization, the language in the customization must match the language selected via the language parameter, and the countries in the customization should match the country codes selected via country_codes.  | [optional] 
**RedirectUri** | Pointer to **string** | A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview.  | [optional] 
**LinkToken** | Pointer to **string** | A link_token, which can be supplied to Link in order to initialize it and receive a public_token, which can be exchanged for an access_token.  | [optional] [readonly] 
**VendorAccessToken** | Pointer to **string** | The access token associated with the Item data is being requested for. | [optional] 
**Expiration** | Pointer to **time.Time** | The expiration date for the link_token. Expires in 4 hours. | [optional] [readonly] 
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. | [optional] [readonly] 
**VerifyOwner** | Pointer to **bool** | Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification will be suppressed by default  | [optional] [default to false]

## Methods

### NewExternalAccountLinkToken

`func NewExternalAccountLinkToken(customerId string, type_ string, clientName string, language string, countryCodes []string, ) *ExternalAccountLinkToken`

NewExternalAccountLinkToken instantiates a new ExternalAccountLinkToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountLinkTokenWithDefaults

`func NewExternalAccountLinkTokenWithDefaults() *ExternalAccountLinkToken`

NewExternalAccountLinkTokenWithDefaults instantiates a new ExternalAccountLinkToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ExternalAccountLinkToken) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExternalAccountLinkToken) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExternalAccountLinkToken) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetVendorInstitutionId

`func (o *ExternalAccountLinkToken) GetVendorInstitutionId() string`

GetVendorInstitutionId returns the VendorInstitutionId field if non-nil, zero value otherwise.

### GetVendorInstitutionIdOk

`func (o *ExternalAccountLinkToken) GetVendorInstitutionIdOk() (*string, bool)`

GetVendorInstitutionIdOk returns a tuple with the VendorInstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInstitutionId

`func (o *ExternalAccountLinkToken) SetVendorInstitutionId(v string)`

SetVendorInstitutionId sets VendorInstitutionId field to given value.

### HasVendorInstitutionId

`func (o *ExternalAccountLinkToken) HasVendorInstitutionId() bool`

HasVendorInstitutionId returns a boolean if a field has been set.

### GetType

`func (o *ExternalAccountLinkToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalAccountLinkToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalAccountLinkToken) SetType(v string)`

SetType sets Type field to given value.


### GetClientName

`func (o *ExternalAccountLinkToken) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ExternalAccountLinkToken) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ExternalAccountLinkToken) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetLanguage

`func (o *ExternalAccountLinkToken) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ExternalAccountLinkToken) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ExternalAccountLinkToken) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCountryCodes

`func (o *ExternalAccountLinkToken) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ExternalAccountLinkToken) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ExternalAccountLinkToken) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.


### GetLinkCustomizationName

`func (o *ExternalAccountLinkToken) GetLinkCustomizationName() string`

GetLinkCustomizationName returns the LinkCustomizationName field if non-nil, zero value otherwise.

### GetLinkCustomizationNameOk

`func (o *ExternalAccountLinkToken) GetLinkCustomizationNameOk() (*string, bool)`

GetLinkCustomizationNameOk returns a tuple with the LinkCustomizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCustomizationName

`func (o *ExternalAccountLinkToken) SetLinkCustomizationName(v string)`

SetLinkCustomizationName sets LinkCustomizationName field to given value.

### HasLinkCustomizationName

`func (o *ExternalAccountLinkToken) HasLinkCustomizationName() bool`

HasLinkCustomizationName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *ExternalAccountLinkToken) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ExternalAccountLinkToken) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ExternalAccountLinkToken) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *ExternalAccountLinkToken) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetLinkToken

`func (o *ExternalAccountLinkToken) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *ExternalAccountLinkToken) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *ExternalAccountLinkToken) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.

### HasLinkToken

`func (o *ExternalAccountLinkToken) HasLinkToken() bool`

HasLinkToken returns a boolean if a field has been set.

### GetVendorAccessToken

`func (o *ExternalAccountLinkToken) GetVendorAccessToken() string`

GetVendorAccessToken returns the VendorAccessToken field if non-nil, zero value otherwise.

### GetVendorAccessTokenOk

`func (o *ExternalAccountLinkToken) GetVendorAccessTokenOk() (*string, bool)`

GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccessToken

`func (o *ExternalAccountLinkToken) SetVendorAccessToken(v string)`

SetVendorAccessToken sets VendorAccessToken field to given value.

### HasVendorAccessToken

`func (o *ExternalAccountLinkToken) HasVendorAccessToken() bool`

HasVendorAccessToken returns a boolean if a field has been set.

### GetExpiration

`func (o *ExternalAccountLinkToken) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ExternalAccountLinkToken) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ExternalAccountLinkToken) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ExternalAccountLinkToken) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRequestId

`func (o *ExternalAccountLinkToken) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ExternalAccountLinkToken) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ExternalAccountLinkToken) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ExternalAccountLinkToken) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetVerifyOwner

`func (o *ExternalAccountLinkToken) GetVerifyOwner() bool`

GetVerifyOwner returns the VerifyOwner field if non-nil, zero value otherwise.

### GetVerifyOwnerOk

`func (o *ExternalAccountLinkToken) GetVerifyOwnerOk() (*bool, bool)`

GetVerifyOwnerOk returns a tuple with the VerifyOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyOwner

`func (o *ExternalAccountLinkToken) SetVerifyOwner(v bool)`

SetVerifyOwner sets VerifyOwner field to given value.

### HasVerifyOwner

`func (o *ExternalAccountLinkToken) HasVerifyOwner() bool`

HasVerifyOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


