/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginInterfaceType : PluginInterfaceType plugin interface type



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginInterfaceType {
    /// capability
    #[serde(rename = "Capability")]
    pub capability: String,
    /// prefix
    #[serde(rename = "Prefix")]
    pub prefix: String,
    /// version
    #[serde(rename = "Version")]
    pub version: String,
}

impl PluginInterfaceType {
    /// PluginInterfaceType plugin interface type
    pub fn new(capability: String, prefix: String, version: String) -> PluginInterfaceType {
        PluginInterfaceType {
            capability,
            prefix,
            version,
        }
    }
}


