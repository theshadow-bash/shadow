package pg

import (
	"fmt"

	"meshnode/internal/config"
)

type Conn struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Pool     int
	Idle     int
	TLSMode  string
}

func FromConfig(cfg config.Store) Conn {
	return Conn{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Name:     cfg.Name,
		User:     cfg.User,
		Password: cfg.Password,
		Pool:     cfg.Pool,
		Idle:     cfg.Idle,
		TLSMode:  cfg.TLSMode,
	}
}

func (c Conn) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s pool_max_conns=%d pool_min_conns=%d",
		c.Host,
		c.Port,
		c.Name,
		c.User,
		c.Password,
		c.TLSMode,
		c.Pool,
		c.Idle,
	)
}

func (c Conn) Ready() bool {
	return c.Host != "" && c.Name != "" && c.User != "" && c.Password != ""
}
