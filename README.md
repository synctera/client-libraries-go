# Go API client for synctera

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./synctera"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.synctera.com/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**CreateAccount**](synctera/docs/AccountsApi.md#createaccount) | **Post** /accounts | Create an account
*AccountsApi* | [**CreateAccountAddress**](synctera/docs/AccountsApi.md#createaccountaddress) | **Post** /accounts/{account_id}/addresses | Create address for an account
*AccountsApi* | [**CreateAccountRelationship**](synctera/docs/AccountsApi.md#createaccountrelationship) | **Post** /accounts/{account_id}/relationships | Create account relationship
*AccountsApi* | [**DeleteAccount**](synctera/docs/AccountsApi.md#deleteaccount) | **Delete** /accounts/{account_id} | Delete account
*AccountsApi* | [**DeleteAccountAddress**](synctera/docs/AccountsApi.md#deleteaccountaddress) | **Delete** /accounts/{account_id}/addresses/{connect_id} | Delete address for an account
*AccountsApi* | [**DeleteAccountRelationship**](synctera/docs/AccountsApi.md#deleteaccountrelationship) | **Delete** /accounts/{account_id}/relationships/{relationship_id} | Delete account relationship
*AccountsApi* | [**GetAccount**](synctera/docs/AccountsApi.md#getaccount) | **Get** /accounts/{account_id} | Get account
*AccountsApi* | [**GetAccountAddress**](synctera/docs/AccountsApi.md#getaccountaddress) | **Get** /accounts/{account_id}/addresses/{connect_id} | Get address for an account
*AccountsApi* | [**GetAccountBalance**](synctera/docs/AccountsApi.md#getaccountbalance) | **Get** /accounts/{account_id}/balance | Get account balance
*AccountsApi* | [**GetAccountRelationship**](synctera/docs/AccountsApi.md#getaccountrelationship) | **Get** /accounts/{account_id}/relationships/{relationship_id} | Get account relationship
*AccountsApi* | [**GetAccountTransactions**](synctera/docs/AccountsApi.md#getaccounttransactions) | **Get** /accounts/{account_id}/transactions | Get account transactions
*AccountsApi* | [**ListAccountRelationship**](synctera/docs/AccountsApi.md#listaccountrelationship) | **Get** /accounts/{account_id}/relationships | List account relationships
*AccountsApi* | [**ListAccounts**](synctera/docs/AccountsApi.md#listaccounts) | **Get** /accounts | List accounts
*AccountsApi* | [**PatchAccount**](synctera/docs/AccountsApi.md#patchaccount) | **Patch** /accounts/{account_id} | Patch account
*AccountsApi* | [**UpdateAccount**](synctera/docs/AccountsApi.md#updateaccount) | **Put** /accounts/{account_id} | Update account
*AccountsApi* | [**UpdateAccountAddress**](synctera/docs/AccountsApi.md#updateaccountaddress) | **Put** /accounts/{account_id}/addresses/{connect_id} | Update address for an account
*AccountsApi* | [**UpdateAccountRelationship**](synctera/docs/AccountsApi.md#updateaccountrelationship) | **Put** /accounts/{account_id}/relationships/{relationship_id} | Update account relationship
*CardTransactionSimulationsApi* | [**SimulateAuthorization**](synctera/docs/CardTransactionSimulationsApi.md#simulateauthorization) | **Post** /cards/transaction_simulations/authorization | Simulate authorization
*CardTransactionSimulationsApi* | [**SimulateAuthorizationAdvice**](synctera/docs/CardTransactionSimulationsApi.md#simulateauthorizationadvice) | **Post** /cards/transaction_simulations/authorization/advice | Simulate authorization advice
*CardTransactionSimulationsApi* | [**SimulateBalanceInquiry**](synctera/docs/CardTransactionSimulationsApi.md#simulatebalanceinquiry) | **Post** /cards/transaction_simulations/financial/balance_inquiry | Simulate balance inquiry
*CardTransactionSimulationsApi* | [**SimulateClearing**](synctera/docs/CardTransactionSimulationsApi.md#simulateclearing) | **Post** /cards/transaction_simulations/clearing | Simulate clearing or refund
*CardTransactionSimulationsApi* | [**SimulateFinancial**](synctera/docs/CardTransactionSimulationsApi.md#simulatefinancial) | **Post** /cards/transaction_simulations/financial | Simulate financial
*CardTransactionSimulationsApi* | [**SimulateFinancialAdvice**](synctera/docs/CardTransactionSimulationsApi.md#simulatefinancialadvice) | **Post** /cards/transaction_simulations/financial/advice | Simulate financial advice
*CardTransactionSimulationsApi* | [**SimulateOriginalCredit**](synctera/docs/CardTransactionSimulationsApi.md#simulateoriginalcredit) | **Post** /cards/transaction_simulations/financial/original_credit | Simulate OCT
*CardTransactionSimulationsApi* | [**SimulateReversal**](synctera/docs/CardTransactionSimulationsApi.md#simulatereversal) | **Post** /cards/transaction_simulations/reversal | Simulate reversal
*CardTransactionSimulationsApi* | [**SimulateWithdrawal**](synctera/docs/CardTransactionSimulationsApi.md#simulatewithdrawal) | **Post** /cards/transaction_simulations/financial/withdrawal | Simulate ATM withdrawal
*CardsApi* | [**ActivateCard**](synctera/docs/CardsApi.md#activatecard) | **Post** /cards/{card_id}/activate | Activate a card
*CardsApi* | [**GetCard**](synctera/docs/CardsApi.md#getcard) | **Get** /cards/{card_id} | Get Card
*CardsApi* | [**GetCardProduct**](synctera/docs/CardsApi.md#getcardproduct) | **Get** /cards/products/{card_product_id} | Get details about a card product
*CardsApi* | [**GetCardProducts**](synctera/docs/CardsApi.md#getcardproducts) | **Get** /cards/products | List Cards Products
*CardsApi* | [**GetClientAccessToken**](synctera/docs/CardsApi.md#getclientaccesstoken) | **Post** /cards/{card_id}/client_token | Get a client token
*CardsApi* | [**GetClientSingleUseToken**](synctera/docs/CardsApi.md#getclientsingleusetoken) | **Post** /cards/single_use_token | Get single-use token
*CardsApi* | [**IssueCard**](synctera/docs/CardsApi.md#issuecard) | **Post** /cards | Issue a Card
*CardsApi* | [**ListCards**](synctera/docs/CardsApi.md#listcards) | **Get** /cards | List Cards
*CardsApi* | [**ListChanges**](synctera/docs/CardsApi.md#listchanges) | **Get** /cards/{card_id}/changes | List Card Changes
*CardsApi* | [**UpdateCard**](synctera/docs/CardsApi.md#updatecard) | **Patch** /cards/{card_id} | Update Card
*CustomersApi* | [**CreateCustomer**](synctera/docs/CustomersApi.md#createcustomer) | **Post** /customers | Create a Customer
*CustomersApi* | [**CreateCustomerEmployment**](synctera/docs/CustomersApi.md#createcustomeremployment) | **Post** /customers/{customer_id}/employment | Create employment record
*CustomersApi* | [**CreateCustomerRiskRating**](synctera/docs/CustomersApi.md#createcustomerriskrating) | **Post** /customers/{customer_id}/risk_ratings | Create customer risk rating
*CustomersApi* | [**GetAllCustomerEmployment**](synctera/docs/CustomersApi.md#getallcustomeremployment) | **Get** /customers/{customer_id}/employment | List customer employment records
*CustomersApi* | [**GetAllCustomerRiskRatings**](synctera/docs/CustomersApi.md#getallcustomerriskratings) | **Get** /customers/{customer_id}/risk_ratings | List customer risk ratings
*CustomersApi* | [**GetCustomer**](synctera/docs/CustomersApi.md#getcustomer) | **Get** /customers/{customer_id} | Get Customer
*CustomersApi* | [**GetCustomerAccount**](synctera/docs/CustomersApi.md#getcustomeraccount) | **Get** /customers/{customer_id}/accounts/{account_id} | Get customer account
*CustomersApi* | [**GetCustomerRiskRating**](synctera/docs/CustomersApi.md#getcustomerriskrating) | **Get** /customers/{customer_id}/risk_ratings/{risk_rating_id} | Get customer risk rating
*CustomersApi* | [**GetPartyEmployment**](synctera/docs/CustomersApi.md#getpartyemployment) | **Get** /customers/{customer_id}/employment/{employment_id} | Get customer employment record
*CustomersApi* | [**ListCustomerAccounts**](synctera/docs/CustomersApi.md#listcustomeraccounts) | **Get** /customers/{customer_id}/accounts | List accounts
*CustomersApi* | [**ListCustomerAddresses**](synctera/docs/CustomersApi.md#listcustomeraddresses) | **Get** /customers/{customer_id}/addresses | List customer addresses
*CustomersApi* | [**ListCustomers**](synctera/docs/CustomersApi.md#listcustomers) | **Get** /customers | List Customers
*CustomersApi* | [**PatchCustomer**](synctera/docs/CustomersApi.md#patchcustomer) | **Patch** /customers/{customer_id} | Patch Customer
*CustomersApi* | [**UpdateCustomer**](synctera/docs/CustomersApi.md#updatecustomer) | **Put** /customers/{customer_id} | Update Customer
*CustomersApi* | [**UpdatePartyEmployment**](synctera/docs/CustomersApi.md#updatepartyemployment) | **Put** /customers/{customer_id}/employment/{employment_id} | Update customer employment record
*DisclosuresApi* | [**CreateDisclosure**](synctera/docs/DisclosuresApi.md#createdisclosure) | **Post** /customers/{customer_id}/disclosures | Create a Disclosure
*DisclosuresApi* | [**ListDisclosures**](synctera/docs/DisclosuresApi.md#listdisclosures) | **Get** /customers/{customer_id}/disclosures | List Disclosures
*KYCVerificationApi* | [**CreateCustomerVerificationResult**](synctera/docs/KYCVerificationApi.md#createcustomerverificationresult) | **Post** /customers/{customer_id}/verifications | Create a Customer Verification Result
*KYCVerificationApi* | [**GetVerification**](synctera/docs/KYCVerificationApi.md#getverification) | **Get** /customers/{customer_id}/verifications/{verification_id} | Get Verification Result
*KYCVerificationApi* | [**ListVerifications**](synctera/docs/KYCVerificationApi.md#listverifications) | **Get** /customers/{customer_id}/verifications | List Verification Results
*KYCVerificationApi* | [**VerifyCustomer**](synctera/docs/KYCVerificationApi.md#verifycustomer) | **Post** /customers/{customer_id}/verify | Verify a customer&#39;s identity
*RDCDepositsApi* | [**CreateRdcDeposit**](synctera/docs/RDCDepositsApi.md#createrdcdeposit) | **Post** /accounts/{account_id}/rdc/deposit | Create an RDC Deposit
*RDCDepositsApi* | [**CreateRdcImage**](synctera/docs/RDCDepositsApi.md#createrdcimage) | **Post** /accounts/{account_id}/rdc/images | Create an RDC Image
*RDCDepositsApi* | [**CreateRdcScan**](synctera/docs/RDCDepositsApi.md#createrdcscan) | **Post** /accounts/{account_id}/rdc/scans | Create an RDC Scan
*RDCDepositsApi* | [**GetRdcDeposit**](synctera/docs/RDCDepositsApi.md#getrdcdeposit) | **Get** /accounts/{account_id}/rdc/deposit/{deposit_id} | Get RDC Deposit
*RDCDepositsApi* | [**GetRdcImage**](synctera/docs/RDCDepositsApi.md#getrdcimage) | **Get** /accounts/{account_id}/rdc/images/{image_id} | Get RDC Image
*RDCDepositsApi* | [**GetRdcScan**](synctera/docs/RDCDepositsApi.md#getrdcscan) | **Get** /accounts/{account_id}/rdc/scans/{scan_id} | Get RDC Scan
*RDCDepositsApi* | [**ListRdcDeposits**](synctera/docs/RDCDepositsApi.md#listrdcdeposits) | **Get** /accounts/{account_id}/rdc/deposit | List RDC Deposits
*RDCDepositsApi* | [**ListRdcImageIds**](synctera/docs/RDCDepositsApi.md#listrdcimageids) | **Get** /accounts/{account_id}/rdc/images | List RDC Image Ids
*RDCDepositsApi* | [**ListRdcScans**](synctera/docs/RDCDepositsApi.md#listrdcscans) | **Get** /accounts/{account_id}/rdc/scans | List RDC Scans
*ReconciliationsApi* | [**CreateReconciliation**](synctera/docs/ReconciliationsApi.md#createreconciliation) | **Post** /reconciliations | Create a reconciliation
*ReconciliationsApi* | [**GetReconciliation**](synctera/docs/ReconciliationsApi.md#getreconciliation) | **Get** /reconciliations/{reconciliation_id} | Get reconciliation
*ReconciliationsApi* | [**ListReconciliations**](synctera/docs/ReconciliationsApi.md#listreconciliations) | **Get** /reconciliations | List reconciliations
*TransactionsApi* | [**A2aTransfer**](synctera/docs/TransactionsApi.md#a2atransfer) | **Post** /transactions/a2a_transfer | Account to account transfer
*TransactionsApi* | [**CreateOutgoingACH**](synctera/docs/TransactionsApi.md#createoutgoingach) | **Post** /transactions/ach | Create outgoing ACH
*TransactionsApi* | [**DeleteOutgoingACH**](synctera/docs/TransactionsApi.md#deleteoutgoingach) | **Delete** /transactions/ach/{payment_id} | Delete pending outgoing ACH
*TransactionsApi* | [**GetA2ATransfer**](synctera/docs/TransactionsApi.md#geta2atransfer) | **Get** /transactions/a2a_transfer/{payment_id} | Get account to account transfer
*TransactionsApi* | [**GetOutgoingACH**](synctera/docs/TransactionsApi.md#getoutgoingach) | **Get** /transactions/ach | Get Pending ACH List
*TransactionsApi* | [**ListA2ATransfer**](synctera/docs/TransactionsApi.md#lista2atransfer) | **Get** /transactions/a2a_transfer/list/{customer_id} | List account to account transfer
*TransactionsApi* | [**ReverseA2ATransfer**](synctera/docs/TransactionsApi.md#reversea2atransfer) | **Post** /transactions/a2a_transfer/{payment_id} | Reverse existing account to account transfer
*TransactionsApi* | [**UpdateOutgoingACH**](synctera/docs/TransactionsApi.md#updateoutgoingach) | **Put** /transactions/ach/{payment_id} | Update outgoing ACH
*WatchlistApi* | [**GetWatchlistAlert**](synctera/docs/WatchlistApi.md#getwatchlistalert) | **Get** /customers/{customer_id}/watchlists/alerts/{alert_id} | Retrieve watchlist monitoring alert
*WatchlistApi* | [**GetWatchlistSubscription**](synctera/docs/WatchlistApi.md#getwatchlistsubscription) | **Get** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Retrieve watchlist monitoring subscription
*WatchlistApi* | [**ListWatchlistAlerts**](synctera/docs/WatchlistApi.md#listwatchlistalerts) | **Get** /customers/{customer_id}/watchlists/alerts | List watchlist monitoring alerts for a customer
*WatchlistApi* | [**ListWatchlistSubscriptions**](synctera/docs/WatchlistApi.md#listwatchlistsubscriptions) | **Get** /customers/{customer_id}/watchlists/subscriptions | List watchlist monitoring subscriptions for a customer
*WatchlistApi* | [**SuppressWatchlistEntityAlert**](synctera/docs/WatchlistApi.md#suppresswatchlistentityalert) | **Post** /customers/{customer_id}/watchlists/suppressions | Suppress entity alert
*WatchlistApi* | [**UpdateWatchlistAlert**](synctera/docs/WatchlistApi.md#updatewatchlistalert) | **Put** /customers/{customer_id}/watchlists/alerts/{alert_id} | Update watchlist alert
*WatchlistApi* | [**UpdateWatchlistSubscription**](synctera/docs/WatchlistApi.md#updatewatchlistsubscription) | **Put** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Update watchlist monitoring subscription
*WatchlistApi* | [**WatchlistSubscribe**](synctera/docs/WatchlistApi.md#watchlistsubscribe) | **Post** /customers/{customer_id}/watchlists/subscriptions | Subscribe a customer to watchlist monitoring
*WebhooksApi* | [**CreateSecret**](synctera/docs/WebhooksApi.md#createsecret) | **Post** /webhooks/secret | Create a secret
*WebhooksApi* | [**CreateWebhook1**](synctera/docs/WebhooksApi.md#createwebhook1) | **Post** /webhooks | Create a webhook
*WebhooksApi* | [**DeleteWebhook**](synctera/docs/WebhooksApi.md#deletewebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
*WebhooksApi* | [**GetEvent**](synctera/docs/WebhooksApi.md#getevent) | **Get** /webhooks/{webhook_id}/events/{event_id} | Get webhook event
*WebhooksApi* | [**GetWebhook1**](synctera/docs/WebhooksApi.md#getwebhook1) | **Get** /webhooks/{webhook_id} | Get a webhook
*WebhooksApi* | [**ListEvents**](synctera/docs/WebhooksApi.md#listevents) | **Get** /webhooks/{webhook_id}/events | List webhook events
*WebhooksApi* | [**ListWebhooks1**](synctera/docs/WebhooksApi.md#listwebhooks1) | **Get** /webhooks | List webhooks
*WebhooksApi* | [**ReplaceSecret**](synctera/docs/WebhooksApi.md#replacesecret) | **Put** /webhooks/secret | Replace an existing secret
*WebhooksApi* | [**ResendEvent**](synctera/docs/WebhooksApi.md#resendevent) | **Post** /webhooks/{webhook_id}/events/{event_id}/resend | Resend an event
*WebhooksApi* | [**RevokeSecret**](synctera/docs/WebhooksApi.md#revokesecret) | **Delete** /webhooks/secret | Revoke the secret
*WebhooksApi* | [**TriggerEvent**](synctera/docs/WebhooksApi.md#triggerevent) | **Post** /webhooks/trigger | Trigger an event
*WebhooksApi* | [**UpdateWebhook**](synctera/docs/WebhooksApi.md#updatewebhook) | **Put** /webhooks/{webhook_id} | Update a webhook


## Documentation For Models

 - [A2aTransfer](synctera/docs/A2aTransfer.md)
 - [A2aTransferList](synctera/docs/A2aTransferList.md)
 - [A2aTransferListAllOf](synctera/docs/A2aTransferListAllOf.md)
 - [A2aTransferTransferReversal](synctera/docs/A2aTransferTransferReversal.md)
 - [Account](synctera/docs/Account.md)
 - [AccountAddress](synctera/docs/AccountAddress.md)
 - [AccountCreation](synctera/docs/AccountCreation.md)
 - [AccountCreationAllOf](synctera/docs/AccountCreationAllOf.md)
 - [AccountList](synctera/docs/AccountList.md)
 - [AccountListAllOf](synctera/docs/AccountListAllOf.md)
 - [AccountType](synctera/docs/AccountType.md)
 - [AchOutgoing](synctera/docs/AchOutgoing.md)
 - [AchOutgoingList](synctera/docs/AchOutgoingList.md)
 - [AchOutgoingListAllOf](synctera/docs/AchOutgoingListAllOf.md)
 - [Address](synctera/docs/Address.md)
 - [Address1](synctera/docs/Address1.md)
 - [AddressList](synctera/docs/AddressList.md)
 - [AddressListAllOf](synctera/docs/AddressListAllOf.md)
 - [Alias](synctera/docs/Alias.md)
 - [AuthRequestModel](synctera/docs/AuthRequestModel.md)
 - [AuthorizationAdviceModel](synctera/docs/AuthorizationAdviceModel.md)
 - [Balance](synctera/docs/Balance.md)
 - [BalanceInquiryRequestModel](synctera/docs/BalanceInquiryRequestModel.md)
 - [BalanceType](synctera/docs/BalanceType.md)
 - [BaseCard](synctera/docs/BaseCard.md)
 - [BaseCardAllOf](synctera/docs/BaseCardAllOf.md)
 - [BasePerson](synctera/docs/BasePerson.md)
 - [BillingAddress](synctera/docs/BillingAddress.md)
 - [CardAcceptorModel](synctera/docs/CardAcceptorModel.md)
 - [CardActivation](synctera/docs/CardActivation.md)
 - [CardChange](synctera/docs/CardChange.md)
 - [CardChangeReasonCode](synctera/docs/CardChangeReasonCode.md)
 - [CardChangeState](synctera/docs/CardChangeState.md)
 - [CardChangesList](synctera/docs/CardChangesList.md)
 - [CardChangesListAllOf](synctera/docs/CardChangesListAllOf.md)
 - [CardEditRequest](synctera/docs/CardEditRequest.md)
 - [CardFormat](synctera/docs/CardFormat.md)
 - [CardFulfillmentStatus](synctera/docs/CardFulfillmentStatus.md)
 - [CardIssuanceRequest](synctera/docs/CardIssuanceRequest.md)
 - [CardListResponse](synctera/docs/CardListResponse.md)
 - [CardListResponseAllOf](synctera/docs/CardListResponseAllOf.md)
 - [CardOptions](synctera/docs/CardOptions.md)
 - [CardPin](synctera/docs/CardPin.md)
 - [CardProduct](synctera/docs/CardProduct.md)
 - [CardProductAllOf](synctera/docs/CardProductAllOf.md)
 - [CardProductList](synctera/docs/CardProductList.md)
 - [CardProductListAllOf](synctera/docs/CardProductListAllOf.md)
 - [CardResponse](synctera/docs/CardResponse.md)
 - [CardStatus](synctera/docs/CardStatus.md)
 - [CardStatusObject](synctera/docs/CardStatusObject.md)
 - [ChangeChannel](synctera/docs/ChangeChannel.md)
 - [ChangeType](synctera/docs/ChangeType.md)
 - [ClearingModel](synctera/docs/ClearingModel.md)
 - [ClientToken](synctera/docs/ClientToken.md)
 - [CreateWebhookRequest](synctera/docs/CreateWebhookRequest.md)
 - [CustomHeaders](synctera/docs/CustomHeaders.md)
 - [Customer](synctera/docs/Customer.md)
 - [CustomerAllOf](synctera/docs/CustomerAllOf.md)
 - [CustomerInPath](synctera/docs/CustomerInPath.md)
 - [CustomerKycStatus](synctera/docs/CustomerKycStatus.md)
 - [CustomerList](synctera/docs/CustomerList.md)
 - [CustomerListAllOf](synctera/docs/CustomerListAllOf.md)
 - [CustomerVerification](synctera/docs/CustomerVerification.md)
 - [CustomerVerificationReasonCodes](synctera/docs/CustomerVerificationReasonCodes.md)
 - [CustomerVerificationResult](synctera/docs/CustomerVerificationResult.md)
 - [CustomerVerificationResultList](synctera/docs/CustomerVerificationResultList.md)
 - [CustomerVerificationResultListAllOf](synctera/docs/CustomerVerificationResultListAllOf.md)
 - [DcSignType](synctera/docs/DcSignType.md)
 - [DeleteResponse](synctera/docs/DeleteResponse.md)
 - [Deposit](synctera/docs/Deposit.md)
 - [DepositList](synctera/docs/DepositList.md)
 - [DepositListAllOf](synctera/docs/DepositListAllOf.md)
 - [Device](synctera/docs/Device.md)
 - [Disclosure](synctera/docs/Disclosure.md)
 - [DisclosureResponse](synctera/docs/DisclosureResponse.md)
 - [Document](synctera/docs/Document.md)
 - [DocumentList](synctera/docs/DocumentList.md)
 - [DocumentListAllOf](synctera/docs/DocumentListAllOf.md)
 - [DocumentType](synctera/docs/DocumentType.md)
 - [EmbossName](synctera/docs/EmbossName.md)
 - [Employment](synctera/docs/Employment.md)
 - [EmploymentList](synctera/docs/EmploymentList.md)
 - [EmploymentListAllOf](synctera/docs/EmploymentListAllOf.md)
 - [Event](synctera/docs/Event.md)
 - [EventList](synctera/docs/EventList.md)
 - [EventListAllOf](synctera/docs/EventListAllOf.md)
 - [EventResend](synctera/docs/EventResend.md)
 - [EventResponseHistory](synctera/docs/EventResponseHistory.md)
 - [EventType](synctera/docs/EventType.md)
 - [EventType1](synctera/docs/EventType1.md)
 - [EventTypeExplicit](synctera/docs/EventTypeExplicit.md)
 - [EventTypeWildcard](synctera/docs/EventTypeWildcard.md)
 - [FinancialRequestModel](synctera/docs/FinancialRequestModel.md)
 - [Form](synctera/docs/Form.md)
 - [Image](synctera/docs/Image.md)
 - [ImageList](synctera/docs/ImageList.md)
 - [ImageListAllOf](synctera/docs/ImageListAllOf.md)
 - [IngestionStatus](synctera/docs/IngestionStatus.md)
 - [InlineObject](synctera/docs/InlineObject.md)
 - [InlineObject1](synctera/docs/InlineObject1.md)
 - [InlineObject2](synctera/docs/InlineObject2.md)
 - [InlineObject3](synctera/docs/InlineObject3.md)
 - [InlineObject4](synctera/docs/InlineObject4.md)
 - [InlineResponse200](synctera/docs/InlineResponse200.md)
 - [InlineResponse201](synctera/docs/InlineResponse201.md)
 - [KycMediaType](synctera/docs/KycMediaType.md)
 - [ModelError](synctera/docs/ModelError.md)
 - [NetworkFeeModel](synctera/docs/NetworkFeeModel.md)
 - [OriginalCreditRequestModel](synctera/docs/OriginalCreditRequestModel.md)
 - [OriginalCreditSenderData](synctera/docs/OriginalCreditSenderData.md)
 - [PaginatedResponse](synctera/docs/PaginatedResponse.md)
 - [PhysicalCard](synctera/docs/PhysicalCard.md)
 - [PhysicalCardAllOf](synctera/docs/PhysicalCardAllOf.md)
 - [PhysicalCardIssuanceRequest](synctera/docs/PhysicalCardIssuanceRequest.md)
 - [PhysicalCardPlusStatus](synctera/docs/PhysicalCardPlusStatus.md)
 - [PhysicalCardResponse](synctera/docs/PhysicalCardResponse.md)
 - [PhysicalCardResponseStatus](synctera/docs/PhysicalCardResponseStatus.md)
 - [PingResponse](synctera/docs/PingResponse.md)
 - [Prospect](synctera/docs/Prospect.md)
 - [ProspectAllOf](synctera/docs/ProspectAllOf.md)
 - [ProviderType](synctera/docs/ProviderType.md)
 - [RawResponse](synctera/docs/RawResponse.md)
 - [RdcMediaType](synctera/docs/RdcMediaType.md)
 - [RecipientName](synctera/docs/RecipientName.md)
 - [Reconciliation](synctera/docs/Reconciliation.md)
 - [ReconciliationInput](synctera/docs/ReconciliationInput.md)
 - [ReconciliationList](synctera/docs/ReconciliationList.md)
 - [ReconciliationListAllOf](synctera/docs/ReconciliationListAllOf.md)
 - [RecurrenceData](synctera/docs/RecurrenceData.md)
 - [Relationship](synctera/docs/Relationship.md)
 - [Relationship1](synctera/docs/Relationship1.md)
 - [RelationshipList](synctera/docs/RelationshipList.md)
 - [RelationshipListAllOf](synctera/docs/RelationshipListAllOf.md)
 - [RelationshipRole](synctera/docs/RelationshipRole.md)
 - [ResendResponse](synctera/docs/ResendResponse.md)
 - [ReversalModel](synctera/docs/ReversalModel.md)
 - [RiskRating](synctera/docs/RiskRating.md)
 - [RiskRatingList](synctera/docs/RiskRatingList.md)
 - [RiskRatingListAllOf](synctera/docs/RiskRatingListAllOf.md)
 - [Scan](synctera/docs/Scan.md)
 - [ScanList](synctera/docs/ScanList.md)
 - [ScanListAllOf](synctera/docs/ScanListAllOf.md)
 - [Shipping](synctera/docs/Shipping.md)
 - [SingleUseTokenRequest](synctera/docs/SingleUseTokenRequest.md)
 - [SingleUseTokenResponse](synctera/docs/SingleUseTokenResponse.md)
 - [SocureEvent](synctera/docs/SocureEvent.md)
 - [SocureEventBody](synctera/docs/SocureEventBody.md)
 - [SocureGlobalWatchlist](synctera/docs/SocureGlobalWatchlist.md)
 - [SocureMatch](synctera/docs/SocureMatch.md)
 - [SocureMatchComments](synctera/docs/SocureMatchComments.md)
 - [Status](synctera/docs/Status.md)
 - [Transaction](synctera/docs/Transaction.md)
 - [TransactionList](synctera/docs/TransactionList.md)
 - [TransactionListAllOf](synctera/docs/TransactionListAllOf.md)
 - [TransactionOptions](synctera/docs/TransactionOptions.md)
 - [VerificationType](synctera/docs/VerificationType.md)
 - [VirtualCard](synctera/docs/VirtualCard.md)
 - [VirtualCardIssuanceRequest](synctera/docs/VirtualCardIssuanceRequest.md)
 - [VirtualCardPlusStatus](synctera/docs/VirtualCardPlusStatus.md)
 - [VirtualCardResponse](synctera/docs/VirtualCardResponse.md)
 - [VirtualCardResponseStatus](synctera/docs/VirtualCardResponseStatus.md)
 - [WatchlistAlert](synctera/docs/WatchlistAlert.md)
 - [WatchlistAlertList](synctera/docs/WatchlistAlertList.md)
 - [WatchlistAlertListAllOf](synctera/docs/WatchlistAlertListAllOf.md)
 - [WatchlistSubscription](synctera/docs/WatchlistSubscription.md)
 - [WatchlistSubscriptionList](synctera/docs/WatchlistSubscriptionList.md)
 - [WatchlistSubscriptionListAllOf](synctera/docs/WatchlistSubscriptionListAllOf.md)
 - [WatchlistSuppress](synctera/docs/WatchlistSuppress.md)
 - [Webhook](synctera/docs/Webhook.md)
 - [WebhookConfig](synctera/docs/WebhookConfig.md)
 - [WebhookList](synctera/docs/WebhookList.md)
 - [WebhookListAllOf](synctera/docs/WebhookListAllOf.md)
 - [WebhookRequest](synctera/docs/WebhookRequest.md)
 - [WebhookResponse](synctera/docs/WebhookResponse.md)
 - [WebhookResponseAllOf](synctera/docs/WebhookResponseAllOf.md)
 - [WithdrawalRequestModel](synctera/docs/WithdrawalRequestModel.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



