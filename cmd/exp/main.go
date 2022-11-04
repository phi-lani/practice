package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DBconfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg DBconfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
func main() {
	cfg := DBconfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "picshare",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}
