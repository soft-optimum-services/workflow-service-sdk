package workflow

import (
	"net/http"
	"strings"
	"time"
)

// Client is the primary facade implementing TemplateReader, ExecutionClient, and WorkflowQueries.
type Client struct {
	baseURL string
	http    *http.Client
}

// New creates a Client from Config. BaseURL must be non-empty before calling API methods
// (helpers return empty/nil or errors when unset).
func New(cfg Config) *Client {
	hc := cfg.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: 12 * time.Second}
	}
	return &Client{
		baseURL: strings.TrimSpace(cfg.BaseURL),
		http:    hc,
	}
}

// NewFromEnv builds a client using WORKFLOW_SERVICE_BASE_URL.
func NewFromEnv() *Client {
	return New(Config{BaseURL: BaseURLFromEnv()})
}
