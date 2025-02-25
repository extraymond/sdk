/*
 * ORY Keto
 *
 * Ory Keto is a cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.
 *
 * The version of the OpenAPI document: v0.6.0-alpha.5
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct InlineResponse400 {
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<i64>,
    #[serde(rename = "details", skip_serializing_if = "Option::is_none")]
    pub details: Option<Vec<serde_json::Value>>,
    #[serde(rename = "message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<String>,
    #[serde(rename = "request", skip_serializing_if = "Option::is_none")]
    pub request: Option<String>,
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
}

impl InlineResponse400 {
    pub fn new() -> InlineResponse400 {
        InlineResponse400 {
            code: None,
            details: None,
            message: None,
            reason: None,
            request: None,
            status: None,
        }
    }
}


