package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/paypal-go/model"
	"net/http"
	"strings"
)

func (c *Client) CreateOrder(ctx context.Context, in *model.PayPal) (*model.PayPal, error) {
	const fn = "CreateOrder"

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v2/checkout/orders", c.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req, http.StatusCreated, fn)
	if err != nil {
		return nil, err
	}

	out := new(model.PayPal)
	if err := json.Unmarshal(res, out); err != nil {
		return nil, err
	}

	return out, nil
}
