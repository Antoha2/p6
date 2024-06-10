package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"4/internal/config"
	"4/internal/helper"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	run()
}

func run() {
	cfg := config.MustLoad()
	dbx, err := initDb(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = SelectData(dbx)
	if err != nil {
		log.Println(err)
	}
}

func SelectData(dbx *gorm.DB) error {

	s := &helper.Test{}
	query := "select id, name from test where id = $1"
	result := dbx.Table("Test").Raw(query, 1).Scan(s)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	log.Println(s)

	return nil
}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%d",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
		5,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open(): %v", err)
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("error to ping connection pool: %v", err)
	}
	log.Printf("Подключение к базе данных на http://127.0.0.1:%s\n", cfg.DB.Port)
	return gormDB, nil
}
