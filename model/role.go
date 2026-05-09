package model

import (
	"time"

	"github.com/google/uuid"
)

// Role represents a workflow role
type Role struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsSystem    bool      `json:"is_system" db:"is_system"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// StateRole represents role permissions on a state
type StateRole struct {
	ID        uuid.UUID `json:"id" db:"id"`
	StateID   uuid.UUID `json:"state_id" db:"state_id"`
	RoleID    uuid.UUID `json:"role_id" db:"role_id"`
	CanView   bool      `json:"can_view" db:"can_view"`
	CanAct    bool      `json:"can_act" db:"can_act"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Role      *Role     `json:"role,omitempty"`
}

// TransitionRole represents role permissions on a transition
type TransitionRole struct {
	ID           uuid.UUID      `json:"id" db:"id"`
	TransitionID uuid.UUID      `json:"transition_id" db:"transition_id"`
	RoleID       uuid.UUID      `json:"role_id" db:"role_id"`
	CanExecute   bool           `json:"can_execute" db:"can_execute"`
	CanOverride  bool           `json:"can_override" db:"can_override"`
	Conditions   map[string]any `json:"conditions" db:"conditions"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	Role         *Role          `json:"role,omitempty"`
}
