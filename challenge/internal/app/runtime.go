package app

import (
	"meshnode/internal/ai"
	"meshnode/internal/config"
	"meshnode/internal/router"
	"meshnode/internal/store/pg"
)

type Runtime struct {
	Config config.Config
	Store  pg.Conn
	AI     ai.Client
	Router router.Router
}

func Load() Runtime {
	cfg := config.Load()
	return Runtime{
		Config: cfg,
		Store:  pg.FromConfig(cfg.Store),
		AI:     ai.FromConfig(cfg.AI),
		Router: router.New(cfg.HTTP),
	}
}

func (r Runtime) Bind() string {
	return r.Config.HTTP.Bind
}

func (r Runtime) StoreName() string {
	return r.Store.Name
}

func (r Runtime) RouteCount() int {
	return len(r.Router.Routes)
}
