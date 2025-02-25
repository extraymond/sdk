/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginMount : PluginMount plugin mount



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginMount {
    /// description
    #[serde(rename = "Description")]
    pub description: String,
    /// destination
    #[serde(rename = "Destination")]
    pub destination: String,
    /// name
    #[serde(rename = "Name")]
    pub name: String,
    /// options
    #[serde(rename = "Options")]
    pub options: Vec<String>,
    /// settable
    #[serde(rename = "Settable")]
    pub settable: Vec<String>,
    /// source
    #[serde(rename = "Source")]
    pub source: String,
    /// type
    #[serde(rename = "Type")]
    pub _type: String,
}

impl PluginMount {
    /// PluginMount plugin mount
    pub fn new(description: String, destination: String, name: String, options: Vec<String>, settable: Vec<String>, source: String, _type: String) -> PluginMount {
        PluginMount {
            description,
            destination,
            name,
            options,
            settable,
            source,
            _type,
        }
    }
}


