package config

import (
	"sync"

	"github.com/caarlos0/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN         string `env:"DB_DSN" envDefault:"host=0.0.0.0 user=postgres password=postgres dbname=messages-service port=5432 sslmode=disable"`
	ServiceName string `env:"SERVICE_NAME" envDefault:"messages-service"`
	RabbitMQUrl string `env:"RABBIT_MQ_URL" envDefault:"amqp://guest:guest@localhost:5672/"`
}

func NewDB(dsn string) (*gorm.DB, error) {
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return client, nil
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
