package config

import (
	"os"
	"strconv"
)

type HTTP struct {
	Bind         string
	ReadTimeout  int
	WriteTimeout int
	MaxBody      int
	Prefix       string
}

type Store struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Pool     int
	Idle     int
	TLSMode  string
}

type AI struct {
	Endpoint string
	Model    string
	Key      string
	Timeout  int
	Batch    int
	Window   int
}

type UI struct {
	Root  string
	Title string
	Mode  string
	Poll  int
}

type Deploy struct {
	Host   string
	User   string
	Port   string
	Target string
}

type Config struct {
	Service string
	HTTP    HTTP
	Store   Store
	AI      AI
	UI      UI
	Deploy  Deploy
	Buffer  int
}

func Load() Config {
	return Config{
		Service: envString("APP_SERVICE", "atlas-gateway"),
		HTTP: HTTP{
			Bind:         envString("APP_BIND", "127.0.0.1:8090"),
			ReadTimeout:  envInt("APP_READ_TIMEOUT", 1400),
			WriteTimeout: envInt("APP_WRITE_TIMEOUT", 2200),
			MaxBody:      envInt("APP_MAX_BODY", 2097152),
			Prefix:       envString("APP_PREFIX", "/node"),
		},
		Store: Store{
			Driver:   "postgres",
			Host:     envString("DB_HOST", "10.12.7.33"),
			Port:     envInt("DB_PORT", 5432),
			Name:     envString("DB_NAME", "atlas"),
			User:     envString("DB_USER", ""),
			Password: envString("DB_PASSWORD", ""),
			Pool:     envInt("DB_POOL", 10),
			Idle:     envInt("DB_IDLE", 4),
			TLSMode:  envString("DB_TLS_MODE", "prefer"),
		},
		AI: AI{
			Endpoint: envString("LLM_ENDPOINT", "https://llm.node.local/v1"),
			Model:    envString("LLM_MODEL", "g-archive"),
			Key:      envString("LLM_KEY", ""),
			Timeout:  envInt("LLM_TIMEOUT", 2100),
			Batch:    envInt("LLM_BATCH", 16),
			Window:   envInt("LLM_WINDOW", 160),
		},
		UI: UI{
			Root:  envString("UI_ROOT", "#app"),
			Title: envString("UI_TITLE", "atlas-node"),
			Mode:  envString("UI_MODE", "matrix"),
			Poll:  envInt("UI_POLL", 700),
		},
		Deploy: Deploy{
			Host:   envString("DEPLOY_HOST", "10.21.4.31"),
			User:   envString("DEPLOY_USER", "release"),
			Port:   envString("DEPLOY_PORT", "2204"),
			Target: envString("DEPLOY_TARGET", "/srv/mesh-flow"),
		},
		Buffer: envInt("APP_BUFFER", 640),
	}
}

func envString(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}
