package config

import (
	"database/sql"
	"sync"

	"github.com/caarlos0/env"
	"github.com/golang-migrate/migrate/v4"
	mpostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName  string `env:"SERVICE_NAME" envDefault:"auth-service"`
	LogLevel     string `env:"API_LOG_LEVEL" envDefault:"info"`
	DSN          string `env:"DB_DSN" envDefault:"host=0.0.0.0 user=postgres password=postgres dbname=go-micro-template-user-service port=5432 sslmode=disable"`
	MigrationURL string `env:"DB_MIGRATION_URL" envDefault:"file://auth/migrations"`

	Secret    string `env:"JWT_SECRET" envDefault:"9caf06bb4436cdbfa20af9121a626bc1093c4f54b31c0fa937957856135345b6"`
	Algorithm string `env:"JWT_ALGORITHM" envDefault:"HS256"`

	UserServiceName string `env:"USER_SERVICE_NAME" envDefault:"auth-service"`
}

func NewDB(dsn, migrationsURL string) (db.Session, error) {
	if err := migrations(dsn, migrationsURL); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	settings, err := postgresql.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	session, err := postgresql.Open(settings)
	if err != nil {
		return nil, err
	}

	return session, nil
}

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
