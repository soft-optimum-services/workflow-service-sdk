package workflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) get(ctx context.Context, rawURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", rawURL, resp.Status)
	}

	return body, nil
}

func (c *Client) postJSON(ctx context.Context, rawURL string, payload any, expectedStatuses ...int) ([]byte, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if !matchesStatus(resp.StatusCode, expectedStatuses...) {
		return nil, fmt.Errorf("POST %s: %s", rawURL, resp.Status)
	}

	return respBody, nil
}

func matchesStatus(status int, expectedStatuses ...int) bool {
	for _, expected := range expectedStatuses {
		if status == expected {
			return true
		}
	}
	return false
}

func parseBaseURL(baseURL string) (*url.URL, error) {
	base := strings.TrimSpace(baseURL)
	if base == "" {
		return nil, fmt.Errorf("base URL is required")
	}
	root, err := url.Parse(base)
	if err != nil || root.Scheme == "" || root.Host == "" {
		return nil, fmt.Errorf("invalid base URL")
	}
	return root, nil
}

func (c *Client) rootURL() (*url.URL, error) {
	return parseBaseURL(c.baseURL)
}
