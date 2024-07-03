package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paypal-go/model"
)

func (c *Client) RefundCapturedPayment(ctx context.Context, captureId string) (*model.Capture, error) {
	url := fmt.Sprintf("%s/v2/payments/captures/%s/refund", c.URL, captureId)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("RefundCapturedPayment: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	out := new(model.Capture)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	return out, nil
}