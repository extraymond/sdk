/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateIdentity {
    /// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
    #[serde(rename = "schema_id")]
    pub schema_id: String,
    /// Traits represent an identity's traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in `schema_url`.
    #[serde(rename = "traits")]
    pub traits: serde_json::Value,
}

impl CreateIdentity {
    pub fn new(schema_id: String, traits: serde_json::Value) -> CreateIdentity {
        CreateIdentity {
            schema_id,
            traits,
        }
    }
}


