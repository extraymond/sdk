/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ImageSummary : ImageSummary image summary



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ImageSummary {
    /// containers
    #[serde(rename = "Containers")]
    pub containers: i64,
    /// created
    #[serde(rename = "Created")]
    pub created: i64,
    /// Id
    #[serde(rename = "Id")]
    pub id: String,
    /// labels
    #[serde(rename = "Labels")]
    pub labels: ::std::collections::HashMap<String, String>,
    /// parent Id
    #[serde(rename = "ParentId")]
    pub parent_id: String,
    /// repo digests
    #[serde(rename = "RepoDigests")]
    pub repo_digests: Vec<String>,
    /// repo tags
    #[serde(rename = "RepoTags")]
    pub repo_tags: Vec<String>,
    /// shared size
    #[serde(rename = "SharedSize")]
    pub shared_size: i64,
    /// size
    #[serde(rename = "Size")]
    pub size: i64,
    /// virtual size
    #[serde(rename = "VirtualSize")]
    pub virtual_size: i64,
}

impl ImageSummary {
    /// ImageSummary image summary
    pub fn new(containers: i64, created: i64, id: String, labels: ::std::collections::HashMap<String, String>, parent_id: String, repo_digests: Vec<String>, repo_tags: Vec<String>, shared_size: i64, size: i64, virtual_size: i64) -> ImageSummary {
        ImageSummary {
            containers,
            created,
            id,
            labels,
            parent_id,
            repo_digests,
            repo_tags,
            shared_size,
            size,
            virtual_size,
        }
    }
}


