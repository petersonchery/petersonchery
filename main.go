package main

import (
	//"FirstProject/controller"
	"FirstProject/dao"
	"net/http"
	"time"
	//"github.com/lib/pq"
	//"fmt"
)

func main() {

	dao.Connect()

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	router := http.NewServeMux()

	server.ListenAndServe()

//	// route.Welcome()
	//route.Menu()
	//rest.GestionDeRoute2()

	//
	//model.Livre = append(model.Liv, controller.Ajouter_livre())

	//.Print(model.Liv)

}
