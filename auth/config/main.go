package config

import (
	"context"
	"sync"

	"github.com/caarlos0/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName         string `env:"SERVICE_NAME" envDefault:"auth-service"`
	LogLevel            string `env:"API_LOG_LEVEL" envDefault:"info"`
	MicroServiceAddress string `env:"MICRO_SERVICE_ADDRESS" envDefault:"localhost:16565"`
	MongoUrl            string `env:"MONGO_URL"`
	DBName              string `env:"MONGO_DB_NAME"`
	Collection          string `env:"MONGO_COLLECTION_NAME"`

	Secret    string `env:"JWT_SECRET" envDefault:"9caf06bb4436cdbfa20af9121a626bc1093c4f54b31c0fa937957856135345b6"`
	Algorithm string `env:"JWT_ALGORITHM" envDefault:"HS256"`

	UserServiceName string `env:"USER_SERVICE_NAME" envDefault:"user-service"`
	UserServiceADDR string `env:"USER_SERVICE_ADDR" envDefault:"localhost:16564"`
}

func NewDB(mongoURL string) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
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
