package rest

import (
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateClient godoc
// @Summary Add a client
// @Description Add a client
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param Client body model.Client true "project"
// @Success 200 {object}  model.Client
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/clients [post]
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var b model.Client
	err := readRequestBody(r.Body, &b)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}
	fmt.Println(b, " client avant rest")
	err = service.CreateClient(&b)
	serverResponse(w, err, b)
	fmt.Println(b, "client apres rest")
}

// GetClientById godoc
// @Summary Get a client by ID
// @Description Get a client by ID
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_client path int true "ID client"
// @Success 200 {object}  model.Client
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/clients/{id} [get]
func GetClientById(w http.ResponseWriter, r *http.Request) {
	// /api/v1/bibliotheques/{id}
	// /api/v1/bibliotheques/15
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide"), nil)
		return
	}
	var client = &model.Client{Id_client: id}
	err = service.GetClientById(client)
	serverResponse(w, err, client)

}

// DeleteClientById godoc
// @Summary Delete a Client By ID
// @Description Delete a Client By ID
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_client path int true "ID client"
// @Success 200 {object}  model.Client
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/clients/sup/{id} [delete]
func DeleteClientById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]

	val, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide "), nil)
		return
	}

	var c = &model.Client{Id_client: val}
	err = service.DeleteClientById(c)
	serverResponse(w, err, c)

}

// UpdateClientById godoc
// @Summary UpdateClient
// @Description Update a Client By ID
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_client path int true "ID client"
// @Param Client body model.Client true "project"
// @Success 200 {object}  model.Client
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/clients/up/{id} [put]
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRep := vars["id"]
	id, err := strconv.Atoi(idRep)
	if err != nil {
		utils.LogWriter()
		fmt.Println("erreur")
		return
	}
	var c = &model.Client{}
	err = readRequestBody(r.Body, c)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
		return
	}
	c.Id_client = id
	err = service.UpdateClient(c)
	serverResponse(w, err, c)
	fmt.Println("update dans res", c)
}

// ListerClient godoc
// @Summary Lister tous les Clients
// @Description Lister tous les Clients
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Client
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/clients/lis [get]
func ListerClient(w http.ResponseWriter, r *http.Request) {

	// var client = &model.Client{}
	// biblios, err := service.ListerBibliotheque(client)

	// for _, bib := range biblios {
	// 	fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	// }
	// serverResponse(w, err, biblios)

}
