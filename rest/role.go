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

// CreateRole godoc
// @Summary Ajouter role
// @Description Ajouter role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param Role body model.Role true "project"
// @Success 200 {object}  model.Role
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/roles [post]
func CreateRole(w http.ResponseWriter, r *http.Request) {
	var rol model.Role
	err := readRequestBody(r.Body, &rol)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}

	err = service.CreateRole(&rol)
	serverResponse(w, err, rol)
	fmt.Println("rest pass")

}

// GetRoleById godoc
// @Summary Update a role by his ID
// @Description Update a role by his ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_role path int true "ID role"
// @Success 200 {object}  model.Role
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/roles/{id} [get]
func GetRoleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide"), nil)
		return
	}
	var rol = &model.Role{Id_role: id}
	err = service.GetRoleById(rol)
	serverResponse(w, err, rol)
}

// DeleteRoleById godoc
// @Summary Delete role by ID
// @Description Delete role by ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_role path int true "ID role"
// @Success 200 {object}  model.Role
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/roles/sup/{id} [delete]
func DeleteRoleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("Id non valide"), nil)
		return
	}
	var role = &model.Role{Id_role: id}
	err = service.DeleteRoleById(role)
	serverResponse(w, err, role)
}

// UpdateRoleById godoc
// @Summary Update role by ID
// @Description Update role by ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_role path int true "ID role"
// @Param Role body model.Role true "project"
// @Success 200 {object}  model.Role
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/roles/up/{id} [put]
func UpdateRoleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		fmt.Println("erreur dans la conversion")
		return
	}
	var role = &model.Role{}
	err = readRequestBody(r.Body, role)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("y a erreur "), nil)
		return
	}
	role.Id_role = id
	err = service.UpdateRoleById(role)
	serverResponse(w, err, role)
}

// ListerRole godoc
// @Summary Lister tous les roles
// @Description Lister tous les roles
// @Tags Role
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Roles
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/roles/lis [get]
func ListerRole(w http.ResponseWriter, r *http.Request) {

	// var biblio = &model.Bibliotheque{}
	// biblios, err := service.ListerBibliotheque(biblio)

	// for _, bib := range biblios {
	// 	fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	// }
	// serverResponse(w, err, biblios)

}
