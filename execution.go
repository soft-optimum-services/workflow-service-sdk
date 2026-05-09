package workflow

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// CanTransition calls POST /api/v1/execution/can-transition.
func (c *Client) CanTransition(ctx context.Context, req CanTransitionRequest) (*CanTransitionResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	body, err := c.postJSON(ctx, root.JoinPath("api", "v1", "execution", "can-transition").String(), req, http.StatusOK)
	if err != nil {
		return nil, err
	}

	var payload CanTransitionResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// ExecuteTransition calls POST /api/v1/execution/execute-transition.
func (c *Client) ExecuteTransition(ctx context.Context, req ExecuteTransitionRequest) (*ExecuteTransitionResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	body, err := c.postJSON(ctx, root.JoinPath("api", "v1", "execution", "execute-transition").String(), req, http.StatusOK)
	if err != nil {
		return nil, err
	}

	var payload ExecuteTransitionResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// GetState calls GET /api/v1/execution/state/{entityType}/{entityId}.
func (c *Client) GetState(ctx context.Context, entityType string, entityID int64) (*GetStateResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	body, err := c.get(ctx, root.JoinPath("api", "v1", "execution", "state", entityType, strconv.FormatInt(entityID, 10)).String())
	if err != nil {
		return nil, err
	}

	var payload GetStateResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// GetHistory calls GET /api/v1/execution/history/{entityType}/{entityId}.
func (c *Client) GetHistory(ctx context.Context, entityType string, entityID int64, limit, offset int) (*GetHistoryResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	u := root.JoinPath("api", "v1", "execution", "history", entityType, strconv.FormatInt(entityID, 10))
	query := u.Query()
	if limit > 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	if offset >= 0 {
		query.Set("offset", strconv.Itoa(offset))
	}
	u.RawQuery = query.Encode()

	body, err := c.get(ctx, u.String())
	if err != nil {
		return nil, err
	}

	var payload GetHistoryResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// BindEntity calls POST /api/v1/execution/bind.
func (c *Client) BindEntity(ctx context.Context, req BindEntityRequest) (*BindEntityResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	body, err := c.postJSON(ctx, root.JoinPath("api", "v1", "execution", "bind").String(), req, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	var payload BindEntityResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// GetAvailableTransitions calls GET /api/v1/execution/transitions/{entityType}/{entityId}?role={role}.
func (c *Client) GetAvailableTransitions(ctx context.Context, entityType string, entityID int64, role string) (*AvailableTransitionsResponse, error) {
	root, err := c.rootURL()
	if err != nil {
		return nil, err
	}

	u := root.JoinPath("api", "v1", "execution", "transitions", entityType, strconv.FormatInt(entityID, 10))
	if role != "" {
		query := u.Query()
		query.Set("role", role)
		u.RawQuery = query.Encode()
	}

	body, err := c.get(ctx, u.String())
	if err != nil {
		return nil, err
	}

	var payload AvailableTransitionsResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
