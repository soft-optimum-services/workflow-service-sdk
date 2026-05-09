package model

import (
	"time"

	"github.com/google/uuid"
)

// Template represents a workflow template
type Template struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	Code            string         `json:"code" db:"code"`
	Name            string         `json:"name" db:"name"`
	Description     *string        `json:"description,omitempty" db:"description"`
	Version         string         `json:"version" db:"version"`
	BusinessDomain  string         `json:"business_domain" db:"business_domain"`
	Status          WorkflowStatus `json:"status" db:"status"`
	IsCurrent       bool           `json:"is_current" db:"is_current"`
	ParentVersionID *uuid.UUID     `json:"parent_version_id,omitempty" db:"parent_version_id"`
	ValidFrom       *time.Time     `json:"valid_from,omitempty" db:"valid_from"`
	ValidUntil      *time.Time     `json:"valid_until,omitempty" db:"valid_until"`
	Metadata        map[string]any `json:"metadata" db:"metadata"`
	CreatedBy       string         `json:"created_by" db:"created_by"`
	UpdatedBy       string         `json:"updated_by" db:"updated_by"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
	States          []State        `json:"states,omitempty"`
	Transitions     []Transition   `json:"transitions,omitempty"`
}

// BusinessDomain represents a business domain configuration
type BusinessDomain struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
