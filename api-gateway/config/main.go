package config

import (
	"net/url"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr     string `env:"HTTP_SERVER_ADDRESS" envDefault:"0.0.0.0:1138"`
	LogLevel string `env:"API_LOG_LEVEL" envDefault:"info"`

	UserServiceName string `env:"USER_SERVICE_NAME" envDefault:"user-service"`
	AuthServiceName string `env:"AUTH_SERVICE_NAME" envDefault:"auth-service"`
	RoomServiceName string `env:"ROOM_SERVICE_NAME" envDefault:"room-service"`

	ApiGatewaySubscribeName string `env:"API_GATEWAY_SUBSCRIBE_CHANNEL_NAME" envDefault:"ApiGatewayResp"`
}

func (c *Config) ServerAddress() *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   c.Addr,
	}
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfig() Config {
	_ = godotenv.Load()

	once.Do(func() {
		c := Config{}
		if err := env.Parse(&c); err != nil {
			panic(err)
		}

		cfg = &c
	})

	return *cfg
}
