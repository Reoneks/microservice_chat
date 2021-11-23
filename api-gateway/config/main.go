package config

import (
	"net/url"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr        string `env:"HTTP_SERVER_ADDRESS" envDefault:"0.0.0.0:1138"`
	LogLevel    string `env:"API_LOG_LEVEL" envDefault:"info"`
	RabbitMQUrl string `env:"RABBIT_MQ_URL" envDefault:"amqp://guest:guest@localhost:5672/"`

	UserServiceName    string `env:"USER_SERVICE_NAME" envDefault:"user-service"`
	AuthServiceName    string `env:"AUTH_SERVICE_NAME" envDefault:"auth-service"`
	RoomServiceName    string `env:"ROOM_SERVICE_NAME" envDefault:"room-service"`
	MessageServiceName string `env:"MESSAGE_SERVICE_NAME" envDefault:"messages-service"`

	ApiGatewaySubscribeName string `env:"API_GATEWAY_SUBSCRIBE_CHANNEL_NAME" envDefault:"ApiGatewayResp"`

	//& publish
	Mandatory bool `env:"MANDATORY" envDefault:"false"`
	Immediate bool `env:"IMMEDIATE" envDefault:"false"`

	//& Send Queue
	SendName             string `env:"NAME" envDefault:"messages"`
	SendDurable          bool   `env:"DURABLE" envDefault:"false"`
	SendDeleteWhenUnused bool   `env:"DELETE_WHEN_UNUSED" envDefault:"false"`
	SendQueueExclusive   bool   `env:"QUEUE_EXCLUSIVE" envDefault:"true"`
	SendQueueNoWait      bool   `env:"QUEUE_NO_WAIT" envDefault:"false"`

	//& Receive Queue
	ReceiveName             string `env:"NAME" envDefault:"messages"`
	ReceiveDurable          bool   `env:"DURABLE" envDefault:"false"`
	ReceiveDeleteWhenUnused bool   `env:"DELETE_WHEN_UNUSED" envDefault:"false"`
	ReceiveQueueExclusive   bool   `env:"QUEUE_EXCLUSIVE" envDefault:"true"`
	ReceiveQueueNoWait      bool   `env:"QUEUE_NO_WAIT" envDefault:"false"`

	//& ExchangeDeclare
	Exchange        string `env:"EXCHANGE" envDefault:"messages-gateway"`
	Type            string `env:"TYPE" envDefault:"fanout"`
	ExchangeDurable bool   `env:"EXCHANGE_DURABLE" envDefault:"true"`
	AutoDelete      bool   `env:"AUTO_DELETE" envDefault:"false"`
	Internal        bool   `env:"INTERNAL" envDefault:"false"`
	ExchangeNoWait  bool   `env:"EXCHANGE_NO_WAIT" envDefault:"false"`
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
