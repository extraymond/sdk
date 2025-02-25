/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
#[serde(tag = "method")]
pub enum SubmitSelfServiceLoginFlowBody {
    #[serde(rename="oidc")]
    SubmitSelfServiceLoginFlowWithOidcMethodBody {
        /// The CSRF Token
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The provider to register with
        #[serde(rename = "traits")]
        traits: String,
    },
    #[serde(rename="password")]
    SubmitSelfServiceLoginFlowWithPasswordMethodBody {
        /// Sending the anti-csrf token is only required for browser login flows.
        #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
        csrf_token: Option<String>,
        /// The user's password.
        #[serde(rename = "password")]
        password: String,
        /// Identifier is the email or username of the user trying to log in.
        #[serde(rename = "password_identifier")]
        password_identifier: String,
    },
}



/// Method should be set to \"password\" when logging in using the identifier and password strategy.
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Method {
    #[serde(rename = "password")]
    Password,
    #[serde(rename = "oidc")]
    Oidc,
}

