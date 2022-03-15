/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// ExternalAccountLinkToken struct for ExternalAccountLinkToken
type ExternalAccountLinkToken struct {
	// The name of your application, as it should be displayed in Link. Maximum length of 30 characters.
	ClientName string `json:"client_name"`
	// Country codes in the ISO-3166-1 alpha-2 country code standard.
	CountryCodes []string `json:"country_codes"`
	// The identifier for the customer associated with this account
	CustomerId string `json:"customer_id"`
	// The expiration date for the link_token. Expires in 4 hours.
	Expiration *time.Time `json:"expiration,omitempty"`
	// The language that corresponds to the link token. For Plaid, see their [documentation](https://plaid.com/docs/api/tokens/#link-token-create-request-language) for a list of allowed values.
	Language string `json:"language"`
	// The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the default customization will be used. When using a Link customization, the language in the customization must match the language selected via the language parameter, and the countries in the customization should match the country codes selected via country_codes.
	LinkCustomizationName *string `json:"link_customization_name,omitempty"`
	// A link_token, which can be supplied to Link in order to initialize it and receive a public_token, which can be exchanged for an access_token.
	LinkToken *string `json:"link_token,omitempty"`
	// A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview.
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting.
	RequestId *string `json:"request_id,omitempty"`
	// The type of the link token. DEPOSITORY for checking and savings accounts, CREDIT for credit card type accounts.
	Type string `json:"type"`
	// The access token associated with the Item data is being requested for.
	VendorAccessToken *string `json:"vendor_access_token,omitempty"`
	// The ID of the institution the access token is requested for. If present the link token will be created in an update mode.
	VendorInstitutionId *string `json:"vendor_institution_id,omitempty"`
	// Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification will be suppressed by default
	VerifyOwner *bool `json:"verify_owner,omitempty"`
}

// NewExternalAccountLinkToken instantiates a new ExternalAccountLinkToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountLinkToken(clientName string, countryCodes []string, customerId string, language string, type_ string) *ExternalAccountLinkToken {
	this := ExternalAccountLinkToken{}
	this.ClientName = clientName
	this.CountryCodes = countryCodes
	this.CustomerId = customerId
	this.Language = language
	this.Type = type_
	var verifyOwner bool = false
	this.VerifyOwner = &verifyOwner
	return &this
}

// NewExternalAccountLinkTokenWithDefaults instantiates a new ExternalAccountLinkToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountLinkTokenWithDefaults() *ExternalAccountLinkToken {
	this := ExternalAccountLinkToken{}
	var verifyOwner bool = false
	this.VerifyOwner = &verifyOwner
	return &this
}

// GetClientName returns the ClientName field value
func (o *ExternalAccountLinkToken) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value
func (o *ExternalAccountLinkToken) SetClientName(v string) {
	o.ClientName = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *ExternalAccountLinkToken) GetCountryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetCountryCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *ExternalAccountLinkToken) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCustomerId returns the CustomerId field value
func (o *ExternalAccountLinkToken) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ExternalAccountLinkToken) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ExternalAccountLinkToken) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetLanguage returns the Language field value
func (o *ExternalAccountLinkToken) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *ExternalAccountLinkToken) SetLanguage(v string) {
	o.Language = v
}

// GetLinkCustomizationName returns the LinkCustomizationName field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetLinkCustomizationName() string {
	if o == nil || o.LinkCustomizationName == nil {
		var ret string
		return ret
	}
	return *o.LinkCustomizationName
}

// GetLinkCustomizationNameOk returns a tuple with the LinkCustomizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetLinkCustomizationNameOk() (*string, bool) {
	if o == nil || o.LinkCustomizationName == nil {
		return nil, false
	}
	return o.LinkCustomizationName, true
}

// HasLinkCustomizationName returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasLinkCustomizationName() bool {
	if o != nil && o.LinkCustomizationName != nil {
		return true
	}

	return false
}

// SetLinkCustomizationName gets a reference to the given string and assigns it to the LinkCustomizationName field.
func (o *ExternalAccountLinkToken) SetLinkCustomizationName(v string) {
	o.LinkCustomizationName = &v
}

// GetLinkToken returns the LinkToken field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetLinkToken() string {
	if o == nil || o.LinkToken == nil {
		var ret string
		return ret
	}
	return *o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetLinkTokenOk() (*string, bool) {
	if o == nil || o.LinkToken == nil {
		return nil, false
	}
	return o.LinkToken, true
}

