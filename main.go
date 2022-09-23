package main

import (
	"net/http"
	xdb "ximlr/go-lang-web/database/init"
	xhttp "ximlr/go-lang-web/http"
)

func main() {

	xdb.DbInit()
	router := xhttp.HttpInit()

	err := http.ListenAndServe(":8888", router)
	if err != nil {
		panic(err.Error())
	}
}
