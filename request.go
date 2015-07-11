package parsesearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WebhookRequest is a Parse Cloud Code Webhook request.
type WebhookRequest struct {
	InstallationID string                 `json:"installationId,omitempty"`
	Master         bool                   `json:"master,omitempty"`
	Object         interface{}            `json:"object,omitempty"`
	Params         map[string]interface{} `json:"params,omitempty"`
	TriggerName    string                 `json:"triggerName,omitempty"`
}

func webhookRequest(r *http.Request, webhookKey string) (*WebhookRequest, error) {
	req := &WebhookRequest{}
	buf := &bytes.Buffer{}
	io.Copy(buf, r.Body)
	defer r.Body.Close()
	if err := json.NewDecoder(buf).Decode(&req); err != nil {
		return nil, err
	}
	if r.Header.Get("X-Parse-Webhook-Key") != webhookKey {
		return nil, fmt.Errorf("invalid webhook key")
	}
	return req, nil
}
