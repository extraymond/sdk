/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ContainerUpdateOkBody : ContainerUpdateOKBody OK response to ContainerUpdate operation



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContainerUpdateOkBody {
    /// warnings
    #[serde(rename = "Warnings")]
    pub warnings: Vec<String>,
}

impl ContainerUpdateOkBody {
    /// ContainerUpdateOKBody OK response to ContainerUpdate operation
    pub fn new(warnings: Vec<String>) -> ContainerUpdateOkBody {
        ContainerUpdateOkBody {
            warnings,
        }
    }
}


