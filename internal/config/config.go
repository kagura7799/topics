package config

import (
	"flag"
	"os"

	"github.com/pkg/errors"
)

// First option is the default value
type Config struct {
	Env  string // local | dev | prod
	Port int    // 80
	Dsn  string
}

func New() (*Config, error) {
	const op = "config.New"
	port := flag.Int("port", 80, "port the server will run on")
	flag.Parse()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("dsn is not set. You can pass DB_DSN as env variable or provide it in '.env' file.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	return &Config{
		Port: *port,
		Dsn:  dsn,
		Env:  env,
	}, nil
}
