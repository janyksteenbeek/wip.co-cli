package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const ENDPOINT = "https://wip.co/graphql"

type Client struct {
	HTTPClient *http.Client
	APIKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		APIKey:     apiKey,
	}
}

func (c *Client) Request(payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", ENDPOINT, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "github.com/janyksteenbeek/wipco-cli")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}
