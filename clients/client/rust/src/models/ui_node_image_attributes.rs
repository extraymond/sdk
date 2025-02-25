/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UiNodeImageAttributes {
    /// The image's source URL.  format: uri
    #[serde(rename = "src")]
    pub src: String,
}

impl UiNodeImageAttributes {
    pub fn new(src: String) -> UiNodeImageAttributes {
        UiNodeImageAttributes {
            src,
        }
    }
}


