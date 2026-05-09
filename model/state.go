package model

import (
	"time"

	"github.com/google/uuid"
)

// State represents a workflow state
type State struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	TemplateID       uuid.UUID       `json:"template_id" db:"template_id"`
	Code             string          `json:"code" db:"code"`
	Name             string          `json:"name" db:"name"`
	Description      *string         `json:"description,omitempty" db:"description"`
	StateType        StateType       `json:"state_type" db:"state_type"`
	IsInitial        bool            `json:"is_initial" db:"is_initial"`
	IsFinal          bool            `json:"is_final" db:"is_final"`
	Color            string          `json:"color" db:"color"`
	Icon             *string         `json:"icon,omitempty" db:"icon"`
	Position         int             `json:"position" db:"position"`
	PositionX        int             `json:"position_x" db:"position_x"`
	PositionY        int             `json:"position_y" db:"position_y"`
	MinRetentionDays *int            `json:"min_retention_days,omitempty" db:"min_retention_days"`
	MaxRetentionDays *int            `json:"max_retention_days,omitempty" db:"max_retention_days"`
	RetentionStatus  RetentionStatus `json:"retention_status" db:"retention_status"`
	AlertOnMinDays   bool            `json:"alert_on_min_days" db:"alert_on_min_days"`
	AlertOnMaxDays   bool            `json:"alert_on_max_days" db:"alert_on_max_days"`
	EscalationRole   *string         `json:"escalation_role,omitempty" db:"escalation_role"`
	RequiredFields   []string        `json:"required_fields" db:"required_fields"`
	Metadata         map[string]any  `json:"metadata" db:"metadata"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
	Statuses         []StateStatus   `json:"statuses,omitempty"`
}

// StateStatus represents a sub-status within a state
type StateStatus struct {
	ID           uuid.UUID `json:"id" db:"id"`
	StateID      uuid.UUID `json:"state_id" db:"state_id"`
	Code         string    `json:"code" db:"code"`
	Name         string    `json:"name" db:"name"`
	Description  *string   `json:"description,omitempty" db:"description"`
	IsDefault    bool      `json:"is_default" db:"is_default"`
	IsTerminal   bool      `json:"is_terminal" db:"is_terminal"`
	Color        string    `json:"color" db:"color"`
	DisplayOrder int       `json:"display_order" db:"display_order"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
