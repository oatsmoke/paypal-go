package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/paypal-go/model"
	"net/http"
)

func (c *Client) RefundCapturedPayment(ctx context.Context, captureId string) (*model.Capture, error) {
	const fn = "RefundCapturedPayment"

	url := fmt.Sprintf("%s/v2/payments/captures/%s/refund", c.URL, captureId)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.SendWithAuth(ctx, req, http.StatusCreated, fn)
	if err != nil {
		return nil, err
	}

	out := new(model.Capture)
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, err
	}

	return out, nil
}
