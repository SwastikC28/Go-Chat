package httpserver

import "github.com/ilyakaznacheev/cleanenv"

type httpserverConfig struct {
	ApiPort int `env:"API_PORT" env-default:"8080"`
}

func newHTTPServerConfig() *httpserverConfig {
	config := httpserverConfig{}
	cleanenv.ReadEnv(&config)
	return &config
}
