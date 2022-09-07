/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// SandboxWipeAlphaApiService SandboxWipeAlphaApi service
type SandboxWipeAlphaApiService service

type ApiWipeWorkspaceRequest struct {
	ctx context.Context
	ApiService *SandboxWipeAlphaApiService
}

func (r ApiWipeWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.WipeWorkspaceExecute(r)
}

/*
WipeWorkspace Delete data

Delete the customer and account related data, leaving other configuration data intact. This enables use cases such as bulk data deletion between tests.
Data associated with below resources will be deleted:
  - Accounts
  - Account applications
  - ACH
  - Businesses
  - Cards
  - Card images
  - Cases
  - Customers
  - Disclosures
  - External Accounts
  - Internal Accounts
  - Payment schedules and history
  - Persons
  - RDC
  - Relationships
  - Transactions (including for Internal Accounts)
  - Verifications

Data associated with below resources will be retained:
  - Account Templates
  - API Keys
  - Bank/Partner data
  - Card product
  - Disclosure document records
  - Egress config
  - Groups
  - PII contract with vault
  - Roles
  - Users
  - Webhooks


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWipeWorkspaceRequest
*/
func (a *SandboxWipeAlphaApiService) WipeWorkspace(ctx context.Context) ApiWipeWorkspaceRequest {
	return ApiWipeWorkspaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SandboxWipeAlphaApiService) WipeWorkspaceExecute(r ApiWipeWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxWipeAlphaApiService.WipeWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wipe"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
