package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	DB *sqlx.DB
}

//Initiate the db connection, and return a Server instance with the db connection
func CreateDbConn(conn string) Server {
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return Server{
		DB: db,
	}
}
