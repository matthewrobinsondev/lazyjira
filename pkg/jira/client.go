package jira

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matthewrobinsondev/lazyjira/pkg/config"
)

const VERSION_2 = "/rest/api/2"
const VERSION_3 = "/rest/api/3"

type ClientInterface interface {
	NewRequest(method, endpoint string, params url.Values, body io.Reader) ([]byte, error)
}

type Client struct {
	httpClient      *http.Client
	BaseURL         string
	BasicAuthHeader string
	version         string
}

func NewClient(cfg *config.Config) *Client {
	httpClient := &http.Client{}
	return &Client{
		httpClient:      httpClient,
		BaseURL:         cfg.JiraURL,
		BasicAuthHeader: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cfg.Email, cfg.AccessToken))),
		version:         VERSION_3,
	}
}

func (c *Client) UseV2() *Client {
	c.version = VERSION_2
	return c
}

func (c *Client) UseV3() *Client {
	c.version = VERSION_3
	return c
}

func (c *Client) NewRequest(method, endpoint string, params url.Values, body io.Reader) ([]byte, error) {
	baseEndpoint := c.BaseURL + "/" + c.version
	reqsUri := fmt.Sprintf("%s%s", baseEndpoint, endpoint)

	if params != nil {
		reqsUri = reqsUri + fmt.Sprintf("?%s", params.Encode())
	}

	req, err := http.NewRequest(method, reqsUri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.BasicAuthHeader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

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
