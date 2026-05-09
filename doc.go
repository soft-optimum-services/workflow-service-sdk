// Package workflow is an HTTP client for workflow-service (template admin + execution APIs).
//
// Domain types live in github.com/soft-optimum-services/workflow-service-sdk/model and mirror
// the JSON payloads returned by workflow-service. When consumer apps keep duplicate structs
// under their own internal packages, Go forbids importing those from another module—types here
// stay standalone. Keep JSON tags in sync when the API model evolves.
package workflow
