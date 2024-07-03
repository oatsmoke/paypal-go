package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/paypal-go/model"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	ID       string    `json:"id"`
	Secret   string    `json:"secret"`
	URL      string    `json:"url"`
	Token    string    `json:"token"`
	TokenExp time.Time `json:"token_exp"`
	HTTP     *http.Client
}

func NewClient(id, secret, url string) *Client {
	return &Client{
		ID:     id,
		Secret: secret,
		URL:    url,
		HTTP:   &http.Client{},
	}
}

func (c *Client) GetAccessToken(ctx context.Context) error {
	url := fmt.Sprintf("%s/v1/oauth2/token", c.URL)
	body := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ID, c.Secret)

	res, err := c.SendRequest(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetAccessToken: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := new(model.Authentication)
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	c.Token = fmt.Sprintf("%s %s", result.TokenType, result.AccessToken)
	c.TokenExp = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second)

	return nil
}

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v", err)
		}
	}(res.Body)

	return res, nil
}

func (c *Client) SendWithAuth(ctx context.Context, req *http.Request) (*http.Response, error) {
	if c.Token == "" || time.Now().After(c.TokenExp) {
		if err := c.GetAccessToken(ctx); err != nil {
			return nil, err
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PayPal-Request-Id", time.Now().String())
	req.Header.Set("Authorization", c.Token)

	return c.SendRequest(req)
}
