/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// Meta : This might include a label and other information that can optionally be used to render UIs.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Meta {
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<Box<crate::models::UiText>>,
}

impl Meta {
    /// This might include a label and other information that can optionally be used to render UIs.
    pub fn new() -> Meta {
        Meta {
            label: None,
        }
    }
}


