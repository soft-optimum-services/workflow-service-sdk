package model

import (
	"time"

	"github.com/google/uuid"
)

// Transition represents a workflow transition
type Transition struct {
	ID                   uuid.UUID         `json:"id" db:"id"`
	TemplateID           uuid.UUID         `json:"template_id" db:"template_id"`
	Code                 string            `json:"code" db:"code"`
	Name                 string            `json:"name" db:"name"`
	Description          *string           `json:"description,omitempty" db:"description"`
	FromStateID          uuid.UUID         `json:"from_state_id" db:"from_state_id"`
	FromStatusID         *uuid.UUID        `json:"from_status_id,omitempty" db:"from_status_id"`
	ToStateID            uuid.UUID         `json:"to_state_id" db:"to_state_id"`
	ToStatusID           *uuid.UUID        `json:"to_status_id,omitempty" db:"to_status_id"`
	TriggerType          TransitionTrigger `json:"trigger_type" db:"trigger_type"`
	Priority             int               `json:"priority" db:"priority"`
	IsEnabled            bool              `json:"is_enabled" db:"is_enabled"`
	RequiresComment      bool              `json:"requires_comment" db:"requires_comment"`
	ConfirmationRequired bool              `json:"confirmation_required" db:"confirmation_required"`
	ConfirmationMessage  *string           `json:"confirmation_message,omitempty" db:"confirmation_message"`
	Metadata             map[string]any    `json:"metadata" db:"metadata"`
	CreatedAt            time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time         `json:"updated_at" db:"updated_at"`
	FromState            *State            `json:"from_state,omitempty"`
	ToState              *State            `json:"to_state,omitempty"`
	Conditions           []Condition       `json:"conditions,omitempty"`
	Actions              []Action          `json:"actions,omitempty"`
	Roles                []TransitionRole  `json:"roles,omitempty"`
}

// Condition represents a transition guard/condition
type Condition struct {
	ID             uuid.UUID         `json:"id" db:"id"`
	TransitionID   uuid.UUID         `json:"transition_id" db:"transition_id"`
	Name           string            `json:"name" db:"name"`
	Description    *string           `json:"description,omitempty" db:"description"`
	ConditionType  ConditionType     `json:"condition_type" db:"condition_type"`
	Field          *string           `json:"field,omitempty" db:"field_path"`
	FieldPath      *string           `json:"field_path,omitempty" db:"field_path"`
	Operator       ConditionOperator `json:"operator,omitempty" db:"operator"`
	ExpectedValue  *string           `json:"expected_value,omitempty" db:"expected_value"`
	Expression     *string           `json:"expression,omitempty" db:"expression"`
	ErrorMessage   *string           `json:"error_message,omitempty" db:"error_message"`
	IsMandatory    bool              `json:"is_mandatory" db:"is_mandatory"`
	ExecutionOrder int               `json:"execution_order" db:"execution_order"`
	IsEnabled      bool              `json:"is_enabled" db:"is_enabled"`
	Metadata       map[string]any    `json:"metadata" db:"metadata"`
	CreatedAt      time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at" db:"updated_at"`
}

// Action represents a post-transition action
type Action struct {
	ID                uuid.UUID      `json:"id" db:"id"`
	TransitionID      uuid.UUID      `json:"transition_id" db:"transition_id"`
	Name              string         `json:"name" db:"name"`
	Description       *string        `json:"description,omitempty" db:"description"`
	ActionType        ActionType     `json:"action_type" db:"action_type"`
	Timing            ActionTiming   `json:"timing" db:"timing"`
	Config            map[string]any   `json:"config" db:"config"`
	ExecutionOrder    int            `json:"execution_order" db:"execution_order"`
	IsEnabled         bool           `json:"is_enabled" db:"is_enabled"`
	IsAsync           bool           `json:"is_async" db:"is_async"`
	RetryCount        int            `json:"retry_count" db:"retry_count"`
	RetryDelaySeconds int            `json:"retry_delay_seconds" db:"retry_delay_seconds"`
	TimeoutSeconds    int            `json:"timeout_seconds" db:"timeout_seconds"`
	OnFailure         string         `json:"on_failure" db:"on_failure"`
	Metadata          map[string]any `json:"metadata" db:"metadata"`
	CreatedAt         time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at" db:"updated_at"`
}
