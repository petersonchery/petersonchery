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

// CreateLivre godoc
// @Summary Add a book
// @Description Add a book
// @Tags Livre
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param Livre body model.Livre true "project"
// @Success 200 {object}  model.Livre
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/livres [post]
func CreateLivre(w http.ResponseWriter, r *http.Request) {
	var liv model.Livre
	err := readRequestBody(r.Body, &liv)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}

	err = service.CreateLivre(&liv)
	serverResponse(w, err, liv)

}

// GetLivreById godoc
// @Summary Get a book by ID
// @Description Get a book by ID
// @Tags Livre
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_Livre path int true "ID Livre"
// @Success 200 {object}  model.Livre
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/livres/{id} [get]
func GetLivreById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide"), nil)
		return
	}
	var livre = &model.Livre{Id_livre: id}
	err = service.GetLivreById(livre)
	serverResponse(w, err, livre)
}

// DeleteLivreById godoc
// @Summary Delete a book by ID
// @Description Delete a book by ID
// @Tags Livre
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_emprunt path int true "ID emprunt"
// @Success 200 {object}  model.Emprunt
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/livres/sup/{id} [delete]
func DeleteLivreById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("Id non valide"), nil)
		return
	}
	var livre = &model.Livre{Id_livre: id}
	err = service.DeleteLivreById(livre)
	serverResponse(w, err, livre)
}

// UpdateEmpruntById godoc
// @Summary Update an emprunt by his ID
// @Description Update an emprunt by his ID
// @Tags Livre
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_livre path int true "ID livre"
// @Param Livre body model.Livre true "project"
// @Success 200 {object}  model.Livre
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/livres/up/{id} [put]
func UpdateLivreById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		fmt.Println("erreur dans la conversion")
		return
	}
	var l = &model.Livre{}
	err = readRequestBody(r.Body, l)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur "), nil)
		return
	}
	l.Id_livre = id
	err = service.UpdateLivreById(l)
	serverResponse(w, err, l)
}

// ListerLivre godoc
// @Summary Lister tous les livres
// @Description Lister tous les livres
// @Tags Livres
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Livres
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/livres/lis [get]
func ListerLivre(w http.ResponseWriter, r *http.Request) {

	// var livre = &model.Livre{}
	// biblios, err := service.ListerBibliotheque(biblio)

	// for _, bib := range biblios {
	// 	fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	// }
	// serverResponse(w, err, biblios)

}
