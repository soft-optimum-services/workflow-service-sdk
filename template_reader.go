package workflow

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	wfmodel "github.com/soft-optimum-services/workflow-service-sdk/model"
)

type statesEnvelope struct {
	States []wfmodel.State `json:"states"`
}

type transitionsEnvelope struct {
	Transitions []wfmodel.Transition `json:"transitions"`
}

// ListStates fetches workflow states for a template code:
// GET {baseURL}/api/v1/templates/by-code/{templateCode}/states.
func (c *Client) ListStates(ctx context.Context, templateCode string) ([]wfmodel.State, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}
	code := strings.TrimSpace(templateCode)
	if code == "" {
		return nil, fmt.Errorf("template code is required")
	}

	body, err := c.get(ctx, root.JoinPath("api", "v1", "templates", "by-code", code, "states").String())
	if err != nil {
		return nil, err
	}

	var payload statesEnvelope
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return payload.States, nil
}

// ListTransitions fetches workflow transitions for a template code:
// GET {baseURL}/api/v1/templates/by-code/{templateCode}/transitions.
func (c *Client) ListTransitions(ctx context.Context, templateCode string) ([]wfmodel.Transition, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}
	code := strings.TrimSpace(templateCode)
	if code == "" {
		return nil, fmt.Errorf("template code is required")
	}

	body, err := c.get(ctx, root.JoinPath("api", "v1", "templates", "by-code", code, "transitions").String())
	if err != nil {
		return nil, err
	}

	var payload transitionsEnvelope
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return payload.Transitions, nil
}

// GetTemplateByCode fetches a workflow template by code:
// GET {baseURL}/api/v1/templates/by-code/{templateCode}.
func (c *Client) GetTemplateByCode(ctx context.Context, templateCode string) (*wfmodel.Template, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}
	code := strings.TrimSpace(templateCode)
	if code == "" {
		return nil, fmt.Errorf("template code is required")
	}

	body, err := c.get(ctx, root.JoinPath("api", "v1", "templates", "by-code", code).String())
	if err != nil {
		return nil, err
	}

	var tpl wfmodel.Template
	if err := json.Unmarshal(body, &tpl); err != nil {
		return nil, err
	}
	return &tpl, nil
}
