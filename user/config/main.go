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
	ServiceName  string `env:"SERVICE_NAME" envDefault:"user-service"`
	MigrationURL string `env:"DB_MIGRATION_URL" envDefault:"file://user/migrations"`
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
