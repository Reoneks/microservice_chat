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
	MicroServiceAddress string `env:"MICRO_SERVICE_ADDRESS" envDefault:"localhost:16566"`
	MongoUrl            string `env:"MONGO_URL"`
	DBName              string `env:"MONGO_DB_NAME"`
	RoomCollection      string `env:"MONGO_ROOM_COLLECTION_NAME"`
	RoomUserCollection  string `env:"MONGO_ROOM_USER_COLLECTION_NAME"`
	ServiceName         string `env:"SERVICE_NAME" envDefault:"room-service"`
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
