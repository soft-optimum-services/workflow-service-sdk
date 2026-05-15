package workflow

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	wfmodel "github.com/soft-optimum-services/workflow-service-sdk/model"
)

// ListManualTransitionsFromState returns enabled manual transitions leaving fromStateCode for templateCode,
// sorted by priority then code. Returns nil when configuration or upstream data is missing.
func (c *Client) ListManualTransitionsFromState(ctx context.Context, templateCode, fromStateCode string) []wfmodel.Transition {
	if c == nil {
		c = NewFromEnv()
	}

	templateCode = strings.TrimSpace(templateCode)
	fromStateCode = strings.TrimSpace(fromStateCode)
	if strings.TrimSpace(c.baseURL) == "" || templateCode == "" || fromStateCode == "" {
		return nil
	}

	states, err := c.ListStates(ctx, templateCode)
	if err != nil || len(states) == 0 {
		return nil
	}
	transitions, err := c.ListTransitions(ctx, templateCode)
	if err != nil || len(transitions) == 0 {
		return nil
	}

	var fromStateID string
	for _, st := range states {
		if st.Code == fromStateCode {
			fromStateID = st.ID.String()
			break
		}
	}
	if fromStateID == "" {
		return nil
	}

	filtered := make([]wfmodel.Transition, 0)
	for _, tr := range transitions {
		if !tr.IsEnabled || string(tr.TriggerType) != "manual" {
			continue
		}
		if tr.FromStateID.String() == fromStateID && tr.Code != "" {
			filtered = append(filtered, tr)
		}
	}
	if len(filtered) == 0 {
		return nil
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		a, b := filtered[i], filtered[j]
		if a.Priority != b.Priority {
			return a.Priority < b.Priority
		}
		return a.Code < b.Code
	})

	return filtered
}

// ResolveTransitionTargetStateCode resolves the target state's code for a manual transition.
func (c *Client) ResolveTransitionTargetStateCode(ctx context.Context, templateCode, fromStateCode, transitionCode string) (string, error) {
	if c == nil {
		c = NewFromEnv()
	}

	if strings.TrimSpace(c.baseURL) == "" {
		return "", errors.New("WORKFLOW_SERVICE_BASE_URL unset")
	}

	templateCode = strings.TrimSpace(templateCode)
	fromStateCode = strings.TrimSpace(fromStateCode)
	if templateCode == "" {
		return "", fmt.Errorf("template code is required")
	}
	if fromStateCode == "" {
		return "", fmt.Errorf("from state code is required")
	}

	states, err := c.ListStates(ctx, templateCode)
	if err != nil || len(states) == 0 {
		return "", err
	}
	transitions, err := c.ListTransitions(ctx, templateCode)
	if err != nil || len(transitions) == 0 {
		return "", err
	}

	stateByID := make(map[string]wfmodel.State, len(states))
	var fromStateID string
	for _, st := range states {
		stateID := st.ID.String()
		if stateID != "" {
			stateByID[stateID] = st
		}
		if st.Code == fromStateCode && stateID != "" {
			fromStateID = stateID
		}
	}
	if fromStateID == "" {
		return "", fmt.Errorf("workflow state %q not found", fromStateCode)
	}

	trCode := strings.TrimSpace(transitionCode)
	var chosen *wfmodel.Transition
	for i := range transitions {
		tr := &transitions[i]
		if tr.Code != trCode {
			continue
		}
		if !tr.IsEnabled || string(tr.TriggerType) != "manual" {
			return "", fmt.Errorf("transition %q is not an enabled manual transition", trCode)
		}
		if tr.FromStateID.String() != fromStateID {
			return "", fmt.Errorf("transition %q does not apply from workflow state %q", trCode, fromStateCode)
		}
		chosen = tr
		break
	}
	if chosen == nil {
		return "", fmt.Errorf("unknown transition %q", trCode)
	}

	to, ok := stateByID[chosen.ToStateID.String()]
	if !ok || to.Code == "" {
		return "", errors.New("workflow to_state not found for transition")
	}
	return to.Code, nil
}

// ResolveTransitionSourceStateCode returns the source state's code for the given manual transition.
// It is the inverse direction of ResolveTransitionTargetStateCode (transition → from-state code rather than from-state + transition → to-state code).
func (c *Client) ResolveTransitionSourceStateCode(ctx context.Context, templateCode, transitionCode string) (string, error) {
	if c == nil {
		c = NewFromEnv()
	}

	if strings.TrimSpace(c.baseURL) == "" {
		return "", errors.New("WORKFLOW_SERVICE_BASE_URL unset")
	}

	templateCode = strings.TrimSpace(templateCode)
	trCode := strings.TrimSpace(transitionCode)
	if templateCode == "" {
		return "", fmt.Errorf("template code is required")
	}
	if trCode == "" {
		return "", fmt.Errorf("transition code is required")
	}

	states, err := c.ListStates(ctx, templateCode)
	if err != nil || len(states) == 0 {
		return "", err
	}
	transitions, err := c.ListTransitions(ctx, templateCode)
	if err != nil || len(transitions) == 0 {
		return "", err
	}

	stateByID := make(map[string]wfmodel.State, len(states))
	for _, st := range states {
		id := st.ID.String()
		if id != "" {
			stateByID[id] = st
		}
	}

	var chosen *wfmodel.Transition
	for i := range transitions {
		tr := &transitions[i]
		if tr.Code != trCode {
			continue
		}
		if !tr.IsEnabled || string(tr.TriggerType) != "manual" {
			return "", fmt.Errorf("transition %q is not an enabled manual transition", trCode)
		}
		if chosen != nil {
			return "", fmt.Errorf("ambiguous transition code %q", trCode)
		}
		chosen = tr
	}
	if chosen == nil {
		return "", fmt.Errorf("unknown transition %q", trCode)
	}

	from, ok := stateByID[chosen.FromStateID.String()]
	if !ok || from.Code == "" {
		return "", errors.New("workflow from_state not found for transition")
	}
	return from.Code, nil
}
