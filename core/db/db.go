package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBSession struct {
	*sql.DB
}

func (s *DBSession) Close() {
	s.DB.Close()
}

func CreateSession() *DBSession {
	psqlInfo := "host=localhost port=5430 user=test password=test dbname=database sslmode=disable"
	if conn, err := sql.Open("postgres", psqlInfo); err != nil {
		log.Fatal("Issue while opening conn", err)
	} else {
		return &DBSession{conn}
	}
	return nil
}
