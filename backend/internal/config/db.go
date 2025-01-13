package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/volatiletech/sqlboiler/boil"
)

type DBConfig interface {
	GetDriver() string
	GetUsername() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetName() string
}

type DB struct {
	Driver   string `envconfig:"DB_DRIVER" required:"true" default:"mysql"`
	Username string `envconfig:"DB_USERNAME" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true" redact:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true" default:"3306"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func NewDB(db *DB) (*bun.DB, error) {
	return newDB(db)
}

func (db DB) GetDriver() string {
	return db.Driver
}

func (db DB) GetUsername() string {
	return db.Username
}

func (db DB) GetPassword() string {
	return db.Password
}

func (db DB) GetHost() string {
	return db.Host
}

func (db DB) GetPort() string {
	return db.Port
}

func (db DB) GetName() string {
	return db.Name
}

func newDB(cfg DBConfig) (*bun.DB, error) {
	sqlDB, err := sql.Open(
		cfg.GetDriver(),
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.GetUsername(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetName(),
		),
	)
	if err != nil {
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(300 * time.Second)
	boil.SetDB(sqlDB)

	return bun.NewDB(sqlDB, mysqldialect.New()), nil
}
