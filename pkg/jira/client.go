package jira

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matthewrobinsondev/lazyjira/pkg/config"
)

type ClientInterface interface {
	NewRequest(method, endpoint string, params url.Values, body io.Reader) ([]byte, error)
}

type Client struct {
	httpClient      *http.Client
	BaseURL         string
	BasicAuthHeader string
}

func NewClient(cfg *config.Config) *Client {
	httpClient := &http.Client{}
	return &Client{
		httpClient:      httpClient,
		BaseURL:         cfg.JiraURL,
		BasicAuthHeader: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cfg.Email, cfg.APIKey))),
	}
}

func (c *Client) NewRequest(method, endpoint string, params url.Values, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	if params != nil {
		url = url + fmt.Sprintf("?%s", params.Encode())
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.BasicAuthHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
