package main

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./ev.db")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(50)
	if err := db.PingContext(context.Background()); err != nil {
		panic("ping failed")
	}
}
