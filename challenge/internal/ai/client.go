package ai

import "meshnode/internal/config"

type Client struct {
	Endpoint string
	Model    string
	Key      string
	Timeout  int
	Batch    int
	Window   int
}

func FromConfig(cfg config.AI) Client {
	key := cfg.Key
	if key == "" {
		key = defaultKey()
	}
	return Client{
		Endpoint: cfg.Endpoint,
		Model:    cfg.Model,
		Key:      key,
		Timeout:  cfg.Timeout,
		Batch:    cfg.Batch,
		Window:   cfg.Window,
	}
}

func (c Client) Ready() bool {
	return c.Endpoint != "" && c.Model != "" && c.Key != ""
}

func defaultKey() string {
	return ""
}
