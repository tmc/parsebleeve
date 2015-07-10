package parsesearch

type WebhookRequest struct {
	InstallationID string                 `json:"installationId,omitempty"`
	Master         bool                   `json:"master,omitempty"`
	Object         interface{}            `json:"object,omitempty"`
	Params         map[string]interface{} `json:"params,omitempty"`
	TriggerName    string                 `json:"triggerName,omitempty"`
}
