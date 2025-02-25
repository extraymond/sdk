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
pub struct UiNodeAttributes {
    /// Sets the input's disabled field to true or false.
    #[serde(rename = "disabled")]
    pub disabled: bool,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<Box<crate::models::UiText>>,
    /// The input's element name.
    #[serde(rename = "name")]
    pub name: String,
    /// The input's pattern.
    #[serde(rename = "pattern", skip_serializing_if = "Option::is_none")]
    pub pattern: Option<String>,
    /// Mark this input field as required.
    #[serde(rename = "required", skip_serializing_if = "Option::is_none")]
    pub required: Option<bool>,
    #[serde(rename = "type")]
    pub _type: String,
    /// The input's value.
    #[serde(rename = "value", skip_serializing_if = "Option::is_none")]
    pub value: Option<serde_json::Value>,
    #[serde(rename = "text")]
    pub text: Box<crate::models::UiText>,
    /// The image's source URL.  format: uri
    #[serde(rename = "src")]
    pub src: String,
    /// The link's href (destination) URL.  format: uri
    #[serde(rename = "href")]
    pub href: String,
    #[serde(rename = "title")]
    pub title: Box<crate::models::UiText>,
}

impl UiNodeAttributes {
    pub fn new(disabled: bool, name: String, _type: String, text: crate::models::UiText, src: String, href: String, title: crate::models::UiText) -> UiNodeAttributes {
        UiNodeAttributes {
            disabled,
            label: None,
            name,
            pattern: None,
            required: None,
            _type,
            value: None,
            text: Box::new(text),
            src,
            href,
            title: Box::new(title),
        }
    }
}


