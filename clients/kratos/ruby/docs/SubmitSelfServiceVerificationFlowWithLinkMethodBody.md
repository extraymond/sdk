# OryHydraClient::SubmitSelfServiceVerificationFlowWithLinkMethodBody

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **csrf_token** | **String** | Sending the anti-csrf token is only required for browser login flows. | [optional] |
| **email** | **String** | Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email |  |
| **method** | **String** | Method supports &#x60;link&#x60; only right now. |  |

## Example

```ruby
require 'ory-kratos-client'

instance = OryHydraClient::SubmitSelfServiceVerificationFlowWithLinkMethodBody.new(
  csrf_token: null,
  email: null,
  method: null
)
```

