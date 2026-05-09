package workflow

import (
	"net/http"
	"os"
	"strings"
)

const envWorkflowServiceBaseURL = "WORKFLOW_SERVICE_BASE_URL"

// Config configures a workflow API client bound to a single origin (scheme + host [+ port]).
type Config struct {
	// BaseURL is the workflow-service root URL, without trailing slash (e.g. https://workflow.example.com).
	BaseURL string
	// HTTPClient is optional; when nil, New uses a client with a 12s timeout.
	HTTPClient *http.Client
}

// BaseURLFromEnv reads WORKFLOW_SERVICE_BASE_URL (trimmed).
func BaseURLFromEnv() string {
	return strings.TrimSpace(os.Getenv(envWorkflowServiceBaseURL))
}
