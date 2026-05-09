package workflow

import (
	"context"

	wfmodel "github.com/soft-optimum-services/workflow-service-sdk/model"
)

// TemplateReader exposes admin/template endpoints that return strongly typed models.
type TemplateReader interface {
	ListStates(ctx context.Context, templateCode string) ([]wfmodel.State, error)
	ListTransitions(ctx context.Context, templateCode string) ([]wfmodel.Transition, error)
	GetTemplateByCode(ctx context.Context, templateCode string) (*wfmodel.Template, error)
}

// ExecutionClient exposes workflow execution endpoints (entity lifecycle).
type ExecutionClient interface {
	CanTransition(ctx context.Context, req CanTransitionRequest) (*CanTransitionResponse, error)
	ExecuteTransition(ctx context.Context, req ExecuteTransitionRequest) (*ExecuteTransitionResponse, error)
	GetState(ctx context.Context, entityType string, entityID int64) (*GetStateResponse, error)
	GetHistory(ctx context.Context, entityType string, entityID int64, limit, offset int) (*GetHistoryResponse, error)
	BindEntity(ctx context.Context, req BindEntityRequest) (*BindEntityResponse, error)
	GetAvailableTransitions(ctx context.Context, entityType string, entityID int64, role string) (*AvailableTransitionsResponse, error)
}

// WorkflowQueries provides derived operations on template states/transitions for UI or orchestration.
type WorkflowQueries interface {
	ListManualTransitionsFromState(ctx context.Context, templateCode, fromStateCode string) []wfmodel.Transition
	ResolveTransitionTargetStateCode(ctx context.Context, templateCode, fromStateCode, transitionCode string) (string, error)
}

// WorkflowClient is the full workflow-service client contract.
//
// Prefer depending on the smaller interfaces (TemplateReader, ExecutionClient, WorkflowQueries)
// when a consumer only needs a subset of functionality.
type WorkflowClient interface {
	TemplateReader
	ExecutionClient
	WorkflowQueries
}

var (
	_ WorkflowClient = (*Client)(nil)
	_ TemplateReader  = (*Client)(nil)
	_ ExecutionClient = (*Client)(nil)
	_ WorkflowQueries = (*Client)(nil)
)
