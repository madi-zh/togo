package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBSession struct {
	*sql.DB
}

func (s *DBSession) Close() {
	s.DB.Close()
}

func CreateSession(cfg Config) *DBSession {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	if conn, err := sql.Open("postgres", psqlInfo); err != nil {
		log.Fatal("Issue while opening conn", err)
	} else {
		return &DBSession{conn}
	}
	return nil
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
