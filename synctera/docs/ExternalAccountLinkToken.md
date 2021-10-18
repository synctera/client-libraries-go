# ExternalAccountLinkToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | The name of your application, as it should be displayed in Link. Maximum length of 30 characters. | 
**CountryCodes** | **[]string** | Country codes in the ISO-3166-1 alpha-2 country code standard. | 
**Expiration** | Pointer to **time.Time** | The expiration date for the link_token. Expires in 4 hours. | [optional] 
**ExternalAccountUser** | [**ExternalAccountUser**](ExternalAccountUser.md) |  | 
**Language** | **string** | The language that Link should be displayed in. | 
**LinkCustomizationName** | Pointer to **string** | The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the default customization will be used. When using a Link customization, the language in the customization must match the language selected via the language parameter, and the countries in the customization should match the country codes selected via country_codes.  | [optional] 
**LinkToken** | Pointer to **string** | A link_token, which can be supplied to Link in order to initialize it and receive a public_token, which can be exchanged for an access_token.  | [optional] [readonly] 
**Products** | Pointer to **[]string** |  | [optional] 
**RedirectUri** | Pointer to **string** | A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview.  | [optional] 
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. | [optional] [readonly] 
**VendorAccessToken** | Pointer to **string** | The access_token associated with the Item to update, used when updating or modifying an existing access_token. Used when launching Link in update mode, when completing the Same-day (manual) Micro-deposit flow, or (optionally) when initializing Link as part of the Payment Initiation (UK and Europe) flow.  | [optional] 
**Webhook** | Pointer to **string** | The destination URL to which any webhooks should be sent. | [optional] 

## Methods

### NewExternalAccountLinkToken

`func NewExternalAccountLinkToken(clientName string, countryCodes []string, externalAccountUser ExternalAccountUser, language string, ) *ExternalAccountLinkToken`

NewExternalAccountLinkToken instantiates a new ExternalAccountLinkToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountLinkTokenWithDefaults

`func NewExternalAccountLinkTokenWithDefaults() *ExternalAccountLinkToken`

NewExternalAccountLinkTokenWithDefaults instantiates a new ExternalAccountLinkToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetExternalAccountUser

`func (o *ExternalAccountLinkToken) GetExternalAccountUser() ExternalAccountUser`

GetExternalAccountUser returns the ExternalAccountUser field if non-nil, zero value otherwise.

### GetExternalAccountUserOk

`func (o *ExternalAccountLinkToken) GetExternalAccountUserOk() (*ExternalAccountUser, bool)`

GetExternalAccountUserOk returns a tuple with the ExternalAccountUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountUser

`func (o *ExternalAccountLinkToken) SetExternalAccountUser(v ExternalAccountUser)`

SetExternalAccountUser sets ExternalAccountUser field to given value.


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

### GetProducts

`func (o *ExternalAccountLinkToken) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ExternalAccountLinkToken) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ExternalAccountLinkToken) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ExternalAccountLinkToken) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

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

### GetWebhook

`func (o *ExternalAccountLinkToken) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ExternalAccountLinkToken) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ExternalAccountLinkToken) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *ExternalAccountLinkToken) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


