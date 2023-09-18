package rest

import (
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBibliothequeById godoc
// @Summary Get a library By ID
// @Description Get a library By ID
// @Tags Bibliotheque
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_bibliotheque path int true "ID bibliotheque"
// @Success 200 {object}  model.Bibliotheque
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/bibliotheques/{id} [get]
func GetBibliothequeById(w http.ResponseWriter, r *http.Request) {
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
	var biblio = &model.Bibliotheque{Id_bibliotheque: id}
	err = service.GetBibliothequeById(biblio)
	serverResponse(w, err, biblio)

}

// CreateBibliotheque godoc
// @Summary Add a library
// @Description Add a library
// @Tags Bibliotheque
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param Bibliotheque body model.Bibliotheque true "project"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/bibliotheques [post]
func CreateBibliotheque(w http.ResponseWriter, r *http.Request) {
	var b model.Bibliotheque
	err := readRequestBody(r.Body, &b)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}

	fmt.Println(b, "avnat rest")
	err = service.CreateBibliotheque(&b)
	serverResponse(w, err, b)
	fmt.Println(b, "apres rest")
}

// DeleteBibliotheque godoc
// @Summary Delete a library By ID
// @Description Delete a library By ID
// @Tags Bibliotheque
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_bibliotheque path int true "ID bibliotheque"
// @Success 200 {object}  model.Bibliotheque
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/bibliotheques/sup/{id} [delete]
func DeleteBibliotheque(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]

	val, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide "), nil)
		return
	}

	var b = &model.Bibliotheque{Id_bibliotheque: val}
	err = service.DeleteBibliothequeById(b)
	serverResponse(w, err, b)

}

// ListerBibliotheque godoc
// @Summary Lister toutes les Bibliotheques
// @Description Lister toutes les Bibliotheques
// @Tags Bibliotheque
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Bibliotheque
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/bibliotheques/lis [get]
func ListerBibliotheque(w http.ResponseWriter, r *http.Request) {

	var biblio = &model.Bibliotheque{}
	biblios, err := service.ListerBibliotheque(biblio)

	for _, bib := range biblios {
		fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	}
	serverResponse(w, err, biblios)

}

// UpdateBibliotheque godoc
// @Summary Update a library By ID
// @Description Update library By Id
// @Tags Bibliotheque
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_bibliotheque path int true "ID bibliotheque"
// @Param Bibliotheque body model.Bibliotheque true "project"
// @Success 200 {object}  model.Bibliotheque
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/bibliotheques/up/{id} [put]
func UpdateBibliotheque(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRep := vars["id"]
	id, err := strconv.Atoi(idRep)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
		return
	}
	var b = &model.Bibliotheque{}

	err = readRequestBody(r.Body, b)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur dans rest"), nil)
		return
	}
	b.Id_bibliotheque = id
	err = service.UpdateBibliotheque(b)
	serverResponse(w, err, b)
	fmt.Println("update dans res", b)
}

func serverResponse(w http.ResponseWriter, err error, data any) {

	encoder := json.NewEncoder(w)
	if err != nil {
		encoder.Encode(map[string]any{"erreur": err.Error()})
		utils.LogWriter()
		return
	}
	encoder.Encode(data)
	fmt.Println("serverResponse ok")

}

func readRequestBody(r io.Reader, dst any) error {
	b, err := io.ReadAll(r)
	if err != nil {
		utils.LogWriter()
		return fmt.Errorf("donne non valide")
	}

	err = json.Unmarshal(b, &dst)
	if err != nil {
		utils.LogWriter()
		return fmt.Errorf("donne non valide")
	}
	return nil
}
