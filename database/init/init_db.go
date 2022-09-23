package db_init

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"ximlr/go-lang-web/database"
)

func DbInit() {

	dbContext := database.CreateContext()

	//
	fmt.Println("Database connecting...")
	fmt.Println(dbContext.ConnectionString)

	//
	err := dbContext.Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection successfully")

	//
	sqlFile := filepath.Join("database", "init", "init_db.sql")

	//
	fileBytes, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		panic("Reading of 'init_db.sql' file  failed")
	}
	sqlScript := string(fileBytes)

	//
	_, err = dbContext.Db.Exec(sqlScript)
	if err != nil {
		panic("'init_db.sql' executing failed")
	}
	fmt.Println("'init_db.sql' executed")
}
