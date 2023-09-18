package main

import (
	"FirstProject/config"
	_ "FirstProject/main/docs"
	"FirstProject/route"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @Gestion de bibliotheque API
// @version 1.0
// @description This is a document for Tainosystems
// @termsOfService http://tainosystems.com
// @contact.name Peterson Chery
// @contact.email Petersonchery33@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 192.168.10.146:8080
// @BasePath /

func main() {

	config.InitDB("user=postgres password=admin dbname=MyBase sslmode=disable host=192.168.10.146 port=5432")

	router := route.InializeRouter()
	router.Methods("GET").PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	server := &http.Server{Handler: router, Addr: ":8080", ReadTimeout: time.Second * 10, WriteTimeout: time.Second * 10}
	server.ListenAndServe()

}
