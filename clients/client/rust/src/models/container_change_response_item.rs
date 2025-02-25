/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ContainerChangeResponseItem : ContainerChangeResponseItem change item in response to ContainerChanges operation



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContainerChangeResponseItem {
    /// Kind of change
    #[serde(rename = "Kind")]
    pub kind: i32,
    /// Path to file that has changed
    #[serde(rename = "Path")]
    pub path: String,
}

impl ContainerChangeResponseItem {
    /// ContainerChangeResponseItem change item in response to ContainerChanges operation
    pub fn new(kind: i32, path: String) -> ContainerChangeResponseItem {
        ContainerChangeResponseItem {
            kind,
            path,
        }
    }
}


