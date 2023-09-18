package rest

import (
	"FirstProject/controller"
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateTenant godoc
// @Summary Add a tenant
// @Description Add a tenant
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param Tenant body model.Emprunt true "project"
// @Success 200 {object}  model.Tenant
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/tenants [post]
func CreateTenant(w http.ResponseWriter, r *http.Request) {
	var ten model.Tenant
	err := readRequestBody(r.Body, &ten)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}

	err = service.CreateTenant(&ten)
	serverResponse(w, err, ten)
	fmt.Println("rest pass")

}

// GetTenantById godoc
// @Summary Get a tenant by his ID
// @Description Get a tenant by his ID
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_tenant path int true "ID tenant"
// @Success 200 {object}  model.Tenant
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/tenants/{id} [get]
func GetTenantById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("id non valide"), nil)
		return
	}
	var ten = &model.Tenant{Id_Tenant: id}
	err = service.GetTenantById(ten)
	serverResponse(w, err, ten)
}

// DeleteTenantById godoc
// @Summary Delete a tenant by his ID
// @Description Delete a tenant by his ID
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_tenant path int true "ID tenant"
// @Success 200 {object}  model.Tenant
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/tenants/sup/{id} [delete]
func DeleteTenantById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("Id non valide"), nil)
		return
	}
	var ten = &model.Tenant{Id_Tenant: id}
	err = service.DeleteTenantById(ten)
	serverResponse(w, err, ten)
}

// UpdateTenantById godoc
// @Summary Update a tenant by his ID
// @Description Update a tenant by his ID
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_tenant path int true "ID tenant"
// @Param Tenant body model.Tenant true "project"
// @Success 200 {object}  model.Tenant
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/tenants/up/{id} [put]
func UpdateTenantById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		fmt.Println("y a erreur")
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
		return
	}
	t := &model.Tenant{}
	err = readRequestBody(r.Body, t)
	t.Id_Tenant = id
	err = controller.UpdateTenant(t)
	serverResponse(w, err, t)

}

// ListerTenant godoc
// @Summary Lister tous les tenants
// @Description Lister tous les tenants
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Tenant
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/tenants/lis [get]
func ListerTenant(w http.ResponseWriter, r *http.Request) {

	var biblio = &model.Bibliotheque{}
	biblios, err := service.ListerBibliotheque(biblio)

	for _, bib := range biblios {
		fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	}
	serverResponse(w, err, biblios)

}
