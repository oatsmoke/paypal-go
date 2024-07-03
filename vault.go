package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/paypal-go/model"
	"io"
	"net/http"
	"strings"
)

func (c *Client) CreatePaymentToken(ctx context.Context, in *model.PaymentSource) (*model.PayPal, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v3/vault/payment-tokens", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreatePaymentToken: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) RetrievePaymentToken(ctx context.Context, id string) (*model.PayPal, error) {
	tokenURL := fmt.Sprintf("%s/v3/vault/payment-tokens/%s", c.URL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", tokenURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RetrievePaymentToken: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeletePaymentToken(ctx context.Context, in string) error {
	tokenURL := fmt.Sprintf("%s/v3/vault/payment-tokens/%s", c.URL, in)
	req, err := http.NewRequestWithContext(ctx, "DELETE", tokenURL, nil)
	if err != nil {
		return err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("DeletePaymentToken: status code %d", res.StatusCode)
	}

	return nil
}

func (c *Client) CreateSetupToken(ctx context.Context, in *model.PayPal) (*model.PayPal, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v3/vault/setup-tokens", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreateSetupToken: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	return out, nil
}
