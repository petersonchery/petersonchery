package main

import (
	"FirstProject/config"
	"FirstProject/route"
	"net/http"
	"time"
	//"FirstProject/dao"
	//"github.com/lib/pq"
	//"fmt"
)

func main() {

	config.InitDB("user=postgres password=admin dbname=MyBase sslmode=disable host=192.168.10.146 port=5432")

	router := route.GetRouter()
	server := &http.Server{Handler: router, Addr: ":8080", ReadTimeout: time.Second * 10, WriteTimeout: time.Second * 10}
	server.ListenAndServe()

}
