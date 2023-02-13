package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)


var connection *pgdriver.Connector

type BaseRepository struct {
	
}

func (base BaseRepository) getConnection(cfg *config.Config) *pgdriver.Connector {
	once := sync.Once{}
	once.Do(func() {
		connection = pgdriver.NewConnector(
			pgdriver.WithNetwork(cfg.FlightDB.Network),
			pgdriver.WithAddr(fmt.Sprintf("%s:%s", cfg.FlightDB.Host, strconv.Itoa(cfg.FlightDB.Port))),
			pgdriver.WithInsecure(true),
			//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
			pgdriver.WithUser(cfg.FlightDB.User),
			pgdriver.WithPassword(os.Getenv(cfg.FlightDB.PasswordKey)),
			pgdriver.WithDatabase(cfg.FlightDB.Database),
			pgdriver.WithApplicationName(cfg.AppName),
			pgdriver.WithTimeout(time.Duration(cfg.FlightDB.Timeout) * time.Second),
			pgdriver.WithDialTimeout(time.Duration(cfg.FlightDB.Timeout) * time.Second),
			pgdriver.WithReadTimeout(time.Duration(cfg.FlightDB.Timeout) * time.Second),
			pgdriver.WithWriteTimeout(time.Duration(cfg.FlightDB.Timeout) * time.Second),
		)
	})
	return connection
}

func (base BaseRepository) GetDB(cfg *config.Config) *bun.DB {
	sqldb := sql.OpenDB(base.getConnection(cfg))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
} 
