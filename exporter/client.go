package exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var DefaultClient = &http.Client{Timeout: 1 * time.Minute}

func HTTPRequest[T any](endpoint string, path string, headers map[string]string) (T, error) {
	var output T

	req, err := http.NewRequest("GET", endpoint+path, nil)
	if err != nil {
		return output, err
	}

	req.Header.Set("User-Agent", "samber/promql_exporter")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return output, fmt.Errorf("status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return output, err
	}

	return output, nil
}
