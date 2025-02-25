/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginConfigNetwork : PluginConfigNetwork plugin config network



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginConfigNetwork {
    /// type
    #[serde(rename = "Type")]
    pub _type: String,
}

impl PluginConfigNetwork {
    /// PluginConfigNetwork plugin config network
    pub fn new(_type: String) -> PluginConfigNetwork {
        PluginConfigNetwork {
            _type,
        }
    }
}


