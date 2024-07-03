package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/paypal-go/model"
	"net/http"
	"strings"
)

func (c *Client) CreatePaymentToken(ctx context.Context, in *model.PaymentSource) (*model.PayPal, error) {
	const fn = "CreatePaymentToken"

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v3/vault/payment-tokens", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req, http.StatusCreated, fn)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) RetrievePaymentToken(ctx context.Context, id string) (*model.PayPal, error) {
	const fn = "RetrievePaymentToken"

	tokenURL := fmt.Sprintf("%s/v3/vault/payment-tokens/%s", c.URL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", tokenURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req, http.StatusOK, fn)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeletePaymentToken(ctx context.Context, in string) error {
	const fn = "DeletePaymentToken"

	tokenURL := fmt.Sprintf("%s/v3/vault/payment-tokens/%s", c.URL, in)
	req, err := http.NewRequestWithContext(ctx, "DELETE", tokenURL, nil)
	if err != nil {
		return err
	}

	if _, err := c.SendWithAuth(ctx, req, http.StatusNoContent, fn); err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateSetupToken(ctx context.Context, in *model.PayPal) (*model.PayPal, error) {
	const fn = "CreateSetupToken"

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v3/vault/setup-tokens", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req, http.StatusCreated, fn)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, err
	}

	return out, nil
}