// HasLinkToken returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasLinkToken() bool {
	if o != nil && o.LinkToken != nil {
		return true
	}

	return false
}

// SetLinkToken gets a reference to the given string and assigns it to the LinkToken field.
func (o *ExternalAccountLinkToken) SetLinkToken(v string) {
	o.LinkToken = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *ExternalAccountLinkToken) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ExternalAccountLinkToken) SetRequestId(v string) {
	o.RequestId = &v
}

// GetType returns the Type field value
func (o *ExternalAccountLinkToken) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalAccountLinkToken) SetType(v string) {
	o.Type = v
}

// GetVendorAccessToken returns the VendorAccessToken field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetVendorAccessToken() string {
	if o == nil || o.VendorAccessToken == nil {
		var ret string
		return ret
	}
	return *o.VendorAccessToken
}

// GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetVendorAccessTokenOk() (*string, bool) {
	if o == nil || o.VendorAccessToken == nil {
		return nil, false
	}
	return o.VendorAccessToken, true
}

// HasVendorAccessToken returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasVendorAccessToken() bool {
	if o != nil && o.VendorAccessToken != nil {
		return true
	}

	return false
}

// SetVendorAccessToken gets a reference to the given string and assigns it to the VendorAccessToken field.
func (o *ExternalAccountLinkToken) SetVendorAccessToken(v string) {
	o.VendorAccessToken = &v
}

// GetVendorInstitutionId returns the VendorInstitutionId field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetVendorInstitutionId() string {
	if o == nil || o.VendorInstitutionId == nil {
		var ret string
		return ret
	}
	return *o.VendorInstitutionId
}

// GetVendorInstitutionIdOk returns a tuple with the VendorInstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetVendorInstitutionIdOk() (*string, bool) {
	if o == nil || o.VendorInstitutionId == nil {
		return nil, false
	}
	return o.VendorInstitutionId, true
}

// HasVendorInstitutionId returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasVendorInstitutionId() bool {
	if o != nil && o.VendorInstitutionId != nil {
		return true
	}

	return false
}

// SetVendorInstitutionId gets a reference to the given string and assigns it to the VendorInstitutionId field.
func (o *ExternalAccountLinkToken) SetVendorInstitutionId(v string) {
	o.VendorInstitutionId = &v
}

// GetVerifyOwner returns the VerifyOwner field value if set, zero value otherwise.
func (o *ExternalAccountLinkToken) GetVerifyOwner() bool {
	if o == nil || o.VerifyOwner == nil {
		var ret bool
		return ret
	}
	return *o.VerifyOwner
}

// GetVerifyOwnerOk returns a tuple with the VerifyOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountLinkToken) GetVerifyOwnerOk() (*bool, bool) {
	if o == nil || o.VerifyOwner == nil {
		return nil, false
	}
	return o.VerifyOwner, true
}

// HasVerifyOwner returns a boolean if a field has been set.
func (o *ExternalAccountLinkToken) HasVerifyOwner() bool {
	if o != nil && o.VerifyOwner != nil {
		return true
	}

	return false
}

// SetVerifyOwner gets a reference to the given bool and assigns it to the VerifyOwner field.
func (o *ExternalAccountLinkToken) SetVerifyOwner(v bool) {
	o.VerifyOwner = &v
}

func (o ExternalAccountLinkToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_name"] = o.ClientName
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if o.LinkCustomizationName != nil {
		toSerialize["link_customization_name"] = o.LinkCustomizationName
	}
	if o.LinkToken != nil {
		toSerialize["link_token"] = o.LinkToken
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.VendorAccessToken != nil {
		toSerialize["vendor_access_token"] = o.VendorAccessToken
	}
	if o.VendorInstitutionId != nil {
		toSerialize["vendor_institution_id"] = o.VendorInstitutionId
	}
	if o.VerifyOwner != nil {
		toSerialize["verify_owner"] = o.VerifyOwner
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccountLinkToken struct {
	value *ExternalAccountLinkToken
	isSet bool
}

func (v NullableExternalAccountLinkToken) Get() *ExternalAccountLinkToken {
	return v.value
}

func (v *NullableExternalAccountLinkToken) Set(val *ExternalAccountLinkToken) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountLinkToken) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountLinkToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountLinkToken(val *ExternalAccountLinkToken) *NullableExternalAccountLinkToken {
	return &NullableExternalAccountLinkToken{value: val, isSet: true}
}

func (v NullableExternalAccountLinkToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountLinkToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
