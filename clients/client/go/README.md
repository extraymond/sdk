# Go API client for client

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed
with a valid Personal Access Token. Public APIs are mostly used in browsers.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.0.1-alpha.15
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
import client "github.com/ory/client-go"
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
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://playground.projects.oryapis.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateIdentityAdmin**](docs/DefaultApi.md#createidentityadmin) | **Post** /api/kratos/admin/identities | Create an Identity
*DefaultApi* | [**CreateRecoveryLinkAdmin**](docs/DefaultApi.md#createrecoverylinkadmin) | **Post** /api/kratos/admin/recovery/link | Create a Recovery Link
*DefaultApi* | [**DeleteIdentityAdmin**](docs/DefaultApi.md#deleteidentityadmin) | **Delete** /api/kratos/admin/identities/{id} | Delete an Identity
*DefaultApi* | [**GetIdentityAdmin**](docs/DefaultApi.md#getidentityadmin) | **Get** /api/kratos/admin/identities/{id} | Get an Identity
*DefaultApi* | [**GetSchema**](docs/DefaultApi.md#getschema) | **Get** /api/kratos/public/schemas/{id} | 
*DefaultApi* | [**GetSchemaAdmin**](docs/DefaultApi.md#getschemaadmin) | **Get** /api/kratos/admin/schemas/{id} | 
*DefaultApi* | [**GetSelfServiceError**](docs/DefaultApi.md#getselfserviceerror) | **Get** /api/kratos/public/self-service/errors | Get User-Facing Self-Service Errors
*DefaultApi* | [**GetSelfServiceErrorAdmin**](docs/DefaultApi.md#getselfserviceerroradmin) | **Get** /api/kratos/admin/self-service/errors | Get User-Facing Self-Service Errors
*DefaultApi* | [**GetSelfServiceLoginFlow**](docs/DefaultApi.md#getselfserviceloginflow) | **Get** /api/kratos/public/self-service/login/flows | Get Login Flow
*DefaultApi* | [**GetSelfServiceLoginFlowAdmin**](docs/DefaultApi.md#getselfserviceloginflowadmin) | **Get** /api/kratos/admin/self-service/login/flows | Get Login Flow
*DefaultApi* | [**GetSelfServiceRecoveryFlow**](docs/DefaultApi.md#getselfservicerecoveryflow) | **Get** /api/kratos/public/self-service/recovery/flows | Get information about a recovery flow
*DefaultApi* | [**GetSelfServiceRecoveryFlowAdmin**](docs/DefaultApi.md#getselfservicerecoveryflowadmin) | **Get** /api/kratos/admin/self-service/recovery/flows | Get information about a recovery flow
*DefaultApi* | [**GetSelfServiceRegistrationFlow**](docs/DefaultApi.md#getselfserviceregistrationflow) | **Get** /api/kratos/public/self-service/registration/flows | Get Registration Flow
*DefaultApi* | [**GetSelfServiceRegistrationFlowAdmin**](docs/DefaultApi.md#getselfserviceregistrationflowadmin) | **Get** /api/kratos/admin/self-service/registration/flows | Get Registration Flow
*DefaultApi* | [**GetSelfServiceSettingsFlow**](docs/DefaultApi.md#getselfservicesettingsflow) | **Get** /api/kratos/public/self-service/settings/flows | Get Settings Flow
*DefaultApi* | [**GetSelfServiceSettingsFlowAdmin**](docs/DefaultApi.md#getselfservicesettingsflowadmin) | **Get** /api/kratos/admin/self-service/settings/flows | Get Settings Flow
*DefaultApi* | [**GetSelfServiceVerificationFlow**](docs/DefaultApi.md#getselfserviceverificationflow) | **Get** /api/kratos/public/self-service/verification/flows | Get Verification Flow
*DefaultApi* | [**GetSelfServiceVerificationFlowAdmin**](docs/DefaultApi.md#getselfserviceverificationflowadmin) | **Get** /api/kratos/admin/self-service/verification/flows | Get Verification Flow
*DefaultApi* | [**GetVersionAdmin**](docs/DefaultApi.md#getversionadmin) | **Get** /api/kratos/admin/version | Return Running Software Version.
*DefaultApi* | [**InitializeSelfServiceBrowserLogoutFlow**](docs/DefaultApi.md#initializeselfservicebrowserlogoutflow) | **Get** /api/kratos/public/self-service/browser/flows/logout | Initialize Browser-Based Logout User Flow
*DefaultApi* | [**InitializeSelfServiceLoginForBrowsers**](docs/DefaultApi.md#initializeselfserviceloginforbrowsers) | **Get** /api/kratos/public/self-service/login/browser | Initialize Login Flow for Browsers
*DefaultApi* | [**InitializeSelfServiceLoginWithoutBrowser**](docs/DefaultApi.md#initializeselfserviceloginwithoutbrowser) | **Get** /api/kratos/public/self-service/login/api | Initialize Login Flow for APIs, Services, Apps, ...
*DefaultApi* | [**InitializeSelfServiceRecoveryForBrowsers**](docs/DefaultApi.md#initializeselfservicerecoveryforbrowsers) | **Get** /api/kratos/public/self-service/recovery/browser | Initialize Recovery Flow for Browser Clients
*DefaultApi* | [**InitializeSelfServiceRecoveryForNativeApps**](docs/DefaultApi.md#initializeselfservicerecoveryfornativeapps) | **Get** /api/kratos/public/self-service/recovery/api | Initialize Recovery Flow for Native Apps and API clients
*DefaultApi* | [**InitializeSelfServiceRegistrationForBrowsers**](docs/DefaultApi.md#initializeselfserviceregistrationforbrowsers) | **Get** /api/kratos/public/self-service/registration/browser | Initialize Registration Flow for Browsers
*DefaultApi* | [**InitializeSelfServiceRegistrationWithoutBrowser**](docs/DefaultApi.md#initializeselfserviceregistrationwithoutbrowser) | **Get** /api/kratos/public/self-service/registration/api | Initialize Registration Flow for APIs, Services, Apps, ...
*DefaultApi* | [**InitializeSelfServiceSettingsForBrowsers**](docs/DefaultApi.md#initializeselfservicesettingsforbrowsers) | **Get** /api/kratos/public/self-service/settings/browser | Initialize Settings Flow for Browsers
*DefaultApi* | [**InitializeSelfServiceSettingsForNativeApps**](docs/DefaultApi.md#initializeselfservicesettingsfornativeapps) | **Get** /api/kratos/public/self-service/settings/api | Initialize Settings Flow for Native Apps and API clients
*DefaultApi* | [**InitializeSelfServiceVerificationForBrowsers**](docs/DefaultApi.md#initializeselfserviceverificationforbrowsers) | **Get** /api/kratos/public/self-service/verification/browser | Initialize Verification Flow for Browser Clients
*DefaultApi* | [**InitializeSelfServiceVerificationForNativeApps**](docs/DefaultApi.md#initializeselfserviceverificationfornativeapps) | **Get** /api/kratos/public/self-service/verification/api | Initialize Verification Flow for Native Apps and API clients
*DefaultApi* | [**IsAliveAdmin**](docs/DefaultApi.md#isaliveadmin) | **Get** /api/kratos/admin/health/alive | Check HTTP Server Status
*DefaultApi* | [**IsReadyAdmin**](docs/DefaultApi.md#isreadyadmin) | **Get** /api/kratos/admin/health/ready | Check HTTP Server and Database Status
*DefaultApi* | [**ListIdentitiesAdmin**](docs/DefaultApi.md#listidentitiesadmin) | **Get** /api/kratos/admin/identities | List Identities
*DefaultApi* | [**PrometheusAdmin**](docs/DefaultApi.md#prometheusadmin) | **Get** /api/kratos/admin/metrics/prometheus | Get snapshot metrics from the Hydra service. If you&#39;re using k8s, you can then add annotations to your deployment like so:
*DefaultApi* | [**RevokeSession**](docs/DefaultApi.md#revokesession) | **Delete** /api/kratos/public/sessions | Initialize Logout Flow for API Clients - Revoke a Session
*DefaultApi* | [**SubmitSelfServiceLoginFlow**](docs/DefaultApi.md#submitselfserviceloginflow) | **Post** /api/kratos/public/self-service/login | Submit a Login Flow
*DefaultApi* | [**SubmitSelfServiceRecoveryFlow**](docs/DefaultApi.md#submitselfservicerecoveryflow) | **Post** /api/kratos/public/self-service/recovery | Complete Recovery Flow
*DefaultApi* | [**SubmitSelfServiceRecoveryFlowWithLinkMethod**](docs/DefaultApi.md#submitselfservicerecoveryflowwithlinkmethod) | **Post** /api/kratos/public/self-service/recovery/methods/link | Complete Recovery Flow with Link Method
*DefaultApi* | [**SubmitSelfServiceRegistrationFlow**](docs/DefaultApi.md#submitselfserviceregistrationflow) | **Post** /api/kratos/public/self-service/registration | Submit a Registration Flow
*DefaultApi* | [**SubmitSelfServiceSettingsFlow**](docs/DefaultApi.md#submitselfservicesettingsflow) | **Post** /api/kratos/public/self-service/settings | Complete Settings Flow
*DefaultApi* | [**SubmitSelfServiceVerificationFlow**](docs/DefaultApi.md#submitselfserviceverificationflow) | **Post** /api/kratos/public/self-service/verification/flows | Complete Verification Flow
*DefaultApi* | [**ToSession**](docs/DefaultApi.md#tosession) | **Get** /api/kratos/public/sessions/whoami | Check Who the Current HTTP Session Belongs To
*DefaultApi* | [**UpdateIdentityAdmin**](docs/DefaultApi.md#updateidentityadmin) | **Put** /api/kratos/admin/identities/{id} | Update an Identity


## Documentation For Models

 - [AuthenticateOKBody](docs/AuthenticateOKBody.md)
 - [ContainerChangeResponseItem](docs/ContainerChangeResponseItem.md)
 - [ContainerCreateCreatedBody](docs/ContainerCreateCreatedBody.md)
 - [ContainerTopOKBody](docs/ContainerTopOKBody.md)
 - [ContainerUpdateOKBody](docs/ContainerUpdateOKBody.md)
 - [ContainerWaitOKBody](docs/ContainerWaitOKBody.md)
 - [ContainerWaitOKBodyError](docs/ContainerWaitOKBodyError.md)
 - [CreateIdentity](docs/CreateIdentity.md)
 - [CreateRecoveryLink](docs/CreateRecoveryLink.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GenericError](docs/GenericError.md)
 - [GraphDriverData](docs/GraphDriverData.md)
 - [HealthNotReadyStatus](docs/HealthNotReadyStatus.md)
 - [HealthStatus](docs/HealthStatus.md)
 - [IdResponse](docs/IdResponse.md)
 - [Identity](docs/Identity.md)
 - [IdentityCredentials](docs/IdentityCredentials.md)
 - [ImageDeleteResponseItem](docs/ImageDeleteResponseItem.md)
 - [ImageSummary](docs/ImageSummary.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse503](docs/InlineResponse503.md)
 - [JsonError](docs/JsonError.md)
 - [LoginFlow](docs/LoginFlow.md)
 - [LoginViaApiResponse](docs/LoginViaApiResponse.md)
 - [Meta](docs/Meta.md)
 - [Plugin](docs/Plugin.md)
 - [PluginConfig](docs/PluginConfig.md)
 - [PluginConfigArgs](docs/PluginConfigArgs.md)
 - [PluginConfigInterface](docs/PluginConfigInterface.md)
 - [PluginConfigLinux](docs/PluginConfigLinux.md)
 - [PluginConfigNetwork](docs/PluginConfigNetwork.md)
 - [PluginConfigRootfs](docs/PluginConfigRootfs.md)
 - [PluginConfigUser](docs/PluginConfigUser.md)
 - [PluginDevice](docs/PluginDevice.md)
 - [PluginEnv](docs/PluginEnv.md)
 - [PluginInterfaceType](docs/PluginInterfaceType.md)
 - [PluginMount](docs/PluginMount.md)
 - [PluginSettings](docs/PluginSettings.md)
 - [Port](docs/Port.md)
 - [RecoveryAddress](docs/RecoveryAddress.md)
 - [RecoveryFlow](docs/RecoveryFlow.md)
 - [RecoveryLink](docs/RecoveryLink.md)
 - [RegistrationFlow](docs/RegistrationFlow.md)
 - [RegistrationViaApiResponse](docs/RegistrationViaApiResponse.md)
 - [RevokeSession](docs/RevokeSession.md)
 - [SelfServiceErrorContainer](docs/SelfServiceErrorContainer.md)
 - [ServiceUpdateResponse](docs/ServiceUpdateResponse.md)
 - [Session](docs/Session.md)
 - [SettingsFlow](docs/SettingsFlow.md)
 - [SettingsProfileFormConfig](docs/SettingsProfileFormConfig.md)
 - [SettingsViaApiResponse](docs/SettingsViaApiResponse.md)
 - [SubmitSelfServiceBrowserSettingsOIDCFlowPayload](docs/SubmitSelfServiceBrowserSettingsOIDCFlowPayload.md)
 - [SubmitSelfServiceLoginFlow](docs/SubmitSelfServiceLoginFlow.md)
 - [SubmitSelfServiceLoginFlowWithPasswordMethod](docs/SubmitSelfServiceLoginFlowWithPasswordMethod.md)
 - [SubmitSelfServiceRecoveryFlowWithLinkMethod](docs/SubmitSelfServiceRecoveryFlowWithLinkMethod.md)
 - [SubmitSelfServiceRegistrationFlow](docs/SubmitSelfServiceRegistrationFlow.md)
 - [SubmitSelfServiceRegistrationFlowWithPasswordMethod](docs/SubmitSelfServiceRegistrationFlowWithPasswordMethod.md)
 - [SubmitSelfServiceSettingsFlow](docs/SubmitSelfServiceSettingsFlow.md)
 - [SubmitSelfServiceSettingsFlowWithPasswordMethod](docs/SubmitSelfServiceSettingsFlowWithPasswordMethod.md)
 - [SubmitSelfServiceSettingsFlowWithProfileMethod](docs/SubmitSelfServiceSettingsFlowWithProfileMethod.md)
 - [SubmitSelfServiceVerificationFlowWithLinkMethod](docs/SubmitSelfServiceVerificationFlowWithLinkMethod.md)
 - [UiContainer](docs/UiContainer.md)
 - [UiNode](docs/UiNode.md)
 - [UiNodeAnchorAttributes](docs/UiNodeAnchorAttributes.md)
 - [UiNodeAttributes](docs/UiNodeAttributes.md)
 - [UiNodeImageAttributes](docs/UiNodeImageAttributes.md)
 - [UiNodeInputAttributes](docs/UiNodeInputAttributes.md)
 - [UiNodeTextAttributes](docs/UiNodeTextAttributes.md)
 - [UiText](docs/UiText.md)
 - [UpdateIdentity](docs/UpdateIdentity.md)
 - [VerifiableIdentityAddress](docs/VerifiableIdentityAddress.md)
 - [VerificationFlow](docs/VerificationFlow.md)
 - [Version](docs/Version.md)
 - [Volume](docs/Volume.md)
 - [VolumeUsageData](docs/VolumeUsageData.md)


## Documentation For Authorization



### oryToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### sessionCookie

- **Type**: API key
- **API key parameter name**: ory_kratos_session
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ory_kratos_session and passed in as the auth context for each request.


### sessionToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
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

support@ory.sh

