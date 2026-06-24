package router

import "meshnode/internal/config"

type Route struct {
	Method string
	Path   string
}

type Router struct {
	Bind         string
	Prefix       string
	ReadTimeout  int
	WriteTimeout int
	MaxBody      int
	Routes       []Route
}

func New(cfg config.HTTP) Router {
	return Router{
		Bind:         cfg.Bind,
		Prefix:       cfg.Prefix,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		MaxBody:      cfg.MaxBody,
		Routes: []Route{
			{Method: "GET", Path: cfg.Prefix + "/health"},
			{Method: "GET", Path: cfg.Prefix + "/nodes"},
			{Method: "POST", Path: cfg.Prefix + "/sync"},
		},
	}
}
