// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defaultAPIURL = "https://api.gridscale.io"

// GridscaleClient is a thin REST client for the gridscale API.
type GridscaleClient struct {
	baseURL    string
	userUUID   string
	token      string
	httpClient *http.Client
}

// NewGridscaleClient creates a new GridscaleClient. If baseURL is empty, the
// default gridscale API URL is used.
func NewGridscaleClient(userUUID, token, baseURL string) *GridscaleClient {
	if baseURL == "" {
		baseURL = defaultAPIURL
	}
	return &GridscaleClient{
		baseURL:    baseURL,
		userUUID:   userUUID,
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// Get performs an authenticated GET request to the given path and decodes the
// JSON response body into out.
func (c *GridscaleClient) Get(ctx context.Context, path string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("X-Auth-UserID", c.userUUID)
	req.Header.Set("X-Auth-Token", c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d for %s", resp.StatusCode, path)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
