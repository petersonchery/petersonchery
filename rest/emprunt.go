package rest

import (
	"FirstProject/controller"
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

// CreateEmprunt godoc
// @Summary Ajouter un Emprunt pour un client
// @Description Ajouter un Emprunt pour un client
// @Tags Emprunt
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_emprunt path int true "ID emprunt"
// @Param Emprunt body model.Emprunt true "project"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/emprunts [post]
func CreateEmprunt(w http.ResponseWriter, r *http.Request) {
	var b model.Emprunt
	err := readRequestBody(r.Body, &b)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}
	err = service.CreateEmprunt(&b)
	serverResponse(w, err, b)

}

// GetEmpruntById godoc
// @Summary Get an emprunt by his ID
// @Description Get an emprunt by his ID
// @Tags Emprunt
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_emprunt path int true "ID emprunt"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/emprunts/{id} [get]
func GetEmpruntById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]

	val, err := strconv.Atoi(idReq)

	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur du cote du serveur pour la requette emprunt"), nil)
		return
	}
	var emp = &model.Emprunt{Id_emprunt: val}
	err = service.GetEmpruntById(emp)

	serverResponse(w, err, emp)

}

// DeleteEmpruntById godoc
// @Summary Delete  emprunt by ID
// @Description Delete emprunt by ID
// @Tags Emprunt
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_emprunt path int true "ID emprunt"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/emprunts/sup/{id} [delete]
func DeleteEmpruntById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]

	val, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide "), nil)
		return
	}

	var e = &model.Emprunt{Id_emprunt: val}
	err = service.DeleteEmpruntById(e)
	serverResponse(w, err, e)

}

// UpdateEmpruntById godoc
// @Summary Update an emprunt by his ID
// @Description Update an emprunt by his ID
// @Tags Emprunt
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_emprunt path int true "ID emprunt"
// @Param Emprunt body model.Emprunt true "project"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/emprunts/up/{id} [put]
func UpdateEmpruntById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		fmt.Println("erreur dans la conversion")
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
	}
	e := &model.Emprunt{}
	err = readRequestBody(r.Body, e)

	e.Id_emprunt = id
	err = controller.UpdateEmpruntById(e)
	serverResponse(w, err, e)
}

// ListerEmprunt godoc
// @Summary Lister tous les emprunts
// @Description Lister tous les Bibliotheques
// @Tags Emprunt
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/emprunts/lis [get]
func ListerEmprunt(w http.ResponseWriter, r *http.Request) {

	var biblio = &model.Bibliotheque{}
	biblios, err := service.ListerBibliotheque(biblio)

	for _, bib := range biblios {
		fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	}
	serverResponse(w, err, biblios)

}
