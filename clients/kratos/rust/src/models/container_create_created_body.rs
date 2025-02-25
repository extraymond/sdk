/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ContainerCreateCreatedBody : ContainerCreateCreatedBody OK response to ContainerCreate operation



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContainerCreateCreatedBody {
    /// The ID of the created container
    #[serde(rename = "Id")]
    pub id: String,
    /// Warnings encountered when creating the container
    #[serde(rename = "Warnings")]
    pub warnings: Vec<String>,
}

impl ContainerCreateCreatedBody {
    /// ContainerCreateCreatedBody OK response to ContainerCreate operation
    pub fn new(id: String, warnings: Vec<String>) -> ContainerCreateCreatedBody {
        ContainerCreateCreatedBody {
            id,
            warnings,
        }
    }
}


