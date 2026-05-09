package workflow

// CanTransitionRequest is the payload for POST /api/v1/execution/can-transition.
type CanTransitionRequest struct {
	EntityType     string         `json:"entity_type"`
	EntityID       int64          `json:"entity_id"`
	TransitionCode string         `json:"transition_code"`
	UserID         string         `json:"user_id"`
	UserRole       string         `json:"user_role,omitempty"`
	Context        map[string]any `json:"context,omitempty"`
}

// CanTransitionResponse is the response from POST /api/v1/execution/can-transition.
type CanTransitionResponse struct {
	Allowed         bool     `json:"allowed"`
	Reason          string   `json:"reason,omitempty"`
	RequiredActions []string `json:"required_actions,omitempty"`
	Warnings        []string `json:"warnings,omitempty"`
}

// ExecuteTransitionRequest is the payload for POST /api/v1/execution/execute-transition.
type ExecuteTransitionRequest struct {
	EntityType     string         `json:"entity_type"`
	EntityID       int64          `json:"entity_id"`
	TransitionCode string         `json:"transition_code"`
	UserID         string         `json:"user_id"`
	UserRole       string         `json:"user_role,omitempty"`
	Comment        string         `json:"comment,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Context        map[string]any `json:"context,omitempty"`
}

// ExecuteTransitionResponse is the response from POST /api/v1/execution/execute-transition.
type ExecuteTransitionResponse struct {
	Success          bool     `json:"success"`
	NewState         string   `json:"new_state"`
	NewStatus        string   `json:"new_status,omitempty"`
	ExecutionID      string   `json:"execution_id"`
	ActionsTriggered []string `json:"actions_triggered,omitempty"`
	ErrorMessage     string   `json:"error_message,omitempty"`
}

// GetStateResponse is the response from GET /api/v1/execution/state/{entityType}/{entityId}.
type GetStateResponse struct {
	CurrentState         string   `json:"current_state"`
	CurrentStatus        string   `json:"current_status,omitempty"`
	TemplateCode         string   `json:"template_code"`
	TemplateVersion      string   `json:"template_version"`
	AvailableTransitions []string `json:"available_transitions"`
	StateEnteredAt       string   `json:"state_entered_at"`
	DaysInState          int      `json:"days_in_state"`
}

// HistoryEntry represents one transition history item.
type HistoryEntry struct {
	ID         string `json:"id"`
	FromState  string `json:"from_state"`
	FromStatus string `json:"from_status,omitempty"`
	ToState    string `json:"to_state"`
	ToStatus   string `json:"to_status,omitempty"`
	Transition string `json:"transition"`
	ExecutedBy string `json:"executed_by"`
	ExecutedAt string `json:"executed_at"`
	Comment    string `json:"comment,omitempty"`
	Status     string `json:"status"`
}

// GetHistoryResponse is the response from GET /api/v1/execution/history/{entityType}/{entityId}.
type GetHistoryResponse struct {
	Transitions []HistoryEntry `json:"transitions"`
	Total       int            `json:"total"`
}

// BindEntityRequest is the payload for POST /api/v1/execution/bind.
type BindEntityRequest struct {
	TemplateCode string         `json:"template_code"`
	EntityType   string         `json:"entity_type"`
	EntityID     int64          `json:"entity_id"`
	InitialState string         `json:"initial_state,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

// BindEntityResponse is the response from POST /api/v1/execution/bind.
type BindEntityResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AvailableTransition contains exposed transition fields for execution context.
type AvailableTransition struct {
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	ConfirmationRequired bool   `json:"confirmation_required"`
	ConfirmationMessage  string `json:"confirmation_message,omitempty"`
}

// AvailableTransitionsResponse is the response from GET /api/v1/execution/transitions/{entityType}/{entityId}.
type AvailableTransitionsResponse struct {
	Transitions []AvailableTransition `json:"transitions"`
}
