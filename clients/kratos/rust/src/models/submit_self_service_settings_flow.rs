/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceSettingsFlow {
    /// The Anti-CSRF Token  This token is only required when performing browser flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to profile when trying to update a profile.  type: string
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<String>,
    /// Password is the updated password  type: string
    #[serde(rename = "password")]
    pub password: String,
    /// Traits contains all of the identity's traits.
    #[serde(rename = "traits")]
    pub traits: serde_json::Value,
}

impl SubmitSelfServiceSettingsFlow {
    pub fn new(password: String, traits: serde_json::Value) -> SubmitSelfServiceSettingsFlow {
        SubmitSelfServiceSettingsFlow {
            csrf_token: None,
            method: None,
            password,
            traits,
        }
    }
}


