package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseContext struct {
	Db               *sql.DB
	ConnectionString string
}

func CreateContext() *DatabaseContext {
	connectionString := ConnectionString()

	//
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return &DatabaseContext{
		Db:               db,
		ConnectionString: connectionString,
	}
}
