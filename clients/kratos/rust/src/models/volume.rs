/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// Volume : Volume volume



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Volume {
    /// Date/Time the volume was created.
    #[serde(rename = "CreatedAt", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    /// Name of the volume driver used by the volume.
    #[serde(rename = "Driver")]
    pub driver: String,
    /// User-defined key/value metadata.
    #[serde(rename = "Labels")]
    pub labels: ::std::collections::HashMap<String, String>,
    /// Mount path of the volume on the host.
    #[serde(rename = "Mountpoint")]
    pub mountpoint: String,
    /// Name of the volume.
    #[serde(rename = "Name")]
    pub name: String,
    /// The driver specific options used when creating the volume.
    #[serde(rename = "Options")]
    pub options: ::std::collections::HashMap<String, String>,
    /// The level at which the volume exists. Either `global` for cluster-wide, or `local` for machine level.
    #[serde(rename = "Scope")]
    pub scope: String,
    /// Low-level details about the volume, provided by the volume driver. Details are returned as a map with key/value pairs: `{\"key\":\"value\",\"key2\":\"value2\"}`.  The `Status` field is optional, and is omitted if the volume driver does not support this feature.
    #[serde(rename = "Status", skip_serializing_if = "Option::is_none")]
    pub status: Option<::std::collections::HashMap<String, serde_json::Value>>,
    #[serde(rename = "UsageData", skip_serializing_if = "Option::is_none")]
    pub usage_data: Option<Box<crate::models::VolumeUsageData>>,
}

impl Volume {
    /// Volume volume
    pub fn new(driver: String, labels: ::std::collections::HashMap<String, String>, mountpoint: String, name: String, options: ::std::collections::HashMap<String, String>, scope: String) -> Volume {
        Volume {
            created_at: None,
            driver,
            labels,
            mountpoint,
            name,
            options,
            scope,
            status: None,
            usage_data: None,
        }
    }
}


