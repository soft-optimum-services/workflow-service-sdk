package model

// WorkflowStatus represents the lifecycle status of a workflow template
type WorkflowStatus string

const (
	WorkflowStatusDraft      WorkflowStatus = "draft"
	WorkflowStatusActive     WorkflowStatus = "active"
	WorkflowStatusDeprecated WorkflowStatus = "deprecated"
	WorkflowStatusArchived   WorkflowStatus = "archived"
)

// StateType represents the type of a workflow state
type StateType string

const (
	StateTypeInitial      StateType = "initial"
	StateTypeIntermediate StateType = "intermediate"
	StateTypeTerminal     StateType = "terminal"
	StateTypeError        StateType = "error"
)

// RetentionStatus represents the retention status of a state
type RetentionStatus string

const (
	RetentionStatusNormal    RetentionStatus = "normal"
	RetentionStatusWarning   RetentionStatus = "warning"
	RetentionStatusCritical  RetentionStatus = "critical"
	RetentionStatusBlocked   RetentionStatus = "blocked"
	RetentionStatusEscalated RetentionStatus = "escalated"
)

// TransitionTrigger represents how a transition is triggered
type TransitionTrigger string

// TriggerType is an alias for TransitionTrigger
type TriggerType = TransitionTrigger

const (
	TransitionTriggerManual      TransitionTrigger = "manual"
	TransitionTriggerAutomatic   TransitionTrigger = "automatic"
	TransitionTriggerScheduled   TransitionTrigger = "scheduled"
	TransitionTriggerConditional TransitionTrigger = "conditional"
	TransitionTriggerExternal    TransitionTrigger = "external"
)

// ConditionType represents the type of condition
type ConditionType string

const (
	ConditionTypeRole        ConditionType = "role"
	ConditionTypeField       ConditionType = "field"
	ConditionTypeTime        ConditionType = "time"
	ConditionTypeExpression  ConditionType = "expression"
	ConditionTypeExternalAPI ConditionType = "external_api"
	ConditionTypeCustom      ConditionType = "custom"
)

// ConditionOperator represents comparison operators
type ConditionOperator string

const (
	OperatorEquals               ConditionOperator = "equals"
	OperatorNotEquals            ConditionOperator = "not_equals"
	OperatorGreaterThan          ConditionOperator = "greater_than"
	OperatorGreaterThanOrEquals ConditionOperator = "greater_than_or_equals"
	OperatorLessThan             ConditionOperator = "less_than"
	OperatorLessThanOrEquals     ConditionOperator = "less_than_or_equals"
	OperatorContains             ConditionOperator = "contains"
	OperatorNotContains          ConditionOperator = "not_contains"
	OperatorStartsWith           ConditionOperator = "starts_with"
	OperatorEndsWith             ConditionOperator = "ends_with"
	OperatorIn                   ConditionOperator = "in"
	OperatorNotIn                ConditionOperator = "not_in"
	OperatorIsNull               ConditionOperator = "is_null"
	OperatorIsNotNull            ConditionOperator = "is_not_null"
	OperatorBetween              ConditionOperator = "between"
	OperatorRegex                ConditionOperator = "regex"
)

// ActionType represents the type of action
type ActionType string

const (
	ActionTypeNotificationEmail ActionType = "notification_email"
	ActionTypeNotificationSMS   ActionType = "notification_sms"
	ActionTypeNotificationPush  ActionType = "notification_push"
	ActionTypeWebhook           ActionType = "webhook"
	ActionTypeFieldUpdate       ActionType = "field_update"
	ActionTypeDocumentGen       ActionType = "document_generation"
	ActionTypeTaskCreation      ActionType = "task_creation"
	ActionTypeCustomScript      ActionType = "custom_script"
)

// ActionTiming represents when an action is executed
type ActionTiming string

const (
	ActionTimingBefore ActionTiming = "before"
	ActionTimingAfter  ActionTiming = "after"
	ActionTimingAsync  ActionTiming = "async"
)

// ExecutionStatus represents the status of an execution
type ExecutionStatus string

const (
	ExecutionStatusPending    ExecutionStatus = "pending"
	ExecutionStatusInProgress ExecutionStatus = "in_progress"
	ExecutionStatusCompleted  ExecutionStatus = "completed"
	ExecutionStatusFailed     ExecutionStatus = "failed"
	ExecutionStatusCancelled  ExecutionStatus = "cancelled"
	ExecutionStatusRolledBack ExecutionStatus = "rolled_back"
)

// AuditAction represents audit log actions
type AuditAction string

const (
	AuditActionCreate    AuditAction = "create"
	AuditActionUpdate    AuditAction = "update"
	AuditActionDelete    AuditAction = "delete"
	AuditActionActivate  AuditAction = "activate"
	AuditActionDeprecate AuditAction = "deprecate"
	AuditActionArchive   AuditAction = "archive"
	AuditActionClone     AuditAction = "clone"
)
