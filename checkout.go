package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paypal-go/model"
	"strings"
)

func (c *Client) CreateOrder(ctx context.Context, in *model.PayPal) (*model.PayPal, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v2/checkout/orders", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreateOrder: status code %d", res.StatusCode)
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
