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
	MicroServiceAddress string `env:"MICRO_SERVICE_ADDRESS" envDefault:"localhost:16567"`
	MongoUrl            string `env:"MONGO_URL"`
	DBName              string `env:"MONGO_DB_NAME"`
	Collection          string `env:"MONGO_COLLECTION_NAME"`
	ServiceName         string `env:"SERVICE_NAME" envDefault:"messages-service"`
	RabbitMQUrl         string `env:"RABBIT_MQ_URL"`

	//& StartConsumer
	Consumer          string `env:"CONSUMER"`
	AutoAsk           bool   `env:"AUTO_ASK" envDefault:"true"`
	ConsumerExclusive bool   `env:"CONSUMER_EXCLUSIVE" envDefault:"false"`
	NoLocal           bool   `env:"NO_LOCAL" envDefault:"false"`
	ConsumerNoWait    bool   `env:"CONSUMER_NO_WAIT" envDefault:"false"`

	//& publish
	Exchange  string `env:"EXCHANGE" envDefault:"messages-gateway"`
	Mandatory bool   `env:"MANDATORY" envDefault:"false"`
	Immediate bool   `env:"IMMEDIATE" envDefault:"false"`

	//& Queue
	Name             string `env:"NAME" envDefault:"messages"`
	Durable          bool   `env:"DURABLE" envDefault:"false"`
	DeleteWhenUnused bool   `env:"DELETE_WHEN_UNUSED" envDefault:"false"`
	QueueExclusive   bool   `env:"QUEUE_EXCLUSIVE" envDefault:"true"`
	QueueNoWait      bool   `env:"QUEUE_NO_WAIT" envDefault:"false"`

	//& QoS
	PrefetchCount int  `env:"PREFETCH_COUNT" envDefault:"1"`
	PrefetchSize  int  `env:"PREFETCH_SIZE" envDefault:"0"`
	Global        bool `env:"GLOBAL" envDefault:"false"`

	//& ExchangeDeclare
	Type            string `env:"TYPE" envDefault:"fanout"`
	ExchangeDurable bool   `env:"EXCHANGE_DURABLE" envDefault:"true"`
	AutoDelete      bool   `env:"AUTO_DELETE" envDefault:"false"`
	Internal        bool   `env:"INTERNAL" envDefault:"false"`
	ExchangeNoWait  bool   `env:"EXCHANGE_NO_WAIT" envDefault:"false"`
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
