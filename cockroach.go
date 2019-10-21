package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func connect() *sql.DB{
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database :", err)
	}
	return db
}
