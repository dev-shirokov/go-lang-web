package database

import "fmt"

const (
	HOST     = "127.0.0.1"
	DATABASE = "notes"
	USER     = "postgres"
	PASSWORD = "example"
)

func ConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	return connectionString
}
