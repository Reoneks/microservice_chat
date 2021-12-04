package config

import (
	"database/sql"
	"sync"

	"github.com/caarlos0/env"
	"github.com/golang-migrate/migrate/v4"
	mpostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN          string `env:"DB_DSN" envDefault:"host=0.0.0.0 user=postgres password=postgres dbname=analytics port=5433 sslmode=disable"`
	ServiceName  string `env:"SERVICE_NAME" envDefault:"messages-service"`
	RabbitMQUrl  string `env:"RABBIT_MQ_URL" envDefault:"amqp://guest:guest@localhost:5672/"`
	MigrationURL string `env:"DB_MIGRATION_URL" envDefault:"file://auth/migrations"`

	//& StartConsumer
	Consumer          string `env:"CONSUMER" envDefault:""`
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

func NewDB(dsn, migrationsURL string) (*gorm.DB, error) {
	if err := migrations(dsn, migrationsURL); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

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

func migrations(dsn, migrationsURL string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	driver, err := mpostgres.WithInstance(db, &mpostgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	return m.Up()
}

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
