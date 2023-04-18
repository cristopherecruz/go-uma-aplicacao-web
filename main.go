package main

import (
	"github.com/goumaaplicacaoweb/routes"
	"net/http"
)

func main() {

	routes.CarregarRotas()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
