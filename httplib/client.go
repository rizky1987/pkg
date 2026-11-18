package httplib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

// NewClient membuat instance HTTP Client dengan default timeout
func NewClient(timeout time.Duration) *Client {
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// PostJSON mengirimkan HTTP POST dengan payload JSON dan men-decode response ke target Struct
func (c *Client) PostJSON(ctx context.Context, url string, headers map[string]string, payload interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("httplib: failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return fmt.Errorf("httplib: failed to create request: %w", err)
	}

	// Default Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Custom Headers (misal: Authorization Bearer untuk WhatsApp)
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("httplib: execution failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("httplib: failed to read response body: %w", err)
	}

	// Validasi Status Code Non-2xx
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("httplib: server returned error status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode ke target struct jika disediakan
	if target != nil && len(bodyBytes) > 0 {
		if err := json.Unmarshal(bodyBytes, target); err != nil {
			return fmt.Errorf("httplib: failed to unmarshal response: %w", err)
		}
	}

	return nil
}
