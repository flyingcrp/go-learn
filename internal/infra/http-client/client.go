package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var Default = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	},
}

func Get[T any](ctx context.Context, url string) (*T, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set("Accept", "application/json")
	return do[T](Default.Do(req))
}

func Post[T any](ctx context.Context, url string, body any) (*T, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return do[T](Default.Do(req))
}

func do[T any](resp *http.Response, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("http %d from %s: %s", resp.StatusCode, resp.Request.URL, strings.TrimSpace(string(body)))
	}

	ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return nil, fmt.Errorf("unexpected content-type: %s", ct)
	}

	var target T
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}
	return &target, nil
}
