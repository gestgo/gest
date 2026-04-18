package indicators

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gestgo/gest/package/technique/health"
)

type HTTPIndicator struct {
	key    string
	url    string
	client *http.Client
}

func NewHTTPIndicator(key, url string, client *http.Client) *HTTPIndicator {
	if client == nil {
		client = http.DefaultClient
	}
	return &HTTPIndicator{key: key, url: url, client: client}
}

func (i *HTTPIndicator) Check(ctx context.Context) health.IndicatorResult {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, i.url, nil)
	if err != nil {
		return health.Down(i.key, err, nil)
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return health.Down(i.key, err, nil)
	}
	resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return health.Down(i.key, fmt.Errorf("unexpected status %d", resp.StatusCode), map[string]any{
			"statusCode": resp.StatusCode,
			"url":        i.url,
		})
	}

	return health.Up(i.key, map[string]any{"statusCode": resp.StatusCode, "url": i.url})
}
